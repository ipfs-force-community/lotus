package messager

import (
	"github.com/urfave/cli/v2"
)

var (
	MessagerFlagUrl = &cli.StringFlag{
		Name:  "messager-url",
		Usage: "messager url",
	}
	MessagerFlagToken = &cli.StringFlag{
		Name:  "messager-token",
		Usage: "messager token",
	}
	WalletUrlFlagWallet = &cli.StringFlag{
		Name:  "wallet-url",
		Usage: "wallet url",
	}
	WalletTokenFlagWallet = &cli.StringFlag{
		Name:  "wallet-token",
		Usage: "wallet token",
	}
)
