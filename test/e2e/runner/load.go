package main

import (
	"context"

	e2e "github.com/cometbft/cometbft/test/e2e/pkg"
	"github.com/cometbft/cometbft/types"
)

const workerPoolSize = 16

// Load generates transactions against the network until the given context is
// canceled.
func Load(ctx context.Context, testnet *e2e.Testnet) error { _ = "STUB: not implemented"; return nil }

// generate run ID on startup

// Monitor successful transactions, and abort on stalls.

// loadGenerate generates jobs until the context is canceled
func loadGenerate(ctx context.Context, txCh chan<- types.Tx, testnet *e2e.Testnet, id []byte) {
	_ = "STUB: not implemented"
	return
}

// A context with a timeout is created here to time the createTxBatch
// function out. If createTxBatch has not completed its work by the time
// the next batch is set to be sent out, then the context is canceled so that
// the current batch is halted, allowing the next batch to begin.

// createTxBatch creates new transactions and sends them into the txCh. createTxBatch
// returns when either a full batch has been sent to the txCh or the context
// is canceled.
func createTxBatch(ctx context.Context, txCh chan<- types.Tx, testnet *e2e.Testnet, id []byte) {
	_ = "STUB: not implemented"
	return
}

// loadProcess processes transactions by sending transactions received on the txCh
// to the client.
func loadProcess(ctx context.Context, txCh <-chan types.Tx, chSuccess chan<- struct{}, n *e2e.Node) {
	_ = "STUB: not implemented"
	return
}
