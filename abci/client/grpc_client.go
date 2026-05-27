package abcicli

import (
	"context"
	"net"
	"sync"

	"google.golang.org/grpc"

	"github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/libs/service"
)

var _ Client = (*grpcClient)(nil)

// A stripped copy of the remoteClient that makes
// synchronous calls using grpc
type grpcClient struct {
	service.BaseService
	mustConnect bool

	client   types.ABCIClient
	conn     *grpc.ClientConn
	chReqRes chan *ReqRes // dispatches "async" responses to callbacks *in order*, needed by mempool

	mtx   sync.Mutex
	addr  string
	err   error
	resCb func(*types.Request, *types.Response) // listens to all callbacks
}

func NewGRPCClient(addr string, mustConnect bool) Client {
	_ = "STUB: not implemented"
	return *new(Client)
}

// Buffering the channel is needed to make calls appear asynchronous,
// which is required when the caller makes multiple async calls before
// processing callbacks (e.g. due to holding locks). 64 means that a
// caller can make up to 64 async calls before a callback must be
// processed (otherwise it deadlocks). It also means that we can make 64
// gRPC calls while processing a slow callback at the channel head.

func dialerFunc(_ context.Context, addr string) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

func (cli *grpcClient) OnStart() error { _ = "STUB: not implemented"; return nil }

// This processes asynchronous request/response messages and dispatches
// them to callbacks.

// Use a separate function to use defer for mutex unlocks (this handles panics)

// Notify client listener if set

// Notify reqRes listener if set

func (cli *grpcClient) OnStop() { _ = "STUB: not implemented"; return }

func (cli *grpcClient) StopForError(err error) { _ = "STUB: not implemented"; return }

func (cli *grpcClient) Error() error { _ = "STUB: not implemented"; return nil }

// Set listener for all responses
// NOTE: callback may get internally generated flush responses.
func (cli *grpcClient) SetResponseCallback(resCb Callback) { _ = "STUB: not implemented"; return }

//----------------------------------------

func (cli *grpcClient) CheckTxAsync(ctx context.Context, req *types.RequestCheckTx) (*ReqRes, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// finishAsyncCall creates a ReqRes for an async call, and immediately populates it
// with the response. We don't complete it until it's been ordered via the channel.
func (cli *grpcClient) finishAsyncCall(req *types.Request, res *types.Response) *ReqRes {
	_ = "STUB: not implemented"
	return nil
}

// use channel for async responses, since they must be ordered

//----------------------------------------

func (cli *grpcClient) Flush(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (cli *grpcClient) Echo(ctx context.Context, msg string) (*types.ResponseEcho, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cli *grpcClient) Info(ctx context.Context, req *types.RequestInfo) (*types.ResponseInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cli *grpcClient) CheckTx(ctx context.Context, req *types.RequestCheckTx) (*types.ResponseCheckTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cli *grpcClient) Query(ctx context.Context, req *types.RequestQuery) (*types.ResponseQuery, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cli *grpcClient) Commit(ctx context.Context, _ *types.RequestCommit) (*types.ResponseCommit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cli *grpcClient) InitChain(ctx context.Context, req *types.RequestInitChain) (*types.ResponseInitChain, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cli *grpcClient) ListSnapshots(ctx context.Context, req *types.RequestListSnapshots) (*types.ResponseListSnapshots, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cli *grpcClient) OfferSnapshot(ctx context.Context, req *types.RequestOfferSnapshot) (*types.ResponseOfferSnapshot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cli *grpcClient) LoadSnapshotChunk(ctx context.Context, req *types.RequestLoadSnapshotChunk) (*types.ResponseLoadSnapshotChunk, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cli *grpcClient) ApplySnapshotChunk(ctx context.Context, req *types.RequestApplySnapshotChunk) (*types.ResponseApplySnapshotChunk, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cli *grpcClient) PrepareProposal(ctx context.Context, req *types.RequestPrepareProposal) (*types.ResponsePrepareProposal, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cli *grpcClient) ProcessProposal(ctx context.Context, req *types.RequestProcessProposal) (*types.ResponseProcessProposal, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cli *grpcClient) ExtendVote(ctx context.Context, req *types.RequestExtendVote) (*types.ResponseExtendVote, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cli *grpcClient) VerifyVoteExtension(ctx context.Context, req *types.RequestVerifyVoteExtension) (*types.ResponseVerifyVoteExtension, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cli *grpcClient) FinalizeBlock(ctx context.Context, req *types.RequestFinalizeBlock) (*types.ResponseFinalizeBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cli *grpcClient) QuerySequence(ctx context.Context, req *types.RequestQuerySequence) (*types.ResponseQuerySequence, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
