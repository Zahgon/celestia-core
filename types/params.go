package types

import (
	"time"

	"github.com/cometbft/cometbft/crypto/ed25519"
	"github.com/cometbft/cometbft/crypto/secp256k1"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
)

const (
	DefaultMaxBlockSizeBytes = 128 * 1024 * 1024
	// ReducedMaxBlockSizeBytes will be used when disabling the propagation reactor in emergency situations
	ReducedMaxBlockSizeBytes = 8 * 1024 * 1024

	// BlockPartSizeBytes is the size of one block part.
	BlockPartSizeBytes uint32 = 65536 // 64kB

	ABCIPubKeyTypeEd25519   = ed25519.KeyType
	ABCIPubKeyTypeSecp256k1 = secp256k1.KeyType
)

var ABCIPubKeyTypesToNames = map[string]string{
	ABCIPubKeyTypeEd25519:   ed25519.PubKeyName,
	ABCIPubKeyTypeSecp256k1: secp256k1.PubKeyName,
}

var (
	// MaxBlockSizeBytes is the maximum permitted size of the blocks.
	MaxBlockSizeBytes = DefaultMaxBlockSizeBytes

	// MaxBlockPartsCount is the maximum number of block parts.
	MaxBlockPartsCount = (uint32(MaxBlockSizeBytes) / BlockPartSizeBytes) + 1
)

// ConsensusParams contains consensus critical parameters that determine the
// validity of blocks.
type ConsensusParams struct {
	Block     BlockParams     `json:"block"`
	Evidence  EvidenceParams  `json:"evidence"`
	Validator ValidatorParams `json:"validator"`
	Version   VersionParams   `json:"version"`
	ABCI      ABCIParams      `json:"abci"`
}

// BlockParams define limits on the block size and gas plus minimum time
// between blocks.
type BlockParams struct {
	MaxBytes int64 `json:"max_bytes"`
	MaxGas   int64 `json:"max_gas"`
}

// EvidenceParams determine how we handle evidence of malfeasance.
type EvidenceParams struct {
	MaxAgeNumBlocks int64         `json:"max_age_num_blocks"` // only accept new evidence more recent than this
	MaxAgeDuration  time.Duration `json:"max_age_duration"`
	MaxBytes        int64         `json:"max_bytes"`
}

// ValidatorParams restrict the public key types validators can use.
// NOTE: uses ABCI pubkey naming, not Amino names.
type ValidatorParams struct {
	PubKeyTypes []string `json:"pub_key_types"`
}

type VersionParams struct {
	App uint64 `json:"app"`
}

// ABCIParams configure ABCI functionality specific to the Application Blockchain
// Interface.
type ABCIParams struct {
	VoteExtensionsEnableHeight int64 `json:"vote_extensions_enable_height"`
}

// VoteExtensionsEnabled returns true if vote extensions are enabled at height h
// and false otherwise.
func (a ABCIParams) VoteExtensionsEnabled(h int64) bool { _ = "STUB: not implemented"; return false }

// DefaultConsensusParams returns a default ConsensusParams.
func DefaultConsensusParams() *ConsensusParams { _ = "STUB: not implemented"; return nil }

// DefaultBlockParams returns a default BlockParams.
func DefaultBlockParams() BlockParams { _ = "STUB: not implemented"; return *new(BlockParams) }

// 21MB

// DefaultEvidenceParams returns a default EvidenceParams.
func DefaultEvidenceParams() EvidenceParams { _ = "STUB: not implemented"; return *new(EvidenceParams) }

// 27.8 hrs at 1block/s

// 1MB

// DefaultValidatorParams returns a default ValidatorParams, which allows
// only ed25519 pubkeys.
func DefaultValidatorParams() ValidatorParams {
	_ = "STUB: not implemented"
	return *new(ValidatorParams)
}

func DefaultVersionParams() VersionParams { _ = "STUB: not implemented"; return *new(VersionParams) }

func DefaultABCIParams() ABCIParams {
	_ = "STUB: not implemented"

	// When set to 0, vote extensions are not required.
	return *new(ABCIParams)
}

func IsValidPubkeyType(params ValidatorParams, pubkeyType string) bool {
	_ = "STUB: not implemented"
	return false
}

// Validate validates the ConsensusParams to ensure all values are within their
// allowed limits, and returns an error if they are not.
func (params ConsensusParams) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// Check if keyType is a known ABCIPubKeyType

// ValidateUpdate validates the updated VoteExtensionsEnableHeight.
// | r | params...EnableHeight | updated...EnableHeight | result (nil == pass)
// |  1 | *                    | (nil)                  | nil
// |  2 | *                    | < 0                    | VoteExtensionsEnableHeight must be positive
// |  3 | <=0                  | 0                      | nil
// |  4 | X                    | X (>=0)                | nil
// |  5 | > 0; <=height        | 0                      | vote extensions cannot be disabled once enabled
// |  6 | > 0; > height        | 0                      | nil (disable a previous proposal)
// |  7 | *                    | <=height               | vote extensions cannot be updated to a past height
// |  8 | <=0                  | > height (*)           | nil
// |  9 | (> 0) <=height       | > height (*)           | vote extensions cannot be modified once enabled
// | 10 | (> 0) > height       | > height (*)           | nil
func (params ConsensusParams) ValidateUpdate(updated *cmtproto.ConsensusParams, h int64) error {
	_ = "STUB: not implemented"
	// 1
	return nil
}

// 2

// 3

// 4 (implicit: updated.Abci.VoteExtensionsEnableHeight >= 0)

// 5 & 6

// 5

// 6

// 7 (implicit: updated.Abci.VoteExtensionsEnableHeight > 0)

// 8 (implicit: updated.Abci.VoteExtensionsEnableHeight > h)

// 9 (implicit: params.ABCI.VoteExtensionsEnableHeight > 0 && updated.Abci.VoteExtensionsEnableHeight > h)

// 10 (implicit: params.ABCI.VoteExtensionsEnableHeight > h && updated.Abci.VoteExtensionsEnableHeight > h)

// Hash returns a hash of a subset of the parameters to store in the block header.
// Only the Block.MaxBytes and Block.MaxGas are included in the hash.
// This allows the ConsensusParams to evolve more without breaking the block
// protocol. No need for a Merkle tree here, just a small struct to hash.
func (params ConsensusParams) Hash() []byte { _ = "STUB: not implemented"; return nil }

// Update returns a copy of the params with updates from the non-zero fields of p2.
// NOTE: note: must not modify the original
func (params ConsensusParams) Update(params2 *cmtproto.ConsensusParams) ConsensusParams {
	_ = "STUB: not implemented"
	// explicit copy
	return *new(ConsensusParams)
}

// we must defensively consider any structs may be nil

// Copy params2.Validator.PubkeyTypes, and set result's value to the copy.
// This avoids having to initialize the slice to 0 values, and then write to it again.

func (params *ConsensusParams) ToProto() cmtproto.ConsensusParams {
	_ = "STUB: not implemented"
	return *new(cmtproto.ConsensusParams)
}

func ConsensusParamsFromProto(pbParams cmtproto.ConsensusParams) ConsensusParams {
	_ = "STUB: not implemented"
	return *new(ConsensusParams)
}
