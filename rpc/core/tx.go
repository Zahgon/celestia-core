package core

import (
	"context"

	abcitypes "github.com/cometbft/cometbft/abci/types"
	cmtquery "github.com/cometbft/cometbft/libs/pubsub/query"
	ctypes "github.com/cometbft/cometbft/rpc/core/types"
	rpctypes "github.com/cometbft/cometbft/rpc/jsonrpc/types"
	"github.com/cometbft/cometbft/state"
	"github.com/cometbft/cometbft/types"
)

const (
	TxStatusUnknown   string = "UNKNOWN"
	TxStatusPending   string = "PENDING"
	TxStatusEvicted   string = "EVICTED"
	TxStatusRejected  string = "REJECTED"
	TxStatusCommitted string = "COMMITTED"

	// MaxTxStatusBatchSize max number of tx hashes queried in a single batch request.
	MaxTxStatusBatchSize = 20
)

// Tx allows you to query the transaction results. `nil` could mean the
// transaction is in the mempool, invalidated, or was not sent in the first
// place.
// More: https://docs.cometbft.com/v0.38.x/rpc/#/Info/tx
//
// Deprecated: The tx endpoint is deprecated and will be removed in a future release.
func (env *Environment) Tx(_ *rpctypes.Context, hash []byte, prove bool) (*ctypes.ResultTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if index is disabled, return error

// TxSearch allows you to query for multiple transactions results. It returns a
// list of transactions (maximum ?per_page entries) and the total count.
// More: https://docs.cometbft.com/v0.38.x/rpc/#/Info/tx_search
//
// Deprecated: The tx_search endpoint is deprecated and will be removed in a future release.
func (env *Environment) TxSearch(
	ctx *rpctypes.Context,
	query string,
	prove bool,
	pagePtr, perPagePtr *int,
	orderBy string,
) (*ctypes.ResultTxSearch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if index is disabled, return error

func (env *Environment) txSearchPage(
	ctx context.Context,
	q *cmtquery.Query,
	pagePtr *int,
	perPage int,
	orderBy string,
) ([]*abcitypes.TxResult, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// sort results (must be done before pagination)

// paginate results

// Deprecated: helper for the deprecated tx and tx_search endpoints.
func (env *Environment) txResultsToRPC(results []*abcitypes.TxResult, prove bool) ([]*ctypes.ResultTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (env *Environment) proveTx(height int64, index uint32) (types.ShareProof, error) {
	_ = "STUB: not implemented"
	return *new(types.ShareProof), nil
}

// ProveShares creates an NMT proof for a set of shares to a set of rows. It is
// end exclusive.
// Deprecated: Use ProveSharesV2 instead.
func (env *Environment) ProveShares(
	_ *rpctypes.Context,
	height int64,
	startShare uint64,
	endShare uint64,
) (types.ShareProof, error) {
	_ = "STUB: not implemented"
	return *new(types.ShareProof), nil
}

// we can make the assumption that for custom queries, if the value is nil
// and some logs have been emitted, then an error happened.

// TxStatus retrieves the status of a transaction by its hash. It returns a ResultTxStatus
// with the transaction's height and index if committed, or its pending, evicted, or unknown status.
// It also includes the execution code and log for failed txs.
func (env *Environment) TxStatus(ctx *rpctypes.Context, hash []byte) (*ctypes.ResultTxStatus, error) {
	_ = "STUB: not implemented"

	// Check if the tx has been committed
	return nil, nil
}

// Get the tx key from the hash

// Check if the tx is in the mempool

// Check if the tx is evicted

// Check if the tx was rejected (this is only the case for recheck-tx)

// If the tx is not in the mempool, evicted, or committed, return unknown.
// This can happen in the following cases:
// - Tx was never submitted to this node
// - Tx was evicted/rejected and has expired from the cache
// - Tx was submitted to a different node and not yet propagated
// - Tx is invalid and was immediately rejected without caching

// TxStatusBatch returns the status of each queried tx and info with their hashes.
func (env *Environment) TxStatusBatch(ctx *rpctypes.Context, hashes [][]byte) (*ctypes.ResultTxStatusBatch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ProveSharesV2 creates a proof for a set of shares to the data root.
// The range is end exclusive.
func (env *Environment) ProveSharesV2(
	ctx *rpctypes.Context,
	height int64,
	startShare uint64,
	endShare uint64,
) (*ctypes.ResultShareProof, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func loadRawBlock(bs state.BlockStore, height int64) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If the part is missing (e.g. since it has been deleted after we
// loaded the block meta) we consider the whole block to be missing.
