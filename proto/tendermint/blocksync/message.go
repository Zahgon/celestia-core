package blocksync

import (
	"github.com/cosmos/gogoproto/proto"

	"github.com/cometbft/cometbft/p2p"
)

var _ p2p.Wrapper = &StatusRequest{}
var _ p2p.Wrapper = &StatusResponse{}
var _ p2p.Wrapper = &NoBlockResponse{}
var _ p2p.Wrapper = &BlockResponse{}
var _ p2p.Wrapper = &BlockRequest{}

const (
	BlockResponseMessagePrefixSize   = 4
	BlockResponseMessageFieldKeySize = 1
)

func (m *BlockRequest) Wrap() proto.Message { _ = "STUB: not implemented"; return *new(proto.Message) }

func (m *BlockResponse) Wrap() proto.Message { _ = "STUB: not implemented"; return *new(proto.Message) }

func (m *NoBlockResponse) Wrap() proto.Message {
	_ = "STUB: not implemented"
	return *new(proto.Message)
}

func (m *StatusRequest) Wrap() proto.Message { _ = "STUB: not implemented"; return *new(proto.Message) }

func (m *StatusResponse) Wrap() proto.Message {
	_ = "STUB: not implemented"
	return *new(proto.Message)
}

// Unwrap implements the p2p Wrapper interface and unwraps a wrapped blockchain
// message.
func (m *Message) Unwrap() (proto.Message, error) {
	_ = "STUB: not implemented"
	return *new(proto.Message), nil
}
