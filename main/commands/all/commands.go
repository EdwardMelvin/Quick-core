package all

import (
	"github.com/edwardmelvin/quick-core/main/commands/all/api"
	"github.com/edwardmelvin/quick-core/main/commands/all/tls"
	"github.com/edwardmelvin/quick-core/main/commands/base"
)

// go:generate go run github.com/edwardmelvin/quick-core/common/errors/errorgen

func init() {
	base.RootCommand.Commands = append(
		base.RootCommand.Commands,
		api.CmdAPI,
		// cmdConvert,
		tls.CmdTLS,
		cmdUUID,
		cmdX25519,
	)
}
