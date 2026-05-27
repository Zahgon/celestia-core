package cat

import (
	tmsync "github.com/cometbft/cometbft/libs/sync"
	"github.com/cometbft/cometbft/mempool"
	"github.com/cometbft/cometbft/p2p"
)

const firstPeerID = mempool.UnknownPeerID + 1

// mempoolIDs is a thread-safe map of peer IDs to shorter uint16 IDs used by the Reactor for tracking peer
// messages and peer state such as what transactions peers have seen
type mempoolIDs struct {
	mtx       tmsync.RWMutex
	peerMap   map[p2p.ID]uint16   // quick lookup table for peer ID to short ID
	nextID    uint16              // assumes that a node will never have over 65536 active peers
	activeIDs map[uint16]p2p.Peer // used to check if a given peerID key is used, the value doesn't matter
}

func newMempoolIDs() *mempoolIDs { _ = "STUB: not implemented"; return nil }

// reserve unknownPeerID(0) for mempoolReactor.BroadcastTx

// ReserveForPeer searches for the next unused ID and assigns it to the
// peer.
func (ids *mempoolIDs) ReserveForPeer(peer p2p.Peer) { _ = "STUB: not implemented"; return }

// Replace the existing peer object with the new one in case of reconnection

// nextPeerID returns the next unused peer ID to use.
// This assumes that ids's mutex is already locked.
func (ids *mempoolIDs) nextPeerID() uint16 { _ = "STUB: not implemented"; return 0 }

// Reclaim returns the ID reserved for the peer back to unused pool.
func (ids *mempoolIDs) Reclaim(peerID p2p.ID) uint16 { _ = "STUB: not implemented"; return 0 }

// GetIDForPeer returns the shorthand ID reserved for the peer.
func (ids *mempoolIDs) GetIDForPeer(peerID p2p.ID) uint16 { _ = "STUB: not implemented"; return 0 }

// GetPeer returns the peer for the given shorthand ID.
func (ids *mempoolIDs) GetPeer(id uint16) p2p.Peer {
	_ = "STUB: not implemented"
	return *new(p2p.Peer)
}

// GetAll returns all active peers.
func (ids *mempoolIDs) GetAll() map[uint16]p2p.Peer { _ = "STUB: not implemented"; return nil }

// make a copy of the map.

// Len returns the number of active peers.
func (ids *mempoolIDs) Len() int { _ = "STUB: not implemented"; return 0 }
