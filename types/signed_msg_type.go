package types

import cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"

// IsVoteTypeValid returns true if t is a valid vote type.
func IsVoteTypeValid(t cmtproto.SignedMsgType) bool { _ = "STUB: not implemented"; return false }

var signedMsgTypeToShortName = map[cmtproto.SignedMsgType]string{
	cmtproto.UnknownType:   "unknown",
	cmtproto.PrevoteType:   "prevote",
	cmtproto.PrecommitType: "precommit",
	cmtproto.ProposalType:  "proposal",
}

// Returns a short lowercase descriptor for a signed message type.
func SignedMsgTypeToShortString(t cmtproto.SignedMsgType) string {
	_ = "STUB: not implemented"
	return ""
}
