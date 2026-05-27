package evidence

import (
	"sync"
	"time"

	dbm "github.com/cometbft/cometbft-db"

	clist "github.com/cometbft/cometbft/libs/clist"
	"github.com/cometbft/cometbft/libs/log"
	sm "github.com/cometbft/cometbft/state"
	"github.com/cometbft/cometbft/types"
)

const (
	baseKeyCommitted = byte(0x00)
	baseKeyPending   = byte(0x01)
)

// Pool maintains a pool of valid evidence to be broadcasted and committed
type Pool struct {
	logger log.Logger

	evidenceStore dbm.DB
	evidenceList  *clist.CList // concurrent linked-list of evidence
	evidenceSize  uint32       // amount of pending evidence

	// needed to load validators to verify evidence
	stateDB sm.Store
	// needed to load headers and commits to verify evidence
	blockStore BlockStore

	mtx sync.Mutex
	// latest state
	state sm.State
	// evidence from consensus is buffered to this slice, awaiting until the next height
	// before being flushed to the pool. This prevents broadcasting and proposing of
	// evidence before the height with which the evidence happened is finished.
	consensusBuffer []duplicateVoteSet

	pruningHeight int64
	pruningTime   time.Time
}

// NewPool creates an evidence pool. If using an existing evidence store,
// it will add all pending evidence to the concurrent list.
func NewPool(evidenceDB dbm.DB, stateDB sm.Store, blockStore BlockStore) (*Pool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if pending evidence already in db, in event of prior failure, then check for expiration,
// update the size and load it back to the evidenceList

// PendingEvidence is used primarily as part of block proposal and returns up to maxNum of uncommitted evidence.
func (evpool *Pool) PendingEvidence(maxBytes int64) ([]types.Evidence, int64) {
	_ = "STUB: not implemented"
	return nil, 0
}

// Update takes both the new state and the evidence committed at that height and performs
// the following operations:
//  1. Take any conflicting votes from consensus and use the state's LastBlockTime to form
//     DuplicateVoteEvidence and add it to the pool.
//  2. Update the pool's state which contains evidence params relating to expiry.
//  3. Moves pending evidence that has now been committed into the committed pool.
//  4. Removes any expired evidence based on both height and time.
func (evpool *Pool) Update(state sm.State, ev types.EvidenceList) {
	_ = "STUB: not implemented"
	// sanity check
	return
}

// flush conflicting vote pairs from the buffer, producing DuplicateVoteEvidence and
// adding it to the pool

// update state

// move committed evidence out from the pending pool and into the committed pool

// prune pending evidence when it has expired. This also updates when the next evidence will expire

// AddEvidence checks the evidence is valid and adds it to the pool.
func (evpool *Pool) AddEvidence(ev types.Evidence) error { _ = "STUB: not implemented"; return nil }

// We have already verified this piece of evidence - no need to do it again

// check that the evidence isn't already committed

// this can happen if the peer that sent us the evidence is behind so we shouldn't
// punish the peer.

// 1) Verify against state.

// 2) Save to store.

// 3) Add evidence to clist.

// ReportConflictingVotes takes two conflicting votes and forms duplicate vote evidence,
// adding it eventually to the evidence pool.
//
// Duplicate vote attacks happen before the block is committed and the timestamp is
// finalized, thus the evidence pool holds these votes in a buffer, forming the
// evidence from them once consensus at that height has been reached and `Update()` with
// the new state called.
//
// Votes are not verified.
func (evpool *Pool) ReportConflictingVotes(voteA, voteB *types.Vote) {
	_ = "STUB: not implemented"
	return
}

// CheckEvidence takes an array of evidence from a block and verifies all the evidence there.
// If it has already verified the evidence then it jumps to the next one. It ensures that no
// evidence has already been committed or is being proposed twice. It also adds any
// evidence that it doesn't currently have so that it can quickly form ABCI Evidence later.
func (evpool *Pool) CheckEvidence(evList types.EvidenceList) error {
	_ = "STUB: not implemented"
	return nil
}

// We must verify light client attack evidence regardless because there could be a
// different conflicting block with the same hash.

// check that the evidence isn't already committed

// Something went wrong with adding the evidence but we already know it is valid
// hence we log an error and continue

// check for duplicate evidence. We cache hashes so we don't have to work them out again.

// EvidenceFront goes to the first evidence in the clist
func (evpool *Pool) EvidenceFront() *clist.CElement { _ = "STUB: not implemented"; return nil }

// EvidenceWaitChan is a channel that closes once the first evidence in the list is there. i.e Front is not nil
func (evpool *Pool) EvidenceWaitChan() <-chan struct{} { _ = "STUB: not implemented"; return nil }

// SetLogger sets the Logger.
func (evpool *Pool) SetLogger(l log.Logger) {
	_ = "STUB: not implemented"

	// Size returns the number of evidence in the pool.
	return
}

func (evpool *Pool) Size() uint32 { _ = "STUB: not implemented"; return 0 }

// State returns the current state of the evpool.
func (evpool *Pool) State() sm.State { _ = "STUB: not implemented"; return *new(sm.State) }

func (evpool *Pool) Close() error { _ = "STUB: not implemented"; return nil }

// IsExpired checks whether evidence or a polc is expired by checking whether a height and time is older
// than set by the evidence consensus parameters
func (evpool *Pool) isExpired(height int64, time time.Time) bool {
	_ = "STUB: not implemented"
	return false
}

// IsCommitted returns true if we have already seen this exact evidence and it is already marked as committed.
func (evpool *Pool) isCommitted(evidence types.Evidence) bool {
	_ = "STUB: not implemented"
	return false
}

// IsPending checks whether the evidence is already pending. DB errors are passed to the logger.
func (evpool *Pool) isPending(evidence types.Evidence) bool {
	_ = "STUB: not implemented"
	return false
}

func (evpool *Pool) addPendingEvidence(ev types.Evidence) error {
	_ = "STUB: not implemented"
	return nil
}

func (evpool *Pool) removePendingEvidence(evidence types.Evidence) {
	_ = "STUB: not implemented"
	return
}

// markEvidenceAsCommitted processes all the evidence in the block, marking it as
// committed and removing it from the pending database.
func (evpool *Pool) markEvidenceAsCommitted(evidence types.EvidenceList) {
	_ = "STUB: not implemented"
	return
}

// Add evidence to the committed list. As the evidence is stored in the block store
// we only need to record the height that it was saved at.

// remove committed evidence from the clist

// listEvidence retrieves lists evidence from oldest to newest within maxBytes.
// If maxBytes is -1, there's no cap on the size of returned evidence.
func (evpool *Pool) listEvidence(prefixKey byte, maxBytes int64) ([]types.Evidence, int64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// used for calculating the bytes size

func (evpool *Pool) removeExpiredPendingEvidence() (int64, time.Time) {
	_ = "STUB: not implemented"
	return 0, *new(time.Time)
}

// return the height and time with which this evidence will have expired so we know when to prune next

// We either have no pending evidence or all evidence has expired

func (evpool *Pool) removeEvidenceFromList(
	blockEvidenceMap map[string]struct{},
) {
	_ = "STUB: not implemented"
	return
}

// Remove from clist

func (evpool *Pool) updateState(state sm.State) { _ = "STUB: not implemented"; return }

// processConsensusBuffer converts all the duplicate votes witnessed from consensus
// into DuplicateVoteEvidence. It sets the evidence timestamp to the block height
// from the most recently committed block.
// Evidence is then added to the pool so as to be ready to be broadcasted and proposed.
func (evpool *Pool) processConsensusBuffer(state sm.State) { _ = "STUB: not implemented"; return }

// Check the height of the conflicting votes and fetch the corresponding time and validator set
// to produce the valid evidence

// evidence pool shouldn't expect to get votes from consensus of a height that is above the current
// state. If this error is seen then perhaps consider keeping the votes in the buffer and retry
// in following heights

// check if we already have this evidence

// check that the evidence is not already committed on chain

// reset consensus buffer

type duplicateVoteSet struct {
	VoteA *types.Vote
	VoteB *types.Vote
}

func bytesToEv(evBytes []byte) (types.Evidence, error) {
	_ = "STUB: not implemented"
	return *new(types.Evidence), nil
}

func evMapKey(ev types.Evidence) string { _ = "STUB: not implemented"; return "" }

// big endian padded hex
func bE(h int64) string { _ = "STUB: not implemented"; return "" }

func keyCommitted(evidence types.Evidence) []byte { _ = "STUB: not implemented"; return nil }

func keyPending(evidence types.Evidence) []byte { _ = "STUB: not implemented"; return nil }

func keySuffix(evidence types.Evidence) []byte { _ = "STUB: not implemented"; return nil }
