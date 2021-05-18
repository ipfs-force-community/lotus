package messager

import (
	"github.com/urfave/cli/v2"
)

var (
	MessagerFlagUrl = &cli.StringFlag{
		Name:  "messager-url",
		Usage: "messager url",
	}
	WalletUrlFlag = &cli.StringFlag{
		Name:  "wallet-url",
		Usage: "wallet url",
	}
	WalletTokenFlag = &cli.StringFlag{
		Name:  "wallet-token",
		Usage: "wallet token",
	}
	GatewayUrlFlag = &cli.StringSliceFlag{
		Name:  "gateway-url",
		Usage: "gateway url",
	}
	AuthTokenFlag = &cli.StringFlag{
		Name:  "auth-token",
		Usage: "auth token",
	}
)
