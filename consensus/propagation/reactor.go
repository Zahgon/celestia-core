package propagation

import (
	"context"
	"sync/atomic"
	"time"

	"github.com/cometbft/cometbft/crypto"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cometbft/cometbft/p2p/conn"

	"github.com/cometbft/cometbft/libs/sync"
	"github.com/cometbft/cometbft/store"
	"github.com/cometbft/cometbft/types"

	proptypes "github.com/cometbft/cometbft/consensus/propagation/types"
	"github.com/cometbft/cometbft/libs/trace"
	"github.com/cometbft/cometbft/p2p"
)

const (
	maxMsgSize = 512 * 1024 // 512kb

	// DataChannel the propagation reactor channel handling the haves, the compact block,
	// and the recovery parts.
	DataChannel = byte(0x50)

	// WantChannel the propagation reactor channel handling the wants.
	WantChannel = byte(0x51)

	// ReactorIncomingMessageQueueSize the size of the reactor's message queue.
	ReactorIncomingMessageQueueSize = 20000

	// RetryTime automatic catchup retry timeout.
	RetryTime = 2500 * time.Millisecond
)

type Reactor struct {
	p2p.BaseReactor // BaseService + p2p.Switch

	peerstate map[p2p.ID]*PeerState

	// ProposalCache temporarily stores recently active proposals and their
	// block data for gossiping.
	*ProposalCache
	currentProposer crypto.PubKey

	privval       types.PrivValidator
	chainID       string
	BlockMaxBytes int64

	// mempool access to read the transactions by hash from the mempool
	// and eventually remove it.
	mempool Mempool

	partChan     chan types.PartInfo
	proposalChan chan ProposalAndSrc

	mtx         *sync.Mutex
	traceClient trace.Tracer
	self        p2p.ID
	started     atomic.Bool
	ticker      *time.Ticker

	ctx    context.Context
	cancel context.CancelFunc
}

type Config struct {
	Store         *store.BlockStore
	Mempool       Mempool
	Privval       types.PrivValidator
	ChainID       string
	BlockMaxBytes int64
}

func NewReactor(
	self p2p.ID,
	config Config,
	options ...ReactorOption,
) *Reactor {
	_ = "STUB: not implemented"
	return nil
}

// start the catchup routine

// run the catchup routine to recover any missing parts for past heights.

type ReactorOption func(*Reactor)

func WithTracer(tracer trace.Tracer) func(r *Reactor) { _ = "STUB: not implemented"; return nil }

func (blockProp *Reactor) SetLogger(logger log.Logger) { _ = "STUB: not implemented"; return }

func (blockProp *Reactor) OnStart() error { _ = "STUB: not implemented"; return nil }

func (blockProp *Reactor) OnStop() { _ = "STUB: not implemented"; return }

func (blockProp *Reactor) GetChannels() []*conn.ChannelDescriptor {
	_ = "STUB: not implemented"
	return nil
}

// InitPeer initializes a new peer by checking if it is different from self or already exists in the peer list.
func (blockProp *Reactor) InitPeer(peer p2p.Peer) (p2p.Peer, error) {
	_ = "STUB: not implemented"
	// Ignore the peer if it is ourselves.
	return *new(p2p.Peer), nil
}

// ignore the peer if it already exists.

// AddPeer adds the peer to the block propagation reactor. This should be called when a peer
// is connected. The proposal is sent to the peer so that it can start catchup
// or request data.
func (blockProp *Reactor) AddPeer(peer p2p.Peer) { _ = "STUB: not implemented"; return }

// this means the compact block was created from catchup and no need to share it.
// otherwise, we need to correctly populate it.

// send the current proposal

func (blockProp *Reactor) RemovePeer(peer p2p.Peer, reason interface{}) {
	_ = "STUB: not implemented"
	return
}

func (blockProp *Reactor) ReceiveEnvelope(e p2p.Envelope) { _ = "STUB: not implemented"; return }

// Trace the received HaveParts message

// Trace the received WantParts message

func (blockProp *Reactor) Receive(e p2p.Envelope) { _ = "STUB: not implemented"; return }

// Prune removes all peer and proposal state from the block propagation reactor.
// This should be called only after a block has been committed.
func (blockProp *Reactor) Prune(committedHeight int64) { _ = "STUB: not implemented"; return }

func (blockProp *Reactor) SetProposer(proposer crypto.PubKey) { _ = "STUB: not implemented"; return }

// Check for cached proposals for the current height.
// This enables fast catchup when a node falls behind and misses proposals.

func (blockProp *Reactor) SetHeightAndRound(height int64, round int32) {
	_ = "STUB: not implemented"
	return
}

// If we already have a verified proposal cached for the current height
// (committed height + 1), treat it as catchup so delayed precommit is skipped.

// todo: delete the old round data as its no longer relevant don't delete
// past round data if it has a POL

// Check for cached proposals that might now be applicable.
// This handles the case where we advance to a new round and have a cached
// proposal for that round waiting to be applied.

func (blockProp *Reactor) ResetRequestCounts() { _ = "STUB: not implemented"; return }

// todo: investigate why nil peers can be present

func (blockProp *Reactor) StartProcessing() { _ = "STUB: not implemented"; return }

func ConcurrentRequestLimit(peersCount, partsCount int) int64 { _ = "STUB: not implemented"; return 0 }

// getPeer returns the peer state for the given peer. If the peer does not exist,
// nil is returned.
func (blockProp *Reactor) getPeer(peer p2p.ID) *PeerState { _ = "STUB: not implemented"; return nil }

// getPeers returns a list of all peers that the data routine is aware of.
func (blockProp *Reactor) getPeers() []*PeerState { _ = "STUB: not implemented"; return nil }

// setPeer sets the peer state for the given peer.
func (blockProp *Reactor) setPeer(peer p2p.ID, state *PeerState) { _ = "STUB: not implemented"; return }

func IsLegacyPropagation(peer p2p.Peer) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// GetPartChan returns the channel used for receiving part information.
func (r *Reactor) GetPartChan() <-chan types.PartInfo {
	_ = "STUB: not implemented"

	// GetProposalChan returns the channel used for receiving proposals.
	return nil
}

func (r *Reactor) GetProposalChan() <-chan ProposalAndSrc { _ = "STUB: not implemented"; return nil }

// GetUnverifiedProposal returns a cached compact block for the given height
// from any peer. Returns nil if none found. The returned compact block has NOT
// been verified via the consensus reactor's verification function - the caller
// must verify before use.
func (r *Reactor) GetUnverifiedProposal(height int64) *proptypes.CompactBlock {
	_ = "STUB: not implemented"
	return nil
}

// IsCatchingUp returns true if the node is catching up on block data
// (has unfinished heights that need to be downloaded).
func (r *Reactor) IsCatchingUp() bool {
	_ = "STUB: not implemented"
	// 1) Any unfinished heights (missing parts or marked catchup) => catching up.
	return false
}
