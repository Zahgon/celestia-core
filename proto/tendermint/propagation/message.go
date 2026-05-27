package propagation

import (
	"github.com/cosmos/gogoproto/proto"

	"github.com/cometbft/cometbft/p2p"
)

var (
	_ p2p.Wrapper   = &CompactBlock{}
	_ p2p.Wrapper   = &HaveParts{}
	_ p2p.Wrapper   = &WantParts{}
	_ p2p.Wrapper   = &RecoveryPart{}
	_ p2p.Unwrapper = &Message{}
)

// Wrap implements the p2p Wrapper interface wraps a propagation message.
func (m *CompactBlock) Wrap() proto.Message { _ = "STUB: not implemented"; return *new(proto.Message) }

// Wrap implements the p2p Wrapper interface and wraps a propagation message.
func (m *HaveParts) Wrap() proto.Message { _ = "STUB: not implemented"; return *new(proto.Message) }

// Wrap implements the p2p Wrapper interface and wraps a propagation message.
func (m *WantParts) Wrap() proto.Message { _ = "STUB: not implemented"; return *new(proto.Message) }

// Wrap implements the p2p Wrapper interface and wraps a propagation message.
func (m *RecoveryPart) Wrap() proto.Message { _ = "STUB: not implemented"; return *new(proto.Message) }

// Unwrap implements the p2p Wrapper interface and unwraps a wrapped mempool
// message.
func (m *Message) Unwrap() (proto.Message, error) {
	_ = "STUB: not implemented"
	return *new(proto.Message), nil
}
