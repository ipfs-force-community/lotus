package messager

import (
	"context"
	"github.com/filecoin-project/go-address"
	jsonrpc "github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/lotus/chain/types"
	logging "github.com/ipfs/go-log/v2"
	"golang.org/x/xerrors"
	"net/http"
	"time"
)

var messagerLog = logging.Logger("messager")

type IMessager interface {
	HasWalletAddress(ctx context.Context, addr address.Address) (bool, error)
	HasMessageByUid(ctx context.Context, id string) (bool, error)
	WaitMessage(ctx context.Context, id string, confidence uint64) (*MsgDetail, error)
	PushMessage(ctx context.Context, msg *types.Message, meta *MsgMeta) (string, error)
	PushMessageWithId(ctx context.Context, id string, msg *types.Message, meta *MsgMeta) (string, error)
	GetMessageByUid(ctx context.Context, id string) (*MsgDetail, error)
}

var _ IMessager = (*Messager)(nil)

type Messager struct {
	Internal struct {
		HasMessageByUid   func(ctx context.Context, id string) (bool, error)
		WaitMessage       func(ctx context.Context, id string, confidence uint64) (*types.Message, error)
		PushMessage       func(ctx context.Context, msg *types.Message, meta *MsgMeta) (string, error)
		PushMessageWithId func(ctx context.Context, id string, msg *types.Message, meta *MsgMeta) (string, error)
		GetMessageByUid   func(ctx context.Context, id string) (*MsgDetail, error)
		HasWalletAddress  func(ctx context.Context, addr address.Address) (bool, error)
	}
}

func (message *Messager) HasMessageByUid(ctx context.Context, id string) (bool, error) {
	return message.Internal.HasMessageByUid(ctx, id)
}

func NewMessager(cfg *MessagerConfig) (*Messager, error) {
	headers := http.Header{}
	if len(cfg.Token) != 0 {
		headers.Add("Authorization", "Bearer "+cfg.Token)
	}

	var messager Messager
	_, err := jsonrpc.NewMergeClient(context.Background(), cfg.Url, "Message",
		[]interface{}{
			&messager.Internal,
		},
		headers,
	)
	return &messager, err
}

func (message *Messager) PushMessage(ctx context.Context, msg *types.Message, meta *MsgMeta) (string, error) {
	return message.Internal.PushMessage(ctx, msg, meta)
}

func (message *Messager) PushMessageWithId(ctx context.Context, id string, msg *types.Message, meta *MsgMeta) (string, error) {
	return message.Internal.PushMessageWithId(ctx, id, msg, meta)
}

func (message *Messager) GetMessageByUid(ctx context.Context, id string) (*MsgDetail, error) {
	return message.Internal.GetMessageByUid(ctx, id)
}

func (message *Messager) HasWalletAddress(ctx context.Context, addr address.Address) (bool, error) {
	return message.Internal.HasWalletAddress(ctx, addr)
}

func (message *Messager) WaitMessage(ctx context.Context, id string, confidence uint64) (*MsgDetail, error) {
	tm := time.NewTicker(time.Second * 30)
	defer tm.Stop()

	doneCh := make(chan struct{}, 1)
	doneCh <- struct{}{}

	for {
		select {
		case <-doneCh:
			msg, err := message.Internal.GetMessageByUid(ctx, id)
			if err != nil {
				return nil, err
			}

			switch msg.State {
			case UnKnown:
				continue
			//OffChain
			case FillMsg:
				fallthrough
			case UnFillMsg:
				fallthrough
			//OnChain
			case ReplacedMsg:
				fallthrough
			case OnChainMsg:
				if msg.Confidence >= int64(confidence) {
					return msg, nil
				}
				continue
			//Error
			case FailedMsg:
				var reason string
				if msg.Receipt != nil {
					reason = string(msg.Receipt.Return)
				}
				return nil, xerrors.Errorf("msg failed due to %s", reason)
			}

		case <-tm.C:
			doneCh <- struct{}{}
		case <-ctx.Done():
			return nil, xerrors.New("exit by client ")
		}
	}
}
