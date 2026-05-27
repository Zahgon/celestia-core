package abcicli

import (
	"context"

	types "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/libs/service"
	cmtsync "github.com/cometbft/cometbft/libs/sync"
)

// NOTE: use defer to unlock mutex because Application might panic (e.g., in
// case of malicious tx or query). It only makes sense for publicly exposed
// methods like CheckTx (/broadcast_tx_* RPC endpoint) or Query (/abci_query
// RPC endpoint), but defers are used everywhere for the sake of consistency.
type localClient struct {
	service.BaseService

	mtx *cmtsync.Mutex
	types.Application
	Callback
}

var _ Client = (*localClient)(nil)

// NewLocalClient creates a local client, which wraps the application interface that
// Tendermint as the client will call to the application as the server. The only
// difference, is that the local client has a global mutex which enforces serialization
// of all the ABCI calls from Tendermint to the Application.
func NewLocalClient(mtx *cmtsync.Mutex, app types.Application) Client {
	_ = "STUB: not implemented"
	return *new(Client)
}

func (app *localClient) SetResponseCallback(cb Callback) { _ = "STUB: not implemented"; return }

func (app *localClient) CheckTxAsync(ctx context.Context, req *types.RequestCheckTx) (*ReqRes, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (app *localClient) callback(req *types.Request, res *types.Response) *ReqRes {
	_ = "STUB: not implemented"
	return nil
}

func newLocalReqRes(req *types.Request, res *types.Response) *ReqRes {
	_ = "STUB: not implemented"
	return nil
}

//-------------------------------------------------------

func (app *localClient) Error() error { _ = "STUB: not implemented"; return nil }

func (app *localClient) Flush(context.Context) error { _ = "STUB: not implemented"; return nil }

func (app *localClient) Echo(_ context.Context, msg string) (*types.ResponseEcho, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (app *localClient) Info(ctx context.Context, req *types.RequestInfo) (*types.ResponseInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (app *localClient) CheckTx(ctx context.Context, req *types.RequestCheckTx) (*types.ResponseCheckTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (app *localClient) Query(ctx context.Context, req *types.RequestQuery) (*types.ResponseQuery, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (app *localClient) Commit(ctx context.Context, req *types.RequestCommit) (*types.ResponseCommit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (app *localClient) InitChain(ctx context.Context, req *types.RequestInitChain) (*types.ResponseInitChain, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (app *localClient) ListSnapshots(ctx context.Context, req *types.RequestListSnapshots) (*types.ResponseListSnapshots, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (app *localClient) OfferSnapshot(ctx context.Context, req *types.RequestOfferSnapshot) (*types.ResponseOfferSnapshot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (app *localClient) LoadSnapshotChunk(ctx context.Context,
	req *types.RequestLoadSnapshotChunk) (*types.ResponseLoadSnapshotChunk, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (app *localClient) ApplySnapshotChunk(ctx context.Context,
	req *types.RequestApplySnapshotChunk) (*types.ResponseApplySnapshotChunk, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (app *localClient) PrepareProposal(ctx context.Context, req *types.RequestPrepareProposal) (*types.ResponsePrepareProposal, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (app *localClient) ProcessProposal(ctx context.Context, req *types.RequestProcessProposal) (*types.ResponseProcessProposal, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (app *localClient) ExtendVote(ctx context.Context, req *types.RequestExtendVote) (*types.ResponseExtendVote, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (app *localClient) VerifyVoteExtension(ctx context.Context, req *types.RequestVerifyVoteExtension) (*types.ResponseVerifyVoteExtension, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (app *localClient) FinalizeBlock(ctx context.Context, req *types.RequestFinalizeBlock) (*types.ResponseFinalizeBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (app *localClient) QuerySequence(ctx context.Context, req *types.RequestQuerySequence) (*types.ResponseQuerySequence, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
