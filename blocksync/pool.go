package blocksync

import (
	"time"

	flow "github.com/cometbft/cometbft/libs/flowrate"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cometbft/cometbft/libs/service"
	cmtsync "github.com/cometbft/cometbft/libs/sync"
	"github.com/cometbft/cometbft/libs/trace"
	"github.com/cometbft/cometbft/p2p"
	"github.com/cometbft/cometbft/types"
)

/*
eg, L = latency = 0.1s
	P = num peers = 10
	FN = num full nodes
	BS = 1kB block size
	CB = 1 Mbit/s = 128 kB/s
	CB/P = 12.8 kB
	B/S = CB/P/BS = 12.8 blocks/s

	12.8 * 0.1 = 1.28 blocks on conn
*/

const (
	requestIntervalMS = 2

	// minBlockSizeBytes is the minimum block size (1 KB) used for dynamic retry timer calculation.
	minBlockSizeBytes = 1024

	// maxBlockSizeBytes is the maximum block size (20 MB) used for dynamic retry timer calculation.
	maxBlockSizeBytes = 20 * 1024 * 1024

	// minRetrySeconds is the minimum retry timeout for small blocks.
	minRetrySeconds = 5

	// maxRetrySeconds is the maximum retry timeout for large blocks.
	maxRetrySeconds = 60

	// minReqLimit is the maximum concurrent requests per peer for large blocks.
	// Large blocks should have fewer to avoid bandwidth saturation.
	minReqLimit = 2

	// maxReqLimit is the maximum concurrent requests per peer for large blocks.
	// Small blocks can have more concurrent requests.
	maxReqLimit = 10

	// blockSizeBufferCapacity is the number of block sizes to track for calculating max or average.
	blockSizeBufferCapacity = 70

	// Minimum recv rate to ensure we're receiving blocks from a peer fast
	// enough. If a peer is not sending us data at at least that rate, we
	// consider them to have timedout and we disconnect.
	//
	// Based on the experiments with [Osmosis](https://osmosis.zone/), the
	// minimum rate could be as high as 500 KB/s. However, we're setting it to
	// 128 KB/s for now to be conservative.
	minRecvRate = 128 * 1024 // 128 KB/s

	// peerConnWait is the time that must have elapsed since the pool routine
	// was created before we start making requests. This is to give the peer
	// routine time to connect to peers.
	peerConnWait = 3 * time.Second

	// defaultMaxRequesters is the default maximum number of concurrent block requesters.
	defaultMaxRequesters = 40
)

var peerTimeout = 120 * time.Second // not const so we can override with tests

/*
	Peers self report their heights when we join the block pool.
	Starting from our latest pool.height, we request blocks
	in sequence from peers that reported higher heights than ours.
	Every so often we ask peers what height they're on so we can keep going.

	Requests are continuously made for blocks of higher heights until
	the limit is reached. If most of the requests have no available peers, and we
	are not at peer limits, we can probably switch to consensus reactor
*/

// BlockPool keeps track of the block sync peers, block requests and block responses.
type BlockPool struct {
	service.BaseService
	startTime   time.Time
	startHeight int64

	mtx cmtsync.Mutex
	// block requests
	requesters map[int64]*bpRequester
	height     int64 // the lowest key in requesters.
	// peers
	peers              map[p2p.ID]*bpPeer
	bannedPeers        map[p2p.ID]time.Time
	sortedPeers        []*bpPeer // sorted by curRate, highest first
	maxPeerHeight      int64     // the biggest reported height
	reqLimit           int
	retryTimeout       time.Duration
	lastReceivedBlocks *blockStats

	// atomic
	numPending int32 // number of requests pending assignment or block response

	requestsCh chan<- BlockRequest
	errorsCh   chan<- peerError

	traceClient trace.Tracer
}

// NewBlockPool returns a new BlockPool with the height equal to start. Block
// requests and errors will be sent to requestsCh and errorsCh accordingly.
func NewBlockPool(start int64, requestsCh chan<- BlockRequest, errorsCh chan<- peerError) *BlockPool {
	_ = "STUB: not implemented"
	return nil
}

// newBlockPoolWithTracer returns a new BlockPool with custom tracer
func newBlockPoolWithTracer(start int64, requestsCh chan<- BlockRequest, errorsCh chan<- peerError, traceClient trace.Tracer) *BlockPool {
	_ = "STUB: not implemented"
	return nil
}

// Initialize with default parameters

// OnStart implements service.Service by spawning requesters routine and recording
// pool's start time.
func (pool *BlockPool) OnStart() error { _ = "STUB: not implemented"; return nil }

// recalculateParams updates request limit and retry timeout based on block size
func (pool *BlockPool) recalculateParams() { _ = "STUB: not implemented"; return }

// setting some value between minReqLimit and maxReqLimit to limit the number of requests we send to one peer
// this is done to maintain peer diversity, using inverse here, because for larger blocks we want to send smaller
// number of concurrent requests

// in the same way finding an adequate timeout given the block size,
// here for larger blocks we want to have larger timeouts

// spawns requesters as needed
func (pool *BlockPool) makeRequestersRoutine() { _ = "STUB: not implemented"; return }

// Check if we are within peerConnWait seconds of start time
// This gives us some time to connect to peers before starting a wave of requests

// Calculate the duration to sleep until peerConnWait seconds have passed since pool.startTime

// If we have enough requesters, wait for them to finish.

// If we're caught up, wait for a bit so reactor could finish or a higher height is reported.

// request for more blocks.

// Sleep for a bit to make the requests more ordered.

func (pool *BlockPool) removeTimedoutPeers() { _ = "STUB: not implemented"; return }

// curRate can be 0 on start

// GetStatus returns pool's height, numPending requests and the number of
// requesters.
func (pool *BlockPool) GetStatus() (height int64, numPending int32, lenRequesters int) {
	_ = "STUB: not implemented"
	return 0, 0, 0
}

// IsCaughtUp returns true if this node is caught up, false - otherwise.
// TODO: relax conditions, prevent abuse.
func (pool *BlockPool) IsCaughtUp() bool { _ = "STUB: not implemented"; return false }

// Need at least 1 peer to be considered caught up.

// Some conditions to determine if we're caught up.
// Ensures we've either received a block or waited some amount of time,
// and that we're synced to the highest known height.
// Note we use maxPeerHeight - 1 because to sync block H requires block H+1
// to verify the LastCommit.
//
// maxPeerHeight == 0 happens in two distinct cases:
//   1. Fresh network: no peer has produced any blocks yet (all peer
//      heights are 0). We are caught up to the network and should switch
//      to consensus to participate in producing the first block.
//   2. Stalled: peers exist with blocks but updateMaxPeerHeight filtered
//      them all out (every peer's base is ahead of pool.height). The
//      network is ahead of us but no peer can serve our height, so we
//      are not caught up and must stay in blocksync.

// anyPeerHasBlocks reports whether any peer in the pool advertises a non-zero
// height. CONTRACT: pool.mtx must be locked.
func (pool *BlockPool) anyPeerHasBlocks() bool { _ = "STUB: not implemented"; return false }

// PeekTwoBlocks returns blocks at pool.height and pool.height+1. We need to
// see the second block's Commit to validate the first block. So we peek two
// blocks at a time. We return an extended commit, containing vote extensions
// and their associated signatures, as this is critical to consensus in ABCI++
// as we switch from block sync to consensus mode.
//
// The caller will verify the commit.
func (pool *BlockPool) PeekTwoBlocks() (first, second *types.Block, firstExtCommit *types.ExtendedCommit) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// SetHeight sets the current pool height under the mutex. Used by the
// blocksync reactor when switching from state sync, where pool.height must be
// advanced past the snapshot height before requesters start firing.
func (pool *BlockPool) SetHeight(height int64) { _ = "STUB: not implemented"; return }

// PopRequest removes the requester at pool.height and increments pool.height.
func (pool *BlockPool) PopRequest() { _ = "STUB: not implemented"; return }

// Re-evaluate maxPeerHeight: peers whose pruned base was just beyond the
// previous pool.height may now be able to contribute.

// RemovePeerAndRedoAllPeerRequests retries the request at the given height and
// all the requests made to the same peer. The peer is removed from the pool.
// Returns the ID of the removed peer.
func (pool *BlockPool) RemovePeerAndRedoAllPeerRequests(height int64) p2p.ID {
	_ = "STUB: not implemented"
	return *new(p2p.ID)
}

// this shouldn't happen, but to be on the safe side, let's ignore

// RemovePeer will redo all requesters associated with this peer.

// RedoRequestFrom retries the request at the given height. It does not remove the
// peer.
func (pool *BlockPool) RedoRequestFrom(height int64, peerID p2p.ID) {
	_ = "STUB: not implemented"
	return
}

// If we requested this block
// From this specific peer

// AddBlock validates that the block comes from the peer it was expected from
// and calls the requester to store it.
//
// This requires an extended commit at the same height as the supplied block -
// the block contains the last commit, but we need the latest commit in case we
// need to switch over from block sync to consensus at this height. If the
// height of the extended commit and the height of the block do not match, we
// do not add the block and return an error.
// TODO: ensure that blocks come in order for each peer.
func (pool *BlockPool) AddBlock(peerID p2p.ID, block *types.Block, extCommit *types.ExtendedCommit, blockSize int) error {
	_ = "STUB: not implemented"
	return nil
}

// Peer sent us an invalid block => remove it.

// If the peer sent us a block we clearly didn't request, we disconnect.

// Check if this peer was recently banned. If so, this is likely a race condition
// where the block arrived after the peer was banned and reset from the requester.
// This is not an error, just a timing issue.

// Height returns the pool's height.
func (pool *BlockPool) Height() int64 { _ = "STUB: not implemented"; return 0 }

// MaxPeerHeight returns the highest reported height.
func (pool *BlockPool) MaxPeerHeight() int64 { _ = "STUB: not implemented"; return 0 }

// SetPeerRange sets the peer's alleged blockchain base and height.
func (pool *BlockPool) SetPeerRange(peerID p2p.ID, base int64, height int64) {
	_ = "STUB: not implemented"
	return
}

// A peer whose own reported base exceeds its own height is structurally
// impossible and treated as malicious.

// RemovePeer will redo all requesters associated with this peer.

// no need to sort because curRate is 0 at start.
// just add to the beginning so it's picked first by pickIncrAvailablePeer.

// RemovePeer removes the peer with peerID from the pool. If there's no peer
// with peerID, function is a no-op.
func (pool *BlockPool) RemovePeer(peerID p2p.ID) { _ = "STUB: not implemented"; return }

// CONTRACT: pool.mtx must be locked.
func (pool *BlockPool) removePeer(peerID p2p.ID) { _ = "STUB: not implemented"; return }

// Find a new peer with the biggest height and update maxPeerHeight if the
// peer's height was the biggest.

// updateMaxPeerHeight sets maxPeerHeight to the highest height among peers
// whose advertised range still covers pool.height. If no peers are left,
// maxPeerHeight is set to 0.
func (pool *BlockPool) updateMaxPeerHeight() { _ = "STUB: not implemented"; return }

// Block a malicious peer from poisoning maxPeerHeight with an
// inflated base/height pair no peer can actually serve, which
// would stall IsCaughtUp forever.

// IsPeerBanned returns true if the peer is banned.
func (pool *BlockPool) IsPeerBanned(peerID p2p.ID) bool { _ = "STUB: not implemented"; return false }

// CONTRACT: pool.mtx must be locked.
func (pool *BlockPool) isPeerBanned(peerID p2p.ID) bool {
	_ = "STUB: not implemented"
	// Todo: replace with cmttime.Since in future versions
	return false
}

// CONTRACT: pool.mtx must be locked.
func (pool *BlockPool) banPeer(peerID p2p.ID) { _ = "STUB: not implemented"; return }

// getRetryTimeout returns the current retry timeout (thread-safe)
func (pool *BlockPool) getRetryTimeout() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// Pick an available peer with the given height available.
// If no peers are available, returns nil.
// prevPeerID is a peer to avoid if possible (e.g., peer used for previous height).
// This promotes peer diversity by alternating between peers across consecutive blocks.
func (pool *BlockPool) pickIncrAvailablePeer(height int64, ignorePeerID, prevPeerID p2p.ID) *bpPeer {
	_ = "STUB: not implemented"
	return nil
}

// Peer that matches all criteria except prevPeerID

// If this is the prevPeerID, save as fallback but continue looking
// the idea is that we want to add diversity to peers, so we will try to alternate peers when loading
// this works a bit better with less amount of peers and large blocks

// Sort peers by curRate, highest first.
//
// CONTRACT: pool.mtx must be locked.
func (pool *BlockPool) sortPeers() { _ = "STUB: not implemented"; return }

func (pool *BlockPool) makeNextRequester(nextHeight int64) { _ = "STUB: not implemented"; return }

func (pool *BlockPool) sendRequest(height int64, peerID p2p.ID) { _ = "STUB: not implemented"; return }

func (pool *BlockPool) sendError(err error, peerID p2p.ID) { _ = "STUB: not implemented"; return }

// for debugging purposes
//
//nolint:unused
func (pool *BlockPool) debug() string { _ = "STUB: not implemented"; return "" }

//-------------------------------------

type bpPeer struct {
	didTimeout  bool
	curRate     int64
	numPending  int32
	height      int64
	base        int64
	pool        *BlockPool
	id          p2p.ID
	recvMonitor *flow.Monitor

	timeout *time.Timer

	logger log.Logger
}

func newBPPeer(pool *BlockPool, peerID p2p.ID, base int64, height int64) *bpPeer {
	_ = "STUB: not implemented"
	return nil
}

func (peer *bpPeer) setLogger(l log.Logger) { _ = "STUB: not implemented"; return }

func (peer *bpPeer) resetMonitor() { _ = "STUB: not implemented"; return }

func (peer *bpPeer) resetTimeout() { _ = "STUB: not implemented"; return }

func (peer *bpPeer) incrPending() { _ = "STUB: not implemented"; return }

func (peer *bpPeer) decrPending(recvSize int) { _ = "STUB: not implemented"; return }

func (peer *bpPeer) onTimeout() { _ = "STUB: not implemented"; return }

//-------------------------------------

// bpRequester requests a block from a peer.
//
// If the height is within minBlocksForSingleRequest blocks of the pool's
// height, it will send an additional request to another peer. This is to avoid
// a situation where blocksync is stuck because of a single slow peer. Note
// that it's okay to send a single request when the requested height is far
// from the pool's height. If the peer is slow, it will timeout and be replaced
// with another peer.
type bpRequester struct {
	service.BaseService

	pool               *BlockPool
	height             int64
	gotBlockCh         chan struct{}
	redoCh             chan p2p.ID         // redo may got multiple messages, add peerId to identify repeat
	peerID             p2p.ID              // peerID is the peer currently being requested from ("" if no active request)
	requestedFromPeers map[p2p.ID]struct{} // requestedFromPeers are all peers used by this requester

	mtx          cmtsync.Mutex
	gotBlockFrom p2p.ID
	block        *types.Block
	extCommit    *types.ExtendedCommit
}

func newBPRequester(pool *BlockPool, height int64) *bpRequester {
	_ = "STUB: not implemented"
	return nil
}

func (bpr *bpRequester) OnStart() error { _ = "STUB: not implemented"; return nil }

// Returns true if the peer(s) match and block doesn't already exist.
func (bpr *bpRequester) setBlock(block *types.Block, extCommit *types.ExtendedCommit, peerID p2p.ID) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// already got a block

func (bpr *bpRequester) getBlock() *types.Block { _ = "STUB: not implemented"; return nil }

func (bpr *bpRequester) getExtendedCommit() *types.ExtendedCommit {
	_ = "STUB: not implemented"
	return nil
}

// Returns true if we've requested a block from the given peer.
func (bpr *bpRequester) didRequestFrom(peerID p2p.ID) bool { _ = "STUB: not implemented"; return false }

// Returns the ID of the peer who sent us the block.
func (bpr *bpRequester) gotBlockFromPeerID() p2p.ID { _ = "STUB: not implemented"; return *new(p2p.ID) }

func (bpr *bpRequester) resetAll() p2p.ID { _ = "STUB: not implemented"; return *new(p2p.ID) }

func (bpr *bpRequester) tryRemoveBlock(peerID p2p.ID) bool { _ = "STUB: not implemented"; return false }

// Removes the block (IF we got it from the given peer) and resets the peer.
func (bpr *bpRequester) reset(peerID p2p.ID) (removedBlock bool) {
	_ = "STUB: not implemented"
	return false
}

// Tells bpRequester to pick another peer and try again.
// NOTE: Nonblocking, and does nothing if another redo
// was already requested.
func (bpr *bpRequester) redo(peerID p2p.ID) { _ = "STUB: not implemented"; return }

// pickPeerAndSendRequest picks a peer and sends a block request.
// Only sends a request if there is no active request already.
func (bpr *bpRequester) pickPeerAndSendRequest(ignorePeerID p2p.ID) {
	_ = "STUB: not implemented"
	// Check if there's already an active request - if so, don't make another one
	return
}

// Try to get the peer used for the previous height to promote diversity

// Responsible for making more requests as necessary
// Returns only when a block is found (e.g. AddBlock() is called)
func (bpr *bpRequester) requestRoutine() { _ = "STUB: not implemented"; return }

// If peers returned NoBlockResponse or bad block, reschedule requests.

// We got a block!
// Continue the for-loop and wait til Quit.

// BlockRequest stores a block request identified by the block Height and the PeerID responsible for
// delivering the block
type BlockRequest struct {
	Height int64
	PeerID p2p.ID
}

// interpolate performs a linear interpolation between two integer values
// use inverse == true if you want to have an inverse dependency, i.e.
// the closer is the value to maxValue, the closer is out value to minOut
func interpolate(isInverse bool, value, minValue, maxValue, minOut, maxOut int) (out int) {
	_ = "STUB: not implemented"

	// normalize into [0, 1]
	return 0
}

// scale back to [minOut, maxOut]

// clamp the out value
