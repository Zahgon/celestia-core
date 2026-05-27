package types

import (
	"github.com/cometbft/cometbft/crypto"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
)

// Volatile state for each Validator
// NOTE: The ProposerPriority is not included in Validator.Hash();
// make sure to update that method if changes are made here
type Validator struct {
	Address     Address       `json:"address"`
	PubKey      crypto.PubKey `json:"pub_key"`
	VotingPower int64         `json:"voting_power"`

	ProposerPriority int64 `json:"proposer_priority"`
}

// NewValidator returns a new validator with the given pubkey and voting power.
func NewValidator(pubKey crypto.PubKey, votingPower int64) *Validator {
	_ = "STUB: not implemented"
	return nil
}

// ValidateBasic performs basic validation.
func (v *Validator) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// Creates a new copy of the validator so we can mutate ProposerPriority.
// Panics if the validator is nil.
func (v *Validator) Copy() *Validator { _ = "STUB: not implemented"; return nil }

// Returns the one with higher ProposerPriority.
func (v *Validator) CompareProposerPriority(other *Validator) *Validator {
	_ = "STUB: not implemented"
	return nil
}

// String returns a string representation of String.
//
// 1. address
// 2. public key
// 3. voting power
// 4. proposer priority
func (v *Validator) String() string { _ = "STUB: not implemented"; return "" }

// ValidatorListString returns a prettified validator list for logging purposes.
func ValidatorListString(vals []*Validator) string { _ = "STUB: not implemented"; return "" }

// Bytes computes the unique encoding of a validator with a given voting power.
// These are the bytes that gets hashed in consensus. It excludes address
// as its redundant with the pubkey. This also excludes ProposerPriority
// which changes every round.
func (v *Validator) Bytes() []byte { _ = "STUB: not implemented"; return nil }

// ToProto converts Valiator to protobuf
func (v *Validator) ToProto() (*cmtproto.Validator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FromProto sets a protobuf Validator to the given pointer.
// It returns an error if the public key is invalid.
func ValidatorFromProto(vp *cmtproto.Validator) (*Validator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//----------------------------------------
// RandValidator

// RandValidator returns a randomized validator, useful for testing.
// UNSTABLE
func RandValidator(randPower bool, minPower int64) (*Validator, PrivValidator) {
	_ = "STUB: not implemented"
	return nil, *new(PrivValidator)
}
