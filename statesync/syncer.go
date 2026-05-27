package statesync

import (
	"context"
	"errors"
	"time"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/config"
	"github.com/cometbft/cometbft/libs/log"
	cmtsync "github.com/cometbft/cometbft/libs/sync"
	"github.com/cometbft/cometbft/p2p"
	"github.com/cometbft/cometbft/proxy"
	sm "github.com/cometbft/cometbft/state"
	"github.com/cometbft/cometbft/types"
)

const (
	// chunkTimeout is the timeout while waiting for the next chunk from the chunk queue.
	chunkTimeout = 2 * time.Minute

	// minimumDiscoveryTime is the lowest allowable time for a
	// SyncAny discovery time.
	minimumDiscoveryTime = 5 * time.Second
)

var (
	// errAbort is returned by Sync() when snapshot restoration is aborted.
	errAbort = errors.New("state sync aborted")
	// errRetrySnapshot is returned by Sync() when the snapshot should be retried.
	errRetrySnapshot = errors.New("retry snapshot")
	// errRejectSnapshot is returned by Sync() when the snapshot is rejected.
	errRejectSnapshot = errors.New("snapshot was rejected")
	// errRejectFormat is returned by Sync() when the snapshot format is rejected.
	errRejectFormat = errors.New("snapshot format was rejected")
	// errRejectSender is returned by Sync() when the snapshot sender is rejected.
	errRejectSender = errors.New("snapshot sender was rejected")
	// errVerifyFailed is returned by Sync() when app hash or last height verification fails.
	errVerifyFailed = errors.New("verification failed")
	// errTimeout is returned by Sync() when we've waited too long to receive a chunk.
	errTimeout = errors.New("timed out waiting for chunk")
	// errNoSnapshots is returned by SyncAny() if no snapshots are found and discovery is disabled.
	errNoSnapshots = errors.New("no suitable snapshots found")
)

// syncer runs a state sync against an ABCI app. Use either SyncAny() to automatically attempt to
// sync all snapshots in the pool (pausing to discover new ones), or Sync() to sync a specific
// snapshot. Snapshots and chunks are fed via AddSnapshot() and AddChunk() as appropriate.
type syncer struct {
	logger        log.Logger
	stateProvider StateProvider
	conn          proxy.AppConnSnapshot
	connQuery     proxy.AppConnQuery
	snapshots     *snapshotPool
	tempDir       string
	chunkFetchers int32
	retryTimeout  time.Duration

	mtx            cmtsync.RWMutex
	chunks         *chunkQueue
	onPeerRejected func(p2p.ID)
}

// newSyncer creates a new syncer.
func newSyncer(
	cfg config.StateSyncConfig,
	logger log.Logger,
	conn proxy.AppConnSnapshot,
	connQuery proxy.AppConnQuery,
	stateProvider StateProvider,
	tempDir string,
) *syncer {
	_ = "STUB: not implemented"
	return nil
}

// setOnPeerRejected sets a callback that is invoked when a peer is rejected by the ABCI
// application during chunk application. This allows the caller to disconnect the peer at the
// P2P layer, preventing them from continuing to send invalid chunks.
func (s *syncer) setOnPeerRejected(fn func(p2p.ID)) { _ = "STUB: not implemented"; return }

// AddChunk adds a chunk to the chunk queue, if any. It returns false if the chunk has already
// been added to the queue, or an error if there's no sync in progress.
func (s *syncer) AddChunk(chunk *chunk) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// AddSnapshot adds a snapshot to the snapshot pool. It returns true if a new, previously unseen
// snapshot was accepted and added.
func (s *syncer) AddSnapshot(peer p2p.Peer, snapshot *snapshot) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// AddPeer adds a peer to the pool. For now we just keep it simple and send a single request
// to discover snapshots, later we may want to do retries and stuff.
func (s *syncer) AddPeer(peer p2p.Peer) { _ = "STUB: not implemented"; return }

// RemovePeer removes a peer from the pool.
func (s *syncer) RemovePeer(peer p2p.Peer) { _ = "STUB: not implemented"; return }

// RejectPeer rejects a peer from the pool.
func (s *syncer) RejectPeer(peer p2p.Peer) { _ = "STUB: not implemented"; return }

// SyncAny tries to sync any of the snapshots in the snapshot pool, waiting to discover further
// snapshots if none were found and discoveryTime > 0. It returns the latest state and block commit
// which the caller must use to bootstrap the node.
func (s *syncer) SyncAny(discoveryTime time.Duration, retryHook func()) (sm.State, *types.Commit, error) {
	_ = "STUB: not implemented"
	return *new(sm.State), nil, nil
}

// The app may ask us to retry a snapshot restoration, in which case we need to reuse
// the snapshot and chunk queue from the previous loop iteration.

// If not nil, we're going to retry restoration of the same snapshot.

// in case we forget to close it elsewhere

// Discard snapshot and chunks for next iteration

// Sync executes a sync for a specific snapshot, returning the latest state and block commit which
// the caller must use to bootstrap the node.
func (s *syncer) Sync(snapshot *snapshot, chunks *chunkQueue) (sm.State, *types.Commit, error) {
	_ = "STUB: not implemented"
	return *new(sm.State), nil, nil
}

// Optimistically build new state, so we don't discover any light client failures at the end.

// Offer snapshot to ABCI app.

// Spawn chunk fetchers. They will terminate when the chunk queue is closed or context canceled.

// Restore snapshot

// Verify app and app version

// offerSnapshot offers a snapshot to the app. It returns various errors depending on the app's
// response, or nil if the snapshot was accepted.
func (s *syncer) offerSnapshot(snapshot *snapshot) error { _ = "STUB: not implemented"; return nil }

// applyChunks applies chunks to the app. It returns various errors depending on the app's
// response, or nil once the snapshot is fully restored.
func (s *syncer) applyChunks(chunks *chunkQueue) error { _ = "STUB: not implemented"; return nil }

// Discard and refetch any chunks as requested by the app

// Reject any senders as requested by the app

// fetchChunks requests chunks from peers, receiving allocations from the chunk queue. Chunks
// will be received from the reactor via syncer.AddChunks() to chunkQueue.Add().
func (s *syncer) fetchChunks(ctx context.Context, snapshot *snapshot, chunks *chunkQueue) {
	_ = "STUB: not implemented"
	return
}

// Keep checking until the context is canceled (restore is done), in case any
// chunks need to be refetched.

// requestChunk requests a chunk from a peer.
func (s *syncer) requestChunk(snapshot *snapshot, chunk uint32) { _ = "STUB: not implemented"; return }

// verifyApp verifies the sync, checking the app hash, last block height and app version
func (s *syncer) verifyApp(snapshot *snapshot, appVersion uint64) (abci.TimeoutInfo, error) {
	_ = "STUB: not implemented"
	return *new(abci.TimeoutInfo), nil
}

// sanity check that the app version in the block matches the application's own record
// of its version

// An error here most likely means that the app hasn't inplemented state sync
// or the Info call correctly
