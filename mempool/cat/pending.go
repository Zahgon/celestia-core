package cat

import (
	"sync"
	"time"

	"github.com/cometbft/cometbft/types"
)

const defaultPendingSeenPerSigner = 128

type pendingSeenTx struct {
	signerKey string
	signer    []byte
	txKey     types.TxKey
	sequence  uint64
	peer      uint16
	requested bool
	lastPeer  uint16
}

func (p *pendingSeenTx) peerIDs() []uint16 { _ = "STUB: not implemented"; return nil }

type pendingSeenTracker struct {
	mu        sync.Mutex
	perSigner map[string][]*pendingSeenTx
	byTx      map[types.TxKey]*pendingSeenTx
	limit     int
}

func newPendingSeenTracker(limit int) *pendingSeenTracker { _ = "STUB: not implemented"; return nil }

func (ps *pendingSeenTracker) add(signer []byte, txKey types.TxKey, sequence uint64, peerID uint16) {
	_ = "STUB: not implemented"
	return
}

// First check if we already have this exact txKey

// Already tracking this tx, keep the first peer

// No existing entry for this (signer, sequence), so create a new one

func (ps *pendingSeenTracker) remove(txKey types.TxKey) *pendingSeenTx {
	_ = "STUB: not implemented"
	return nil
}

// get returns the pending entry for a txKey without removing it.
// Returns nil if not found.
func (ps *pendingSeenTracker) get(txKey types.TxKey) *pendingSeenTx {
	_ = "STUB: not implemented"
	return nil
}

func (ps *pendingSeenTracker) entriesForSigner(signer []byte) []*pendingSeenTx {
	_ = "STUB: not implemented"
	return nil
}

func (ps *pendingSeenTracker) removePeer(peerID uint16) { _ = "STUB: not implemented"; return }

func (ps *pendingSeenTracker) signerKeys() [][]byte { _ = "STUB: not implemented"; return nil }

func (ps *pendingSeenTracker) markRequested(txKey types.TxKey, peerID uint16, at time.Time) {
	_ = "STUB: not implemented"
	return
}

func (ps *pendingSeenTracker) markRequestFailed(txKey types.TxKey, peerID uint16) {
	_ = "STUB: not implemented"
	return
}
