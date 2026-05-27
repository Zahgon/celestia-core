package schema

import (
	"time"

	"github.com/cometbft/cometbft/libs/trace"
	"github.com/cometbft/cometbft/types"
)

// ConsensusTables returns the list of tables that are used for consensus
// tracing.
func ConsensusTables() []string { _ = "STUB: not implemented"; return nil }

// Schema constants for the consensus round state tracing database.
const (
	// RoundStateTable is the name of the table that stores the consensus
	// state traces.
	RoundStateTable = "consensus_round_state"
)

// RoundState describes schema for the "consensus_round_state" table.
type RoundState struct {
	Height int64  `json:"height"`
	Round  int32  `json:"round"`
	Step   string `json:"step"`
}

// Table returns the table name for the RoundState struct.
func (RoundState) Table() string { _ = "STUB: not implemented"; return "" }

// WriteRoundState writes a tracing point for a tx using the predetermined
// schema for consensus state tracing.
func WriteRoundState(client trace.Tracer, height int64, round int32, step string) {
	_ = "STUB: not implemented"
	return
}

// Schema constants for the "consensus_block_parts" table.
const (
	// BlockPartsTable is the name of the table that stores the consensus block
	// parts.
	BlockPartsTable = "consensus_block_parts"
)

// BlockPart describes schema for the "consensus_block_parts" table.
type BlockPart struct {
	Height       int64        `json:"height"`
	Round        int32        `json:"round"`
	Index        int32        `json:"index"`
	Catchup      bool         `json:"catchup"`
	Peer         string       `json:"peer"`
	TransferType TransferType `json:"transfer_type"`
}

// Table returns the table name for the BlockPart struct.
func (BlockPart) Table() string { _ = "STUB: not implemented"; return "" }

// WriteBlockPart writes a tracing point for a BlockPart using the predetermined
// schema for consensus state tracing.
func WriteBlockPart(
	client trace.Tracer,
	height int64,
	round int32,
	index uint32,
	catchup bool,
	peer string,
	transferType TransferType,
) {
	_ = "STUB: not implemented"
	// this check is redundant to what is checked during client.Write, although it
	// is an optimization to avoid allocations from the map of fields.
	return
}

//nolint:gosec

// Schema constants for the consensus votes tracing database.
const (
	// VoteTable is the name of the table that stores the consensus
	// voting traces.
	VoteTable = "consensus_vote"
)

// Vote describes schema for the "consensus_vote" table.
type Vote struct {
	Height           int64        `json:"height"`
	Round            int32        `json:"round"`
	VoteType         string       `json:"vote_type"`
	VoteHeight       int64        `json:"vote_height"`
	VoteRound        int32        `json:"vote_round"`
	VoteTimestamp    time.Time    `json:"vote_timestamp"`
	ValidatorAddress string       `json:"vote_validator_address"`
	Peer             string       `json:"peer"`
	TransferType     TransferType `json:"transfer_type"`
}

func (Vote) Table() string {
	_ = "STUB: not implemented"

	// WriteVote writes a tracing point for a vote using the predetermined
	// schema for consensus vote tracing.
	return ""
}

func WriteVote(client trace.Tracer,
	height int64, // height of the current peer when it received/sent the vote
	round int32, // round of the current peer when it received/sent the vote
	vote *types.Vote, // vote received by the current peer
	peer string, // the peer from which it received the vote or the peer to which it sent the vote
	transferType TransferType, // download (received) or upload(sent)
) {
	_ = "STUB: not implemented"
	return
}

const (
	// BlockTable is the name of the table that stores metadata about consensus blocks.
	BlockTable = "consensus_block"
)

// BlockSummary describes schema for the "consensus_block" table.
type BlockSummary struct {
	Height                   int64  `json:"height"`
	UnixMillisecondTimestamp int64  `json:"unix_millisecond_timestamp"`
	TxCount                  int    `json:"tx_count"`
	SquareSize               uint64 `json:"square_size"`
	BlockSize                int    `json:"block_size"`
	Proposer                 string `json:"proposer"`
	LastCommitRound          int32  `json:"last_commit_round"`
}

func (BlockSummary) Table() string {
	_ = "STUB: not implemented"

	// WriteBlockSummary writes a tracing point for a block using the predetermined.
	return ""
}

func WriteBlockSummary(client trace.Tracer, block *types.Block, size int) {
	_ = "STUB: not implemented"
	return
}

//nolint:staticcheck

const (
	ConsensusStateTable = "consensus_state"
)

type ConsensusStateUpdateType string

const (
	ConsensusNewValidBlock      ConsensusStateUpdateType = "new_valid_block"
	ConsensusNewRoundStep       ConsensusStateUpdateType = "new_round_step"
	ConsensusVoteSetBits        ConsensusStateUpdateType = "vote_set_bits"
	ConsensusVoteSet23Prevote   ConsensusStateUpdateType = "vote_set_23_prevote"
	ConsensusVoteSet23Precommit ConsensusStateUpdateType = "vote_set_23_precommit"
	ConsensusHasVote            ConsensusStateUpdateType = "has_vote"
	ConsensusPOL                ConsensusStateUpdateType = "pol"
)

type ConsensusState struct {
	Height       int64        `json:"height"`
	Round        int32        `json:"round"`
	UpdateType   string       `json:"update_type"`
	Peer         string       `json:"peer"`
	TransferType TransferType `json:"transfer_type"`
	Data         []string     `json:"data,omitempty"`
}

func (ConsensusState) Table() string { _ = "STUB: not implemented"; return "" }

func WriteConsensusState(
	client trace.Tracer,
	height int64,
	round int32,
	peer string,
	updateType ConsensusStateUpdateType,
	transferType TransferType,
	data ...string,
) {
	_ = "STUB: not implemented"
	return
}

const (
	ProposalTable = "consensus_proposal"
)

type Proposal struct {
	Height       int64        `json:"height"`
	Round        int32        `json:"round"`
	PeerID       string       `json:"peer_id"`
	TransferType TransferType `json:"transfer_type"`
}

func (Proposal) Table() string { _ = "STUB: not implemented"; return "" }

func WriteProposal(
	client trace.Tracer,
	height int64,
	round int32,
	peerID string,
	transferType TransferType,
) {
	_ = "STUB: not implemented"
	return
}

// Schema constants for the "consensus_block_parts" table.
const (
	// BlockPartsTable is the name of the table that stores the consensus block
	// parts.
	BlockPartStateTable = "bp_state"
)

// BlockPart describes schema for the "consensus_block_parts" table.
type BlockPartState struct {
	Height       int64        `json:"height"`
	Round        int32        `json:"round"`
	Indexes      []int        `json:"indexes"`
	Have         bool         `json:"have"`
	Peer         string       `json:"peer"`
	TransferType TransferType `json:"transfer_type"`
	MessageType  string       `json:"message_type"`
}

// Table returns the table name for the BlockPart struct.
func (b BlockPartState) Table() string { _ = "STUB: not implemented"; return "" }

// WriteBlockPart writes a tracing point for a BlockPart using the predetermined
// schema for consensus state tracing.
func WriteBlockPartState(
	client trace.Tracer,
	height int64,
	round int32,
	indexes []int,
	have bool,
	peer string,
	transferType TransferType,
	messageType string,
) {
	_ = "STUB: not implemented"
	// this check is redundant to what is checked during client.Write, although it
	// is an optimization to avoid allocations from the map of fields.
	return
}

//nolint:gosec

const (
	NotesTable = "notes"
)

type Note struct {
	Note     string `json:"note"`
	Height   int64  `json:"height"`
	Round    int32  `json:"round"`
	NoteType string `json:"note_type"`
}

func (p Note) Table() string { _ = "STUB: not implemented"; return "" }

func WriteNote(
	client trace.Tracer,
	height int64,
	round int32,
	noteType string,
	note string,
	items ...interface{},
) {
	_ = "STUB: not implemented"
	return
}

const (
	CatchupRequestsTable = "catch_reqs"
)

type CatchupRequest struct {
	Height int64  `json:"height"`
	Round  int32  `json:"round"`
	Parts  string `json:"parts"`
	Peer   string `json:"peer"`
}

func (b CatchupRequest) Table() string { _ = "STUB: not implemented"; return "" }

func WriteCatchupRequest(
	client trace.Tracer,
	height int64,
	round int32,
	parts string,
	peer string,
) {
	_ = "STUB: not implemented"
	// this check is redundant to what is checked during client.Write, although it
	// is an optimization to avoid allocations from the map of fields.
	return
}

const (
	RetriesTable = "retries"
)

type Retries struct {
	Height  int64  `json:"height"`
	Round   int32  `json:"round"`
	Missing string `json:"missing"`
}

func (b Retries) Table() string { _ = "STUB: not implemented"; return "" }

func WriteRetries(
	client trace.Tracer,
	height int64,
	round int32,
	missing string,
) {
	_ = "STUB: not implemented"
	return
}

const (
	GapTable = "gap"
)

type Gap struct {
	Height int64 `json:"height"`
	Round  int32 `json:"round"`
}

func (b Gap) Table() string { _ = "STUB: not implemented"; return "" }

func WriteGap(
	client trace.Tracer,
	height int64,
	round int32,
) {
	_ = "STUB: not implemented"
	return
}

const (
	MissedProposalsTable = "missed_proposals"
)

type MissedProposal struct {
	Height   int64  `json:"height"`
	Round    int32  `json:"round"`
	Proposer string `json:"proposer"`
}

func (b MissedProposal) Table() string { _ = "STUB: not implemented"; return "" }

func WriteMissedProposal(
	client trace.Tracer,
	height int64,
	round int32,
	proposerAddress string,
) {
	_ = "STUB: not implemented"
	return
}

const (
	SigningLatencyTable = "signing_latency"
	PrecommitType       = "precommit"
	PrevoteType         = "prevote"
	ProposalType        = "proposal"
)

type SignatureLatency struct {
	Height      int64  `json:"height"`
	Round       int32  `json:"round"`
	Latency     int64  `json:"latency"`
	MessageType string `json:"message_type"`
}

func (b SignatureLatency) Table() string { _ = "STUB: not implemented"; return "" }

func WriteSignatureLatency(
	client trace.Tracer,
	height int64,
	round int32,
	latency int64,
	msgType string,
) {
	_ = "STUB: not implemented"
	return
}

const (
	FullBlockReceivingTimeTable = "full_block_receiving_time"
)

type FullBlockReceivingTime struct {
	Height int64 `json:"height"`
	Round  int32 `json:"round"`
	// Duration in milliseconds from when the proposal was received to when the full block was received
	DurationMs int64 `json:"duration_ms"`
}

func (b FullBlockReceivingTime) Table() string { _ = "STUB: not implemented"; return "" }

func WriteFullBlockReceivingTime(
	client trace.Tracer,
	height int64,
	round int32,
	duration time.Duration,
) {
	_ = "STUB: not implemented"
	return
}
