package gateway

import (
	"context"
	"fmt"
	"time"

	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
	"go.opencensus.io/stats"
	"golang.org/x/time/rate"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-bitfield"
	"github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/filecoin-project/go-state-types/dline"
	"github.com/filecoin-project/go-state-types/network"

	"github.com/filecoin-project/lotus/api"
	apitypes "github.com/filecoin-project/lotus/api/types"
	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/actors/builtin/miner"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/chain/types/ethtypes"
	"github.com/filecoin-project/lotus/lib/sigs"
	_ "github.com/filecoin-project/lotus/lib/sigs/bls"
	_ "github.com/filecoin-project/lotus/lib/sigs/delegated"
	_ "github.com/filecoin-project/lotus/lib/sigs/secp"
	"github.com/filecoin-project/lotus/metrics"
	"github.com/filecoin-project/lotus/node/impl/full"
	"github.com/filecoin-project/lotus/node/modules/dtypes"
)

const (
	DefaultLookbackCap            = time.Hour * 24
	DefaultStateWaitLookbackLimit = abi.ChainEpoch(20)
	DefaultRateLimitTimeout       = time.Second * 5
	basicRateLimitTokens          = 1
	walletRateLimitTokens         = 1
	chainRateLimitTokens          = 2
	stateRateLimitTokens          = 3
)

// TargetAPI defines the API methods that the Node depends on
// (to make it easy to mock for tests)
type TargetAPI interface {
	MpoolPending(context.Context, types.TipSetKey) ([]*types.SignedMessage, error)
	ChainGetBlock(context.Context, cid.Cid) (*types.BlockHeader, error)
	MinerGetBaseInfo(context.Context, address.Address, abi.ChainEpoch, types.TipSetKey) (*api.MiningBaseInfo, error)
	GasEstimateGasPremium(context.Context, uint64, address.Address, int64, types.TipSetKey) (types.BigInt, error)
	StateReplay(context.Context, types.TipSetKey, cid.Cid) (*api.InvocResult, error)
	StateMinerSectorCount(context.Context, address.Address, types.TipSetKey) (api.MinerSectors, error)
	Version(context.Context) (api.APIVersion, error)
	ChainGetParentMessages(context.Context, cid.Cid) ([]api.Message, error)
	ChainGetParentReceipts(context.Context, cid.Cid) ([]*types.MessageReceipt, error)
	ChainGetBlockMessages(context.Context, cid.Cid) (*api.BlockMessages, error)
	ChainGetMessage(ctx context.Context, mc cid.Cid) (*types.Message, error)
	ChainGetNode(ctx context.Context, p string) (*api.IpldObject, error)
	ChainGetTipSet(ctx context.Context, tsk types.TipSetKey) (*types.TipSet, error)
	ChainGetTipSetByHeight(ctx context.Context, h abi.ChainEpoch, tsk types.TipSetKey) (*types.TipSet, error)
	ChainGetTipSetAfterHeight(ctx context.Context, h abi.ChainEpoch, tsk types.TipSetKey) (*types.TipSet, error)
	ChainHasObj(context.Context, cid.Cid) (bool, error)
	ChainHead(ctx context.Context) (*types.TipSet, error)
	ChainNotify(context.Context) (<-chan []*api.HeadChange, error)
	ChainGetPath(ctx context.Context, from, to types.TipSetKey) ([]*api.HeadChange, error)
	ChainReadObj(context.Context, cid.Cid) ([]byte, error)
	ChainPutObj(context.Context, blocks.Block) error
	ChainGetGenesis(context.Context) (*types.TipSet, error)
	GasEstimateMessageGas(ctx context.Context, msg *types.Message, spec *api.MessageSendSpec, tsk types.TipSetKey) (*types.Message, error)
	MpoolGetNonce(ctx context.Context, addr address.Address) (uint64, error)
	GasBatchEstimateMessageGas(ctx context.Context, estimateMessages []*api.EstimateMessage, fromNonce uint64, tsk types.TipSetKey) ([]*api.EstimateResult, error)
	MpoolPushUntrusted(ctx context.Context, sm *types.SignedMessage) (cid.Cid, error)
	MsigGetAvailableBalance(ctx context.Context, addr address.Address, tsk types.TipSetKey) (types.BigInt, error)
	MsigGetVested(ctx context.Context, addr address.Address, start types.TipSetKey, end types.TipSetKey) (types.BigInt, error)
	MsigGetVestingSchedule(context.Context, address.Address, types.TipSetKey) (api.MsigVesting, error)
	MsigGetPending(ctx context.Context, addr address.Address, ts types.TipSetKey) ([]*api.MsigTransaction, error)
	StateAccountKey(ctx context.Context, addr address.Address, tsk types.TipSetKey) (address.Address, error)
	StateCall(ctx context.Context, msg *types.Message, tsk types.TipSetKey) (*api.InvocResult, error)
	StateDealProviderCollateralBounds(ctx context.Context, size abi.PaddedPieceSize, verified bool, tsk types.TipSetKey) (api.DealCollateralBounds, error)
	StateDecodeParams(ctx context.Context, toAddr address.Address, method abi.MethodNum, params []byte, tsk types.TipSetKey) (interface{}, error)
	StateGetActor(ctx context.Context, actor address.Address, ts types.TipSetKey) (*types.Actor, error)
	StateLookupID(ctx context.Context, addr address.Address, tsk types.TipSetKey) (address.Address, error)
	StateListMiners(ctx context.Context, tsk types.TipSetKey) ([]address.Address, error)
	StateMarketBalance(ctx context.Context, addr address.Address, tsk types.TipSetKey) (api.MarketBalance, error)
	StateMarketStorageDeal(ctx context.Context, dealId abi.DealID, tsk types.TipSetKey) (*api.MarketDeal, error)
	StateNetworkName(context.Context) (dtypes.NetworkName, error)
	StateNetworkVersion(context.Context, types.TipSetKey) (network.Version, error)
	StateSearchMsg(ctx context.Context, from types.TipSetKey, msg cid.Cid, limit abi.ChainEpoch, allowReplaced bool) (*api.MsgLookup, error)
	StateWaitMsg(ctx context.Context, cid cid.Cid, confidence uint64, limit abi.ChainEpoch, allowReplaced bool) (*api.MsgLookup, error)
	StateReadState(ctx context.Context, actor address.Address, tsk types.TipSetKey) (*api.ActorState, error)
	StateMinerPower(context.Context, address.Address, types.TipSetKey) (*api.MinerPower, error)
	StateMinerFaults(context.Context, address.Address, types.TipSetKey) (bitfield.BitField, error)
	StateMinerRecoveries(context.Context, address.Address, types.TipSetKey) (bitfield.BitField, error)
	StateMinerInfo(context.Context, address.Address, types.TipSetKey) (api.MinerInfo, error)
	StateMinerDeadlines(context.Context, address.Address, types.TipSetKey) ([]api.Deadline, error)
	StateMinerAvailableBalance(context.Context, address.Address, types.TipSetKey) (types.BigInt, error)
	StateMinerProvingDeadline(context.Context, address.Address, types.TipSetKey) (*dline.Info, error)
	StateCirculatingSupply(context.Context, types.TipSetKey) (abi.TokenAmount, error)
	StateSectorGetInfo(ctx context.Context, maddr address.Address, n abi.SectorNumber, tsk types.TipSetKey) (*miner.SectorOnChainInfo, error)
	StateVerifiedClientStatus(ctx context.Context, addr address.Address, tsk types.TipSetKey) (*abi.StoragePower, error)
	StateVerifierStatus(ctx context.Context, addr address.Address, tsk types.TipSetKey) (*abi.StoragePower, error)
	StateVMCirculatingSupplyInternal(context.Context, types.TipSetKey) (api.CirculatingSupply, error)
	WalletBalance(context.Context, address.Address) (types.BigInt, error)

	EthAddressToFilecoinAddress(ctx context.Context, ethAddress ethtypes.EthAddress) (address.Address, error)
	FilecoinAddressToEthAddress(ctx context.Context, filecoinAddress address.Address) (ethtypes.EthAddress, error)
	EthBlockNumber(ctx context.Context) (ethtypes.EthUint64, error)
	EthGetBlockTransactionCountByNumber(ctx context.Context, blkNum ethtypes.EthUint64) (ethtypes.EthUint64, error)
	EthGetBlockTransactionCountByHash(ctx context.Context, blkHash ethtypes.EthHash) (ethtypes.EthUint64, error)
	EthGetBlockByHash(ctx context.Context, blkHash ethtypes.EthHash, fullTxInfo bool) (ethtypes.EthBlock, error)
	EthGetBlockByNumber(ctx context.Context, blkNum string, fullTxInfo bool) (ethtypes.EthBlock, error)
	EthGetTransactionByHashLimited(ctx context.Context, txHash *ethtypes.EthHash, limit abi.ChainEpoch) (*ethtypes.EthTx, error)
	EthGetTransactionHashByCid(ctx context.Context, cid cid.Cid) (*ethtypes.EthHash, error)
	EthGetMessageCidByTransactionHash(ctx context.Context, txHash *ethtypes.EthHash) (*cid.Cid, error)
	EthGetTransactionCount(ctx context.Context, sender ethtypes.EthAddress, blkParam ethtypes.EthBlockNumberOrHash) (ethtypes.EthUint64, error)
	EthGetTransactionReceiptLimited(ctx context.Context, txHash ethtypes.EthHash, limit abi.ChainEpoch) (*api.EthTxReceipt, error)
	EthGetTransactionByBlockHashAndIndex(ctx context.Context, blkHash ethtypes.EthHash, txIndex ethtypes.EthUint64) (ethtypes.EthTx, error)
	EthGetTransactionByBlockNumberAndIndex(ctx context.Context, blkNum ethtypes.EthUint64, txIndex ethtypes.EthUint64) (ethtypes.EthTx, error)
	EthGetCode(ctx context.Context, address ethtypes.EthAddress, blkParam ethtypes.EthBlockNumberOrHash) (ethtypes.EthBytes, error)
	EthGetStorageAt(ctx context.Context, address ethtypes.EthAddress, position ethtypes.EthBytes, blkParam ethtypes.EthBlockNumberOrHash) (ethtypes.EthBytes, error)
	EthGetBalance(ctx context.Context, address ethtypes.EthAddress, blkParam ethtypes.EthBlockNumberOrHash) (ethtypes.EthBigInt, error)
	EthChainId(ctx context.Context) (ethtypes.EthUint64, error)
	EthSyncing(ctx context.Context) (ethtypes.EthSyncingResult, error)
	NetVersion(ctx context.Context) (string, error)
	NetListening(ctx context.Context) (bool, error)
	EthProtocolVersion(ctx context.Context) (ethtypes.EthUint64, error)
	EthGasPrice(ctx context.Context) (ethtypes.EthBigInt, error)
	EthFeeHistory(ctx context.Context, p jsonrpc.RawParams) (ethtypes.EthFeeHistory, error)
	EthMaxPriorityFeePerGas(ctx context.Context) (ethtypes.EthBigInt, error)
	EthEstimateGas(ctx context.Context, tx ethtypes.EthCall) (ethtypes.EthUint64, error)
	EthCall(ctx context.Context, tx ethtypes.EthCall, blkParam ethtypes.EthBlockNumberOrHash) (ethtypes.EthBytes, error)
	EthSendRawTransaction(ctx context.Context, rawTx ethtypes.EthBytes) (ethtypes.EthHash, error)
	EthGetLogs(ctx context.Context, filter *ethtypes.EthFilterSpec) (*ethtypes.EthFilterResult, error)
	EthGetFilterChanges(ctx context.Context, id ethtypes.EthFilterID) (*ethtypes.EthFilterResult, error)
	EthGetFilterLogs(ctx context.Context, id ethtypes.EthFilterID) (*ethtypes.EthFilterResult, error)
	EthNewFilter(ctx context.Context, filter *ethtypes.EthFilterSpec) (ethtypes.EthFilterID, error)
	EthNewBlockFilter(ctx context.Context) (ethtypes.EthFilterID, error)
	EthNewPendingTransactionFilter(ctx context.Context) (ethtypes.EthFilterID, error)
	EthUninstallFilter(ctx context.Context, id ethtypes.EthFilterID) (bool, error)
	EthSubscribe(ctx context.Context, params jsonrpc.RawParams) (ethtypes.EthSubscriptionID, error)
	EthUnsubscribe(ctx context.Context, id ethtypes.EthSubscriptionID) (bool, error)
	Web3ClientVersion(ctx context.Context) (string, error)
}

var _ TargetAPI = *new(api.FullNode) // gateway depends on latest

type Node struct {
	target                 TargetAPI
	subHnd                 *EthSubHandler
	lookbackCap            time.Duration
	stateWaitLookbackLimit abi.ChainEpoch
	rateLimiter            *rate.Limiter
	rateLimitTimeout       time.Duration
	errLookback            error
}

var (
	_ api.Gateway         = (*Node)(nil)
	_ full.ChainModuleAPI = (*Node)(nil)
	_ full.GasModuleAPI   = (*Node)(nil)
	_ full.MpoolModuleAPI = (*Node)(nil)
	_ full.StateModuleAPI = (*Node)(nil)
)

// NewNode creates a new gateway node.
func NewNode(api TargetAPI, sHnd *EthSubHandler, lookbackCap time.Duration, stateWaitLookbackLimit abi.ChainEpoch, rateLimit int64, rateLimitTimeout time.Duration) *Node {
	var limit rate.Limit
	if rateLimit == 0 {
		limit = rate.Inf
	} else {
		limit = rate.Every(time.Second / time.Duration(rateLimit))
	}
	return &Node{
		target:                 api,
		subHnd:                 sHnd,
		lookbackCap:            lookbackCap,
		stateWaitLookbackLimit: stateWaitLookbackLimit,
		rateLimiter:            rate.NewLimiter(limit, stateRateLimitTokens),
		rateLimitTimeout:       rateLimitTimeout,
		errLookback:            fmt.Errorf("lookbacks of more than %s are disallowed", lookbackCap),
	}
}

func (gw *Node) checkTipsetKey(ctx context.Context, tsk types.TipSetKey) error {
	if tsk.IsEmpty() {
		return nil
	}

	ts, err := gw.target.ChainGetTipSet(ctx, tsk)
	if err != nil {
		return err
	}

	return gw.checkTipset(ts)
}

func (gw *Node) checkTipset(ts *types.TipSet) error {
	at := time.Unix(int64(ts.Blocks()[0].Timestamp), 0)
	if err := gw.checkTimestamp(at); err != nil {
		return fmt.Errorf("bad tipset: %w", err)
	}
	return nil
}

func (gw *Node) checkTipsetHeight(ts *types.TipSet, h abi.ChainEpoch) error {
	if h > ts.Height() {
		return fmt.Errorf("tipset height in future")
	}
	tsBlock := ts.Blocks()[0]
	heightDelta := time.Duration(uint64(tsBlock.Height-h)*build.BlockDelaySecs) * time.Second
	timeAtHeight := time.Unix(int64(tsBlock.Timestamp), 0).Add(-heightDelta)

	if err := gw.checkTimestamp(timeAtHeight); err != nil {
		return fmt.Errorf("bad tipset height: %w", err)
	}
	return nil
}

func (gw *Node) checkTimestamp(at time.Time) error {
	if time.Since(at) > gw.lookbackCap {
		return gw.errLookback
	}
	return nil
}

func (gw *Node) limit(ctx context.Context, tokens int) error {
	ctx2, cancel := context.WithTimeout(ctx, gw.rateLimitTimeout)
	defer cancel()
	if perConnLimiter, ok := ctx2.Value(perConnLimiterKey).(*rate.Limiter); ok {
		err := perConnLimiter.WaitN(ctx2, tokens)
		if err != nil {
			return fmt.Errorf("connection limited. %w", err)
		}
	}

	err := gw.rateLimiter.WaitN(ctx2, tokens)
	if err != nil {
		stats.Record(ctx, metrics.RateLimitCount.M(1))
		return fmt.Errorf("server busy. %w", err)
	}
	return nil
}

func (gw *Node) Discover(ctx context.Context) (apitypes.OpenRPCDocument, error) {
	return build.OpenRPCDiscoverJSON_Gateway(), nil
}

func (gw *Node) Version(ctx context.Context) (api.APIVersion, error) {
	if err := gw.limit(ctx, basicRateLimitTokens); err != nil {
		return api.APIVersion{}, err
	}
	return gw.target.Version(ctx)
}

func (gw *Node) ChainGetParentMessages(ctx context.Context, c cid.Cid) ([]api.Message, error) {
	if err := gw.limit(ctx, chainRateLimitTokens); err != nil {
		return nil, err
	}
	return gw.target.ChainGetParentMessages(ctx, c)
}

func (gw *Node) ChainGetParentReceipts(ctx context.Context, c cid.Cid) ([]*types.MessageReceipt, error) {
	if err := gw.limit(ctx, chainRateLimitTokens); err != nil {
		return nil, err
	}
	return gw.target.ChainGetParentReceipts(ctx, c)
}

func (gw *Node) ChainGetBlockMessages(ctx context.Context, c cid.Cid) (*api.BlockMessages, error) {
	if err := gw.limit(ctx, chainRateLimitTokens); err != nil {
		return nil, err
	}
	return gw.target.ChainGetBlockMessages(ctx, c)
}

func (gw *Node) ChainHasObj(ctx context.Context, c cid.Cid) (bool, error) {
	if err := gw.limit(ctx, chainRateLimitTokens); err != nil {
		return false, err
	}
	return gw.target.ChainHasObj(ctx, c)
}

func (gw *Node) ChainHead(ctx context.Context) (*types.TipSet, error) {
	if err := gw.limit(ctx, chainRateLimitTokens); err != nil {
		return nil, err
	}

	return gw.target.ChainHead(ctx)
}

func (gw *Node) ChainGetMessage(ctx context.Context, mc cid.Cid) (*types.Message, error) {
	if err := gw.limit(ctx, chainRateLimitTokens); err != nil {
		return nil, err
	}
	return gw.target.ChainGetMessage(ctx, mc)
}

func (gw *Node) ChainGetTipSet(ctx context.Context, tsk types.TipSetKey) (*types.TipSet, error) {
	if err := gw.limit(ctx, chainRateLimitTokens); err != nil {
		return nil, err
	}
	return gw.target.ChainGetTipSet(ctx, tsk)
}

func (gw *Node) ChainGetTipSetByHeight(ctx context.Context, h abi.ChainEpoch, tsk types.TipSetKey) (*types.TipSet, error) {
	if err := gw.limit(ctx, chainRateLimitTokens); err != nil {
		return nil, err
	}
	if err := gw.checkTipSetHeight(ctx, h, tsk); err != nil {
		return nil, err
	}
	return gw.target.ChainGetTipSetByHeight(ctx, h, tsk)
}

func (gw *Node) ChainGetTipSetAfterHeight(ctx context.Context, h abi.ChainEpoch, tsk types.TipSetKey) (*types.TipSet, error) {
	if err := gw.limit(ctx, chainRateLimitTokens); err != nil {
		return nil, err
	}
	if err := gw.checkTipSetHeight(ctx, h, tsk); err != nil {
		return nil, err
	}
	return gw.target.ChainGetTipSetAfterHeight(ctx, h, tsk)
}

func (gw *Node) checkTipSetHeight(ctx context.Context, h abi.ChainEpoch, tsk types.TipSetKey) error {
	var ts *types.TipSet
	if tsk.IsEmpty() {
		head, err := gw.target.ChainHead(ctx)
		if err != nil {
			return err
		}
		ts = head
	} else {
		gts, err := gw.target.ChainGetTipSet(ctx, tsk)
		if err != nil {
			return err
		}
		ts = gts
	}

	// Check if the tipset key refers to gw tipset that's too far in the past
	if err := gw.checkTipset(ts); err != nil {
		return err
	}

	// Check if the height is too far in the past
	if err := gw.checkTipsetHeight(ts, h); err != nil {
		return err
	}

	return nil
}

func (gw *Node) ChainGetNode(ctx context.Context, p string) (*api.IpldObject, error) {
	if err := gw.limit(ctx, chainRateLimitTokens); err != nil {
		return nil, err
	}
	return gw.target.ChainGetNode(ctx, p)
}

func (gw *Node) ChainNotify(ctx context.Context) (<-chan []*api.HeadChange, error) {
	if err := gw.limit(ctx, chainRateLimitTokens); err != nil {
		return nil, err
	}
	return gw.target.ChainNotify(ctx)
}

func (gw *Node) ChainGetPath(ctx context.Context, from, to types.TipSetKey) ([]*api.HeadChange, error) {
	if err := gw.limit(ctx, chainRateLimitTokens); err != nil {
		return nil, err
	}
	if err := gw.checkTipsetKey(ctx, from); err != nil {
		return nil, xerrors.Errorf("gateway: checking 'from' tipset: %w", err)
	}
	if err := gw.checkTipsetKey(ctx, to); err != nil {
		return nil, xerrors.Errorf("gateway: checking 'to' tipset: %w", err)
	}
	return gw.target.ChainGetPath(ctx, from, to)
}

func (gw *Node) ChainGetGenesis(ctx context.Context) (*types.TipSet, error) {
	if err := gw.limit(ctx, chainRateLimitTokens); err != nil {
		return nil, err
	}
	return gw.target.ChainGetGenesis(ctx)
}

func (gw *Node) ChainReadObj(ctx context.Context, c cid.Cid) ([]byte, error) {
	if err := gw.limit(ctx, chainRateLimitTokens); err != nil {
		return nil, err
	}
	return gw.target.ChainReadObj(ctx, c)
}

func (gw *Node) ChainPutObj(context.Context, blocks.Block) error {
	return xerrors.New("not supported")
}

func (gw *Node) GasEstimateMessageGas(ctx context.Context, msg *types.Message, spec *api.MessageSendSpec, tsk types.TipSetKey) (*types.Message, error) {
	if err := gw.limit(ctx, stateRateLimitTokens); err != nil {
		return nil, err
	}
	if err := gw.checkTipsetKey(ctx, tsk); err != nil {
		return nil, err
	}
	return gw.target.GasEstimateMessageGas(ctx, msg, spec, tsk)
}

func (gw *Node) MpoolPush(ctx context.Context, sm *types.SignedMessage) (cid.Cid, error) {
	if err := gw.limit(ctx, stateRateLimitTokens); err != nil {
		return cid.Cid{}, err
	}
	// TODO: additional anti-spam checks
	return gw.target.MpoolPushUntrusted(ctx, sm)
}

func (gw *Node) MsigGetAvailableBalance(ctx context.Context, addr address.Address, tsk types.TipSetKey) (types.BigInt, error) {
	if err := gw.limit(ctx, walletRateLimitTokens); err != nil {
		return types.BigInt{}, err
	}
	if err := gw.checkTipsetKey(ctx, tsk); err != nil {
		return types.NewInt(0), err
	}
	return gw.target.MsigGetAvailableBalance(ctx, addr, tsk)
}

func (gw *Node) MsigGetVested(ctx context.Context, addr address.Address, start types.TipSetKey, end types.TipSetKey) (types.BigInt, error) {
	if err := gw.limit(ctx, walletRateLimitTokens); err != nil {
		return types.BigInt{}, err
	}
	if err := gw.checkTipsetKey(ctx, start); err != nil {
		return types.NewInt(0), err
	}
	if err := gw.checkTipsetKey(ctx, end); err != nil {
		return types.NewInt(0), err
	}
	return gw.target.MsigGetVested(ctx, addr, start, end)
}

func (gw *Node) MsigGetVestingSchedule(ctx context.Context, addr address.Address, tsk types.TipSetKey) (api.MsigVesting, error) {
	if err := gw.limit(ctx, walletRateLimitTokens); err != nil {
		return api.MsigVesting{}, err
	}
	if err := gw.checkTipsetKey(ctx, tsk); err != nil {
		return api.MsigVesting{}, err
	}
	return gw.target.MsigGetVestingSchedule(ctx, addr, tsk)
}

func (gw *Node) MsigGetPending(ctx context.Context, addr address.Address, tsk types.TipSetKey) ([]*api.MsigTransaction, error) {
	if err := gw.limit(ctx, walletRateLimitTokens); err != nil {
		return nil, err
	}
	if err := gw.checkTipsetKey(ctx, tsk); err != nil {
		return nil, err
	}
	return gw.target.MsigGetPending(ctx, addr, tsk)
}

func (gw *Node) StateAccountKey(ctx context.Context, addr address.Address, tsk types.TipSetKey) (address.Address, error) {
	if err := gw.limit(ctx, stateRateLimitTokens); err != nil {
		return address.Address{}, err
	}
	if err := gw.checkTipsetKey(ctx, tsk); err != nil {
		return address.Undef, err
	}
	return gw.target.StateAccountKey(ctx, addr, tsk)
}

func (gw *Node) StateDealProviderCollateralBounds(ctx context.Context, size abi.PaddedPieceSize, verified bool, tsk types.TipSetKey) (api.DealCollateralBounds, error) {
	if err := gw.limit(ctx, stateRateLimitTokens); err != nil {
		return api.DealCollateralBounds{}, err
	}
	if err := gw.checkTipsetKey(ctx, tsk); err != nil {
		return api.DealCollateralBounds{}, err
	}
	return gw.target.StateDealProviderCollateralBounds(ctx, size, verified, tsk)
}

func (gw *Node) StateGetActor(ctx context.Context, actor address.Address, tsk types.TipSetKey) (*types.Actor, error) {
	if err := gw.limit(ctx, stateRateLimitTokens); err != nil {
		return nil, err
	}
	if err := gw.checkTipsetKey(ctx, tsk); err != nil {
		return nil, err
	}
	return gw.target.StateGetActor(ctx, actor, tsk)
}

func (gw *Node) StateListMiners(ctx context.Context, tsk types.TipSetKey) ([]address.Address, error) {
	if err := gw.limit(ctx, stateRateLimitTokens); err != nil {
		return nil, err
	}
	if err := gw.checkTipsetKey(ctx, tsk); err != nil {
		return nil, err
	}
	return gw.target.StateListMiners(ctx, tsk)
}

func (gw *Node) StateLookupID(ctx context.Context, addr address.Address, tsk types.TipSetKey) (address.Address, error) {
	if err := gw.limit(ctx, stateRateLimitTokens); err != nil {
		return address.Address{}, err
	}
	if err := gw.checkTipsetKey(ctx, tsk); err != nil {
		return address.Undef, err
	}
	return gw.target.StateLookupID(ctx, addr, tsk)
}

func (gw *Node) StateMarketBalance(ctx context.Context, addr address.Address, tsk types.TipSetKey) (api.MarketBalance, error) {
	if err := gw.limit(ctx, stateRateLimitTokens); err != nil {
		return api.MarketBalance{}, err
	}
	if err := gw.checkTipsetKey(ctx, tsk); err != nil {
		return api.MarketBalance{}, err
	}
	return gw.target.StateMarketBalance(ctx, addr, tsk)
}

func (gw *Node) StateMarketStorageDeal(ctx context.Context, dealId abi.DealID, tsk types.TipSetKey) (*api.MarketDeal, error) {
	if err := gw.limit(ctx, stateRateLimitTokens); err != nil {
		return nil, err
	}
	if err := gw.checkTipsetKey(ctx, tsk); err != nil {
		return nil, err
	}
	return gw.target.StateMarketStorageDeal(ctx, dealId, tsk)
}

func (gw *Node) StateNetworkVersion(ctx context.Context, tsk types.TipSetKey) (network.Version, error) {
	if err := gw.limit(ctx, stateRateLimitTokens); err != nil {
		return network.VersionMax, err
	}
	if err := gw.checkTipsetKey(ctx, tsk); err != nil {
		return network.VersionMax, err
	}
	return gw.target.StateNetworkVersion(ctx, tsk)
}

func (gw *Node) StateSearchMsg(ctx context.Context, from types.TipSetKey, msg cid.Cid, limit abi.ChainEpoch, allowReplaced bool) (*api.MsgLookup, error) {
	if err := gw.limit(ctx, stateRateLimitTokens); err != nil {
		return nil, err
	}
	if limit == api.LookbackNoLimit {
		limit = gw.stateWaitLookbackLimit
	}
	if gw.stateWaitLookbackLimit != api.LookbackNoLimit && limit > gw.stateWaitLookbackLimit {
		limit = gw.stateWaitLookbackLimit
	}
	if err := gw.checkTipsetKey(ctx, from); err != nil {
		return nil, err
	}
	return gw.target.StateSearchMsg(ctx, from, msg, limit, allowReplaced)
}

func (gw *Node) StateWaitMsg(ctx context.Context, msg cid.Cid, confidence uint64, limit abi.ChainEpoch, allowReplaced bool) (*api.MsgLookup, error) {
	if err := gw.limit(ctx, stateRateLimitTokens); err != nil {
		return nil, err
	}
	if limit == api.LookbackNoLimit {
		limit = gw.stateWaitLookbackLimit
	}
	if gw.stateWaitLookbackLimit != api.LookbackNoLimit && limit > gw.stateWaitLookbackLimit {
		limit = gw.stateWaitLookbackLimit
	}
	return gw.target.StateWaitMsg(ctx, msg, confidence, limit, allowReplaced)
}

func (gw *Node) StateReadState(ctx context.Context, actor address.Address, tsk types.TipSetKey) (*api.ActorState, error) {
	if err := gw.limit(ctx, stateRateLimitTokens); err != nil {
		return nil, err
	}
	if err := gw.checkTipsetKey(ctx, tsk); err != nil {
		return nil, err
	}
	return gw.target.StateReadState(ctx, actor, tsk)
}

func (gw *Node) StateMinerPower(ctx context.Context, m address.Address, tsk types.TipSetKey) (*api.MinerPower, error) {
	if err := gw.limit(ctx, stateRateLimitTokens); err != nil {
		return nil, err
	}
	if err := gw.checkTipsetKey(ctx, tsk); err != nil {
		return nil, err
	}
	return gw.target.StateMinerPower(ctx, m, tsk)
}

func (gw *Node) StateMinerFaults(ctx context.Context, m address.Address, tsk types.TipSetKey) (bitfield.BitField, error) {
	if err := gw.limit(ctx, stateRateLimitTokens); err != nil {
		return bitfield.BitField{}, err
	}
	if err := gw.checkTipsetKey(ctx, tsk); err != nil {
		return bitfield.BitField{}, err
	}
	return gw.target.StateMinerFaults(ctx, m, tsk)
}

func (gw *Node) StateMinerRecoveries(ctx context.Context, m address.Address, tsk types.TipSetKey) (bitfield.BitField, error) {
	if err := gw.limit(ctx, stateRateLimitTokens); err != nil {
		return bitfield.BitField{}, err
	}
	if err := gw.checkTipsetKey(ctx, tsk); err != nil {
		return bitfield.BitField{}, err
	}
	return gw.target.StateMinerRecoveries(ctx, m, tsk)
}

func (gw *Node) StateMinerInfo(ctx context.Context, m address.Address, tsk types.TipSetKey) (api.MinerInfo, error) {
	if err := gw.limit(ctx, stateRateLimitTokens); err != nil {
		return api.MinerInfo{}, err
	}
	if err := gw.checkTipsetKey(ctx, tsk); err != nil {
		return api.MinerInfo{}, err
	}
	return gw.target.StateMinerInfo(ctx, m, tsk)
}

func (gw *Node) StateMinerDeadlines(ctx context.Context, m address.Address, tsk types.TipSetKey) ([]api.Deadline, error) {
	if err := gw.limit(ctx, stateRateLimitTokens); err != nil {
		return nil, err
	}
	if err := gw.checkTipsetKey(ctx, tsk); err != nil {
		return nil, err
	}
	return gw.target.StateMinerDeadlines(ctx, m, tsk)
}

func (gw *Node) StateMinerAvailableBalance(ctx context.Context, m address.Address, tsk types.TipSetKey) (types.BigInt, error) {
	if err := gw.limit(ctx, stateRateLimitTokens); err != nil {
		return types.BigInt{}, err
	}
	if err := gw.checkTipsetKey(ctx, tsk); err != nil {
		return types.BigInt{}, err
	}
	return gw.target.StateMinerAvailableBalance(ctx, m, tsk)
}

func (gw *Node) StateMinerProvingDeadline(ctx context.Context, m address.Address, tsk types.TipSetKey) (*dline.Info, error) {
	if err := gw.limit(ctx, stateRateLimitTokens); err != nil {
		return nil, err
	}
	if err := gw.checkTipsetKey(ctx, tsk); err != nil {
		return nil, err
	}
	return gw.target.StateMinerProvingDeadline(ctx, m, tsk)
}

func (gw *Node) StateCirculatingSupply(ctx context.Context, tsk types.TipSetKey) (abi.TokenAmount, error) {
	if err := gw.limit(ctx, stateRateLimitTokens); err != nil {
		return abi.TokenAmount{}, err
	}
	if err := gw.checkTipsetKey(ctx, tsk); err != nil {
		return abi.TokenAmount{}, err
	}
	return gw.target.StateCirculatingSupply(ctx, tsk)
}

func (gw *Node) StateSectorGetInfo(ctx context.Context, maddr address.Address, n abi.SectorNumber, tsk types.TipSetKey) (*miner.SectorOnChainInfo, error) {
	if err := gw.limit(ctx, stateRateLimitTokens); err != nil {
		return nil, err
	}
	if err := gw.checkTipsetKey(ctx, tsk); err != nil {
		return nil, err
	}
	return gw.target.StateSectorGetInfo(ctx, maddr, n, tsk)
}

func (gw *Node) StateVerifiedClientStatus(ctx context.Context, addr address.Address, tsk types.TipSetKey) (*abi.StoragePower, error) {
	if err := gw.limit(ctx, stateRateLimitTokens); err != nil {
		return nil, err
	}
	if err := gw.checkTipsetKey(ctx, tsk); err != nil {
		return nil, err
	}
	return gw.target.StateVerifiedClientStatus(ctx, addr, tsk)
}

func (gw *Node) StateVMCirculatingSupplyInternal(ctx context.Context, tsk types.TipSetKey) (api.CirculatingSupply, error) {
	if err := gw.limit(ctx, stateRateLimitTokens); err != nil {
		return api.CirculatingSupply{}, err
	}
	if err := gw.checkTipsetKey(ctx, tsk); err != nil {
		return api.CirculatingSupply{}, err
	}
	return gw.target.StateVMCirculatingSupplyInternal(ctx, tsk)
}

func (gw *Node) WalletVerify(ctx context.Context, k address.Address, msg []byte, sig *crypto.Signature) (bool, error) {
	if err := gw.limit(ctx, stateRateLimitTokens); err != nil {
		return false, err
	}
	return sigs.Verify(sig, k, msg) == nil, nil
}

func (gw *Node) WalletBalance(ctx context.Context, k address.Address) (types.BigInt, error) {
	if err := gw.limit(ctx, stateRateLimitTokens); err != nil {
		return types.BigInt{}, err
	}
	return gw.target.WalletBalance(ctx, k)
}

func (gw *Node) GasBatchEstimateMessageGas(ctx context.Context, estimateMessages []*api.EstimateMessage, fromNonce uint64, tsk types.TipSetKey) ([]*api.EstimateResult, error) {
	return gw.target.GasBatchEstimateMessageGas(ctx, estimateMessages, fromNonce, tsk)
}
