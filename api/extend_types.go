package api

import "github.com/filecoin-project/lotus/chain/types"

// 接口使用到的自定义的类型，减少 cherry-pick 冲突

type EstimateMessage struct {
	Msg  *types.Message
	Spec *MessageSendSpec
}

type EstimateResult struct {
	Msg *types.Message
	Err string
}
