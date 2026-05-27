package cat

import (
	"time"

	tmsync "github.com/cometbft/cometbft/libs/sync"
	"github.com/cometbft/cometbft/types"
)

// SeenTxSet records transactions that have been
// seen by other peers but not yet by us
type SeenTxSet struct {
	mtx tmsync.Mutex
	set map[types.TxKey]timestampedPeerSet
}

type timestampedPeerSet struct {
	peers map[uint16]struct{}
	time  time.Time
}

func NewSeenTxSet() *SeenTxSet { _ = "STUB: not implemented"; return nil }

// maxSeenTxSetSize limits the number of unique tx keys tracked in the SeenTxSet
// to prevent unbounded memory growth from malicious peers flooding SeenTx messages.
// Each entry costs ~500 bytes (including Go map overhead and GC pressure),
// so 10M entries ≈ 5 GB worst-case.
const maxSeenTxSetSize = 10_000_000

func (s *SeenTxSet) Add(txKey types.TxKey, peer uint16) { _ = "STUB: not implemented"; return }

// Evict one random entry to make room.

func (s *SeenTxSet) RemoveKey(txKey types.TxKey) { _ = "STUB: not implemented"; return }

func (s *SeenTxSet) Remove(txKey types.TxKey, peer uint16) { _ = "STUB: not implemented"; return }

func (s *SeenTxSet) RemovePeer(peer uint16) { _ = "STUB: not implemented"; return }

func (s *SeenTxSet) Prune(limit time.Time) { _ = "STUB: not implemented"; return }

func (s *SeenTxSet) Has(txKey types.TxKey, peer uint16) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *SeenTxSet) Get(txKey types.TxKey) map[uint16]struct{} {
	_ = "STUB: not implemented"
	return nil
}

// make a copy of the struct to avoid concurrency issues

// Len returns the amount of cached items. Mostly used for testing.
func (s *SeenTxSet) Len() int { _ = "STUB: not implemented"; return 0 }

func (s *SeenTxSet) Reset() { _ = "STUB: not implemented"; return }
