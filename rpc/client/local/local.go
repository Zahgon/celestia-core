package local

import (
	"context"

	"github.com/cometbft/cometbft/libs/bytes"
	"github.com/cometbft/cometbft/libs/log"
	cmtpubsub "github.com/cometbft/cometbft/libs/pubsub"
	nm "github.com/cometbft/cometbft/node"
	rpcclient "github.com/cometbft/cometbft/rpc/client"
	"github.com/cometbft/cometbft/rpc/core"
	ctypes "github.com/cometbft/cometbft/rpc/core/types"
	rpctypes "github.com/cometbft/cometbft/rpc/jsonrpc/types"
	"github.com/cometbft/cometbft/types"
)

/*
Local is a Client implementation that directly executes the rpc
functions on a given node, without going through HTTP or GRPC.

This implementation is useful for:

* Running tests against a node in-process without the overhead
of going through an http server
* Communication between an ABCI app and CometBFT when they
are compiled in process.

For real clients, you probably want to use client.HTTP.  For more
powerful control during testing, you probably want the "client/mock" package.

You can subscribe for any event published by CometBFT using Subscribe method.
Note delivery is best-effort. If you don't read events fast enough, CometBFT
might cancel the subscription. The client will attempt to resubscribe (you
don't need to do anything). It will keep trying indefinitely with exponential
backoff (10ms -> 20ms -> 40ms) until successful.
*/
type Local struct {
	*types.EventBus
	Logger log.Logger
	ctx    *rpctypes.Context
	env    *core.Environment
}

// NewLocal configures a client that calls the Node directly.
func New(node *nm.Node) *Local { _ = "STUB: not implemented"; return nil }

var _ rpcclient.Client = (*Local)(nil)

// SetLogger allows to set a logger on the client.
func (c *Local) SetLogger(l log.Logger) { _ = "STUB: not implemented"; return }

func (c *Local) Status(context.Context) (*ctypes.ResultStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Local) ABCIInfo(context.Context) (*ctypes.ResultABCIInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Local) ABCIQuery(ctx context.Context, path string, data bytes.HexBytes) (*ctypes.ResultABCIQuery, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Local) ABCIQueryWithOptions(
	_ context.Context,
	path string,
	data bytes.HexBytes,
	opts rpcclient.ABCIQueryOptions,
) (*ctypes.ResultABCIQuery, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Local) BroadcastTxCommit(_ context.Context, tx types.Tx) (*ctypes.ResultBroadcastTxCommit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Local) BroadcastTxAsync(_ context.Context, tx types.Tx) (*ctypes.ResultBroadcastTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Local) BroadcastTxSync(_ context.Context, tx types.Tx) (*ctypes.ResultBroadcastTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Local) UnconfirmedTxs(_ context.Context, limit *int) (*ctypes.ResultUnconfirmedTxs, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Local) NumUnconfirmedTxs(context.Context) (*ctypes.ResultUnconfirmedTxs, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Local) CheckTx(_ context.Context, tx types.Tx) (*ctypes.ResultCheckTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Local) NetInfo(context.Context) (*ctypes.ResultNetInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Local) DumpConsensusState(context.Context) (*ctypes.ResultDumpConsensusState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Local) ConsensusState(context.Context) (*ctypes.ResultConsensusState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Local) ConsensusParams(_ context.Context, height *int64) (*ctypes.ResultConsensusParams, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Local) Health(context.Context) (*ctypes.ResultHealth, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Local) DialSeeds(_ context.Context, seeds []string) (*ctypes.ResultDialSeeds, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Local) DialPeers(
	_ context.Context,
	peers []string,
	persistent,
	unconditional,
	private bool,
) (*ctypes.ResultDialPeers, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Local) BlockchainInfo(_ context.Context, minHeight, maxHeight int64) (*ctypes.ResultBlockchainInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Local) Genesis(context.Context) (*ctypes.ResultGenesis, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Local) GenesisChunked(_ context.Context, id uint) (*ctypes.ResultGenesisChunk, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Local) Block(_ context.Context, height *int64) (*ctypes.ResultBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Local) BlockByHash(_ context.Context, hash []byte) (*ctypes.ResultBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Local) BlockResults(_ context.Context, height *int64) (*ctypes.ResultBlockResults, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Local) Header(_ context.Context, height *int64) (*ctypes.ResultHeader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Local) HeaderByHash(_ context.Context, hash bytes.HexBytes) (*ctypes.ResultHeader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Local) Commit(_ context.Context, height *int64) (*ctypes.ResultCommit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Local) Validators(_ context.Context, height *int64, page, perPage *int) (*ctypes.ResultValidators, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Deprecated: The tx endpoint is deprecated and will be removed in a future release.
func (c *Local) Tx(_ context.Context, hash []byte, prove bool) (*ctypes.ResultTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Deprecated: The tx_search endpoint is deprecated and will be removed in a future release.
func (c *Local) TxSearch(
	_ context.Context,
	query string,
	prove bool,
	page,
	perPage *int,
	orderBy string,
) (*ctypes.ResultTxSearch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Deprecated: The block_search endpoint is deprecated and will be removed in a future release.
func (c *Local) BlockSearch(
	_ context.Context,
	query string,
	page, perPage *int,
	orderBy string,
) (*ctypes.ResultBlockSearch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Local) BroadcastEvidence(_ context.Context, ev types.Evidence) (*ctypes.ResultBroadcastEvidence, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Local) Subscribe(
	ctx context.Context,
	subscriber,
	query string,
	outCapacity ...int,
) (out <-chan ctypes.ResultEvent, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:staticcheck

func (c *Local) eventsRoutine(
	sub types.Subscription,
	subscriber string,
	q cmtpubsub.Query,
	outc chan<- ctypes.ResultEvent,
) {
	_ = "STUB: not implemented"
	return
}

// client was stopped

// Try to resubscribe with exponential backoff.
func (c *Local) resubscribe(subscriber string, q cmtpubsub.Query) types.Subscription {
	_ = "STUB: not implemented"
	return *new(types.Subscription)
}

// 10ms -> 20ms -> 40ms

func (c *Local) Unsubscribe(ctx context.Context, subscriber, query string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Local) UnsubscribeAll(ctx context.Context, subscriber string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Local) SignedBlock(ctx context.Context, height *int64) (*ctypes.ResultSignedBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Local) DataCommitment(
	_ context.Context,
	start uint64,
	end uint64,
) (*ctypes.ResultDataCommitment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Local) DataRootInclusionProof(
	_ context.Context,
	height uint64,
	start uint64,
	end uint64,
) (*ctypes.ResultDataRootInclusionProof, error) {
	_ = "STUB: not implemented"
	//nolint:gosec
	return nil, nil
}

// ProveShares
// Deprecated: Use ProveSharesV2 instead.
func (c *Local) ProveShares(
	ctx context.Context,
	height uint64,
	startShare uint64,
	endShare uint64,
) (types.ShareProof, error) {
	_ = "STUB: not implemented"
	//nolint:gosec
	return *new(types.ShareProof), nil
}

func (c *Local) ProveSharesV2(
	ctx context.Context,
	height uint64,
	startShare uint64,
	endShare uint64,
) (*ctypes.ResultShareProof, error) {
	_ = "STUB: not implemented"
	//nolint:gosec
	return nil, nil
}

func (c *Local) TxStatus(ctx context.Context, hash []byte) (*ctypes.ResultTxStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Local) TxStatusBatch(ctx context.Context, hashes [][]byte) (*ctypes.ResultTxStatusBatch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
