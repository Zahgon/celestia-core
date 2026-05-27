package types

import (
	"github.com/cometbft/cometbft/crypto/merkle"
	tmbytes "github.com/cometbft/cometbft/libs/bytes"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
)

// RowProof is a Merkle proof that a set of rows exist in a Merkle tree with a
// given data root.
type RowProof struct {
	// RowRoots are the roots of the rows being proven.
	RowRoots []tmbytes.HexBytes `json:"row_roots"`
	// Proofs is a list of Merkle proofs where each proof proves that a row
	// exists in a Merkle tree with a given data root.
	Proofs []*merkle.Proof `json:"proofs"`
	// StartRow the index of the start row.
	// Note: currently, StartRow is not validated as part of the proof verification.
	// If this field is used downstream, Validate(root) should be called along with
	// extra validation depending on how it's used.
	StartRow uint32 `json:"start_row"`
	// EndRow the index of the end row.
	// Note: currently, EndRow is not validated as part of the proof verification.
	// If this field is used downstream, Validate(root) should be called along with
	// extra validation depending on how it's used.
	EndRow uint32 `json:"end_row"`
}

// Validate performs checks on the fields of this RowProof. Returns an error if
// the proof fails validation. If the proof passes validation, this function
// attempts to verify the proof. It returns nil if the proof is valid.
func (rp RowProof) Validate(root []byte) error { _ = "STUB: not implemented"; return nil }

// Use uint64 arithmetic to prevent uint32 overflow when computing the
// expected number of rows (e.g. StartRow=0, EndRow=MaxUint32).

// VerifyProof verifies that all the row roots in this RowProof exist in a
// Merkle tree with the given root. Returns true if all proofs are valid.
func (rp RowProof) VerifyProof(root []byte) bool { _ = "STUB: not implemented"; return false }

func RowProofFromProto(p *tmproto.RowProof) RowProof {
	_ = "STUB: not implemented"
	return *new(RowProof)
}
