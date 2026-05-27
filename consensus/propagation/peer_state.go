package propagation

import (
	"context"
	"sync/atomic"

	proptypes "github.com/cometbft/cometbft/consensus/propagation/types"
	"github.com/cometbft/cometbft/libs/bits"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cometbft/cometbft/libs/sync"
	"github.com/cometbft/cometbft/p2p"
)

const (
	// MaxUnverifiedProposals is the maximum number of compact blocks
	// that can be cached per peer. Keeps the lowest heights when full.
	MaxUnverifiedProposals = 24
)

type request struct {
	height int64
	round  int32
	index  uint32
}

// PeerState keeps track of haves and wants for each peer. This is used for
// block prop and catchup.
type PeerState struct {
	ctx    context.Context
	cancel context.CancelFunc
	peer   p2p.Peer

	mtx *sync.RWMutex
	// state organized the haves and wants for each data is indexed by height
	// and round.
	state map[int64]map[int32]*partState

	concurrentReqs    atomic.Int64
	receivedParts     chan partData
	receivedHaves     chan request
	canRequest        chan struct{}
	remainingRequests map[int64]map[int32]int

	logger log.Logger

	// consensusPeerState allows the propagation reactor to update peer state
	// in the consensus reactor. This enables both reactors to gossip data
	// while minimizing redundant bandwidth.
	consensusPeerState PeerStateEditor

	// unverifiedProposals stores compact blocks received from this peer
	// for heights we weren't ready to process. Indexed by height.
	// These have NOT been verified via the consensus reactor's verification
	// function. Limited to MaxUnverifiedProposals entries, keeping lowest heights.
	unverifiedProposals map[int64]*proptypes.CompactBlock
}

type partData struct {
	height int64
	round  int32
}

// newPeerState initializes and returns a new PeerState. This should be
// called for each peer.
func newPeerState(ctx context.Context, peer p2p.Peer, logger log.Logger) *PeerState {
	_ = "STUB: not implemented"
	return nil
}

// SetConsensusPeerState sets the consensus peer state editor for this peer
func (ps *PeerState) SetConsensusPeerState(editor PeerStateEditor) {
	_ = "STUB: not implemented"
	return
}

// GetConsensusPeerState returns the consensus peer state editor if available
func (ps *PeerState) GetConsensusPeerState() PeerStateEditor {
	_ = "STUB: not implemented"
	return *new(PeerStateEditor)
}

// MaxUnverifiedProposalHeight returns the highest cached unverified proposal height.
// Returns 0 if none exist.
func (ps *PeerState) MaxUnverifiedProposalHeight() int64 { _ = "STUB: not implemented"; return 0 }

// Initialize initializes the state for a given height and round in a
// thread-safe way.
func (d *PeerState) Initialize(height int64, round int32, size int) {
	_ = "STUB: not implemented"
	return
}

// initialize initializes the state for a given height and round. This method is
// not thread-safe.
func (d *PeerState) initialize(height int64, round int32, size int) {
	_ = "STUB: not implemented"
	// Initialize the inner map if it doesn't exist
	return
}

func (d *PeerState) IncreaseConcurrentReqs(add int64) { _ = "STUB: not implemented"; return }

func (d *PeerState) SetConcurrentReqs(count int64) { _ = "STUB: not implemented"; return }

func (d *PeerState) DecreaseRemainingRequests(height int64, round int32, sub int) {
	_ = "STUB: not implemented"
	return
}

func (d *PeerState) SetRemainingRequests(height int64, round int32, count int) {
	_ = "STUB: not implemented"
	return
}

func (d *PeerState) GetRemainingRequests(height int64, round int32) int {
	_ = "STUB: not implemented"
	return 0
}

func (d *PeerState) DecreaseConcurrentReqs(sub int64) { _ = "STUB: not implemented"; return }

// AddHaves sets the haves for a given height and round.
func (d *PeerState) AddHaves(height int64, round int32, haves *bits.BitArray) {
	_ = "STUB: not implemented"
	return
}

// AddWants sets the wants for a given height and round.
func (d *PeerState) AddWants(height int64, round int32, wants *bits.BitArray) {
	_ = "STUB: not implemented"
	return
}

// AddRequests sets the requests for a given height and round.
func (d *PeerState) AddRequests(height int64, round int32, requests *bits.BitArray) {
	_ = "STUB: not implemented"
	return
}

// SetHave sets the have bit for a given part.
// Returns an error if the state is not initialized.
func (d *PeerState) SetHave(height int64, round int32, part int) error {
	_ = "STUB: not implemented"
	return nil
}

// SetWant sets the want bit for a given part.
// Returns an error if the state is not initialized.
func (d *PeerState) SetWant(height int64, round int32, part int, wants bool) error {
	_ = "STUB: not implemented"
	return nil
}

// GetHaves retrieves the haves for a given height and round.
func (d *PeerState) GetHaves(height int64, round int32) (empty *bits.BitArray, has bool) {
	_ = "STUB: not implemented"
	return nil, false

	// create the maps if they don't exist
}

// GetWants retrieves the wants for a given height and round.
func (d *PeerState) GetWants(height int64, round int32) (empty *bits.BitArray, has bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// create the maps if they don't exist

// GetRequests retrieves the requests for a given height and round.
func (d *PeerState) GetRequests(height int64, round int32) (empty *bits.BitArray, has bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// create the maps if they don't exist

// WantsPart checks if the peer wants a given part.
func (d *PeerState) WantsPart(height int64, round int32, part uint32) bool {
	_ = "STUB: not implemented"
	return false
}

// DeleteHeight removes all haves and wants for a given height.
func (d *PeerState) DeleteHeight(height int64) { _ = "STUB: not implemented"; return }

func (d *PeerState) RequestsReady() { _ = "STUB: not implemented"; return }

func (d *PeerState) CanRequest() chan struct{} { _ = "STUB: not implemented"; return nil }

// prune removes all haves and wants for heights less than the given height,
// while keeping the last keepRecentRounds for the current height.
func (d *PeerState) prune(prunePastHeight int64) { _ = "STUB: not implemented"; return }

// Prune unverified proposals for heights <= prunePastHeight

// todo: prune rounds separately from heights

// StoreUnverifiedProposal caches a compact block for a future height.
// Returns true if stored, false if rejected (cache full with lower heights).
func (d *PeerState) StoreUnverifiedProposal(cb *proptypes.CompactBlock) bool {
	_ = "STUB: not implemented"
	return false
}

// If we already have a proposal for this height, replace it (last write wins)

// If cache has room, store directly

// Cache is full - find the highest height

// Only store if this height is lower than the highest cached
// (we prioritize catching up on nearest heights first)

// Reject - cache is full with lower heights

// GetUnverifiedProposal returns cached compact block for height, or nil.
func (d *PeerState) GetUnverifiedProposal(height int64) *proptypes.CompactBlock {
	_ = "STUB: not implemented"
	return nil
}

// DeleteUnverifiedProposal removes a cached compact block for a specific height.
func (d *PeerState) DeleteUnverifiedProposal(height int64) { _ = "STUB: not implemented"; return }

type partState struct {
	haves    *bits.BitArray
	wants    *bits.BitArray
	requests *bits.BitArray
}

// newpartState initializes and returns a new partState
func newpartState(size int, _ int64, _ int32) *partState { _ = "STUB: not implemented"; return nil }

func (p *partState) addHaves(haves *bits.BitArray) { _ = "STUB: not implemented"; return }

func (p *partState) addWants(wants *bits.BitArray) { _ = "STUB: not implemented"; return }

func (p *partState) addRequests(requests *bits.BitArray) { _ = "STUB: not implemented"; return }

// SetHave sets the have bit for a given part.
// TODO support setting the hash and the proof
func (p *partState) setHave(index int, has bool) { _ = "STUB: not implemented"; return }

// SetWant sets the want bit for a given part.
func (p *partState) setWant(part int, wants bool) { _ = "STUB: not implemented"; return }
