package kv

import (
	"context"

	dbm "github.com/cometbft/cometbft-db"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cometbft/cometbft/libs/pubsub/query"
	"github.com/cometbft/cometbft/libs/pubsub/query/syntax"
	"github.com/cometbft/cometbft/state/indexer"
	"github.com/cometbft/cometbft/types"
)

var _ indexer.BlockIndexer = (*BlockerIndexer)(nil)

// BlockerIndexer implements a block indexer, indexing FinalizeBlock
// events with an underlying KV store. Block events are indexed by their height,
// such that matching search criteria returns the respective block height(s).
type BlockerIndexer struct {
	store dbm.DB

	// Add unique event identifier to use when querying
	// Matching will be done both on height AND eventSeq
	eventSeq int64
	log      log.Logger
}

func New(store dbm.DB) *BlockerIndexer { _ = "STUB: not implemented"; return nil }

func (idx *BlockerIndexer) SetLogger(l log.Logger) {
	_ = "STUB: not implemented"

	// Has returns true if the given height has been indexed. An error is returned
	// upon database query failure.
	return
}

func (idx *BlockerIndexer) Has(height int64) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Index indexes FinalizeBlock events for a given block by its height.
// The following is indexed:
//
// primary key: encode(block.height | height) => encode(height)
// FinalizeBlock events: encode(eventType.eventAttr|eventValue|height|finalize_block|eventSeq) => encode(height)
func (idx *BlockerIndexer) Index(bh types.EventDataNewBlockEvents) error {
	_ = "STUB: not implemented"
	return nil
}

// 1. index by height

// 2. index block events

// Search performs a query for block heights that match a given FinalizeBlock
// event search criteria. The given query can match against zero,
// one or more block heights. In the case of height queries, i.e. block.height=H,
// if the height is indexed, that height alone will be returned. An error and
// nil slice is returned. Otherwise, a non-nil slice and nil error is returned.
func (idx *BlockerIndexer) Search(ctx context.Context, q *query.Query) ([]int64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// conditions to skip because they're handled before "everything else"

// If we are not matching events and block.height occurs more than once, the later value will
// overwrite the first one.

// Extract ranges. If both upper and lower bounds exist, it's better to get
// them in order as to not iterate over kvs that are not within range.

// If we have additional constraints and want to query per event
// attributes, we cannot simply return all blocks for a height.
// But we remember the height we want to find and forward it to
// match(). If we only have the height constraint
// in the query (the second part of the ||), we don't need to query
// per event conditions and return all events within the height range.

// If we have a query range over height and want to still look for
// specific event values we do not want to simply return all
// blocks in this height range. We remember the height range info
// and pass it on to match() to take into account when processing events.

// If the query contains ranges other than the height then we need to treat the height
// range when querying the conditions of the other range.
// Otherwise we can just return all the blocks within the height range (as there is no
// additional constraint on events)

// Ignore any remaining conditions if the first condition resulted in no
// matches (assuming implicit AND operand).

// for all other conditions

// Ignore any remaining conditions if the first condition resulted in no
// matches (assuming implicit AND operand).

// fetch matching heights

// matchRange returns all matching block heights that match a given QueryRange
// and start key. An already filtered result (filteredHeights) is provided such
// that any non-intersecting matches are removed.
//
// NOTE: The provided filteredHeights may be empty if no previous condition has
// matched.
func (idx *BlockerIndexer) matchRange(
	ctx context.Context,
	qr indexer.QueryRange,
	startKey []byte,
	filteredHeights map[string][]byte,
	firstRun bool,
	heightInfo HeightInfo,
) (map[string][]byte, error) {
	_ = "STUB: not implemented"
	// A previous match was attempted but resulted in no matches, so we return
	// no matches (assuming AND operand).
	return nil, nil
}

// The precision here is 125. For numbers bigger than this, the value
// will not be parsed properly

// Either:
//
// 1. Regardless if a previous match was attempted, which may have had
// results, but no match was found for the current condition, then we
// return no matches (assuming AND operand).
//
// 2. A previous match was not attempted, so we return all results.

// Remove/reduce matches in filteredHashes that were not found in this
// match (tmpHashes).

// Check whether in this iteration we have not found an overlapping height (tmpHeight == nil)
// or whether the events in which the attributed occurred do not match (first part of the condition)

func (idx *BlockerIndexer) setTmpHeights(tmpHeights map[string][]byte, it dbm.Iterator) {
	_ = "STUB: not implemented"
	// If we return attributes that occur within the same events, then store the
	// event sequence in the result map as well.
	return
}

// Copy the value because the iterator will be reused.

// match returns all matching heights that meet a given query condition and start
// key. An already filtered result (filteredHeights) is provided such that any
// non-intersecting matches are removed.
//
// NOTE: The provided filteredHeights may be empty if no previous condition has
// matched.
func (idx *BlockerIndexer) match(
	ctx context.Context,
	c syntax.Condition,
	startKeyBz []byte,
	filteredHeights map[string][]byte,
	firstRun bool,
	heightInfo HeightInfo,
) (map[string][]byte, error) {
	_ = "STUB: not implemented"
	// A previous match was attempted but resulted in no matches, so we return
	// no matches (assuming AND operand).
	return nil, nil
}

//nolint:staticcheck

// Either:
//
// 1. Regardless if a previous match was attempted, which may have had
// results, but no match was found for the current condition, then we
// return no matches (assuming AND operand).
//
// 2. A previous match was not attempted, so we return all results.

// Remove/reduce matches in filteredHeights that were not found in this
// match (tmpHeights).

func (idx *BlockerIndexer) indexEvents(batch dbm.Batch, events []abci.Event, height int64) error {
	_ = "STUB: not implemented"
	return nil
}

// only index events with a non-empty type

// index iff the event specified index:true and it's not a reserved event
