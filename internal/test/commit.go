package test

import (
	"time"

	"github.com/cometbft/cometbft/types"
)

func MakeCommitFromVoteSet(blockID types.BlockID, voteSet *types.VoteSet, validators []types.PrivValidator, now time.Time) (*types.Commit, error) {
	_ = "STUB: not implemented"
	// all sign
	return nil, nil
}

func MakeCommit(blockID types.BlockID, height int64, round int32, valSet *types.ValidatorSet, privVals []types.PrivValidator, chainID string, now time.Time) (*types.Commit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
