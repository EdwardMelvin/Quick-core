package wireguard

import (
	"context"
	"errors"
	"io"

	"github.com/edwardmelvin/quick-core/common"
	"github.com/edwardmelvin/quick-core/common/buf"
	"github.com/edwardmelvin/quick-core/common/log"
	"github.com/edwardmelvin/quick-core/common/net"
	"github.com/edwardmelvin/quick-core/common/session"
	"github.com/edwardmelvin/quick-core/common/signal"
	"github.com/edwardmelvin/quick-core/common/task"
	"github.com/edwardmelvin/quick-core/core"
	"github.com/edwardmelvin/quick-core/features/dns"
	"github.com/edwardmelvin/quick-core/features/policy"
	"github.com/edwardmelvin/quick-core/features/routing"
	"github.com/edwardmelvin/quick-core/transport/internet/stat"
)

var nullDestination = net.TCPDestination(net.AnyIP, 0)

type Server struct {
	bindServer *netBindServer

	info          routingInfo
	policyManager policy.Manager
}

type routingInfo struct {
	ctx         context.Context
	dispatcher  routing.Dispatcher
	inboundTag  *session.Inbound
	outboundTag *session.Outbound
	contentTag  *session.Content
}

func NewServer(ctx context.Context, conf *DeviceConfig) (*Server, error) {
	v := core.MustFromContext(ctx)

	endpoints, hasIPv4, hasIPv6, err := parseEndpoints(conf)
	if err != nil {
		return nil, err
	}

	server := &Server{
		bindServer: &netBindServer{
			netBind: netBind{
				dns: v.GetFeature(dns.ClientType()).(dns.Client),
				dnsOption: dns.IPOption{
					IPv4Enable: hasIPv4,
					IPv6Enable: hasIPv6,
				},
			},
		},
		policyManager: v.GetFeature(policy.ManagerType()).(policy.Manager),
	}

	tun, err := conf.createTun()(endpoints, int(conf.Mtu), server.forwardConnection)
	if err != nil {
		return nil, err
	}

	if err = tun.BuildDevice(createIPCRequest(conf), server.bindServer); err != nil {
		_ = tun.Close()
		return nil, err
	}

	return server, nil
}

// Network implements proxy.Inbound.
func (*Server) Network() []net.Network {
	return []net.Network{net.Network_UDP}
}

// Process implements proxy.Inbound.
func (s *Server) Process(ctx context.Context, network net.Network, conn stat.Connection, dispatcher routing.Dispatcher) error {
	inbound := session.InboundFromContext(ctx)
	inbound.Name = "wireguard"
	inbound.SetCanSpliceCopy(3)

	s.info = routingInfo{
		ctx:         core.ToBackgroundDetachedContext(ctx),
		dispatcher:  dispatcher,
		inboundTag:  session.InboundFromContext(ctx),
		outboundTag: session.OutboundFromContext(ctx),
		contentTag:  session.ContentFromContext(ctx),
	}

	ep, err := s.bindServer.ParseEndpoint(conn.RemoteAddr().String())
	if err != nil {
		return err
	}

	nep := ep.(*netEndpoint)
	nep.conn = conn

	reader := buf.NewPacketReader(conn)
	for {
		mpayload, err := reader.ReadMultiBuffer()
		if err != nil {
			return err
		}

		for _, payload := range mpayload {
			v, ok := <-s.bindServer.readQueue
			if !ok {
				return nil
			}
			i, err := payload.Read(v.buff)

			v.bytes = i
			v.endpoint = nep
			v.err = err
			v.waiter.Done()
			if err != nil && errors.Is(err, io.EOF) {
				nep.conn = nil
				return nil
			}
		}
	}
}

func (s *Server) forwardConnection(dest net.Destination, conn net.Conn) {
	if s.info.dispatcher == nil {
		newError("unexpected: dispatcher == nil").AtError().WriteToLog()
		return
	}
	defer conn.Close()

	ctx, cancel := context.WithCancel(core.ToBackgroundDetachedContext(s.info.ctx))
	plcy := s.policyManager.ForLevel(0)
	timer := signal.CancelAfterInactivity(ctx, cancel, plcy.Timeouts.ConnectionIdle)

	ctx = log.ContextWithAccessMessage(ctx, &log.AccessMessage{
		From:   nullDestination,
		To:     dest,
		Status: log.AccessAccepted,
		Reason: "",
	})

	if s.info.inboundTag != nil {
		ctx = session.ContextWithInbound(ctx, s.info.inboundTag)
	}
	if s.info.outboundTag != nil {
		ctx = session.ContextWithOutbound(ctx, s.info.outboundTag)
	}
	if s.info.contentTag != nil {
		ctx = session.ContextWithContent(ctx, s.info.contentTag)
	}

	link, err := s.info.dispatcher.Dispatch(ctx, dest)
	if err != nil {
		newError("dispatch connection").Base(err).AtError().WriteToLog()
	}
	defer cancel()

	requestDone := func() error {
		defer timer.SetTimeout(plcy.Timeouts.DownlinkOnly)
		if err := buf.Copy(buf.NewReader(conn), link.Writer, buf.UpdateActivity(timer)); err != nil {
			return newError("failed to transport all TCP request").Base(err)
		}

		return nil
	}

	responseDone := func() error {
		defer timer.SetTimeout(plcy.Timeouts.UplinkOnly)
		if err := buf.Copy(link.Reader, buf.NewWriter(conn), buf.UpdateActivity(timer)); err != nil {
			return newError("failed to transport all TCP response").Base(err)
		}

		return nil
	}

	requestDonePost := task.OnSuccess(requestDone, task.Close(link.Writer))
	if err := task.Run(ctx, requestDonePost, responseDone); err != nil {
		common.Interrupt(link.Reader)
		common.Interrupt(link.Writer)
		newError("connection ends").Base(err).AtDebug().WriteToLog()
		return
	}
}
