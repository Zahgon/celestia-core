package tmhash

import (
	"crypto/sha256"
	"hash"
)

const (
	Size      = sha256.Size
	BlockSize = sha256.BlockSize
)

// New returns a new hash.Hash.
func New() hash.Hash {
	_ = "STUB: not implemented"
	return *

	// Sum returns the SHA256 of the bz.
	new(hash.Hash)
}

func Sum(bz []byte) []byte { _ = "STUB: not implemented"; return nil }

// SumMany takes at least 1 byteslice along with a variadic
// number of other byteslices and produces the SHA256 sum from
// hashing them as if they were 1 joined slice.
func SumMany(data []byte, rest ...[]byte) []byte { _ = "STUB: not implemented"; return nil }

//-------------------------------------------------------------

const (
	TruncatedSize = 20
)

type sha256trunc struct {
	sha256 hash.Hash
}

func (h sha256trunc) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (h sha256trunc) Sum(b []byte) []byte { _ = "STUB: not implemented"; return nil }

func (h sha256trunc) Reset() { _ = "STUB: not implemented"; return }

func (h sha256trunc) Size() int { _ = "STUB: not implemented"; return 0 }

func (h sha256trunc) BlockSize() int { _ = "STUB: not implemented"; return 0 }

// NewTruncated returns a new hash.Hash.
func NewTruncated() hash.Hash { _ = "STUB: not implemented"; return *new(hash.Hash) }

// SumTruncated returns the first 20 bytes of SHA256 of the bz.
func SumTruncated(bz []byte) []byte { _ = "STUB: not implemented"; return nil }
