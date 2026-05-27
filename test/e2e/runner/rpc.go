package main

import (
	"context"
	"time"

	rpctypes "github.com/cometbft/cometbft/rpc/core/types"
	e2e "github.com/cometbft/cometbft/test/e2e/pkg"
	"github.com/cometbft/cometbft/types"
)

// waitForHeight waits for the network to reach a certain height (or above),
// returning the highest height seen. Errors if the network is not making
// progress at all.
func waitForHeight(ctx context.Context, testnet *e2e.Testnet, height int64) (*types.Block, *types.BlockID, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// waitForNode waits for a node to become available and catch up to the given block height.
func waitForNode(ctx context.Context, node *e2e.Node, height int64, timeout time.Duration) (*rpctypes.ResultStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// waitForAllNodes waits for all nodes to become available and catch up to the given block height.
func waitForAllNodes(ctx context.Context, testnet *e2e.Testnet, height int64, timeout time.Duration) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
