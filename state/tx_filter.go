package state

import (
	mempl "github.com/cometbft/cometbft/mempool"
)

// TxPreCheck returns a function to filter transactions before processing.
// The function limits the size of a transaction to the block's maximum data size.
func TxPreCheck(state State) mempl.PreCheckFunc {
	_ = "STUB: not implemented"
	return *new(mempl.PreCheckFunc)
}

// TxPostCheck returns a function to filter transactions after processing.
// The function limits the gas wanted by a transaction to the block's maximum total gas.
func TxPostCheck(state State) mempl.PostCheckFunc {
	_ = "STUB: not implemented"
	return *new(mempl.PostCheckFunc)
}
