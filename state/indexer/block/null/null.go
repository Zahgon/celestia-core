package null

import (
	"context"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cometbft/cometbft/libs/pubsub/query"
	"github.com/cometbft/cometbft/state/indexer"
	"github.com/cometbft/cometbft/types"
)

var _ indexer.BlockIndexer = (*BlockerIndexer)(nil)

// TxIndex implements a no-op block indexer.
type BlockerIndexer struct{}

func (idx *BlockerIndexer) Has(int64) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (idx *BlockerIndexer) Index(types.EventDataNewBlockEvents) error {
	_ = "STUB: not implemented"
	return nil
}

func (idx *BlockerIndexer) Search(context.Context, *query.Query) ([]int64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (idx *BlockerIndexer) SetLogger(log.Logger) { _ = "STUB: not implemented"; return }
