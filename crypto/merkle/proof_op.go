package merkle

import (
	cmtcrypto "github.com/cometbft/cometbft/proto/tendermint/crypto"
)

//----------------------------------------
// ProofOp gets converted to an instance of ProofOperator:

// ProofOperator is a layer for calculating intermediate Merkle roots
// when a series of Merkle trees are chained together.
// Run() takes leaf values from a tree and returns the Merkle
// root for the corresponding tree. It takes and returns a list of bytes
// to allow multiple leaves to be part of a single proof, for instance in a range proof.
// ProofOp() encodes the ProofOperator in a generic way so it can later be
// decoded with OpDecoder.
type ProofOperator interface {
	Run([][]byte) ([][]byte, error)
	GetKey() []byte
	ProofOp() cmtcrypto.ProofOp
}

//----------------------------------------
// Operations on a list of ProofOperators

// ProofOperators is a slice of ProofOperator(s).
// Each operator will be applied to the input value sequentially
// and the last Merkle root will be verified with already known data
type ProofOperators []ProofOperator

func (poz ProofOperators) VerifyValue(root []byte, keypath string, value []byte) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (poz ProofOperators) Verify(root []byte, keypath string, args [][]byte) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// VerifyFromKeys performs the same verification logic as the normal Verify
// method, except it does not perform any processing on the keypath. This is
// useful when using keys that have split or escape points as a part of the key.
func (poz ProofOperators) VerifyFromKeys(root []byte, keys [][]byte, args [][]byte) (err error) {
	_ = "STUB: not implemented"
	return nil
}

//----------------------------------------
// ProofRuntime - main entrypoint

type OpDecoder func(cmtcrypto.ProofOp) (ProofOperator, error)

type ProofRuntime struct {
	decoders map[string]OpDecoder
}

func NewProofRuntime() *ProofRuntime { _ = "STUB: not implemented"; return nil }

func (prt *ProofRuntime) RegisterOpDecoder(typ string, dec OpDecoder) {
	_ = "STUB: not implemented"
	return
}

func (prt *ProofRuntime) Decode(pop cmtcrypto.ProofOp) (ProofOperator, error) {
	_ = "STUB: not implemented"
	return *new(ProofOperator), nil
}

func (prt *ProofRuntime) DecodeProof(proof *cmtcrypto.ProofOps) (ProofOperators, error) {
	_ = "STUB: not implemented"
	return *new(ProofOperators), nil
}

func (prt *ProofRuntime) VerifyValue(proof *cmtcrypto.ProofOps, root []byte, keypath string, value []byte) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (prt *ProofRuntime) VerifyValueFromKeys(proof *cmtcrypto.ProofOps, root []byte, keys [][]byte, value []byte) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// TODO In the long run we'll need a method of classifcation of ops,
// whether existence or absence or perhaps a third?
func (prt *ProofRuntime) VerifyAbsence(proof *cmtcrypto.ProofOps, root []byte, keypath string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (prt *ProofRuntime) Verify(proof *cmtcrypto.ProofOps, root []byte, keypath string, args [][]byte) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// DefaultProofRuntime only knows about value proofs.
// To use e.g. IAVL proofs, register op-decoders as
// defined in the IAVL package.
func DefaultProofRuntime() (prt *ProofRuntime) { _ = "STUB: not implemented"; return nil }

// VerifyFromKeys performs the same verification logic as the normal Verify
// method, except it does not perform any processing on the keypath. This is
// useful when using keys that have split or escape points as a part of the key.
func (prt *ProofRuntime) VerifyFromKeys(proof *cmtcrypto.ProofOps, root []byte, keys [][]byte, args [][]byte) (err error) {
	_ = "STUB: not implemented"
	return nil
}
