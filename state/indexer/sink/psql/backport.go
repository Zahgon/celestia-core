package psql

// This file adds code to the psql package that is needed for integration with
// v0.34, but which is not part of the original implementation.
//
// In v0.35, ADR 65 was implemented in which the TxIndexer and BlockIndexer
// interfaces were merged into a hybrid EventSink interface. The Backport*
// types defined here bridge the psql EventSink (which was built in terms of
// the v0.35 interface) to the old interfaces.
//
// We took this narrower approach to backporting to avoid pulling in a much
// wider-reaching set of changes in v0.35 that would have broken several of the
// v0.34.x APIs. The result is sufficient to work with the node plumbing as it
// exists in the v0.34 branch.

import (
	"context"

	"github.com/cometbft/cometbft/libs/log"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/libs/pubsub/query"
	"github.com/cometbft/cometbft/state/txindex"
	"github.com/cometbft/cometbft/types"
)

const (
	eventTypeFinalizeBlock = "finalize_block"
)

// TxIndexer returns a bridge from es to the CometBFT v0.34 transaction indexer.
func (es *EventSink) TxIndexer() BackportTxIndexer {
	_ = "STUB: not implemented"
	return *new(BackportTxIndexer)
}

// BackportTxIndexer implements the txindex.TxIndexer interface by delegating
// indexing operations to an underlying PostgreSQL event sink.
type BackportTxIndexer struct{ psql *EventSink }

// AddBatch indexes a batch of transactions in Postgres, as part of TxIndexer.
func (b BackportTxIndexer) AddBatch(batch *txindex.Batch) error {
	_ = "STUB: not implemented"
	return nil
}

// Index indexes a single transaction result in Postgres, as part of TxIndexer.
func (b BackportTxIndexer) Index(txr *abci.TxResult) error { _ = "STUB: not implemented"; return nil }

// Get is implemented to satisfy the TxIndexer interface, but is not supported
// by the psql event sink and reports an error for all inputs.
func (BackportTxIndexer) Get([]byte) (*abci.TxResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Search is implemented to satisfy the TxIndexer interface, but it is not
// supported by the psql event sink and reports an error for all inputs.
func (BackportTxIndexer) Search(context.Context, *query.Query) ([]*abci.TxResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (BackportTxIndexer) SetLogger(log.Logger) {
	_ = "STUB: not implemented"

	// BlockIndexer returns a bridge that implements the CometBFT v0.34 block
	// indexer interface, using the Postgres event sink as a backing store.
	return
}

func (es *EventSink) BlockIndexer() BackportBlockIndexer {
	_ = "STUB: not implemented"
	return *new(BackportBlockIndexer)
}

// BackportBlockIndexer implements the indexer.BlockIndexer interface by
// delegating indexing operations to an underlying PostgreSQL event sink.
type BackportBlockIndexer struct{ psql *EventSink }

// Has is implemented to satisfy the BlockIndexer interface, but it is not
// supported by the psql event sink and reports an error for all inputs.
func (BackportBlockIndexer) Has(_ int64) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Index indexes block begin and end events for the specified block.  It is
// part of the BlockIndexer interface.
func (b BackportBlockIndexer) Index(block types.EventDataNewBlockEvents) error {
	_ = "STUB: not implemented"
	return nil
}

// Search is implemented to satisfy the BlockIndexer interface, but it is not
// supported by the psql event sink and reports an error for all inputs.
func (BackportBlockIndexer) Search(context.Context, *query.Query) ([]int64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (BackportBlockIndexer) SetLogger(log.Logger) { _ = "STUB: not implemented"; return }
