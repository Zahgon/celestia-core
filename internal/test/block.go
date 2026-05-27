package test

import (
	"testing"
	"time"

	"github.com/cometbft/cometbft/types"
)

const (
	DefaultTestChainID = "test-chain"
)

var (
	DefaultTestTime = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
)

func RandomAddress() []byte { _ = "STUB: not implemented"; return nil }

func RandomHash() []byte { _ = "STUB: not implemented"; return nil }

func MakeBlockID() types.BlockID { _ = "STUB: not implemented"; return *new(types.BlockID) }

func MakeBlockIDWithHash(hash []byte) types.BlockID {
	_ = "STUB: not implemented"
	return *new(types.BlockID)
}

// MakeHeader fills the rest of the contents of the header such that it passes
// validate basic
func MakeHeader(t *testing.T, h *types.Header) *types.Header { _ = "STUB: not implemented"; return nil }
