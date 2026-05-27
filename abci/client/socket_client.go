package abcicli

import (
	"container/list"
	"context"
	"io"
	"net"
	"sync"

	"github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/libs/service"
	"github.com/cometbft/cometbft/libs/timer"
)

const (
	reqQueueSize    = 256 // TODO make configurable
	flushThrottleMS = 20  // Don't wait longer than...
)

// socketClient is the client side implementation of the Tendermint
// Socket Protocol (TSP). It is used by an instance of Tendermint to pass
// ABCI requests to an out of process application running the socketServer.
//
// This is goroutine-safe. All calls are serialized to the server through an unbuffered queue. The socketClient
// tracks responses and expects them to respect the order of the requests sent.
type socketClient struct {
	service.BaseService

	addr        string
	mustConnect bool
	conn        net.Conn

	reqQueue   chan *ReqRes
	flushTimer *timer.ThrottleTimer

	mtx     sync.Mutex
	err     error
	reqSent *list.List                            // list of requests sent, waiting for response
	resCb   func(*types.Request, *types.Response) // called on all requests, if set.
}

var _ Client = (*socketClient)(nil)

// NewSocketClient creates a new socket client, which connects to a given
// address. If mustConnect is true, the client will return an error upon start
// if it fails to connect else it will continue to retry.
func NewSocketClient(addr string, mustConnect bool) Client {
	_ = "STUB: not implemented"
	return *new(Client)
}

// OnStart implements Service by connecting to the server and spawning reading
// and writing goroutines.
func (cli *socketClient) OnStart() error { _ = "STUB: not implemented"; return nil }

// OnStop implements Service by closing connection and flushing all queues.
func (cli *socketClient) OnStop() { _ = "STUB: not implemented"; return }

// Error returns an error if the client was stopped abruptly.
func (cli *socketClient) Error() error { _ = "STUB: not implemented"; return nil }

//----------------------------------------

// SetResponseCallback sets a callback, which will be executed for each
// non-error & non-empty response from the server.
//
// NOTE: callback may get internally generated flush responses.
func (cli *socketClient) SetResponseCallback(resCb Callback) { _ = "STUB: not implemented"; return }

func (cli *socketClient) CheckTxAsync(ctx context.Context, req *types.RequestCheckTx) (*ReqRes, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//----------------------------------------

func (cli *socketClient) sendRequestsRoutine(conn io.Writer) { _ = "STUB: not implemented"; return }

// N.B. We must enqueue before sending out the request, otherwise the
// server may reply before we do it, and the receiver will fail for an
// unsolicited reply.

// If it's a flush request, flush the current buffer.

// flush queue

// Probably will fill the buffer, or retry later.

func (cli *socketClient) recvResponseRoutine(conn io.Reader) { _ = "STUB: not implemented"; return }

// app responded with error
// XXX After setting cli.err, release waiters (e.g. reqres.Done())

func (cli *socketClient) trackRequest(reqres *ReqRes) {
	_ = "STUB: not implemented"
	// N.B. We must NOT hold the client state lock while checking this, or we
	// may deadlock with shutdown.
	return
}

func (cli *socketClient) didRecvResponse(res *types.Response) error {
	_ = "STUB: not implemented"
	return nil
}

// Get the first ReqRes.

// release waiters
// pop first item from linked list

// Notify client listener if set (global callback).

// Notify reqRes listener if set (request specific callback).
//
// NOTE: It is possible this callback isn't set on the reqres object. At this
// point, in which case it will be called after, when it is set.

//----------------------------------------

func (cli *socketClient) Flush(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (cli *socketClient) Echo(ctx context.Context, msg string) (*types.ResponseEcho, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cli *socketClient) Info(ctx context.Context, req *types.RequestInfo) (*types.ResponseInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cli *socketClient) CheckTx(ctx context.Context, req *types.RequestCheckTx) (*types.ResponseCheckTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cli *socketClient) Query(ctx context.Context, req *types.RequestQuery) (*types.ResponseQuery, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cli *socketClient) Commit(ctx context.Context, _ *types.RequestCommit) (*types.ResponseCommit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cli *socketClient) InitChain(ctx context.Context, req *types.RequestInitChain) (*types.ResponseInitChain, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cli *socketClient) ListSnapshots(ctx context.Context, req *types.RequestListSnapshots) (*types.ResponseListSnapshots, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cli *socketClient) OfferSnapshot(ctx context.Context, req *types.RequestOfferSnapshot) (*types.ResponseOfferSnapshot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cli *socketClient) LoadSnapshotChunk(ctx context.Context, req *types.RequestLoadSnapshotChunk) (*types.ResponseLoadSnapshotChunk, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cli *socketClient) ApplySnapshotChunk(ctx context.Context, req *types.RequestApplySnapshotChunk) (*types.ResponseApplySnapshotChunk, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cli *socketClient) PrepareProposal(ctx context.Context, req *types.RequestPrepareProposal) (*types.ResponsePrepareProposal, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cli *socketClient) ProcessProposal(ctx context.Context, req *types.RequestProcessProposal) (*types.ResponseProcessProposal, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cli *socketClient) ExtendVote(ctx context.Context, req *types.RequestExtendVote) (*types.ResponseExtendVote, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cli *socketClient) VerifyVoteExtension(ctx context.Context, req *types.RequestVerifyVoteExtension) (*types.ResponseVerifyVoteExtension, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cli *socketClient) FinalizeBlock(ctx context.Context, req *types.RequestFinalizeBlock) (*types.ResponseFinalizeBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cli *socketClient) QuerySequence(ctx context.Context, req *types.RequestQuerySequence) (*types.ResponseQuerySequence, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cli *socketClient) queueRequest(ctx context.Context, req *types.Request) (*ReqRes, error) {
	_ = "STUB: not implemented"
	return nil,

		// TODO: set cli.err if reqQueue times out
		nil
}

// Maybe auto-flush, or unset auto-flush

// flushQueue marks as complete and discards all remaining pending requests
// from the queue.
func (cli *socketClient) flushQueue() { _ = "STUB: not implemented"; return }

// mark all in-flight messages as resolved (they will get cli.Error())

// mark all queued messages as resolved

//----------------------------------------

func resMatchesReq(req *types.Request, res *types.Response) (ok bool) {
	_ = "STUB: not implemented"
	return false
}

func (cli *socketClient) stopForError(err error) { _ = "STUB: not implemented"; return }
