package state

import (
	"context"

	"github.com/cometbft/cometbft/libs/trace"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cometbft/cometbft/mempool"
	"github.com/cometbft/cometbft/proxy"
	"github.com/cometbft/cometbft/types"
)

//-----------------------------------------------------------------------------
// BlockExecutor handles block execution and state updates.
// It exposes ApplyBlock(), which validates & executes the block, updates state w/ ABCI responses,
// then commits and updates the mempool atomically, then saves state.

// BlockExecutor provides the context and accessories for properly executing a block.
type BlockExecutor struct {
	// save state, validators, consensus params, abci responses here
	store Store

	// use blockstore for the pruning functions.
	blockStore BlockStore

	// execute the app against this
	proxyApp proxy.AppConnConsensus

	// events
	eventBus types.BlockEventPublisher

	// manage the mempool lock during commit
	// and update both with block results after commit.
	mempool mempool.Mempool
	evpool  EvidencePool

	logger log.Logger

	metrics *Metrics

	// root directory for debug file saving
	rootDir string

	// tracer optional tracer
	tracer trace.Tracer

	// running count of state entries pruned, used to decide when to trigger
	// state-store compaction via PruneStates.
	prunedStates uint64
}

type BlockExecutorOption func(executor *BlockExecutor)

func BlockExecutorWithMetrics(metrics *Metrics) BlockExecutorOption {
	_ = "STUB: not implemented"
	return *new(BlockExecutorOption)
}

func BlockExecutorWithRootDir(rootDir string) BlockExecutorOption {
	_ = "STUB: not implemented"
	return *new(BlockExecutorOption)
}

func BlockExecutorWithTracer(tracer trace.Tracer) BlockExecutorOption {
	_ = "STUB: not implemented"
	return *new(BlockExecutorOption)
}

// NewBlockExecutor returns a new BlockExecutor with a NopEventBus.
// Call SetEventBus to provide one.
func NewBlockExecutor(
	stateStore Store,
	logger log.Logger,
	proxyApp proxy.AppConnConsensus,
	mempool mempool.Mempool,
	evpool EvidencePool,
	blockStore BlockStore,
	options ...BlockExecutorOption,
) *BlockExecutor {
	_ = "STUB: not implemented"
	return nil
}

func (blockExec *BlockExecutor) Store() Store {
	_ = "STUB: not implemented"
	return *

	// SetEventBus - sets the event bus for publishing block related events.
	// If not called, it defaults to types.NopEventBus.
	new(Store)
}

func (blockExec *BlockExecutor) SetEventBus(eventBus types.BlockEventPublisher) {
	_ = "STUB: not implemented"
	return
}

// CreateProposalBlock calls state.MakeBlock with evidence from the evpool
// and txs from the mempool. The max bytes must be big enough to fit the commit.
// The block space is first allocated to outstanding evidence.
// The rest is given to txs, up to the max gas.
//
// Contract: application will not return more bytes than are sent over the wire.
func (blockExec *BlockExecutor) CreateProposalBlock(
	ctx context.Context,
	height int64,
	state State,
	lastExtCommit *types.ExtendedCommit,
	proposerAddr []byte,
) (*types.Block, *types.PartSet, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Fetch a limited amount of valid txs

// Reap 25% more transactions than can fit in the block to provide
// buffer for PrepareProposal to have sufficient transactions to work with

// For non-panic errors, also save the failed proposal block

// The App MUST ensure that only valid (and hence 'processable') transactions
// enter the mempool. Hence, at this point, we can't have any non-processable
// transaction causing an error.
//
// Also, the App can simply skip any transaction that could cause any kind of trouble.
// Either way, we cannot recover in a meaningful way, unless we skip proposing
// this block, repair what caused the error and try again. Hence, we return an
// error for now (the production code calling this function is expected to panic).

// get the cached hashes
// TODO: make sure that the hashes are correct here
// via also removing hashes that the application removed!

func (blockExec *BlockExecutor) ProcessProposal(
	block *types.Block,
	initialHeight int64,
) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

//nolint:staticcheck
//nolint:staticcheck
//nolint:staticcheck
//nolint:staticcheck

// needed for v3 to sync with multiplexer as the header is stored in state

// ValidateBlock validates the given block against the given state.
// If the block is invalid, it returns an error.
// Validation does not mutate state, but does require historical information from the stateDB,
// ie. to verify evidence from a validator at an old height.
func (blockExec *BlockExecutor) ValidateBlock(state State, block *types.Block) error {
	_ = "STUB: not implemented"
	return nil
}

// ApplyVerifiedBlock does the same as `ApplyBlock`, but skips verification.
func (blockExec *BlockExecutor) ApplyVerifiedBlock(
	state State, blockID types.BlockID, block *types.Block, lastCommit *types.Commit,
) (State, error) {
	_ = "STUB: not implemented"
	return *new(State), nil
}

// ApplyBlock validates the block against the state, executes it against the app,
// fires the relevant events, commits the app, and saves the new state and responses.
// It returns the new state.
// It's the only function that needs to be called
// from outside this package to process and commit an entire block.
// It takes a blockID to avoid recomputing the parts hash.
func (blockExec *BlockExecutor) ApplyBlock(
	state State, blockID types.BlockID, block *types.Block, lastCommit *types.Commit,
) (State, error) {
	_ = "STUB: not implemented"
	return *new(State), nil
}

func (blockExec *BlockExecutor) applyBlock(state State, blockID types.BlockID, block *types.Block, lastCommit *types.Commit) (State, error) {
	_ = "STUB: not implemented"
	return *new(State), nil
}

// Unmarshal blob txs

// needed for v3 to sync with multiplexer as the header is stored in state

// Assert that the application correctly returned tx results for each of the transactions provided in the block
//nolint:staticcheck
//nolint:staticcheck

// Save indexing info of the transaction.
// This needs to be done prior to saving state
// for correct crash recovery

// XXX

// Save the results before we commit.

// XXX

// validate the validator updates and convert to CometBFT types

// Update the state with the block and responses.

// Lock mempool, commit app state, update mempoool.

// Update evpool with the latest state.

// XXX

// Update the app hash and save the state.

// XXX

// Prune old heights, if requested by ABCI app.

// Events are fired after everything else.
// NOTE: if we crash between Commit and Save, events wont be fired during replay

func (blockExec *BlockExecutor) ExtendVote(
	ctx context.Context,
	vote *types.Vote,
	block *types.Block,
	state State,
) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (blockExec *BlockExecutor) VerifyVoteExtension(ctx context.Context, vote *types.Vote) error {
	_ = "STUB: not implemented"
	return nil
}

// Commit locks the mempool, runs the ABCI Commit message, and updates the
// mempool.
// It returns the result of calling abci.Commit which is the height to retain (if any)).
// The application is expected to have persisted its state (if any) before returning
// from the ABCI Commit call. This is the only place where the application should
// persist its state.
// The Mempool must be locked during commit and update because state is
// typically reset on Commit and old txs must be replayed against committed
// state before new txs are run in the mempool, lest they be invalid.
func (blockExec *BlockExecutor) Commit(
	state State,
	block *types.Block,
	abciResponse *abci.ResponseFinalizeBlock,
) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// while mempool is Locked, flush to ensure all async requests have completed
// in the ABCI app before Commit.

// Commit block, get hash back

// ResponseCommit has no error code - just data

// Update mempool.

//---------------------------------------------------------
// Helper functions for executing blocks and updating state

func buildLastCommitInfoFromStore(block *types.Block, store Store, initialHeight int64) abci.CommitInfo {
	_ = "STUB: not implemented"
	return *new(abci.CommitInfo)
}

// check for initial height before loading validators
// there is no last commit for the initial height.
// return an empty value.

// BuildLastCommitInfo builds a CommitInfo from the given block and validator set.
// If you want to load the validator set from the store instead of providing it,
// use buildLastCommitInfoFromStore.
func BuildLastCommitInfo(block *types.Block, lastValSet *types.ValidatorSet, initialHeight int64) abci.CommitInfo {
	_ = "STUB: not implemented"
	return *new(abci.CommitInfo)
}

// there is no last commit for the initial height.
// return an empty value.

// ensure that the size of the validator set in the last commit matches
// the size of the validator set in the state store.

// buildExtendedCommitInfoFromStore populates an ABCI extended commit from the
// corresponding CometBFT extended commit ec, using the stored validator set
// from ec.  It requires ec to include the original precommit votes along with
// the vote extensions from the last commit.
//
// For heights below the initial height, for which we do not have the required
// data, it returns an empty record.
//
// Assumes that the commit signatures are sorted according to validator index.
func buildExtendedCommitInfoFromStore(ec *types.ExtendedCommit, store Store, initialHeight int64, ap types.ABCIParams) abci.ExtendedCommitInfo {
	_ = "STUB: not implemented"
	return *new(abci.ExtendedCommitInfo)
}

// There are no extended commits for heights below the initial height.

// BuildExtendedCommitInfo builds an ExtendedCommitInfo from the given block and validator set.
// If you want to load the validator set from the store instead of providing it,
// use buildExtendedCommitInfoFromStore.
func BuildExtendedCommitInfo(ec *types.ExtendedCommit, valSet *types.ValidatorSet, initialHeight int64, ap types.ABCIParams) abci.ExtendedCommitInfo {
	_ = "STUB: not implemented"
	return *new(abci.ExtendedCommitInfo)
}

// There are no extended commits for heights below the initial height.

// Ensure that the size of the validator set in the extended commit matches
// the size of the validator set in the state store.

// Absent signatures have empty validator addresses, but otherwise we
// expect the validator addresses to be the same.

// Check if vote extensions were enabled during the commit's height: ec.Height.
// ec is the commit from the previous height, so if extensions were enabled
// during that height, we ensure they are present and deliver the data to
// the proposer. If they were not enabled during this previous height, we
// will not deliver extension data.

func validateValidatorUpdates(abciUpdates []abci.ValidatorUpdate,
	params types.ValidatorParams) error {
	_ = "STUB: not implemented"
	return nil
}

// continue, since this is deleting the validator, and thus there is no
// pubkey to check

// Check if validator's pubkey matches an ABCI type in the consensus params

// updateState returns a new State updated according to the header and responses.
func updateState(
	state State,
	blockID types.BlockID,
	header *types.Header,
	abciResponse *abci.ResponseFinalizeBlock,
	validatorUpdates []*types.Validator,
) (State, error) {
	_ = "STUB: not implemented"

	// Copy the valset so we can apply changes from EndBlock
	// and update s.LastValidators and s.Validators.
	return *new(State), nil
}

// Update the validator set with the latest abciResponse.

// Change results from this height but only applies to the next next height.

// Update validator proposer priority and set state variables.

// Update the params with the latest abciResponse.

// NOTE: must not mutate state.ConsensusParams

// Change results from this height but only applies to the next height.

// NOTE: the AppHash and the VoteExtension has not been populated.
// It will be filled on state.Save.

// Fire NewBlock, NewBlockHeader.
// Fire TxEvent for every tx.
// NOTE: if CometBFT crashes before commit, some or all of these events may be published again.
func fireEvents(
	logger log.Logger,
	eventBus types.BlockEventPublisher,
	block *types.Block,
	blockID types.BlockID,
	abciResponse *abci.ResponseFinalizeBlock,
	validatorUpdates []*types.Validator,
	currentValidators *types.ValidatorSet,
	lastCommit *types.Commit,
) {
	_ = "STUB: not implemented"
	return
}

//nolint:staticcheck

//----------------------------------------------------------------------------------------------------
// Execute block without state. TODO: eliminate

// ExecCommitBlock executes and commits a block on the proxyApp without validating or mutating the state.
// It returns the application root hash (result of abci.Commit).
func ExecCommitBlock(
	appConnConsensus proxy.AppConnConsensus,
	block *types.Block,
	logger log.Logger,
	store Store,
	initialHeight int64,
) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Unmarshal blob txs to match the behavior of applyBlock, which strips
// BlobTx wrappers before sending transactions to FinalizeBlock.

// needed for v3 to sync with multiplexer as the header is stored in state

// Assert that the application correctly returned tx results for each of the transactions provided in the block
//nolint:staticcheck
//nolint:staticcheck

// Commit block

// ResponseCommit has no error or log

func (blockExec *BlockExecutor) pruneBlocks(retainHeight int64, state State) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// PruneStates flushes in 1000-entry batches, so partial deletions persist
// to disk even on error. Account for them so the compaction-trigger
// bucket math stays aligned with what's actually on disk.

// saveFailedProposalBlock saves a failed proposal block to the debug directory
func (blockExec *BlockExecutor) saveFailedProposalBlock(state State, block *types.Block, reason string) {
	_ = "STUB: not implemented"
	return
}
