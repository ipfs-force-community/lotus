package proof_client

import (
	"context"
	"encoding/json"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/chain/gen"
	"github.com/filecoin-project/lotus/node/modules/dtypes"
	proof2 "github.com/filecoin-project/specs-actors/v2/actors/runtime/proof"
	"github.com/google/uuid"
	"github.com/ipfs-force-community/venus-gateway/proofevent"
	"github.com/ipfs-force-community/venus-gateway/types"
	logging "github.com/ipfs/go-log/v2"
	"golang.org/x/xerrors"
	"time"
)

var log = logging.Logger("proof_event")

type ProofEvent struct {
	epp    gen.WinningPoStProver
	client *ProofEventClient
	mAddr  dtypes.MinerAddress
}

func (e *ProofEvent) listenProofRequest(ctx context.Context) {
	for {
		if err := e.listenProofRequestOnce(ctx); err != nil {
			log.Errorf("listen head changes errored: %s", err)
		} else {
			log.Warn("listenHeadChanges quit")
		}
		select {
		case <-time.After(time.Second):
		case <-ctx.Done():
			log.Warnf("not restarting listenHeadChanges: context error: %s", ctx.Err())
			return
		}

		log.Info("restarting listenHeadChanges")
	}
}

func (e *ProofEvent) listenProofRequestOnce(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	policy := &proofevent.ProofRegisterPolicy{
		MinerAddress: address.Address(e.mAddr),
	}
	proofEventCh, err := e.client.ListenProofEvent(ctx, policy)
	if err != nil {
		// Retry is handled by caller
		return xerrors.Errorf("listenHeadChanges ChainNotify call failed: %w", err)
	}

	for proofEvent := range proofEventCh {
		switch proofEvent.Method {
		case "InitConnect":
			req := types.ConnectedCompleted{}
			err := json.Unmarshal(proofEvent.Payload, &req)
			if err != nil {
				return xerrors.Errorf("odd error in connect %v", err)
			}
			log.Infof("success to connect with proof %s", req.ChannelId)
		case "ComputeProof":
			req := types.ComputeProofRequest{}
			err := json.Unmarshal(proofEvent.Payload, &req)
			if err != nil {
				e.client.ResponseProofEvent(ctx, &types.ResponseEvent{
					Id:      proofEvent.Id,
					Payload: nil,
					Error:   err.Error(),
				})
				continue
			}
			e.processComputeProof(ctx, proofEvent.Id, req)
		default:
			log.Errorf("unexpect proof event type %s", proofEvent.Method)
		}
	}

	return nil
}

func (e *ProofEvent) processComputeProof(ctx context.Context, reqId uuid.UUID, req types.ComputeProofRequest) {
	sectorInfos := make([]proof2.SectorInfo, len(req.SectorInfos))
	for index, sinfo := range req.SectorInfos {
		sectorInfos[index] = proof2.SectorInfo{
			SealProof:    sinfo.SealProof,
			SectorNumber: sinfo.SectorNumber,
			SealedCID:    sinfo.SealedCID,
		}
	}
	proof, err := e.epp.ComputeProof(ctx, sectorInfos, req.Rand)
	if err != nil {
		e.client.ResponseProofEvent(ctx, &types.ResponseEvent{
			Id:      reqId,
			Payload: nil,
			Error:   err.Error(),
		})
		return
	}

	proofBytes, err := json.Marshal(proof)
	if err != nil {
		e.client.ResponseProofEvent(ctx, &types.ResponseEvent{
			Id:      reqId,
			Payload: nil,
			Error:   err.Error(),
		})
		return
	}

	err = e.client.ResponseProofEvent(ctx, &types.ResponseEvent{
		Id:      reqId,
		Payload: proofBytes,
		Error:   "",
	})
	if err != nil {
		log.Errorf("response proof event %s failed", reqId)
	}
}
