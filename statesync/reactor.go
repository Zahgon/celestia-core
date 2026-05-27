package statesync

import (
	"time"

	"github.com/cometbft/cometbft/config"
	cmtsync "github.com/cometbft/cometbft/libs/sync"
	"github.com/cometbft/cometbft/p2p"
	"github.com/cometbft/cometbft/proxy"
	sm "github.com/cometbft/cometbft/state"
	"github.com/cometbft/cometbft/types"
)

const (
	// SnapshotChannel exchanges snapshot metadata
	SnapshotChannel = byte(0x60)
	// ChunkChannel exchanges chunk contents
	ChunkChannel = byte(0x61)
	// recentSnapshots is the number of recent snapshots to send and receive per peer.
	recentSnapshots = 10

	// ReactorIncomingMessageQueueSize the size of the reactor's message queue.
	ReactorIncomingMessageQueueSize = 100
)

// Reactor handles state sync, both restoring snapshots for the local node and serving snapshots
// for other nodes.
type Reactor struct {
	p2p.BaseReactor

	cfg       config.StateSyncConfig
	conn      proxy.AppConnSnapshot
	connQuery proxy.AppConnQuery
	tempDir   string
	metrics   *Metrics

	// This will only be set when a state sync is in progress. It is used to feed received
	// snapshots and chunks into the sync.
	mtx    cmtsync.RWMutex
	syncer *syncer
}

// NewReactor creates a new state sync reactor.
func NewReactor(
	cfg config.StateSyncConfig,
	conn proxy.AppConnSnapshot,
	connQuery proxy.AppConnQuery,
	metrics *Metrics,
) *Reactor {
	_ = "STUB: not implemented"
	return nil
}

// GetChannels implements p2p.Reactor.
func (r *Reactor) GetChannels() []*p2p.ChannelDescriptor { _ = "STUB: not implemented"; return nil }

// OnStart implements p2p.Reactor.
func (r *Reactor) OnStart() error {
	_ = "STUB: not implemented"

	// AddPeer implements p2p.Reactor.
	return nil
}

func (r *Reactor) AddPeer(peer p2p.Peer) { _ = "STUB: not implemented"; return }

// RemovePeer implements p2p.Reactor.
func (r *Reactor) RemovePeer(peer p2p.Peer, _ interface{}) { _ = "STUB: not implemented"; return }

// Receive implements p2p.Reactor.
func (r *Reactor) Receive(e p2p.Envelope) { _ = "STUB: not implemented"; return }

// TODO: We may want to consider punishing the peer for certain errors

// recentSnapshots fetches the n most recent snapshots from the app
func (r *Reactor) recentSnapshots(n uint32) ([]*snapshot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Sync runs a state sync, returning the new state and last commit at the snapshot height.
// The caller must store the state and commit in the state database and block store.
func (r *Reactor) Sync(stateProvider StateProvider, discoveryTime time.Duration) (sm.State, *types.Commit, error) {
	_ = "STUB: not implemented"
	return *new(sm.State), nil, nil
}

// Request snapshots from all currently connected peers
