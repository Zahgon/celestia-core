package state

import (
	"errors"
)

type (
	ErrInvalidBlock error
	ErrProxyAppConn error

	ErrUnknownBlock struct {
		Height int64
	}

	ErrBlockHashMismatch struct {
		CoreHash []byte
		AppHash  []byte
		Height   int64
	}

	ErrAppBlockHeightTooHigh struct {
		CoreHeight int64
		AppHeight  int64
	}

	ErrAppBlockHeightTooLow struct {
		AppHeight int64
		StoreBase int64
	}

	ErrLastStateMismatch struct {
		Height int64
		Core   []byte
		App    []byte
	}

	ErrStateMismatch struct {
		Got      *State
		Expected *State
	}

	ErrNoValSetForHeight struct {
		Height int64
	}

	ErrNoConsensusParamsForHeight struct {
		Height int64
	}

	ErrNoABCIResponsesForHeight struct {
		Height int64
	}

	ErrABCIResponseResponseUnmarshalForHeight struct {
		Height int64
	}

	ErrABCIResponseCorruptedOrSpecChangeForHeight struct {
		Err    error
		Height int64
	}
)

func (e ErrUnknownBlock) Error() string { _ = "STUB: not implemented"; return "" }

func (e ErrBlockHashMismatch) Error() string { _ = "STUB: not implemented"; return "" }

func (e ErrAppBlockHeightTooHigh) Error() string { _ = "STUB: not implemented"; return "" }

func (e ErrAppBlockHeightTooLow) Error() string { _ = "STUB: not implemented"; return "" }

func (e ErrLastStateMismatch) Error() string { _ = "STUB: not implemented"; return "" }

func (e ErrStateMismatch) Error() string { _ = "STUB: not implemented"; return "" }

func (e ErrNoValSetForHeight) Error() string { _ = "STUB: not implemented"; return "" }

func (e ErrNoConsensusParamsForHeight) Error() string { _ = "STUB: not implemented"; return "" }

func (e ErrNoABCIResponsesForHeight) Error() string { _ = "STUB: not implemented"; return "" }

func (e ErrABCIResponseResponseUnmarshalForHeight) Error() string {
	_ = "STUB: not implemented"
	return ""
}

func (e ErrABCIResponseCorruptedOrSpecChangeForHeight) Error() string {
	_ = "STUB: not implemented"
	return ""
}

func (e ErrABCIResponseCorruptedOrSpecChangeForHeight) Unwrap() error {
	_ = "STUB: not implemented"
	return nil
}

var ErrFinalizeBlockResponsesNotPersisted = errors.New("node is not persisting finalize block responses")
