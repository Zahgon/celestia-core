package block

import (
	"github.com/cometbft/cometbft/config"
	"github.com/cometbft/cometbft/state/indexer"
	"github.com/cometbft/cometbft/state/txindex"
)

// IndexerFromConfig constructs a slice of indexer.EventSink using the provided
// configuration.
func IndexerFromConfig(cfg *config.Config, dbProvider config.DBProvider, chainID string) (
	txIdx txindex.TxIndexer, blockIdx indexer.BlockIndexer, err error,
) {
	_ = "STUB: not implemented"
	return *new(txindex.TxIndexer), *new(indexer.BlockIndexer), nil
}

// IndexerFromConfigWithDisabledIndexers constructs a slice of indexer.EventSink using the provided
// configuration. If all indexers are disabled in the configuration, it returns null indexers.
// Otherwise, it creates the appropriate indexers based on the configuration.
func IndexerFromConfigWithDisabledIndexers(cfg *config.Config, dbProvider config.DBProvider, chainID string) (
	txIdx txindex.TxIndexer, blockIdx indexer.BlockIndexer, allIndexersDisabled bool, err error,
) {
	_ = "STUB: not implemented"
	return *new(txindex.TxIndexer), *new(indexer.BlockIndexer), false, nil
}

// Deprecated: the "kv" indexer is deprecated and will be removed in a future release.
