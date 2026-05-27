package types

import (
	"sync"

	"github.com/cometbft/cometbft/crypto/merkle"
	"github.com/cometbft/cometbft/libs/bits"
	protoprop "github.com/cometbft/cometbft/proto/tendermint/propagation"
	"github.com/cometbft/cometbft/types"
)

const (
	ParityRatio = 2
)

// TxMetaData keeps track of the hash of a transaction and its location within the
// protobuf encoded block.
// Range is [start, end).
type TxMetaData struct {
	Hash  []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Start uint32 `protobuf:"varint,2,opt,name=start,proto3" json:"start,omitempty"`
	End   uint32 `protobuf:"varint,3,opt,name=end,proto3" json:"end,omitempty"`
}

// ToProto converts TxMetaData to its protobuf representation.
func (t *TxMetaData) ToProto() *protoprop.TxMetaData { _ = "STUB: not implemented"; return nil }

// TxMetaDataFromProto converts a protobuf TxMetaData to its Go representation.
func TxMetaDataFromProto(t *protoprop.TxMetaData) *TxMetaData {
	_ = "STUB: not implemented"
	return nil
}

// ValidateBasic checks if the TxMetaData is valid. It fails if Start > End or
// if the hash is invalid.
func (t *TxMetaData) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// CompactBlock contains commitments and metadata for reusing transactions that
// have already been distributed.
type CompactBlock struct {
	// BpHash is the block propagation hash. It's the root of the mekle tree where the leaves are the concatenation of the original partset elements and the parity one.
	BpHash    []byte         `json:"bp_hash,omitempty"`
	Blobs     []TxMetaData   `json:"blobs,omitempty"`
	Signature []byte         `json:"signature,omitempty"`
	Proposal  types.Proposal `json:"proposal,omitempty"`
	// length of the last part
	LastLen uint32 `json:"last_len,omitempty"`
	// the original + parity part set parts hashes.
	PartsHashes [][]byte `json:"parts_hashes,omitempty"`

	mtx sync.Mutex
	// proofsCache is local storage from generated proofs from the PartsHashes.
	// It must not be included in any serialization.
	proofsCache []*merkle.Proof
}

// SignBytes returns the compact block commitment data that
// needs to be signed.
// The sign bytes are the field-delimited protobuf encoding of the compact block.
func (c *CompactBlock) SignBytes() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

//nolint:prealloc

// ValidateBasic checks if the CompactBlock is valid. It fails if the height is
// negative, if the round is negative, if the BpHash is invalid, or if any of
// the Blobs are invalid.
func (c *CompactBlock) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// validate tx metadata

// hasOverlappingRanges checks whether any ranges in the provided slice of TxMetaData overlap.
// Returns an error if overlapping ranges are found, otherwise returns nil.
func hasOverlappingRanges(blobs []TxMetaData) error { _ = "STUB: not implemented"; return nil }

// Create a copy of the blobs slice to avoid mutating the original

// If current range starts before previous range ends, there's an overlap
// using < instead of <= because the ranges are [start:end)

// ToProto converts CompactBlock to its protobuf representation.
func (c *CompactBlock) ToProto() *protoprop.CompactBlock { _ = "STUB: not implemented"; return nil }

// Proofs returns the proofs to each part. If the proofs are not already
// generated, then they are done so during the first call. An error is only
// thrown if the proofs are generated and the resulting hashes don't match those
// in the compact block. This method should be called upon first receiving a
// compact block.
func (c *CompactBlock) Proofs() ([]*merkle.Proof, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *CompactBlock) GetProof(i uint32) *merkle.Proof { _ = "STUB: not implemented"; return nil }

func (c *CompactBlock) SetProofCache(proofs []*merkle.Proof) { _ = "STUB: not implemented"; return }

// CompactBlockFromProto converts a protobuf CompactBlock to its Go representation.
func CompactBlockFromProto(c *protoprop.CompactBlock) (*CompactBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PartMetaData keeps track of the hash of each part, its location via the
// index, along with the proof of inclusion to either the PartSetHeader hash or
// the BPRoot in the CompactBlock.
type PartMetaData struct {
	Index uint32 `json:"index,omitempty"`
	Hash  []byte `json:"hash,omitempty"`
}

// ValidateBasic checks if the PartMetaData is valid. It fails if the hash or
// the proof is invalid.
func (p *PartMetaData) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// HaveParts is the go representation of the wire message for determining the
// route of parts.
type HaveParts struct {
	Height int64          `json:"height,omitempty"`
	Round  int32          `json:"round,omitempty"`
	Parts  []PartMetaData `json:"parts,omitempty"`
}

// BitArrary returns a bit array of the provided size with the indexes of the
// parts set to true.
func (h *HaveParts) BitArray(size int) *bits.BitArray { _ = "STUB: not implemented"; return nil }

// ValidateBasic checks if the HaveParts is valid. It fails if Parts is nil or
// empty, or if any of the parts are invalid.
func (h *HaveParts) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// ValidatePartHashes verifies that each part's hash in the HaveParts struct matches the corresponding expected hash.
// Returns an error if any hash does not match, indicating the index of the first mismatch.
func (h *HaveParts) ValidatePartHashes(expectedHashes [][]byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *HaveParts) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func (h *HaveParts) GetIndex(i uint32) bool { _ = "STUB: not implemented"; return false }

// ToProto converts HaveParts to its protobuf representation.
func (h *HaveParts) ToProto() *protoprop.HaveParts { _ = "STUB: not implemented"; return nil }

// HavePartFromProto converts a protobuf HaveParts to its Go representation.
func HavePartFromProto(h *protoprop.HaveParts) (*HaveParts, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// WantParts is a message that requests a set of parts from a peer.
type WantParts struct {
	Parts             *bits.BitArray `json:"parts"`
	Height            int64          `json:"height,omitempty"`
	Round             int32          `json:"round,omitempty"`
	Prove             bool           `json:"prove,omitempty"`
	MissingPartsCount int32          `json:"missing_parts_count,omitempty"`
}

func (w *WantParts) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// ToProto converts WantParts to its protobuf representation.
func (w *WantParts) ToProto() *protoprop.WantParts { _ = "STUB: not implemented"; return nil }

// WantPartsFromProto converts a protobuf WantParts to its Go representation.
func WantPartsFromProto(w *protoprop.WantParts) (*WantParts, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type RecoveryPart struct {
	Height int64
	Round  int32
	Index  uint32
	Data   []byte
	Proof  *merkle.Proof
}

func (p *RecoveryPart) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

func RecoveryPartFromProto(r *protoprop.RecoveryPart) (*RecoveryPart, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MsgFromProto takes a consensus proto message and returns the native go type
func MsgFromProto(p *protoprop.Message) (Message, error) {
	_ = "STUB: not implemented"
	return *new(Message), nil
}

// Message is a message that can be sent and received on the Reactor
type Message interface {
	ValidateBasic() error
}
