package kvstore

import (
	"context"

	dbm "github.com/cometbft/cometbft-db"

	"github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/libs/log"
	cryptoproto "github.com/cometbft/cometbft/proto/tendermint/crypto"
)

var (
	stateKey        = []byte("stateKey")
	kvPairPrefixKey = []byte("kvPairKey:")
)

const (
	ValidatorPrefix        = "val="
	AppVersion      uint64 = 1
)

var _ types.Application = (*Application)(nil)

// Application is the kvstore state machine. It complies with the abci.Application interface.
// It takes transactions in the form of key=value and saves them in a database. This is
// a somewhat trivial example as there is no real state execution
type Application struct {
	types.BaseApplication

	state        State
	RetainBlocks int64 // blocks to retain after commit (via ResponseCommit.RetainHeight)
	stagedTxs    [][]byte
	logger       log.Logger

	// validator set
	valUpdates         []types.ValidatorUpdate
	valAddrToPubKeyMap map[string]cryptoproto.PublicKey

	// If true, the app will generate block events in BeginBlock. Used to test the event indexer
	// Should be false by default to avoid generating too much data.
	genBlockEvents bool
}

// NewApplication creates an instance of the kvstore from the provided database
func NewApplication(db dbm.DB) *Application { _ = "STUB: not implemented"; return nil }

// NewPersistentApplication creates a new application using the pebbledb database engine
func NewPersistentApplication(dbDir string) *Application { _ = "STUB: not implemented"; return nil }

// NewInMemoryApplication creates a new application from an in memory database.
// Nothing will be persisted.
func NewInMemoryApplication() *Application { _ = "STUB: not implemented"; return nil }

func (app *Application) SetGenBlockEvents() { _ = "STUB: not implemented"; return }

// Info returns information about the state of the application. This is generally used everytime a Tendermint instance
// begins and let's the application know what Tendermint versions it's interacting with. Based from this information,
// Tendermint will ensure it is in sync with the application by potentially replaying the blocks it has. If the
// Application returns a 0 appBlockHeight, Tendermint will call InitChain to initialize the application with consensus related data
func (app *Application) Info(context.Context, *types.RequestInfo) (*types.ResponseInfo, error) {
	_ = "STUB: not implemented"
	// Tendermint expects the application to persist validators, on start-up we need to reload them to memory if they exist
	return nil, nil
}

func (app *Application) timeoutInfo() types.TimeoutInfo {
	_ = "STUB: not implemented"
	return *new(types.TimeoutInfo)
}

// InitChain takes the genesis validators and stores them in the kvstore. It returns the application hash in the
// case that the application starts prepopulated with values. This method is called whenever a new instance of the application
// starts (i.e. app height = 0).
func (app *Application) InitChain(_ context.Context, req *types.RequestInitChain) (*types.ResponseInitChain, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CheckTx handles inbound transactions or in the case of recheckTx assesses old transaction validity after a state transition.
// As this is called frequently, it's preferably to keep the check as stateless and as quick as possible.
// Here we check that the transaction has the correctly key=value format.
// For the KVStore we check that each transaction has the valid tx format:
// - Contains one and only one `=`
// - `=` is not the first or last byte.
// - if key is `val` that the validator update transaction is also valid
func (app *Application) CheckTx(_ context.Context, req *types.RequestCheckTx) (*types.ResponseCheckTx, error) {
	_ = "STUB: not implemented"
	// If it is a validator update transaction, check that it is correctly formatted
	return nil, nil
}

//nolint:nilerr

// Tx must have a format like key:value or key=value. That is:
// - it must have one and only one ":" or "="
// - It must not begin or end with these special characters
func isValidTx(tx []byte) bool { _ = "STUB: not implemented"; return false }

// PrepareProposal is called when the node is a proposer. CometBFT stages a set of transactions to the application. As the
// KVStore has two accepted formats, `:` and `=`, we modify all instances of `:` with `=` to make it consistent. Note: this is
// quite a trivial example of transaction modification.
// NOTE: we assume that CometBFT will never provide more transactions than can fit in a block.
func (app *Application) PrepareProposal(ctx context.Context, req *types.RequestPrepareProposal) (*types.ResponsePrepareProposal, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// formatTxs validates and excludes invalid transactions
// also substitutes all the transactions with x:y to x=y
func (app *Application) formatTxs(ctx context.Context, blockData [][]byte) [][]byte {
	_ = "STUB: not implemented"
	return nil
}

// ProcessProposal is called whenever a node receives a complete proposal. It allows the application to validate the proposal.
// Only validators who can vote will have this method called. For the KVstore we reuse CheckTx.
func (app *Application) ProcessProposal(ctx context.Context, req *types.RequestProcessProposal) (*types.ResponseProcessProposal, error) {
	_ = "STUB: not implemented"
	return nil,

		// As CheckTx is a full validity check we can simply reuse this
		nil
}

// FinalizeBlock executes the block against the application state. It punishes validators who equivocated and
// updates validators according to transactions in a block. The rest of the transactions are regular key value
// updates and are cached in memory and will be persisted once Commit is called.
// ConsensusParams are never changed.
func (app *Application) FinalizeBlock(_ context.Context, req *types.RequestFinalizeBlock) (*types.ResponseFinalizeBlock, error) {
	_ = "STUB: not implemented"
	// reset valset changes
	return nil, nil
}

// Punish validators who committed equivocation.

// With every transaction we can emit a series of events. To make it simple, we just emit the same events.

// set signers for the tx

// Commit is called after FinalizeBlock and after Tendermint state which includes the updates to
// AppHash, ConsensusParams and ValidatorSet has occurred.
// The KVStore persists the validator updates and the new key values
func (app *Application) Commit(context.Context, *types.RequestCommit) (*types.ResponseCommit, error) {
	_ = "STUB: not implemented"
	// apply the validator updates to state (note this is really the validator set at h + 2)
	return nil, nil
}

// persist all the staged txs in the kvstore

// persist the state (i.e. size and height)

// Returns an associated value or nil if missing.
func (app *Application) Query(_ context.Context, reqQuery *types.RequestQuery) (*types.ResponseQuery, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO make Proof return index

func (app *Application) Close() error { _ = "STUB: not implemented"; return nil }

func isValidatorTx(tx []byte) bool { _ = "STUB: not implemented"; return false }

func parseValidatorTx(tx []byte) (string, []byte, int64, error) {
	_ = "STUB: not implemented"
	return "", nil,

		//  get the pubkey and power
		0, nil
}

// decode the pubkey

// decode the power

// add, update, or remove a validator
func (app *Application) updateValidator(v types.ValidatorUpdate) { _ = "STUB: not implemented"; return }

// remove validator

// add or update validator

func (app *Application) getValidators() (validators []types.ValidatorUpdate) {
	_ = "STUB: not implemented"
	return nil
}

// -----------------------------

type State struct {
	db dbm.DB
	// Size is essentially the amount of transactions that have been processes.
	// This is used for the appHash
	Size   int64 `json:"size"`
	Height int64 `json:"height"`
}

func loadState(db dbm.DB) State { _ = "STUB: not implemented"; return *new(State) }

func saveState(state State) { _ = "STUB: not implemented"; return }

// Hash returns the hash of the application state. This is computed
// as the size or number of transactions processed within the state. Note that this isn't
// a strong guarantee of state machine replication because states could
// have different kv values but still have the same size.
// This function is used as the "AppHash"
func (s State) Hash() []byte { _ = "STUB: not implemented"; return nil }

func prefixKey(key []byte) []byte { _ = "STUB: not implemented"; return nil }
