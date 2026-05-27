package cat

import (
	"sync"
	"time"

	"github.com/cometbft/cometbft/types"
)

const defaultGlobalRequestTimeout = 1 * time.Hour

// requestScheduler tracks the lifecycle of outbound transaction requests.
type requestScheduler struct {
	mtx sync.Mutex

	// responseTime is the time the scheduler
	// waits for a response from a peer before
	// invoking the callback
	responseTime time.Duration

	// globalTimeout represents the longest duration
	// to wait for any late response (after the reponseTime).
	// After this period the request is garbage collected.
	globalTimeout time.Duration

	// requestsByPeer is a lookup table of requests by peer.
	// Multiple tranasctions can be requested by a single peer at one
	requestsByPeer map[uint16]requestSet

	// requestsByTx is a lookup table for requested txs.
	// There can only be one request per tx.
	requestsByTx map[types.TxKey]uint16
}

type requestSet map[types.TxKey]*time.Timer

func newRequestScheduler(responseTime, globalTimeout time.Duration) *requestScheduler {
	_ = "STUB: not implemented"
	return nil
}

func (r *requestScheduler) Add(key types.TxKey, peer uint16, onTimeout func(key types.TxKey, peer uint16)) bool {
	_ = "STUB: not implemented"
	return false
}

// not allowed to have more than one outgoing transaction at once

// trigger callback. Callback can `Add` the tx back to the scheduler

// We set another timeout because the peer could still send
// a late response after the first timeout and it's important
// to recognize that it is a transaction in response to a
// request and not a new transaction being broadcasted to the entire
// network. This timer cannot be stopped and is used to ensure
// garbage collection.

func (r *requestScheduler) ForTx(key types.TxKey) uint16 { _ = "STUB: not implemented"; return 0 }

func (r *requestScheduler) Has(peer uint16, key types.TxKey) bool {
	_ = "STUB: not implemented"
	return false
}

// CountForPeer returns the number of active requests to a specific peer.
func (r *requestScheduler) CountForPeer(peer uint16) int { _ = "STUB: not implemented"; return 0 }

func (r *requestScheduler) ClearAllRequestsFrom(peer uint16) requestSet {
	_ = "STUB: not implemented"
	return *new(requestSet)
}

func (r *requestScheduler) MarkReceived(peer uint16, key types.TxKey) bool {
	_ = "STUB: not implemented"
	return false
}

// Close stops all timers and clears all requests.
// Add should never be called after `Close`.
func (r *requestScheduler) Close() { _ = "STUB: not implemented"; return }
