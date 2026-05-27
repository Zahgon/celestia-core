package types

import (
	"sync"
	"sync/atomic"

	"github.com/cometbft/cometbft/crypto/merkle"
	"github.com/cometbft/cometbft/libs/bits"
	"github.com/cometbft/cometbft/types"
)

// CombinedPartSet wraps two PartSet instances: one for original block data and one for parity data.
type CombinedPartSet struct {
	mtx      *sync.Mutex
	totalMap *bits.BitArray
	original *types.PartSet // holds the original parts (indexes: 0 to original.Total()-1)
	parity   *types.PartSet // holds parity parts (logical indexes start at original.Total())
	lastLen  uint32
	catchup  bool

	IsDecoding atomic.Bool
}

// NewCombinedSetFromCompactBlock creates a new CombinedPartSet from a
// CompactBlock using the PartSetHeader in the proposal and the BpHash from the
// CompactBlock.
func NewCombinedSetFromCompactBlock(cb *CompactBlock) *CombinedPartSet {
	_ = "STUB: not implemented"
	return nil
}

func NewCombinedPartSetFromOriginal(original *types.PartSet, catchup bool) *CombinedPartSet {
	_ = "STUB: not implemented"
	return nil
}

func (cps *CombinedPartSet) SetProposalData(original, parity *types.PartSet) {
	_ = "STUB: not implemented"
	return
}

func (cps *CombinedPartSet) Original() *types.PartSet { _ = "STUB: not implemented"; return nil }

func (cps *CombinedPartSet) Parity() *types.PartSet { _ = "STUB: not implemented"; return nil }

func (cps *CombinedPartSet) BitArray() *bits.BitArray { _ = "STUB: not implemented"; return nil }

// OringinalBitArray returns a BitArray that only missing parts if they are in the original
// part set.
func (cps *CombinedPartSet) MissingOriginal() *bits.BitArray { _ = "STUB: not implemented"; return nil }

func (cps *CombinedPartSet) Total() uint32 { _ = "STUB: not implemented"; return 0 }

func (cps *CombinedPartSet) IsComplete() bool { _ = "STUB: not implemented"; return false }

// CanDecode determines if enough parts have been added to decode the block.
func (cps *CombinedPartSet) CanDecode() bool { _ = "STUB: not implemented"; return false }

func (cps *CombinedPartSet) Decode() error { _ = "STUB: not implemented"; return nil }

// AddPart adds a recovery part to the combined part set. It assumes that the parts being
// added have already been verified.
func (cps *CombinedPartSet) AddPart(part *RecoveryPart, proof merkle.Proof) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Adjust the index to be relative to the parity set.

// AddOriginalPart adds an original part to the combined partset.
func (cps *CombinedPartSet) AddOriginalPart(part *types.Part) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (cps *CombinedPartSet) HasPart(index int) bool { _ = "STUB: not implemented"; return false }

func (cps *CombinedPartSet) GetPart(index uint32) (*types.Part, bool) {
	_ = "STUB: not implemented"
	return nil, false
}
