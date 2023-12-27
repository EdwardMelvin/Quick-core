package udp

import (
	"github.com/edwardmelvin/quick-core/common"
	"github.com/edwardmelvin/quick-core/transport/internet"
)

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
