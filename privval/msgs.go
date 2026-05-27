package privval

import (
	"github.com/cosmos/gogoproto/proto"

	privvalproto "github.com/cometbft/cometbft/proto/tendermint/privval"
)

// TODO: Add ChainIDRequest

func mustWrapMsg(pb proto.Message) privvalproto.Message {
	_ = "STUB: not implemented"
	return *new(privvalproto.Message)
}
