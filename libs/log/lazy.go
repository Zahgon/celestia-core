package log

import (
	cmtbytes "github.com/cometbft/cometbft/libs/bytes"
)

type LazySprintf struct {
	format string
	args   []interface{}
}

// NewLazySprintf defers fmt.Sprintf until the Stringer interface is invoked.
// This is particularly useful for avoiding calling Sprintf when debugging is not
// active.
func NewLazySprintf(format string, args ...interface{}) *LazySprintf {
	_ = "STUB: not implemented"
	return nil
}

func (l *LazySprintf) String() string { _ = "STUB: not implemented"; return "" }

type LazyBlockHash struct {
	block hashable
}

type hashable interface {
	Hash() cmtbytes.HexBytes
}

// NewLazyBlockHash defers block Hash until the Stringer interface is invoked.
// This is particularly useful for avoiding calling Sprintf when debugging is not
// active.
func NewLazyBlockHash(block hashable) *LazyBlockHash { _ = "STUB: not implemented"; return nil }

func (l *LazyBlockHash) String() string { _ = "STUB: not implemented"; return "" }
