package types

import (
	"time"

	"github.com/cometbft/cometbft/libs/bits"
	cmtbytes "github.com/cometbft/cometbft/libs/bytes"
	cmtsync "github.com/cometbft/cometbft/libs/sync"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	cmtversion "github.com/cometbft/cometbft/proto/tendermint/version"
)

const (
	// MaxHeaderBytes is a maximum header size.
	// NOTE: Because app hash can be of arbitrary size, the header is therefore not
	// capped in size and thus this number should be seen as a soft max
	MaxHeaderBytes int64 = 623

	// MaxOverheadForBlock - maximum overhead to encode a block (up to
	// MaxBlockSizeBytes in size) not including it's parts except Data.
	// This means it also excludes the overhead for individual transactions.
	//
	// Uvarint length of MaxBlockSizeBytes: 4 bytes
	// 2 fields (2 embedded):               2 bytes
	// Uvarint length of Data.Txs:          4 bytes
	// Data.Txs field:                      1 byte
	MaxOverheadForBlock int64 = 11
)

// Block defines the atomic unit of a CometBFT blockchain.
type Block struct {
	mtx cmtsync.Mutex

	verifiedHash cmtbytes.HexBytes // Verified block hash (not included in the struct hash)
	Header       `json:"header"`
	Data         `json:"data"`
	Evidence     EvidenceData `json:"evidence"`
	LastCommit   *Commit      `json:"last_commit"`

	// cachedHashes is used purely for passing the hashes of the tx alongside
	// the block. This are not included in any encoding of this struct.
	cachedHashes [][]byte
}

// ValidateBasic performs basic validation that doesn't involve state data.
// It checks the internal consistency of the block.
// Further validation is done using state#ValidateBlock.
func (b *Block) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// Validate the last commit and its hash.

// NOTE: b.Data.Txs may be nil, but b.Data.Hash() still works fine.

// NOTE: b.Evidence.Evidence may be nil, but we're just looping.

// fillHeader fills in any remaining header fields that are a function of the block data
func (b *Block) fillHeader() { _ = "STUB: not implemented"; return }

// Hash computes and returns the block hash.
// If the block is incomplete, block hash is nil for safety.
func (b *Block) Hash() cmtbytes.HexBytes { _ = "STUB: not implemented"; return *new(cmtbytes.HexBytes) }

// MakePartSet returns a PartSet containing parts of a serialized block.
// This is the form in which the block is gossipped to peers.
// CONTRACT: partSize is greater than zero.
func (b *Block) MakePartSet(partSize uint32) (*PartSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// HashesTo is a convenience function that checks if a block hashes to the given argument.
// Returns false if the block is nil or the hash is empty.
func (b *Block) HashesTo(hash []byte) bool { _ = "STUB: not implemented"; return false }

// Size returns size of the block in bytes.
func (b *Block) Size() int { _ = "STUB: not implemented"; return 0 }

// String returns a string representation of the block
//
// See StringIndented.
func (b *Block) String() string { _ = "STUB: not implemented"; return "" }

// StringIndented returns an indented String.
//
// Header
// Data
// Evidence
// LastCommit
// Hash
func (b *Block) StringIndented(indent string) string { _ = "STUB: not implemented"; return "" }

// StringShort returns a shortened string representation of the block.
func (b *Block) StringShort() string { _ = "STUB: not implemented"; return "" }

// ToProto converts Block to protobuf
func (b *Block) ToProto() (*cmtproto.Block, error) { _ = "STUB: not implemented"; return nil, nil }

// CachedHashes return any hashes of transactions that were included in this
// block. This is used for passing the hashes of the txs alongside the block,
// they are not included in any validity rule or encoding of the block.
func (b *Block) CachedHashes() [][]byte { _ = "STUB: not implemented"; return nil }

// SetCachedHashes sets the cached hashes of the block. This is used for passing
// the hashes of the txs alongside the block, they are not included in any
// validity rule or encoding of the block.
func (b *Block) SetCachedHashes(hashes [][]byte) { _ = "STUB: not implemented"; return }

// FromProto sets a protobuf Block to the given pointer.
// It returns an error if the block is invalid.
func BlockFromProto(bp *cmtproto.Block) (*Block, error) { _ = "STUB: not implemented"; return nil, nil }

//-----------------------------------------------------------------------------

// MaxDataBytes returns the maximum size of block's data.
//
// XXX: Panics on negative result.
func MaxDataBytes(maxBytes, evidenceBytes int64, valsCount int) int64 {
	_ = "STUB: not implemented"
	return 0
}

// MaxDataBytesNoEvidence returns the maximum size of block's data when
// evidence count is unknown (will be assumed to be 0).
//
// XXX: Panics on negative result.
func MaxDataBytesNoEvidence(maxBytes int64, valsCount int) int64 {
	_ = "STUB: not implemented"
	return 0
}

//-----------------------------------------------------------------------------

// Header defines the structure of a CometBFT block header.
// NOTE: changes to the Header should be duplicated in:
// - header.Hash()
// - abci.Header
// - https://github.com/cometbft/cometbft/blob/v0.38.x/spec/blockchain/blockchain.md
type Header struct {
	// basic block info
	Version cmtversion.Consensus `json:"version"`
	ChainID string               `json:"chain_id"`
	Height  int64                `json:"height"`
	Time    time.Time            `json:"time"`

	// prev block info
	LastBlockID BlockID `json:"last_block_id"`

	// hashes of block data
	LastCommitHash cmtbytes.HexBytes `json:"last_commit_hash"` // commit from validators from the last block
	DataHash       cmtbytes.HexBytes `json:"data_hash"`        // transactions

	// hashes from the app output from the prev block
	ValidatorsHash     cmtbytes.HexBytes `json:"validators_hash"`      // validators for the current block
	NextValidatorsHash cmtbytes.HexBytes `json:"next_validators_hash"` // validators for the next block
	ConsensusHash      cmtbytes.HexBytes `json:"consensus_hash"`       // consensus params for current block
	AppHash            cmtbytes.HexBytes `json:"app_hash"`             // state after txs from the previous block
	// root hash of all results from the txs from the previous block
	// see `deterministicExecTxResult` to understand which parts of a tx is hashed into here
	LastResultsHash cmtbytes.HexBytes `json:"last_results_hash"`

	// consensus info
	EvidenceHash    cmtbytes.HexBytes `json:"evidence_hash"`    // evidence included in the block
	ProposerAddress Address           `json:"proposer_address"` // original proposer of the block
}

// Populate the Header with state-derived data.
// Call this after MakeBlock to complete the Header.
func (h *Header) Populate(
	version cmtversion.Consensus, chainID string,
	timestamp time.Time, lastBlockID BlockID,
	valHash, nextValHash []byte,
	consensusHash, appHash, lastResultsHash []byte,
	proposerAddress Address,
) {
	_ = "STUB: not implemented"
	return
}

// ValidateBasic performs stateless validation on a Header returning an error
// if any validation fails.
//
// NOTE: Timestamp validation is subtle and handled elsewhere.
func (h Header) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// Basic validation of hashes related to application data.
// Will validate fully against state in state#ValidateBlock.

// NOTE: AppHash is arbitrary length

// Hash returns the hash of the header.
// It computes a Merkle tree from the header fields
// ordered as they appear in the Header.
// Returns nil if ValidatorHash is missing,
// since a Header is not valid unless there is
// a ValidatorsHash (corresponding to the validator set).
func (h *Header) Hash() cmtbytes.HexBytes {
	_ = "STUB: not implemented"
	return *new(cmtbytes.HexBytes)
}

// StringIndented returns an indented string representation of the header.
func (h *Header) StringIndented(indent string) string { _ = "STUB: not implemented"; return "" }

// ToProto converts Header to protobuf
func (h *Header) ToProto() *cmtproto.Header { _ = "STUB: not implemented"; return nil }

// FromProto sets a protobuf Header to the given pointer.
// It returns an error if the header is invalid.
func HeaderFromProto(ph *cmtproto.Header) (Header, error) {
	_ = "STUB: not implemented"
	return *new(Header), nil
}

//-------------------------------------

// BlockIDFlag indicates which BlockID the signature is for.
type BlockIDFlag byte

const (
	// BlockIDFlagAbsent - no vote was received from a validator.
	BlockIDFlagAbsent BlockIDFlag = iota + 1
	// BlockIDFlagCommit - voted for the Commit.BlockID.
	BlockIDFlagCommit
	// BlockIDFlagNil - voted for nil.
	BlockIDFlagNil
)

const (
	// Max size of commit without any commitSigs -> 79 for BlockID, 8 for Height, 4 for Round.
	MaxCommitOverheadBytes int64 = 91
	// Commit sig size is made up of 64 bytes for the signature, 20 bytes for the address,
	// 1 byte for the flag and 14 bytes for the timestamp
	MaxCommitSigBytes int64 = 109
)

// CommitSig is a part of the Vote included in a Commit.
type CommitSig struct {
	BlockIDFlag      BlockIDFlag `json:"block_id_flag"`
	ValidatorAddress Address     `json:"validator_address"`
	Timestamp        time.Time   `json:"timestamp"`
	Signature        []byte      `json:"signature"`
}

func MaxCommitBytes(valCount int) int64 {
	_ = "STUB: not implemented"
	// From the repeated commit sig field
	return 0
}

// NewCommitSigAbsent returns new CommitSig with BlockIDFlagAbsent. Other
// fields are all empty.
func NewCommitSigAbsent() CommitSig { _ = "STUB: not implemented"; return *new(CommitSig) }

// CommitSig returns a string representation of CommitSig.
//
// 1. first 6 bytes of signature
// 2. first 6 bytes of validator address
// 3. block ID flag
// 4. timestamp
func (cs CommitSig) String() string { _ = "STUB: not implemented"; return "" }

// BlockID returns the Commit's BlockID if CommitSig indicates signing,
// otherwise - empty BlockID.
func (cs CommitSig) BlockID(commitBlockID BlockID) BlockID {
	_ = "STUB: not implemented"
	return *new(BlockID)
}

// ValidateBasic performs basic validation.
func (cs CommitSig) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// NOTE: Timestamp validation is subtle and handled elsewhere.

// ToProto converts CommitSig to protobuf
func (cs *CommitSig) ToProto() *cmtproto.CommitSig { _ = "STUB: not implemented"; return nil }

// FromProto sets a protobuf CommitSig to the given pointer.
// It returns an error if the CommitSig is invalid.
func (cs *CommitSig) FromProto(csp cmtproto.CommitSig) error { _ = "STUB: not implemented"; return nil }

//-------------------------------------

// ExtendedCommitSig contains a commit signature along with its corresponding
// vote extension and vote extension signature.
type ExtendedCommitSig struct {
	CommitSig                 // Commit signature
	Extension          []byte // Vote extension
	ExtensionSignature []byte // Vote extension signature
}

// NewExtendedCommitSigAbsent returns new ExtendedCommitSig with
// BlockIDFlagAbsent. Other fields are all empty.
func NewExtendedCommitSigAbsent() ExtendedCommitSig {
	_ = "STUB: not implemented"
	return *new(ExtendedCommitSig)
}

// String returns a string representation of an ExtendedCommitSig.
//
// 1. commit sig
// 2. first 6 bytes of vote extension
// 3. first 6 bytes of vote extension signature
func (ecs ExtendedCommitSig) String() string { _ = "STUB: not implemented"; return "" }

// ValidateBasic checks whether the structure is well-formed.
func (ecs ExtendedCommitSig) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// EnsureExtensions validates that a vote extensions signature is present for
// this ExtendedCommitSig.
func (ecs ExtendedCommitSig) EnsureExtension(extEnabled bool) error {
	_ = "STUB: not implemented"
	return nil
}

// ToProto converts the ExtendedCommitSig to its Protobuf representation.
func (ecs *ExtendedCommitSig) ToProto() *cmtproto.ExtendedCommitSig {
	_ = "STUB: not implemented"
	return nil
}

// FromProto populates the ExtendedCommitSig with values from the given
// Protobuf representation. Returns an error if the ExtendedCommitSig is
// invalid.
func (ecs *ExtendedCommitSig) FromProto(ecsp cmtproto.ExtendedCommitSig) error {
	_ = "STUB: not implemented"
	return nil
}

//-------------------------------------

// Commit contains the evidence that a block was committed by a set of validators.
// NOTE: Commit is empty for height 1, but never nil.
type Commit struct {
	// NOTE: The signatures are in order of address to preserve the bonded
	// ValidatorSet order.
	// Any peer with a block can gossip signatures by index with a peer without
	// recalculating the active ValidatorSet.
	Height     int64       `json:"height"`
	Round      int32       `json:"round"`
	BlockID    BlockID     `json:"block_id"`
	Signatures []CommitSig `json:"signatures"`

	// Memoized in first call to corresponding method.
	// NOTE: can't memoize in constructor because constructor isn't used for
	// unmarshaling.
	hash cmtbytes.HexBytes
}

// Clone creates a deep copy of this commit.
func (commit *Commit) Clone() *Commit { _ = "STUB: not implemented"; return nil }

// GetVote converts the CommitSig for the given valIdx to a Vote. Commits do
// not contain vote extensions, so the vote extension and vote extension
// signature will not be present in the returned vote.
// Returns nil if the precommit at valIdx is nil.
// Panics if valIdx >= commit.Size().
func (commit *Commit) GetVote(valIdx int32) *Vote { _ = "STUB: not implemented"; return nil }

// VoteSignBytes returns the bytes of the Vote corresponding to valIdx for
// signing.
//
// The only unique part is the Timestamp - all other fields signed over are
// otherwise the same for all validators.
//
// Panics if valIdx >= commit.Size().
//
// See VoteSignBytes
func (commit *Commit) VoteSignBytes(chainID string, valIdx int32) []byte {
	_ = "STUB: not implemented"
	return nil
}

// Size returns the number of signatures in the commit.
func (commit *Commit) Size() int { _ = "STUB: not implemented"; return 0 }

// ValidateBasic performs basic validation that doesn't involve state data.
// Does not actually check the cryptographic signatures.
func (commit *Commit) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// Hash returns the hash of the commit
func (commit *Commit) Hash() cmtbytes.HexBytes {
	_ = "STUB: not implemented"
	return *new(cmtbytes.HexBytes)
}

// WrappedExtendedCommit wraps a commit as an ExtendedCommit.
// The VoteExtension fields of the resulting value will by nil.
// Wrapping a Commit as an ExtendedCommit is useful when an API
// requires an ExtendedCommit wire type but does not
// need the VoteExtension data.
func (commit *Commit) WrappedExtendedCommit() *ExtendedCommit {
	_ = "STUB: not implemented"
	return nil
}

// StringIndented returns a string representation of the commit.
func (commit *Commit) StringIndented(indent string) string { _ = "STUB: not implemented"; return "" }

// ToProto converts Commit to protobuf
func (commit *Commit) ToProto() *cmtproto.Commit { _ = "STUB: not implemented"; return nil }

// FromProto sets a protobuf Commit to the given pointer.
// It returns an error if the commit is invalid.
func CommitFromProto(cp *cmtproto.Commit) (*Commit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//-------------------------------------

// ExtendedCommit is similar to Commit, except that its signatures also retain
// their corresponding vote extensions and vote extension signatures.
type ExtendedCommit struct {
	Height             int64
	Round              int32
	BlockID            BlockID
	ExtendedSignatures []ExtendedCommitSig

	bitArray *bits.BitArray
}

// Clone creates a deep copy of this extended commit.
func (ec *ExtendedCommit) Clone() *ExtendedCommit { _ = "STUB: not implemented"; return nil }

// ToExtendedVoteSet constructs a VoteSet from the Commit and validator set.
// Panics if signatures from the ExtendedCommit can't be added to the voteset.
// Panics if any of the votes have invalid or absent vote extension data.
// Inverse of VoteSet.MakeExtendedCommit().
func (ec *ExtendedCommit) ToExtendedVoteSet(chainID string, vals *ValidatorSet) *VoteSet {
	_ = "STUB: not implemented"
	return nil
}

// addSigsToVoteSet adds all of the signature to voteSet.
func (ec *ExtendedCommit) addSigsToVoteSet(voteSet *VoteSet) { _ = "STUB: not implemented"; return }

// OK, some precommits can be missing.

// ToVoteSet constructs a VoteSet from the Commit and validator set.
// Panics if signatures from the commit can't be added to the voteset.
// Inverse of VoteSet.MakeCommit().
func (commit *Commit) ToVoteSet(chainID string, vals *ValidatorSet) *VoteSet {
	_ = "STUB: not implemented"
	return nil
}

// OK, some precommits can be missing.

// EnsureExtensions validates that a vote extensions signature is present for
// every ExtendedCommitSig in the ExtendedCommit.
func (ec *ExtendedCommit) EnsureExtensions(extEnabled bool) error {
	_ = "STUB: not implemented"
	return nil
}

// ToCommit converts an ExtendedCommit to a Commit by removing all vote
// extension-related fields.
func (ec *ExtendedCommit) ToCommit() *Commit { _ = "STUB: not implemented"; return nil }

// GetExtendedVote converts the ExtendedCommitSig for the given validator
// index to a Vote with a vote extensions.
// It panics if valIndex is out of range.
func (ec *ExtendedCommit) GetExtendedVote(valIndex int32) *Vote {
	_ = "STUB: not implemented"
	return nil
}

// Type returns the vote type of the extended commit, which is always
// VoteTypePrecommit
// Implements VoteSetReader.
func (ec *ExtendedCommit) Type() byte { _ = "STUB: not implemented"; return 0 }

// GetHeight returns height of the extended commit.
// Implements VoteSetReader.
func (ec *ExtendedCommit) GetHeight() int64 {
	_ = "STUB: not implemented"

	// GetRound returns height of the extended commit.
	// Implements VoteSetReader.
	return 0
}

func (ec *ExtendedCommit) GetRound() int32 {
	_ = "STUB: not implemented"

	// Size returns the number of signatures in the extended commit.
	// Implements VoteSetReader.
	return 0
}

func (ec *ExtendedCommit) Size() int { _ = "STUB: not implemented"; return 0 }

// BitArray returns a BitArray of which validators voted for BlockID or nil in
// this extended commit.
// Implements VoteSetReader.
func (ec *ExtendedCommit) BitArray() *bits.BitArray { _ = "STUB: not implemented"; return nil }

// TODO: need to check the BlockID otherwise we could be counting conflicts,
//       not just the one with +2/3 !

// GetByIndex returns the vote corresponding to a given validator index.
// Panics if `index >= extCommit.Size()`.
// Implements VoteSetReader.
func (ec *ExtendedCommit) GetByIndex(valIdx int32) *Vote { _ = "STUB: not implemented"; return nil }

// IsCommit returns true if there is at least one signature.
// Implements VoteSetReader.
func (ec *ExtendedCommit) IsCommit() bool { _ = "STUB: not implemented"; return false }

// ValidateBasic checks whether the extended commit is well-formed. Does not
// actually check the cryptographic signatures.
func (ec *ExtendedCommit) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// ToProto converts ExtendedCommit to protobuf
func (ec *ExtendedCommit) ToProto() *cmtproto.ExtendedCommit { _ = "STUB: not implemented"; return nil }

// ExtendedCommitFromProto constructs an ExtendedCommit from the given Protobuf
// representation. It returns an error if the extended commit is invalid.
func ExtendedCommitFromProto(ecp *cmtproto.ExtendedCommit) (*ExtendedCommit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//-------------------------------------

// Data contains the set of transactions included in the block
type Data struct {
	// Txs that will be applied by state @ block.Height+1.
	// NOTE: not all txs here are valid.  We're just agreeing on the order first.
	// This means that block.AppHash does not include these txs.
	Txs Txs `json:"txs"`

	// SquareSize is the size of the square after splitting all the block data
	// into shares. The erasure data is discarded after generation, and keeping this
	// value avoids unnecessarily regenerating all of the shares when returning
	// proofs that some element was included in the block
	SquareSize uint64 `json:"square_size"`

	// Volatile
	hash cmtbytes.HexBytes
}

func NewData(txs Txs, squareSize uint64, hash cmtbytes.HexBytes) Data {
	_ = "STUB: not implemented"
	return *new(Data)
}

// Hash returns the hash of the data
func (data *Data) Hash() cmtbytes.HexBytes {
	_ = "STUB: not implemented"
	return *new(cmtbytes.HexBytes)
}

// NOTE: leaves of merkle tree are TxIDs

// StringIndented returns an indented string representation of the transactions.
func (data *Data) StringIndented(indent string) string { _ = "STUB: not implemented"; return "" }

// ToProto converts Data to protobuf
func (data *Data) ToProto() cmtproto.Data { _ = "STUB: not implemented"; return *new(cmtproto.Data) }

// DataFromProto takes a protobuf representation of Data &
// returns the native type.
func DataFromProto(dp *cmtproto.Data) (Data, error) {
	_ = "STUB: not implemented"
	return *new(Data), nil
}

//-----------------------------------------------------------------------------

type Blob struct {
	// NamespaceVersion is the version of the namespace. Used in conjunction
	// with NamespaceID to determine the namespace of this blob.
	NamespaceVersion uint8

	// NamespaceID defines the namespace ID of this blob. Used in conjunction
	// with NamespaceVersion to determine the namespace of this blob.
	NamespaceID []byte

	// Data is the actual data of the blob.
	// (e.g. a block of a virtual sidechain).
	Data []byte

	// ShareVersion is the version of the share format that this blob should use
	// when encoded into shares.
	ShareVersion uint8
}

// Namespace returns the namespace of this blob encoded as a byte slice.
func (b Blob) Namespace() []byte { _ = "STUB: not implemented"; return nil }

// -----------------------------------------------------------------------------

// EvidenceData contains any evidence of malicious wrong-doing by validators
type EvidenceData struct {
	Evidence EvidenceList `json:"evidence"`

	// Volatile. Used as cache
	hash     cmtbytes.HexBytes
	byteSize int64
}

// Hash returns the hash of the data.
func (data *EvidenceData) Hash() cmtbytes.HexBytes {
	_ = "STUB: not implemented"
	return *new(cmtbytes.HexBytes)
}

// ByteSize returns the total byte size of all the evidence
func (data *EvidenceData) ByteSize() int64 { _ = "STUB: not implemented"; return 0 }

// StringIndented returns a string representation of the evidence.
func (data *EvidenceData) StringIndented(indent string) string {
	_ = "STUB: not implemented"
	return ""
}

// ToProto converts EvidenceData to protobuf
func (data *EvidenceData) ToProto() (*cmtproto.EvidenceList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FromProto sets a protobuf EvidenceData to the given pointer.
func (data *EvidenceData) FromProto(eviData *cmtproto.EvidenceList) error {
	_ = "STUB: not implemented"
	return nil
}

//--------------------------------------------------------------------------------

// BlockID
type BlockID struct {
	Hash          cmtbytes.HexBytes `json:"hash"`
	PartSetHeader PartSetHeader     `json:"parts"`
}

// Equals returns true if the BlockID matches the given BlockID
func (blockID BlockID) Equals(other BlockID) bool { _ = "STUB: not implemented"; return false }

// Key returns a machine-readable string representation of the BlockID
func (blockID BlockID) Key() string { _ = "STUB: not implemented"; return "" }

// ValidateBasic performs basic validation.
func (blockID BlockID) ValidateBasic() error {
	_ = "STUB: not implemented"
	// Hash can be empty in case of POLBlockID in Proposal.
	return nil
}

// IsZero returns true if this is the BlockID of a nil block.
func (blockID BlockID) IsZero() bool { _ = "STUB: not implemented"; return false }

// IsComplete returns true if this is a valid BlockID of a non-nil block.
func (blockID BlockID) IsComplete() bool { _ = "STUB: not implemented"; return false }

// String returns a human readable string representation of the BlockID.
//
// 1. hash
// 2. part set header
//
// See PartSetHeader#String
func (blockID BlockID) String() string { _ = "STUB: not implemented"; return "" }

// ToProto converts BlockID to protobuf
func (blockID *BlockID) ToProto() cmtproto.BlockID {
	_ = "STUB: not implemented"
	return *new(cmtproto.BlockID)
}

// FromProto sets a protobuf BlockID to the given pointer.
// It returns an error if the block id is invalid.
func BlockIDFromProto(bID *cmtproto.BlockID) (*BlockID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ProtoBlockIDIsNil is similar to the IsNil function on BlockID, but for the
// Protobuf representation.
func ProtoBlockIDIsNil(bID *cmtproto.BlockID) bool { _ = "STUB: not implemented"; return false }
