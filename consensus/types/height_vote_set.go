package types

import (
	"errors"
	"sync"

	"github.com/cometbft/cometbft/p2p"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cometbft/cometbft/types"
)

type RoundVoteSet struct {
	Prevotes   *types.VoteSet
	Precommits *types.VoteSet
}

var (
	ErrGotVoteFromUnwantedRound = errors.New(
		"peer has sent a vote that does not match our round for more than one round",
	)
)

/*
Keeps track of all VoteSets from round 0 to round 'round'.

Also keeps track of up to one RoundVoteSet greater than
'round' from each peer, to facilitate catchup syncing of commits.

A commit is +2/3 precommits for a block at a round,
but which round is not known in advance, so when a peer
provides a precommit for a round greater than mtx.round,
we create a new entry in roundVoteSets but also remember the
peer to prevent abuse.
We let each peer provide us with up to 2 unexpected "catchup" rounds.
One for their LastCommit round, and another for the official commit round.
*/
type HeightVoteSet struct {
	chainID           string
	height            int64
	valSet            *types.ValidatorSet
	extensionsEnabled bool

	mtx               sync.Mutex
	round             int32                  // max tracked round
	roundVoteSets     map[int32]RoundVoteSet // keys: [0...round]
	peerCatchupRounds map[p2p.ID][]int32     // keys: peer.ID; values: at most 2 rounds
}

func NewHeightVoteSet(chainID string, height int64, valSet *types.ValidatorSet) *HeightVoteSet {
	_ = "STUB: not implemented"
	return nil
}

func NewExtendedHeightVoteSet(chainID string, height int64, valSet *types.ValidatorSet) *HeightVoteSet {
	_ = "STUB: not implemented"
	return nil
}

func (hvs *HeightVoteSet) Reset(height int64, valSet *types.ValidatorSet) {
	_ = "STUB: not implemented"
	return
}

func (hvs *HeightVoteSet) Height() int64 { _ = "STUB: not implemented"; return 0 }

func (hvs *HeightVoteSet) Round() int32 { _ = "STUB: not implemented"; return 0 }

// Create more RoundVoteSets up to round.
func (hvs *HeightVoteSet) SetRound(round int32) { _ = "STUB: not implemented"; return }

// Already exists because peerCatchupRounds.

func (hvs *HeightVoteSet) addRound(round int32) { _ = "STUB: not implemented"; return }

// log.Debug("addRound(round)", "round", round)

// Duplicate votes return added=false, err=nil.
// By convention, peerID is "" if origin is self.
func (hvs *HeightVoteSet) AddVote(vote *types.Vote, peerID p2p.ID, extEnabled bool) (added bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

// punish peer

func (hvs *HeightVoteSet) Prevotes(round int32) *types.VoteSet {
	_ = "STUB: not implemented"
	return nil
}

func (hvs *HeightVoteSet) Precommits(round int32) *types.VoteSet {
	_ = "STUB: not implemented"
	return nil
}

// Last round and blockID that has +2/3 prevotes for a particular block or nil.
// Returns -1 if no such round exists.
func (hvs *HeightVoteSet) POLInfo() (polRound int32, polBlockID types.BlockID) {
	_ = "STUB: not implemented"
	return 0, *new(types.BlockID)
}

func (hvs *HeightVoteSet) getVoteSet(round int32, voteType cmtproto.SignedMsgType) *types.VoteSet {
	_ = "STUB: not implemented"
	return nil
}

// If a peer claims that it has 2/3 majority for given blockKey, call this.
// NOTE: if there are too many peers, or too much peer churn,
// this can cause memory issues.
// TODO: implement ability to remove peers too
func (hvs *HeightVoteSet) SetPeerMaj23(
	round int32,
	voteType cmtproto.SignedMsgType,
	peerID p2p.ID,
	blockID types.BlockID) error {
	_ = "STUB: not implemented"
	return nil
}

// something we don't know about yet

//---------------------------------------------------------
// string and json

func (hvs *HeightVoteSet) String() string { _ = "STUB: not implemented"; return "" }

func (hvs *HeightVoteSet) StringIndented(indent string) string {
	_ = "STUB: not implemented"
	return ""
}

// rounds 0 ~ hvs.round inclusive

// all other peer catchup rounds

func (hvs *HeightVoteSet) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (hvs *HeightVoteSet) toAllRoundVotes() []roundVotes { _ = "STUB: not implemented"; return nil }

// rounds 0 ~ hvs.round inclusive

// TODO: all other peer catchup rounds

type roundVotes struct {
	Round              int32    `json:"round"`
	Prevotes           []string `json:"prevotes"`
	PrevotesBitArray   string   `json:"prevotes_bit_array"`
	Precommits         []string `json:"precommits"`
	PrecommitsBitArray string   `json:"precommits_bit_array"`
}
