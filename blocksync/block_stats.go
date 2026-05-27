package blocksync

// blockStats is a circular blockSizes that holds at most n elements
// and provides O(1) average calculation and O(1) max retrieval
type blockStats struct {
	blockSizes []int
	capacity   int
	size       int // current number of elements
	head       int // index where the next element will be written
	max        int // maximum value in the blockSizes
}

const defaultCapacity = 10

// newBlockStats creates a new rotating blockSizes with given capacity
func newBlockStats(capacity int) *blockStats { _ = "STUB: not implemented"; return nil }

// Add adds a new element to the blockSizes
// If the blockSizes is full, it removes the oldest element
func (rb *blockStats) Add(value int) { _ = "STUB: not implemented"; return }

// Buffer is full, replace the oldest element

// recalculateMax recalculates the maximum value by scanning the blockSizes
// This is only called when the previous max value is removed
func (rb *blockStats) recalculateMax() { _ = "STUB: not implemented"; return }

// GetMax returns the maximum value in the blockSizes
func (rb *blockStats) GetMax() int {
	_ = "STUB: not implemented"

	// GetSize returns the current number of elements in the blockSizes
	return 0
}

func (rb *blockStats) GetSize() int {
	_ = "STUB: not implemented"

	// GetCapacity returns the maximum capacity of the blockSizes
	return 0
}

func (rb *blockStats) GetCapacity() int { _ = "STUB: not implemented"; return 0 }
