package bits

import (
	"regexp"
	"sync"

	cmtprotobits "github.com/cometbft/cometbft/proto/tendermint/libs/bits"
)

// BitArray is a thread-safe implementation of a bit array.
type BitArray struct {
	mtx   sync.Mutex
	Bits  int      `json:"bits"`  // NOTE: persisted via reflect, must be exported
	Elems []uint64 `json:"elems"` // NOTE: persisted via reflect, must be exported
}

// NewBitArray returns a new bit array.
// It returns nil if the number of bits is zero.
func NewBitArray(bits int) *BitArray { _ = "STUB: not implemented"; return nil }

// NewBitArrayFromFn returns a new bit array.
// It returns nil if the number of bits is zero.
// It initializes the `i`th bit to the value of `fn(i)`.
func NewBitArrayFromFn(bits int, fn func(int) bool) *BitArray {
	_ = "STUB: not implemented"
	return nil
}

// Size returns the number of bits in the bitarray.
func (bA *BitArray) Size() int { _ = "STUB: not implemented"; return 0 }

// GetIndex returns the bit at index i within the bit array.
// The behavior is undefined if i >= bA.Bits
func (bA *BitArray) GetIndex(i int) bool { _ = "STUB: not implemented"; return false }

// Fill sets all bits to true
func (bA *BitArray) Fill() { _ = "STUB: not implemented"; return }

func (bA *BitArray) getIndex(i int) bool { _ = "STUB: not implemented"; return false }

// SetIndex sets the bit at index i within the bit array.
// The behavior is undefined if i >= bA.Bits
func (bA *BitArray) SetIndex(i int, v bool) bool { _ = "STUB: not implemented"; return false }

func (bA *BitArray) setIndex(i int, v bool) bool { _ = "STUB: not implemented"; return false }

// AddBitArray combines two bit arrays by taking the bitwise OR of the two. If
// the two bit arrays have different lengths, AddBitArray right-pads the smaller
// of the two bit-arrays with zeroes. Thus the size of the return value is the
// maximum of the two provided bit arrays.
func (bA *BitArray) AddBitArray(b *BitArray) { _ = "STUB: not implemented"; return }

// Update bA's Bits count to be the max of the two

// Resize bA's Elems if b is longer

// Perform bitwise OR operation up to the length of b

// Copy returns a copy of the provided bit array.
func (bA *BitArray) Copy() *BitArray { _ = "STUB: not implemented"; return nil }

func (bA *BitArray) copy() *BitArray { _ = "STUB: not implemented"; return nil }

func (bA *BitArray) copyBits(bits int) *BitArray { _ = "STUB: not implemented"; return nil }

// Or returns a bit array resulting from a bitwise OR of the two bit arrays.
// If the two bit-arrys have different lengths, Or right-pads the smaller of the two bit-arrays with zeroes.
// Thus the size of the return value is the maximum of the two provided bit arrays.
func (bA *BitArray) Or(o *BitArray) *BitArray { _ = "STUB: not implemented"; return nil }

// And returns a bit array resulting from a bitwise AND of the two bit arrays.
// If the two bit-arrys have different lengths, this truncates the larger of the two bit-arrays from the right.
// Thus the size of the return value is the minimum of the two provided bit arrays.
func (bA *BitArray) And(o *BitArray) *BitArray { _ = "STUB: not implemented"; return nil }

func (bA *BitArray) and(o *BitArray) *BitArray { _ = "STUB: not implemented"; return nil }

// Not returns a bit array resulting from a bitwise Not of the provided bit array.
func (bA *BitArray) Not() *BitArray { _ = "STUB: not implemented"; return nil }

// Degenerate

func (bA *BitArray) not() *BitArray { _ = "STUB: not implemented"; return nil }

// Sub subtracts the two bit-arrays bitwise, without carrying the bits.
// Note that carryless subtraction of a - b is (a and not b).
// The output is the same as bA, regardless of o's size.
// If bA is longer than o, o is right padded with zeroes
func (bA *BitArray) Sub(o *BitArray) *BitArray { _ = "STUB: not implemented"; return nil }

// TODO: Decide if we should do 1's complement here?

// output is the same size as bA

// Only iterate to the minimum size between the two.
// If o is longer, those bits are ignored.
// If bA is longer, then skipping those iterations is equivalent
// to right padding with 0's

// &^ is and not in golang

// IsEmpty returns true iff all bits in the bit array are 0
func (bA *BitArray) IsEmpty() bool { _ = "STUB: not implemented"; return false }

// should this be opposite?

// IsFull returns true iff all bits in the bit array are 1.
func (bA *BitArray) IsFull() bool { _ = "STUB: not implemented"; return false }

// Check all elements except the last

// Check that the last element has (lastElemBits) 1's

// PickRandom returns a random index for a set bit in the bit array.
// If there is no such value, it returns 0, false.
// It uses the global randomness in `random.go` to get this index.
func (bA *BitArray) PickRandom() (int, bool) { _ = "STUB: not implemented"; return 0, false }

// no bits set to true

func (bA *BitArray) GetTrueIndices() []int { _ = "STUB: not implemented"; return nil }

// set all true indices

//nolint:gosec

// handle last element

//nolint:gosec

func (bA *BitArray) getNumTrueIndices() int { _ = "STUB: not implemented"; return 0 }

// size and elements must be valid to do this calc

// handle all elements except the last one

// handle last element

// getNthTrueIndex returns the index of the nth true bit in the bit array.
// n is 0 indexed. (e.g. for bitarray x__x, getNthTrueIndex(0) returns 0).
// If there is no such value, it returns -1.
func (bA *BitArray) getNthTrueIndex(n int) int { _ = "STUB: not implemented"; return 0 }

// Iterate over each element

// Count set bits in the current element

// If the count of set bits in this element plus the count so far
// is greater than or equal to n, then the nth bit must be in this element

// Find the index of the nth set bit within this element

// Calculate the absolute index of the set bit

// If the count is not enough, continue to the next element

// If we reach here, it means n is out of range

// String returns a string representation of BitArray: BA{<bit-string>},
// where <bit-string> is a sequence of 'x' (1) and '_' (0).
// The <bit-string> includes spaces and newlines to help people.
// For a simple sequence of 'x' and '_' characters with no spaces or newlines,
// see the MarshalJSON() method.
// Example: "BA{_x_}" or "nil-BitArray" for nil.
func (bA *BitArray) String() string { _ = "STUB: not implemented"; return "" }

// StringIndented returns the same thing as String(), but applies the indent
// at every 10th bit, and twice at every 50th bit.
func (bA *BitArray) StringIndented(indent string) string { _ = "STUB: not implemented"; return "" }

func (bA *BitArray) stringIndented(indent string) string { _ = "STUB: not implemented"; return "" }

// Bytes returns the byte representation of the bits within the bitarray.
func (bA *BitArray) Bytes() []byte { _ = "STUB: not implemented"; return nil }

// Update sets the bA's bits to be that of the other bit array.
// The copying begins from the begin of both bit arrays.
func (bA *BitArray) Update(o *BitArray) { _ = "STUB: not implemented"; return }

// MarshalJSON implements json.Marshaler interface by marshaling bit array
// using a custom format: a string of '-' or 'x' where 'x' denotes the 1 bit.
func (bA *BitArray) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

var bitArrayJSONRegexp = regexp.MustCompile(`\A"([_x]*)"\z`)

// UnmarshalJSON implements json.Unmarshaler interface by unmarshaling a custom
// JSON description.
func (bA *BitArray) UnmarshalJSON(bz []byte) error { _ = "STUB: not implemented"; return nil }

// This is required e.g. for encoding/json when decoding
// into a pointer with pre-allocated BitArray.

// Validate 'b'.

// Construct new BitArray and copy over.

// Treat it as if we encountered the case: b == "null"

//nolint:govet

// ToProto converts BitArray to protobuf
func (bA *BitArray) ToProto() *cmtprotobits.BitArray { _ = "STUB: not implemented"; return nil }

// FromProto sets a protobuf BitArray to the given pointer.
func (bA *BitArray) FromProto(protoBitArray *cmtprotobits.BitArray) {
	_ = "STUB: not implemented"
	return
}

// ValidateBasic validates a BitArray. Note that a nil BitArray and BitArray of
// size 0 bits is valid. However the number of Bits and Elems be valid based on
// each other.
func (bA *BitArray) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

func numElements(bits int) int { _ = "STUB: not implemented"; return 0 }
