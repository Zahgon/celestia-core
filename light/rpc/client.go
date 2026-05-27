package rpc

import (
	"context"
	"errors"
	"time"

	"github.com/cometbft/cometbft/crypto/merkle"
	cmtbytes "github.com/cometbft/cometbft/libs/bytes"
	service "github.com/cometbft/cometbft/libs/service"
	rpcclient "github.com/cometbft/cometbft/rpc/client"
	ctypes "github.com/cometbft/cometbft/rpc/core/types"
	rpctypes "github.com/cometbft/cometbft/rpc/jsonrpc/types"
	"github.com/cometbft/cometbft/types"
)

var errNegOrZeroHeight = errors.New("negative or zero height")

// KeyPathFunc builds a merkle path out of the given path and key.
type KeyPathFunc func(path string, key []byte) (merkle.KeyPath, error)

// LightClient is an interface that contains functionality needed by Client from the light client.
//
//go:generate ../../scripts/mockery_generate.sh LightClient
type LightClient interface {
	ChainID() string
	Update(ctx context.Context, now time.Time) (*types.LightBlock, error)
	VerifyLightBlockAtHeight(ctx context.Context, height int64, now time.Time) (*types.LightBlock, error)
	TrustedLightBlock(height int64) (*types.LightBlock, error)
}

var _ rpcclient.Client = (*Client)(nil)

// Client is an RPC client, which uses light#Client to verify data (if it can
// be proved). Note, merkle.DefaultProofRuntime is used to verify values
// returned by ABCI#Query.
type Client struct {
	service.BaseService

	next rpcclient.Client
	lc   LightClient

	// proof runtime used to verify values returned by ABCIQuery
	prt       *merkle.ProofRuntime
	keyPathFn KeyPathFunc
}

var _ rpcclient.Client = (*Client)(nil)

// Option allow you to tweak Client.
type Option func(*Client)

// KeyPathFn option can be used to set a function, which parses a given path
// and builds the merkle path for the prover. It must be provided if you want
// to call ABCIQuery or ABCIQueryWithOptions.
func KeyPathFn(fn KeyPathFunc) Option { _ = "STUB: not implemented"; return *new(Option) }

// DefaultMerkleKeyPathFn creates a function used to generate merkle key paths
// from a path string and a key. This is the default used by the cosmos SDK.
// This merkle key paths are required when verifying /abci_query calls
func DefaultMerkleKeyPathFn() KeyPathFunc {
	_ = "STUB: not implemented"
	// regexp for extracting store name from /abci_query path
	return *new(KeyPathFunc)
}

// NewClient returns a new client.
func NewClient(next rpcclient.Client, lc LightClient, opts ...Option) *Client {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) OnStart() error { _ = "STUB: not implemented"; return nil }

func (c *Client) OnStop() { _ = "STUB: not implemented"; return }

func (c *Client) Status(ctx context.Context) (*ctypes.ResultStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) ABCIInfo(ctx context.Context) (*ctypes.ResultABCIInfo, error) {
	_ = "STUB: not implemented"
	return nil,

		// ABCIQuery requests proof by default.
		nil
}

func (c *Client) ABCIQuery(ctx context.Context, path string, data cmtbytes.HexBytes) (*ctypes.ResultABCIQuery, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ABCIQueryWithOptions returns an error if opts.Prove is false.
func (c *Client) ABCIQueryWithOptions(ctx context.Context, path string, data cmtbytes.HexBytes,
	opts rpcclient.ABCIQueryOptions) (*ctypes.ResultABCIQuery, error) {
	_ = "STUB: not implemented"

	// always request the proof
	return nil, nil
}

// Validate the response.

// Update the light client if we're behind.
// NOTE: AppHash for height H is in header H+1.

// Validate the value proof against the trusted header.

// 1) build a Merkle key path from path and resp.Key

// 2) verify value

// OR validate the absence proof against the trusted header.

func (c *Client) BroadcastTxCommit(ctx context.Context, tx types.Tx) (*ctypes.ResultBroadcastTxCommit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) BroadcastTxAsync(ctx context.Context, tx types.Tx) (*ctypes.ResultBroadcastTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) BroadcastTxSync(ctx context.Context, tx types.Tx) (*ctypes.ResultBroadcastTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) UnconfirmedTxs(ctx context.Context, limit *int) (*ctypes.ResultUnconfirmedTxs, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) NumUnconfirmedTxs(ctx context.Context) (*ctypes.ResultUnconfirmedTxs, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) CheckTx(ctx context.Context, tx types.Tx) (*ctypes.ResultCheckTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) NetInfo(ctx context.Context) (*ctypes.ResultNetInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) DumpConsensusState(ctx context.Context) (*ctypes.ResultDumpConsensusState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) ConsensusState(ctx context.Context) (*ctypes.ResultConsensusState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) ConsensusParams(ctx context.Context, height *int64) (*ctypes.ResultConsensusParams, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Validate res.

// Update the light client if we're behind.

// Verify hash.

func (c *Client) Health(ctx context.Context) (*ctypes.ResultHealth, error) {
	_ = "STUB: not implemented"
	return nil,

		// BlockchainInfo calls rpcclient#BlockchainInfo and then verifies every header
		// returned.
		nil
}

func (c *Client) BlockchainInfo(ctx context.Context, minHeight, maxHeight int64) (*ctypes.ResultBlockchainInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Validate res.

// Update the light client if we're behind.

// Verify each of the BlockMetas.

func (c *Client) Genesis(ctx context.Context) (*ctypes.ResultGenesis, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) GenesisChunked(ctx context.Context, id uint) (*ctypes.ResultGenesisChunk, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Block calls rpcclient#Block and then verifies the result.
func (c *Client) Block(ctx context.Context, height *int64) (*ctypes.ResultBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Validate res.

// Update the light client if we're behind.

// Verify block.

// BlockByHash calls rpcclient#BlockByHash and then verifies the result.
func (c *Client) BlockByHash(ctx context.Context, hash []byte) (*ctypes.ResultBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Validate res.

// Update the light client if we're behind.

// Verify block.

// BlockResults returns the block results for the given height. If no height is
// provided, the results of the block preceding the latest are returned.
// NOTE: Light client only verifies the tx results
func (c *Client) BlockResults(ctx context.Context, height *int64) (*ctypes.ResultBlockResults, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Can't return the latest block results here because we won't be able to
// prove them. Return the results for the previous block instead.

// Validate res.

// Update the light client if we're behind.

// Build a Merkle tree out of the above 3 binary slices.

// Verify block results.

// Header fetches and verifies the header directly via the light client
func (c *Client) Header(ctx context.Context, height *int64) (*ctypes.ResultHeader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// HeaderByHash calls rpcclient#HeaderByHash and updates the client if it's falling behind.
func (c *Client) HeaderByHash(ctx context.Context, hash cmtbytes.HexBytes) (*ctypes.ResultHeader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:staticcheck

//nolint:staticcheck

func (c *Client) Commit(ctx context.Context, height *int64) (*ctypes.ResultCommit, error) {
	_ = "STUB: not implemented"
	// Update the light client if we're behind and retrieve the light block at the requested height
	// or at the latest height if no height is provided.
	return nil, nil
}

// Tx calls rpcclient#Tx method and then verifies the proof if such was
// requested.
//
// Deprecated: The tx endpoint is deprecated and will be removed in a future release.
func (c *Client) Tx(ctx context.Context, hash []byte, prove bool) (*ctypes.ResultTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Validate res.

// Update the light client if we're behind.

// Validate the proof.

// Deprecated: The tx_search endpoint is deprecated and will be removed in a future release.
func (c *Client) TxSearch(
	ctx context.Context,
	query string,
	prove bool,
	page, perPage *int,
	orderBy string,
) (*ctypes.ResultTxSearch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Deprecated: The block_search endpoint is deprecated and will be removed in a future release.
func (c *Client) BlockSearch(
	ctx context.Context,
	query string,
	page, perPage *int,
	orderBy string,
) (*ctypes.ResultBlockSearch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Validators fetches and verifies validators.
func (c *Client) Validators(
	ctx context.Context,
	height *int64,
	pagePtr, perPagePtr *int,
) (*ctypes.ResultValidators, error) {
	_ = "STUB: not implemented"

	// Update the light client if we're behind and retrieve the light block at the
	// requested height or at the latest height if no height is provided.
	return nil, nil
}

func (c *Client) BroadcastEvidence(ctx context.Context, ev types.Evidence) (*ctypes.ResultBroadcastEvidence, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) Subscribe(ctx context.Context, subscriber, query string,
	outCapacity ...int) (out <-chan ctypes.ResultEvent, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) Unsubscribe(ctx context.Context, subscriber, query string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) UnsubscribeAll(ctx context.Context, subscriber string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) updateLightClientIfNeededTo(ctx context.Context, height *int64) (*types.LightBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) RegisterOpDecoder(typ string, dec merkle.OpDecoder) {
	_ = "STUB: not implemented"
	return
}

// SubscribeWS subscribes for events using the given query and remote address as
// a subscriber, but does not verify responses (UNSAFE)!
// TODO: verify data
func (c *Client) SubscribeWS(ctx *rpctypes.Context, query string) (*ctypes.ResultSubscribe, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// We should have a switch here that performs a validation
// depending on the event's type.

// UnsubscribeWS calls original client's Unsubscribe using remote address as a
// subscriber.
func (c *Client) UnsubscribeWS(ctx *rpctypes.Context, query string) (*ctypes.ResultUnsubscribe, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnsubscribeAllWS calls original client's UnsubscribeAll using remote address
// as a subscriber.
func (c *Client) UnsubscribeAllWS(ctx *rpctypes.Context) (*ctypes.ResultUnsubscribe, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// XXX: Copied from rpc/core/env.go
const (
	// see README
	defaultPerPage = 30
	maxPerPage     = 100
)

func validatePage(pagePtr *int, perPage, totalCount int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// no page parameter

// one page (even if it's empty)

func validatePerPage(perPagePtr *int) int { _ = "STUB: not implemented"; return 0 }

// no per_page parameter

func validateSkipCount(page, perPage int) int { _ = "STUB: not implemented"; return 0 }

// DataCommitment returns the data commitment for the given height.
func (c *Client) DataCommitment(ctx context.Context, start, end uint64) (*ctypes.ResultDataCommitment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) DataRootInclusionProof(ctx context.Context, height, start, end uint64) (*ctypes.ResultDataRootInclusionProof, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) ProveShares(ctx context.Context, height, start, end uint64) (types.ShareProof, error) {
	_ = "STUB: not implemented"
	return *new(types.ShareProof), nil
}

func (c *Client) ProveSharesV2(ctx context.Context, height, start, end uint64) (*ctypes.ResultShareProof, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) TxStatus(ctx context.Context, hash []byte) (*ctypes.ResultTxStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) TxStatusBatch(ctx context.Context, hashes [][]byte) (*ctypes.ResultTxStatusBatch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SignedBlock calls rpcclient#SignedBlock and then verifies the result.
func (c *Client) SignedBlock(ctx context.Context, height *int64) (*ctypes.ResultSignedBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Validate res.

// NOTE: this will re-request the header and commit from the primary. Ideally, you'd just
// fetch the data from the primary and use the light client to verify it.

//nolint:staticcheck

//nolint:staticcheck
