package types

import (
	"errors"
	"time"

	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
)

var (
	ErrInvalidBlockPartSignature = errors.New("error invalid block part signature")
	ErrInvalidBlockPartHash      = errors.New("error invalid block part hash")
)

// Proposal defines a block proposal for the consensus.
// It refers to the block by BlockID field.
// It must be signed by the correct proposer for the given Height/Round
// to be considered valid. It may depend on votes from a previous round,
// a so-called Proof-of-Lock (POL) round, as noted in the POLRound.
// If POLRound >= 0, then BlockID corresponds to the block that is locked in POLRound.
type Proposal struct {
	Type      cmtproto.SignedMsgType
	Height    int64     `json:"height"`
	Round     int32     `json:"round"`     // there can not be greater than 2_147_483_647 rounds
	POLRound  int32     `json:"pol_round"` // -1 if null.
	BlockID   BlockID   `json:"block_id"`
	Timestamp time.Time `json:"timestamp"`
	Signature []byte    `json:"signature"`
}

// NewProposal returns a new Proposal.
// If there is no POLRound, polRound should be -1.
func NewProposal(height int64, round int32, polRound int32, blockID BlockID) *Proposal {
	_ = "STUB: not implemented"
	return nil
}

// ValidateBasic performs basic validation.
func (p *Proposal) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// ValidateBasic above would pass even if the BlockID was empty:

// NOTE: Timestamp validation is subtle and handled elsewhere.

// String returns a string representation of the Proposal.
//
// 1. height
// 2. round
// 3. block ID
// 4. POL round
// 5. first 6 bytes of signature
// 6. timestamp
//
// See BlockID#String.
func (p *Proposal) String() string { _ = "STUB: not implemented"; return "" }

// ProposalSignBytes returns the proto-encoding of the canonicalized Proposal,
// for signing. Panics if the marshaling fails.
//
// The encoded Protobuf message is varint length-prefixed (using MarshalDelimited)
// for backwards-compatibility with the Amino encoding, due to e.g. hardware
// devices that rely on this encoding.
//
// See CanonicalizeProposal
func ProposalSignBytes(chainID string, p *cmtproto.Proposal) []byte {
	_ = "STUB: not implemented"
	return nil
}

// ToProto converts Proposal to protobuf
func (p *Proposal) ToProto() *cmtproto.Proposal { _ = "STUB: not implemented"; return nil }

// FromProto sets a protobuf Proposal to the given pointer.
// It returns an error if the proposal is invalid.
func ProposalFromProto(pp *cmtproto.Proposal) (*Proposal, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
