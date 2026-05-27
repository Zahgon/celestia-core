package mempool

import (
	"container/list"

	cmtsync "github.com/cometbft/cometbft/libs/sync"
	"github.com/cometbft/cometbft/types"
)

// TxCache defines an interface for raw transaction caching in a mempool.
// Currently, a TxCache does not allow direct reading or getting of transaction
// values. A TxCache is used primarily to push transactions and removing
// transactions. Pushing via Push returns a boolean telling the caller if the
// transaction already exists in the cache or not.
type TxCache interface {
	// Reset resets the cache to an empty state.
	Reset()

	// Push adds the given raw transaction to the cache and returns true if it was
	// newly added. Otherwise, it returns false.
	Push(key types.TxKey) bool

	// Remove removes the given raw transaction from the cache.
	Remove(key types.TxKey)

	// Has reports whether tx is present in the cache. Checking for presence is
	// not treated as an access of the value.
	Has(key types.TxKey) bool

	// HasKey reports whether the given key is present in the cache.
	HasKey(key types.TxKey) bool
}

var _ TxCache = (*LRUTxCache)(nil)

// LRUTxCache maintains a thread-safe LRU cache of raw transactions. The cache
// only stores the hash of the raw transaction.
type LRUTxCache struct {
	mtx      cmtsync.Mutex
	size     int
	cacheMap map[types.TxKey]*list.Element
	list     *list.List
}

func NewLRUTxCache(cacheSize int) *LRUTxCache { _ = "STUB: not implemented"; return nil }

// GetList returns the underlying linked-list that backs the LRU cache. Note,
// this should be used for testing purposes only!
func (c *LRUTxCache) GetList() *list.List { _ = "STUB: not implemented"; return nil }

func (c *LRUTxCache) Reset() { _ = "STUB: not implemented"; return }

func (c *LRUTxCache) Push(key types.TxKey) bool { _ = "STUB: not implemented"; return false }

func (c *LRUTxCache) Remove(key types.TxKey) { _ = "STUB: not implemented"; return }

func (c *LRUTxCache) Has(key types.TxKey) bool { _ = "STUB: not implemented"; return false }

func (c *LRUTxCache) HasKey(key types.TxKey) bool { _ = "STUB: not implemented"; return false }

// NopTxCache defines a no-op raw transaction cache.
type NopTxCache struct{}

var _ TxCache = (*NopTxCache)(nil)

func (NopTxCache) Reset()                { _ = "STUB: not implemented"; return }
func (NopTxCache) Push(types.TxKey) bool { _ = "STUB: not implemented"; return false }
func (NopTxCache) Remove(types.TxKey)    { _ = "STUB: not implemented"; return }
func (NopTxCache) Has(types.TxKey) bool  { _ = "STUB: not implemented"; return false }
func (NopTxCache) HasKey(types.TxKey) bool {
	_ = "STUB: not implemented"

	// cacheEntry stores both the transaction key and error code
	return false
}

type cacheEntry struct {
	key  types.TxKey
	code uint32
	log  string
}

// RejectedTxCache is a cache of rejected transactions. It wraps LRUTxCache
// to store the error code for a transaction that has been rejected.
type RejectedTxCache struct {
	cache *LRUTxCache
}

// NewRejectedTxCache creates a new rejected tx cache.
func NewRejectedTxCache(cacheSize int) *RejectedTxCache { _ = "STUB: not implemented"; return nil }

// Reset resets the cache to an empty state.
func (c *RejectedTxCache) Reset() {
	_ = "STUB: not implemented"

	// Push adds a tx key and error code to the cache.
	return
}

func (c *RejectedTxCache) Push(key types.TxKey, code uint32, log string) bool {
	_ = "STUB: not implemented"
	return false
}

// Get returns the error code for a tx key if it exists in the cache.
func (c *RejectedTxCache) Get(key types.TxKey) (uint32, string, bool) {
	_ = "STUB: not implemented"
	return 0, "", false
}

// HasKey returns true if the tx key is present in the cache.
func (c *RejectedTxCache) HasKey(key types.TxKey) bool { _ = "STUB: not implemented"; return false }

// Remove removes a tx from the cache.
func (c *RejectedTxCache) Remove(key types.TxKey) { _ = "STUB: not implemented"; return }
