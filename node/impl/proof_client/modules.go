package proof_client

import (
	"context"
	"github.com/filecoin-project/lotus/chain/gen"
	"github.com/filecoin-project/lotus/node/modules/dtypes"
)

func StartProofEvent(ctx context.Context, prover gen.WinningPoStProver, client *ProofEventClient, mAddr dtypes.MinerAddress) error {
	proofEvent := &ProofEvent{epp: prover, client: client, mAddr: mAddr}
	go proofEvent.listenProofRequest(ctx)
	return nil
}
