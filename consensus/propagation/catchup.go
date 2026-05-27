package propagation

import (
	proptypes "github.com/cometbft/cometbft/consensus/propagation/types"
	"github.com/cometbft/cometbft/types"
)

// retryWants ensure that all data for all unpruned compact blocks is requested.
func (blockProp *Reactor) retryWants() { _ = "STUB: not implemented"; return }

// only re-request original parts that are missing, not parity parts.

// make requests from different peers

// subtract the parts we just requested

// keep track of which requests we've made this attempt.

func (blockProp *Reactor) AddCommitment(height int64, round int32, psh *types.PartSetHeader) {
	_ = "STUB: not implemented"
	return
}

// this assumes that the parity parts are the same size

// increment the local copies of the height and round

func shuffle[T any](slice []T) []T { _ = "STUB: not implemented"; return nil }

// applyCachedProposalIfAvailable checks for cached proposals at the current height/round
// and applies the first valid one. Called automatically after SetProposer or SetHeightAndRound
// to enable fast catchup when a node falls behind.
//
// This function iterates through ALL peers' cached proposals for the current height/round,
// trying each one until it finds a valid proposal. This ensures a single invalid proposal
// from one peer doesn't block valid proposals from other peers.
func (blockProp *Reactor) applyCachedProposalIfAvailable() { _ = "STUB: not implemented"; return }

// Check if we already have a proposal for this height/round (normal case)

// Already have proposal, no need to check cache

// Iterate through all peers looking for a valid cached proposal

// This peer has no cached proposal for this height

// Skip proposals for different rounds - they'll be tried when we advance

// Try to validate this proposal

// Try next peer's cached proposal

// Found a valid proposal - apply it

// Clean up the cache entry for this peer only if we successfully applied it.

// handleCachedCompactBlock processes a verified cached compact block.
// Similar to handleCompactBlock but skips validation (already verified) and triggers immediate catchup.
// Returns true if the cached block was applied.
func (blockProp *Reactor) handleCachedCompactBlock(cb *proptypes.CompactBlock) bool {
	_ = "STUB: not implemented"
	return false
}

// generate (and cache) the proofs from the partset hashes in the compact block

// Send proposal to consensus reactor

// From self since it's from cache

// Add to proposal cache

// Mark as catchup to skip parity requests in retryWants

// Recover any parts from mempool

// Immediately trigger part requests (like AddCommitment)
