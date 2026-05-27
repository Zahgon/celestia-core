package merkle

import (
	"hash"
)

// TODO: make these have a large predefined capacity
var (
	leafPrefix  = []byte{0}
	innerPrefix = []byte{1}
)

// returns tmhash(<empty>)
func emptyHash() []byte { _ = "STUB: not implemented"; return nil }

// returns tmhash(0x00 || leaf)
func leafHash(leaf []byte) []byte { _ = "STUB: not implemented"; return nil }

// returns tmhash(0x00 || leaf)
func leafHashOpt(s hash.Hash, leaf []byte) []byte { _ = "STUB: not implemented"; return nil }

func LeafHash(leaf []byte) []byte { _ = "STUB: not implemented"; return nil }

// returns tmhash(0x01 || left || right)
func innerHash(left []byte, right []byte) []byte { _ = "STUB: not implemented"; return nil }

func innerHashOpt(s hash.Hash, left []byte, right []byte) []byte {
	_ = "STUB: not implemented"
	return nil
}
