package state

import (
	"sync"
	"sync/atomic"

	dbm "github.com/cometbft/cometbft-db"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/libs/log"
	cmtstate "github.com/cometbft/cometbft/proto/tendermint/state"
	"github.com/cometbft/cometbft/types"
)

const (
	// persist validators every valSetCheckpointInterval blocks to avoid
	// LoadValidators taking too much time.
	// https://github.com/tendermint/tendermint/pull/3438
	// 100000 results in ~ 100ms to get 100 validators (see BenchmarkLoadValidators)
	valSetCheckpointInterval = 100000
)

//------------------------------------------------------------------------

func calcValidatorsKey(height int64) []byte { _ = "STUB: not implemented"; return nil }

func calcConsensusParamsKey(height int64) []byte { _ = "STUB: not implemented"; return nil }

func calcABCIResponsesKey(height int64) []byte { _ = "STUB: not implemented"; return nil }

//----------------------

var lastABCIResponseKey = []byte("lastABCIResponseKey")
var offlineStateSyncHeight = []byte("offlineStateSyncHeightKey")

//go:generate ../scripts/mockery_generate.sh Store

// Store defines the state store interface
//
// It is used to retrieve current state and save and load ABCI responses,
// validators and consensus parameters
type Store interface {
	// LoadFromDBOrGenesisFile loads the most recent state.
	// If the chain is new it will use the genesis file from the provided genesis file path as the current state.
	LoadFromDBOrGenesisFile(string) (State, error)
	// LoadFromDBOrGenesisDoc loads the most recent state.
	// If the chain is new it will use the genesis doc as the current state.
	LoadFromDBOrGenesisDoc(*types.GenesisDoc) (State, error)
	// Load loads the current state of the blockchain
	Load() (State, error)
	// LoadValidators loads the validator set at a given height
	LoadValidators(int64) (*types.ValidatorSet, error)
	// LoadFinalizeBlockResponse loads the abciResponse for a given height
	LoadFinalizeBlockResponse(int64) (*abci.ResponseFinalizeBlock, error)
	// LoadLastFinalizeBlockResponse loads the last abciResponse for a given height
	LoadLastFinalizeBlockResponse(int64) (*abci.ResponseFinalizeBlock, error)
	// LoadConsensusParams loads the consensus params for a given height
	LoadConsensusParams(int64) (types.ConsensusParams, error)
	// Save overwrites the previous state with the updated one
	Save(State) error
	// SaveFinalizeBlockResponse saves ABCIResponses for a given height
	SaveFinalizeBlockResponse(int64, *abci.ResponseFinalizeBlock) error
	// Bootstrap is used for bootstrapping state when not starting from a initial height.
	Bootstrap(State) error
	// PruneStates takes the height from which to start pruning and which height stop at
	PruneStates(fromHeight, toHeight, evidenceThresholdHeight int64, previouslyPrunedStates uint64) (uint64, error)
	// Saves the height at which the store is bootstrapped after out of band statesync
	SetOfflineStateSyncHeight(height int64) error
	// Gets the height at which the store is bootstrapped after out of band statesync
	GetOfflineStateSyncHeight() (int64, error)
	// Close closes the connection with the database
	Close() error
}

// dbStore wraps a db (github.com/cometbft/cometbft-db).
type dbStore struct {
	db dbm.DB

	StoreOptions

	// compaction is kept behind a pointer so dbStore stays copyable (its
	// fields contain atomic.Bool and sync.WaitGroup which must not be
	// copied). Must always be non-nil; every dbStore literal in this
	// package initializes it.
	compaction *compactionState
}

// compactionState single-flights forced compaction (`inFlight`) and lets
// Close wait for the background goroutine (`wg`).
type compactionState struct {
	inFlight atomic.Bool
	wg       sync.WaitGroup
}

type StoreOptions struct {
	// DiscardABCIResponses determines whether or not the store
	// retains all ABCIResponses. If DiscardABCIResponses is enabled,
	// the store will maintain only the response object from the latest
	// height.
	DiscardABCIResponses bool

	Compact bool

	CompactionInterval int64

	Logger log.Logger
}

var _ Store = (*dbStore)(nil)

func IsEmpty(store dbStore) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// NewStore creates the dbStore of the state pkg.
func NewStore(db dbm.DB, options StoreOptions) Store { _ = "STUB: not implemented"; return *new(Store) }

// LoadStateFromDBOrGenesisFile loads the most recent state from the database,
// or creates a new one from the given genesisFilePath.
func (store dbStore) LoadFromDBOrGenesisFile(genesisFilePath string) (State, error) {
	_ = "STUB: not implemented"
	return *new(State), nil
}

// LoadStateFromDBOrGenesisDoc loads the most recent state from the database,
// or creates a new one from the given genesisDoc.
func (store dbStore) LoadFromDBOrGenesisDoc(genesisDoc *types.GenesisDoc) (State, error) {
	_ = "STUB: not implemented"
	return *new(State), nil
}

// LoadState loads the State from the database.
func (store dbStore) Load() (State, error) { _ = "STUB: not implemented"; return *new(State), nil }

func (store dbStore) loadState(key []byte) (state State, err error) {
	_ = "STUB: not implemented"
	return *new(State), nil
}

// DATA HAS BEEN CORRUPTED OR THE SPEC HAS CHANGED

// Save persists the State, the ValidatorsInfo, and the ConsensusParamsInfo to the database.
// This flushes the writes (e.g. calls SetSync).
func (store dbStore) Save(state State) error { _ = "STUB: not implemented"; return nil }

func (store dbStore) save(state State, key []byte) error { _ = "STUB: not implemented"; return nil }

// If first block, save validators for the block.

// This extra logic due to validator set changes being delayed 1 block.
// It may get overwritten due to InitChain validator updates.

// Save next validators.

// Save next consensus params.

// BootstrapState saves a new state, used e.g. by state sync when starting from non-zero height.
func (store dbStore) Bootstrap(state State) error { _ = "STUB: not implemented"; return nil }

// PruneStates deletes states between the given heights (including from, excluding to). It is not
// guaranteed to delete all states, since the last checkpointed state and states being pointed to by
// e.g. `LastHeightChanged` must remain. The state at to must also exist.
//
// The from parameter is necessary since we can't do a key scan in a performant way due to the key
// encoding not preserving ordering: https://github.com/tendermint/tendermint/issues/4567
// This will cause some old states to be left behind when doing incremental partial prunes,
// specifically older checkpoints and LastHeightChanged targets.
func (store dbStore) PruneStates(from int64, to int64, evidenceThresholdHeight int64, previouslyPrunedStates uint64) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// keep last checkpoint too

// We have to delete in reverse order, to avoid deleting previous heights that have validator
// sets and consensus params that we may need to retrieve.

// For heights we keep, we must make sure they have the full validator set or consensus
// params, otherwise they will panic if they're retrieved directly (instead of
// indirectly via a LastHeightChanged pointer).

// else we keep the validator set because we might need
// it later on for evidence verification

// avoid batches growing too large by flushing to database regularly

// triggerCompactionAsync launches a background compaction of the state DB.
// The state store is small enough that compacting the whole DB is cheap, so
// no range scoping is needed. `inFlight` single-flights triggers and `wg`
// lets Close wait for the goroutine to finish.
func (store dbStore) triggerCompactionAsync(to int64) { _ = "STUB: not implemented"; return }

//------------------------------------------------------------------------

// TxResultsHash returns the root hash of a Merkle tree of
// ExecTxResulst responses (see ABCIResults.Hash)
//
// See merkle.SimpleHashFromByteSlices
func TxResultsHash(txResults []*abci.ExecTxResult) []byte { _ = "STUB: not implemented"; return nil }

// LoadFinalizeBlockResponse loads the DiscardABCIResponses for the given height from the
// database. If the node has D set to true, ErrABCIResponsesNotPersisted
// is persisted. If not found, ErrNoABCIResponsesForHeight is returned.
func (store dbStore) LoadFinalizeBlockResponse(height int64) (*abci.ResponseFinalizeBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check for an error or if the resp.AppHash is nil if so
// this means the unmarshalling should be a LegacyABCIResponses
// Depending on a source message content (serialized as ABCIResponses)
// there are instances where it can be deserialized as a FinalizeBlockResponse
// without causing an error. But the values will not be deserialized properly
// and, it will contain zero values, and one of them is an AppHash == nil
// This can be verified in the /state/compatibility_test.go file

// The data might be of the legacy ABCI response type, so
// we try to unmarshal that

// only return an error, this method is only invoked through the `/block_results` not for state logic and
// some tests, so no need to exit cometbft if there's an error, just return it.

// The state store contains the old format. Migrate to
// the new ResponseFinalizeBlock format. Note that the
// new struct expects the AppHash which we don't have.

// TODO: ensure that buf is completely read.

// LoadLastFinalizeBlockResponse loads the FinalizeBlockResponses from the most recent height.
// The height parameter is used to ensure that the response corresponds to the latest height.
// If not, an error is returned.
//
// This method is used for recovering in the case that we called the Commit ABCI
// method on the application but crashed before persisting the results.
func (store dbStore) LoadLastFinalizeBlockResponse(height int64) (*abci.ResponseFinalizeBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Here we validate the result by comparing its height to the expected height.

// It is possible if this is called directly after an upgrade that
// ResponseFinalizeBlock is nil. In which case we use the legacy
// ABCI responses

// sanity check

// SaveFinalizeBlockResponse persists the ResponseFinalizeBlock to the database.
// This is useful in case we crash after app.Commit and before s.Save().
// Responses are indexed by height so they can also be loaded later to produce
// Merkle proofs.
//
// CONTRACT: height must be monotonically increasing every time this is called.
func (store dbStore) SaveFinalizeBlockResponse(height int64, resp *abci.ResponseFinalizeBlock) error {
	_ = "STUB: not implemented"
	return nil
}

// strip nil values,

// If the flag is false then we save the ABCIResponse. This can be used for the /BlockResults
// query or to reindex an event using the command line.

// We always save the last ABCI response for crash recovery.
// This overwrites the previous saved ABCI Response.

//-----------------------------------------------------------------------------

// LoadValidators loads the ValidatorSet for a given height.
// Returns ErrNoValSetForHeight if the validator set can't be found for this height.
func (store dbStore) LoadValidators(height int64) (*types.ValidatorSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// mutate

func lastStoredHeightFor(height, lastHeightChanged int64) int64 {
	_ = "STUB: not implemented"
	return 0
}

// CONTRACT: Returned ValidatorsInfo can be mutated.
func loadValidatorsInfo(db dbm.DB, height int64) (*cmtstate.ValidatorsInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DATA HAS BEEN CORRUPTED OR THE SPEC HAS CHANGED

// TODO: ensure that buf is completely read.

// saveValidatorsInfo persists the validator set.
//
// `height` is the effective height for which the validator is responsible for
// signing. It should be called from s.Save(), right before the state itself is
// persisted.
func (store dbStore) saveValidatorsInfo(height, lastHeightChanged int64, valSet *types.ValidatorSet, batch dbm.Batch) error {
	_ = "STUB: not implemented"
	return nil
}

// Only persist validator set if it was updated or checkpoint height (see
// valSetCheckpointInterval) is reached.

//-----------------------------------------------------------------------------

// ConsensusParamsInfo represents the latest consensus params, or the last height it changed

// LoadConsensusParams loads the ConsensusParams for a given height.
func (store dbStore) LoadConsensusParams(height int64) (types.ConsensusParams, error) {
	_ = "STUB: not implemented"
	return *new(types.ConsensusParams), nil
}

func (store dbStore) loadConsensusParamsInfo(height int64) (*cmtstate.ConsensusParamsInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DATA HAS BEEN CORRUPTED OR THE SPEC HAS CHANGED

// TODO: ensure that buf is completely read.

// saveConsensusParamsInfo persists the consensus params for the next block to disk.
// It should be called from s.Save(), right before the state itself is persisted.
// If the consensus params did not change after processing the latest block,
// only the last height for which they changed is persisted.
func (store dbStore) saveConsensusParamsInfo(nextHeight, changeHeight int64, params types.ConsensusParams, batch dbm.Batch) error {
	_ = "STUB: not implemented"
	return nil
}

func (store dbStore) SetOfflineStateSyncHeight(height int64) error {
	_ = "STUB: not implemented"
	return nil
}

// Gets the height at which the store is bootstrapped after out of band statesync
func (store dbStore) GetOfflineStateSyncHeight() (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (store dbStore) Close() error { _ = "STUB: not implemented"; return nil }

// responseFinalizeBlockFromLegacy is a convenience function that takes the old abci responses and morphs
// it to the finalize block response. Note that the app hash is missing
func responseFinalizeBlockFromLegacy(legacyResp *cmtstate.LegacyABCIResponses) *abci.ResponseFinalizeBlock {
	_ = "STUB: not implemented"
	return nil
}

// Check for begin block and end block and only append events or assign values if they are not nil

// Add BeginBlock attribute to BeginBlock events

// Add EndBlock attribute to BeginBlock events

// NOTE: AppHash is missing in the response but will
// be caught and filled in consensus/replay.go

func int64FromBytes(bz []byte) int64 { _ = "STUB: not implemented"; return 0 }

func int64ToBytes(i int64) []byte { _ = "STUB: not implemented"; return nil }
