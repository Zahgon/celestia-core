package blocksync

import (
	"github.com/cosmos/gogoproto/proto"

	"github.com/cometbft/cometbft/types"
)

const (
	// NOTE: keep up to date with bcproto.BlockResponse
	BlockResponseMessagePrefixSize   = 4
	BlockResponseMessageFieldKeySize = 1
)

var (
	MaxMsgSize = types.MaxBlockSizeBytes +
		BlockResponseMessagePrefixSize +
		BlockResponseMessageFieldKeySize
)

// ValidateMsg validates a message.
func ValidateMsg(pb proto.Message) error { _ = "STUB: not implemented"; return nil }

// Avoid double-calling `types.BlockFromProto` for performance reasons.
// See https://github.com/cometbft/cometbft/issues/1964
