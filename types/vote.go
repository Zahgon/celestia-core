package types

import (
	"errors"
	"time"

	"github.com/cometbft/cometbft/crypto"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
)

const (
	nilVoteStr string = "nil-Vote"

	// The maximum supported number of bytes in a vote extension.
	MaxVoteExtensionSize int = 1024 * 1024
)

var (
	ErrVoteUnexpectedStep            = errors.New("unexpected step")
	ErrVoteInvalidValidatorIndex     = errors.New("invalid validator index")
	ErrVoteInvalidValidatorAddress   = errors.New("invalid validator address")
	ErrVoteInvalidSignature          = errors.New("invalid signature")
	ErrVoteInvalidBlockHash          = errors.New("invalid block hash")
	ErrVoteNonDeterministicSignature = errors.New("non-deterministic signature")
	ErrVoteNil                       = errors.New("nil vote")
	ErrVoteExtensionAbsent           = errors.New("vote extension absent")
	ErrInvalidVoteExtension          = errors.New("invalid vote extension")
)

type ErrVoteConflictingVotes struct {
	VoteA *Vote
	VoteB *Vote
}

func (err *ErrVoteConflictingVotes) Error() string { _ = "STUB: not implemented"; return "" }

func NewConflictingVoteError(vote1, vote2 *Vote) *ErrVoteConflictingVotes {
	_ = "STUB: not implemented"
	return nil
}

// The vote extension is only valid for non-nil precommits.
type ErrVoteExtensionInvalid struct {
	ExtSignature []byte
}

func (err *ErrVoteExtensionInvalid) Error() string { _ = "STUB: not implemented"; return "" }

// Address is hex bytes.
type Address = crypto.Address

// Vote represents a prevote, precommit, or commit vote from validators for
// consensus.
type Vote struct {
	Type               cmtproto.SignedMsgType `json:"type"`
	Height             int64                  `json:"height"`
	Round              int32                  `json:"round"`    // assume there will not be greater than 2_147_483_647 rounds
	BlockID            BlockID                `json:"block_id"` // zero if vote is nil.
	Timestamp          time.Time              `json:"timestamp"`
	ValidatorAddress   Address                `json:"validator_address"`
	ValidatorIndex     int32                  `json:"validator_index"`
	Signature          []byte                 `json:"signature"`
	Extension          []byte                 `json:"extension"`
	ExtensionSignature []byte                 `json:"extension_signature"`
}

// VoteFromProto attempts to convert the given serialization (Protobuf) type to
// our Vote domain type. No validation is performed on the resulting vote -
// this is left up to the caller to decide whether to call ValidateBasic or
// ValidateWithExtension.
func VoteFromProto(pv *cmtproto.Vote) (*Vote, error) { _ = "STUB: not implemented"; return nil, nil }

// CommitSig converts the Vote to a CommitSig.
func (vote *Vote) CommitSig() CommitSig { _ = "STUB: not implemented"; return *new(CommitSig) }

// ExtendedCommitSig attempts to construct an ExtendedCommitSig from this vote.
// Panics if either the vote extension signature is missing or if the block ID
// is not either empty or complete.
func (vote *Vote) ExtendedCommitSig() ExtendedCommitSig {
	_ = "STUB: not implemented"
	return *new(ExtendedCommitSig)
}

// VoteSignBytes returns the proto-encoding of the canonicalized Vote, for
// signing. Panics if the marshaling fails.
//
// The encoded Protobuf message is varint length-prefixed (using MarshalDelimited)
// for backwards-compatibility with the Amino encoding, due to e.g. hardware
// devices that rely on this encoding.
//
// See CanonicalizeVote
func VoteSignBytes(chainID string, vote *cmtproto.Vote) []byte {
	_ = "STUB: not implemented"
	return nil
}

// VoteExtensionSignBytes returns the proto-encoding of the canonicalized vote
// extension for signing. Panics if the marshaling fails.
//
// Similar to VoteSignBytes, the encoded Protobuf message is varint
// length-prefixed for backwards-compatibility with the Amino encoding.
func VoteExtensionSignBytes(chainID string, vote *cmtproto.Vote) []byte {
	_ = "STUB: not implemented"
	return nil
}

func (vote *Vote) Copy() *Vote { _ = "STUB: not implemented"; return nil }

// String returns a string representation of Vote.
//
// 1. validator index
// 2. first 6 bytes of validator address
// 3. height
// 4. round,
// 5. type byte
// 6. type string
// 7. first 6 bytes of block hash
// 8. first 6 bytes of signature
// 9. first 6 bytes of vote extension
// 10. timestamp
func (vote *Vote) String() string { _ = "STUB: not implemented"; return "" }

func (vote *Vote) verifyAndReturnProto(chainID string, pubKey crypto.PubKey) (*cmtproto.Vote, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Verify checks whether the signature associated with this vote corresponds to
// the given chain ID and public key. This function does not validate vote
// extension signatures - to do so, use VerifyWithExtension instead.
func (vote *Vote) Verify(chainID string, pubKey crypto.PubKey) error {
	_ = "STUB: not implemented"
	return nil
}

// VerifyVoteAndExtension performs the same verification as Verify, but
// additionally checks whether the vote extension signature corresponds to the
// given chain ID and public key. We only verify vote extension signatures for
// precommits.
func (vote *Vote) VerifyVoteAndExtension(chainID string, pubKey crypto.PubKey) error {
	_ = "STUB: not implemented"
	return nil
}

// We only verify vote extension signatures for non-nil precommits.

// VerifyExtension checks whether the vote extension signature corresponds to the
// given chain ID and public key.
func (vote *Vote) VerifyExtension(chainID string, pubKey crypto.PubKey) error {
	_ = "STUB: not implemented"
	return nil
}

// ValidateBasic checks whether the vote is well-formed. It does not, however,
// check vote extensions - for vote validation with vote extension validation,
// use ValidateWithExtension.
func (vote *Vote) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// NOTE: Timestamp validation is subtle and handled elsewhere.

// BlockID.ValidateBasic would not err if we for instance have an empty hash but a
// non-empty PartsSetHeader:

// We should only ever see vote extensions in non-nil precommits, otherwise
// this is a violation of the specification.
// https://github.com/tendermint/tendermint/issues/8487

// It's possible that this vote has vote extensions but
// they could also be disabled and thus not present thus
// we can't do all checks

// NOTE: extended votes should have a signature regardless of
// of whether there is any data in the extension or not however
// we don't know if extensions are enabled so we can only
// enforce the signature when extension size is not nil

// EnsureExtension checks for the presence of extensions signature data
// on precommit vote types.
func (vote *Vote) EnsureExtension() error {
	_ = "STUB: not implemented"
	// We should always see vote extension signatures in non-nil precommits
	return nil
}

// ToProto converts the handwritten type to proto generated type
// return type, nil if everything converts safely, otherwise nil, error
func (vote *Vote) ToProto() *cmtproto.Vote { _ = "STUB: not implemented"; return nil }

func VotesToProto(votes []*Vote) []*cmtproto.Vote { _ = "STUB: not implemented"; return nil }

// protobuf crashes when serializing "repeated" fields with nil elements

// SignAndCheckVote signs the vote with the given privVal and checks the vote.
// It returns an error if the vote is invalid and a boolean indicating if the
// error is recoverable or not.
func SignAndCheckVote(
	vote *Vote,
	privVal PrivValidator,
	chainID string,
	extensionsEnabled bool,
) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Failing to sign a vote has always been a recoverable error, this
// function keeps it that way.

// Non-recoverable because the caller passed parameters that don't make sense

// Error if prevote contains an extension signature

// Non-recoverable because the vote is malformed

// Error if missing extension signature for non-nil Precommit

// Non-recoverable because the vote is malformed
