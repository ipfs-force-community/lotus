package main

import (
	"context"
	"github.com/filecoin-project/go-jsonrpc/auth"
	"github.com/filecoin-project/lotus/api/v1api"
	"github.com/filecoin-project/venus-auth/cmd/jwtclient"
)

var _ jwtclient.IJwtAuthClient = (*WrapClient)(nil)

type WrapClient struct {
	a v1api.FullNode
}

func (w *WrapClient) Verify(ctx context.Context, token string) ([]auth.Permission, error) {
	return w.a.AuthVerify(ctx, token)
}
