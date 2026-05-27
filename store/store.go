package store

import (
	"sync"
	"sync/atomic"

	"github.com/cosmos/gogoproto/proto"
	lru "github.com/hashicorp/golang-lru/v2"

	dbm "github.com/cometbft/cometbft-db"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/libs/log"
	cmtsync "github.com/cometbft/cometbft/libs/sync"
	cmtstore "github.com/cometbft/cometbft/proto/tendermint/store"
	sm "github.com/cometbft/cometbft/state"
	"github.com/cometbft/cometbft/types"
)

// maxBlockPartsToBatch is set to 600 to batch a whole block. Now that we're switching to pebbleDB,
// setting this to 600 is more performant.
const maxBlockPartsToBatch = 600

/*
BlockStore is a simple low level store for blocks.

There are three types of information stored:
  - BlockMeta:   Meta information about each block
  - Block part:  Parts of each block, aggregated w/ PartSet
  - Commit:      The commit part of each block, for gossiping precommit votes

Currently the precommit signatures are duplicated in the Block parts as
well as the Commit.  In the future this may change, perhaps by moving
the Commit data outside the Block. (TODO)

The store can be assumed to contain all contiguous blocks between base and height (inclusive).

// NOTE: BlockStore methods will panic if they encounter errors
// deserializing loaded data, indicating probable corruption on disk.
*/
type BlockStore struct {
	db dbm.DB

	// mtx guards access to the struct fields listed below it. Although we rely on the database
	// to enforce fine-grained concurrency control for its data, we need to make sure that
	// no external observer can get data from the database that is not in sync with the fields below,
	// and vice-versa. Hence, when updating the fields below, we use the mutex to make sure
	// that the database is also up to date. This prevents any concurrent external access from
	// obtaining inconsistent data.
	// The only reason for keeping these fields in the struct is that the data
	// can't efficiently be queried from the database since the key encoding we use is not
	// lexicographically ordered (see https://github.com/tendermint/tendermint/issues/4567).
	mtx    cmtsync.RWMutex
	base   int64
	height int64

	seenCommitCache          *lru.Cache[int64, *types.Commit]
	blockCommitCache         *lru.Cache[int64, *types.Commit]
	blockExtendedCommitCache *lru.Cache[int64, *types.ExtendedCommit]

	// blocksDeleted, compact, compactionInterval, and compactionFrom are only
	// read/written from PruneBlocks, which has a single production caller (the
	// consensus goroutine via BlockExecutor.ApplyBlock), so no synchronization
	// is needed for these fields. compacting and compactionWg coordinate with
	// the background compaction goroutine and are intentionally lock-free.
	// compactionFrom holds the lowest height not yet covered by a range-scoped
	// forced compaction. It is initialized to base on construction and
	// advanced only on successful Compact.
	blocksDeleted      int64
	compact            bool
	compactionInterval int64
	compactionFrom     int64
	compacting         atomic.Bool
	compactionWg       sync.WaitGroup

	logger log.Logger
}

type BlockStoreOption func(*BlockStore)

// WithCompaction sets the compaction parameters.
func WithCompaction(compact bool, compactionInterval int64) BlockStoreOption {
	_ = "STUB: not implemented"
	return *new(BlockStoreOption)
}

// WithLogger sets the logger used by the BlockStore.
func WithLogger(logger log.Logger) BlockStoreOption {
	_ = "STUB: not implemented"
	return *new(BlockStoreOption)
}

// NewBlockStore returns a new BlockStore with the given DB,
// initialized to the last height that was committed to the DB.
func NewBlockStore(db dbm.DB, options ...BlockStoreOption) *BlockStore {
	_ = "STUB: not implemented"
	return nil
}

func (bs *BlockStore) addCaches() {
	_ = "STUB: not implemented"

	// err can only occur if the argument is non-positive, so is impossible in context.
	return
}

func (bs *BlockStore) IsEmpty() bool { _ = "STUB: not implemented"; return false }

// Base returns the first known contiguous block height, or 0 for empty block stores.
func (bs *BlockStore) Base() int64 { _ = "STUB: not implemented"; return 0 }

// Height returns the last known contiguous block height, or 0 for empty block stores.
func (bs *BlockStore) Height() int64 { _ = "STUB: not implemented"; return 0 }

// Size returns the number of blocks in the block store.
func (bs *BlockStore) Size() int64 { _ = "STUB: not implemented"; return 0 }

// LoadBase atomically loads the base block meta, or returns nil if no base is found.
func (bs *BlockStore) LoadBaseMeta() *types.BlockMeta { _ = "STUB: not implemented"; return nil }

// LoadPartSet returns the partset for a given height.
func (bs *BlockStore) LoadPartSet(height int64) (*types.PartSet, *types.BlockMeta, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// LoadBlock returns the block with the given height.
// If no block is found for that height, it returns nil.
func (bs *BlockStore) LoadBlock(height int64) *types.Block { _ = "STUB: not implemented"; return nil }

// If the part is missing (e.g. since it has been deleted after we
// loaded the block meta) we consider the whole block to be missing.

// NOTE: The existence of meta should imply the existence of the
// block. So, make sure meta is only saved after blocks are saved.

// LoadBlockByHash returns the block with the given hash.
// If no block is found for that hash, it returns nil.
// Panics if it fails to parse height associated with the given hash.
func (bs *BlockStore) LoadBlockByHash(hash []byte) *types.Block {
	_ = "STUB: not implemented"
	return nil
}

// LoadBlockPart returns the Part at the given index
// from the block at the given height.
// If no part is found for the given height and index, it returns nil.
func (bs *BlockStore) LoadBlockPart(height int64, index int) *types.Part {
	_ = "STUB: not implemented"
	return nil
}

// LoadBlockMeta returns the BlockMeta for the given height.
// If no block is found for the given height, it returns nil.
func (bs *BlockStore) LoadBlockMeta(height int64) *types.BlockMeta {
	_ = "STUB: not implemented"
	return nil
}

// LoadBlockMetaByHash returns the blockmeta who's header corresponds to the given
// hash. If none is found, returns nil.
func (bs *BlockStore) LoadBlockMetaByHash(hash []byte) *types.BlockMeta {
	_ = "STUB: not implemented"
	return nil
}

// LoadBlockCommit returns the Commit for the given height.
// This commit consists of the +2/3 and other Precommit-votes for block at `height`,
// and it comes from the block.LastCommit for `height+1`.
// If no commit is found for the given height, it returns nil.
func (bs *BlockStore) LoadBlockCommit(height int64) *types.Commit {
	_ = "STUB: not implemented"
	return nil
}

// LoadExtendedCommit returns the ExtendedCommit for the given height.
// The extended commit is not guaranteed to contain the same +2/3 precommits data
// as the commit in the block.
func (bs *BlockStore) LoadBlockExtendedCommit(height int64) *types.ExtendedCommit {
	_ = "STUB: not implemented"
	return nil
}

// LoadSeenCommit returns the locally seen Commit for the given height.
// This is useful when we've seen a commit, but there has not yet been
// a new block at `height + 1` that includes this commit in its block.LastCommit.
func (bs *BlockStore) LoadSeenCommit(height int64) *types.Commit {
	_ = "STUB: not implemented"
	return nil
}

// PruneBlocks removes block up to (but not including) a height. It returns number of blocks pruned and the evidence retain height - the height at which data needed to prove evidence must not be removed.
func (bs *BlockStore) PruneBlocks(height int64, state sm.State) (uint64, int64, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

// We can't trust batches to be atomic, so update base first to make sure noone
// tries to access missing blocks.

// assume already deleted

// block may be nil if parts were deleted by a previous pruning pass
// that preserved meta for evidence. Skip tx hash cleanup since those
// tx hashes would have been deleted in that previous pass.

// This logic is in place to protect data that proves malicious behavior.
// If the height is within the evidence age, we continue to persist the header and commit data.

// if height is beyond the evidence point we dont delete the header

// if height is beyond the evidence point we dont delete the commit data

// flush every 1000 blocks to avoid batches becoming too large

// triggerCompactionAsync launches a background compaction over the height
// range that has been pruned since the last successful compaction
// ([compactionFrom, retainHeight)). If a compaction is already in flight, it
// logs and returns; the caller is responsible for resetting the deletion
// counter regardless. The goroutine is tracked by compactionWg so Close can
// wait for it.
//
// Range scoping is applied per height-keyed key family. Hash-keyed families
// (BH:, TH:) are left to pebble's natural background compaction.
func (bs *BlockStore) triggerCompactionAsync(retainHeight int64) { _ = "STUB: not implemented"; return }

// Single-writer: the next trigger has to wait for `compacting` to
// clear, and we set this before releasing it.

// compactBlockStoreRange issues one Compact call per height-keyed key family
// for the byte range [prefix+from, prefix+to). Heights are encoded as
// decimal ASCII, so the byte range matches the integer range only when from
// and to share a digit count; at digit-count boundaries some pruned heights
// fall outside the range and are left to pebble's background compaction.
// This is safe — pebble.Compact never drops live data; the range is only
// a hint for which sstables to rewrite.
func compactBlockStoreRange(db dbm.DB, from, to int64) error { _ = "STUB: not implemented"; return nil }

// SaveBlock persists the given block, blockParts, and seenCommit to the underlying db.
// blockParts: Must be parts of the block
// seenCommit: The +2/3 precommits that were seen which committed at height.
//
//	If all the nodes restart after committing a block,
//	we need this to reload the precommits to catch-up nodes to the
//	most recent height.  Otherwise they'd stall at H-1.
func (bs *BlockStore) SaveBlock(block *types.Block, blockParts *types.PartSet, seenCommit *types.Commit) {
	_ = "STUB: not implemented"
	return
}

// Save new BlockStoreState descriptor. This also flushes the database.

// SaveBlockWithExtendedCommit persists the given block, blockParts, and
// seenExtendedCommit to the underlying db. seenExtendedCommit is stored under
// two keys in the database: as the seenCommit and as the ExtendedCommit data for the
// height. This allows the vote extension data to be persisted for all blocks
// that are saved.
func (bs *BlockStore) SaveBlockWithExtendedCommit(block *types.Block, blockParts *types.PartSet, seenExtendedCommit *types.ExtendedCommit) {
	_ = "STUB: not implemented"
	return
}

// Save new BlockStoreState descriptor. This also flushes the database.

func (bs *BlockStore) saveBlockToBatch(
	block *types.Block,
	blockParts *types.PartSet,
	seenCommit *types.Commit,
	batch dbm.Batch,
) error {
	_ = "STUB: not implemented"
	return nil
}

// If the block is small, batch save the block parts. Otherwise, save the
// parts individually.

// Save block parts. This must be done before the block meta, since callers
// typically load the block meta first as an indication that the block exists
// and then go on to load block parts - we must make sure the block is
// complete as soon as the block meta is written.

// Save block meta

// Save block commit (duplicate and separate from the Block)

// Save seen commit (seen +2/3 precommits for block)
// NOTE: we can delete this at a later height

func (bs *BlockStore) saveBlockPart(height int64, index int, part *types.Part, batch dbm.Batch, saveBlockPartsToBatch bool) {
	_ = "STUB: not implemented"
	return
}

// Contract: the caller MUST have, at least, a read lock on `bs`.
func (bs *BlockStore) saveStateAndWriteDB(batch dbm.Batch, errMsg string) error {
	_ = "STUB: not implemented"
	return nil
}

// SaveSeenCommit saves a seen commit, used by e.g. the state sync reactor when bootstrapping node.
func (bs *BlockStore) SaveSeenCommit(height int64, seenCommit *types.Commit) error {
	_ = "STUB: not implemented"
	return nil
}

func (bs *BlockStore) Close() error { _ = "STUB: not implemented"; return nil }

//-----------------------------------------------------------------------------

func calcBlockMetaKey(height int64) []byte { _ = "STUB: not implemented"; return nil }

func calcBlockPartKey(height int64, partIndex int) []byte { _ = "STUB: not implemented"; return nil }

func calcBlockCommitKey(height int64) []byte { _ = "STUB: not implemented"; return nil }

func calcSeenCommitKey(height int64) []byte { _ = "STUB: not implemented"; return nil }

func calcExtCommitKey(height int64) []byte { _ = "STUB: not implemented"; return nil }

func calcBlockHashKey(hash []byte) []byte { _ = "STUB: not implemented"; return nil }

func calcTxHashKey(hash []byte) []byte { _ = "STUB: not implemented"; return nil }

//-----------------------------------------------------------------------------

var blockStoreKey = []byte("blockStore")

// SaveBlockStoreState persists the blockStore state to the database.
// deprecated: still present in this version for API compatibility
func SaveBlockStoreState(bsj *cmtstore.BlockStoreState, db dbm.DB) {
	_ = "STUB: not implemented"
	return
}

// SaveBlockStoreStateBatch persists the blockStore state to the database.
// It uses the DB batch passed as parameter
func SaveBlockStoreStateBatch(bsj *cmtstore.BlockStoreState, batch dbm.Batch) {
	_ = "STUB: not implemented"
	return
}

func saveBlockStoreStateBatchInternal(bsj *cmtstore.BlockStoreState, db dbm.DB, batch dbm.Batch) {
	_ = "STUB: not implemented"
	return
}

// LoadBlockStoreState returns the BlockStoreState as loaded from disk.
// If no BlockStoreState was previously persisted, it returns the zero value.
func LoadBlockStoreState(db dbm.DB) cmtstore.BlockStoreState {
	_ = "STUB: not implemented"
	return *new(cmtstore.BlockStoreState)
}

// Backwards compatibility with persisted data from before Base existed.

// mustEncode proto encodes a proto.message and panics if fails
func mustEncode(pb proto.Message) []byte { _ = "STUB: not implemented"; return nil }

//-----------------------------------------------------------------------------

// DeleteLatestBlock removes the block pointed to by height,
// lowering height by one.
func (bs *BlockStore) DeleteLatestBlock() error { _ = "STUB: not implemented"; return nil }

// delete what we can, skipping what's already missing, to ensure partial
// blocks get deleted fully.

// delete last, so as to not leave keys built on meta.BlockID dangling

// SaveTxInfo indexes the txs from the block with the given execution results.
// Only the error logs are saved for failed transactions.
func (bs *BlockStore) SaveTxInfo(block *types.Block, execTxRes []*abci.ExecTxResult) error {
	_ = "STUB: not implemented"
	return nil
}

// Create a new batch

// Batch and save txs from the block

//nolint:gosec

// Set error log for failed txs

// Write the batch to the db

// LoadTxInfo loads the TxInfo from disk given its hash.
func (bs *BlockStore) LoadTxInfo(txHash []byte) *cmtstore.TxInfo {
	_ = "STUB: not implemented"
	return nil
}
