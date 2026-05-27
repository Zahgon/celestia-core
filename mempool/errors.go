package mempool

import (
	"errors"
)

// ErrTxNotFound is returned to the client if tx is not found in mempool
var ErrTxNotFound = errors.New("transaction not found in mempool")

// ErrTxInCache is returned to the client if we saw tx earlier
var ErrTxInCache = errors.New("tx already exists in cache")

// ErrRecheckFull is returned when checking if the mempool is full and
// rechecking is still in progress after a new block was committed.
var ErrRecheckFull = errors.New("mempool is still rechecking after a new committed block, so it is considered as full")

// ErrTxTooLarge defines an error when a transaction is too big to be sent in a
// message to other peers.
type ErrTxTooLarge struct {
	Max    int
	Actual int
}

func (e ErrTxTooLarge) Error() string { _ = "STUB: not implemented"; return "" }

// ErrMempoolIsFull defines an error where CometBFT and the application cannot
// handle that much load.
type ErrMempoolIsFull struct {
	NumTxs      int
	MaxTxs      int
	TxsBytes    int64
	MaxTxsBytes int64
	RecheckFull bool
}

func (e ErrMempoolIsFull) Error() string { _ = "STUB: not implemented"; return "" }

// ErrPreCheck defines an error where a transaction fails a pre-check.
type ErrPreCheck struct {
	Err error
}

func (e ErrPreCheck) Error() string { _ = "STUB: not implemented"; return "" }

func (e ErrPreCheck) Unwrap() error {
	_ = "STUB: not implemented"

	// IsPreCheckError returns true if err is due to pre check failure.
	return nil
}

func IsPreCheckError(err error) bool { _ = "STUB: not implemented"; return false }

type ErrAppConnMempool struct {
	Err error
}

func (e ErrAppConnMempool) Error() string { _ = "STUB: not implemented"; return "" }

func (e ErrAppConnMempool) Unwrap() error { _ = "STUB: not implemented"; return nil }

type ErrFlushAppConn struct {
	Err error
}

func (e ErrFlushAppConn) Error() string { _ = "STUB: not implemented"; return "" }

func (e ErrFlushAppConn) Unwrap() error { _ = "STUB: not implemented"; return nil }
