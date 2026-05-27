package types

import (
	"testing"
	"time"

	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
)

func MakeExtCommit(blockID BlockID, height int64, round int32,
	voteSet *VoteSet, validators []PrivValidator, now time.Time, extEnabled bool) (*ExtendedCommit, error) {
	_ = "STUB: not implemented"

	// all sign
	return nil, nil
}

func signAddVote(privVal PrivValidator, vote *Vote, voteSet *VoteSet) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func MakeVote(
	val PrivValidator,
	chainID string,
	valIndex int32,
	height int64,
	round int32,
	step cmtproto.SignedMsgType,
	blockID BlockID,
	time time.Time,
) (*Vote, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func MakeVoteNoError(
	t *testing.T,
	val PrivValidator,
	chainID string,
	valIndex int32,
	height int64,
	round int32,
	step cmtproto.SignedMsgType,
	blockID BlockID,
	time time.Time,
) *Vote {
	_ = "STUB: not implemented"
	return nil
}

// MakeBlock returns a new block with an empty header, except what can be
// computed from itself.
// It populates the same set of fields validated by ValidateBasic.
func MakeBlock(height int64, data Data, lastCommit *Commit, evidence []Evidence) *Block {
	_ = "STUB: not implemented"
	return nil
}

// MakeTxs is a helper function to generate mock transactions by given the block height
// and the transaction numbers.
func MakeTxs(height int64, num int) (txs []Tx) { _ = "STUB: not implemented"; return nil }

func MakeTenTxs(height int64) (txs []Tx) { _ = "STUB: not implemented"; return nil }

func MakeData(txs []Tx) Data { _ = "STUB: not implemented"; return *new(Data) }

func RandCommit(now time.Time) *Commit { _ = "STUB: not implemented"; return nil }

func RandVoteSet(
	height int64,
	round int32,
	signedMsgType cmtproto.SignedMsgType,
	numValidators int,
	votingPower int64,
) (*VoteSet, *ValidatorSet, []PrivValidator) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func MakeBlockIDRandom() BlockID { _ = "STUB: not implemented"; return *new(BlockID) }

//nolint: errcheck // ignore errcheck for read
//nolint: errcheck // ignore errcheck for read
