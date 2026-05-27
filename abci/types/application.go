package types

import "context"

//go:generate ../../scripts/mockery_generate.sh Application

// Application is an interface that enables any finite, deterministic state machine
// to be driven by a blockchain-based replication engine via the ABCI.
type Application interface {
	// Info/Query Connection
	Info(context.Context, *RequestInfo) (*ResponseInfo, error)    // Return application info
	Query(context.Context, *RequestQuery) (*ResponseQuery, error) // Query for state
	QuerySequence(context.Context, *RequestQuerySequence) (*ResponseQuerySequence, error)

	// Mempool Connection
	CheckTx(context.Context, *RequestCheckTx) (*ResponseCheckTx, error) // Validate a tx for the mempool

	// Consensus Connection
	InitChain(context.Context, *RequestInitChain) (*ResponseInitChain, error) // Initialize blockchain w validators/other info from CometBFT
	PrepareProposal(context.Context, *RequestPrepareProposal) (*ResponsePrepareProposal, error)
	ProcessProposal(context.Context, *RequestProcessProposal) (*ResponseProcessProposal, error)
	// Deliver the decided block with its txs to the Application
	FinalizeBlock(context.Context, *RequestFinalizeBlock) (*ResponseFinalizeBlock, error)
	// Create application specific vote extension
	ExtendVote(context.Context, *RequestExtendVote) (*ResponseExtendVote, error)
	// Verify application's vote extension data
	VerifyVoteExtension(context.Context, *RequestVerifyVoteExtension) (*ResponseVerifyVoteExtension, error)
	// Commit the state and return the application Merkle root hash
	Commit(context.Context, *RequestCommit) (*ResponseCommit, error)

	// State Sync Connection
	ListSnapshots(context.Context, *RequestListSnapshots) (*ResponseListSnapshots, error)                // List available snapshots
	OfferSnapshot(context.Context, *RequestOfferSnapshot) (*ResponseOfferSnapshot, error)                // Offer a snapshot to the application
	LoadSnapshotChunk(context.Context, *RequestLoadSnapshotChunk) (*ResponseLoadSnapshotChunk, error)    // Load a snapshot chunk
	ApplySnapshotChunk(context.Context, *RequestApplySnapshotChunk) (*ResponseApplySnapshotChunk, error) // Apply a shapshot chunk
}

//-------------------------------------------------------
// BaseApplication is a base form of Application

var _ Application = (*BaseApplication)(nil)

type BaseApplication struct{}

func NewBaseApplication() *BaseApplication { _ = "STUB: not implemented"; return nil }

func (BaseApplication) Info(context.Context, *RequestInfo) (*ResponseInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (BaseApplication) CheckTx(context.Context, *RequestCheckTx) (*ResponseCheckTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (BaseApplication) Commit(context.Context, *RequestCommit) (*ResponseCommit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (BaseApplication) Query(context.Context, *RequestQuery) (*ResponseQuery, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (BaseApplication) QuerySequence(context.Context, *RequestQuerySequence) (*ResponseQuerySequence, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (BaseApplication) InitChain(context.Context, *RequestInitChain) (*ResponseInitChain, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (BaseApplication) ListSnapshots(context.Context, *RequestListSnapshots) (*ResponseListSnapshots, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (BaseApplication) OfferSnapshot(context.Context, *RequestOfferSnapshot) (*ResponseOfferSnapshot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (BaseApplication) LoadSnapshotChunk(context.Context, *RequestLoadSnapshotChunk) (*ResponseLoadSnapshotChunk, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (BaseApplication) ApplySnapshotChunk(context.Context, *RequestApplySnapshotChunk) (*ResponseApplySnapshotChunk, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (BaseApplication) PrepareProposal(_ context.Context, req *RequestPrepareProposal) (*ResponsePrepareProposal, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (BaseApplication) ProcessProposal(context.Context, *RequestProcessProposal) (*ResponseProcessProposal, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (BaseApplication) ExtendVote(context.Context, *RequestExtendVote) (*ResponseExtendVote, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (BaseApplication) VerifyVoteExtension(context.Context, *RequestVerifyVoteExtension) (*ResponseVerifyVoteExtension, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (BaseApplication) FinalizeBlock(_ context.Context, req *RequestFinalizeBlock) (*ResponseFinalizeBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
