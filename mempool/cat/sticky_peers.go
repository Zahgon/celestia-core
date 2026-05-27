package cat

import (
	"github.com/cometbft/cometbft/p2p"
)

const stickyHashNamespace = "cat/sticky/v1"

type stickyPeer struct {
	id   uint16
	peer p2p.Peer
}

type stickyPeerScore struct {
	stickyPeer
	score uint64
}

// selectStickyPeers returns up to limit peers deterministically ranked for the signer.
// Uses rendezvous (highest random weight) hashing to compute the ranking.
func selectStickyPeers(signer []byte, peers map[uint16]p2p.Peer, limit int, salt []byte) []stickyPeer {
	_ = "STUB: not implemented"
	return nil
}

// stickyScore64 hashes (signer, peerID, salt) -> uint64.
func stickyScore64(signer []byte, peerID string, salt []byte) uint64 {
	_ = "STUB: not implemented"
	return 0
}
