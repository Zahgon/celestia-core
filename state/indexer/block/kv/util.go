package kv

import (
	"github.com/cometbft/cometbft/libs/pubsub/query/syntax"
	"github.com/cometbft/cometbft/state/indexer"
)

type HeightInfo struct {
	heightRange     indexer.QueryRange
	height          int64
	heightEqIdx     int
	onlyHeightRange bool
	onlyHeightEq    bool
}

func intInSlice(a int, list []int) bool { _ = "STUB: not implemented"; return false }

func int64FromBytes(bz []byte) int64 { _ = "STUB: not implemented"; return 0 }

func int64ToBytes(i int64) []byte { _ = "STUB: not implemented"; return nil }

func heightKey(height int64) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func eventKey(compositeKey, eventValue string, height int64, eventSeq int64) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseValueFromPrimaryKey(key []byte) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func parseValueFromEventKey(key []byte) (string, error) { _ = "STUB: not implemented"; return "", nil }

func parseHeightFromEventKey(key []byte) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func parseEventSeqFromEventKey(key []byte) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// We either have an event sequence or a function type (potentially) followed by an event sequence.
// Potential scenarios:
// 1. Events indexed with v0.38.x and later, will only have an event sequence
// 2. Events indexed between v0.34.27 and v0.37.x will have a function type and an event sequence
// 3. Events indexed before v0.34.27 will only have a function type
// function_type = 'being_block_event' | 'end_block_event'

// The event was not properly indexed

// Check if we have scenarios 2. or 3. (described above).
// If it cannot parse the event function type, it could be 1.

// We should not have anything else after the eventSeq.

// Are we in case 2 or 3
// the event follows the scenario in 2.,
// retrieve the eventSeq
// there should be no error
// We should not have anything else after the eventSeq if in 2.

// Remove all occurrences of height equality queries except one. While we are traversing the conditions, check whether the only condition in
// addition to match events is the height equality or height range query. At the same time, if we do have a height range condition
// ignore the height equality condition. If a height equality exists, place the condition index in the query and the desired height
// into the heightInfo struct
func dedupHeight(conditions []syntax.Condition) (dedupConditions []syntax.Condition, heightInfo HeightInfo, found bool) {
	_ = "STUB: not implemented"
	return nil, *new(HeightInfo), false
}

// If we found a range make sure we set the hegiht idx to -1 as the height equality
// will be removed

func checkHeightConditions(heightInfo HeightInfo, keyHeight int64) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
