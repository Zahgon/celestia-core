package types

import (
	"encoding/json"
	"time"

	"github.com/cometbft/cometbft/libs/bytes"
	"github.com/cometbft/cometbft/types"
)

//-----------------------------------------------------------------------------
// RoundStepType enum type

// RoundStepType enumerates the state of the consensus state machine
type RoundStepType uint8 // These must be numeric, ordered.

// RoundStepType
const (
	RoundStepNewHeight     = RoundStepType(0x01) // Wait til CommitTime + timeoutCommit
	RoundStepNewRound      = RoundStepType(0x02) // Setup new round and go to RoundStepPropose
	RoundStepPropose       = RoundStepType(0x03) // Did propose, gossip proposal
	RoundStepPrevote       = RoundStepType(0x04) // Did prevote, gossip prevotes
	RoundStepPrevoteWait   = RoundStepType(0x05) // Did receive any +2/3 prevotes, start timeout
	RoundStepPrecommit     = RoundStepType(0x06) // Did precommit, gossip precommits
	RoundStepPrecommitWait = RoundStepType(0x07) // Did receive any +2/3 precommits, start timeout
	RoundStepCommit        = RoundStepType(0x08) // Entered commit state machine
	// NOTE: RoundStepNewHeight acts as RoundStepCommitWait.

	// NOTE: Update IsValid method if you change this!
)

// IsValid returns true if the step is valid, false if unknown/undefined.
func (rs RoundStepType) IsValid() bool { _ = "STUB: not implemented"; return false }

// String returns a string
func (rs RoundStepType) String() string { _ = "STUB: not implemented"; return "" }

// Cannot panic.

//-----------------------------------------------------------------------------

// RoundState defines the internal consensus state.
// NOTE: Not thread safe. Should only be manipulated by functions downstream
// of the cs.receiveRoutine
type RoundState struct {
	Height    int64         `json:"height"` // Height we are working on
	Round     int32         `json:"round"`
	Step      RoundStepType `json:"step"`
	StartTime time.Time     `json:"start_time"`

	// Subjective time when +2/3 precommits for Block at Round were found
	CommitTime         time.Time           `json:"commit_time"`
	Validators         *types.ValidatorSet `json:"validators"`
	Proposal           *types.Proposal     `json:"proposal"`
	ProposalBlock      *types.Block        `json:"proposal_block"`
	ProposalBlockParts *types.PartSet      `json:"proposal_block_parts"`
	LockedRound        int32               `json:"locked_round"`
	LockedBlock        *types.Block        `json:"locked_block"`
	LockedBlockParts   *types.PartSet      `json:"locked_block_parts"`

	// The variables below starting with "Valid..." derive their name from
	// the algorithm presented in this paper:
	// [The latest gossip on BFT consensus](https://arxiv.org/abs/1807.04938).
	// Therefore, "Valid...":
	//   * means that the block or round that the variable refers to has
	//     received 2/3+ non-`nil` prevotes (a.k.a. a *polka*)
	//   * has nothing to do with whether the Application returned "Accept" in its
	//     response to `ProcessProposal`, or "Reject"

	// Last known round with POL for non-nil valid block.
	ValidRound int32        `json:"valid_round"`
	ValidBlock *types.Block `json:"valid_block"` // Last known block of POL mentioned above.

	// Last known block parts of POL mentioned above.
	ValidBlockParts           *types.PartSet      `json:"valid_block_parts"`
	Votes                     *HeightVoteSet      `json:"votes"`
	CommitRound               int32               `json:"commit_round"` //
	LastCommit                *types.VoteSet      `json:"last_commit"`  // Last precommits at Height-1
	LastValidators            *types.ValidatorSet `json:"last_validators"`
	TriggeredTimeoutPrecommit bool                `json:"triggered_timeout_precommit"`
}

// Compressed version of the RoundState for use in RPC
type RoundStateSimple struct {
	HeightRoundStep   string              `json:"height/round/step"`
	StartTime         time.Time           `json:"start_time"`
	ProposalBlockHash bytes.HexBytes      `json:"proposal_block_hash"`
	LockedBlockHash   bytes.HexBytes      `json:"locked_block_hash"`
	ValidBlockHash    bytes.HexBytes      `json:"valid_block_hash"`
	Votes             json.RawMessage     `json:"height_vote_set"`
	Proposer          types.ValidatorInfo `json:"proposer"`
}

// Compress the RoundState to RoundStateSimple
func (rs *RoundState) RoundStateSimple() RoundStateSimple {
	_ = "STUB: not implemented"
	return *new(RoundStateSimple)
}

// NewRoundEvent returns the RoundState with proposer information as an event.
func (rs *RoundState) NewRoundEvent() types.EventDataNewRound {
	_ = "STUB: not implemented"
	return *new(types.EventDataNewRound)
}

// CompleteProposalEvent returns information about a proposed block as an event.
func (rs *RoundState) CompleteProposalEvent() types.EventDataCompleteProposal {
	_ = "STUB: not implemented"
	// We must construct BlockID from ProposalBlock and ProposalBlockParts
	// cs.Proposal is not guaranteed to be set when this function is called
	return *new(types.EventDataCompleteProposal)
}

// RoundStateEvent returns the H/R/S of the RoundState as an event.
func (rs *RoundState) RoundStateEvent() types.EventDataRoundState {
	_ = "STUB: not implemented"
	return *new(types.EventDataRoundState)
}

// String returns a string
func (rs *RoundState) String() string { _ = "STUB: not implemented"; return "" }

// StringIndented returns a string
func (rs *RoundState) StringIndented(indent string) string { _ = "STUB: not implemented"; return "" }

// StringShort returns a string
func (rs *RoundState) StringShort() string { _ = "STUB: not implemented"; return "" }
