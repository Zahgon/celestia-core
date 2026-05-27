package merkle

import (
	cmtcrypto "github.com/cometbft/cometbft/proto/tendermint/crypto"
)

const ProofOpValue = "simple:v"

// ValueOp takes a key and a single value as argument and
// produces the root hash.  The corresponding tree structure is
// the SimpleMap tree.  SimpleMap takes a Hasher, and currently
// CometBFT uses tmhash.  SimpleValueOp should support
// the hash function as used in tmhash.  TODO support
// additional hash functions here as options/args to this
// operator.
//
// If the produced root hash matches the expected hash, the
// proof is good.
type ValueOp struct {
	// Encoded in ProofOp.Key.
	key []byte

	// To encode in ProofOp.Data
	Proof *Proof `json:"proof"`
}

var _ ProofOperator = ValueOp{}

func NewValueOp(key []byte, proof *Proof) ValueOp { _ = "STUB: not implemented"; return *new(ValueOp) }

func ValueOpDecoder(pop cmtcrypto.ProofOp) (ProofOperator, error) {
	_ = "STUB: not implemented"
	return *new(ProofOperator), nil
}

// a bit strange as we'll discard this, but it works.

func (op ValueOp) ProofOp() cmtcrypto.ProofOp {
	_ = "STUB: not implemented"
	return *new(cmtcrypto.ProofOp)
}

func (op ValueOp) String() string { _ = "STUB: not implemented"; return "" }

func (op ValueOp) Run(args [][]byte) ([][]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Wrap <op.Key, vhash> to hash the KVPair.
//nolint: errcheck // does not error
//nolint: errcheck // does not error

func (op ValueOp) GetKey() []byte { _ = "STUB: not implemented"; return nil }
