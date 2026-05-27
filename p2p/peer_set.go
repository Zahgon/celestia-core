package p2p

import (
	"net"

	cmtsync "github.com/cometbft/cometbft/libs/sync"
)

// IPeerSet has a (immutable) subset of the methods of PeerSet.
type IPeerSet interface {
	Has(key ID) bool
	HasIP(ip net.IP) bool
	Get(key ID) Peer
	List() []Peer
	Size() int
}

//-----------------------------------------------------------------------------

// PeerSet is a special structure for keeping a table of peers.
// Iteration over the peers is super fast and thread-safe.
type PeerSet struct {
	mtx    cmtsync.Mutex
	lookup map[ID]*peerSetItem
	list   []Peer
}

type peerSetItem struct {
	peer  Peer
	index int
}

// NewPeerSet creates a new peerSet with a list of initial capacity of 256 items.
func NewPeerSet() *PeerSet { _ = "STUB: not implemented"; return nil }

// Add adds the peer to the PeerSet.
// It returns an error carrying the reason, if the peer is already present.
func (ps *PeerSet) Add(peer Peer) error { _ = "STUB: not implemented"; return nil }

// Appending is safe even with other goroutines
// iterating over the ps.list slice.

// Has returns true if the set contains the peer referred to by this
// peerKey, otherwise false.
func (ps *PeerSet) Has(peerKey ID) bool { _ = "STUB: not implemented"; return false }

// HasIP returns true if the set contains the peer referred to by this IP
// address, otherwise false.
func (ps *PeerSet) HasIP(peerIP net.IP) bool { _ = "STUB: not implemented"; return false }

// hasIP does not acquire a lock so it can be used in public methods which
// already lock.
func (ps *PeerSet) hasIP(peerIP net.IP) bool { _ = "STUB: not implemented"; return false }

// Get looks up a peer by the provided peerKey. Returns nil if peer is not
// found.
func (ps *PeerSet) Get(peerKey ID) Peer { _ = "STUB: not implemented"; return *new(Peer) }

// Remove discards peer by its Key, if the peer was previously memoized.
// Returns true if the peer was removed, and false if it was not found.
// in the set.
func (ps *PeerSet) Remove(peer Peer) bool { _ = "STUB: not implemented"; return false }

// Removing the peer has failed so we set a flag to mark that a removal was attempted.
// This can happen when the peer add routine from the switch is running in
// parallel to the receive routine of MConn.
// There is an error within MConn but the switch has not actually added the peer to the peer set yet.
// Setting this flag will prevent a peer from being added to a node's peer set afterwards.

// Create a new copy of the list but with one less item.
// (we must copy because we'll be mutating the list).

// If it's the last peer, that's an easy special case.

// Replace the popped item with the last item in the old list.

// Size returns the number of unique items in the peerSet.
func (ps *PeerSet) Size() int { _ = "STUB: not implemented"; return 0 }

// List returns the threadsafe list of peers.
func (ps *PeerSet) List() []Peer { _ = "STUB: not implemented"; return nil }
