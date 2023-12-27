package command_test

import (
	"context"
	"testing"

	"github.com/edwardmelvin/quick-core/app/dispatcher"
	"github.com/edwardmelvin/quick-core/app/log"
	. "github.com/edwardmelvin/quick-core/app/log/command"
	"github.com/edwardmelvin/quick-core/app/proxyman"
	_ "github.com/edwardmelvin/quick-core/app/proxyman/inbound"
	_ "github.com/edwardmelvin/quick-core/app/proxyman/outbound"
	"github.com/edwardmelvin/quick-core/common"
	"github.com/edwardmelvin/quick-core/common/serial"
	"github.com/edwardmelvin/quick-core/core"
)

func TestLoggerRestart(t *testing.T) {
	v, err := core.New(&core.Config{
		App: []*serial.TypedMessage{
			serial.ToTypedMessage(&log.Config{}),
			serial.ToTypedMessage(&dispatcher.Config{}),
			serial.ToTypedMessage(&proxyman.InboundConfig{}),
			serial.ToTypedMessage(&proxyman.OutboundConfig{}),
		},
	})
	common.Must(err)
	common.Must(v.Start())

	server := &LoggerServer{
		V: v,
	}
	common.Must2(server.RestartLogger(context.Background(), &RestartLoggerRequest{}))
}
