package all

import (
	// The following are necessary as they register handlers in their init functions.

	// Mandatory features. Can't remove unless there are replacements.
	_ "github.com/edwardmelvin/quick-core/app/dispatcher"
	_ "github.com/edwardmelvin/quick-core/app/proxyman/inbound"
	_ "github.com/edwardmelvin/quick-core/app/proxyman/outbound"

	// Default commander and all its services. This is an optional feature.
	_ "github.com/edwardmelvin/quick-core/app/commander"
	_ "github.com/edwardmelvin/quick-core/app/log/command"
	_ "github.com/edwardmelvin/quick-core/app/proxyman/command"
	_ "github.com/edwardmelvin/quick-core/app/stats/command"

	// Developer preview services
	_ "github.com/edwardmelvin/quick-core/app/observatory/command"

	// Other optional features.
	_ "github.com/edwardmelvin/quick-core/app/dns"
	_ "github.com/edwardmelvin/quick-core/app/dns/fakedns"
	_ "github.com/edwardmelvin/quick-core/app/log"
	_ "github.com/edwardmelvin/quick-core/app/metrics"
	_ "github.com/edwardmelvin/quick-core/app/policy"
	_ "github.com/edwardmelvin/quick-core/app/reverse"
	_ "github.com/edwardmelvin/quick-core/app/router"
	_ "github.com/edwardmelvin/quick-core/app/stats"

	// Fix dependency cycle caused by core import in internet package
	_ "github.com/edwardmelvin/quick-core/transport/internet/tagged/taggedimpl"

	// Developer preview features
	_ "github.com/edwardmelvin/quick-core/app/observatory"

	// Inbound and outbound proxies.
	_ "github.com/edwardmelvin/quick-core/proxy/blackhole"
	_ "github.com/edwardmelvin/quick-core/proxy/dns"
	_ "github.com/edwardmelvin/quick-core/proxy/dokodemo"
	_ "github.com/edwardmelvin/quick-core/proxy/freedom"
	_ "github.com/edwardmelvin/quick-core/proxy/http"
	_ "github.com/edwardmelvin/quick-core/proxy/loopback"
	_ "github.com/edwardmelvin/quick-core/proxy/shadowsocks"
	_ "github.com/edwardmelvin/quick-core/proxy/socks"
	_ "github.com/edwardmelvin/quick-core/proxy/trojan"
	_ "github.com/edwardmelvin/quick-core/proxy/vless/inbound"
	_ "github.com/edwardmelvin/quick-core/proxy/vless/outbound"
	_ "github.com/edwardmelvin/quick-core/proxy/vmess/inbound"
	_ "github.com/edwardmelvin/quick-core/proxy/vmess/outbound"
	_ "github.com/edwardmelvin/quick-core/proxy/wireguard"

	// Transports
	_ "github.com/edwardmelvin/quick-core/transport/internet/domainsocket"
	_ "github.com/edwardmelvin/quick-core/transport/internet/grpc"
	_ "github.com/edwardmelvin/quick-core/transport/internet/http"
	_ "github.com/edwardmelvin/quick-core/transport/internet/kcp"
	_ "github.com/edwardmelvin/quick-core/transport/internet/quic"
	_ "github.com/edwardmelvin/quick-core/transport/internet/reality"
	_ "github.com/edwardmelvin/quick-core/transport/internet/tcp"
	_ "github.com/edwardmelvin/quick-core/transport/internet/tls"
	_ "github.com/edwardmelvin/quick-core/transport/internet/udp"
	_ "github.com/edwardmelvin/quick-core/transport/internet/websocket"

	// Transport headers
	_ "github.com/edwardmelvin/quick-core/transport/internet/headers/http"
	_ "github.com/edwardmelvin/quick-core/transport/internet/headers/noop"
	_ "github.com/edwardmelvin/quick-core/transport/internet/headers/srtp"
	_ "github.com/edwardmelvin/quick-core/transport/internet/headers/tls"
	_ "github.com/edwardmelvin/quick-core/transport/internet/headers/utp"
	_ "github.com/edwardmelvin/quick-core/transport/internet/headers/wechat"
	_ "github.com/edwardmelvin/quick-core/transport/internet/headers/wireguard"

	// JSON & TOML & YAML
	_ "github.com/edwardmelvin/quick-core/main/json"
	_ "github.com/edwardmelvin/quick-core/main/toml"
	_ "github.com/edwardmelvin/quick-core/main/yaml"

	// Load config from file or http(s)
	_ "github.com/edwardmelvin/quick-core/main/confloader/external"

	// Commands
	_ "github.com/edwardmelvin/quick-core/main/commands/all"
)
