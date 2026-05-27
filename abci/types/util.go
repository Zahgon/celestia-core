package types

import (
	"sort"
)

//------------------------------------------------------------------------------

// ValidatorUpdates is a list of validators that implements the Sort interface
type ValidatorUpdates []ValidatorUpdate

var _ sort.Interface = (ValidatorUpdates)(nil)

// All these methods for ValidatorUpdates:
//    Len, Less and Swap
// are for ValidatorUpdates to implement sort.Interface
// which will be used by the sort package.
// See Issue https://github.com/tendermint/abci/issues/212

func (v ValidatorUpdates) Len() int {
	_ = "STUB: not implemented"

	// XXX: doesn't distinguish same validator with different power
	return 0
}

func (v ValidatorUpdates) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (v ValidatorUpdates) Swap(i, j int) { _ = "STUB: not implemented"; return }
