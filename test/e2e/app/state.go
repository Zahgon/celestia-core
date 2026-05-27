package app

import (
	"sync"
)

const (
	stateFileName     = "app_state.json"
	prevStateFileName = "prev_app_state.json"
)

// Intermediate type used exclusively in serialization/deserialization of
// State, such that State need not expose any of its internal values publicly.
type serializedState struct {
	Height uint64
	Values map[string]string
	Hash   []byte
}

// State is the application state.
type State struct {
	sync.RWMutex
	height uint64
	values map[string]string
	hash   []byte

	currentFile string
	// app saves current and previous state for rollback functionality
	previousFile    string
	persistInterval uint64
	initialHeight   uint64
}

// NewState creates a new state.
func NewState(dir string, persistInterval uint64) (*State, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// load loads state from disk. It does not take out a lock, since it is called
// during construction.
func (s *State) load() error { _ = "STUB: not implemented"; return nil }

// if the current state doesn't exist then we try recover from the previous state

// save saves the state to disk. It does not take out a lock since it is called
// internally by Commit which does lock.
func (s *State) save() error { _ = "STUB: not implemented"; return nil }

// We write the state to a separate file and move it to the destination, to
// make it atomic.

//nolint:gosec

// We take the current state and move it to the previous state, replacing it

// Finally, we take the new state and replace the current state.

// GetHash provides a thread-safe way of accessing a copy of the current state
// hash.
func (s *State) GetHash() []byte { _ = "STUB: not implemented"; return nil }

// Info returns both the height and hash simultaneously, and is used in the
// ABCI Info call.
func (s *State) Info() (uint64, []byte) { _ = "STUB: not implemented"; return 0, nil }

// Export exports key/value pairs as JSON, used for state sync snapshots.
// Additionally returns the current height and hash of the state.
func (s *State) Export() ([]byte, uint64, []byte, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil, nil
}

// Import imports key/value pairs from JSON bytes, used for InitChain.AppStateBytes and
// state sync snapshots. It also saves the state once imported.
func (s *State) Import(height uint64, jsonBytes []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// Get fetches a value. A missing value is returned as an empty string.
func (s *State) Get(key string) string { _ = "STUB: not implemented"; return "" }

// Set sets a value. Setting an empty value is equivalent to deleting it.
func (s *State) Set(key, value string) { _ = "STUB: not implemented"; return }

// Query is used in the ABCI Query call, and provides both the current height
// and the value associated with the given key.
func (s *State) Query(key string) (string, uint64) { _ = "STUB: not implemented"; return "", 0 }

// Finalize is called after applying a block, updating the height and returning the new app_hash
func (s *State) Finalize() []byte { _ = "STUB: not implemented"; return nil }

// Commit commits the current state.
func (s *State) Commit() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *State) Rollback() error { _ = "STUB: not implemented"; return nil }

func (s *State) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

func (s *State) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// hashItems hashes a set of key/value items.
func hashItems(items map[string]string, height uint64) []byte {
	_ = "STUB: not implemented"
	return nil
}
