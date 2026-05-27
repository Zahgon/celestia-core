package http

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/cometbft/cometbft/libs/bytes"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cometbft/cometbft/libs/service"
	cmtsync "github.com/cometbft/cometbft/libs/sync"
	rpcclient "github.com/cometbft/cometbft/rpc/client"
	ctypes "github.com/cometbft/cometbft/rpc/core/types"
	jsonrpcclient "github.com/cometbft/cometbft/rpc/jsonrpc/client"
	"github.com/cometbft/cometbft/types"
)

/*
HTTP is a Client implementation that communicates with a CometBFT node over
JSON RPC and WebSockets.

This is the main implementation you probably want to use in production code.
There are other implementations when calling the CometBFT node in-process
(Local), or when you want to mock out the server for test code (mock).

You can subscribe for any event published by CometBFT using Subscribe method.
Note delivery is best-effort. If you don't read events fast enough or network is
slow, CometBFT might cancel the subscription. The client will attempt to
resubscribe (you don't need to do anything). It will keep trying every second
indefinitely until successful.

Request batching is available for JSON RPC requests over HTTP, which conforms to
the JSON RPC specification (https://www.jsonrpc.org/specification#batch). See
the example for more details.

Example:

	c, err := New("http://192.168.1.10:26657", "/websocket")
	if err != nil {
		// handle error
	}

	// call Start/Stop if you're subscribing to events
	err = c.Start()
	if err != nil {
		// handle error
	}
	defer c.Stop()

	res, err := c.Status()
	if err != nil {
		// handle error
	}

	// handle result
*/
type HTTP struct {
	remote string
	rpc    *jsonrpcclient.Client

	*baseRPCClient
	*WSEvents
}

// BatchHTTP provides the same interface as `HTTP`, but allows for batching of
// requests (as per https://www.jsonrpc.org/specification#batch). Do not
// instantiate directly - rather use the HTTP.NewBatch() method to create an
// instance of this struct.
//
// Batching of HTTP requests is thread-safe in the sense that multiple
// goroutines can each create their own batches and send them using the same
// HTTP client. Multiple goroutines could also enqueue transactions in a single
// batch, but ordering of transactions in the batch cannot be guaranteed in such
// an example.
type BatchHTTP struct {
	rpcBatch *jsonrpcclient.RequestBatch
	*baseRPCClient
}

// rpcClient is an internal interface to which our RPC clients (batch and
// non-batch) must conform. Acts as an additional code-level sanity check to
// make sure the implementations stay coherent.
type rpcClient interface {
	rpcclient.ABCIClient
	rpcclient.HistoryClient
	rpcclient.NetworkClient
	rpcclient.SignClient
	rpcclient.StatusClient
}

// baseRPCClient implements the basic RPC method logic without the actual
// underlying RPC call functionality, which is provided by `caller`.
type baseRPCClient struct {
	caller jsonrpcclient.Caller
}

var (
	_ rpcClient = (*HTTP)(nil)
	_ rpcClient = (*BatchHTTP)(nil)
	_ rpcClient = (*baseRPCClient)(nil)
)

//-----------------------------------------------------------------------------
// HTTP

// New takes a remote endpoint in the form <protocol>://<host>:<port> and
// the websocket path (which always seems to be "/websocket")
// An error is returned on invalid remote. The function panics when remote is nil.
func New(remote, wsEndpoint string) (*HTTP, error) { _ = "STUB: not implemented"; return nil, nil }

// Create timeout enabled http client
func NewWithTimeout(remote, wsEndpoint string, timeout uint) (*HTTP, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewWithClient allows for setting a custom http client (See New).
// An error is returned on invalid remote. The function panics when remote is nil.
func NewWithClient(remote, wsEndpoint string, client *http.Client) (*HTTP, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var _ rpcclient.Client = (*HTTP)(nil)

// SetLogger sets a logger.
func (c *HTTP) SetLogger(l log.Logger) { _ = "STUB: not implemented"; return }

// Remote returns the remote network address in a string form.
func (c *HTTP) Remote() string {
	_ = "STUB: not implemented"

	// NewBatch creates a new batch client for this HTTP client.
	return ""
}

func (c *HTTP) NewBatch() *BatchHTTP { _ = "STUB: not implemented"; return nil }

//-----------------------------------------------------------------------------
// BatchHTTP

// Send is a convenience function for an HTTP batch that will trigger the
// compilation of the batched requests and send them off using the client as a
// single request. On success, this returns a list of the deserialized results
// from each request in the sent batch.
func (b *BatchHTTP) Send(ctx context.Context) ([]interface{}, error) {
	_ = "STUB: not implemented"
	return nil,

		// Clear will empty out this batch of requests and return the number of requests
		// that were cleared out.
		nil
}

func (b *BatchHTTP) Clear() int { _ = "STUB: not implemented"; return 0 }

// Count returns the number of enqueued requests waiting to be sent.
func (b *BatchHTTP) Count() int { _ = "STUB: not implemented"; return 0 }

//-----------------------------------------------------------------------------
// baseRPCClient

func (c *baseRPCClient) Status(ctx context.Context) (*ctypes.ResultStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *baseRPCClient) ABCIInfo(ctx context.Context) (*ctypes.ResultABCIInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *baseRPCClient) ABCIQuery(
	ctx context.Context,
	path string,
	data bytes.HexBytes,
) (*ctypes.ResultABCIQuery, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *baseRPCClient) ABCIQueryWithOptions(
	ctx context.Context,
	path string,
	data bytes.HexBytes,
	opts rpcclient.ABCIQueryOptions,
) (*ctypes.ResultABCIQuery, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *baseRPCClient) BroadcastTxCommit(
	ctx context.Context,
	tx types.Tx,
) (*ctypes.ResultBroadcastTxCommit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *baseRPCClient) BroadcastTxAsync(
	ctx context.Context,
	tx types.Tx,
) (*ctypes.ResultBroadcastTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *baseRPCClient) BroadcastTxSync(
	ctx context.Context,
	tx types.Tx,
) (*ctypes.ResultBroadcastTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *baseRPCClient) broadcastTX(
	ctx context.Context,
	route string,
	tx types.Tx,
) (*ctypes.ResultBroadcastTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *baseRPCClient) UnconfirmedTxs(
	ctx context.Context,
	limit *int,
) (*ctypes.ResultUnconfirmedTxs, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *baseRPCClient) NumUnconfirmedTxs(ctx context.Context) (*ctypes.ResultUnconfirmedTxs, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *baseRPCClient) CheckTx(ctx context.Context, tx types.Tx) (*ctypes.ResultCheckTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *baseRPCClient) NetInfo(ctx context.Context) (*ctypes.ResultNetInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *baseRPCClient) DumpConsensusState(ctx context.Context) (*ctypes.ResultDumpConsensusState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *baseRPCClient) ConsensusState(ctx context.Context) (*ctypes.ResultConsensusState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *baseRPCClient) ConsensusParams(
	ctx context.Context,
	height *int64,
) (*ctypes.ResultConsensusParams, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *baseRPCClient) Health(ctx context.Context) (*ctypes.ResultHealth, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *baseRPCClient) BlockchainInfo(
	ctx context.Context,
	minHeight,
	maxHeight int64,
) (*ctypes.ResultBlockchainInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *baseRPCClient) Genesis(ctx context.Context) (*ctypes.ResultGenesis, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *baseRPCClient) GenesisChunked(ctx context.Context, id uint) (*ctypes.ResultGenesisChunk, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *baseRPCClient) Block(ctx context.Context, height *int64) (*ctypes.ResultBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *baseRPCClient) BlockByHash(ctx context.Context, hash []byte) (*ctypes.ResultBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *baseRPCClient) BlockResults(
	ctx context.Context,
	height *int64,
) (*ctypes.ResultBlockResults, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *baseRPCClient) Header(ctx context.Context, height *int64) (*ctypes.ResultHeader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *baseRPCClient) HeaderByHash(ctx context.Context, hash bytes.HexBytes) (*ctypes.ResultHeader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *baseRPCClient) Commit(ctx context.Context, height *int64) (*ctypes.ResultCommit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Deprecated: The tx endpoint is deprecated and will be removed in a future release.
func (c *baseRPCClient) Tx(ctx context.Context, hash []byte, prove bool) (*ctypes.ResultTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Deprecated: The tx_search endpoint is deprecated and will be removed in a future release.
func (c *baseRPCClient) TxSearch(
	ctx context.Context,
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
func (c *baseRPCClient) BlockSearch(
	ctx context.Context,
	query string,
	page, perPage *int,
	orderBy string,
) (*ctypes.ResultBlockSearch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *baseRPCClient) Validators(
	ctx context.Context,
	height *int64,
	page,
	perPage *int,
) (*ctypes.ResultValidators, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *baseRPCClient) BroadcastEvidence(
	ctx context.Context,
	ev types.Evidence,
) (*ctypes.ResultBroadcastEvidence, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *baseRPCClient) SignedBlock(ctx context.Context, height *int64) (*ctypes.ResultSignedBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *baseRPCClient) DataCommitment(
	ctx context.Context,
	start uint64,
	end uint64,
) (*ctypes.ResultDataCommitment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *baseRPCClient) TxStatus(
	ctx context.Context,
	hash []byte,
) (*ctypes.ResultTxStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *baseRPCClient) TxStatusBatch(
	ctx context.Context,
	hashes [][]byte,
) (*ctypes.ResultTxStatusBatch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *baseRPCClient) DataRootInclusionProof(
	ctx context.Context,
	height uint64,
	start uint64,
	end uint64,
) (*ctypes.ResultDataRootInclusionProof, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ProveShares
// Deprecated: Use ProveSharesV2 instead.
func (c *baseRPCClient) ProveShares(
	ctx context.Context,
	height uint64,
	startShare uint64,
	endShare uint64,
) (types.ShareProof, error) {
	_ = "STUB: not implemented"
	return *new(types.ShareProof), nil
}

func (c *baseRPCClient) ProveSharesV2(
	ctx context.Context,
	height uint64,
	startShare uint64,
	endShare uint64,
) (*ctypes.ResultShareProof, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//-----------------------------------------------------------------------------
// WSEvents

var errNotRunning = errors.New("client is not running. Use .Start() method to start")

// WSEvents is a wrapper around WSClient, which implements EventsClient.
type WSEvents struct {
	service.BaseService
	remote   string
	endpoint string
	ws       *jsonrpcclient.WSClient

	mtx           cmtsync.RWMutex
	subscriptions map[string]chan ctypes.ResultEvent // query -> chan
}

func newWSEvents(remote, endpoint string) (*WSEvents, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// resubscribe immediately

// OnStart implements service.Service by starting WSClient and event loop.
func (w *WSEvents) OnStart() error { _ = "STUB: not implemented"; return nil }

// OnStop implements service.Service by stopping WSClient.
func (w *WSEvents) OnStop() { _ = "STUB: not implemented"; return }

// Subscribe implements EventsClient by using WSClient to subscribe given
// subscriber to query. By default, returns a channel with cap=1. Error is
// returned if it fails to subscribe.
//
// Channel is never closed to prevent clients from seeing an erroneous event.
//
// It returns an error if WSEvents is not running.
func (w *WSEvents) Subscribe(ctx context.Context, _, query string,
	outCapacity ...int,
) (out <-chan ctypes.ResultEvent, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// subscriber param is ignored because CometBFT will override it with
// remote IP anyway.

// Unsubscribe implements EventsClient by using WSClient to unsubscribe given
// subscriber from query.
//
// It returns an error if WSEvents is not running.
func (w *WSEvents) Unsubscribe(ctx context.Context, _, query string) error {
	_ = "STUB: not implemented"
	return nil
}

// UnsubscribeAll implements EventsClient by using WSClient to unsubscribe
// given subscriber from all the queries.
//
// It returns an error if WSEvents is not running.
func (w *WSEvents) UnsubscribeAll(ctx context.Context, _ string) error {
	_ = "STUB: not implemented"
	return nil
}

// After being reconnected, it is necessary to redo subscription to server
// otherwise no data will be automatically received.
func (w *WSEvents) redoSubscriptionsAfter(d time.Duration) { _ = "STUB: not implemented"; return }

func isErrAlreadySubscribed(err error) bool { _ = "STUB: not implemented"; return false }

func (w *WSEvents) eventListener() { _ = "STUB: not implemented"; return }

// Error can be ErrAlreadySubscribed or max client (subscriptions per
// client) reached or CometBFT exited.
// We can ignore ErrAlreadySubscribed, but need to retry in other
// cases.

// Resubscribe after 1 second to give CometBFT time to restart (if
// crashed).
