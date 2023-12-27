package core_test

import (
	"testing"

	"github.com/edwardmelvin/quick-core/app/dispatcher"
	"github.com/edwardmelvin/quick-core/app/proxyman"
	"github.com/edwardmelvin/quick-core/common"
	"github.com/edwardmelvin/quick-core/common/net"
	"github.com/edwardmelvin/quick-core/common/protocol"
	"github.com/edwardmelvin/quick-core/common/serial"
	"github.com/edwardmelvin/quick-core/common/uuid"
	. "github.com/edwardmelvin/quick-core/core"
	"github.com/edwardmelvin/quick-core/features/dns"
	"github.com/edwardmelvin/quick-core/features/dns/localdns"
	_ "github.com/edwardmelvin/quick-core/main/distro/all"
	"github.com/edwardmelvin/quick-core/proxy/dokodemo"
	"github.com/edwardmelvin/quick-core/proxy/vmess"
	"github.com/edwardmelvin/quick-core/proxy/vmess/outbound"
	"github.com/edwardmelvin/quick-core/testing/servers/tcp"
	"google.golang.org/protobuf/proto"
)

func TestXrayDependency(t *testing.T) {
	instance := new(Instance)

	wait := make(chan bool, 1)
	instance.RequireFeatures(func(d dns.Client) {
		if d == nil {
			t.Error("expected dns client fulfilled, but actually nil")
		}
		wait <- true
	})
	instance.AddFeature(localdns.New())
	<-wait
}

func TestXrayClose(t *testing.T) {
	port := tcp.PickPort()

	userID := uuid.New()
	config := &Config{
		App: []*serial.TypedMessage{
			serial.ToTypedMessage(&dispatcher.Config{}),
			serial.ToTypedMessage(&proxyman.InboundConfig{}),
			serial.ToTypedMessage(&proxyman.OutboundConfig{}),
		},
		Inbound: []*InboundHandlerConfig{
			{
				ReceiverSettings: serial.ToTypedMessage(&proxyman.ReceiverConfig{
					PortList: &net.PortList{
						Range: []*net.PortRange{net.SinglePortRange(port)},
					},
					Listen: net.NewIPOrDomain(net.LocalHostIP),
				}),
				ProxySettings: serial.ToTypedMessage(&dokodemo.Config{
					Address: net.NewIPOrDomain(net.LocalHostIP),
					Port:    uint32(0),
					NetworkList: &net.NetworkList{
						Network: []net.Network{net.Network_TCP},
					},
				}),
			},
		},
		Outbound: []*OutboundHandlerConfig{
			{
				ProxySettings: serial.ToTypedMessage(&outbound.Config{
					Receiver: []*protocol.ServerEndpoint{
						{
							Address: net.NewIPOrDomain(net.LocalHostIP),
							Port:    uint32(0),
							User: []*protocol.User{
								{
									Account: serial.ToTypedMessage(&vmess.Account{
										Id: userID.String(),
									}),
								},
							},
						},
					},
				}),
			},
		},
	}

	cfgBytes, err := proto.Marshal(config)
	common.Must(err)

	server, err := StartInstance("protobuf", cfgBytes)
	common.Must(err)
	server.Close()
}
