package db

import (
	"regexp"

	dbm "github.com/cometbft/cometbft-db"

	cmtsync "github.com/cometbft/cometbft/libs/sync"
	"github.com/cometbft/cometbft/light/store"
	"github.com/cometbft/cometbft/types"
)

var sizeKey = []byte("size")

type dbs struct {
	db     dbm.DB
	prefix string

	mtx  cmtsync.RWMutex
	size uint16
}

// New returns a Store that wraps any DB (with an optional prefix in case you
// want to use one DB with many light clients).
func New(db dbm.DB, prefix string) store.Store { _ = "STUB: not implemented"; return *new(store.Store) }

// SaveLightBlock persists LightBlock to the db.
//
// Safe for concurrent use by multiple goroutines.
func (s *dbs) SaveLightBlock(lb *types.LightBlock) error { _ = "STUB: not implemented"; return nil }

// DeleteLightBlockAndValidatorSet deletes the LightBlock from
// the db.
//
// Safe for concurrent use by multiple goroutines.
func (s *dbs) DeleteLightBlock(height int64) error { _ = "STUB: not implemented"; return nil }

// LightBlock retrieves the LightBlock at the given height.
//
// Safe for concurrent use by multiple goroutines.
func (s *dbs) LightBlock(height int64) (*types.LightBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// LastLightBlockHeight returns the last LightBlock height stored.
//
// Safe for concurrent use by multiple goroutines.
func (s *dbs) LastLightBlockHeight() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// FirstLightBlockHeight returns the first LightBlock height stored.
//
// Safe for concurrent use by multiple goroutines.
func (s *dbs) FirstLightBlockHeight() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// LightBlockBefore iterates over light blocks until it finds a block before
// the given height. It returns ErrLightBlockNotFound if no such block exists.
//
// Safe for concurrent use by multiple goroutines.
func (s *dbs) LightBlockBefore(height int64) (*types.LightBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Prune prunes header & validator set pairs until there are only size pairs
// left.
//
// Safe for concurrent use by multiple goroutines.
func (s *dbs) Prune(size uint16) error {
	_ = "STUB: not implemented"
	// 1) Check how many we need to prune.
	return nil
}

// nothing to prune

// 2) Iterate over headers and perform a batch operation.

// 3) Update size.

// Size returns the number of header & validator set pairs.
//
// Safe for concurrent use by multiple goroutines.
func (s *dbs) Size() uint16 { _ = "STUB: not implemented"; return 0 }

func (s *dbs) lbKey(height int64) []byte { _ = "STUB: not implemented"; return nil }

var keyPattern = regexp.MustCompile(`^(lb)/([^/]*)/([0-9]+)$`)

func parseKey(key []byte) (part string, prefix string, height int64, ok bool) {
	_ = "STUB: not implemented"
	return "", "", 0, false
}

// good!

func parseLbKey(key []byte) (prefix string, height int64, ok bool) {
	_ = "STUB: not implemented"
	return "", 0, false
}

func marshalSize(size uint16) []byte { _ = "STUB: not implemented"; return nil }

func unmarshalSize(bz []byte) uint16 { _ = "STUB: not implemented"; return 0 }
