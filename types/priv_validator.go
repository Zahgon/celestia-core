package types

import (
	"errors"

	cmtbytes "github.com/cometbft/cometbft/libs/bytes"

	"github.com/cometbft/cometbft/crypto"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
)

// PrivValidator defines the functionality of a local CometBFT validator
// that signs votes and proposals, and never double signs.
type PrivValidator interface {
	GetPubKey() (crypto.PubKey, error)

	SignVote(chainID string, vote *cmtproto.Vote) error
	SignProposal(chainID string, proposal *cmtproto.Proposal) error
	SignRawBytes(chainID, uniqueID string, rawBytes []byte) ([]byte, error)
}

// RawBytesSignBytesPrefix defines a domain separator prefix added to raw bytes to ensure the resulting
// signed message can't be confused with a consensus message, which could lead to double signing
const RawBytesSignBytesPrefix = "COMET::RAW_BYTES::SIGN"

// RawBytesMessageSignBytes returns the canonical bytes for signing raw data messages.
// It requires non-empty chainID, uniqueID, and rawBytes to prevent security issues.
// Returns error if any required parameter is empty or if marshaling fails.
func RawBytesMessageSignBytes(chainID, uniqueID string, rawBytes []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type PrivValidatorsByAddress []PrivValidator

func (pvs PrivValidatorsByAddress) Len() int { _ = "STUB: not implemented"; return 0 }

func (pvs PrivValidatorsByAddress) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (pvs PrivValidatorsByAddress) Swap(i, j int) { _ = "STUB: not implemented"; return }

func P2PMessageSignBytes(chainID, uID string, hash cmtbytes.HexBytes) []byte {
	_ = "STUB: not implemented"
	return nil
}

//----------------------------------------
// MockPV

const (
	MockChainID = "incorrect-chain-id"
)

var _ PrivValidator = &MockPV{}

// MockPV implements PrivValidator without any safety or persistence.
// Only use it for testing.
type MockPV struct {
	PrivKey              crypto.PrivKey
	breakProposalSigning bool
	breakVoteSigning     bool
}

func NewMockPV() MockPV { _ = "STUB: not implemented"; return *new(MockPV) }

// NewMockPVWithParams allows one to create a MockPV instance, but with finer
// grained control over the operation of the mock validator. This is useful for
// mocking test failures.
func NewMockPVWithParams(privKey crypto.PrivKey, breakProposalSigning, breakVoteSigning bool) MockPV {
	_ = "STUB: not implemented"
	return *new(MockPV)
}

// Implements PrivValidator.
func (pv MockPV) GetPubKey() (crypto.PubKey, error) {
	_ = "STUB: not implemented"
	return *new(crypto.PubKey), nil
}

// Implements PrivValidator.
func (pv MockPV) SignVote(chainID string, vote *cmtproto.Vote) error {
	_ = "STUB: not implemented"
	return nil
}

// We only sign vote extensions for non-nil precommits

// Implements PrivValidator.
func (pv MockPV) SignProposal(chainID string, proposal *cmtproto.Proposal) error {
	_ = "STUB: not implemented"
	return nil
}

func (pv MockPV) SignRawBytes(chainID, uniqueID string, rawBytes []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (pv MockPV) ExtractIntoValidator(votingPower int64) *Validator {
	_ = "STUB: not implemented"
	return nil
}

// String returns a string representation of the MockPV.
func (pv MockPV) String() string { _ = "STUB: not implemented"; return "" }

// mockPV will never return an error, ignored here

// XXX: Implement.
func (pv MockPV) DisableChecks() {
	_ = "STUB: not implemented"
	// Currently this does nothing,
	// as MockPV has no safety checks at all.
	return
}

type ErroringMockPV struct {
	MockPV
}

var ErroringMockPVErr = errors.New("erroringMockPV always returns an error")

// Implements PrivValidator.
func (pv *ErroringMockPV) SignVote(string, *cmtproto.Vote) error {
	_ = "STUB: not implemented"
	return nil

	// Implements PrivValidator.
}

func (pv *ErroringMockPV) SignProposal(string, *cmtproto.Proposal) error {
	_ = "STUB: not implemented"
	return nil
}

func (pv *ErroringMockPV) SignP2PMessage(chainID, uID string, hash cmtbytes.HexBytes) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Implements PrivValidator.
}

func (pv *ErroringMockPV) SignRawBytes(chainID, uniqueID string, rawBytes []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// NewErroringMockPV returns a MockPV that fails on each signing request. Again, for testing only.
}

func NewErroringMockPV() *ErroringMockPV { _ = "STUB: not implemented"; return nil }
