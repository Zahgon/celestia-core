package kv

import (
	cmtsyntax "github.com/cometbft/cometbft/libs/pubsub/query/syntax"
	"github.com/cometbft/cometbft/state/indexer"
)

type HeightInfo struct {
	heightRange     indexer.QueryRange
	height          int64
	heightEqIdx     int
	onlyHeightRange bool
	onlyHeightEq    bool
}

// IntInSlice returns true if a is found in the list.
func intInSlice(a int, list []int) bool { _ = "STUB: not implemented"; return false }

func ParseEventSeqFromEventKey(key []byte) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func dedupHeight(conditions []cmtsyntax.Condition) (dedupConditions []cmtsyntax.Condition, heightInfo HeightInfo) {
	_ = "STUB: not implemented"
	return nil, *new(HeightInfo)
}

// If we found a range make sure we set the height idx to -1 as the height equality
// will be removed

func checkHeightConditions(heightInfo HeightInfo, keyHeight int64) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
