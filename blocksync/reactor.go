package blocksync

import (
	"sync"

	"github.com/cometbft/cometbft/crypto"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cometbft/cometbft/libs/trace"
	"github.com/cometbft/cometbft/p2p"
	bcproto "github.com/cometbft/cometbft/proto/tendermint/blocksync"
	sm "github.com/cometbft/cometbft/state"
	"github.com/cometbft/cometbft/store"
)

const (
	// BlocksyncChannel is a channel for blocks and status updates (`BlockStore` height)
	BlocksyncChannel = byte(0x40)

	trySyncIntervalMS = 10

	// stop syncing when last block's time is
	// within this much of the system time.
	// stopSyncingDurationMinutes = 10

	// ask for best height every 10s
	statusUpdateIntervalSeconds = 10
	// check if we should switch to consensus reactor
	switchToConsensusIntervalSeconds = 1

	// ReactorIncomingMessageQueueSize the size of the reactor's message queue.
	ReactorIncomingMessageQueueSize = 100
)

type consensusReactor interface {
	// for when we switch from blocksync reactor and block sync to
	// the consensus machine
	SwitchToConsensus(state sm.State, skipWAL bool)
}

// ReactorOption defines a function argument for Reactor.
type ReactorOption func(*Reactor)

// ReactorVerifyData sets the verifyData field of the reactor.
func ReactorVerifyData(verifyData bool) ReactorOption {
	_ = "STUB: not implemented"
	return *new(ReactorOption)
}

type peerError struct {
	err    error
	peerID p2p.ID
}

func (e peerError) Error() string { _ = "STUB: not implemented"; return "" }

// Reactor handles long-term catchup syncing.
type Reactor struct {
	p2p.BaseReactor

	// immutable
	initialState sm.State

	blockExec     *sm.BlockExecutor
	store         sm.BlockStore
	pool          *BlockPool
	traceClient   trace.Tracer
	blockSync     bool
	verifyData    bool
	localAddr     crypto.Address
	poolRoutineWg sync.WaitGroup

	requestsCh <-chan BlockRequest
	errorsCh   <-chan peerError

	switchToConsensusMs int

	metrics *Metrics
}

// NewReactor returns new reactor instance.
func NewReactor(state sm.State, blockExec *sm.BlockExecutor, store *store.BlockStore,
	blockSync bool, metrics *Metrics, offlineStateSyncHeight int64, options ...ReactorOption,
) *Reactor {
	_ = "STUB: not implemented"
	return nil
}

// Function added to keep existing API.
func NewReactorWithAddr(state sm.State, blockExec *sm.BlockExecutor, store *store.BlockStore,
	blockSync bool, localAddr crypto.Address, metrics *Metrics, offlineStateSyncHeight int64, traceClient trace.Tracer, options ...ReactorOption,
) *Reactor {
	_ = "STUB: not implemented"
	return nil
}

// If state sync was performed offline and the stores were bootstrapped to height H
// the state store's lastHeight will be H while blockstore's Height and Base are still 0
// 1. This scenario should not lead to a panic in this case, which is indicated by
// having a OfflineStateSyncHeight > 0
// 2. We need to instruct the blocksync reactor to start fetching blocks from H+1
// instead of 0.

// It's okay to block since sendRequest is called from a separate goroutine
// (bpRequester#requestRoutine; 1 per each peer).

// must be bigger than peers count
// so we don't block in #Receive#pool.AddBlock

// SetLogger implements service.Service by setting the logger on reactor and pool.
func (bcR *Reactor) SetLogger(l log.Logger) { _ = "STUB: not implemented"; return }

//nolint:staticcheck

// OnStart implements service.Service.
func (bcR *Reactor) OnStart() error { _ = "STUB: not implemented"; return nil }

// SwitchToBlockSync is called by the state sync reactor when switching to block sync.
func (bcR *Reactor) SwitchToBlockSync(state sm.State) error { _ = "STUB: not implemented"; return nil }

// OnStop implements service.Service.
func (bcR *Reactor) OnStop() { _ = "STUB: not implemented"; return }

// GetChannels implements Reactor
func (bcR *Reactor) GetChannels() []*p2p.ChannelDescriptor { _ = "STUB: not implemented"; return nil }

// AddPeer implements Reactor by sending our state to peer.
func (bcR *Reactor) AddPeer(peer p2p.Peer) { _ = "STUB: not implemented"; return }

// it's OK if send fails. will try later in poolRoutine

// peer is added to the pool once we receive the first
// bcStatusResponseMessage from the peer and call pool.SetPeerRange

// RemovePeer implements Reactor by removing peer from the pool.
func (bcR *Reactor) RemovePeer(peer p2p.Peer, _ interface{}) { _ = "STUB: not implemented"; return }

// respondToPeer loads a block and sends it to the requesting peer,
// if we have it. Otherwise, we'll respond saying we don't have it.
func (bcR *Reactor) respondToPeer(msg *bcproto.BlockRequest, src p2p.Peer) (queued bool) {
	_ = "STUB: not implemented"
	return false
}

// Receive implements Reactor by handling 4 types of messages (look below).
func (bcR *Reactor) Receive(e p2p.Envelope) { _ = "STUB: not implemented"; return }

// Send peer our state.

// Got a peer status. Unverified.

func (bcR *Reactor) localNodeBlocksTheChain(state sm.State) bool {
	_ = "STUB: not implemented"
	return false
}

// Handle messages from the poolReactor telling the reactor what to do.
// NOTE: Don't sleep in the FOR_LOOP or otherwise slow it down!
func (bcR *Reactor) poolRoutine(stateSynced bool) { _ = "STUB: not implemented"; return }

// ask for status updates

// The "if" statement below is a bit confusing, so here is a breakdown
// of its logic and purpose:
//
// If we are at genesis (no block in the chain), we don't need VoteExtensions
// because the first block's LastCommit is empty anyway.
//
// If VoteExtensions were disabled for the previous height then we don't need
// VoteExtensions.
//
// If we have sync'd at least one block, then we are guaranteed to have extensions
// if we need them by the logic inside loop FOR_LOOP: it requires that the blocks
// it fetches have extensions if extensions were enabled during the height.
//
// If we already had extensions for the initial height (e.g. we are recovering),
// then we are guaranteed to have extensions for the last block (if required) even
// if we did not blocksync any block.
//

// If require extensions, but since we don't have them yet, then we cannot switch to consensus yet.

// else {
// should only happen during testing
// }

// chan time

// NOTE: It is a subtle mistake to process more than a single block
// at a time (e.g. 10) here, because we only TrySend 1 request per
// loop.  The ratio mismatch can result in starving of blocks, a
// sudden burst of requests and responses, and repeat.
// Consequently, it is better to split these routines rather than
// coupling them as it's written here.  TODO uncouple from request
// routine.

// See if there are any blocks to sync.

// we need to have fetched two consecutive blocks in order to
// perform blocksync verification

// Some sanity checks on heights

// Panicking because the block pool's height  MUST keep consistent with the state; the block pool is totally under our control

// Panicking because this is an obvious bug in the block pool, which is totally under our control

// Before priming didProcessCh for another check on the next
// iteration, break the loop if the BlockPool or the Reactor itself
// has quit. This avoids case ambiguity of the outer select when two
// channels are ready.

// Try again quickly next loop.

// Start timing validation

// Finally, verify the first block using the second's commit
// NOTE: we can probably make this more efficient, but note that calling
// first.Hash() doesn't verify the tx contents, so MakePartSet() is
// currently necessary.

// validate the block before we persist it

// Block sync doesn't check that the `Data` in a block is valid.
// Since celestia-core can't determine if the `Data` in a block
// is valid, the next line asks celestia-app to check if the
// block is valid via ProcessProposal. If this minRequesterIncrease wasn't
// performed, a malicious node could fabricate an alternative
// set of transactions that would cause a different app hash and
// thus cause this node to panic.

// Calculate validation duration

// if vote extensions were required at this height, ensure they exist.

// Verify all signatures in the extended commit since it will
// be persisted to the blockstore. Without this check, a
// malicious peer could provide a corrupted ExtendedCommit
// (e.g. with wrong ValidatorAddress fields) that would cause
// a panic on consensus restart when reconstructLastCommit
// calls ToExtendedVoteSet.

// NOTE: we've already removed the peer's request, but we
// still need to clean up the rest.

// NOTE: we've already removed the peer's request, but we
// still need to clean up the rest.

// Start timing block save

// TODO: batch saves so we dont persist to disk every block

// We use LastCommit here instead of extCommit. extCommit is not
// guaranteed to be populated by the peer if extensions are not enabled.
// Currently, the peer should provide an extCommit even if the vote extension data are absent
// but this may change so using second.LastCommit is safer.

// Calculate save duration

// Trace block saved after successful validation

// TODO: same thing for app - but we would need a way to
// get the hash without persisting the state

// TODO This is bad, are we zombie?

// BroadcastStatusRequest broadcasts `BlockStore` base and height.
func (bcR *Reactor) BroadcastStatusRequest() { _ = "STUB: not implemented"; return }
