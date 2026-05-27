package state

import (
	"github.com/cometbft/cometbft/types"
)

//-----------------------------------------------------
// Validate block

func validateBlock(state State, block *types.Block) error {
	_ = "STUB: not implemented"
	// Validate internal consistency.
	return nil
}

// Validate basic info.

// Validate prev block info.

// Validate app info

// Validate block LastCommit.

// LastCommit.Signatures length is checked in VerifyCommit.

// NOTE: We can't actually verify it's the right proposer because we don't
// know what round the block was first proposed. So just check that it's
// a legit address and a known validator.

// Validate block Time

// Check evidence doesn't exceed the limit amount of bytes.
