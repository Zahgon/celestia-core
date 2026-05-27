package propagation

import (
	"sync/atomic"

	proptypes "github.com/cometbft/cometbft/consensus/propagation/types"
	"github.com/cometbft/cometbft/libs/bits"
	"github.com/cometbft/cometbft/libs/sync"
	"github.com/cometbft/cometbft/store"
	"github.com/cometbft/cometbft/types"
)

type proposalData struct {
	compactBlock *proptypes.CompactBlock
	block        *proptypes.CombinedPartSet
	maxRequests  *bits.BitArray
	catchup      bool
}

type ProposalCache struct {
	store     *store.BlockStore
	pmtx      *sync.Mutex
	proposals map[int64]map[int32]*proposalData

	// height the height we're trying to get consensus on.
	// the last committed height is height-1.
	height int64
	round  int32

	// currentProposalPartsCount is the maximum number of concurrent requests allowed for the current proposal.
	currentProposalPartsCount atomic.Int64
}

func NewProposalCache(bs *store.BlockStore) *ProposalCache { _ = "STUB: not implemented"; return nil }

// if there is a block saved in the store, set the current height and round.

// getCurrentProposalPartsCount returns the current proposal number of parts.
func (p *ProposalCache) getCurrentProposalPartsCount() int64 { _ = "STUB: not implemented"; return 0 }

// setCurrentProposalPartsCount sets the current proposal number of parts.
func (p *ProposalCache) setCurrentProposalPartsCount(limit int64) {
	_ = "STUB: not implemented"
	return
}

func (p *ProposalCache) AddProposal(cb *proptypes.CompactBlock) (added bool) {
	_ = "STUB: not implemented"
	return false
}

// GetProposal returns the proposal and block for a given height and round if
// this node has it stored or cached.
func (p *ProposalCache) GetProposal(height int64, round int32) (*types.Proposal, *types.PartSet, bool) {
	_ = "STUB: not implemented"
	return nil, nil, false
}

func (p *ProposalCache) unfinishedHeights() []*proposalData { _ = "STUB: not implemented"; return nil }

// We only consider heights strictly higher than the committed height as
// "in flight". Heights that are already committed are ignored even if they
// remain cached.

// Treat catchup proposals as "unfinished" even if all parts are present.
// This lets the consensus layer know it should skip delayed precommit when
// replaying cached proposals to catch up.

// relevant determines if a height or round is currently actionable. For
// example, passing the height that was already committed is not actionable.
// Passing a round that has already been surpassed is not actionable.
func (p *ProposalCache) relevant(height int64, round int32) bool {
	_ = "STUB: not implemented"
	return false
}

// safeRelevant determines if a have message is relevant in a thread safe way.
func (p *ProposalCache) safeRelevant(height int64, round int32) bool {
	_ = "STUB: not implemented"
	return false
}

// GetProposal returns the proposal and block for a given height and round if
// this node has it stored or cached. It also return the max requests for that
// block.
func (p *ProposalCache) getAllState(height int64, round int32, catchup bool) (*proptypes.CompactBlock, *proptypes.CombinedPartSet, *bits.BitArray, bool) {
	_ = "STUB: not implemented"
	return nil, nil, nil, false
}

// if the round is less than -1, then they're asking for the latest
// proposal

// get the latest round

// GetCurrentProposal returns the current proposal and block for the current
// height and round.
func (p *ProposalCache) GetCurrentProposal() (*types.Proposal, *proptypes.CombinedPartSet, bool) {
	_ = "STUB: not implemented"
	return nil, nil, false
}

// GetCurrentCompactBlock returns the current compact block for the current
// height and round.
func (p *ProposalCache) GetCurrentCompactBlock() (*proptypes.CompactBlock, *proptypes.CombinedPartSet, bool) {
	_ = "STUB: not implemented"
	return nil, nil, false
}

func (p *ProposalCache) DeleteHeight(height int64) { _ = "STUB: not implemented"; return }

func (p *ProposalCache) DeleteRound(height int64, round int32) { _ = "STUB: not implemented"; return }

// prune deletes all cached compact blocks for heights less than the provided
// height and round.
//
// todo: also prune rounds. this requires prune in the consensus reactor after
// moving rounds.
func (p *ProposalCache) prune(pruneHeight int64) { _ = "STUB: not implemented"; return }
