package proof_client

import (
	"context"
	"github.com/filecoin-project/lotus/chain/gen"
	"github.com/filecoin-project/lotus/node/config"
	"github.com/filecoin-project/lotus/node/modules/dtypes"
	"go.uber.org/fx"
)

func StartProofEvent(prover gen.WinningPoStProver, lc fx.Lifecycle, cfg *config.RegisterProofConfig, mAddr dtypes.MinerAddress) error {
	for _, addr := range cfg.Urls {
		client, err := NewProofEventClient(lc, addr, cfg.Token)
		if err != nil {
			return err
		}
		proofEvent := ProofEvent{
			prover: prover,
			client: client,
			mAddr:  mAddr,
		}
		go proofEvent.listenProofRequest(context.Background())
	}

	return nil
}
