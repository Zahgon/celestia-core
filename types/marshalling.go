package types

import (
	"github.com/cosmos/gogoproto/proto"
)

// TxPosition holds the start and end indexes (in the overall encoded []byte)
// for a given txs field.
type TxPosition struct {
	Start uint32
	// End exclusive position of the transaction
	End uint32
}

// safeAddUint32 performs checked addition of two uint32 numbers.
func safeAddUint32(a, b uint32) (uint32, error) { _ = "STUB: not implemented"; return 0, nil }

// MarshalBlockWithTxPositions marshals the given Block message using protobuf
// and returns both the encoded []byte and a slice of positions marking the
// boundaries of each nested tx (repeated []byte field) inside Data (field number 1).
func MarshalBlockWithTxPositions(block proto.Message, txsCount int) ([]byte, []TxPosition, error) {
	_ = "STUB: not implemented"
	// First, marshal the entire message normally.
	return nil, nil, nil
}

// In our Block proto, field number 2 is the Data message.

// Read the field tag (a varint).

// length-delimited

// For non length-delimited fields, skip appropriately.

// varint

// 64-bit

// 32-bit

// findField scans the encoded message b for a length-delimited field with the given targetField number.
func findField(b []byte, targetField int) (contentStart int, contentEnd int, content []byte, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil, nil
}

// readVarint reads a varint-encoded unsigned integer from b.
func readVarint(b []byte) (uint64, int, error) { _ = "STUB: not implemented"; return 0, 0, nil }
