package null

import (
	"context"

	"github.com/cometbft/cometbft/libs/log"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/libs/pubsub/query"
	"github.com/cometbft/cometbft/state/txindex"
)

var _ txindex.TxIndexer = (*TxIndex)(nil)

// TxIndex acts as a /dev/null.
type TxIndex struct{}

// Get on a TxIndex is disabled and panics when invoked.
func (txi *TxIndex) Get(_ []byte) (*abci.TxResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AddBatch is a noop and always returns nil.
func (txi *TxIndex) AddBatch(_ *txindex.Batch) error {
	_ = "STUB: not implemented"

	// Index is a noop and always returns nil.
	return nil
}

func (txi *TxIndex) Index(_ *abci.TxResult) error { _ = "STUB: not implemented"; return nil }

func (txi *TxIndex) Search(_ context.Context, _ *query.Query) ([]*abci.TxResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (txi *TxIndex) SetLogger(log.Logger) { _ = "STUB: not implemented"; return }
