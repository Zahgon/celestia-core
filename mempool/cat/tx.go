package cat

import (
	"time"

	"github.com/cometbft/cometbft/types"
)

// wrappedTx defines a wrapper around a raw transaction with additional metadata
// that is used for indexing. With the exception of the map of peers who have
// seen this transaction, this struct should never be modified
type wrappedTx struct {
	// these fields are immutable
	tx            *types.CachedTx // the original transaction data
	height        int64           // height when this transaction was initially checked (for expiry)
	timestamp     time.Time       // time when transaction was entered (for TTL)
	gasWanted     int64           // app: gas required to execute this transaction
	priority      int64           // app: priority value for this transaction
	sender        []byte          // app: assigned sender label
	sequence      uint64          // app: sequence number for this transaction
	fromBroadcast bool            // true if submitted via local RPC broadcast
}

func newWrappedTx(tx *types.CachedTx, height, gasWanted, priority int64, sender []byte, sequence uint64, fromBroadcast bool) *wrappedTx {
	_ = "STUB: not implemented"
	return nil
}

// Size reports the size of the raw transaction in bytes.
func (w *wrappedTx) size() int64 { _ = "STUB: not implemented"; return 0 }

// key returns the underlying tx key.
func (w *wrappedTx) key() types.TxKey {
	_ = "STUB: not implemented"

	// txSet groups transactions from the same signer and carries an aggregated priority.
	// Transactions within a set are ordered by sequence (ascending). If sequences are
	// equal or unset, order by arrival timestamp.
	return *new(types.TxKey)
}

type txSet struct {
	signerKey string
	signer    []byte
	// this should be ordered by sequence (ascending)
	txs                []*wrappedTx
	aggregatedPriority int64
	bytes              int64
	firstTimestamp     time.Time
	firstHeight        int64
	// gas-weighted aggregation tracking
	totalGasWanted      int64
	weightedPrioritySum int64
}

func newTxSet(wtx ...*wrappedTx) *txSet { _ = "STUB: not implemented"; return nil }

// addTxToSet inserts wtx into the set maintaining sequence order (ascending).
// If sequence equal, order by timestamp.
func (set *txSet) addTxToSet(wtx *wrappedTx) { _ = "STUB: not implemented"; return }

// update gas-weighted aggregation

// removeTx removes the provided wrappedTx from the set and updates aggregation.
func (set *txSet) removeTx(wtx *wrappedTx) bool { _ = "STUB: not implemented"; return false }

// Remove the tx from the set

// If the set is empty, or the total gas wanted is zero
// set the aggregated priority to zero

// Recompute earliest timestamp and lowest height

// sliceTxsByBytesAndGas slices the transactions set into a possible subset
// that fits within the given bytes and gas constraints ordered by sequence
func (set *txSet) sliceTxsByBytesAndGas(numBytes int64, numGas int64) *txSet {
	_ = "STUB: not implemented"
	// check if we have no budget left
	return nil
}

// return empty set if no budget

// check if we have unlimited budget or enough budget for the whole set

// return full set if budget is sufficient

func (set *txSet) rawTxs() []*types.CachedTx { _ = "STUB: not implemented"; return nil }

func aggregatePriorityAcrossSets(txSets []*txSet) int64 { _ = "STUB: not implemented"; return 0 }
