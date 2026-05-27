package core

import (
	"github.com/cometbft/cometbft/p2p"
	ctypes "github.com/cometbft/cometbft/rpc/core/types"
	rpctypes "github.com/cometbft/cometbft/rpc/jsonrpc/types"
	"github.com/cometbft/cometbft/types"
)

// Status returns CometBFT status including node info, pubkey, latest block
// hash, app hash, block height and time.
// More: https://docs.cometbft.com/v0.38.x/rpc/#/Info/status
func (env *Environment) Status(*rpctypes.Context) (*ctypes.ResultStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Return the very last voting power, not the voting power of this validator
// during the last block.

func (env *Environment) validatorAtHeight(h int64) *types.Validator {
	_ = "STUB: not implemented"
	return nil
}

// GetNodeInfo returns the node info with the app version set to the latest app
// version from the state store.
//
// This function is necessary because upstream CometBFT does not support
// upgrading app versions for a running binary. Therefore the
// env.P2PTransport.NodeInfo.ProtocolVersion.App is expected to be set on node
// start-up and never updated. Celestia supports upgrading the app version for a
// running binary so the env.P2PTransport.NodeInfo.ProtocolVersion.App will be
// incorrect if a node upgraded app versions without restarting. This function
// corrects that issue by fetching the latest app version from the state store.
func GetNodeInfo(env *Environment, latestHeight int64) p2p.DefaultNodeInfo {
	_ = "STUB: not implemented"
	return *new(p2p.DefaultNodeInfo)
}

// use the default app version if we can't load the consensus params (i.e. height 0)

// override the default app version with the latest app version
