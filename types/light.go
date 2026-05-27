package types

import (
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
)

// LightBlock is a SignedHeader and a ValidatorSet.
// It is the basis of the light client
type LightBlock struct {
	*SignedHeader `json:"signed_header"`
	ValidatorSet  *ValidatorSet `json:"validator_set"`
}

// ValidateBasic checks that the data is correct and consistent
//
// This does no verification of the signatures
func (lb LightBlock) ValidateBasic(chainID string) error { _ = "STUB: not implemented"; return nil }

// make sure the validator set is consistent with the header
//nolint:staticcheck

//nolint:staticcheck

// String returns a string representation of the LightBlock
func (lb LightBlock) String() string { _ = "STUB: not implemented"; return "" }

// StringIndented returns an indented string representation of the LightBlock
//
// SignedHeader
// ValidatorSet
func (lb LightBlock) StringIndented(indent string) string { _ = "STUB: not implemented"; return "" }

// ToProto converts the LightBlock to protobuf
func (lb *LightBlock) ToProto() (*cmtproto.LightBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// LightBlockFromProto converts from protobuf back into the Lightblock.
// An error is returned if either the validator set or signed header are invalid
func LightBlockFromProto(pb *cmtproto.LightBlock) (*LightBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//-----------------------------------------------------------------------------

// SignedHeader is a header along with the commits that prove it.
type SignedHeader struct {
	*Header `json:"header"`

	Commit *Commit `json:"commit"`
}

// IsEmpty returns true if both the header and commit are nil.
func (sh SignedHeader) IsEmpty() bool { _ = "STUB: not implemented"; return false }

// ValidateBasic does basic consistency checks and makes sure the header
// and commit are consistent.
//
// NOTE: This does not actually check the cryptographic signatures.  Make sure
// to use a Verifier to validate the signatures actually provide a
// significantly strong proof for this header's validity.
func (sh SignedHeader) ValidateBasic(chainID string) error { _ = "STUB: not implemented"; return nil }

// Make sure the header is consistent with the commit.

//nolint:staticcheck

// String returns a string representation of SignedHeader.
func (sh SignedHeader) String() string { _ = "STUB: not implemented"; return "" }

// StringIndented returns an indented string representation of SignedHeader.
//
// Header
// Commit
func (sh SignedHeader) StringIndented(indent string) string { _ = "STUB: not implemented"; return "" }

// ToProto converts SignedHeader to protobuf
func (sh *SignedHeader) ToProto() *cmtproto.SignedHeader { _ = "STUB: not implemented"; return nil }

// FromProto sets a protobuf SignedHeader to the given pointer.
// It returns an error if the header or the commit is invalid.
func SignedHeaderFromProto(shp *cmtproto.SignedHeader) (*SignedHeader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
