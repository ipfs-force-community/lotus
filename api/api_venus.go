package api

import (
	"context"

	"github.com/filecoin-project/go-address"

	"github.com/filecoin-project/lotus/chain/types"
)

type VenusAPI interface {
	// MethodGroup: Venus

	MpoolSelects(context.Context, types.TipSetKey, []float64) ([][]*types.SignedMessage, error) //perm:read

	MpoolPublishMessage(ctx context.Context, smsg *types.SignedMessage) error //perm:write

	MpoolPublishByAddr(context.Context, address.Address) error //perm:write

	GasBatchEstimateMessageGas(ctx context.Context, estimateMessages []*EstimateMessage, fromNonce uint64, tsk types.TipSetKey) ([]*EstimateResult, error) //perm:read
}
