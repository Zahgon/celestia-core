package statesync

import (
	"github.com/cosmos/gogoproto/proto"

	"github.com/cometbft/cometbft/p2p"
)

var _ p2p.Wrapper = &ChunkRequest{}
var _ p2p.Wrapper = &ChunkResponse{}
var _ p2p.Wrapper = &SnapshotsRequest{}
var _ p2p.Wrapper = &SnapshotsResponse{}

func (m *SnapshotsResponse) Wrap() proto.Message {
	_ = "STUB: not implemented"
	return *new(proto.Message)
}

func (m *SnapshotsRequest) Wrap() proto.Message {
	_ = "STUB: not implemented"
	return *new(proto.Message)
}

func (m *ChunkResponse) Wrap() proto.Message { _ = "STUB: not implemented"; return *new(proto.Message) }

func (m *ChunkRequest) Wrap() proto.Message { _ = "STUB: not implemented"; return *new(proto.Message) }

// Unwrap implements the p2p Wrapper interface and unwraps a wrapped state sync
// proto message.
func (m *Message) Unwrap() (proto.Message, error) {
	_ = "STUB: not implemented"
	return *new(proto.Message), nil
}
