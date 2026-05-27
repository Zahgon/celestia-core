package consensus

import (
	"errors"
	"sync"
	"sync/atomic"

	"github.com/cometbft/cometbft/consensus/propagation"

	cstypes "github.com/cometbft/cometbft/consensus/types"
	"github.com/cometbft/cometbft/libs/bits"
	cmtjson "github.com/cometbft/cometbft/libs/json"
	"github.com/cometbft/cometbft/libs/log"
	cmtsync "github.com/cometbft/cometbft/libs/sync"
	"github.com/cometbft/cometbft/libs/trace"
	"github.com/cometbft/cometbft/p2p"
	cmtcons "github.com/cometbft/cometbft/proto/tendermint/consensus"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	sm "github.com/cometbft/cometbft/state"
	"github.com/cometbft/cometbft/types"
)

const (
	StateChannel       = byte(0x20)
	DataChannel        = byte(0x21)
	VoteChannel        = byte(0x22)
	VoteSetBitsChannel = byte(0x23)

	maxMsgSize = 512 * 1024 // 512kb; NOTE/TODO: keep in sync with types.PartSet sizes.

	blocksToContributeToBecomeGoodPeer = 10000
	votesToContributeToBecomeGoodPeer  = 10000

	// ReactorIncomingMessageQueueSize the size of the reactor's message queue.
	ReactorIncomingMessageQueueSize = 5000
)

//-----------------------------------------------------------------------------

// Reactor defines a reactor for the consensus service.
type Reactor struct {
	p2p.BaseReactor // BaseService + p2p.Switch

	conS *State

	mtx      cmtsync.RWMutex
	waitSync bool
	eventBus *types.EventBus
	rs       *cstypes.RoundState

	Metrics     *Metrics
	traceClient trace.Tracer

	propagator propagation.Propagator

	// gossipDataEnabled controls whether the gossipDataRoutine should run
	gossipDataEnabled atomic.Bool
}

type ReactorOption func(*Reactor)

// NewReactor returns a new Reactor with the given
// consensusState.
func NewReactor(consensusState *State, propagator propagation.Propagator, waitSync bool, options ...ReactorOption) *Reactor {
	_ = "STUB: not implemented"
	return nil
}

// WithGossipDataEnabled sets whether the gossipDataRoutine should run
func WithGossipDataEnabled(enabled bool) ReactorOption {
	_ = "STUB: not implemented"
	return *new(ReactorOption)
}

// IsGossipDataEnabled returns whether the gossipDataRoutine should run
func (conR *Reactor) IsGossipDataEnabled() bool { _ = "STUB: not implemented"; return false }

// OnStart implements BaseService by subscribing to events, which later will be
// broadcasted to other peers and starting state if we're not in block sync.
func (conR *Reactor) OnStart() error { _ = "STUB: not implemented"; return nil }

// start routine that computes peer statistics for evaluating peer quality

// OnStop implements BaseService by unsubscribing from events and stopping
// state.
func (conR *Reactor) OnStop() { _ = "STUB: not implemented"; return }

// SwitchToConsensus switches from block_sync mode to consensus mode.
// It resets the state, turns off block_sync, and starts the consensus state-machine
func (conR *Reactor) SwitchToConsensus(state sm.State, skipWAL bool) {
	_ = "STUB: not implemented"
	return
}

// We need to lock, as we are not entering consensus state from State's `handleMsg` or `handleTimeout`

// We have no votes, so reconstruct LastCommit from SeenCommit

// NOTE: The line below causes broadcastNewRoundStepRoutine() to broadcast a
// NewRoundStepMessage.

// GetChannels implements Reactor
func (conR *Reactor) GetChannels() []*p2p.ChannelDescriptor {
	_ = "STUB: not implemented"
	// TODO optimize
	return nil
}

// maybe split between gossiping current block and catchup stuff
// once we gossip the whole block there's nothing left to send until next height or round

// InitPeer implements Reactor by creating a state for the peer.
func (conR *Reactor) InitPeer(peer p2p.Peer) (p2p.Peer, error) {
	_ = "STUB: not implemented"
	return *new(p2p.Peer), nil
}

// AddPeer implements Reactor by spawning multiple gossiping goroutines for the
// peer.
func (conR *Reactor) AddPeer(peer p2p.Peer) { _ = "STUB: not implemented"; return }

// Begin routines for this peer.

// Only start gossip data routine if the peer supports the DataChannel

// Send our state to peer.
// If we're block_syncing, broadcast a RoundStepMessage later upon SwitchToConsensus().

// RemovePeer is a noop.
func (conR *Reactor) RemovePeer(p2p.Peer, interface{}) { _ = "STUB: not implemented"; return }

// TODO
// ps, ok := peer.Get(PeerStateKey).(*PeerState)
// if !ok {
// 	panic(fmt.Sprintf("Peer %v has no state", peer))
// }
// ps.Disconnect()

// Receive implements Reactor
// NOTE: We process these messages even when we're block_syncing.
// Messages affect either a peer state or the consensus state.
// Peer state updates can happen in parallel, but processing of
// proposals, block parts, and votes are ordered by the receiveRoutine
// NOTE: blocks on consensus state for proposals, block parts, and votes
func (conR *Reactor) Receive(e p2p.Envelope) { _ = "STUB: not implemented"; return }

// receive is an inner function that returns an error if the message is incorrect
// but doesn't disconnect the peer (this allows us to test the behavior of the function)
func (conR *Reactor) receive(e p2p.Envelope) error { _ = "STUB: not implemented"; return nil }

// Get peer states

// Peer claims to have a maj23 for some BlockID at H,R,S,

// Respond with a VoteSetBitsMessage showing which votes we have.
// (and consequently shows which we don't have)

// TODO handle the proposal message case in the propagation reactor

// don't punish (leave room for soft upgrades)

// don't punish (leave room for soft upgrades)

// SetEventBus sets event bus.
func (conR *Reactor) SetEventBus(b *types.EventBus) { _ = "STUB: not implemented"; return }

// WaitSync returns whether the consensus reactor is waiting for state/block sync.
func (conR *Reactor) WaitSync() bool { _ = "STUB: not implemented"; return false }

//--------------------------------------

// subscribeToBroadcastEvents subscribes for new round steps and votes
// using internal pubsub defined on state to broadcast
// them to peers upon receiving.
func (conR *Reactor) subscribeToBroadcastEvents() { _ = "STUB: not implemented"; return }

func (conR *Reactor) unsubscribeFromBroadcastEvents() { _ = "STUB: not implemented"; return }

func (conR *Reactor) broadcastNewRoundStepMessage(rs *cstypes.RoundState) {
	_ = "STUB: not implemented"
	return
}

func (conR *Reactor) broadcastNewValidBlockMessage(rs *cstypes.RoundState) {
	_ = "STUB: not implemented"
	return
}

// Broadcasts HasVoteMessage to peers that care.
func (conR *Reactor) broadcastHasVoteMessage(vote *types.Vote) { _ = "STUB: not implemented"; return }

/*
	// TODO: Make this broadcast more selective.
	for _, peer := range conR.Switch.Peers().List() {
		ps, ok := peer.Get(PeerStateKey).(*PeerState)
		if !ok {
			panic(fmt.Sprintf("Peer %v has no state", peer))
		}
		prs := ps.GetRoundState()
		if prs.Height == vote.Height {
			// TODO: Also filter on round?
			e := p2p.Envelope{
				ChannelID: StateChannel, struct{ ConsensusMessage }{msg},
				Message: p,
			}
			peer.TrySend(e)
		} else {
			// Height doesn't match
			// TODO: check a field, maybe CatchupCommitRound?
			// TODO: But that requires changing the struct field comment.
		}
	}
*/

func makeRoundStepMessage(rs *cstypes.RoundState) (nrsMsg *cmtcons.NewRoundStep) {
	_ = "STUB: not implemented"
	return nil
}

func (conR *Reactor) sendNewRoundStepMessage(peer p2p.Peer) { _ = "STUB: not implemented"; return }

func (conR *Reactor) updateRoundStateRoutine() { _ = "STUB: not implemented"; return }

func (conR *Reactor) getRoundState() *cstypes.RoundState { _ = "STUB: not implemented"; return nil }

func (conR *Reactor) gossipDataRoutine(peer p2p.Peer, ps *PeerState) {
	_ = "STUB: not implemented"
	return
}

// Exit early if gossip data is disabled

// Manage disconnects from self or peer.

// Send proposal Block parts?

// This tells peer that this part applies to us.
// This tells peer that this part applies to us.

// If the peer is on a previous height that we have, help catch up.

// if we never received the commit message from the peer, the block parts wont be initialized

// continue the loop since prs is a copy and not effected by this initialization

// If height and round don't match, sleep.

// logger.Info("Peer Height|Round mismatch, sleeping",
// "peerHeight", prs.Height, "peerRound", prs.Round, "peer", peer)

// By here, height and round match.
// Proposal block parts were already matched and sent if any were wanted.
// (These can match on hash so the round doesn't matter)
// Now consider sending other things, like the Proposal itself.

// Send Proposal && ProposalPOL BitArray?

// Proposal: share the proposal metadata with peer.

// NOTE[ZM]: A peer might have received different proposal msg so this Proposal msg will be rejected!

// ProposalPOL: lets peer know which POL votes we have so far.
// Peer must receive ProposalMessage first.
// rs.Proposal was validated, so rs.Proposal.POLRound <= rs.Round,
// so we definitely have rs.Votes.Prevotes(rs.Proposal.POLRound).

// Nothing to do. Sleep.

func (conR *Reactor) gossipDataForCatchup(logger log.Logger, rs *cstypes.RoundState,
	prs *cstypes.PeerRoundState, ps *PeerState, peer p2p.Peer,
) {
	_ = "STUB: not implemented"
	return
}

// Ensure that the peer's PartSetHeader is correct

// Load the part

// Send the part

// Not our height, so it doesn't matter.
// Not our height, so it doesn't matter.

//nolint:gosec

// sleep to avoid retrying too fast

//  logger.Info("No parts to send in catch-up, sleeping")

func (conR *Reactor) gossipVotesRoutine(peer p2p.Peer, ps *PeerState) {
	_ = "STUB: not implemented"
	return
}

// Simple hack to throttle logs upon sleep.

// Manage disconnects from self or peer.

// First sleep

// No more sleep

// logger.Debug("gossipVotesRoutine", "rsHeight", rs.Height, "rsRound", rs.Round,
// "prsHeight", prs.Height, "prsRound", prs.Round, "prsStep", prs.Step)

// If height matches, then send LastCommit, Prevotes, Precommits.

// Special catchup logic.
// If peer is lagging by height 1, send LastCommit.

// Catchup logic
// If peer is lagging by more than 1, send Commit.

// Load the block's extended commit for prs.Height,
// which contains precommit signatures for prs.Height.

//nolint:staticcheck
// We sent nothing. Sleep...

// Continued sleep...

// pickSendVoteAndTrace picks a vote to send and traces it.
// It returns true if a vote is sent.
// Note that it is a wrapper around PickSendVote with the addition of tracing the vote.
func (conR *Reactor) pickSendVoteAndTrace(votes types.VoteSetReader, rs *cstypes.RoundState, ps *PeerState) bool {
	_ = "STUB: not implemented"
	return false
}

// if a vote is sent, trace it

func (conR *Reactor) gossipVotesForHeight(
	logger log.Logger,
	rs *cstypes.RoundState,
	prs *cstypes.PeerRoundState,
	ps *PeerState,
) bool {
	_ = "STUB: not implemented"
	// If there are lastCommits to send...
	return false
}

// If there are POL prevotes to send...

// If there are prevotes to send...

// If there are precommits to send...

// If there are prevotes to send...Needed because of validBlock mechanism

// If there are POLPrevotes to send...

// NOTE: `queryMaj23Routine` has a simple crude design since it only comes
// into play for liveness when there's a signature DDoS attack happening.
func (conR *Reactor) queryMaj23Routine(peer p2p.Peer, ps *PeerState) {
	_ = "STUB: not implemented"

	// Manage disconnects from self or peer.
	return
}

// Maybe send Height/Round/Prevotes

// Maybe send Height/Round/Precommits

// Maybe send Height/Round/ProposalPOL

// Little point sending LastCommitRound/LastCommit,
// These are fleeting and non-blocking.

// Maybe send Height/CatchupCommitRound/CatchupCommit.

func (conR *Reactor) peerStatsRoutine() { _ = "STUB: not implemented"; return }

// Get peer

// Get peer state

// String returns a string representation of the Reactor.
// NOTE: For now, it is just a hard-coded string to avoid accessing unprotected shared variables.
// TODO: improve!
func (conR *Reactor) String() string {
	_ = "STUB: not implemented"
	// better not to access shared variables
	return ""
}

// conR.StringIndented("")

// ReactorMetrics sets the metrics
func ReactorMetrics(metrics *Metrics) ReactorOption {
	_ = "STUB: not implemented"
	return *new(ReactorOption)
}

func ReactorTracing(traceClient trace.Tracer) ReactorOption {
	_ = "STUB: not implemented"
	return *new(ReactorOption)
}

//-----------------------------------------------------------------------------

var (
	ErrPeerStateHeightRegression = errors.New("error peer state height regression")
	ErrPeerStateInvalidStartTime = errors.New("error peer state invalid startTime")
)

// PeerState contains the known state of a peer, including its connection and
// threadsafe access to its PeerRoundState.
// NOTE: THIS GETS DUMPED WITH rpc/core/consensus.go.
// Be mindful of what you Expose.
type PeerState struct {
	peer   p2p.Peer
	logger log.Logger

	mtx   sync.Mutex             // NOTE: Modify below using setters, never directly.
	PRS   cstypes.PeerRoundState `json:"round_state"` // Exposed.
	Stats *peerStateStats        `json:"stats"`       // Exposed.
}

// peerStateStats holds internal statistics for a peer.
type peerStateStats struct {
	Votes      int `json:"votes"`
	BlockParts int `json:"block_parts"`
}

func (pss peerStateStats) String() string { _ = "STUB: not implemented"; return "" }

// NewPeerState returns a new PeerState for the given Peer
func NewPeerState(peer p2p.Peer) *PeerState { _ = "STUB: not implemented"; return nil }

// SetLogger allows to set a logger on the peer state. Returns the peer state
// itself.
func (ps *PeerState) SetLogger(logger log.Logger) *PeerState { _ = "STUB: not implemented"; return nil }

// GetRoundState returns an shallow copy of the PeerRoundState.
// There's no point in mutating it since it won't change PeerState.
func (ps *PeerState) GetRoundState() *cstypes.PeerRoundState { _ = "STUB: not implemented"; return nil }

// copy

// MarshalJSON implements the json.Marshaler interface.
func (ps *PeerState) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// GetHeight returns an atomic snapshot of the PeerRoundState's height
// used by the mempool to ensure peers are caught up before broadcasting new txs
func (ps *PeerState) GetHeight() int64 { _ = "STUB: not implemented"; return 0 }

// SetHasProposal sets the given proposal as known for the peer.
func (ps *PeerState) SetHasProposal(proposal *types.Proposal) { _ = "STUB: not implemented"; return }

// ps.PRS.ProposalBlockParts is set due to NewValidBlockMessage

// Nil until ProposalPOLMessage received.

// InitProposalBlockParts initializes the peer's proposal block parts header and bit array.
func (ps *PeerState) InitProposalBlockParts(partSetHeader types.PartSetHeader) {
	_ = "STUB: not implemented"
	return
}

// SetHasProposalBlockPart sets the given block part index as known for the peer.
func (ps *PeerState) SetHasProposalBlockPart(height int64, round int32, index int) {
	_ = "STUB: not implemented"
	return
}

// PickSendVote picks a vote and sends it to the peer.
// Returns true if vote was sent.
func (ps *PeerState) PickSendVote(votes types.VoteSetReader) *types.Vote {
	_ = "STUB: not implemented"
	return nil
}

// PickVoteToSend picks a vote to send to the peer.
// Returns true if a vote was picked.
// NOTE: `votes` must be the correct Size() for the Height().
func (ps *PeerState) PickVoteToSend(votes types.VoteSetReader) (vote *types.Vote, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// Lazily set data using 'votes'.

// Not something worth sending

func (ps *PeerState) getVoteBitArray(height int64, round int32, votesType cmtproto.SignedMsgType) *bits.BitArray {
	_ = "STUB: not implemented"
	return nil
}

// 'round': A round for which we have a +2/3 commit.
func (ps *PeerState) ensureCatchupCommitRound(height int64, round int32, numValidators int) {
	_ = "STUB: not implemented"
	return
}

/*
	NOTE: This is wrong, 'round' could change.
	e.g. if orig round is not the same as block LastCommit round.
	if ps.CatchupCommitRound != -1 && ps.CatchupCommitRound != round {
		panic(fmt.Sprintf(
			"Conflicting CatchupCommitRound. Height: %v,
			Orig: %v,
			New: %v",
			height,
			ps.CatchupCommitRound,
			round))
	}
*/

// Nothing to do!

// EnsureVoteBitArrays ensures the bit-arrays have been allocated for tracking
// what votes this peer has received.
// NOTE: It's important to make sure that numValidators actually matches
// what the node sees as the number of validators for height.
func (ps *PeerState) EnsureVoteBitArrays(height int64, numValidators int) {
	_ = "STUB: not implemented"
	return
}

func (ps *PeerState) ensureVoteBitArrays(height int64, numValidators int) {
	_ = "STUB: not implemented"
	return
}

//nolint:staticcheck

// RecordVote increments internal votes related statistics for this peer.
// It returns the total number of added votes.
func (ps *PeerState) RecordVote() int { _ = "STUB: not implemented"; return 0 }

// VotesSent returns the number of blocks for which peer has been sending us
// votes.
func (ps *PeerState) VotesSent() int { _ = "STUB: not implemented"; return 0 }

// RecordBlockPart increments internal block part related statistics for this peer.
// It returns the total number of added block parts.
func (ps *PeerState) RecordBlockPart() int { _ = "STUB: not implemented"; return 0 }

// BlockPartsSent returns the number of useful block parts the peer has sent us.
func (ps *PeerState) BlockPartsSent() int { _ = "STUB: not implemented"; return 0 }

// SetHasVote sets the given vote as known by the peer
func (ps *PeerState) SetHasVote(vote *types.Vote) { _ = "STUB: not implemented"; return }

func (ps *PeerState) setHasVote(height int64, round int32, voteType cmtproto.SignedMsgType, index int32) {
	_ = "STUB: not implemented"
	return
}

// NOTE: some may be nil BitArrays -> no side effects.

// ApplyNewRoundStepMessage updates the peer state for the new round.
func (ps *PeerState) ApplyNewRoundStepMessage(msg *NewRoundStepMessage) {
	_ = "STUB: not implemented"
	return
}

// Ignore duplicates or decreases

// Just remember these values.

// We'll update the BitArray capacity later.

// Peer caught up to CatchupCommitRound.
// Preserve psCatchupCommit!
// NOTE: We prefer to use prs.Precommits if
// pr.Round matches pr.CatchupCommitRound.

// Shift Precommits to LastCommit.

// We'll update the BitArray capacity later.

// ApplyNewValidBlockMessage updates the peer state for the new valid block.
func (ps *PeerState) ApplyNewValidBlockMessage(msg *NewValidBlockMessage) {
	_ = "STUB: not implemented"
	return
}

// ApplyProposalPOLMessage updates the peer state for the new proposal POL.
func (ps *PeerState) ApplyProposalPOLMessage(msg *ProposalPOLMessage) {
	_ = "STUB: not implemented"
	return
}

// TODO: Merge onto existing ps.PRS.ProposalPOL?
// We might have sent some prevotes in the meantime.

// ApplyHasVoteMessage updates the peer state for the new vote.
func (ps *PeerState) ApplyHasVoteMessage(msg *HasVoteMessage) { _ = "STUB: not implemented"; return }

// ApplyVoteSetBitsMessage updates the peer state for the bit-array of votes
// it claims to have for the corresponding BlockID.
// `ourVotes` is a BitArray of votes we have for msg.BlockID
// NOTE: if ourVotes is nil (e.g. msg.Height < rs.Height),
// we conservatively overwrite ps's votes w/ msg.Votes.
func (ps *PeerState) ApplyVoteSetBitsMessage(msg *VoteSetBitsMessage, ourVotes *bits.BitArray) {
	_ = "STUB: not implemented"
	return
}

// String returns a string representation of the PeerState
func (ps *PeerState) String() string { _ = "STUB: not implemented"; return "" }

// StringIndented returns a string representation of the PeerState
func (ps *PeerState) StringIndented(indent string) string { _ = "STUB: not implemented"; return "" }

//-----------------------------------------------------------------------------
// Messages

// Message is a message that can be sent and received on the Reactor
type Message interface {
	ValidateBasic() error
}

func init() {
	cmtjson.RegisterType(&NewRoundStepMessage{}, "tendermint/NewRoundStepMessage")
	cmtjson.RegisterType(&NewValidBlockMessage{}, "tendermint/NewValidBlockMessage")
	cmtjson.RegisterType(&ProposalMessage{}, "tendermint/Proposal")
	cmtjson.RegisterType(&ProposalPOLMessage{}, "tendermint/ProposalPOL")
	cmtjson.RegisterType(&BlockPartMessage{}, "tendermint/BlockPart")
	cmtjson.RegisterType(&VoteMessage{}, "tendermint/Vote")
	cmtjson.RegisterType(&HasVoteMessage{}, "tendermint/HasVote")
	cmtjson.RegisterType(&VoteSetMaj23Message{}, "tendermint/VoteSetMaj23")
	cmtjson.RegisterType(&VoteSetBitsMessage{}, "tendermint/VoteSetBits")
}

//-------------------------------------

// NewRoundStepMessage is sent for every step taken in the ConsensusState.
// For every height/round/step transition
type NewRoundStepMessage struct {
	Height                int64
	Round                 int32
	Step                  cstypes.RoundStepType
	SecondsSinceStartTime int64
	LastCommitRound       int32
}

// ValidateBasic performs basic validation.
func (m *NewRoundStepMessage) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// NOTE: SecondsSinceStartTime may be negative

// LastCommitRound will be -1 for the initial height, but we don't know what height this is
// since it can be specified in genesis. The reactor will have to validate this via
// ValidateHeight().

// ValidateHeight validates the height given the chain's initial height.
func (m *NewRoundStepMessage) ValidateHeight(initialHeight int64) error {
	_ = "STUB: not implemented"
	return nil
}

// String returns a string representation.
func (m *NewRoundStepMessage) String() string { _ = "STUB: not implemented"; return "" }

//-------------------------------------

// NewValidBlockMessage is sent when a validator observes a valid block B in some round r,
// i.e., there is a Proposal for block B and 2/3+ prevotes for the block B in the round r.
// In case the block is also committed, then IsCommit flag is set to true.
type NewValidBlockMessage struct {
	Height             int64
	Round              int32
	BlockPartSetHeader types.PartSetHeader
	BlockParts         *bits.BitArray
	IsCommit           bool
}

// ValidateBasic performs basic validation.
func (m *NewValidBlockMessage) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// String returns a string representation.
func (m *NewValidBlockMessage) String() string { _ = "STUB: not implemented"; return "" }

//-------------------------------------

// ProposalMessage is sent when a new block is proposed.
type ProposalMessage struct {
	Proposal *types.Proposal
}

// ValidateBasic performs basic validation.
func (m *ProposalMessage) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// String returns a string representation.
func (m *ProposalMessage) String() string { _ = "STUB: not implemented"; return "" }

//-------------------------------------

// ProposalPOLMessage is sent when a previous proposal is re-proposed.
type ProposalPOLMessage struct {
	Height           int64
	ProposalPOLRound int32
	ProposalPOL      *bits.BitArray
}

// ValidateBasic performs basic validation.
func (m *ProposalPOLMessage) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// String returns a string representation.
func (m *ProposalPOLMessage) String() string { _ = "STUB: not implemented"; return "" }

//-------------------------------------

// BlockPartMessage is sent when gossipping a piece of the proposed block.
type BlockPartMessage struct {
	Height int64
	Round  int32
	Part   *types.Part
}

// ValidateBasic performs basic validation.
func (m *BlockPartMessage) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// String returns a string representation.
func (m *BlockPartMessage) String() string { _ = "STUB: not implemented"; return "" }

//-------------------------------------

// VoteMessage is sent when voting for a proposal (or lack thereof).
type VoteMessage struct {
	Vote *types.Vote
}

// ValidateBasic checks whether the vote within the message is well-formed.
func (m *VoteMessage) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// String returns a string representation.
func (m *VoteMessage) String() string { _ = "STUB: not implemented"; return "" }

//-------------------------------------

// HasVoteMessage is sent to indicate that a particular vote has been received.
type HasVoteMessage struct {
	Height int64
	Round  int32
	Type   cmtproto.SignedMsgType
	Index  int32
}

// ValidateBasic performs basic validation.
func (m *HasVoteMessage) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// String returns a string representation.
func (m *HasVoteMessage) String() string { _ = "STUB: not implemented"; return "" }

//-------------------------------------

// VoteSetMaj23Message is sent to indicate that a given BlockID has seen +2/3 votes.
type VoteSetMaj23Message struct {
	Height  int64
	Round   int32
	Type    cmtproto.SignedMsgType
	BlockID types.BlockID
}

// ValidateBasic performs basic validation.
func (m *VoteSetMaj23Message) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// String returns a string representation.
func (m *VoteSetMaj23Message) String() string { _ = "STUB: not implemented"; return "" }

//-------------------------------------

// VoteSetBitsMessage is sent to communicate the bit-array of votes seen for the BlockID.
type VoteSetBitsMessage struct {
	Height  int64
	Round   int32
	Type    cmtproto.SignedMsgType
	BlockID types.BlockID
	Votes   *bits.BitArray
}

// ValidateBasic performs basic validation.
func (m *VoteSetBitsMessage) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// NOTE: Votes.Size() can be zero if the node does not have any

// String returns a string representation.
func (m *VoteSetBitsMessage) String() string { _ = "STUB: not implemented"; return "" }

//-------------------------------------
