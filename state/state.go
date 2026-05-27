package state

import (
	"time"

	cmtstate "github.com/cometbft/cometbft/proto/tendermint/state"
	"github.com/cometbft/cometbft/types"
)

// database keys
var (
	stateKey = []byte("stateKey")
)

//-----------------------------------------------------------------------------

// InitStateVersion sets the Consensus and Software versions.
func InitStateVersion(appVersion uint64) cmtstate.Version {
	_ = "STUB: not implemented"
	return *new(cmtstate.Version)
}

//-----------------------------------------------------------------------------

// State is a short description of the latest committed block of the consensus protocol.
// It keeps all information necessary to validate new blocks,
// including the last validator set and the consensus params.
// All fields are exposed so the struct can be easily serialized,
// but none of them should be mutated directly.
// Instead, use state.Copy() or state.NextState(...).
// NOTE: not goroutine-safe.
type State struct {
	Version cmtstate.Version

	// immutable
	ChainID       string
	InitialHeight int64 // should be 1, not 0, when starting from height 1

	// LastBlockHeight=0 at genesis (ie. block(H=0) does not exist)
	LastBlockHeight int64
	LastBlockID     types.BlockID
	LastBlockTime   time.Time

	// LastValidators is used to validate block.LastCommit.
	// Validators are persisted to the database separately every time they change,
	// so we can query for historical validator sets.
	// Note that if s.LastBlockHeight causes a valset change,
	// we set s.LastHeightValidatorsChanged = s.LastBlockHeight + 1 + 1
	// Extra +1 due to nextValSet delay.
	NextValidators              *types.ValidatorSet
	Validators                  *types.ValidatorSet
	LastValidators              *types.ValidatorSet
	LastHeightValidatorsChanged int64

	// Consensus parameters used for validating blocks.
	// Changes returned by FinalizeBlock and updated after Commit.
	ConsensusParams                  types.ConsensusParams
	LastHeightConsensusParamsChanged int64

	// Merkle root of the results from executing prev block
	LastResultsHash []byte

	// the latest AppHash we've received from calling abci.Commit()
	AppHash []byte

	// Timeouts from the application
	Timeouts cmtstate.TimeoutInfo
}

// Propose returns the amount of time to wait for a proposal using application timeouts
func (state State) Propose(round int32) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// Prevote returns the amount of time to wait for straggler votes after receiving any +2/3 prevotes using application timeouts
func (state State) Prevote(round int32) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// Precommit returns the amount of time to wait for straggler votes after receiving any +2/3 precommits using application timeouts
func (state State) Precommit(round int32) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// Commit returns the amount of time to wait for straggler votes after receiving +2/3 precommits using application timeouts
func (state State) Commit(t time.Time) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// Copy makes a copy of the State for mutating.
func (state State) Copy() State { _ = "STUB: not implemented"; return *new(State) }

// Equals returns true if the States are identical.
func (state State) Equals(state2 State) bool { _ = "STUB: not implemented"; return false }

// Bytes serializes the State using protobuf.
// It panics if either casting to protobuf or serialization fails.
func (state State) Bytes() []byte { _ = "STUB: not implemented"; return nil }

// IsEmpty returns true if the State is equal to the empty State.
func (state State) IsEmpty() bool { _ = "STUB: not implemented"; return false }

// XXX can't compare to Empty

// ToProto takes the local state type and returns the equivalent proto type
func (state *State) ToProto() (*cmtstate.State, error) { _ = "STUB: not implemented"; return nil, nil }

// At Block 1 LastValidators is nil

// FromProto takes a state proto message & returns the local state type
func FromProto(pb *cmtstate.State) (*State, error) { _ = "STUB: not implemented"; return nil, nil }

// At Block 1 LastValidators is nil

//------------------------------------------------------------------------
// Create a block from the latest state

// MakeBlock builds a block from the current state with the given txs, commit,
// and evidence. Note it also takes a proposerAddress because the state does not
// track rounds, and hence does not know the correct proposer. TODO: fix this!
func (state State) MakeBlock(
	height int64,
	data types.Data,
	lastCommit *types.Commit,
	evidence []types.Evidence,
	proposerAddress []byte,
) (*types.Block, *types.PartSet, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (state State) MakeBlockWithoutPartset(
	height int64,
	data types.Data,
	lastCommit *types.Commit,
	evidence []types.Evidence,
	proposerAddress []byte,
) (*types.Block, error) {
	_ = "STUB: not implemented"
	// Build base block with block data.
	return nil, nil
}

// Set time.

// genesis time

// Fill rest of header with state data.
//nolint:staticcheck

// MedianTime computes a median time for a given Commit (based on Timestamp field of votes messages) and the
// corresponding validator set. The computed time is always between timestamps of
// the votes sent by honest processes, i.e., a faulty processes can not arbitrarily increase or decrease the
// computed value.
func MedianTime(commit *types.Commit, validators *types.ValidatorSet) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

// If there's no condition, TestValidateBlockCommit panics; not needed normally.

//------------------------------------------------------------------------
// Genesis

// MakeGenesisStateFromFile reads and unmarshals state from the given
// file.
//
// Used during replay and in tests.
func MakeGenesisStateFromFile(genDocFile string) (State, error) {
	_ = "STUB: not implemented"
	return *new(State), nil
}

// MakeGenesisDocFromFile reads and unmarshals genesis doc from the given file.
func MakeGenesisDocFromFile(genDocFile string) (*types.GenesisDoc, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MakeGenesisState creates state from types.GenesisDoc.
func MakeGenesisState(genDoc *types.GenesisDoc) (State, error) {
	_ = "STUB: not implemented"
	return *new(State), nil
}

func getAppVersion(genDoc *types.GenesisDoc) uint64 { _ = "STUB: not implemented"; return 0 }

// Default to app version 1 because some chains (e.g. mocha-4) did not set
// an explicit app version in genesis.json.
