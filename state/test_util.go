package state

import (
	"testing"
	"time"

	"github.com/cometbft/cometbft/types"

	dbm "github.com/cometbft/cometbft-db"
)

func SetupTestCase(t testing.TB) (func(t testing.TB), dbm.DB, State) {
	_ = "STUB: not implemented"
	return nil, *new(dbm.DB), *new(State)
}

// SetupTestCaseWithPrivVal is like SetupTestCase but also returns the private validator
// that matches the genesis validator. This is useful for tests that need to create
// commits matching the validator set.
func SetupTestCaseWithPrivVal(t testing.TB) (func(t testing.TB), dbm.DB, State, types.PrivValidator) {
	_ = "STUB: not implemented"
	return nil, *new(dbm.DB), *new(State), *new(types.PrivValidator)
}

// Load the private validator that matches the genesis validator

// MakeTestCommit creates a commit for testing.
// This creates an empty commit (with only absent signatures) which is valid for
// state.MakeBlock because:
// - For height == InitialHeight, MedianTime is not called (genesis time is used)
// - For height > InitialHeight, MedianTime iterates over signatures but skips absent ones
//
// This approach is simpler than creating valid signatures because:
// 1. When state.LastBlockHeight == 0, state.LastValidators is empty
// 2. FilePV tracks signed heights and refuses to sign the same height twice
//
// The returned commit has a random BlockID which doesn't matter for these tests
// since we're not validating the commit against an actual previous block.
func MakeTestCommit(_ testing.TB, _ State, _ types.PrivValidator, _ time.Time) *types.Commit {
	_ = "STUB: not implemented"
	// Create a random block ID for the commit
	return nil
}

// Return an empty commit - no signatures
// This works because MedianTime only processes non-absent signatures,
// and state.MakeBlock uses genesis time when height == InitialHeight
