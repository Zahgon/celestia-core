package types

import (
	"errors"

	"github.com/cometbft/cometbft/crypto/merkle"
	"github.com/cometbft/cometbft/libs/bits"
	cmtbytes "github.com/cometbft/cometbft/libs/bytes"
	cmtsync "github.com/cometbft/cometbft/libs/sync"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
)

var (
	ErrPartSetUnexpectedIndex   = errors.New("error part set unexpected index")
	ErrPartSetInvalidProof      = errors.New("error part set invalid proof")
	ErrPartSetInvalidProofHash  = errors.New("error part set invalid proof: wrong hash")
	ErrPartSetInvalidProofTotal = errors.New("error part set invalid proof: wrong total")
	ErrPartTooBig               = errors.New("error part size too big")
	ErrPartInvalidSize          = errors.New("error inner part with invalid size")
)

// ErrInvalidPart is an error type for invalid parts.
type ErrInvalidPart struct {
	Reason error
}

func (e ErrInvalidPart) Error() string { _ = "STUB: not implemented"; return "" }

func (e ErrInvalidPart) Unwrap() error { _ = "STUB: not implemented"; return nil }

type PartInfo struct {
	*Part
	Height int64
	Round  int32
}

type Part struct {
	mtx   cmtsync.Mutex
	Index uint32            `json:"index"`
	Bytes cmtbytes.HexBytes `json:"bytes"`
	Proof merkle.Proof      `json:"proof"`
}

// ValidateBasic performs basic validation.
func (part *Part) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// All parts except the last one should have the same constant size.

// GetProof returns the merkle proof while safely handling concurrent access
func (part *Part) GetProof() merkle.Proof { _ = "STUB: not implemented"; return *new(merkle.Proof) }

// SetProof sets the merkle proof while safely handling concurrent access
func (part *Part) SetProof(proof merkle.Proof) { _ = "STUB: not implemented"; return }

// String returns a string representation of Part.
//
// See StringIndented.
func (part *Part) String() string { _ = "STUB: not implemented"; return "" }

// StringIndented returns an indented Part.
//
// See merkle.Proof#StringIndented
func (part *Part) StringIndented(indent string) string { _ = "STUB: not implemented"; return "" }

func (part *Part) ToProto() (*cmtproto.Part, error) { _ = "STUB: not implemented"; return nil, nil }

func PartFromProto(pb *cmtproto.Part) (*Part, error) { _ = "STUB: not implemented"; return nil, nil }

//-------------------------------------

type PartSetHeader struct {
	Total uint32            `json:"total"`
	Hash  cmtbytes.HexBytes `json:"hash"`
}

// String returns a string representation of PartSetHeader.
//
// 1. total number of parts
// 2. first 6 bytes of the hash
func (psh PartSetHeader) String() string { _ = "STUB: not implemented"; return "" }

func (psh PartSetHeader) IsZero() bool { _ = "STUB: not implemented"; return false }

func (psh PartSetHeader) Equals(other PartSetHeader) bool { _ = "STUB: not implemented"; return false }

// ValidateBasic performs basic validation.
func (psh PartSetHeader) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// Hash can be empty in case of POLBlockID.PartSetHeader in Proposal.

// ToProto converts PartSetHeader to protobuf
func (psh *PartSetHeader) ToProto() cmtproto.PartSetHeader {
	_ = "STUB: not implemented"
	return *new(cmtproto.PartSetHeader)
}

// FromProto sets a protobuf PartSetHeader to the given pointer
func PartSetHeaderFromProto(ppsh *cmtproto.PartSetHeader) (*PartSetHeader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ProtoPartSetHeaderIsZero is similar to the IsZero function for
// PartSetHeader, but for the Protobuf representation.
func ProtoPartSetHeaderIsZero(ppsh *cmtproto.PartSetHeader) bool {
	_ = "STUB: not implemented"
	return false
}

//-------------------------------------

type PartSet struct {
	total uint32
	hash  []byte

	mtx cmtsync.Mutex

	// buffer is a single preallocated buffer for all parts
	buffer []byte
	// partSize specifies the size of each part in bytes (excluding last part)
	partSize int
	// lastPartSize is the size of the last part (which may be smaller than partSize)
	lastPartSize int
	// proofs stores part proofs separately
	proofs        []merkle.Proof
	partsBitArray *bits.BitArray
	count         uint32
	// byteSize is the total size (in bytes). Used to ensure that the
	// part set doesn't exceed the maximum block bytes
	byteSize int64

	TxPos []TxPosition
}

// Returns an immutable, full PartSet from the data bytes.
// The data bytes are split into "partSize" chunks, and merkle tree computed.
// CONTRACT: partSize is greater than zero.
func NewPartSetFromData(data []byte, partSize uint32) (ops *PartSet, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Compute merkle proofs

// Fill the buffer in a single copy and populate metadata.

// Set sizes and bookkeeping.

// newPartSetFromChunks creates a new PartSet from given data chunks, and other data.
func newPartSetFromChunks(chunks [][]byte, root cmtbytes.HexBytes, proofs []*merkle.Proof, partSize int) (*PartSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// create a new partset using the new parity parts.

// access ps directly, without mutex, because we know it is not used elsewhere

// Ensure we don't exceed buffer bounds

// Encode Extend erasure encodes the block parts. Only the original parts should be
// provided. The parity data is formed into its own PartSet and returned
// alongside the length of the last part. The length of the last part is
// necessary because the last part may be padded with zeros after decoding. These zeros must be removed before computi
func Encode(ops *PartSet, partSize uint32) (*PartSet, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// allocate once, only slice later

// pad ONLY the last chunk and not the part with zeros if necessary AFTER the root has been generated

// init an encoder if it is not already initialized using the original
// number of parts.

// Encode the parts.

// only the parity data is needed for the new partset.

// IsReadyForDecoding returns true if the PartSet has every single part, not just
// ready to be decoded.
// TODO: this here only requires 2/3rd. We need all the data now because we have no erasure encoding.
func (ps *PartSet) IsReadyForDecoding() bool { _ = "STUB: not implemented"; return false }

// pruneLastPart trims the last original part to its true length, removing any
// Reed-Solomon padding bytes that were added during encoding.
func pruneLastPart(data [][]byte, lastIndex uint32, lastPartLen int) {
	_ = "STUB: not implemented"
	return
}

// Decode uses the block parts that are provided to reconstruct the original
// data. It throws an error if the PartSet is incomplete or the resulting root
// is different from that in the PartSetHeader. Parts are fully complete with
// proofs after decoding.
func Decode(ops, eps *PartSet, lastPartLen int) (*PartSet, *PartSet, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// prune the last part if we need to

// recalculate all of the proofs since we apparently don't have a function
// to generate a single proof... TODO: don't generate proofs for block parts
// we already have...

// recalculate all of the proofs since we apparently don't have a function
// to generate a single proof... TODO: don't generate proofs for block parts
// we already have.

// Returns an empty PartSet ready to be populated.
func NewPartSetFromHeader(header PartSetHeader, partSize uint32) *PartSet {
	_ = "STUB: not implemented"
	// Calculate buffer size: (total-1) full parts + potentially smaller last part
	return nil
}

// default to 0; will be updated when the last part is added

func (ps *PartSet) Header() PartSetHeader { _ = "STUB: not implemented"; return *new(PartSetHeader) }

func (ps *PartSet) HasHeader(header PartSetHeader) bool { _ = "STUB: not implemented"; return false }

func (ps *PartSet) BitArray() *bits.BitArray { _ = "STUB: not implemented"; return nil }

func (ps *PartSet) Hash() []byte { _ = "STUB: not implemented"; return nil }

func (ps *PartSet) HashesTo(hash []byte) bool { _ = "STUB: not implemented"; return false }

func (ps *PartSet) Count() uint32 { _ = "STUB: not implemented"; return 0 }

func (ps *PartSet) ByteSize() int64 { _ = "STUB: not implemented"; return 0 }

func (ps *PartSet) Total() uint32 { _ = "STUB: not implemented"; return 0 }

// CONTRACT: part is validated using ValidateBasic.
func (ps *PartSet) AddPart(part *Part) (bool, error) {
	_ = "STUB: not implemented"
	// TODO: remove this? would be preferable if this only returned (false, nil)
	// when its a duplicate block part
	return false, nil
}

// Defense-in-depth: ensure the part's index is bound to its proof's index
// before any further processing. Callers are expected to enforce this via
// Part.ValidateBasic, but enforcing it at the sink prevents a structurally
// valid proof for index j from being installed in slot i (i != j).

// If part already exists, return false.

// The proof should be compatible with the number of parts.

func (ps *PartSet) AddPartWithoutProof(part *Part) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (ps *PartSet) addPart(part *Part) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// Invalid part index

// If part already exists, return false.

// Calculate buffer position and copy part data

// Ensure we don't exceed buffer bounds

// Track last part size if this is the last part

// getPartBytes returns the bytes for a part from the internal buffer.
// Assumes mutex is already locked.
//
// Returns nil if:
// - The part at given index doesn't exist (not yet received)
// - The index is out of bounds
// - Buffer access would be out of bounds
func (ps *PartSet) getPartBytes(index int) []byte { _ = "STUB: not implemented"; return nil }

// Calculate buffer position

// For the last part, use the actual size

// GetPartBytes returns only the bytes for a part, without the proof.
// This is more efficient when only the data is needed.
func (ps *PartSet) GetPartBytes(index int) []byte { _ = "STUB: not implemented"; return nil }

func (ps *PartSet) GetPart(index int) *Part { _ = "STUB: not implemented"; return nil }

func (ps *PartSet) HasPart(index int) bool { _ = "STUB: not implemented"; return false }

func (ps *PartSet) IsComplete() bool { _ = "STUB: not implemented"; return false }

func (ps *PartSet) GetBytes() []byte { _ = "STUB: not implemented"; return nil }

// StringShort returns a short version of String.
//
// (Count of Total)
func (ps *PartSet) StringShort() string { _ = "STUB: not implemented"; return "" }

func (ps *PartSet) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
