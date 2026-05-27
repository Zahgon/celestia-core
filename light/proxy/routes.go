package proxy

import (
	"github.com/cometbft/cometbft/libs/bytes"
	lrpc "github.com/cometbft/cometbft/light/rpc"
	ctypes "github.com/cometbft/cometbft/rpc/core/types"
	rpcserver "github.com/cometbft/cometbft/rpc/jsonrpc/server"
	rpctypes "github.com/cometbft/cometbft/rpc/jsonrpc/types"
	"github.com/cometbft/cometbft/types"
)

func RPCRoutes(c *lrpc.Client) map[string]*rpcserver.RPCFunc { _ = "STUB: not implemented"; return nil }

// Subscribe/unsubscribe are reserved for websocket events.

// info API

// tx broadcast API

// abci API

// evidence API

// Celestia Specific RPCs

type rpcHealthFunc func(ctx *rpctypes.Context) (*ctypes.ResultHealth, error)

func makeHealthFunc(c *lrpc.Client) rpcHealthFunc {
	_ = "STUB: not implemented"
	return *new(rpcHealthFunc)
}

type rpcStatusFunc func(ctx *rpctypes.Context) (*ctypes.ResultStatus, error)

func makeStatusFunc(c *lrpc.Client) rpcStatusFunc {
	_ = "STUB: not implemented"
	return *new(rpcStatusFunc)
}

type rpcNetInfoFunc func(ctx *rpctypes.Context, minHeight, maxHeight int64) (*ctypes.ResultNetInfo, error)

func makeNetInfoFunc(c *lrpc.Client) rpcNetInfoFunc {
	_ = "STUB: not implemented"
	return *new(rpcNetInfoFunc)
}

type rpcBlockchainInfoFunc func(ctx *rpctypes.Context, minHeight, maxHeight int64) (*ctypes.ResultBlockchainInfo, error)

func makeBlockchainInfoFunc(c *lrpc.Client) rpcBlockchainInfoFunc {
	_ = "STUB: not implemented"
	return *new(rpcBlockchainInfoFunc)
}

type rpcGenesisFunc func(ctx *rpctypes.Context) (*ctypes.ResultGenesis, error)

func makeGenesisFunc(c *lrpc.Client) rpcGenesisFunc {
	_ = "STUB: not implemented"
	return *new(rpcGenesisFunc)
}

type rpcGenesisChunkedFunc func(ctx *rpctypes.Context, chunk uint) (*ctypes.ResultGenesisChunk, error)

func makeGenesisChunkedFunc(c *lrpc.Client) rpcGenesisChunkedFunc {
	_ = "STUB: not implemented"
	return *new(rpcGenesisChunkedFunc)
}

type rpcBlockFunc func(ctx *rpctypes.Context, height *int64) (*ctypes.ResultBlock, error)

func makeBlockFunc(c *lrpc.Client) rpcBlockFunc {
	_ = "STUB: not implemented"
	return *new(rpcBlockFunc)
}

type rpcHeaderFunc func(ctx *rpctypes.Context, height *int64) (*ctypes.ResultHeader, error)

func makeHeaderFunc(c *lrpc.Client) rpcHeaderFunc {
	_ = "STUB: not implemented"
	return *new(rpcHeaderFunc)
}

type rpcHeaderByHashFunc func(ctx *rpctypes.Context, hash []byte) (*ctypes.ResultHeader, error)

func makeHeaderByHashFunc(c *lrpc.Client) rpcHeaderByHashFunc {
	_ = "STUB: not implemented"
	return *new(rpcHeaderByHashFunc)
}

type rpcBlockByHashFunc func(ctx *rpctypes.Context, hash []byte) (*ctypes.ResultBlock, error)

func makeBlockByHashFunc(c *lrpc.Client) rpcBlockByHashFunc {
	_ = "STUB: not implemented"
	return *new(rpcBlockByHashFunc)
}

type rpcBlockResultsFunc func(ctx *rpctypes.Context, height *int64) (*ctypes.ResultBlockResults, error)

func makeBlockResultsFunc(c *lrpc.Client) rpcBlockResultsFunc {
	_ = "STUB: not implemented"
	return *new(rpcBlockResultsFunc)
}

type rpcCommitFunc func(ctx *rpctypes.Context, height *int64) (*ctypes.ResultCommit, error)

func makeCommitFunc(c *lrpc.Client) rpcCommitFunc {
	_ = "STUB: not implemented"
	return *new(rpcCommitFunc)
}

//nolint:staticcheck // Deprecated: will be removed in a future release.
type rpcTxFunc func(ctx *rpctypes.Context, hash []byte, prove bool) (*ctypes.ResultTx, error)

//nolint:staticcheck // Deprecated: will be removed in a future release.
func makeTxFunc(c *lrpc.Client) rpcTxFunc { _ = "STUB: not implemented"; return *new(rpcTxFunc) }

//nolint:staticcheck // Deprecated: will be removed in a future release.
type rpcTxSearchFunc func(
	ctx *rpctypes.Context,
	query string,
	prove bool,
	page, perPage *int,
	orderBy string,
) (*ctypes.ResultTxSearch, error)

//nolint:staticcheck // Deprecated: will be removed in a future release.
func makeTxSearchFunc(c *lrpc.Client) rpcTxSearchFunc {
	_ = "STUB: not implemented"
	return *new(rpcTxSearchFunc)
}

//nolint:staticcheck // Deprecated: will be removed in a future release.
type rpcBlockSearchFunc func(
	ctx *rpctypes.Context,
	query string,
	prove bool,
	page, perPage *int,
	orderBy string,
) (*ctypes.ResultBlockSearch, error)

//nolint:staticcheck // Deprecated: will be removed in a future release.
func makeBlockSearchFunc(c *lrpc.Client) rpcBlockSearchFunc {
	_ = "STUB: not implemented"
	return *new(rpcBlockSearchFunc)
}

type rpcValidatorsFunc func(ctx *rpctypes.Context, height *int64,
	page, perPage *int) (*ctypes.ResultValidators, error)

func makeValidatorsFunc(c *lrpc.Client) rpcValidatorsFunc {
	_ = "STUB: not implemented"
	return *new(rpcValidatorsFunc)
}

type rpcDumpConsensusStateFunc func(ctx *rpctypes.Context) (*ctypes.ResultDumpConsensusState, error)

func makeDumpConsensusStateFunc(c *lrpc.Client) rpcDumpConsensusStateFunc {
	_ = "STUB: not implemented"
	return *new(rpcDumpConsensusStateFunc)
}

type rpcConsensusStateFunc func(ctx *rpctypes.Context) (*ctypes.ResultConsensusState, error)

func makeConsensusStateFunc(c *lrpc.Client) rpcConsensusStateFunc {
	_ = "STUB: not implemented"
	return *new(rpcConsensusStateFunc)
}

type rpcConsensusParamsFunc func(ctx *rpctypes.Context, height *int64) (*ctypes.ResultConsensusParams, error)

func makeConsensusParamsFunc(c *lrpc.Client) rpcConsensusParamsFunc {
	_ = "STUB: not implemented"
	return *new(rpcConsensusParamsFunc)
}

type rpcUnconfirmedTxsFunc func(ctx *rpctypes.Context, limit *int) (*ctypes.ResultUnconfirmedTxs, error)

func makeUnconfirmedTxsFunc(c *lrpc.Client) rpcUnconfirmedTxsFunc {
	_ = "STUB: not implemented"
	return *new(rpcUnconfirmedTxsFunc)
}

type rpcNumUnconfirmedTxsFunc func(ctx *rpctypes.Context) (*ctypes.ResultUnconfirmedTxs, error)

func makeNumUnconfirmedTxsFunc(c *lrpc.Client) rpcNumUnconfirmedTxsFunc {
	_ = "STUB: not implemented"
	return *new(rpcNumUnconfirmedTxsFunc)
}

type rpcBroadcastTxCommitFunc func(ctx *rpctypes.Context, tx types.Tx) (*ctypes.ResultBroadcastTxCommit, error)

func makeBroadcastTxCommitFunc(c *lrpc.Client) rpcBroadcastTxCommitFunc {
	_ = "STUB: not implemented"
	return *new(rpcBroadcastTxCommitFunc)
}

type rpcBroadcastTxSyncFunc func(ctx *rpctypes.Context, tx types.Tx) (*ctypes.ResultBroadcastTx, error)

func makeBroadcastTxSyncFunc(c *lrpc.Client) rpcBroadcastTxSyncFunc {
	_ = "STUB: not implemented"
	return *new(rpcBroadcastTxSyncFunc)
}

type rpcBroadcastTxAsyncFunc func(ctx *rpctypes.Context, tx types.Tx) (*ctypes.ResultBroadcastTx, error)

func makeBroadcastTxAsyncFunc(c *lrpc.Client) rpcBroadcastTxAsyncFunc {
	_ = "STUB: not implemented"
	return *new(rpcBroadcastTxAsyncFunc)
}

type rpcABCIQueryFunc func(ctx *rpctypes.Context, path string,
	data bytes.HexBytes, height int64, prove bool) (*ctypes.ResultABCIQuery, error)

func makeABCIQueryFunc(c *lrpc.Client) rpcABCIQueryFunc {
	_ = "STUB: not implemented"
	return *new(rpcABCIQueryFunc)
}

type rpcABCIInfoFunc func(ctx *rpctypes.Context) (*ctypes.ResultABCIInfo, error)

func makeABCIInfoFunc(c *lrpc.Client) rpcABCIInfoFunc {
	_ = "STUB: not implemented"
	return *new(rpcABCIInfoFunc)
}

type rpcBroadcastEvidenceFunc func(ctx *rpctypes.Context, ev types.Evidence) (*ctypes.ResultBroadcastEvidence, error)

func makeBroadcastEvidenceFunc(c *lrpc.Client) rpcBroadcastEvidenceFunc {
	_ = "STUB: not implemented"
	return *new(rpcBroadcastEvidenceFunc)
}

type rpcTxStatusFunc func(ctx *rpctypes.Context, hash []byte) (*ctypes.ResultTxStatus, error)

func makeTxStatusFunc(c *lrpc.Client) rpcTxStatusFunc {
	_ = "STUB: not implemented"
	return *new(rpcTxStatusFunc)
}

type rpcTxStatusBatchFunc func(ctx *rpctypes.Context, hashes [][]byte) (*ctypes.ResultTxStatusBatch, error)

func makeTxStatusBatchFunc(c *lrpc.Client) rpcTxStatusBatchFunc {
	_ = "STUB: not implemented"
	return *new(rpcTxStatusBatchFunc)
}
