package test

import (
	"github.com/cometbft/cometbft/types"
)

func MakeNTxs(height, n int64) types.Txs { _ = "STUB: not implemented"; return *new(types.Txs) }

func MakeNTxsWithSize(height, n int, size int) types.Txs {
	_ = "STUB: not implemented"
	return *new(types.Txs)
}
