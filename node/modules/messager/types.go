package messager

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
)

type MessageState int

const (
	UnKnown MessageState = iota
	UnFillMsg
	FillMsg
	OnChainMsg
	FailedMsg
	ReplacedMsg
)

type MessagerConfig struct {
	Url   string
	Token string
}

type MessageWithUID struct {
	UnsignedMessage types.Message
	ID              string
}

type Signature struct {
	Type byte
	Data []byte
}

type MsgDetail struct {
	ID string

	UnsignedCid *cid.Cid
	SignedCid   *cid.Cid
	types.Message
	*Signature

	Height     int64
	Confidence int64
	Receipt    *types.MessageReceipt
	TipSetKey  types.TipSetKey

	Meta *MsgMeta

	WalletName string

	State MessageState
}

func (msgDetail *MsgDetail) Cid() cid.Cid {
	if msgDetail.UnsignedCid == nil {
		return cid.Undef
	}

	if msgDetail.From.Protocol() == address.BLS {
		return *msgDetail.UnsignedCid
	}
	return *msgDetail.SignedCid
}

type MsgMeta struct {
	ExpireEpoch       abi.ChainEpoch `json:"expireEpoch"`
	GasOverEstimation float64        `json:"gasOverEstimation"`
	MaxFee            big.Int        `json:"maxFee,omitempty"`
	MaxFeeCap         big.Int        `json:"maxFeeCap"`
}
