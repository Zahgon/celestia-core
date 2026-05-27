package kv

import (
	"context"

	"github.com/cometbft/cometbft/libs/log"

	dbm "github.com/cometbft/cometbft-db"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/libs/pubsub/query"
	"github.com/cometbft/cometbft/libs/pubsub/query/syntax"
	"github.com/cometbft/cometbft/state/indexer"
	"github.com/cometbft/cometbft/state/txindex"
)

const (
	tagKeySeparator     = "/"
	tagKeySeparatorRune = '/'
	eventSeqSeparator   = "$es$"
)

var _ txindex.TxIndexer = (*TxIndex)(nil)
var _ txindex.PagedTxIndexer = (*TxIndex)(nil)

type txRef struct {
	hash   []byte
	height int64
	index  uint32
}

// TxIndex is the simplest possible indexer, backed by key-value storage (levelDB).
type TxIndex struct {
	store dbm.DB
	// Number the events in the event list
	eventSeq int64

	log log.Logger
}

// NewTxIndex creates new KV indexer.
func NewTxIndex(store dbm.DB) *TxIndex { _ = "STUB: not implemented"; return nil }

func (txi *TxIndex) SetLogger(l log.Logger) {
	_ = "STUB: not implemented"

	// Get gets transaction from the TxIndex storage and returns it or nil if the
	// transaction is not found.
	return
}

func (txi *TxIndex) Get(hash []byte) (*abci.TxResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AddBatch indexes a batch of transactions using the given list of events. Each
// key that indexed from the tx's events is a composite of the event type and
// the respective attribute's key delimited by a "." (eg. "account.number").
// Any event with an empty type is not indexed.
func (txi *TxIndex) AddBatch(b *txindex.Batch) error { _ = "STUB: not implemented"; return nil }

// Index indexes a single transaction using the given list of events. Each key
// that indexed from the tx's events is a composite of the event type and the
// respective attribute's key delimited by a "." (eg. "account.number").
// Any event with an empty type is not indexed.
//
// If a transaction is indexed with the same hash as a previous transaction, it will
// be overwritten unless the tx result was NOT OK and the prior result was OK i.e.
// more transactions that successfully executed overwrite transactions that failed
// or successful yet older transactions.
func (txi *TxIndex) Index(result *abci.TxResult) error { _ = "STUB: not implemented"; return nil }

// if the new transaction failed and it's already indexed in an older block and was successful
// we skip it as we want users to get the older successful transaction when they query.

// index tx by events

// index by height (always)

// index by hash (always)

func (txi *TxIndex) indexEvents(result *abci.TxResult, hash []byte, store dbm.Batch) error {
	_ = "STUB: not implemented"
	return nil
}

// only index events with a non-empty type

// index if `index: true` is set

// ensure event does not conflict with a reserved prefix key

// Search performs a search using the given query.
//
// It breaks the query into conditions (like "tx.height > 5"). For each
// condition, it queries the DB index. One special use cases here: (1) if
// "tx.hash" is found, it returns tx result for it (2) for range queries it is
// better for the client to provide both lower and upper bounds, so we are not
// performing a full scan. Results from querying indexes are then intersected
// and returned to the caller, in no particular order.
//
// Search will exit early and return any result fetched so far,
// when a message is received on the context chan.
func (txi *TxIndex) Search(ctx context.Context, q *query.Query) ([]*abci.TxResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SearchPaged performs the same match as Search, but applies ordering and
// pagination before loading full TxResult protobufs.
func (txi *TxIndex) SearchPaged(
	ctx context.Context,
	q *query.Query,
	orderBy string,
	skipCount, pageSize int,
) ([]*abci.TxResult, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (txi *TxIndex) loadResults(ctx context.Context, refs []txRef) ([]*abci.TxResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (txi *TxIndex) searchRefs(ctx context.Context, q *query.Query) (map[string]txRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// get a list of conditions (like "tx.height > 5")

// if there is a hash condition, return the result immediately

// conditions to skip because they're handled before "everything else"

// If we are not matching events and tx.height = 3 occurs more than once, the later value will
// overwrite the first one.

// extract ranges
// if both upper and lower bounds exist, it's better to get them in order not
// no iterate over kvs that are not within range.

// If we have a query range over height and want to still look for
// specific event values we do not want to simply return all
// transactios in this height range. We remember the height range info
// and pass it on to match() to take into account when processing events.

// Ignore any remaining conditions if the first condition resulted
// in no matches (assuming implicit AND operand).

// if there is a height condition ("tx.height=3"), extract it

// for all other conditions

// Ignore any remaining conditions if the first condition resulted
// in no matches (assuming implicit AND operand).

func lookForHash(conditions []syntax.Condition) (hash []byte, ok bool, err error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func (txi *TxIndex) setTmpHashes(tmpHeights map[string]txRef, key, value []byte) {
	_ = "STUB: not implemented"
	return
}

// Copy the value because the iterator will be reused.

// match returns all matching txs by hash that meet a given condition and start
// key. An already filtered result (filteredHashes) is provided such that any
// non-intersecting matches are removed.
//
// NOTE: filteredHashes may be empty if no previous condition has matched.
func (txi *TxIndex) match(
	ctx context.Context,
	c syntax.Condition,
	startKeyBz []byte,
	filteredHashes map[string]txRef,
	firstRun bool,
	heightInfo HeightInfo,
) map[string]txRef {
	_ = "STUB: not implemented"
	// A previous match was attempted but resulted in no matches, so we return
	// no matches (assuming AND operand).
	return nil
}

//nolint:staticcheck

// If we have a height range in a query, we need only transactions
// for this height

// Potentially exit early.

// XXX: can't use startKeyBz here because c.Operand is nil
// (e.g. "account.owner/<nil>/" won't match w/ a single row)

// Potentially exit early.

// XXX: startKey does not apply here.
// For example, if startKey = "account.owner/an/" and search query = "account.owner CONTAINS an"
// we can't iterate with prefix "account.owner/an/" because we might miss keys like "account.owner/Ulan/"

// Potentially exit early.

// Either:
//
// 1. Regardless if a previous match was attempted, which may have had
// results, but no match was found for the current condition, then we
// return no matches (assuming AND operand).
//
// 2. A previous match was not attempted, so we return all results.

// Remove/reduce matches in filteredHashes that were not found in this
// match (tmpHashes).

// Potentially exit early.

// matchRange returns all matching txs by hash that meet a given queryRange and
// start key. An already filtered result (filteredHashes) is provided such that
// any non-intersecting matches are removed.
//
// NOTE: filteredHashes may be empty if no previous condition has matched.
func (txi *TxIndex) matchRange(
	ctx context.Context,
	qr indexer.QueryRange,
	startKey []byte,
	filteredHashes map[string]txRef,
	firstRun bool,
	heightInfo HeightInfo,
) map[string]txRef {
	_ = "STUB: not implemented"
	// A previous match was attempted but resulted in no matches, so we return
	// no matches (assuming AND operand).
	return nil
}

// TODO: We need to make a function for getting it.Key() as a byte slice with no copies.
// It currently copies the source data (which can change on a subsequent .Next() call) but that
// is not an issue for us.

// XXX: passing time in a ABCI Events is not yet implemented
// case time.Time:
// 	v := strconv.ParseInt(extractValueFromKey(it.Key()), 10, 64)
// 	if v == r.upperBound {
// 		break
// 	}

// Potentially exit early.

// Either:
//
// 1. Regardless if a previous match was attempted, which may have had
// results, but no match was found for the current condition, then we
// return no matches (assuming AND operand).
//
// 2. A previous match was not attempted, so we return all results.

// Remove/reduce matches in filteredHashes that were not found in this
// match (tmpHashes).

// Potentially exit early.

func uniqueRefs(filteredHashes map[string]txRef) []txRef { _ = "STUB: not implemented"; return nil }

func sortRefs(refs []txRef, orderBy string) error { _ = "STUB: not implemented"; return nil }

func txRefLess(a, b txRef) bool { _ = "STUB: not implemented"; return false }

// Keys

func isTagKey(key []byte) bool {
	_ = "STUB: not implemented"
	// Normally, if the event was indexed with an event sequence, the number of
	// tags should 4. Alternatively it should be 3 if the event was not indexed
	// with the corresponding event sequence. However, some attribute values in
	// production can contain the tag separator. Therefore, the condition is >= 3.
	return false
}

func extractHeightFromKey(key []byte) (int64, error) {
	_ = "STUB: not implemented"
	// the height is the second last element in the key.
	// Find the position of the last occurrence of tagKeySeparator
	return 0, nil
}

// Find the position of the second last occurrence of tagKeySeparator

// Extract the height part of the key

func extractIndexFromKey(key []byte) (uint32, error) { _ = "STUB: not implemented"; return 0, nil }

func txRefFromKeyValue(key, value []byte) (txRef, error) {
	_ = "STUB: not implemented"
	return *new(txRef), nil
}

func extractValueFromKey(key []byte) string {
	_ = "STUB: not implemented"
	// Find the positions of tagKeySeparator in the byte slice
	return ""
}

// If there are less than 2 occurrences of tagKeySeparator, return an empty string

// Extract the value between the first and second last occurrence of tagKeySeparator

// Trim any leading or trailing whitespace

// TODO: Do an unsafe cast to avoid an extra allocation here

func extractEventSeqFromKey(key []byte) string { _ = "STUB: not implemented"; return "" }

//nolint:staticcheck

func keyForEvent(key string, value string, result *abci.TxResult, eventSeq int64) []byte {
	_ = "STUB: not implemented"
	return nil
}

func keyForHeight(result *abci.TxResult) []byte { _ = "STUB: not implemented"; return nil }

// Added to facilitate having the eventSeq in event keys
// Otherwise queries break expecting 5 entries

func startKeyForCondition(c syntax.Condition, height int64) []byte {
	_ = "STUB: not implemented"
	return nil
}

func startKey(fields ...interface{}) []byte { _ = "STUB: not implemented"; return nil }

func (txi *TxIndex) indexResult(batch dbm.Batch, result *abci.TxResult) error {
	_ = "STUB: not implemented"
	return nil
}

// if the new transaction failed and it's already indexed in an older block and was successful
// we skip it as we want users to get the older successful transaction when they query.

// index tx by events

// index by height (always)

// index by hash (always)
