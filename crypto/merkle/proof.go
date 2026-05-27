package merkle

import (
	cmtcrypto "github.com/cometbft/cometbft/proto/tendermint/crypto"
)

const (
	// MaxAunts is the maximum number of aunts that can be included in a Proof.
	// This corresponds to a tree of size 2^100, which should be sufficient for all conceivable purposes.
	// This maximum helps prevent Denial-of-Service attacks by limitting the size of the proofs.
	MaxAunts = 100
)

// Proof represents a Merkle proof.
// NOTE: The convention for proofs is to include leaf hashes but to
// exclude the root hash.
// This convention is implemented across IAVL range proofs as well.
// Keep this consistent unless there's a very good reason to change
// everything.  This also affects the generalized proof system as
// well.
type Proof struct {
	Total    int64    `json:"total"`           // Total number of items.
	Index    int64    `json:"index"`           // Index of item to prove.
	LeafHash []byte   `json:"leaf_hash"`       // Hash of item value.
	Aunts    [][]byte `json:"aunts,omitempty"` // Hashes from leaf's sibling to a root's child.
}

// ProofsFromByteSlices computes inclusion proof for given items.
// proofs[0] is the proof for items[0].
func ProofsFromByteSlices(items [][]byte) (rootHash []byte, proofs []*Proof) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Verify that the Proof proves the root hash.
// Check sp.Index/sp.Total manually if needed
func (sp *Proof) Verify(rootHash []byte, leaf []byte) error { _ = "STUB: not implemented"; return nil }

// Compute the root hash given a leaf hash.  Panics in case of errors.
func (sp *Proof) ComputeRootHash() []byte { _ = "STUB: not implemented"; return nil }

// Compute the root hash given a leaf hash.
func (sp *Proof) computeRootHash() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// String implements the stringer interface for Proof.
// It is a wrapper around StringIndented.
func (sp *Proof) String() string { _ = "STUB: not implemented"; return "" }

// StringIndented generates a canonical string representation of a Proof.
func (sp *Proof) StringIndented(indent string) string { _ = "STUB: not implemented"; return "" }

// ValidateBasic performs basic validation.
// NOTE: it expects the LeafHash and the elements of Aunts to be of size tmhash.Size,
// and it expects at most MaxAunts elements in Aunts.
func (sp *Proof) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

func (sp *Proof) ToProto() *cmtcrypto.Proof { _ = "STUB: not implemented"; return nil }

func ProofFromProto(pb *cmtcrypto.Proof, optional bool) (*Proof, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Use the leafHash and innerHashes to get the root merkle hash.
// If the length of the innerHashes slice isn't exactly correct, the result is nil.
// Recursive impl.
func computeHashFromAunts(index, total int64, leafHash []byte, innerHashes [][]byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ProofNode is a helper structure to construct merkle proof.
// The node and the tree is thrown away afterwards.
// Exactly one of node.Left and node.Right is nil, unless node is the root, in which case both are nil.
// node.Parent.Hash = hash(node.Hash, node.Right.Hash) or
// hash(node.Left.Hash, node.Hash), depending on whether node is a left/right child.
type ProofNode struct {
	Hash   []byte
	Parent *ProofNode
	Left   *ProofNode // Left sibling  (only one of Left,Right is set)
	Right  *ProofNode // Right sibling (only one of Left,Right is set)
}

// FlattenAunts will return the inner hashes for the item corresponding to the leaf,
// starting from a leaf ProofNode.
func (spn *ProofNode) FlattenAunts() [][]byte {
	_ = "STUB: not implemented"
	// Nonrecursive impl.
	return nil
}

// count the length of the path from the root to the leaf.

// trails[0].Hash is the leaf hash for items[0].
// trails[i].Parent.Parent....Parent == root for all i.
func trailsFromByteSlices(items [][]byte) (trails []*ProofNode, root *ProofNode) {
	_ = "STUB: not implemented"
	// Recursive impl.
	return nil, nil
}

func trailsFromLeafHashes(leafHashes [][]byte) (trails []*ProofNode, root *ProofNode) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ProofsFromLeafHashes(leafHashes [][]byte) (rootHash []byte, proofs []*Proof) {
	_ = "STUB: not implemented"
	return nil, nil
}
