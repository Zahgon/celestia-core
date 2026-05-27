package cat

import (
	"errors"
	"sync"
	"time"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/config"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cometbft/cometbft/libs/trace"
	"github.com/cometbft/cometbft/mempool"
	"github.com/cometbft/cometbft/proxy"
	"github.com/cometbft/cometbft/types"
)

// enforce compile-time satisfaction of the Mempool interface
var _ mempool.Mempool = (*TxPool)(nil)

var (
	ErrTxInMempool = errors.New("tx already exists in mempool")
)

// TxPoolOption sets an optional parameter on the TxPool.
type TxPoolOption func(*TxPool)

// TxPool implemements the Mempool interface and allows the application to
// set priority values on transactions in the CheckTx response. When selecting
// transactions to include in a block, higher-priority transactions are chosen
// first.  When evicting transactions from the mempool for size constraints,
// lower-priority transactions are evicted first. Transactions themselves are
// unordered (A map is used). They can be broadcast in an order different from
// the order to which transactions are entered. There is no guarantee when CheckTx
// passes that a transaction has been successfully broadcast to any of its peers.
//
// A TTL can be set to remove transactions after a period of time or a number
// of heights.
//
// A cache of rejectedTxs can be set in the mempool config. Transactions that
// are rejected because of `CheckTx` or other validity checks will be instantly
// rejected if they are seen again. Committed transactions are also added to
// this cache. This serves somewhat as replay protection but applications should
// implement something more comprehensive
type TxPool struct {
	// Immutable fields
	logger       log.Logger
	config       *config.MempoolConfig
	proxyAppConn proxy.AppConnMempool
	metrics      *mempool.Metrics
	traceClient  trace.Tracer

	// these values are modified once per height
	mtx                  sync.Mutex
	notifiedTxsAvailable bool
	txsAvailable         chan struct{} // one value sent per height when mempool is not empty
	preCheckFn           mempool.PreCheckFunc
	postCheckFn          mempool.PostCheckFunc
	height               int64     // the latest height passed to Update
	lastPurgeTime        time.Time // the last time we attempted to purge transactions via the TTL

	// Thread-safe cache of rejected transactions for quick look-up
	rejectedTxCache *mempool.RejectedTxCache
	// Thread-safe cache of evicted transactions for quick look-up
	evictedTxCache *mempool.LRUTxCache
	// Thread-safe list of transactions peers have seen that we have not yet seen
	seenByPeersSet *SeenTxSet

	// Store of wrapped transactions
	store *store

	// broadcastCh is an unbuffered channel of new transactions that need to
	// be broadcasted to peers. Only populated if `broadcast` in the config is enabled
	broadcastCh      chan *wrappedTx
	broadcastMtx     sync.Mutex
	txsToBeBroadcast []types.TxKey

	// heightSignal notifies listeners whenever the mempool height advances.
	heightSignal chan struct{}
}

// NewTxPool constructs a new, empty content addressable txpool at the specified
// initial height and using the given config and options.
func NewTxPool(
	logger log.Logger,
	cfg *config.MempoolConfig,
	proxyAppConn proxy.AppConnMempool,
	height int64,
	options ...TxPoolOption,
) *TxPool {
	_ = "STUB: not implemented"
	return nil
}

// WithPreCheck sets a filter for the mempool to reject a transaction if f(tx)
// returns an error. This is executed before CheckTx. It only applies to the
// first created block. After that, Update() overwrites the existing value.
func WithPreCheck(f mempool.PreCheckFunc) TxPoolOption {
	_ = "STUB: not implemented"
	return *new(TxPoolOption)
}

// WithPostCheck sets a filter for the mempool to reject a transaction if
// f(tx, resp) returns an error. This is executed after CheckTx. It only applies
// to the first created block. After that, Update overwrites the existing value.
func WithPostCheck(f mempool.PostCheckFunc) TxPoolOption {
	_ = "STUB: not implemented"
	return *new(TxPoolOption)
}

// WithMetrics sets the mempool's metrics collector.
func WithMetrics(metrics *mempool.Metrics) TxPoolOption {
	_ = "STUB: not implemented"
	return *new(TxPoolOption)
}

// WithTracer sets the mempool's trace client.
func WithTracer(tracer trace.Tracer) TxPoolOption {
	_ = "STUB: not implemented"
	return *new(TxPoolOption)
}

// Lock locks the mempool, no new transactions can be processed
func (txmp *TxPool) Lock() {
	_ = "STUB: not implemented"

	// Unlock unlocks the mempool
	return
}

func (txmp *TxPool) Unlock() {
	_ = "STUB: not implemented"

	// Size returns the number of valid transactions in the mempool. It is
	// thread-safe.
	return
}

func (txmp *TxPool) Size() int { _ = "STUB: not implemented"; return 0 }

// SizeBytes returns the total sum in bytes of all the valid transactions in the
// mempool. It is thread-safe.
func (txmp *TxPool) SizeBytes() int64 { _ = "STUB: not implemented"; return 0 }

// FlushAppConn executes FlushSync on the mempool's proxyAppConn.
//
// The caller must hold an exclusive mempool lock (by calling txmp.Lock) before
// calling FlushAppConn.
func (txmp *TxPool) FlushAppConn() error { _ = "STUB: not implemented"; return nil }

// EnableTxsAvailable enables the mempool to trigger events when transactions
// are available on a block by block basis.
func (txmp *TxPool) EnableTxsAvailable() { _ = "STUB: not implemented"; return }

// TxsAvailable returns a channel which fires once for every height, and only
// when transactions are available in the mempool. It is thread-safe.
func (txmp *TxPool) TxsAvailable() <-chan struct{} { _ = "STUB: not implemented"; return nil }

// Height returns the latest height that the mempool is at
func (txmp *TxPool) Height() int64 { _ = "STUB: not implemented"; return 0 }

// HeightSignal returns a channel that fires whenever the mempool height advances.
func (txmp *TxPool) HeightSignal() <-chan struct{} { _ = "STUB: not implemented"; return nil }

// Has returns true if the transaction is currently in the mempool
func (txmp *TxPool) Has(txKey types.TxKey) bool { _ = "STUB: not implemented"; return false }

// Get retrieves a transaction based on the key.
// Deprecated: use GetTxByKey instead.
func (txmp *TxPool) Get(txKey types.TxKey) (*types.CachedTx, bool) {
	_ = "STUB: not implemented"
	return nil, false

	// GetTxByKey retrieves a transaction based on the key. It returns a bool
	// indicating whether transaction was found in the cache.
}

func (txmp *TxPool) GetTxByKey(txKey types.TxKey) (*types.CachedTx, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// WasRecentlyEvicted returns a bool indicating whether the transaction with
// the specified key was recently evicted and is currently within the cache.
func (txmp *TxPool) WasRecentlyEvicted(txKey types.TxKey) bool {
	_ = "STUB: not implemented"
	return false
}

// WasRecentlyRejected returns a bool indicating if the transaction was recently rejected and is
// currently within the cache. It also returns the rejection code and log.
func (txmp *TxPool) WasRecentlyRejected(txKey types.TxKey) (bool, uint32, string) {
	_ = "STUB: not implemented"
	return false, 0, ""
}

// CheckTx adds the given transaction to the mempool if it fits and passes the
// application's ABCI CheckTx method. This should be viewed as the entry method for new transactions
// into the network. In practice this happens via an RPC endpoint
func (txmp *TxPool) CheckTx(tx types.Tx, cb func(*abci.ResponseCheckTx), txInfo mempool.TxInfo) error {
	_ = "STUB: not implemented"
	// Reject transactions in excess of the configured maximum transaction size.
	return nil
}

// This is a new transaction that we haven't seen before. Verify it against the app and attempt
// to add it to the transaction pool.

// call the callback if it is set

// push to the broadcast queue that a new transaction is ready

// next is used by the reactor to get the next transaction to broadcast
// to all other peers.
func (txmp *TxPool) next() <-chan *wrappedTx { _ = "STUB: not implemented"; return nil }

// markToBeBroadcast marks a transaction to be broadcasted to peers.
// This should never block so we use a map to create an unbounded queue
// of transactions that need to be gossiped.
func (txmp *TxPool) markToBeBroadcast(key types.TxKey) { _ = "STUB: not implemented"; return }

// TryAddNewTx attempts to add a tx that has not already been seen before. It first marks it as seen
// to avoid races with the same tx. It then call `CheckTx` so that the application can validate it.
// If it passes `CheckTx`, the new transaction is added to the mempool as long as it has
// sufficient priority and space else if evicted it will return an error
func (txmp *TxPool) TryAddNewTx(tx *types.CachedTx, key types.TxKey, txInfo mempool.TxInfo) (*abci.ResponseCheckTx, error) {
	_ = "STUB: not implemented"
	// First check the cache to see if we can conclude early. We may have already seen and processed
	// the transaction if:
	// - We are connected to nodes running v0 or v1 which simply flood the network
	// - If a client submits a transaction to multiple nodes (via RPC)
	// - We send multiple requests and the first peer eventually responds after the second peer has already provided the tx
	return nil, nil
}

// The peer has sent us a transaction that we have already seen

// reserve the key

// If a precheck hook is defined, call it before invoking the application.

// Early exit if the proxy connection has an error.

// Invoke an ABCI CheckTx for this transaction.

// Create wrapped tx

// Perform the post check

// Now we consider the transaction to be valid. Once a transaction is valid, it
// can only become invalid if recheckTx is enabled and RecheckTx returns a non zero code

// RemoveTxByKey removes the transaction with the specified key from the
// mempool. It adds it to the rejectedTxCache so it will not be added again
func (txmp *TxPool) RemoveTxByKey(txKey types.TxKey) error { _ = "STUB: not implemented"; return nil }

func (txmp *TxPool) removeTxByKey(txKey types.TxKey) { _ = "STUB: not implemented"; return }

// Flush purges the contents of the mempool and the cache, leaving both empty.
// The current height is not modified by this operation.
func (txmp *TxPool) Flush() {
	_ = "STUB: not implemented"
	// Remove all the transactions in the list explicitly, so that the sizes
	// and indexes get updated properly.
	return
}

// PeerHasTx marks that the transaction has been seen by a peer.
func (txmp *TxPool) PeerHasTx(peer uint16, txKey types.TxKey) { _ = "STUB: not implemented"; return }

// ReapMaxBytesMaxGas returns a slice of valid transactions that fit within the
// size and gas constraints. The results are ordered by decreasing priority,
// with ties broken by increasing order of arrival. Transactions are also
// grouped together by signer in order of sequence to preserve sequence ordering within a signer.
//
// # Reaping transactions does not remove them from the mempool
//
// If maxBytes < 0, no limit is set on the total size in bytes.
// If maxGas < 0, no limit is set on the total gas cost.
//
// If the mempool is empty or has no transactions fitting within the given
// constraints, the result will also be empty.
func (txmp *TxPool) ReapMaxBytesMaxGas(maxBytes, maxGas int64) []*types.CachedTx {
	_ = "STUB: not implemented"
	return nil
}

// if the next transaction set can not fit, then we need to break down the indidual transactions
// and work out the residual set that has the highest accumulative priority and append that

// this function iterates over remaining txSets starting at all possible offsets i.e. from
// the first, second, third etc. transaction and working out what permutation given the remaining
// available bytes and gas has the highest aggregated priority
func (txmp *TxPool) determineLeftoverTxs(txSets []*txSet, remainingBytes, remainingGas int64) []*types.CachedTx {
	_ = "STUB: not implemented"
	return nil
}

// getAggregatedPriorityAndTxs return the first n txs in the provided txSets that fit within
// the remaining bytes and gas and returns the aggregated priority of the resulting txs
func (txmp *TxPool) getAggregatedPriorityAndTxs(txSets []*txSet, remainingBytes, remainingGas int64) (int64, []*types.CachedTx) {
	_ = "STUB: not implemented"
	return 0, nil
}

// if the full tx set is not returned, that indicates there are no more bytes or gas left

func flattenTxSets(txSets []*txSet) []*types.CachedTx { _ = "STUB: not implemented"; return nil }

// ReapMaxTxs returns up to max transactions from the mempool. The results are
// ordered by decreasing priority with ties broken by increasing order of
// arrival. Transactions are also ordered to preserve sequence numbers within a signer.
// Reaping transactions does not remove them from the mempool.
//
// If max < 0, all transactions in the mempool are reaped.
//
// The result may have fewer than max elements (possibly zero) if the mempool
// does not have that many transactions available.
func (txmp *TxPool) ReapMaxTxs(max int) []*types.CachedTx { _ = "STUB: not implemented"; return nil }

// Update removes all the given transactions from the mempool and the cache,
// and updates the current block height. The blockTxs and deliverTxResponses
// must have the same length with each response corresponding to the tx at the
// same offset.
//
// If the configuration enables recheck, Update sends each remaining
// transaction after removing blockTxs to the ABCI CheckTx method.  Any
// transactions marked as invalid during recheck are also removed.
//
// The caller must hold an exclusive mempool lock (by calling txmp.Lock) before
// calling Update.
func (txmp *TxPool) Update(
	blockHeight int64,
	blockTxs []*types.CachedTx,
	deliverTxResponses []*abci.ExecTxResult,
	newPreFn mempool.PreCheckFunc,
	newPostFn mempool.PostCheckFunc,
) error {
	_ = "STUB: not implemented"
	// Safety check: Transactions and responses must match in number.
	return nil
}

// Regardless of success, remove the transaction from the mempool.

// Record user-submitted tx latency = block receive time - tx receive time

// Write trace for confirmed transaction if we have the wrapped tx info

// Purge expired transactions based on TTL. This is the only place where
// TTL purging should happen to maintain checkTxState consistency.

// If there any uncommitted transactions left in the mempool, we either
// initiate re-CheckTx per remaining transaction or notify that remaining
// transactions are left.

// addNewTransaction handles the ABCI CheckTx response for the first time a
// transaction is added to the mempool.  A recheck after a block is committed
// goes to handleRecheckResult.
//
// If either the application rejected the transaction or a post-check hook is
// defined and rejects the transaction, it is discarded.
//
// Otherwise, if the mempool is full, check for lower-priority transactions
// that can be evicted to make room for the new one. If no such transactions
// exist, this transaction is logged and dropped; otherwise the selected
// transactions are evicted.
//
// Finally, the new transaction is added and size stats updated.
func (txmp *TxPool) addNewTransaction(wtx *wrappedTx) error {
	_ = "STUB: not implemented"
	// At this point the application has ruled the transaction valid, but the
	// mempool might be full. If so, find the lowest-priority items with lower
	// priority than the application assigned to this new one, and evict as many
	// of them as necessary to make room for tx. If no such items exist, we
	// discard tx.
	return nil
}

// Set-level eviction: aggregate by signer and compare against the new set's aggregated priority

// If there are no suitable eviction candidates, or the total size is insufficient, drop the new one.

// Sort lowest aggregated-priority sets first; ties: newer sets first to preserve FIFO within set groups

// Evict as many sets as needed to make room for the incoming tx

// Iterate in reverse order removing the higher sequence numbers first

func (txmp *TxPool) evictTx(wtx *wrappedTx) { _ = "STUB: not implemented"; return }

// handleRecheckResult handles the responses from ABCI CheckTx calls issued
// during the recheck phase of a block Update.  It removes any transactions
// invalidated by the application.
//
// This method is NOT executed for the initial CheckTx on a new transaction;
// that case is handled by addNewTransaction instead.
func (txmp *TxPool) handleRecheckResult(wtx *wrappedTx, checkTxRes *abci.ResponseCheckTx) {
	_ = "STUB: not implemented"
	return
}

// If a postcheck hook is defined, call it before checking the result.

// Note that we do not update the transaction with any of the values returned in
// recheck tx

// N.B. Size of mempool did not change

// recheckTransactions initiates re-CheckTx ABCI calls for all the transactions
// currently in the mempool. It reports the number of recheck calls that were
// successfully initiated.
//
// Precondition: The mempool is not empty.
// The caller must hold txmp.mtx exclusively.
func (txmp *TxPool) recheckTransactions() { _ = "STUB: not implemented"; return }

// Get all transactions currently in the mempool requiring recheck.
// This avoids holding the store lock during the CheckTx calls which could
// cause a deadlock when handleRecheckResult tries to modify the store.

// Issue CheckTx calls for each remaining transaction, and when all the
// rechecks are complete signal watchers that transactions may be available.

// The response for this CheckTx is handled by the default recheckTxCallback.

// When recheck is complete, trigger a notification for more transactions.

// availableBytes returns the number of bytes available in the mempool.
func (txmp *TxPool) availableBytes() int64 { _ = "STUB: not implemented"; return 0 }

// canAddTx returns an error if we cannot insert the provided *wrappedTx into
// the mempool due to mempool configured constraints. Otherwise, nil is
// returned and the transaction can be inserted into the mempool.
func (txmp *TxPool) canAddTx(size int64) bool { _ = "STUB: not implemented"; return false }

// purgeExpiredTxs removes all transactions from the mempool that have exceeded
// their respective height or time-based limits as of the given blockHeight.
// Transactions removed by this operation are not removed from the rejectedTxCache.
func (txmp *TxPool) purgeExpiredTxs(blockHeight int64) { _ = "STUB: not implemented"; return }

// nothing to do

// Add the purged transactions to the evicted cache

func (txmp *TxPool) notifyTxsAvailable() { _ = "STUB: not implemented"; return }

// nothing to do

// channel cap is 1, so this will send once

func (txmp *TxPool) preCheck(tx *types.CachedTx) error { _ = "STUB: not implemented"; return nil }

func (txmp *TxPool) postCheck(tx *types.CachedTx, res *abci.ResponseCheckTx) error {
	_ = "STUB: not implemented"
	return nil
}
