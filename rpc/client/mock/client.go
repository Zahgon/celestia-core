package mock

/*
package mock returns a Client implementation that
accepts various (mock) implementations of the various methods.

This implementation is useful for using in tests, when you don't
need a real server, but want a high-level of control about
the server response you want to mock (eg. error handling),
or if you just want to record the calls to verify in your tests.

For real clients, you probably want the "http" package.  If you
want to directly call a CometBFT node in process, you can use the
"local" package.
*/

import (
	"context"

	"github.com/cometbft/cometbft/libs/bytes"
	"github.com/cometbft/cometbft/libs/service"
	"github.com/cometbft/cometbft/rpc/client"
	"github.com/cometbft/cometbft/rpc/core"
	ctypes "github.com/cometbft/cometbft/rpc/core/types"
	"github.com/cometbft/cometbft/types"
)

// Client wraps arbitrary implementations of the various interfaces.
type Client struct {
	client.ABCIClient
	client.SignClient
	client.HistoryClient
	client.StatusClient
	client.EventsClient
	client.EvidenceClient
	client.MempoolClient
	service.Service

	env *core.Environment
}

func New() Client { _ = "STUB: not implemented"; return *new(Client) }

var _ client.Client = Client{}

// Call is used by recorders to save a call and response.
// It can also be used to configure mock responses.
type Call struct {
	Name     string
	Args     interface{}
	Response interface{}
	Error    error
}

// GetResponse will generate the apporiate response for us, when
// using the Call struct to configure a Mock handler.
//
// When configuring a response, if only one of Response or Error is
// set then that will always be returned. If both are set, then
// we return Response if the Args match the set args, Error otherwise.
func (c Call) GetResponse(args interface{}) (interface{}, error) {
	_ = "STUB: not implemented"
	// handle the case with no response
	return nil, nil
}

// response without error

// have both, we must check args....

func (c Client) Status(context.Context) (*ctypes.ResultStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c Client) ABCIInfo(context.Context) (*ctypes.ResultABCIInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c Client) ABCIQuery(ctx context.Context, path string, data bytes.HexBytes) (*ctypes.ResultABCIQuery, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c Client) ABCIQueryWithOptions(
	_ context.Context,
	path string,
	data bytes.HexBytes,
	opts client.ABCIQueryOptions,
) (*ctypes.ResultABCIQuery, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c Client) BroadcastTxCommit(_ context.Context, tx types.Tx) (*ctypes.ResultBroadcastTxCommit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c Client) BroadcastTxAsync(_ context.Context, tx types.Tx) (*ctypes.ResultBroadcastTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c Client) BroadcastTxSync(_ context.Context, tx types.Tx) (*ctypes.ResultBroadcastTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c Client) CheckTx(_ context.Context, tx types.Tx) (*ctypes.ResultCheckTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c Client) NetInfo(_ context.Context) (*ctypes.ResultNetInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c Client) ConsensusState(_ context.Context) (*ctypes.ResultConsensusState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c Client) DumpConsensusState(_ context.Context) (*ctypes.ResultDumpConsensusState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c Client) ConsensusParams(_ context.Context, height *int64) (*ctypes.ResultConsensusParams, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c Client) Health(_ context.Context) (*ctypes.ResultHealth, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c Client) DialSeeds(_ context.Context, seeds []string) (*ctypes.ResultDialSeeds, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c Client) DialPeers(
	_ context.Context,
	peers []string,
	persistent,
	unconditional,
	private bool,
) (*ctypes.ResultDialPeers, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c Client) BlockchainInfo(_ context.Context, minHeight, maxHeight int64) (*ctypes.ResultBlockchainInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c Client) Genesis(context.Context) (*ctypes.ResultGenesis, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c Client) Block(_ context.Context, height *int64) (*ctypes.ResultBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c Client) BlockByHash(_ context.Context, hash []byte) (*ctypes.ResultBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c Client) Commit(_ context.Context, height *int64) (*ctypes.ResultCommit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c Client) Validators(_ context.Context, height *int64, page, perPage *int) (*ctypes.ResultValidators, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c Client) BroadcastEvidence(_ context.Context, ev types.Evidence) (*ctypes.ResultBroadcastEvidence, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c Client) DataCommitment(
	ctx context.Context,
	start uint64,
	end uint64,
) (*ctypes.ResultDataCommitment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c Client) DataRootInclusionProof(
	ctx context.Context,
	height uint64,
	start uint64,
	end uint64,
) (*ctypes.ResultDataRootInclusionProof, error) {
	_ = "STUB: not implemented"
	//nolint:gosec
	return nil, nil
}
