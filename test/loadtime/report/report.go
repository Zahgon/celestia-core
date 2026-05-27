package report

import (
	"time"

	"github.com/gofrs/uuid"

	"github.com/cometbft/cometbft/types"
)

// BlockStore defines the set of methods needed by the report generator from
// CometBFT's store.Blockstore type. Using an interface allows for tests to
// more easily simulate the required behavior without having to use the more
// complex real API.
type BlockStore interface {
	Height() int64
	Base() int64
	LoadBlock(int64) *types.Block
}

// DataPoint contains the set of data collected for each transaction.
type DataPoint struct {
	Duration  time.Duration
	BlockTime time.Time
	Hash      []byte
}

// Report contains the data calculated from reading the timestamped transactions
// of each block found in the blockstore.
type Report struct {
	ID                      uuid.UUID
	Rate, Connections, Size uint64
	Max, Min, Avg, StdDev   time.Duration

	// NegativeCount is the number of negative durations encountered while
	// reading the transaction data. A negative duration means that
	// a transaction timestamp was greater than the timestamp of the block it
	// was included in and likely indicates an issue with the experimental
	// setup.
	NegativeCount int

	// All contains all data points gathered from all valid transactions.
	// The order of the contents of All is not guaranteed to be match the order of transactions
	// in the chain.
	All []DataPoint

	// used for calculating average during report creation.
	sum int64
}

type Reports struct {
	s map[uuid.UUID]Report
	l []Report

	// errorCount is the number of parsing errors encountered while reading the
	// transaction data. Parsing errors may occur if a transaction not generated
	// by the payload package is submitted to the chain.
	errorCount int
}

func (rs *Reports) List() []Report { _ = "STUB: not implemented"; return nil }

func (rs *Reports) ErrorCount() int { _ = "STUB: not implemented"; return 0 }

func (rs *Reports) addDataPoint(id uuid.UUID, l time.Duration, bt time.Time, hash []byte, conns, rate, size uint64) {
	_ = "STUB: not implemented"
	return
}

// Using an int64 here makes an assumption about the scale and quantity of the data we are processing.
// If all latencies were 2 seconds, we would need around 4 billion records to overflow this.
// We are therefore assuming that the data does not exceed these bounds.

func (rs *Reports) calculateAll() { _ = "STUB: not implemented"; return }

func (rs *Reports) addError() {
	_ = "STUB: not implemented"

	// GenerateFromBlockStore creates a Report using the data in the provided
	// BlockStore.
	return
}

func GenerateFromBlockStore(s BlockStore) (*Reports, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Deserializing to proto can be slow but does not depend on other data
// and can therefore be done in parallel.
// Deserializing in parallel does mean that the resulting data is
// not guaranteed to be delivered in the same order it was given to the
// worker pool.

// Data from two adjacent block are used here simultaneously,
// blocks of height H and H+1. The transactions of the block of
// height H are used with the timestamp from the block of height
// H+1. This is done because the timestamp from H+1 is calculated
// by using the precommits submitted at height H. The timestamp in
// block H+1 represents the time at which block H was committed.
//
// In the (very unlikely) event that the very last block of the
// chain contains payload transactions, those transactions will not
// be used in the latency calculations because the last block whose
// transactions are used is the block one before the last.

//nolint:staticcheck

func toFloat(in []DataPoint) []float64 { _ = "STUB: not implemented"; return nil }
