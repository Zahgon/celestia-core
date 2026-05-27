package app

import (
	"context"
	"time"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/libs/log"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
)

const (
	appVersion                 = 1
	voteExtensionKey    string = "extensionSum"
	voteExtensionMaxVal int64  = 128
	prefixReservedKey   string = "reservedTxKey_"
	suffixChainID       string = "ChainID"
	suffixVoteExtHeight string = "VoteExtensionsHeight"
	suffixInitialHeight string = "InitialHeight"
)

// Application is an ABCI application for use by end-to-end tests. It is a
// simple key/value store for strings, storing data in memory and persisting
// to disk as JSON, taking state sync snapshots if requested.
type Application struct {
	abci.BaseApplication
	logger          log.Logger
	state           *State
	snapshots       *SnapshotStore
	cfg             *Config
	restoreSnapshot *abci.Snapshot
	restoreChunks   [][]byte
}

// Config allows for the setting of high level parameters for running the e2e Application
// KeyType and ValidatorUpdates must be the same for all nodes running the same application.
type Config struct {
	// The directory with which state.json will be persisted in. Usually $HOME/.cometbft/data
	Dir string `toml:"dir"`

	// SnapshotInterval specifies the height interval at which the application
	// will take state sync snapshots. Defaults to 0 (disabled).
	SnapshotInterval uint64 `toml:"snapshot_interval"`

	// RetainBlocks specifies the number of recent blocks to retain. Defaults to
	// 0, which retains all blocks. Must be greater that PersistInterval,
	// SnapshotInterval and EvidenceAgeHeight.
	RetainBlocks uint64 `toml:"retain_blocks"`

	// KeyType sets the curve that will be used by validators.
	// Options are ed25519 & secp256k1
	KeyType string `toml:"key_type"`

	// PersistInterval specifies the height interval at which the application
	// will persist state to disk. Defaults to 1 (every height), setting this to
	// 0 disables state persistence.
	PersistInterval uint64 `toml:"persist_interval"`

	// ValidatorUpdates is a map of heights to validator names and their power,
	// and will be returned by the ABCI application. For example, the following
	// changes the power of validator01 and validator02 at height 1000:
	//
	// [validator_update.1000]
	// validator01 = 20
	// validator02 = 10
	//
	// Specifying height 0 returns the validator update during InitChain. The
	// application returns the validator updates as-is, i.e. removing a
	// validator must be done by returning it with power 0, and any validators
	// not specified are not changed.
	//
	// height <-> pubkey <-> voting power
	ValidatorUpdates map[string]map[string]uint8 `toml:"validator_update"`

	// Add artificial delays to each of the main ABCI calls to mimic computation time
	// of the application
	PrepareProposalDelay time.Duration `toml:"prepare_proposal_delay"`
	ProcessProposalDelay time.Duration `toml:"process_proposal_delay"`
	CheckTxDelay         time.Duration `toml:"check_tx_delay"`
	FinalizeBlockDelay   time.Duration `toml:"finalize_block_delay"`
	VoteExtensionDelay   time.Duration `toml:"vote_extension_delay"`

	// VoteExtensionsEnableHeight configures the first height during which
	// the chain will use and require vote extension data to be present
	// in precommit messages.
	VoteExtensionsEnableHeight int64 `toml:"vote_extensions_enable_height"`

	// VoteExtensionsUpdateHeight configures the height at which consensus
	// param VoteExtensionsEnableHeight will be set.
	// -1 denotes it is set at genesis.
	// 0 denotes it is set at InitChain.
	VoteExtensionsUpdateHeight int64 `toml:"vote_extensions_update_height"`
}

func DefaultConfig(dir string) *Config { _ = "STUB: not implemented"; return nil }

// timeoutInfo returns timeout configuration for the e2e application
// Uses the same timeouts as the default consensus config for production-like testing
func (app *Application) timeoutInfo() abci.TimeoutInfo {
	_ = "STUB: not implemented"
	return *new(abci.TimeoutInfo)
}

// e2e uses this specific value

// NewApplication creates the application.
func NewApplication(cfg *Config) (*Application, error) { _ = "STUB: not implemented"; return nil, nil }

// Info implements ABCI.
func (app *Application) Info(context.Context, *abci.RequestInfo) (*abci.ResponseInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (app *Application) updateVoteExtensionEnableHeight(currentHeight int64) *cmtproto.ConsensusParams {
	_ = "STUB: not implemented"
	return nil
}

// Info implements ABCI.
func (app *Application) InitChain(_ context.Context, req *abci.RequestInitChain) (*abci.ResponseInitChain, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get validators from genesis

// CheckTx implements ABCI.
func (app *Application) CheckTx(_ context.Context, req *abci.RequestCheckTx) (*abci.ResponseCheckTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FinalizeBlock implements ABCI.
func (app *Application) FinalizeBlock(_ context.Context, req *abci.RequestFinalizeBlock) (*abci.ResponseFinalizeBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// shouldn't happen since we verified it in CheckTx and ProcessProposal

// Commit implements ABCI.
func (app *Application) Commit(_ context.Context, _ *abci.RequestCommit) (*abci.ResponseCommit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Query implements ABCI.
func (app *Application) Query(_ context.Context, req *abci.RequestQuery) (*abci.ResponseQuery, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ListSnapshots implements ABCI.
func (app *Application) ListSnapshots(context.Context, *abci.RequestListSnapshots) (*abci.ResponseListSnapshots, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// LoadSnapshotChunk implements ABCI.
func (app *Application) LoadSnapshotChunk(_ context.Context, req *abci.RequestLoadSnapshotChunk) (*abci.ResponseLoadSnapshotChunk, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// OfferSnapshot implements ABCI.
func (app *Application) OfferSnapshot(_ context.Context, req *abci.RequestOfferSnapshot) (*abci.ResponseOfferSnapshot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ApplySnapshotChunk implements ABCI.
func (app *Application) ApplySnapshotChunk(_ context.Context, req *abci.RequestApplySnapshotChunk) (*abci.ResponseApplySnapshotChunk, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:prealloc

// PrepareProposal will take the given transactions and attempt to prepare a
// proposal from them when it's our turn to do so. If the current height has
// vote extension enabled, this method will use vote extensions from the previous
// height, passed from CometBFT as parameters to construct a special transaction
// whose value is the sum of all of the vote extensions from the previous round.
//
// Additionally, we verify the vote extension signatures passed from CometBFT and
// include all data necessary for such verification in the special transaction's
// payload so that ProcessProposal at other nodes can also verify the proposer
// constructed the special transaction correctly.
//
// If vote extensions are enabled for the current height, PrepareProposal makes
// sure there was at least one non-empty vote extension whose signature it could verify.
// If vote extensions are not enabled for the current height, PrepareProposal makes
// sure non-empty vote extensions are not present.
//
// The special vote extension-generated transaction must fit within an empty block
// and takes precedence over all other transactions coming from the mempool.
func (app *Application) PrepareProposal(
	_ context.Context, req *abci.RequestPrepareProposal,
) (*abci.ResponsePrepareProposal, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Coherence: No need to call parseTx, as the check is stateless and has been performed by CheckTx

// When vote extensions are enabled, our generated transaction takes precedence
// over any supplied transaction that attempts to modify the "extensionSum" value.

// Coherence: No need to call parseTx, as the check is stateless and has been performed by CheckTx

// ProcessProposal implements part of the Application interface.
// It accepts any proposal that does not contain a malformed transaction.
// NOTE It is up to real Applications to effect punitive behavior in the cases ProcessProposal
// returns ResponseProcessProposal_REJECT, as it is evidence of misbehavior.
func (app *Application) ProcessProposal(_ context.Context, req *abci.RequestProcessProposal) (*abci.ResponseProcessProposal, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Additional check for vote extension-related txs

// ExtendVote will produce vote extensions in the form of random numbers to
// demonstrate vote extension nondeterminism.
//
// In the next block, if there are any vote extensions from the previous block,
// a new transaction will be proposed that updates a special value in the
// key/value store ("extensionSum") with the sum of all of the numbers collected
// from the vote extensions.
func (app *Application) ExtendVote(_ context.Context, req *abci.RequestExtendVote) (*abci.ResponseExtendVote, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// We don't care that these values are generated by a weak random number
// generator. It's just for test purposes.
//nolint:gosec // G404: Use of weak random number generator

// VerifyVoteExtension simply validates vote extensions from other validators
// without doing anything about them. In this case, it just makes sure that the
// vote extension is a well-formed integer value.
func (app *Application) VerifyVoteExtension(_ context.Context, req *abci.RequestVerifyVoteExtension) (*abci.ResponseVerifyVoteExtension, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// We don't allow vote extensions to be optional

func (app *Application) Rollback() error { _ = "STUB: not implemented"; return nil }

func (app *Application) getAppHeight() int64 { _ = "STUB: not implemented"; return 0 }

func (app *Application) checkHeightAndExtensions(isPrepareProcessProposal bool, height int64, callsite string) (int64, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

// at exactly voteExtHeight, PrepareProposal still has no extensions, see RFC100

func (app *Application) storeValidator(valUpdate *abci.ValidatorUpdate) error {
	_ = "STUB: not implemented"
	// Store validator data to verify extensions
	return nil
}

// validatorUpdates generates a validator set update.
func (app *Application) validatorUpdates(height uint64) (abci.ValidatorUpdates, error) {
	_ = "STUB: not implemented"
	return *new(abci.ValidatorUpdates), nil
}

// parseTx parses a tx in 'key=value' format into a key and value.
func parseTx(tx []byte) (string, string, error) { _ = "STUB: not implemented"; return "", "", nil }

func (app *Application) verifyAndSum(
	areExtensionsEnabled bool,
	currentHeight int64,
	extCommit *abci.ExtendedCommitInfo,
	callsite string,
) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Only interested in votes that can have extensions

// Vote extension signatures are always provided. Apps can use them to verify the integrity of extensions

// Reconstruct vote extension's signed bytes...

// the vote extension was signed in the previous height

//... and verify

// The extension's format should have been verified in VerifyVoteExtension

// verifyExtensionTx parses and verifies the payload of a vote extension-generated tx
func (app *Application) verifyExtensionTx(height int64, payload string) error {
	_ = "STUB: not implemented"
	return nil
}

// Final check that the proposer behaved correctly

// parseVoteExtension attempts to parse the given extension data into a positive
// integer value.
func parseVoteExtension(ext []byte) (int64, error) { _ = "STUB: not implemented"; return 0, nil }
