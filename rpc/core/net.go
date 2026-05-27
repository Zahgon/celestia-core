package core

import (
	ctypes "github.com/cometbft/cometbft/rpc/core/types"
	rpctypes "github.com/cometbft/cometbft/rpc/jsonrpc/types"
)

// NetInfo returns network info.
// More: https://docs.cometbft.com/v0.38.x/rpc/#/Info/net_info
func (env *Environment) NetInfo(*rpctypes.Context) (*ctypes.ResultNetInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: Should we include PersistentPeers and Seeds in here?
// PRO: useful info
// CON: privacy

// UnsafeDialSeeds dials the given seeds (comma-separated id@IP:PORT).
func (env *Environment) UnsafeDialSeeds(_ *rpctypes.Context, seeds []string) (*ctypes.ResultDialSeeds, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnsafeDialPeers dials the given peers (comma-separated id@IP:PORT),
// optionally making them persistent.
func (env *Environment) UnsafeDialPeers(
	_ *rpctypes.Context,
	peers []string,
	persistent, unconditional, private bool,
) (*ctypes.ResultDialPeers, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Genesis returns genesis file.
// More: https://docs.cometbft.com/v0.38.x/rpc/#/Info/genesis
func (env *Environment) Genesis(*rpctypes.Context) (*ctypes.ResultGenesis, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (env *Environment) GenesisChunked(_ *rpctypes.Context, chunk uint) (*ctypes.ResultGenesisChunk, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getIDs(peers []string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }
