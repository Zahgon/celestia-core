package test

import (
	"time"

	"github.com/cometbft/cometbft/types"
)

func GenesisDoc(
	time time.Time,
	validators []*types.Validator,
	consensusParams *types.ConsensusParams,
	chainID string,
) *types.GenesisDoc {
	_ = "STUB: not implemented"
	return nil
}
