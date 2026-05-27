package app

import (
	"sync"

	abci "github.com/cometbft/cometbft/abci/types"
)

const (
	snapshotChunkSize = 1e6

	// Keep only the most recent 10 snapshots. Older snapshots are pruned
	maxSnapshotCount = 10
)

// SnapshotStore stores state sync snapshots. Snapshots are stored simply as
// JSON files, and chunks are generated on-the-fly by splitting the JSON data
// into fixed-size chunks.
type SnapshotStore struct {
	sync.RWMutex
	dir      string
	metadata []abci.Snapshot
}

// NewSnapshotStore creates a new snapshot store.
func NewSnapshotStore(dir string) (*SnapshotStore, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// loadMetadata loads snapshot metadata. Does not take out locks, since it's
// called internally on construction.
func (s *SnapshotStore) loadMetadata() error { _ = "STUB: not implemented"; return nil }

// saveMetadata saves snapshot metadata. Does not take out locks, since it's
// called internally from e.g. Create().
func (s *SnapshotStore) saveMetadata() error { _ = "STUB: not implemented"; return nil }

// save the file to a new file and move it to make saving atomic.

//nolint: gosec

// Create creates a snapshot of the given application state's key/value pairs.
func (s *SnapshotStore) Create(state *State) (abci.Snapshot, error) {
	_ = "STUB: not implemented"
	return *new(abci.Snapshot), nil
}

//nolint:gosec

// Prune removes old snapshots ensuring only the most recent n snapshots remain
func (s *SnapshotStore) Prune(n int) error { _ = "STUB: not implemented"; return nil }

// snapshots are appended to the metadata struct, hence pruning removes from
// the front of the array

// update metadata by removing the deleted snapshots

// List lists available snapshots.
func (s *SnapshotStore) List() ([]*abci.Snapshot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// LoadChunk loads a snapshot chunk.
func (s *SnapshotStore) LoadChunk(height uint64, format uint32, chunk uint32) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// byteChunk returns the chunk at a given index from the full byte slice.
func byteChunk(bz []byte, index uint32) []byte { _ = "STUB: not implemented"; return nil }

// byteChunks calculates the number of chunks in the byte slice.
func byteChunks(bz []byte) uint32 { _ = "STUB: not implemented"; return 0 }
