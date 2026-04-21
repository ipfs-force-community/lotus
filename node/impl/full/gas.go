package full

import (
	"context"
	"fmt"
	stdbig "math/big"
	"os"

	"go.uber.org/fx"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/go-state-types/exitcode"

	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/build/buildconstants"
	lbuiltin "github.com/filecoin-project/lotus/chain/actors/builtin"
	"github.com/filecoin-project/lotus/chain/messagepool"
	"github.com/filecoin-project/lotus/chain/stmgr"
	"github.com/filecoin-project/lotus/chain/store"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/node/impl/gasutils"
	"github.com/filecoin-project/lotus/node/modules/dtypes"
)

type GasModuleAPI interface {
	GasEstimateMessageGas(ctx context.Context, msg *types.Message, spec *api.MessageSendSpec, tsk types.TipSetKey) (*types.Message, error)
	GasBatchEstimateMessageGas(ctx context.Context, estimateMessages []*api.EstimateMessage, fromNonce uint64, tsk types.TipSetKey) ([]*api.EstimateResult, error)
}

var _ GasModuleAPI = *new(api.FullNode)

// GasModule provides a default implementation of GasModuleAPI.
// It can be swapped out with another implementation through Dependency
// Injection (for example with a thin RPC client).
type GasModule struct {
	fx.In
	Stmgr     *stmgr.StateManager
	Chain     *store.ChainStore
	Mpool     *messagepool.MessagePool
	GetMaxFee dtypes.DefaultMaxFeeFunc

	PriceCache *gasutils.GasPriceCache
}

var _ GasModuleAPI = (*GasModule)(nil)

type GasAPI struct {
	fx.In

	GasModuleAPI

	Stmgr *stmgr.StateManager
	Chain *store.ChainStore
	Mpool *messagepool.MessagePool

	PriceCache *gasutils.GasPriceCache
}

func (a *GasAPI) GasEstimateFeeCap(
	ctx context.Context,
	msg *types.Message,
	maxqueueblks int64,
	tsk types.TipSetKey,
) (types.BigInt, error) {
	return gasutils.GasEstimateFeeCap(ctx, a.Chain, msg, maxqueueblks, tsk)
}

func (m *GasModule) GasEstimateFeeCap(
	ctx context.Context,
	msg *types.Message,
	maxqueueblks int64,
	tsk types.TipSetKey,
) (types.BigInt, error) {
	return gasutils.GasEstimateFeeCap(ctx, m.Chain, msg, maxqueueblks, tsk)
}

func (a *GasAPI) GasEstimateGasPremium(
	ctx context.Context,
	nblocksincl uint64,
	sender address.Address,
	gaslimit int64,
	tsk types.TipSetKey,
) (types.BigInt, error) {
	return gasutils.GasEstimateGasPremium(ctx, a.Chain, a.PriceCache, nblocksincl, tsk)
}

func (m *GasModule) GasEstimateGasPremium(
	ctx context.Context,
	nblocksincl uint64,
	sender address.Address,
	gaslimit int64,
	tsk types.TipSetKey,
) (types.BigInt, error) {
	return gasutils.GasEstimateGasPremium(ctx, m.Chain, m.PriceCache, nblocksincl, tsk)
}

func (a *GasAPI) GasEstimateGasLimit(ctx context.Context, msgIn *types.Message, tsk types.TipSetKey) (int64, error) {
	ts, err := a.Chain.GetTipSetFromKey(ctx, tsk)
	if err != nil {
		return -1, xerrors.Errorf("getting tipset: %w", err)
	}
	return gasutils.GasEstimateGasLimit(ctx, a.Chain, a.Stmgr, a.Mpool, msgIn, ts)
}

func (m *GasModule) GasEstimateGasLimit(ctx context.Context, msgIn *types.Message, tsk types.TipSetKey) (int64, error) {
	ts, err := m.Chain.GetTipSetFromKey(ctx, tsk)
	if err != nil {
		return -1, xerrors.Errorf("getting tipset: %w", err)
	}
	return gasutils.GasEstimateGasLimit(ctx, m.Chain, m.Stmgr, m.Mpool, msgIn, ts)
}
func evalMessageGasLimit(ctx context.Context, smgr *stmgr.StateManager, cstore *store.ChainStore, msgIn *types.Message, priorMsgs []types.ChainMsg, ts *types.TipSet) (int64, error) {
	msg := *msgIn
	msg.GasLimit = buildconstants.BlockGasLimit
	msg.GasFeeCap = types.NewInt(uint64(buildconstants.MinimumBaseFee) + 1)
	msg.GasPremium = types.NewInt(1)

	applyTSMessages := true
	if os.Getenv("LOTUS_SKIP_APPLY_TS_MESSAGE_CALL_WITH_GAS") == "1" {
		applyTSMessages = false
	}

	// Try calling until we find a height with no migration.
	var res *api.InvocResult
	var err error
	for {
		res, err = smgr.CallWithGas(ctx, &msg, priorMsgs, ts, applyTSMessages)
		if err != stmgr.ErrExpensiveFork {
			break
		}
		ts, err = cstore.GetTipSetFromKey(ctx, ts.Parents())
		if err != nil {
			return -1, xerrors.Errorf("getting parent tipset: %w", err)
		}
	}
	if err != nil {
		return -1, xerrors.Errorf("CallWithGas failed: %w", err)
	}
	if res.MsgRct.ExitCode != exitcode.Ok {
		return -1, xerrors.Errorf("message execution failed: exit %s, reason: %s", res.MsgRct.ExitCode, res.Error)
	}

	ret := res.MsgRct.GasUsed

	log.Infow("evalMessageGasLimit CallWithGas Result", "GasUsed", ret, "ExitCode", res.MsgRct.ExitCode)

	transitionalMulti := 1.0
	// Overestimate gas around the upgrade
	if ts.Height() <= buildconstants.UpgradeHyggeHeight && (buildconstants.UpgradeHyggeHeight-ts.Height() <= 20) {
		func() {

			// Bare transfers get about 3x more expensive: https://github.com/filecoin-project/FIPs/blob/master/FIPS/fip-0057.md#product-considerations
			if msgIn.Method == builtin.MethodSend {
				transitionalMulti = 3.0
				return
			}

			st, err := smgr.ParentState(ts)
			if err != nil {
				return
			}
			act, err := st.GetActor(msg.To)
			if err != nil {
				return
			}

			if lbuiltin.IsStorageMinerActor(act.Code) {
				switch msgIn.Method {
				case 3:
					transitionalMulti = 1.92
				case 4:
					transitionalMulti = 1.72
				case 6:
					transitionalMulti = 1.06
				case 7:
					transitionalMulti = 1.2
				case 16:
					transitionalMulti = 1.19
				case 18:
					transitionalMulti = 1.73
				case 23:
					transitionalMulti = 1.73
				case 26:
					transitionalMulti = 1.15
				case 27:
					transitionalMulti = 1.18
				default:
				}
			}
		}()
	}
	ret = (ret * int64(transitionalMulti*1024)) >> 10

	// Special case for PaymentChannel collect, which is deleting actor
	// We ignore errors in this special case since they CAN occur,
	// and we just want to detect existing payment channel actors
	st, err := smgr.ParentState(ts)
	if err == nil {
		act, err := st.GetActor(msg.To)
		if err == nil && lbuiltin.IsPaymentChannelActor(act.Code) && msgIn.Method == builtin.MethodsPaych.Collect {
			// add the refunded gas for DestroyActor back into the gas used
			ret += 76e3
		}
	}

	return ret, nil
}

func (m *GasModule) GasEstimateMessageGas(ctx context.Context, msg *types.Message, spec *api.MessageSendSpec, ts types.TipSetKey) (*types.Message, error) {
	if msg.GasLimit == 0 {
		gasLimit, err := m.GasEstimateGasLimit(ctx, msg, ts)
		if err != nil {
			return nil, err
		}
		msg.GasLimit = int64(float64(gasLimit) * m.Mpool.GetConfig().GasLimitOverestimation)

		// Gas overestimation can cause us to exceed the block gas limit, cap it.
		if msg.GasLimit > buildconstants.BlockGasLimit {
			msg.GasLimit = buildconstants.BlockGasLimit
		}
	}

	if msg.GasPremium == types.EmptyInt || types.BigCmp(msg.GasPremium, types.NewInt(0)) == 0 {
		gasPremium, err := m.GasEstimateGasPremium(ctx, 10, msg.From, msg.GasLimit, ts)
		if err != nil {
			return nil, xerrors.Errorf("estimating gas price: %w", err)
		}
		msg.GasPremium = gasPremium
	}

	if msg.GasFeeCap == types.EmptyInt || types.BigCmp(msg.GasFeeCap, types.NewInt(0)) == 0 {
		feeCap, err := m.GasEstimateFeeCap(ctx, msg, 20, ts)
		if err != nil {
			return nil, xerrors.Errorf("estimating fee cap: %w", err)
		}
		msg.GasFeeCap = feeCap
	}

	messagepool.CapGasFee(m.GetMaxFee, msg, spec)

	return msg, nil
}

func (m *GasModule) GasBatchEstimateMessageGas(ctx context.Context, estimateMessages []*api.EstimateMessage, fromNonce uint64, tsk types.TipSetKey) ([]*api.EstimateResult, error) {
	if len(estimateMessages) == 0 {
		return nil, nil
	}

	ts, err := m.Chain.GetTipSetFromKey(ctx, tsk)
	if err != nil {
		return nil, xerrors.Errorf("getting tipset: %w", err)
	}

	fromA, err := m.Stmgr.ResolveToDeterministicAddress(ctx, estimateMessages[0].Msg.From, ts)
	if err != nil {
		return nil, xerrors.Errorf("getting key address: %w", err)
	}

	pending, ts := m.Mpool.PendingFor(ctx, fromA)
	priorMsgs := make([]types.ChainMsg, 0, len(pending))
	for _, m := range pending {
		priorMsgs = append(priorMsgs, m)
	}

	var estimateResults []*api.EstimateResult
	for _, estimateMessage := range estimateMessages {
		estimateMsg := estimateMessage.Msg
		estimateMsg.Nonce = fromNonce

		log.Debugf("call GasBatchEstimateMessageGas msg %v, spec %v", estimateMsg, estimateMessage.Spec)

		if estimateMsg.GasLimit == 0 {
			gasUsed, err := evalMessageGasLimit(ctx, m.Stmgr, m.Chain, estimateMsg, priorMsgs, ts)
			if err != nil {
				estimateMsg.Nonce = 0
				estimateResults = append(estimateResults, &api.EstimateResult{
					Msg: estimateMsg,
					Err: fmt.Sprintf("estimating gas price: %v", err),
				})
				continue
			}
			gasLimitOverestimation := m.Mpool.GetConfig().GasLimitOverestimation
			if estimateMessage.Spec != nil && estimateMessage.Spec.GasOverEstimation > 0 {
				gasLimitOverestimation = estimateMessage.Spec.GasOverEstimation
			}
			estimateMsg.GasLimit = int64(float64(gasUsed) * gasLimitOverestimation)
		}

		if estimateMsg.GasPremium == types.EmptyInt || types.BigCmp(estimateMsg.GasPremium, types.NewInt(0)) == 0 {
			gasPremium, err := m.GasEstimateGasPremium(ctx, 10, estimateMsg.From, estimateMsg.GasLimit, types.TipSetKey{})
			if err != nil {
				estimateMsg.Nonce = 0
				estimateResults = append(estimateResults, &api.EstimateResult{
					Msg: estimateMsg,
					Err: fmt.Sprintf("estimating gas price: %v", err),
				})
				continue
			}
			if estimateMessage.Spec != nil && estimateMessage.Spec.GasOverPremium > 0 {
				olgGasPremium := gasPremium
				newGasPremium, _ := new(stdbig.Float).Mul(new(stdbig.Float).SetInt(stdbig.NewInt(gasPremium.Int64())), stdbig.NewFloat(estimateMessage.Spec.GasOverPremium)).Int(nil)
				gasPremium = big.NewFromGo(newGasPremium)
				log.Debugf("call GasBatchEstimateMessageGas old premium %v, new premium %v, premium ration %f", olgGasPremium, newGasPremium, estimateMessage.Spec.GasOverPremium)
			}
			estimateMsg.GasPremium = gasPremium
		}

		if estimateMsg.GasFeeCap == types.EmptyInt || types.BigCmp(estimateMsg.GasFeeCap, types.NewInt(0)) == 0 {
			feeCap, err := m.GasEstimateFeeCap(ctx, estimateMsg, 20, types.EmptyTSK)
			if err != nil {
				estimateMsg.Nonce = 0
				estimateResults = append(estimateResults, &api.EstimateResult{
					Msg: estimateMsg,
					Err: fmt.Sprintf("estimating fee cap: %v", err),
				})
				continue
			}
			estimateMsg.GasFeeCap = feeCap
		}

		messagepool.CapGasFee(m.GetMaxFee, estimateMsg, estimateMessage.Spec)

		estimateResults = append(estimateResults, &api.EstimateResult{
			Msg: estimateMsg,
		})
		priorMsgs = append(priorMsgs, estimateMsg)
		fromNonce++
	}
	return estimateResults, nil
}
