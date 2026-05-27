package p2p

import (
	"github.com/cosmos/gogoproto/proto"
)

func (m *PexAddrs) Wrap() proto.Message { _ = "STUB: not implemented"; return *new(proto.Message) }

func (m *PexRequest) Wrap() proto.Message { _ = "STUB: not implemented"; return *new(proto.Message) }

// Unwrap implements the p2p Wrapper interface and unwraps a wrapped PEX
// message.
func (m *Message) Unwrap() (proto.Message, error) {
	_ = "STUB: not implemented"
	return *new(proto.Message), nil
}
