package core

import (
	ctypes "github.com/cometbft/cometbft/rpc/core/types"
	rpctypes "github.com/cometbft/cometbft/rpc/jsonrpc/types"
)

// Validators gets the validator set at the given block height.
//
// If no height is provided, it will fetch the latest validator set. Note the
// validators are sorted by their voting power - this is the canonical order
// for the validators in the set as used in computing their Merkle root.
//
// More: https://docs.cometbft.com/v0.38.x/rpc/#/Info/validators
func (env *Environment) Validators(
	_ *rpctypes.Context,
	heightPtr *int64,
	pagePtr, perPagePtr *int,
) (*ctypes.ResultValidators, error) {
	_ = "STUB: not implemented"
	// The latest validator that we know is the NextValidator of the last block.
	return nil, nil
}

// DumpConsensusState dumps consensus state.
// UNSTABLE
// More: https://docs.cometbft.com/v0.38.x/rpc/#/Info/dump_consensus_state
func (env *Environment) DumpConsensusState(*rpctypes.Context) (*ctypes.ResultDumpConsensusState, error) {
	_ = "STUB: not implemented"
	// Get Peer consensus states.
	return nil, nil
}

// peer does not have a state yet

// Peer basic info.

// Peer consensus state.

// Get self round state.

// ConsensusState returns a concise summary of the consensus state.
// UNSTABLE
// More: https://docs.cometbft.com/v0.38.x/rpc/#/Info/consensus_state
func (env *Environment) GetConsensusState(*rpctypes.Context) (*ctypes.ResultConsensusState, error) {
	_ = "STUB: not implemented"
	// Get self round state.
	return nil, nil
}

// ConsensusParams gets the consensus parameters at the given block height.
// If no height is provided, it will fetch the latest consensus params.
// More: https://docs.cometbft.com/v0.38.x/rpc/#/Info/consensus_params
func (env *Environment) ConsensusParams(
	_ *rpctypes.Context,
	heightPtr *int64,
) (*ctypes.ResultConsensusParams, error) {
	_ = "STUB: not implemented"
	// The latest consensus params that we know is the consensus params after the
	// last block.
	return nil, nil
}
