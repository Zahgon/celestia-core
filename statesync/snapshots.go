package statesync

import (
	"crypto/sha256"

	cmtsync "github.com/cometbft/cometbft/libs/sync"
	"github.com/cometbft/cometbft/p2p"
)

// snapshotKey is a snapshot key used for lookups.
type snapshotKey [sha256.Size]byte

// snapshot contains data about a snapshot.
type snapshot struct {
	Height   uint64
	Format   uint32
	Chunks   uint32
	Hash     []byte
	Metadata []byte

	trustedAppHash    []byte // populated by light client
	trustedAppVersion uint64 // populated by light client
}

// Key generates a snapshot key, used for lookups. It takes into account not only the height and
// format, but also the chunks, hash, and metadata in case peers have generated snapshots in a
// non-deterministic manner. All fields must be equal for the snapshot to be considered the same.
func (s *snapshot) Key() snapshotKey {
	_ = "STUB: not implemented"
	// Hash.Write() never returns an error.
	return *new(snapshotKey)
}

//nolint:staticcheck

// snapshotPool discovers and aggregates snapshots across peers.
type snapshotPool struct {
	cmtsync.Mutex
	snapshots     map[snapshotKey]*snapshot
	snapshotPeers map[snapshotKey]map[p2p.ID]p2p.Peer

	// indexes for fast searches
	formatIndex map[uint32]map[snapshotKey]bool
	heightIndex map[uint64]map[snapshotKey]bool
	peerIndex   map[p2p.ID]map[snapshotKey]bool

	// blacklists for rejected items
	formatBlacklist   map[uint32]bool
	peerBlacklist     map[p2p.ID]bool
	snapshotBlacklist map[snapshotKey]bool
}

// newSnapshotPool creates a new snapshot pool. The state source is used for
func newSnapshotPool() *snapshotPool { _ = "STUB: not implemented"; return nil }

// Add adds a snapshot to the pool, unless the peer has already sent recentSnapshots snapshots. It
// returns true if this was a new, non-blacklisted snapshot. The snapshot height is verified using
// the light client, and the expected app hash is set for the snapshot.
func (p *snapshotPool) Add(peer p2p.Peer, snapshot *snapshot) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Best returns the "best" currently known snapshot, if any.
func (p *snapshotPool) Best() *snapshot { _ = "STUB: not implemented"; return nil }

// GetPeer returns a random peer for a snapshot, if any.
func (p *snapshotPool) GetPeer(snapshot *snapshot) p2p.Peer {
	_ = "STUB: not implemented"
	return *new(p2p.Peer)
}

//nolint:gosec // G404: Use of weak random number generator

// GetPeers returns the peers for a snapshot.
func (p *snapshotPool) GetPeers(snapshot *snapshot) []p2p.Peer {
	_ = "STUB: not implemented"
	return nil
}

// sort results, for testability (otherwise order is random, so tests randomly fail)

// Ranked returns a list of snapshots ranked by preference. The current heuristic is very naïve,
// preferring the snapshot with the greatest height, then greatest format, then greatest number of
// peers. This can be improved quite a lot.
func (p *snapshotPool) Ranked() []*snapshot { _ = "STUB: not implemented"; return nil }

// Reject rejects a snapshot. Rejected snapshots will never be used again.
func (p *snapshotPool) Reject(snapshot *snapshot) { _ = "STUB: not implemented"; return }

// RejectFormat rejects a snapshot format. It will never be used again.
func (p *snapshotPool) RejectFormat(format uint32) { _ = "STUB: not implemented"; return }

// RejectPeer rejects a peer. It will never be used again.
func (p *snapshotPool) RejectPeer(peerID p2p.ID) { _ = "STUB: not implemented"; return }

// RemovePeer removes a peer from the pool, and any snapshots that no longer have peers.
func (p *snapshotPool) RemovePeer(peerID p2p.ID) { _ = "STUB: not implemented"; return }

// removePeer removes a peer. The caller must hold the mutex lock.
func (p *snapshotPool) removePeer(peerID p2p.ID) { _ = "STUB: not implemented"; return }

// removeSnapshot removes a snapshot. The caller must hold the mutex lock.
func (p *snapshotPool) removeSnapshot(key snapshotKey) { _ = "STUB: not implemented"; return }
