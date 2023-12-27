package tcp

import (
	"github.com/edwardmelvin/quick-core/common"
	"github.com/edwardmelvin/quick-core/transport/internet"
)

const protocolName = "tcp"

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
