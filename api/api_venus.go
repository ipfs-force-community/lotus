package api

import (
	"context"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/filecoin-project/lotus/chain/types"
)

type VenusAPI interface {
	// MethodGroup: Venus

	MpoolSelects(context.Context, types.TipSetKey, []float64) ([][]*types.SignedMessage, error) //perm:read

	MpoolPublishMessage(ctx context.Context, smsg *types.SignedMessage) error //perm:write

	MpoolPublishByAddr(context.Context, address.Address) error //perm:write

	GasBatchEstimateMessageGas(ctx context.Context, estimateMessages []*EstimateMessage, fromNonce uint64, tsk types.TipSetKey) ([]*EstimateResult, error) //perm:read

	// ChainGetRandomnessFromTickets is used to sample the chain for randomness.
	ChainGetRandomnessFromTickets(ctx context.Context, tsk types.TipSetKey, personalization crypto.DomainSeparationTag, randEpoch abi.ChainEpoch, entropy []byte) (abi.Randomness, error) //perm:read

	// ChainGetRandomnessFromBeacon is used to sample the beacon for randomness.
	ChainGetRandomnessFromBeacon(ctx context.Context, tsk types.TipSetKey, personalization crypto.DomainSeparationTag, randEpoch abi.ChainEpoch, entropy []byte) (abi.Randomness, error) //perm:read

	// BeaconGetEntry returns the beacon entry for the given filecoin epoch. If
	// the entry has not yet been produced, the call will block until the entry
	// becomes available
	BeaconGetEntry(ctx context.Context, epoch abi.ChainEpoch) (*types.BeaconEntry, error) //perm:read
}
