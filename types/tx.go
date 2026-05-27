package types

import (
	"crypto/sha256"

	"github.com/cometbft/cometbft/crypto/merkle"
	cmtbytes "github.com/cometbft/cometbft/libs/bytes"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
)

// TxKeySize is the size of the transaction key index
const TxKeySize = sha256.Size

type (
	// Tx is an arbitrary byte array.
	// NOTE: Tx has no types at this level, so when wire encoded it's just length-prefixed.
	// Might we want types here ?
	Tx []byte

	// TxKey is the fixed length array key used as an index.
	TxKey [TxKeySize]byte
)

func (tx TxKey) String() string { _ = "STUB: not implemented"; return "" }

// Hash computes the TMHASH hash of the wire encoded transaction.
type CachedTx struct {
	Tx
	hash []byte
}

// Hash returns the cached hash if available, otherwise it computes the hash
// using the normal Tx.Hash method.
func (tx *CachedTx) Hash() []byte { _ = "STUB: not implemented"; return nil }

// Key returns the cached key if available, otherwise it computes the key
// using the normal Tx.Key method.
func (tx *CachedTx) Key() TxKey { _ = "STUB: not implemented"; return *new(TxKey) }

// NewCachedTx creates a new CachedTx with the provided transaction and hash.
func NewCachedTx(tx Tx, hash []byte) *CachedTx { _ = "STUB: not implemented"; return nil }

func TxsFromCachedTxs(cachedTxs []*CachedTx) Txs { _ = "STUB: not implemented"; return *new(Txs) }

// CachedTxFromTxs creates a slice of CachedTx from a slice of Tx.
func CachedTxFromTxs(txs Txs) []*CachedTx { _ = "STUB: not implemented"; return nil }

func CachedTxToSliceOfBytes(cachedTxs []*CachedTx) [][]byte { _ = "STUB: not implemented"; return nil }

// Hash computes the TMHASH hash of the wire encoded transaction. It attempts to
// unwrap the transaction if it is a IndexWrapper or a BlobTx.
func (tx Tx) Hash() []byte { _ = "STUB: not implemented"; return nil }

func (tx Tx) ToCachedTx() *CachedTx { _ = "STUB: not implemented"; return nil }

// Key returns the sha256 hash of the wire encoded transaction. It attempts to
// unwrap the transaction if it is a BlobTx or a IndexWrapper.
func (tx Tx) Key() TxKey { _ = "STUB: not implemented"; return *new(TxKey) }

// String returns the hex-encoded transaction as a string.
func (tx Tx) String() string { _ = "STUB: not implemented"; return "" }

func TxKeyFromBytes(bytes []byte) (TxKey, error) {
	_ = "STUB: not implemented"
	return *new(TxKey), nil
}

// Txs is a slice of Tx.
type Txs []Tx

// Hash returns the Merkle root hash of the transaction hashes.
// i.e. the leaves of the tree are the hashes of the txs.
func (txs Txs) Hash() []byte { _ = "STUB: not implemented"; return nil }

// Index returns the index of this transaction in the list, or -1 if not found
func (txs Txs) Index(tx Tx) int { _ = "STUB: not implemented"; return 0 }

// IndexByHash returns the index of this transaction hash in the list, or -1 if not found
func (txs Txs) IndexByHash(hash []byte) int { _ = "STUB: not implemented"; return 0 }

func (txs Txs) Proof(i int) TxProof { _ = "STUB: not implemented"; return *new(TxProof) }

func (txs Txs) hashList() [][]byte { _ = "STUB: not implemented"; return nil }

// Txs is a slice of transactions. Sorting a Txs value orders the transactions
// lexicographically.
func (txs Txs) Len() int           { _ = "STUB: not implemented"; return 0 }
func (txs Txs) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (txs Txs) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func ToTxs(txl [][]byte) Txs { _ = "STUB: not implemented"; return *new(Txs) }

func (txs Txs) Validate(maxSizeBytes int64) error { _ = "STUB: not implemented"; return nil }

// ToSliceOfBytes converts a Txs to slice of byte slices.
func (txs Txs) ToSliceOfBytes() [][]byte { _ = "STUB: not implemented"; return nil }

// TxProof represents a Merkle proof of the presence of a transaction in the Merkle tree.
type TxProof struct {
	RootHash cmtbytes.HexBytes `json:"root_hash"`
	Data     Tx                `json:"data"`
	Proof    merkle.Proof      `json:"proof"`
}

// Leaf returns the hash(tx), which is the leaf in the merkle tree which this proof refers to.
func (tp TxProof) Leaf() []byte { _ = "STUB: not implemented"; return nil }

// Validate verifies the proof. It returns nil if the RootHash matches the dataHash argument,
// and if the proof is internally consistent. Otherwise, it returns a sensible error.
func (tp TxProof) Validate(dataHash []byte) error { _ = "STUB: not implemented"; return nil }

func (tp TxProof) ToProto() cmtproto.TxProof {
	_ = "STUB: not implemented"
	return *new(cmtproto.TxProof)
}

func TxProofFromProto(pb cmtproto.TxProof) (TxProof, error) {
	_ = "STUB: not implemented"
	return *new(TxProof), nil
}

// ComputeProtoSizeForTxs wraps the transactions in cmtproto.Data{} and calculates the size.
// https://developers.google.com/protocol-buffers/docs/encoding
func ComputeProtoSizeForTxs(txs []Tx) int64 { _ = "STUB: not implemented"; return 0 }

// UnmarshalIndexWrapper attempts to unmarshal the provided transaction into an
// IndexWrapper transaction. It returns true if the provided transaction is an
// IndexWrapper transaction. An IndexWrapper transaction is a transaction that contains
// a MsgPayForBlob that has been wrapped with a share index.
//
// NOTE: protobuf sometimes does not throw an error if the transaction passed is
// not a tmproto.IndexWrapper, since the protobuf definition for MsgPayForBlob is
// kept in the app, we cannot perform further checks without creating an import
// cycle.
func UnmarshalIndexWrapper(tx Tx) (indexWrapper cmtproto.IndexWrapper, isIndexWrapper bool) {
	_ = "STUB: not implemented"
	// attempt to unmarshal into an IndexWrapper transaction
	return *new(cmtproto.IndexWrapper), false
}

// MarshalIndexWrapper creates a wrapped Tx that includes the original transaction
// and the share index of the start of its blob.
//
// NOTE: must be unwrapped to be a viable sdk.Tx.
func MarshalIndexWrapper(tx Tx, shareIndexes ...uint32) (Tx, error) {
	_ = "STUB: not implemented"
	return *new(Tx), nil
}

// UnmarshalBlobTx attempts to unmarshal a transaction into blob transaction. If an
// error is thrown, false is returned.
func UnmarshalBlobTx(tx Tx) (bTx cmtproto.BlobTx, isBlob bool) {
	_ = "STUB: not implemented"
	return *new(cmtproto.BlobTx), false
}

// perform some quick basic checks to prevent false positives

// MarshalBlobTx creates a BlobTx using a normal transaction and some number of
// blobs.
//
// NOTE: Any checks on the blobs or the transaction must be performed in the
// application.
func MarshalBlobTx(tx []byte, blobs ...*cmtproto.Blob) (Tx, error) {
	_ = "STUB: not implemented"
	return *new(Tx), nil
}
