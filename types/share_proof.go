package types

import (
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
)

// ShareProof is an NMT proof that a set of shares exist in a set of rows and a
// Merkle proof that those rows exist in a Merkle tree with a given data root.
type ShareProof struct {
	// Data are the raw shares that are being proven.
	Data [][]byte `json:"data"`
	// ShareProofs are NMT proofs that the shares in Data exist in a set of
	// rows. There will be one ShareProof per row that the shares occupy.
	ShareProofs []*tmproto.NMTProof `json:"share_proofs"`
	// NamespaceID is the namespace id of the shares being proven. This
	// namespace id is used when verifying the proof. If the namespace id doesn't
	// match the namespace of the shares, the proof will fail verification.
	NamespaceID      []byte   `json:"namespace_id"`
	RowProof         RowProof `json:"row_proof"`
	NamespaceVersion uint32   `json:"namespace_version"`
}

func (sp ShareProof) ToProto() tmproto.ShareProof {
	_ = "STUB: not implemented"
	// TODO consider extracting a ToProto function for RowProof
	return *new(tmproto.ShareProof)
}

// ShareProofFromProto creates a ShareProof from a proto message.
// Expects the proof to be pre-validated.
func ShareProofFromProto(pb tmproto.ShareProof) (ShareProof, error) {
	_ = "STUB: not implemented"
	return *new(ShareProof), nil
}

// Validate runs basic validations on the proof then verifies if it is consistent.
// It returns nil if the proof is valid. Otherwise, it returns a sensible error.
// The `root` is the block data root that the shares to be proven belong to.
// Note: these proofs are tested on the app side.
func (sp ShareProof) Validate(root []byte) error { _ = "STUB: not implemented"; return nil }

// the range is not inclusive from the left.

func (sp ShareProof) VerifyProof() bool { _ = "STUB: not implemented"; return false }

// Consider extracting celestia-app's namespace package. We can't use it
// here because that would introduce a circulcar import.
//nolint:gosec
