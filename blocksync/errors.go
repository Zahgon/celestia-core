package blocksync

import (
	"errors"

	"github.com/cosmos/gogoproto/proto"
)

var (
	// ErrNilMessage is returned when provided message is empty
	ErrNilMessage = errors.New("message cannot be nil")
)

// ErrInvalidBase is returned when peer informs of a status with invalid height
type ErrInvalidHeight struct {
	Height int64
	Reason string
}

func (e ErrInvalidHeight) Error() string { _ = "STUB: not implemented"; return "" }

// ErrInvalidBase is returned when peer informs of a status with invalid base
type ErrInvalidBase struct {
	Base   int64
	Reason string
}

func (e ErrInvalidBase) Error() string { _ = "STUB: not implemented"; return "" }

type ErrUnknownMessageType struct {
	Msg proto.Message
}

func (e ErrUnknownMessageType) Error() string { _ = "STUB: not implemented"; return "" }

type ErrReactorValidation struct {
	Err error
}

func (e ErrReactorValidation) Error() string { _ = "STUB: not implemented"; return "" }

func (e ErrReactorValidation) Unwrap() error { _ = "STUB: not implemented"; return nil }
