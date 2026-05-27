package statesync

import (
	"errors"

	"github.com/cosmos/gogoproto/proto"
)

const (
	// snapshotMsgSize is the maximum size of a snapshotResponseMessage
	snapshotMsgSize = int(4e6)
	// chunkMsgSize is the maximum size of a chunkResponseMessage
	chunkMsgSize = int(16e6)
)

var (
	ErrExceedsMaxSnapshotChunks = errors.New("amount of chunks in the snapshot exceeds the maximum allowed number of chunks")
)

// validateMsg validates a message.
func validateMsg(pb proto.Message, maxSnapshotChunks uint32) error {
	_ = "STUB: not implemented"
	return nil
}
