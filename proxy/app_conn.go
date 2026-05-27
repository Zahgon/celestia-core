package proxy

import (
	"context"

	"github.com/go-kit/kit/metrics"

	abcicli "github.com/cometbft/cometbft/abci/client"
	"github.com/cometbft/cometbft/abci/types"
)

//go:generate ../scripts/mockery_generate.sh AppConnConsensus|AppConnMempool|AppConnQuery|AppConnSnapshot

//----------------------------------------------------------------------------------------
// Enforce which abci msgs can be sent on a connection at the type level

type AppConnConsensus interface {
	Error() error
	InitChain(context.Context, *types.RequestInitChain) (*types.ResponseInitChain, error)
	PrepareProposal(context.Context, *types.RequestPrepareProposal) (*types.ResponsePrepareProposal, error)
	ProcessProposal(context.Context, *types.RequestProcessProposal) (*types.ResponseProcessProposal, error)
	ExtendVote(context.Context, *types.RequestExtendVote) (*types.ResponseExtendVote, error)
	VerifyVoteExtension(context.Context, *types.RequestVerifyVoteExtension) (*types.ResponseVerifyVoteExtension, error)
	FinalizeBlock(context.Context, *types.RequestFinalizeBlock) (*types.ResponseFinalizeBlock, error)
	Commit(context.Context) (*types.ResponseCommit, error)
}

type AppConnMempool interface {
	SetResponseCallback(abcicli.Callback)
	Error() error

	CheckTx(context.Context, *types.RequestCheckTx) (*types.ResponseCheckTx, error)
	CheckTxAsync(context.Context, *types.RequestCheckTx) (*abcicli.ReqRes, error)
	Flush(context.Context) error
	QuerySequence(context.Context, *types.RequestQuerySequence) (*types.ResponseQuerySequence, error)
}

type AppConnQuery interface {
	Error() error

	Echo(context.Context, string) (*types.ResponseEcho, error)
	Info(context.Context, *types.RequestInfo) (*types.ResponseInfo, error)
	Query(context.Context, *types.RequestQuery) (*types.ResponseQuery, error)
}

type AppConnSnapshot interface {
	Error() error

	ListSnapshots(context.Context, *types.RequestListSnapshots) (*types.ResponseListSnapshots, error)
	OfferSnapshot(context.Context, *types.RequestOfferSnapshot) (*types.ResponseOfferSnapshot, error)
	LoadSnapshotChunk(context.Context, *types.RequestLoadSnapshotChunk) (*types.ResponseLoadSnapshotChunk, error)
	ApplySnapshotChunk(context.Context, *types.RequestApplySnapshotChunk) (*types.ResponseApplySnapshotChunk, error)
}

//-----------------------------------------------------------------------------------------
// Implements AppConnConsensus (subset of abcicli.Client)

type appConnConsensus struct {
	metrics *Metrics
	appConn abcicli.Client
}

var _ AppConnConsensus = (*appConnConsensus)(nil)

func NewAppConnConsensus(appConn abcicli.Client, metrics *Metrics) AppConnConsensus {
	_ = "STUB: not implemented"
	return *new(AppConnConsensus)
}

func (app *appConnConsensus) Error() error { _ = "STUB: not implemented"; return nil }

func (app *appConnConsensus) InitChain(ctx context.Context, req *types.RequestInitChain) (*types.ResponseInitChain, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (app *appConnConsensus) PrepareProposal(ctx context.Context,
	req *types.RequestPrepareProposal) (*types.ResponsePrepareProposal, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (app *appConnConsensus) ProcessProposal(ctx context.Context, req *types.RequestProcessProposal) (*types.ResponseProcessProposal, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (app *appConnConsensus) ExtendVote(ctx context.Context, req *types.RequestExtendVote) (*types.ResponseExtendVote, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (app *appConnConsensus) VerifyVoteExtension(ctx context.Context, req *types.RequestVerifyVoteExtension) (*types.ResponseVerifyVoteExtension, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (app *appConnConsensus) FinalizeBlock(ctx context.Context, req *types.RequestFinalizeBlock) (*types.ResponseFinalizeBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (app *appConnConsensus) Commit(ctx context.Context) (*types.ResponseCommit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//------------------------------------------------
// Implements AppConnMempool (subset of abcicli.Client)

type appConnMempool struct {
	metrics *Metrics
	appConn abcicli.Client
}

func NewAppConnMempool(appConn abcicli.Client, metrics *Metrics) AppConnMempool {
	_ = "STUB: not implemented"
	return *new(AppConnMempool)
}

func (app *appConnMempool) SetResponseCallback(cb abcicli.Callback) {
	_ = "STUB: not implemented"
	return
}

func (app *appConnMempool) Error() error { _ = "STUB: not implemented"; return nil }

func (app *appConnMempool) Flush(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (app *appConnMempool) CheckTx(ctx context.Context, req *types.RequestCheckTx) (*types.ResponseCheckTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (app *appConnMempool) CheckTxAsync(ctx context.Context, req *types.RequestCheckTx) (*abcicli.ReqRes, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (app *appConnMempool) QuerySequence(ctx context.Context, req *types.RequestQuerySequence) (*types.ResponseQuerySequence, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//------------------------------------------------
// Implements AppConnQuery (subset of abcicli.Client)

type appConnQuery struct {
	metrics *Metrics
	appConn abcicli.Client
}

func NewAppConnQuery(appConn abcicli.Client, metrics *Metrics) AppConnQuery {
	_ = "STUB: not implemented"
	return *new(AppConnQuery)
}

func (app *appConnQuery) Error() error { _ = "STUB: not implemented"; return nil }

func (app *appConnQuery) Echo(ctx context.Context, msg string) (*types.ResponseEcho, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (app *appConnQuery) Info(ctx context.Context, req *types.RequestInfo) (*types.ResponseInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (app *appConnQuery) Query(ctx context.Context, req *types.RequestQuery) (*types.ResponseQuery, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//------------------------------------------------
// Implements AppConnSnapshot (subset of abcicli.Client)

type appConnSnapshot struct {
	metrics *Metrics
	appConn abcicli.Client
}

func NewAppConnSnapshot(appConn abcicli.Client, metrics *Metrics) AppConnSnapshot {
	_ = "STUB: not implemented"
	return *new(AppConnSnapshot)
}

func (app *appConnSnapshot) Error() error { _ = "STUB: not implemented"; return nil }

func (app *appConnSnapshot) ListSnapshots(ctx context.Context, req *types.RequestListSnapshots) (*types.ResponseListSnapshots, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (app *appConnSnapshot) OfferSnapshot(ctx context.Context, req *types.RequestOfferSnapshot) (*types.ResponseOfferSnapshot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (app *appConnSnapshot) LoadSnapshotChunk(ctx context.Context, req *types.RequestLoadSnapshotChunk) (*types.ResponseLoadSnapshotChunk, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (app *appConnSnapshot) ApplySnapshotChunk(ctx context.Context, req *types.RequestApplySnapshotChunk) (*types.ResponseApplySnapshotChunk, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// addTimeSample returns a function that, when called, adds an observation to m.
// The observation added to m is the number of seconds ellapsed since addTimeSample
// was initially called. addTimeSample is meant to be called in a defer to calculate
// the amount of time a function takes to complete.
func addTimeSample(m metrics.Histogram) func() { _ = "STUB: not implemented"; return nil }
