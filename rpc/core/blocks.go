package core

import (
	"github.com/cometbft/cometbft/crypto/merkle"
	"github.com/cometbft/cometbft/libs/bytes"
	ctypes "github.com/cometbft/cometbft/rpc/core/types"
	rpctypes "github.com/cometbft/cometbft/rpc/jsonrpc/types"
)

const (
	asc  = "asc"
	desc = "desc"
)

// BlockchainInfo gets block headers for minHeight <= height <= maxHeight.
//
// If maxHeight does not yet exist, blocks up to the current height will be
// returned. If minHeight does not exist (due to pruning), earliest existing
// height will be used.
//
// At most 20 items will be returned. Block headers are returned in descending
// order (highest first).
//
// More: https://docs.cometbft.com/v0.38.x/rpc/#/Info/blockchain
func (env *Environment) BlockchainInfo(
	_ *rpctypes.Context,
	minHeight, maxHeight int64,
) (*ctypes.ResultBlockchainInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// error if either min or max are negative or min > max
// if 0, use blockstore base for min, latest block height for max
// enforce limit.
func filterMinMax(base, height, min, max, limit int64) (int64, int64, error) {
	_ = "STUB: not implemented"
	// filter negatives
	return 0, 0, nil
}

// adjust for default values

// limit max to the height

// limit min to the base

// limit min to within `limit` of max
// so the total number of blocks returned will be `limit`

// Header gets block header at a given height.
// If no height is provided, it will fetch the latest header.
// More: https://docs.cometbft.com/v0.38.x/rpc/#/Info/header
func (env *Environment) Header(_ *rpctypes.Context, heightPtr *int64) (*ctypes.ResultHeader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// HeaderByHash gets header by hash.
// More: https://docs.cometbft.com/v0.38.x/rpc/#/Info/header_by_hash
func (env *Environment) HeaderByHash(_ *rpctypes.Context, hash bytes.HexBytes) (*ctypes.ResultHeader, error) {
	_ = "STUB: not implemented"
	// N.B. The hash parameter is HexBytes so that the reflective parameter
	// decoding logic in the HTTP service will correctly translate from JSON.
	// See https://github.com/tendermint/tendermint/issues/6802 for context.
	return nil, nil
}

// Block gets block at a given height.
// If no height is provided, it will fetch the latest block.
// More: https://docs.cometbft.com/v0.38.x/rpc/#/Info/block
func (env *Environment) Block(_ *rpctypes.Context, heightPtr *int64) (*ctypes.ResultBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BlockByHash gets block by hash.
// More: https://docs.cometbft.com/v0.38.x/rpc/#/Info/block_by_hash
func (env *Environment) BlockByHash(_ *rpctypes.Context, hash []byte) (*ctypes.ResultBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If block is not nil, then blockMeta can't be nil.

// Commit gets block commit at a given height.
// If no height is provided, it will fetch the commit for the latest block.
// More: https://docs.cometbft.com/v0.38.x/rpc/#/Info/commit
func (env *Environment) Commit(_ *rpctypes.Context, heightPtr *int64) (*ctypes.ResultCommit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If the next block has not been committed yet,
// use a non-canonical commit

// Return the canonical commit (comes from the block at height+1)

// BlockResults gets ABCIResults at a given height.
// If no height is provided, it will fetch results for the latest block.
//
// Results are for the height of the block containing the txs.
// Thus response.results.deliver_tx[5] is the results of executing
// getBlock(h).Txs[5]
// More: https://docs.cometbft.com/v0.38.x/rpc/#/Info/block_results
func (env *Environment) BlockResults(_ *rpctypes.Context, heightPtr *int64) (*ctypes.ResultBlockResults, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BlockSearch searches for a paginated set of blocks matching
// FinalizeBlock event search criteria.
//
// Deprecated: The block_search endpoint is deprecated and will be removed in a future release.
func (env *Environment) BlockSearch(
	ctx *rpctypes.Context,
	query string,
	pagePtr, perPagePtr *int,
	orderBy string,
) (*ctypes.ResultBlockSearch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// skip if block indexing is disabled

// sort results (must be done before pagination)

// paginate results

// SignedBlock fetches the set of transactions at a specified height and all the relevant
// data to verify the transactions (i.e. using light client verification).
func (env *Environment) SignedBlock(ctx *rpctypes.Context, heightPtr *int64) (*ctypes.ResultSignedBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DataCommitment collects the data roots over a provided ordered range of blocks,
// and then creates a new Merkle root of those data roots. The range is end exclusive.
func (env *Environment) DataCommitment(ctx *rpctypes.Context, start, end uint64) (*ctypes.ResultDataCommitment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create data commitment

// padBytes Pad bytes to given length
func padBytes(byt []byte, length int) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// To32PaddedHexBytes takes a number and returns its hex representation padded to 32 bytes.
// Used to mimic the result of `abi.encode(number)` in Ethereum.
func To32PaddedHexBytes(number uint64) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Make sure hex representation has even length.
// The `strconv.FormatUint` can return odd length hex encodings.
// For example, `strconv.FormatUint(10, 16)` returns `a`.
// Thus, we need to pad it.

// DataRootTuple contains the data that will be used to create the QGB commitments.
// The commitments will be signed by orchestrators and submitted to an EVM chain via a relayer.
// For more information: https://github.com/celestiaorg/quantum-gravity-bridge/blob/master/src/DataRootTuple.sol
type DataRootTuple struct {
	height   uint64
	dataRoot [32]byte
}

// EncodeDataRootTuple takes a height and a data root, and returns the equivalent of
// `abi.encode(...)` in Ethereum.
// The encoded type is a DataRootTuple, which has the following ABI:
//
//	{
//	  "components":[
//	     {
//	        "internalType":"uint256",
//	        "name":"height",
//	        "type":"uint256"
//	     },
//	     {
//	        "internalType":"bytes32",
//	        "name":"dataRoot",
//	        "type":"bytes32"
//	     },
//	     {
//	        "internalType":"structDataRootTuple",
//	        "name":"_tuple",
//	        "type":"tuple"
//	     }
//	  ]
//	}
//
// padding the hex representation of the height padded to 32 bytes concatenated to the data root.
// For more information, refer to:
// https://github.com/celestiaorg/quantum-gravity-bridge/blob/master/src/DataRootTuple.sol
func EncodeDataRootTuple(height uint64, dataRoot [32]byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// dataCommitmentBlocksLimit The maximum number of blocks to be used to create a data commitment.
// It's a local parameter to protect the API from creating unnecessarily large commitments.
const dataCommitmentBlocksLimit = 10_000 // ~33 hours of blocks assuming 12-second blocks.

// validateDataCommitmentRange runs basic checks on the asc sorted list of
// heights that will be used subsequently in generating data commitments over
// the defined set of heights.
func (env *Environment) validateDataCommitmentRange(start uint64, end uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// the data commitment range is end exclusive
//nolint:gosec

// hashDataRootTuples hashes a list of blocks data root tuples, i.e. height, data root and square size,
// then returns their merkle root.
func hashDataRootTuples(tuples []DataRootTuple) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// validateDataRootInclusionProofRequest validates the request to generate a data root
// inclusion proof.
func (env *Environment) validateDataRootInclusionProofRequest(height uint64, start uint64, end uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// proveDataRootTuples returns the merkle inclusion proof for a height.
func (env *Environment) proveDataRootTuples(tuples []DataRootTuple, height int64) (*merkle.Proof, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:gosec

// fetchDataRootTuples takes an end exclusive range of heights and fetches its
// corresponding data root tuples.
func (env *Environment) fetchDataRootTuples(start, end uint64) ([]DataRootTuple, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:gosec

//nolint:gosec

// DataRootInclusionProof creates an inclusion proof for the data root of block
// height `height` in the set of blocks defined by `start` and `end`. The range
// is end exclusive.
func (env *Environment) DataRootInclusionProof(
	ctx *rpctypes.Context,
	height int64,
	start,
	end uint64,
) (*ctypes.ResultDataRootInclusionProof, error) {
	_ = "STUB: not implemented"
	//nolint:gosec
	return nil, nil
}

func (env *Environment) GenerateDataRootInclusionProof(height int64, start, end uint64) (*merkle.Proof, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
