package types

import (
	"github.com/cometbft/cometbft/libs/bits"
	cmtsync "github.com/cometbft/cometbft/libs/sync"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
)

const (
	// MaxVotesCount is the maximum number of votes in a set. Used in ValidateBasic funcs for
	// protection against DOS attacks. Note this implies a corresponding equal limit to
	// the number of validators.
	MaxVotesCount = 10000
)

// UNSTABLE
// XXX: duplicate of p2p.ID to avoid dependence between packages.
// Perhaps we can have a minimal types package containing this (and other things?)
// that both `types` and `p2p` import ?
type P2PID string

/*
VoteSet helps collect signatures from validators at each height+round for a
predefined vote type.

We need VoteSet to be able to keep track of conflicting votes when validators
double-sign.  Yet, we can't keep track of *all* the votes seen, as that could
be a DoS attack vector.

There are two storage areas for votes.
1. voteSet.votes
2. voteSet.votesByBlock

`.votes` is the "canonical" list of votes.  It always has at least one vote,
if a vote from a validator had been seen at all.  Usually it keeps track of
the first vote seen, but when a 2/3 majority is found, votes for that get
priority and are copied over from `.votesByBlock`.

`.votesByBlock` keeps track of a list of votes for a particular block.  There
are two ways a &blockVotes{} gets created in `.votesByBlock`.
1. the first vote seen by a validator was for the particular block.
2. a peer claims to have seen 2/3 majority for the particular block.

Since the first vote from a validator will always get added in `.votesByBlock`
, all votes in `.votes` will have a corresponding entry in `.votesByBlock`.

When a &blockVotes{} in `.votesByBlock` reaches a 2/3 majority quorum, its
votes are copied into `.votes`.

All this is memory bounded because conflicting votes only get added if a peer
told us to track that block, each peer only gets to tell us 1 such block, and,
there's only a limited number of peers.

NOTE: Assumes that the sum total of voting power does not exceed MaxUInt64.
*/
type VoteSet struct {
	chainID           string
	height            int64
	round             int32
	signedMsgType     cmtproto.SignedMsgType
	valSet            *ValidatorSet
	extensionsEnabled bool

	mtx           cmtsync.Mutex
	votesBitArray *bits.BitArray
	votes         []*Vote                // Primary votes to share
	sum           int64                  // Sum of voting power for seen votes, discounting conflicts
	maj23         *BlockID               // First 2/3 majority seen
	votesByBlock  map[string]*blockVotes // string(blockHash|blockParts) -> blockVotes
	peerMaj23s    map[P2PID]BlockID      // Maj23 for each peer
}

// NewVoteSet instantiates all fields of a new vote set. This constructor requires
// that no vote extension data be present on the votes that are added to the set.
func NewVoteSet(chainID string, height int64, round int32,
	signedMsgType cmtproto.SignedMsgType, valSet *ValidatorSet) *VoteSet {
	_ = "STUB: not implemented"
	return nil
}

// NewExtendedVoteSet constructs a vote set with additional vote verification logic.
// The VoteSet constructed with NewExtendedVoteSet verifies the vote extension
// data for every vote added to the set.
func NewExtendedVoteSet(chainID string, height int64, round int32,
	signedMsgType cmtproto.SignedMsgType, valSet *ValidatorSet) *VoteSet {
	_ = "STUB: not implemented"
	return nil
}

func (voteSet *VoteSet) ChainID() string { _ = "STUB: not implemented"; return "" }

// Implements VoteSetReader.
func (voteSet *VoteSet) GetHeight() int64 { _ = "STUB: not implemented"; return 0 }

// Implements VoteSetReader.
func (voteSet *VoteSet) GetRound() int32 { _ = "STUB: not implemented"; return 0 }

// Implements VoteSetReader.
func (voteSet *VoteSet) Type() byte { _ = "STUB: not implemented"; return 0 }

// Implements VoteSetReader.
func (voteSet *VoteSet) Size() int { _ = "STUB: not implemented"; return 0 }

// Returns added=true if vote is valid and new.
// Otherwise returns err=ErrVote[
//
//	UnexpectedStep | InvalidIndex | InvalidAddress |
//	InvalidSignature | InvalidBlockHash | ConflictingVotes ]
//
// Duplicate votes return added=false, err=nil.
// Conflicting votes return added=*, err=ErrVoteConflictingVotes.
// NOTE: vote should not be mutated after adding.
// NOTE: VoteSet must not be nil
// NOTE: Vote must not be nil
func (voteSet *VoteSet) AddVote(vote *Vote) (added bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

// NOTE: Validates as much as possible before attempting to verify the signature.
func (voteSet *VoteSet) addVote(vote *Vote) (added bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Ensure that validator index was set

// Make sure the step matches.

// Ensure that signer is a validator.

// Ensure that the signer has the right address.

// If we already know of this vote, return false.

// duplicate

// Check signature.

// Add vote and get conflicting vote if any.

// Returns (vote, true) if vote exists for valIndex and blockKey.
func (voteSet *VoteSet) getVote(valIndex int32, blockKey string) (vote *Vote, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// Assumes signature is valid.
// If conflicting vote exists, returns it.
func (voteSet *VoteSet) addVerifiedVote(
	vote *Vote,
	blockKey string,
	votingPower int64,
) (added bool, conflicting *Vote) {
	_ = "STUB: not implemented"
	return false, nil

	// Already exists in voteSet.votes?
}

// Replace vote if blockKey matches voteSet.maj23.

// Otherwise don't add it to voteSet.votes

// Add to voteSet.votes and incr .sum

// There's a conflict and no peer claims that this block is special.

// We'll add the vote in a bit.

// .votesByBlock doesn't exist...

// ... and there's a conflicting vote.
// We're not even tracking this blockKey, so just forget it.

// ... and there's no conflicting vote.
// Start tracking this blockKey

// We'll add the vote in a bit.

// Before adding to votesByBlock, see if we'll exceed quorum

// Add vote to votesByBlock

// If we just crossed the quorum threshold and have 2/3 majority...

// Only consider the first quorum reached

// And also copy votes over to voteSet.votes

// If a peer claims that it has 2/3 majority for given blockKey, call this.
// NOTE: if there are too many peers, or too much peer churn,
// this can cause memory issues.
// TODO: implement ability to remove peers too
// NOTE: VoteSet must not be nil
func (voteSet *VoteSet) SetPeerMaj23(peerID P2PID, blockID BlockID) error {
	_ = "STUB: not implemented"
	return nil
}

// Make sure peer hasn't already told us something.

// Nothing to do

// Create .votesByBlock entry if needed.

// Nothing to do

// No need to copy votes, already there.

// No need to copy votes, no votes to copy over.

// Implements VoteSetReader.
func (voteSet *VoteSet) BitArray() *bits.BitArray { _ = "STUB: not implemented"; return nil }

func (voteSet *VoteSet) BitArrayByBlockID(blockID BlockID) *bits.BitArray {
	_ = "STUB: not implemented"
	return nil
}

// NOTE: if validator has conflicting votes, returns "canonical" vote
// Implements VoteSetReader.
func (voteSet *VoteSet) GetByIndex(valIndex int32) *Vote { _ = "STUB: not implemented"; return nil }

// List returns a copy of the list of votes stored by the VoteSet.
func (voteSet *VoteSet) List() []Vote { _ = "STUB: not implemented"; return nil }

func (voteSet *VoteSet) GetByAddress(address []byte) *Vote { _ = "STUB: not implemented"; return nil }

func (voteSet *VoteSet) HasTwoThirdsMajority() bool { _ = "STUB: not implemented"; return false }

// Implements VoteSetReader.
func (voteSet *VoteSet) IsCommit() bool { _ = "STUB: not implemented"; return false }

func (voteSet *VoteSet) HasTwoThirdsAny() bool { _ = "STUB: not implemented"; return false }

func (voteSet *VoteSet) HasAll() bool { _ = "STUB: not implemented"; return false }

// If there was a +2/3 majority for blockID, return blockID and true.
// Else, return the empty BlockID{} and false.
func (voteSet *VoteSet) TwoThirdsMajority() (blockID BlockID, ok bool) {
	_ = "STUB: not implemented"
	return *new(BlockID), false
}

//--------------------------------------------------------------------------------
// Strings and JSON

const nilVoteSetString = "nil-VoteSet"

// String returns a string representation of VoteSet.
//
// See StringIndented.
func (voteSet *VoteSet) String() string { _ = "STUB: not implemented"; return "" }

// StringIndented returns an indented String.
//
// Height Round Type
// Votes
// Votes bit array
// 2/3+ majority
//
// See Vote#String.
func (voteSet *VoteSet) StringIndented(indent string) string { _ = "STUB: not implemented"; return "" }

// Marshal the VoteSet to JSON. Same as String(), just in JSON,
// and without the height/round/signedMsgType (since its already included in the votes).
func (voteSet *VoteSet) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// More human readable JSON of the vote set
// NOTE: insufficient for unmarshalling from (compressed votes)
// TODO: make the peerMaj23s nicer to read (eg just the block hash)
type VoteSetJSON struct {
	Votes         []string          `json:"votes"`
	VotesBitArray string            `json:"votes_bit_array"`
	PeerMaj23s    map[P2PID]BlockID `json:"peer_maj_23s"`
}

// Return the bit-array of votes including
// the fraction of power that has voted like:
// "BA{29:xx__x__x_x___x__x_______xxx__} 856/1304 = 0.66"
func (voteSet *VoteSet) BitArrayString() string { _ = "STUB: not implemented"; return "" }

func (voteSet *VoteSet) bitArrayString() string { _ = "STUB: not implemented"; return "" }

// Returns a list of votes compressed to more readable strings.
func (voteSet *VoteSet) VoteStrings() []string { _ = "STUB: not implemented"; return nil }

func (voteSet *VoteSet) voteStrings() []string { _ = "STUB: not implemented"; return nil }

// StringShort returns a short representation of VoteSet.
//
// 1. height
// 2. round
// 3. signed msg type
// 4. first 2/3+ majority
// 5. fraction of voted power
// 6. votes bit array
// 7. 2/3+ majority for each peer
func (voteSet *VoteSet) StringShort() string { _ = "STUB: not implemented"; return "" }

// LogString produces a logging suitable string representation of the
// vote set.
func (voteSet *VoteSet) LogString() string { _ = "STUB: not implemented"; return "" }

// return the power voted, the total, and the fraction
func (voteSet *VoteSet) sumTotalFrac() (int64, int64, float64) {
	_ = "STUB: not implemented"
	return 0, 0, 0
}

//--------------------------------------------------------------------------------
// Commit

// MakeExtendedCommit constructs a Commit from the VoteSet. It only includes
// precommits for the block, which has 2/3+ majority, and nil.
//
// Panics if the vote type is not PrecommitType or if there's no +2/3 votes for
// a single block.
func (voteSet *VoteSet) MakeExtendedCommit(ap ABCIParams) *ExtendedCommit {
	_ = "STUB: not implemented"
	return nil
}

// Make sure we have a 2/3 majority

// For every validator, get the precommit with extensions

// if block ID exists but doesn't match, exclude sig

//--------------------------------------------------------------------------------

/*
Votes for a particular block
There are two ways a *blockVotes gets created for a blockKey.
1. first (non-conflicting) vote of a validator w/ blockKey (peerMaj23=false)
2. A peer claims to have a 2/3 majority w/ blockKey (peerMaj23=true)
*/
type blockVotes struct {
	peerMaj23 bool           // peer claims to have maj23
	bitArray  *bits.BitArray // valIndex -> hasVote?
	votes     []*Vote        // valIndex -> *Vote
	sum       int64          // vote sum
}

func newBlockVotes(peerMaj23 bool, numValidators int) *blockVotes {
	_ = "STUB: not implemented"
	return nil
}

func (vs *blockVotes) addVerifiedVote(vote *Vote, votingPower int64) {
	_ = "STUB: not implemented"
	return
}

func (vs *blockVotes) getByIndex(index int32) *Vote { _ = "STUB: not implemented"; return nil }

//--------------------------------------------------------------------------------

// Common interface between *consensus.VoteSet and types.Commit
type VoteSetReader interface {
	GetHeight() int64
	GetRound() int32
	Type() byte
	Size() int
	BitArray() *bits.BitArray
	GetByIndex(int32) *Vote
	IsCommit() bool
}
