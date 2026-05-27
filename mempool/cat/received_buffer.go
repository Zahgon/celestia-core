package cat

import (
	"sync"
	"time"

	"github.com/cometbft/cometbft/mempool"
	"github.com/cometbft/cometbft/types"
)

// bufferedTx holds a transaction that are requested out-of-order within certain bounds, waiting to be processed
type bufferedTx struct {
	tx       *types.CachedTx
	txKey    types.TxKey
	txInfo   mempool.TxInfo
	peerID   string
	sequence uint64
	addedAt  time.Time
}

// receivedTxBuffer holds transactions that are requested out-of-order.
// Transactions are buffered by (signer, sequence) and processed in order
// once earlier sequences complete.
type receivedTxBuffer struct {
	mu sync.Mutex
	// buffers contains mapping: signer (as string) -> sequence -> buffered tx
	buffers map[string]map[uint64]*bufferedTx
	// countByPeer tracks how many transactions are buffered from each peer
	countByPeer map[string]int
}

// newReceivedTxBuffer creates a new buffer for out-of-order transactions
func newReceivedTxBuffer() *receivedTxBuffer { _ = "STUB: not implemented"; return nil }

// add stores a transaction in the buffer for later processing.
// Returns false if the buffer is full for this signer, peer is at capacity, or tx already exists.
func (b *receivedTxBuffer) add(signer []byte, seq uint64, tx *types.CachedTx, txKey types.TxKey, txInfo mempool.TxInfo, peerID string) bool {
	_ = "STUB: not implemented"
	return false
}

// Check buffer limit per peer

// Check if already buffered

// Check buffer limit per signer

// get retrieves a buffered transaction for the given signer and sequence.
// Returns nil if not found.
func (b *receivedTxBuffer) get(signer []byte, seq uint64) *bufferedTx {
	_ = "STUB: not implemented"
	return nil
}

// removeLowerSeqs deletes all buffered transactions with sequence <= seq
func (b *receivedTxBuffer) removeLowerSeqs(signer []byte, seq uint64) {
	_ = "STUB: not implemented"
	return
}

// Decrement peer count before removing

// Clean up empty signer map

// signerKeys returns all signers that have buffered transactions
func (b *receivedTxBuffer) signerKeys() [][]byte { _ = "STUB: not implemented"; return nil }

// countForPeer returns the number of buffered transactions from a specific peer
func (b *receivedTxBuffer) countForPeer(peerID string) int { _ = "STUB: not implemented"; return 0 }
