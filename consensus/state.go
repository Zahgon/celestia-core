package consensus

import (
	"context"
	"errors"
	"sync/atomic"
	"time"

	"github.com/cometbft/cometbft/consensus/propagation"

	cfg "github.com/cometbft/cometbft/config"
	cstypes "github.com/cometbft/cometbft/consensus/types"
	"github.com/cometbft/cometbft/crypto"
	cmtevents "github.com/cometbft/cometbft/libs/events"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cometbft/cometbft/libs/service"
	cmtsync "github.com/cometbft/cometbft/libs/sync"
	"github.com/cometbft/cometbft/libs/trace"
	"github.com/cometbft/cometbft/p2p"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	sm "github.com/cometbft/cometbft/state"
	"github.com/cometbft/cometbft/types"
)

// Consensus sentinel errors
var (
	ErrInvalidProposalSignature   = errors.New("error invalid proposal signature")
	ErrInvalidProposalPOLRound    = errors.New("error invalid proposal POL round")
	ErrAddingVote                 = errors.New("error adding vote")
	ErrSignatureFoundInPastBlocks = errors.New("found signature from the same key")
	ErrProposalTooManyParts       = errors.New("proposal block has too many parts")

	errPubKeyIsNotSet             = errors.New("pubkey is not set. Look for \"Can't get private validator pubkey\" errors")
	errInvalidProposalHeightRound = errors.New("invalid proposal height/round")
)

var msgQueueSize = 20_000

// msgs from the reactor which may update the state
type msgInfo struct {
	Msg    Message `json:"msg"`
	PeerID p2p.ID  `json:"peer_key"`
}

// internally generated messages which may update the state
type timeoutInfo struct {
	Duration time.Duration         `json:"duration"`
	Height   int64                 `json:"height"`
	Round    int32                 `json:"round"`
	Step     cstypes.RoundStepType `json:"step"`
}

func (ti *timeoutInfo) String() string { _ = "STUB: not implemented"; return "" }

// interface to the mempool
type txNotifier interface {
	TxsAvailable() <-chan struct{}
}

// interface to the evidence pool
type evidencePool interface {
	// reports conflicting votes to the evidence pool to be processed into evidence
	ReportConflictingVotes(voteA, voteB *types.Vote)
}

// State handles execution of the consensus algorithm.
// It processes votes and proposals, and upon reaching agreement,
// commits blocks to the chain and executes them against the application.
// The internal state machine receives input from peers, the internal validator, and from a timer.
type State struct {
	service.BaseService

	// config details
	config        *cfg.ConsensusConfig
	privValidator types.PrivValidator // for signing votes

	// store blocks and commits
	blockStore sm.BlockStore

	// create and execute blocks
	blockExec *sm.BlockExecutor

	// notify us if txs are available
	txNotifier txNotifier

	// add evidence to the pool
	// when it's detected
	evpool evidencePool

	// rsMtx protects only access to fields of rs
	rsMtx cmtsync.RWMutex
	rs    cstypes.RoundState

	// stateMtx protects only access to fields of state
	stateMtx cmtsync.RWMutex
	state    sm.State // State until height-1.

	// mtx protects access to internal fields of State (excluding rs and state!)
	mtx cmtsync.Mutex
	// privValidator pubkey, memoized for the duration of one block
	// to avoid extra requests to HSM
	privValidatorPubKey crypto.PubKey

	// state changes may be triggered by: msgs from peers,
	// msgs from ourself, or by timeouts
	peerMsgQueue     chan msgInfo
	internalMsgQueue chan msgInfo
	timeoutTicker    TimeoutTicker

	// information about about added votes and block parts are written on this channel
	// so statistics can be computed by reactor
	statsMsgQueue chan msgInfo

	// we use eventBus to trigger msg broadcasts in the reactor,
	// and to notify external subscribers, eg. through a websocket
	eventBus *types.EventBus

	// a Write-Ahead Log ensures we can recover from any kind of crash
	// and helps us avoid signing conflicting votes
	wal          WAL
	replayMode   bool // so we don't log signing errors during replay
	doWALCatchup bool // determines if we even try to do the catchup

	// for tests where we want to limit the number of transitions the state makes
	nSteps int

	// some functions can be overwritten for testing
	decideProposal        func(height int64, round int32)
	doPrevote             func(height int64, round int32)
	setProposal           func(proposal *types.Proposal) error
	StartedPrecommitSleep atomic.Bool

	// closed when we finish shutting down
	done chan struct{}

	// afterPanicFn, if non-nil, is called after onExit when receiveRoutine
	// recovers from a panic, to trigger a full node shutdown.
	afterPanicFn func()

	// synchronous pubsub between consensus state and reactor.
	// state only emits EventNewRoundStep and EventVote
	evsw cmtevents.EventSwitch

	propagator           propagation.Propagator
	newHeightOrRoundChan chan struct{}

	// for reporting metrics
	metrics *Metrics

	// offline state sync height indicating to which height the node synced offline
	offlineStateSyncHeight int64

	// traceClient is used to trace the state machine.
	traceClient trace.Tracer

	// gossipDataEnabled controls whether the gossipDataRoutine should run
	gossipDataEnabled atomic.Bool

	// proposalReceivedTime tracks when the proposal was received for the current height/round
	// Used to calculate the duration until full block is received
	proposalReceivedTime time.Time
}

// StateOption sets an optional parameter on the State.
type StateOption func(*State)

// NewState returns a new State.
func NewState(
	config *cfg.ConsensusConfig,
	state sm.State,
	blockExec *sm.BlockExecutor,
	blockStore sm.BlockStore,
	propagator propagation.Propagator,
	txNotifier txNotifier,
	evpool evidencePool,
	options ...StateOption,
) *State {
	_ = "STUB: not implemented"
	return nil
}

// set function defaults (may be overwritten before calling Start)

// We have no votes, so reconstruct LastCommit from SeenCommit.

// In case of out of band performed statesync, the state store
// will have a state but no extended commit (as no block has been downloaded).
// If the height at which the vote extensions are enabled is lower
// than the height at which we statesync, consensus will panic because
// it will try to reconstruct the extended commit here.

// NOTE: we do not call scheduleRound0 yet, we do that upon Start()

// SetLogger implements Service.
func (cs *State) SetLogger(l log.Logger) { _ = "STUB: not implemented"; return }

//nolint:staticcheck

// SetAfterPanicFn sets a callback that is invoked after onExit when
// receiveRoutine recovers from a panic. This is used to trigger a full
// node shutdown so the process doesn't linger as a zombie.
func (cs *State) SetAfterPanicFn(fn func()) { _ = "STUB: not implemented"; return }

// SetEventBus sets event bus.
func (cs *State) SetEventBus(b *types.EventBus) { _ = "STUB: not implemented"; return }

// StateMetrics sets the metrics.
func StateMetrics(metrics *Metrics) StateOption {
	_ = "STUB: not implemented"
	return *new(StateOption)
}

// SetTraceClient sets the remote event collector.
func SetTraceClient(ec trace.Tracer) StateOption {
	_ = "STUB: not implemented"
	return *new(StateOption)
}

// SetGossipDataEnabled specifies whether the legacy block prop is enabled.
func SetGossipDataEnabled(enabled bool) StateOption {
	_ = "STUB: not implemented"
	return *new(StateOption)
}

// OfflineStateSyncHeight indicates the height at which the node
// statesync offline - before booting sets the metrics.
func OfflineStateSyncHeight(height int64) StateOption {
	_ = "STUB: not implemented"
	return *new(StateOption)
}

// String returns a string.
func (cs *State) String() string {
	_ = "STUB: not implemented"
	// better not to access shared variables
	return ""
}

// GetState returns a copy of the chain state.
func (cs *State) GetState() sm.State { _ = "STUB: not implemented"; return *new(sm.State) }

// GetLastHeight returns the last height committed.
// If there were no blocks, returns 0.
func (cs *State) GetLastHeight() int64 { _ = "STUB: not implemented"; return 0 }

// GetRoundState returns a shallow copy of the internal consensus state.
func (cs *State) GetRoundState() *cstypes.RoundState { _ = "STUB: not implemented"; return nil }

// copy

// GetRoundStateJSON returns a json of RoundState.
func (cs *State) GetRoundStateJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// GetRoundStateSimpleJSON returns a json of RoundStateSimple
func (cs *State) GetRoundStateSimpleJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetValidators returns a copy of the current validators.
func (cs *State) GetValidators() (int64, []*types.Validator) {
	_ = "STUB: not implemented"
	return 0, nil
}

// SetPrivValidator sets the private validator account for signing votes. It
// immediately requests pubkey and caches it.
func (cs *State) SetPrivValidator(priv types.PrivValidator) { _ = "STUB: not implemented"; return }

// SetTimeoutTicker sets the local timer. It may be useful to overwrite for
// testing.
func (cs *State) SetTimeoutTicker(timeoutTicker TimeoutTicker) { _ = "STUB: not implemented"; return }

// LoadCommit loads the commit for a given height.
func (cs *State) LoadCommit(height int64) *types.Commit { _ = "STUB: not implemented"; return nil }

// OnStart loads the latest state via the WAL, and starts the timeout and
// receive routines.
func (cs *State) OnStart() error {
	_ = "STUB: not implemented"
	// We may set the WAL in testing before calling Start, so only OpenWAL if its
	// still the nilWAL.
	return nil
}

// we need the timeoutRoutine for replay so
// we don't block on the tick chan.
// NOTE: we will get a build up of garbage go routines
// firing on the tockChan until the receiveRoutine is started
// to deal with them (by that point, at most one will be valid)

// We may have lost some votes if the process crashed reload from consensus
// log to catchup.

// 1) prep work

// 2) backup original WAL file

// 3) try to repair (WAL file will be overwritten!)

// reload WAL file

// Double Signing Risk Reduction

// now start the receiveRoutine

// schedule the first round!
// use GetRoundState so we don't race the receiveRoutine for access

// timeoutRoutine: receive requests for timeouts on tickChan and fire timeouts on tockChan
// receiveRoutine: serializes processing of proposoals, block parts, votes; coordinates state transitions
func (cs *State) startRoutines(maxSteps int) { _ = "STUB: not implemented"; return }

// loadWalFile loads WAL data from file. It overwrites cs.wal.
func (cs *State) loadWalFile() error { _ = "STUB: not implemented"; return nil }

// OnStop implements service.Service.
func (cs *State) OnStop() { _ = "STUB: not implemented"; return }

// WAL is stopped in receiveRoutine.

// Wait waits for the the main routine to return.
// NOTE: be sure to Stop() the event switch and drain
// any event channels or this may deadlock
func (cs *State) Wait() {
	_ = "STUB: not implemented"

	// OpenWAL opens a file to log all consensus messages and timeouts for
	// deterministic accountability.
	return
}

func (cs *State) OpenWAL(walFile string) (WAL, error) {
	_ = "STUB: not implemented"
	return *new(WAL), nil
}

//------------------------------------------------------------
// Public interface for passing messages into the consensus state, possibly causing a state transition.
// If peerID == "", the msg is considered internal.
// Messages are added to the appropriate queue (peer or internal).
// If the queue is full, the function may block.
// TODO: should these return anything or let callers just use events?

// AddVote inputs a vote.
func (cs *State) AddVote(vote *types.Vote, peerID p2p.ID) (added bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

// TODO: wait for event?!

// SetProposal inputs a proposal.
func (cs *State) SetProposal(proposal *types.Proposal, peerID p2p.ID) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: wait for event?!

// AddProposalBlockPart inputs a part of the proposal block.
func (cs *State) AddProposalBlockPart(height int64, round int32, part *types.Part, peerID p2p.ID) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: wait for event?!

// SetProposalAndBlock inputs the proposal and all block parts.
func (cs *State) SetProposalAndBlock(
	proposal *types.Proposal,
	block *types.Block, //nolint:revive
	parts *types.PartSet,
	peerID p2p.ID,
) error {
	_ = "STUB: not implemented"
	// TODO: Since the block parameter is not used, we should instead expose just a SetProposal method.
	return nil
}

//------------------------------------------------------------
// internal functions for managing the state

func (cs *State) updateHeight(height int64) { _ = "STUB: not implemented"; return }

func (cs *State) updateRoundStep(round int32, step cstypes.RoundStepType) {
	_ = "STUB: not implemented"
	return
}

// enterNewRound(height, 0) at cs.StartTime.
func (cs *State) scheduleRound0(rs *cstypes.RoundState) {
	_ = "STUB: not implemented"
	// cs.Logger.Info("scheduleRound0", "now", cmttime.Now(), "startTime", cs.StartTime)
	return
}

// Attempt to schedule a timeout (by sending timeoutInfo on the tickChan)
func (cs *State) scheduleTimeout(duration time.Duration, height int64, round int32, step cstypes.RoundStepType) {
	_ = "STUB: not implemented"
	return
}

// Propose returns the amount of time to wait for a proposal, using application timeouts
// and falling back to config timeouts if application timeouts are zero
func (cs *State) Propose(round int32) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// Fallback to config values if state timeouts are zero

// Prevote returns the amount of time to wait for straggler votes after receiving any +2/3 prevotes,
// using application timeouts and falling back to config timeouts if application timeouts are zero
func (cs *State) Prevote(round int32) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// Fallback to config values if state timeouts are zero

// Precommit returns the amount of time to wait for straggler votes after receiving any +2/3 precommits,
// using application timeouts and falling back to config timeouts if application timeouts are zero
func (cs *State) Precommit(round int32) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// Fallback to config values if state timeouts are zero

// send a msg into the receiveRoutine regarding our own proposal, block part, or vote
func (cs *State) sendInternalMessage(mi msgInfo) { _ = "STUB: not implemented"; return }

// NOTE: using the go-routine means our votes can
// be processed out of order.
// TODO: use CList here for strict determinism and
// attempt push to internalMsgQueue in receiveRoutine

// ReconstructSeenCommit reconstructs the seen commit
// This function is meant to be called after statesync
// that was performed offline as to avoid interfering with vote
// extensions.
func (cs *State) reconstructSeenCommit(state sm.State) { _ = "STUB: not implemented"; return }

// Reconstruct the LastCommit from either SeenCommit or the ExtendedCommit. SeenCommit
// and ExtendedCommit are saved along with the block. If VoteExtensions are required
// the method will panic on an absent ExtendedCommit or an ExtendedCommit without
// extension data.
func (cs *State) reconstructLastCommit(state sm.State) { _ = "STUB: not implemented"; return }

func (cs *State) votesFromExtendedCommit(state sm.State) (*types.VoteSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cs *State) votesFromSeenCommit(state sm.State) (*types.VoteSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates State and increments height to match that of state.
// The round becomes 0 and cs.rs.Step becomes cstypes.RoundStepNewHeight.
func (cs *State) updateToState(state sm.State) { _ = "STUB: not implemented"; return }

// This might happen when someone else is mutating cs.state.
// Someone forgot to pass in state.Copy() somewhere?!

// If state isn't further out than cs.state, just ignore.
// This happens when SwitchToConsensus() is called in the reactor.
// We don't want to reset e.g. the Votes, but we still want to
// signal the new round step, because other services (eg. txNotifier)
// depend on having an up-to-date peer state!

// Reset fields based on state.

// Very first commit should be empty.

// Otherwise, use cs.rs.Votes

// NOTE: when consensus starts, it has no votes. reconstructLastCommit
// must be called to reconstruct LastCommit from SeenCommit.

// Next desired block height

// RoundState fields

// "Now" makes it easier to sync up dev nodes.
// We add timeoutCommit to allow transactions
// to be gathered for the first block.
// And alternative solution that relies on clocks:
// cs.StartTime = state.LastBlockTime.Add(timeoutCommit)

// Don't use cs.state.Timeouts.TimeoutCommit because that is zero

// Finally, broadcast RoundState

func (cs *State) newStep() { _ = "STUB: not implemented"; return }

// newStep is called by updateToState in NewState before the eventBus is set!

//-----------------------------------------
// the main go routines

// receiveRoutine handles messages which may cause state transitions.
// it's argument (n) is the number of messages to process before exiting - use 0 to run forever
// It keeps the RoundState and is the only thing that updates it.
// Updates (state transitions) happen on timeouts, complete proposals, and 2/3 majorities.
// State must be locked before any internal state is updated.
func (cs *State) receiveRoutine(maxSteps int) { _ = "STUB: not implemented"; return }

// NOTE: the internalMsgQueue may have signed messages from our
// priv_val that haven't hit the WAL, but its ok because
// priv_val tracks LastSig

// close wal now that we're done writing to it

// stop gracefully
//
// NOTE: We most probably shouldn't be running any further when there is
// some unexpected panic. Some unknown error happened, and so we don't
// know if that will result in the validator signing an invalid thing. It
// might be worthwhile to explore a mechanism for manual resuming via
// some console or secure RPC system, but for now, halting the chain upon
// unexpected consensus bugs sounds like the better option.

// handles proposals, block parts, votes
// may generate internal events (votes, complete proposals, 2/3 majorities)

// NOTE: fsync

// we actually want to simulate failing during
// the previous WriteSync, but this isn't easy to do.
// Equivalent would be to fail here and manually remove
// some bytes from the end of the wal.
// XXX

// handles proposals, block parts, votes

// tockChan:

// if the timeout is relevant to the rs
// go to the next step

// state transitions on complete-proposal, 2/3-any, 2/3-one
func (cs *State) handleMsg(mi msgInfo) { _ = "STUB: not implemented"; return }

// will not cause transition.
// once proposal is set, we can receive block parts

// if the proposal is complete, we'll enterPrevote or tryFinalizeCommit

// We unlock here to yield to any routines that need to read the the RoundState.
// Previously, this code held the lock from the point at which the final block
// part was received until the block executed against the application.
// This prevented the reactor from being able to retrieve the most updated
// version of the RoundState. The reactor needs the updated RoundState to
// gossip the now completed block.
//
// This code can be further improved by either always operating on a copy
// of RoundState and only locking when switching out State's copy of
// RoundState with the updated copy or by emitting RoundState events in
// more places for routines depending on it to listen for.

// attempt to add the vote and dupeout the validator if its a duplicate signature
// if the vote gives us a 2/3-any or 2/3-one, we transition

// if err == ErrAddingVote {
// TODO: punish peer
// We probably don't want to stop the peer here. The vote does not
// necessarily comes from a malicious peer but can be just broadcasted by
// a typical peer.
// https://github.com/tendermint/tendermint/issues/1281
// }

// NOTE: the vote is broadcast to peers by the reactor listening
// for vote events

// TODO: If rs.Height == vote.Height && rs.Round < vote.Round,
// the peer is sending us CatchupCommit precommits.
// We could make note of this and help filter in broadcastHasVoteMessage().

func (cs *State) handleTimeout(ti timeoutInfo, rs cstypes.RoundState) {
	_ = "STUB: not implemented"
	return
}

// timeouts must be for current height, round, step

// the timeout will now cause a state transition

// NewRound event fired from enterNewRound.
// XXX: should we fire timeout here (for timeout commit)?

func (cs *State) handleTxsAvailable() { _ = "STUB: not implemented"; return }

// We only need to do this for round 0.

// timeoutCommit phase

// enterPropose will be called by enterNewRound

// +1ms to ensure RoundStepNewRound timeout always happens after RoundStepNewHeight

// after timeoutCommit

//-----------------------------------------------------------------------------
// State functions
// Used internally by handleTimeout and handleMsg to make state transitions

// Enter: `timeoutNewHeight` by startTime (commitTime+timeoutCommit),
//
//	or, if SkipTimeoutCommit==true, after receiving all precommits from (height,round-1)
//
// Enter: `timeoutPrecommits` after any +2/3 precommits from (height,round-1)
// Enter: +2/3 precommits for nil at (height,round-1)
// Enter: +2/3 prevotes any or +2/3 precommits for block or any from (height, round)
// NOTE: cs.StartTime was already set for height.
func (cs *State) enterNewRound(height int64, round int32) { _ = "STUB: not implemented"; return }

// If moving to a new round (not round 0), check if previous proposer missed
// Only count as miss if proposal was never received (Proposal == nil)

// increment validators if necessary

// Setup new round
// we don't fire newStep for this step,
// but we fire an event, so update the round step first

// If round == 0, we've already reset these upon new height, and meanwhile
// we might have received a proposal for round 0.

// also track next round (round+1) to allow round-skipping

// Wait for txs to be available in the mempool
// before we enterPropose in round 0. If the last block changed the app hash,
// we may need an empty "proof" block, and enterPropose immediately.

// needProofBlock returns true on the first height (so the genesis app hash is signed right away)
// and where the last block (height-1) caused the app hash to change
func (cs *State) needProofBlock(height int64) bool { _ = "STUB: not implemented"; return false }

// See https://github.com/cometbft/cometbft/issues/370

// Enter (CreateEmptyBlocks): from enterNewRound(height,round)
// Enter (CreateEmptyBlocks, CreateEmptyBlocksInterval > 0 ):
//
//	after enterNewRound(height,round), after timeout of CreateEmptyBlocksInterval
//
// Enter (!CreateEmptyBlocks) : after enterNewRound(height,round), once txs are in the mempool
func (cs *State) enterPropose(height int64, round int32) { _ = "STUB: not implemented"; return }

// Done enterPropose:

// If we have the whole proposal + POL, then goto Prevote now.
// else, we'll enterPrevote when the rest of the proposal is received (in AddProposalBlockPart),
// or else after timeoutPropose

// If we don't get the proposal and all block parts quick enough, enterPrevote

// Nothing more to do if we're not a validator

// If this node is a validator & proposer in the current round, it will
// miss the opportunity to create a block.

// if not a validator, we're done

func (cs *State) isProposer(address []byte) bool { _ = "STUB: not implemented"; return false }

func (cs *State) defaultDecideProposal(height int64, round int32) {
	_ = "STUB: not implemented"
	return
}

// Decide on block

// If there is valid block, choose that.

// set the recovery related fields if using an existing block

// Create a new proposal block from state/txs from the mempool.

// Flush the WAL. Otherwise, we may not recompute the same proposal to sign,
// and the privValidator will refuse to sign anything.

// Make proposal

// Verify whether this proposal corresponds to the returned signature.
// This fixes the edge case where a KMS, in a sentry setup, returns the signature of a different proposal

// send proposal and block parts on internal msg queue

// Returns true if the proposal block is complete &&
// (if POLRound was proposed, we have +2/3 prevotes from there).
func (cs *State) isProposalComplete() bool { _ = "STUB: not implemented"; return false }

// we have the proposal. if there's a POLRound,
// make sure we have the prevotes from it too

// if this is false the proposer is lying or we haven't received the POL yet

// Create the next block to propose and return it. Returns nil block upon error.
//
// We really only need to return the parts, but the block is returned for
// convenience so we can log the proposal block.
//
// NOTE: keep it side-effect free for clarity.
// CONTRACT: cs.privValidator is not nil.
func (cs *State) createProposalBlock(ctx context.Context) (block *types.Block, blockParts *types.PartSet, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// TODO(sergio): wouldn't it be easier if CreateProposalBlock accepted cs.LastCommit directly?

// We're creating a proposal for the first block.
// The commit is empty, but not nil.

// Make the commit from LastCommit

// This shouldn't happen.

// If this node is a validator & proposer in the current round, it will
// miss the opportunity to create a block.

// Enter: `timeoutPropose` after entering Propose.
// Enter: proposal block and POL is ready.
// Prevote for LockedBlock if we're locked, or ProposalBlock if valid.
// Otherwise vote nil.
func (cs *State) enterPrevote(height int64, round int32) { _ = "STUB: not implemented"; return }

// Done enterPrevote:

// Sign and broadcast vote as necessary

// Once `addVote` hits any +2/3 prevotes, we will go to PrevoteWait
// (so we have more time to try and collect +2/3 prevotes for a single block)

func (cs *State) defaultDoPrevote(height int64, round int32) { _ = "STUB: not implemented"; return }

// If a block is locked, prevote that.

// If ProposalBlock is nil, prevote nil.

// Validate proposal block, from consensus' perspective

// ProposalBlock is invalid, prevote nil.

/*
	Before prevoting on the block received from the proposer for the current round and height,
	we request the Application, via `ProcessProposal` ABCI call, to confirm that the block is
	valid. If the Application does not accept the block, consensus prevotes `nil`.

	WARNING: misuse of block rejection by the Application can seriously compromise
	the liveness properties of consensus.
	Please see `PrepareProosal`-`ProcessProposal` coherence and determinism properties
	in the ABCI++ specification.
*/

// Vote nil if the Application rejected the block

// save block as a json

// Prevote cs.rs.ProposalBlock
// NOTE: the proposal signature is validated when it is received,
// and the proposal block parts are validated as they are received (against the merkle hash in the proposal)

// Enter: any +2/3 prevotes at next round.
func (cs *State) enterPrevoteWait(height int64, round int32) { _ = "STUB: not implemented"; return }

// Done enterPrevoteWait:

// Wait for some more prevotes; enterPrecommit

// Enter: `timeoutPrevote` after any +2/3 prevotes.
// Enter: `timeoutPrecommit` after any +2/3 precommits.
// Enter: +2/3 precomits for block or nil.
// Lock & precommit the ProposalBlock if we have enough prevotes for it (a POL in this round)
// else, unlock an existing lock and precommit nil if +2/3 of prevotes were nil,
// else, precommit nil otherwise.
func (cs *State) enterPrecommit(height int64, round int32) { _ = "STUB: not implemented"; return }

// if any other routine tries to enter precommit, we just return

// Done enterPrecommit:

// check for a polka

// If we don't have a polka, we must precommit nil.

// At this point +2/3 prevoted for a particular block or nil.

// the latest POLRound should be this round.

// +2/3 prevoted nil. Unlock and precommit nil.

// At this point, +2/3 prevoted for a particular block.

// If we're already locked on that block, precommit it, and update the LockedRound

// If +2/3 prevoted for proposal block, stage and precommit it

// Validate the block.

// There was a polka in this round for a block we don't have.
// Fetch that block, unlock, and precommit nil.
// The +2/3 prevotes for this round is the POL for our unlock.

// Enter: any +2/3 precommits for next round.
func (cs *State) enterPrecommitWait(height int64, round int32) { _ = "STUB: not implemented"; return }

// Done enterPrecommitWait:

// wait for some more precommits; enterNewRound

// Enter: +2/3 precommits for block
func (cs *State) enterCommit(height int64, commitRound int32) { _ = "STUB: not implemented"; return }

// Done enterCommit:
// keep cs.rs.Round the same, commitRound points to the right Precommits set.

// Maybe finalize immediately.

// The Locked* fields no longer matter.
// Move them over to ProposalBlock if they match the commit hash,
// otherwise they'll be cleared in updateToState.

// If we don't have the block being committed, set up to get it.

// We're getting the wrong block.
// Set up ProposalBlockParts and keep waiting.

// If we have the block AND +2/3 commits for it, finalize.
func (cs *State) tryFinalizeCommit(height int64) { _ = "STUB: not implemented"; return }

// TODO: this happens every time if we're not a validator (ugly logs)
// TODO: ^^ wait, why does it matter that we're a validator?

// Increment height and goto cstypes.RoundStepNewHeight
func (cs *State) finalizeCommit(height int64) { _ = "STUB: not implemented"; return }

// XXX

// Save to blockStore.

// NOTE: the seenCommit is local justification to commit this block,
// but may differ from the LastCommit included in the next block

// Happens during replay if we already saved the block but didn't commit

// XXX

// Write EndHeightMessage{} for this height, implying that the blockstore
// has saved the block.
//
// If we crash before writing this EndHeightMessage{}, we will recover by
// running ApplyBlock during the ABCI handshake when we restart.  If we
// didn't save the block to the blockstore before writing
// EndHeightMessage{}, we'd have to change WAL replay -- currently it
// complains about replaying for heights where an #ENDHEIGHT entry already
// exists.
//
// Either way, the State should not be resumed until we
// successfully call ApplyBlock (ie. later here, or in Handshake after
// restart).

// NOTE: fsync

// XXX

// Create a copy of the state for staging and an event cache for txs.

// Execute and commit the block, update and save the state, and update the mempool.
// We use apply verified block here because we have verified the block in this function already.
// NOTE The block.AppHash won't reflect these txs until the next block.

// XXX

// must be called before we update state

// NewHeightStep!

// XXX

// Private validator might have changed it's key pair => refetch pubkey.

// prune the propagation reactor

// cs.StartTime is already set.
// Schedule Round0 to start soon.

// By here,
// * cs.rs.Height has been increment to height+1
// * cs.rs.Step is now cstypes.RoundStepNewHeight
// * cs.StartTime is set to when we will start round0.

func (cs *State) recordMetrics(height int64, block *types.Block) { _ = "STUB: not implemented"; return }

// height=0 -> MissingValidators and MissingValidatorsPower are both 0.
// Remember that the first LastCommit is intentionally empty, so it's not
// fair to increment missing validators number.

// Sanity check that commit size matches validator set size - only applies
// after first block.

// Metrics won't be updated, but it's not critical.

// NOTE: byzantine validators power and count is only for consensus evidence i.e. duplicate vote

// trace some metadata about the block

//nolint:staticcheck
//nolint:staticcheck

// KMSSigningDelay is a constant representing a delay used primarily to adjust for KMS signing latencies.
const KMSSigningDelay = 200 * time.Millisecond

// precommitDelay calculates if the process has waited at least a certain number of seconds
// from their start time before they can vote.
// If the application's DelayedPrecommitTimeout is set to 0, no precommit wait is done.
// When catching up on block data, returns 0 to speed up block processing.
func (cs *State) precommitDelay() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// setting 0 as a special case not to reschedule the pre-commit

// When catching up, skip the delayed precommit timeout to speed up block processing.

//-----------------------------------------------------------------------------

func (cs *State) defaultSetProposal(proposal *types.Proposal) error {
	_ = "STUB: not implemented"
	// Already have one
	// TODO: possibly catch double proposals
	return nil
}

// Does not apply

// Verify POLRound, which must be -1 or in range [0, proposal.Round).

// Verify signature

// Validate the proposed block size, derived from its PartSetHeader

// We don't update cs.rs.ProposalBlockParts if it is already set.
// This happens if we're already in cstypes.RoundStepCommit or if there is a valid block in the current round.
// TODO: We can check if Proposal is for a different block as this is a sign of misbehavior!

// NOTE: block is not necessarily valid.
// Asynchronously triggers either enterPrevote (before we timeout of propose) or tryFinalizeCommit,
// once we have the full block.
func (cs *State) addProposalBlockPart(msg *BlockPartMessage, peerID p2p.ID) (added bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Blocks might be reused, so round mismatch is OK

// We're not expecting a block part.

// NOTE: this can happen when we've gone to a higher round and
// then receive parts from the previous round - not necessarily a bad peer.

// NOTE: we are disregarding possible duplicates above where heights dont match or we're not expecting block parts yet
// but between the matches_current = true and false, we have all the info.

// NOTE: it's possible to receive complete proposal blocks for future rounds without having the proposal

func (cs *State) handleCompleteProposal(blockHeight int64) {
	_ = "STUB: not implemented"
	// Update Valid* if we can.
	return
}

// TODO: In case there is +2/3 majority in Prevotes set for some
// block and cs.rs.ProposalBlock contains different block, either
// proposer is faulty or voting power of faulty processes is more
// than 1/3. We should trigger in the future accountability
// procedure at this point.

// Move onto the next step

// this is optimisation as this will be triggered when prevote is added

// If we're waiting on the proposal block...

// Attempt to add the vote. if its a duplicate signature, dupeout the validator
func (cs *State) tryAddVote(vote *types.Vote, peerID p2p.ID) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// NOTE: some of these errors are swallowed here

// If the vote height is off, we'll just ignore it,
// But if it's a conflicting sig, add it to the cs.evpool.
// If it's otherwise invalid, punish peer.
//nolint: gocritic

// report conflicting votes to the evidence pool

// Either
// 1) bad peer OR
// 2) not a bad peer? this can also err sometimes with "Unexpected step" OR
// 3) tmkms use with multiple validators connecting to a single tmkms instance
// 		(https://github.com/tendermint/tendermint/issues/3839).

func (cs *State) addVote(vote *types.Vote, peerID p2p.ID) (added bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

// A precommit for the previous height?
// These come in while we wait timeoutCommit

// Late precommit at prior height is ignored

// If the vote wasnt added but there's no error, its a duplicate vote

// if we can skip timeoutCommit and have all the votes now,

// go straight to new round (skip timeout commit)
// cs.scheduleTimeout(time.Duration(0), cs.rs.Height, 0, cstypes.RoundStepNewHeight)

// Height mismatch is ignored.
// Not necessarily a bad peer, but not favorable behavior.

// Check to see if the chain is configured to extend votes.

// The chain is configured to extend votes, check that the vote is
// not for a nil block and verify the extensions signature against the
// corresponding public key.

// Verify VoteExtension if precommit and not nil
// https://github.com/tendermint/tendermint/issues/8487

// Skip the VerifyVoteExtension call if the vote was issued by this validator.

// The core fields of the vote message were already validated in the
// consensus reactor when the vote was received.
// Here, we verify the signature of the vote extension included in the vote
// message.

// TODO: we should disconnect from this malicious peer

// Vote extensions are not enabled on the network.
// Reject the vote, as it is malformed
//
// TODO punish a peer if it sent a vote with an extension when the feature
// is disabled on the network.
// https://github.com/tendermint/tendermint/issues/8565

// Either duplicate, or error upon cs.rs.Votes.AddByIndex()

// If the vote wasnt added but there's no error, its a duplicate vote

// If +2/3 prevotes for a block or nil for *any* round:

// There was a polka!
// If we're locked but this is a recent polka, unlock.
// If it matches our ProposalBlock, update the ValidBlock

// Unlock if `cs.rs.LockedRound < vote.Round <= cs.rs.Round`
// NOTE: If vote.Round > cs.rs.Round, we'll deal with it when we get to vote.Round

// Update Valid* if we can.
// NOTE: our proposal block may be nil or not what received a polka..

// we're getting the wrong block

// If +2/3 prevotes for *anything* for future round:

// Round-skip if there is any 2/3+ of votes ahead of us

// current round

// If the proposal is now complete, enter prevote of cs.rs.Round.

// Executed as TwoThirdsMajority could be from a higher round

// CONTRACT: cs.privValidator is not nil.
func (cs *State) signVote(
	msgType cmtproto.SignedMsgType,
	hash []byte,
	header types.PartSetHeader,
	block *types.Block,
) (*types.Vote, error) {
	_ = "STUB: not implemented"
	// Flush the WAL. Otherwise, we may not recompute the same vote to sign,
	// and the privValidator will refuse to sign anything.
	return nil, nil
}

// if the signedMessage type is for a non-nil precommit, add
// VoteExtension

func (cs *State) voteTime() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// Minimum time increment between blocks

// TODO: We should remove next line in case we don't vote for v in case cs.rs.ProposalBlock == nil,
// even if cs.rs.LockedBlock != nil. See https://github.com/cometbft/cometbft/tree/v0.38.x/spec/.

// See the BFT time spec
// https://github.com/cometbft/cometbft/blob/v0.38.x/spec/consensus/bft-time.md

// sign the vote and publish on internalMsgQueue
// block information is only used to extend votes (precommit only); should be nil in all other cases
func (cs *State) signAddVote(
	msgType cmtproto.SignedMsgType,
	hash []byte,
	header types.PartSetHeader,
	block *types.Block,
) {
	_ = "STUB: not implemented"
	return
}

// the node does not have a key

// Vote won't be signed, but it's not critical.

// If the node not in the validator set, do nothing.

// TODO: pass pubKey to signVote

// updatePrivValidatorPubKey get's the private validator public key and
// memoizes it. This func returns an error if the private validator is not
// responding or responds with an error.
func (cs *State) updatePrivValidatorPubKey() error { _ = "STUB: not implemented"; return nil }

// look back to check existence of the node's consensus votes before joining consensus
func (cs *State) checkDoubleSigningRisk(height int64) error { _ = "STUB: not implemented"; return nil }

func (cs *State) calculatePrevoteMessageDelayMetrics() { _ = "STUB: not implemented"; return }

//---------------------------------------------------------

func CompareHRS(h1 int64, r1 int32, s1 cstypes.RoundStepType, h2 int64, r2 int32, s2 cstypes.RoundStepType) int {
	_ = "STUB: not implemented"
	return 0
}

// repairWalFile decodes messages from src (until the decoder errors) and
// writes them to dst.
func repairWalFile(src, dst string) error { _ = "STUB: not implemented"; return nil }

// best-case repair (until first error is encountered)

// syncData continuously listens for and processes block parts or proposals from the propagation reactor.
// It stops execution when the service is terminated or the channels are closed.
func (cs *State) syncData() { _ = "STUB: not implemented"; return }

func (cs *State) lockAll() { _ = "STUB: not implemented"; return }

func (cs *State) unlockAll() { _ = "STUB: not implemented"; return }
