// Package psql implements an event sink backed by a PostgreSQL database.
package psql

import (
	"context"
	"database/sql"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/libs/pubsub/query"
	"github.com/cometbft/cometbft/types"
)

const (
	tableBlocks     = "blocks"
	tableTxResults  = "tx_results"
	tableEvents     = "events"
	tableAttributes = "attributes"
	driverName      = "postgres"
)

// EventSink is an indexer backend providing the tx/block index services.  This
// implementation stores records in a PostgreSQL database using the schema
// defined in state/indexer/sink/psql/schema.sql.
type EventSink struct {
	store   *sql.DB
	chainID string
}

// NewEventSink constructs an event sink associated with the PostgreSQL
// database specified by connStr. Events written to the sink are attributed to
// the specified chainID.
func NewEventSink(connStr, chainID string) (*EventSink, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DB returns the underlying Postgres connection used by the sink.
// This is exported to support testing.
func (es *EventSink) DB() *sql.DB {
	_ = "STUB: not implemented"

	// runInTransaction executes query in a fresh database transaction.
	// If query reports an error, the transaction is rolled back and the
	// error from query is reported to the caller.
	// Otherwise, the result of committing the transaction is returned.
	return nil
}

func runInTransaction(db *sql.DB, query func(*sql.Tx) error) error {
	_ = "STUB: not implemented"
	return nil
}

// report the initial error, not the rollback

// queryWithID executes the specified SQL query with the given arguments,
// expecting a single-row, single-column result containing an ID. If the query
// succeeds, the ID from the result is returned.
func queryWithID(tx *sql.Tx, query string, args ...interface{}) (uint32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// insertEvents inserts a slice of events and any indexed attributes of those
// events into the database associated with dbtx.
//
// If txID > 0, the event is attributed to the transaction with that
// ID; otherwise it is recorded as a block event.
func insertEvents(dbtx *sql.Tx, blockID, txID uint32, evts []abci.Event) error {
	_ = "STUB: not implemented"
	// Populate the transaction ID field iff one is defined (> 0).
	return nil
}

// Add each event to the events table, and retrieve its row ID to use when
// adding any attributes the event provides.

// Skip events with an empty type.

// Add any attributes flagged for indexing.

// makeIndexedEvent constructs an event from the specified composite key and
// value. If the key has the form "type.name", the event will have a single
// attribute with that name and the value; otherwise the event will have only
// a type and no attributes.
func makeIndexedEvent(compositeKey, value string) abci.Event {
	_ = "STUB: not implemented"
	return *new(abci.Event)
}

// IndexBlockEvents indexes the specified block header, part of the
// indexer.EventSink interface.
func (es *EventSink) IndexBlockEvents(h types.EventDataNewBlockEvents) error {
	_ = "STUB: not implemented"
	return nil
}

// Add the block to the blocks table and report back its row ID for use
// in indexing the events for the block.

// we already saw this block; quietly succeed

// Insert the special block meta-event for height.

// Insert all the block events. Order is important here,

func (es *EventSink) IndexTxEvents(txrs []*abci.TxResult) error {
	_ = "STUB: not implemented"
	return nil
}

// Encode the result message in protobuf wire format for indexing.

// Index the hash of the underlying transaction as a hex string.

// Find the block associated with this transaction. The block header
// must have been indexed prior to the transactions belonging to it.

// Insert a record for this tx_result and capture its ID for indexing events.

// we already saw this transaction; quietly succeed

// Insert the special transaction meta-events for hash and height.

// Index any events packaged with the transaction.

// SearchBlockEvents is not implemented by this sink, and reports an error for all queries.
func (es *EventSink) SearchBlockEvents(_ context.Context, _ *query.Query) ([]int64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SearchTxEvents is not implemented by this sink, and reports an error for all queries.
func (es *EventSink) SearchTxEvents(_ context.Context, _ *query.Query) ([]*abci.TxResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetTxByHash is not implemented by this sink, and reports an error for all queries.
func (es *EventSink) GetTxByHash(_ []byte) (*abci.TxResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// HasBlock is not implemented by this sink, and reports an error for all queries.
func (es *EventSink) HasBlock(_ int64) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// Stop closes the underlying PostgreSQL database.
func (es *EventSink) Stop() error { _ = "STUB: not implemented"; return nil }
