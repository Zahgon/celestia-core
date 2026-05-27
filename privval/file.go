package privval

import (
	"time"

	"github.com/cometbft/cometbft/crypto"
	cmtbytes "github.com/cometbft/cometbft/libs/bytes"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cometbft/cometbft/types"
)

// TODO: type ?
const (
	stepNone      int8 = 0 // Used to distinguish the initial state
	stepPropose   int8 = 1
	stepPrevote   int8 = 2
	stepPrecommit int8 = 3
)

// A vote is either stepPrevote or stepPrecommit.
func voteToStep(vote *cmtproto.Vote) int8 { _ = "STUB: not implemented"; return 0 }

//-------------------------------------------------------------------------------

// FilePVKey stores the immutable part of PrivValidator.
type FilePVKey struct {
	Address types.Address  `json:"address"`
	PubKey  crypto.PubKey  `json:"pub_key"`
	PrivKey crypto.PrivKey `json:"priv_key"`

	filePath string
}

// Save persists the FilePVKey to its filePath.
func (pvKey FilePVKey) Save() { _ = "STUB: not implemented"; return }

//-------------------------------------------------------------------------------

// FilePVLastSignState stores the mutable part of PrivValidator.
type FilePVLastSignState struct {
	Height    int64             `json:"height"`
	Round     int32             `json:"round"`
	Step      int8              `json:"step"`
	Signature []byte            `json:"signature,omitempty"`
	SignBytes cmtbytes.HexBytes `json:"signbytes,omitempty"`

	filePath string
}

func (lss *FilePVLastSignState) reset() { _ = "STUB: not implemented"; return }

// CheckHRS checks the given height, round, step (HRS) against that of the
// FilePVLastSignState. It returns an error if the arguments constitute a regression,
// or if they match but the SignBytes are empty.
// The returned boolean indicates whether the last Signature should be reused -
// it returns true if the HRS matches the arguments and the SignBytes are not empty (indicating
// we have already signed for this HRS, and can reuse the existing signature).
// It panics if the HRS matches the arguments, there's a SignBytes, but no Signature.
func (lss *FilePVLastSignState) CheckHRS(height int64, round int32, step int8) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Save persists the FilePvLastSignState to its filePath.
func (lss *FilePVLastSignState) Save() { _ = "STUB: not implemented"; return }

//-------------------------------------------------------------------------------

// FilePV implements PrivValidator using data persisted to disk
// to prevent double signing.
// NOTE: the directories containing pv.Key.filePath and pv.LastSignState.filePath must already exist.
// It includes the LastSignature and LastSignBytes so we don't lose the signature
// if the process crashes after signing but before the resulting consensus message is processed.
type FilePV struct {
	Key           FilePVKey
	LastSignState FilePVLastSignState
}

// NewFilePV generates a new validator from the given key and paths.
func NewFilePV(privKey crypto.PrivKey, keyFilePath, stateFilePath string) *FilePV {
	_ = "STUB: not implemented"
	return nil
}

// GenFilePV generates a new validator with randomly generated private key
// and sets the filePaths, but does not call Save().
func GenFilePV(keyFilePath, stateFilePath string) *FilePV { _ = "STUB: not implemented"; return nil }

// LoadFilePV loads a FilePV from the filePaths.  The FilePV handles double
// signing prevention by persisting data to the stateFilePath.  If either file path
// does not exist, the program will exit.
func LoadFilePV(keyFilePath, stateFilePath string) *FilePV { _ = "STUB: not implemented"; return nil }

// LoadFilePVEmptyState loads a FilePV from the given keyFilePath, with an empty LastSignState.
// If the keyFilePath does not exist, the program will exit.
func LoadFilePVEmptyState(keyFilePath, stateFilePath string) *FilePV {
	_ = "STUB: not implemented"
	return nil
}

// If loadState is true, we load from the stateFilePath. Otherwise, we use an empty LastSignState.
func loadFilePV(keyFilePath, stateFilePath string, loadState bool) *FilePV {
	_ = "STUB: not implemented"
	return nil
}

// overwrite pubkey and address for convenience

// LoadOrGenFilePV loads a FilePV from the given filePaths
// or else generates a new one and saves it to the filePaths.
func LoadOrGenFilePV(keyFilePath, stateFilePath string) *FilePV {
	_ = "STUB: not implemented"
	return nil
}

// GetAddress returns the address of the validator.
// Implements PrivValidator.
func (pv *FilePV) GetAddress() types.Address {
	_ = "STUB: not implemented"
	return *

	// GetPubKey returns the public key of the validator.
	// Implements PrivValidator.
	new(types.Address)
}

func (pv *FilePV) GetPubKey() (crypto.PubKey, error) {
	_ = "STUB: not implemented"
	return *new(crypto.PubKey), nil
}

// SignVote signs a canonical representation of the vote, along with the
// chainID. Implements PrivValidator.
func (pv *FilePV) SignVote(chainID string, vote *cmtproto.Vote) error {
	_ = "STUB: not implemented"
	return nil
}

// SignProposal signs a canonical representation of the proposal, along with
// the chainID. Implements PrivValidator.
func (pv *FilePV) SignProposal(chainID string, proposal *cmtproto.Proposal) error {
	_ = "STUB: not implemented"
	return nil
}

func (pv *FilePV) SignRawBytes(chainID, uniqueID string, rawBytes []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Save persists the FilePV to disk.
func (pv *FilePV) Save() { _ = "STUB: not implemented"; return }

// Reset resets all fields in the FilePV.
// NOTE: Unsafe!
func (pv *FilePV) Reset() { _ = "STUB: not implemented"; return }

// String returns a string representation of the FilePV.
func (pv *FilePV) String() string { _ = "STUB: not implemented"; return "" }

//------------------------------------------------------------------------------------

// signVote checks if the vote is good to sign and sets the vote signature.
// It may need to set the timestamp as well if the vote is otherwise the same as
// a previously signed vote (ie. we crashed after signing but before the vote hit the WAL).
// Extension signatures are always signed for non-nil precommits (even if the data is empty).
func (pv *FilePV) signVote(chainID string, vote *cmtproto.Vote) error {
	_ = "STUB: not implemented"
	return nil
}

// Vote extensions are non-deterministic, so it is possible that an
// application may have created a different extension. We therefore always
// re-sign the vote extensions of precommits. For prevotes and nil
// precommits, the extension signature will always be empty.
// Even if the signed over data is empty, we still add the signature

// We might crash before writing to the wal,
// causing us to try to re-sign for the same HRS.
// If signbytes are the same, use the last signature.
// If they only differ by timestamp, use last timestamp and signature
// Otherwise, return error

// Compares the canonicalized votes (i.e. without vote extensions
// or vote extension signatures).

// It passed the checks. Sign the vote

// signProposal checks if the proposal is good to sign and sets the proposal signature.
// It may need to set the timestamp as well if the proposal is otherwise the same as
// a previously signed proposal ie. we crashed after signing but before the proposal hit the WAL).
func (pv *FilePV) signProposal(chainID string, proposal *cmtproto.Proposal) error {
	_ = "STUB: not implemented"
	return nil
}

// We might crash before writing to the wal,
// causing us to try to re-sign for the same HRS.
// If signbytes are the same, use the last signature.
// If they only differ by timestamp, use last timestamp and signature
// Otherwise, return error

// It passed the checks. Sign the proposal

// Persist height/round/step and signature
func (pv *FilePV) saveSigned(height int64, round int32, step int8,
	signBytes []byte, sig []byte) {
	_ = "STUB: not implemented"
	return
}

//-----------------------------------------------------------------------------------------

// Returns the timestamp from the lastSignBytes.
// Returns true if the only difference in the votes is their timestamp.
// Performs these checks on the canonical votes (excluding the vote extension
// and vote extension signatures).
func checkVotesOnlyDifferByTimestamp(lastSignBytes, newSignBytes []byte) (time.Time, bool) {
	_ = "STUB: not implemented"
	return *new(time.Time), false
}

// set the times to the same value and check equality

// returns the timestamp from the lastSignBytes.
// returns true if the only difference in the proposals is their timestamp
func checkProposalsOnlyDifferByTimestamp(lastSignBytes, newSignBytes []byte) (time.Time, bool) {
	_ = "STUB: not implemented"
	return *new(time.Time), false
}

// set the times to the same value and check equality
