package consensus

import (
	"github.com/cosmos/gogoproto/proto"

	"github.com/cometbft/cometbft/p2p"
)

var _ p2p.Wrapper = &VoteSetBits{}
var _ p2p.Wrapper = &VoteSetMaj23{}
var _ p2p.Wrapper = &Vote{}
var _ p2p.Wrapper = &ProposalPOL{}
var _ p2p.Wrapper = &Proposal{}
var _ p2p.Wrapper = &NewValidBlock{}
var _ p2p.Wrapper = &NewRoundStep{}
var _ p2p.Wrapper = &HasVote{}
var _ p2p.Wrapper = &BlockPart{}

func (m *VoteSetBits) Wrap() proto.Message { _ = "STUB: not implemented"; return *new(proto.Message) }

func (m *VoteSetMaj23) Wrap() proto.Message { _ = "STUB: not implemented"; return *new(proto.Message) }

func (m *HasVote) Wrap() proto.Message { _ = "STUB: not implemented"; return *new(proto.Message) }

func (m *Vote) Wrap() proto.Message { _ = "STUB: not implemented"; return *new(proto.Message) }

func (m *BlockPart) Wrap() proto.Message { _ = "STUB: not implemented"; return *new(proto.Message) }

func (m *ProposalPOL) Wrap() proto.Message { _ = "STUB: not implemented"; return *new(proto.Message) }

func (m *Proposal) Wrap() proto.Message { _ = "STUB: not implemented"; return *new(proto.Message) }

func (m *NewValidBlock) Wrap() proto.Message { _ = "STUB: not implemented"; return *new(proto.Message) }

func (m *NewRoundStep) Wrap() proto.Message { _ = "STUB: not implemented"; return *new(proto.Message) }

// Unwrap implements the p2p Wrapper interface and unwraps a wrapped consensus
// proto message.
func (m *Message) Unwrap() (proto.Message, error) {
	_ = "STUB: not implemented"
	return *new(proto.Message), nil
}
