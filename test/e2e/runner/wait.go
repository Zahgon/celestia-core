package main

import (
	"context"
	"time"

	e2e "github.com/cometbft/cometbft/test/e2e/pkg"
)

// Wait waits for a number of blocks to be produced, and for all nodes to catch
// up with it.
func Wait(ctx context.Context, testnet *e2e.Testnet, blocks int64) error {
	_ = "STUB: not implemented"
	return nil
}

// WaitUntil waits until a given height has been reached.
func WaitUntil(ctx context.Context, testnet *e2e.Testnet, height int64) error {
	_ = "STUB: not implemented"
	return nil
}

// waitingTime estimates how long it should take for a node to reach the height.
// More nodes in a network implies we may expect a slower network and may have to wait longer.
func waitingTime(nodes int, height int64) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}
