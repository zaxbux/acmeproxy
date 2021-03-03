module github.com/zaxbux/acmeproxy

go 1.15

require (
	github.com/abbot/go-http-auth v0.4.0
	github.com/cloudflare/cloudflare-go v0.13.2
	github.com/codeskyblue/realip v0.1.0
	github.com/go-acme/lego/v4 v4.2.0
	github.com/mattn/go-colorable v0.1.8 // indirect
	github.com/mgutz/ansi v0.0.0-20200706080929-d51e80ef957d // indirect
	github.com/mholt/certmagic v0.9.3
	github.com/orange-cloudfoundry/ipfiltering v0.0.0-20170823192147-f48f1b767f82
	github.com/sirupsen/logrus v1.4.2
	github.com/stretchr/testify v1.6.1
	github.com/x-cray/logrus-prefixed-formatter v0.5.2
	golang.org/x/net v0.0.0-20200822124328-c89045814202
	gopkg.in/urfave/cli.v1 v1.20.0
)

replace github.com/go-acme/lego/v4/providers/dns/cloudflare => ./lego/providers/dns/cloudflare
