package main

import (
	"context"

	e2e "github.com/cometbft/cometbft/test/e2e/pkg"
	"github.com/cometbft/cometbft/test/e2e/pkg/infra"
)

func Start(ctx context.Context, testnet *e2e.Testnet, p infra.Provider) error {
	_ = "STUB: not implemented"
	return nil
}

// Nodes are already sorted by name. Sort them by name then startAt,
// which gives the overall order startAt, mode, name.

// Start initial nodes (StartAt: 0)

// Wait for initial height

// Update any state sync nodes with a trusted height and hash

// if we're starting a node that's ahead of
// the last known height of the network, then
// we should make sure that the rest of the
// network has reached at least the height
// that this node will start at before we
// start the node.
