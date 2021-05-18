package cli

import (
	"context"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/go-state-types/crypto"
	cliutil "github.com/filecoin-project/lotus/cli/util"
	"github.com/filecoin-project/lotus/node/config"
	"github.com/filecoin-project/venus-wallet/core"
)

type IWalletClient interface {
	WalletList(context.Context) ([]address.Address, error)
	WalletHas(context.Context, address.Address) (bool, error)
	WalletSign(ctx context.Context, signer address.Address, toSign []byte, meta core.MsgMeta) (*crypto.Signature, error)
}

var _ IWalletClient = (*WalletClient)(nil)

type WalletClient struct {
	Internal struct {
		WalletList func(context.Context) ([]address.Address, error)
		WalletHas  func(ctx context.Context, address address.Address) (bool, error)
		WalletSign func(ctx context.Context, signer address.Address, toSign []byte, meta core.MsgMeta) (*crypto.Signature, error)
	}
}

func (walletClient *WalletClient) WalletList(ctx context.Context) ([]address.Address, error) {
	return walletClient.Internal.WalletList(ctx)
}

func (walletClient *WalletClient) WalletHas(ctx context.Context, addr address.Address) (bool, error) {
	return walletClient.Internal.WalletHas(ctx, addr)
}

func (walletClient *WalletClient) WalletSign(ctx context.Context, signer address.Address, toSign []byte, meta core.MsgMeta) (*crypto.Signature, error) {
	return walletClient.Internal.WalletSign(ctx, signer, toSign, meta)
}

func NewWalletClient(cfg *config.WalletConfig) (WalletClient, error) {
	apiInfo := cliutil.APIInfo{
		Addr:  cfg.Url,
		Token: []byte(cfg.Token),
	}

	addr, err := apiInfo.DialArgs()
	if err != nil {
		return WalletClient{}, err
	}

	var res WalletClient
	_, err = jsonrpc.NewMergeClient(context.Background(), addr, "Filecoin", []interface{}{&res.Internal}, apiInfo.AuthHeader())
	return res, err
}
