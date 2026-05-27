package types

import (
	"time"

	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
)

// Canonical* wraps the structs in types for amino encoding them for use in SignBytes / the Signable interface.

// TimeFormat is used for generating the sigs
const TimeFormat = time.RFC3339Nano

//-----------------------------------
// Canonicalize the structs

func CanonicalizeBlockID(bid cmtproto.BlockID) *cmtproto.CanonicalBlockID {
	_ = "STUB: not implemented"
	return nil
}

// CanonicalizeVote transforms the given PartSetHeader to a CanonicalPartSetHeader.
func CanonicalizePartSetHeader(psh cmtproto.PartSetHeader) cmtproto.CanonicalPartSetHeader {
	_ = "STUB: not implemented"
	return *new(cmtproto.CanonicalPartSetHeader)
}

// CanonicalizeVote transforms the given Proposal to a CanonicalProposal.
func CanonicalizeProposal(chainID string, proposal *cmtproto.Proposal) cmtproto.CanonicalProposal {
	_ = "STUB: not implemented"
	return *new(cmtproto.CanonicalProposal)
}

// encoded as sfixed64
// encoded as sfixed64

// CanonicalizeVote transforms the given Vote to a CanonicalVote, which does
// not contain ValidatorIndex and ValidatorAddress fields, or any fields
// relating to vote extensions.
func CanonicalizeVote(chainID string, vote *cmtproto.Vote) cmtproto.CanonicalVote {
	_ = "STUB: not implemented"
	return *new(cmtproto.CanonicalVote)
}

// encoded as sfixed64
// encoded as sfixed64

// CanonicalizeVoteExtension extracts the vote extension from the given vote
// and constructs a CanonicalizeVoteExtension struct, whose representation in
// bytes is what is signed in order to produce the vote extension's signature.
func CanonicalizeVoteExtension(chainID string, vote *cmtproto.Vote) cmtproto.CanonicalVoteExtension {
	_ = "STUB: not implemented"
	return *new(cmtproto.CanonicalVoteExtension)
}

// CanonicalTime can be used to stringify time in a canonical way.
func CanonicalTime(t time.Time) string {
	_ = "STUB: not implemented"
	// Note that sending time over amino resets it to
	// local time, we need to force UTC here, so the
	// signatures match
	return ""
}
