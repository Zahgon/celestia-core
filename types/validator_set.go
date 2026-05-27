package types

import (
	"errors"
	"fmt"
	"math"

	cmtmath "github.com/cometbft/cometbft/libs/math"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
)

const (
	// MaxTotalVotingPower - the maximum allowed total voting power.
	// It needs to be sufficiently small to, in all cases:
	// 1. prevent clipping in incrementProposerPriority()
	// 2. let (diff+diffMax-1) not overflow in IncrementProposerPriority()
	// (Proof of 1 is tricky, left to the reader).
	// It could be higher, but this is sufficiently large for our purposes,
	// and leaves room for defensive purposes.
	MaxTotalVotingPower = int64(math.MaxInt64) / 8

	// PriorityWindowSizeFactor - is a constant that when multiplied with the
	// total voting power gives the maximum allowed distance between validator
	// priorities.
	PriorityWindowSizeFactor = 2
)

// ErrTotalVotingPowerOverflow is returned if the total voting power of the
// resulting validator set exceeds MaxTotalVotingPower.
var ErrTotalVotingPowerOverflow = fmt.Errorf("total voting power of resulting valset exceeds max %d",
	MaxTotalVotingPower)

// ErrProposerNotInVals is returned if the proposer is not in the validator set.
var ErrProposerNotInVals = errors.New("proposer not in validator set")

// ValidatorSet represent a set of *Validator at a given height.
//
// The validators can be fetched by address or index.
// The index is in order of .VotingPower, so the indices are fixed for all
// rounds of a given blockchain height - ie. the validators are sorted by their
// voting power (descending). Secondary index - .Address (ascending).
//
// On the other hand, the .ProposerPriority of each validator and the
// designated .GetProposer() of a set changes every round, upon calling
// .IncrementProposerPriority().
//
// NOTE: Not goroutine-safe.
// NOTE: All get/set to validators should copy the value for safety.
type ValidatorSet struct {
	// NOTE: persisted via reflect, must be exported.
	Validators []*Validator `json:"validators"`
	Proposer   *Validator   `json:"proposer"`

	// cached (unexported)
	totalVotingPower int64
	// true if all validators have the same type of public key or if the set is empty.
	allKeysHaveSameType bool
}

// NewValidatorSet initializes a ValidatorSet by copying over the values from
// `valz`, a list of Validators. If valz is nil or empty, the new ValidatorSet
// will have an empty list of Validators.
//
// The addresses of validators in `valz` must be unique otherwise the function
// panics.
//
// Note the validator set size has an implied limit equal to that of the
// MaxVotesCount - commits by a validator set larger than this will fail
// validation.
func NewValidatorSet(valz []*Validator) *ValidatorSet { _ = "STUB: not implemented"; return nil }

func (vals *ValidatorSet) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// IsNilOrEmpty returns true if validator set is nil or empty.
func (vals *ValidatorSet) IsNilOrEmpty() bool { _ = "STUB: not implemented"; return false }

// CopyIncrementProposerPriority increments ProposerPriority and updates the
// proposer on a copy, and returns it.
func (vals *ValidatorSet) CopyIncrementProposerPriority(times int32) *ValidatorSet {
	_ = "STUB: not implemented"
	return nil
}

// IncrementProposerPriority increments ProposerPriority of each validator and
// updates the proposer. Panics if validator set is empty.
// `times` must be positive.
func (vals *ValidatorSet) IncrementProposerPriority(times int32) { _ = "STUB: not implemented"; return }

// Cap the difference between priorities to be proportional to 2*totalPower by
// re-normalizing priorities, i.e., rescale all priorities by multiplying with:
//  2*totalVotingPower/(maxPriority - minPriority)

// Call IncrementProposerPriority(1) times times.

// RescalePriorities rescales the priorities such that the distance between the
// maximum and minimum is smaller than `diffMax`. Panics if validator set is
// empty.
func (vals *ValidatorSet) RescalePriorities(diffMax int64) { _ = "STUB: not implemented"; return }

// NOTE: This check is merely a sanity check which could be
// removed if all tests would init. voting power appropriately;
// i.e. diffMax should always be > 0

// Calculating ceil(diff/diffMax):
// Re-normalization is performed by dividing by an integer for simplicity.
// NOTE: This may make debugging priority issues easier as well.

func (vals *ValidatorSet) incrementProposerPriority() *Validator {
	_ = "STUB: not implemented"
	return nil
}

// Check for overflow for sum.

// Decrement the validator with most ProposerPriority.

// Mind the underflow.

// Should not be called on an empty validator set.
func (vals *ValidatorSet) computeAvgProposerPriority() int64 { _ = "STUB: not implemented"; return 0 }

// This should never happen: each val.ProposerPriority is in bounds of int64.

// Compute the difference between the max and min ProposerPriority of that set.
func computeMaxMinPriorityDiff(vals *ValidatorSet) int64 { _ = "STUB: not implemented"; return 0 }

func (vals *ValidatorSet) getValWithMostPriority() *Validator {
	_ = "STUB: not implemented"
	return nil
}

func (vals *ValidatorSet) shiftByAvgProposerPriority() { _ = "STUB: not implemented"; return }

// Makes a copy of the validator list.
func validatorListCopy(valsList []*Validator) []*Validator { _ = "STUB: not implemented"; return nil }

// Copy each validator into a new ValidatorSet.
func (vals *ValidatorSet) Copy() *ValidatorSet { _ = "STUB: not implemented"; return nil }

// HasAddress returns true if address given is in the validator set, false -
// otherwise.
func (vals *ValidatorSet) HasAddress(address []byte) bool { _ = "STUB: not implemented"; return false }

// GetByAddress returns an index of the validator with address and validator
// itself (copy) if found. Otherwise, -1 and nil are returned.
func (vals *ValidatorSet) GetByAddress(address []byte) (index int32, val *Validator) {
	_ = "STUB: not implemented"
	return 0, nil
}

// GetByIndex returns the validator's address and validator itself (copy) by
// index.
// It returns nil values if index is less than 0 or greater or equal to
// len(ValidatorSet.Validators).
func (vals *ValidatorSet) GetByIndex(index int32) (address []byte, val *Validator) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Size returns the length of the validator set.
func (vals *ValidatorSet) Size() int { _ = "STUB: not implemented"; return 0 }

// Forces recalculation of the set's total voting power.
// Returns an error if total voting power exceeds MaxTotalVotingPower.
func (vals *ValidatorSet) updateTotalVotingPower() error { _ = "STUB: not implemented"; return nil }

// mind overflow

// TotalVotingPowerSafe returns the sum of the voting powers of all validators,
// or an error if the total exceeds MaxTotalVotingPower.
func (vals *ValidatorSet) TotalVotingPowerSafe() (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// TotalVotingPower returns the sum of the voting powers of all validators.
// It recomputes the total voting power if required.
func (vals *ValidatorSet) TotalVotingPower() int64 { _ = "STUB: not implemented"; return 0 }

// GetProposer returns the current proposer. If the validator set is empty, nil
// is returned.
func (vals *ValidatorSet) GetProposer() (proposer *Validator) {
	_ = "STUB: not implemented"
	return nil
}

func (vals *ValidatorSet) findProposer() *Validator { _ = "STUB: not implemented"; return nil }

// Hash returns the Merkle root hash build using validators (as leaves) in the
// set.
//
// See merkle.HashFromByteSlices.
func (vals *ValidatorSet) Hash() []byte { _ = "STUB: not implemented"; return nil }

// ProposerPriorityHash returns the tmhash of the proposer priorities.
// Validator set must be sorted to get the same hash.
// If the validator set is empty, nil is returned.
func (vals *ValidatorSet) ProposerPriorityHash() []byte { _ = "STUB: not implemented"; return nil }

// Iterate will run the given function over the set.
func (vals *ValidatorSet) Iterate(fn func(index int, val *Validator) bool) {
	_ = "STUB: not implemented"
	return
}

// Checks changes against duplicates, splits the changes in updates and
// removals, sorts them by address.
//
// Returns:
// updates, removals - the sorted lists of updates and removals
// err - non-nil if duplicate entries or entries with negative voting power are seen
//
// No changes are made to 'origChanges'.
func processChanges(origChanges []*Validator) (updates, removals []*Validator, err error) {
	_ = "STUB: not implemented"
	// Make a deep copy of the changes and sort by address.
	return nil, nil, nil
}

// Scan changes by address and append valid validators to updates or removals lists.

// verifyUpdates verifies a list of updates against a validator set, making sure the allowed
// total voting power would not be exceeded if these updates would be applied to the set.
//
// Inputs:
// updates - a list of proper validator changes, i.e. they have been verified by processChanges for duplicates
//
//	and invalid values.
//
// vals - the original validator set. Note that vals is NOT modified by this function.
// removedPower - the total voting power that will be removed after the updates are verified and applied.
//
// Returns:
// tvpAfterUpdatesBeforeRemovals -  the new total voting power if these updates would be applied without the removals.
//
//	Note that this will be < 2 * MaxTotalVotingPower in case high power validators are removed and
//	validators are added/ updated with high power values.
//
// err - non-nil if the maximum allowed total voting power would be exceeded
func verifyUpdates(
	updates []*Validator,
	vals *ValidatorSet,
	removedPower int64,
) (tvpAfterUpdatesBeforeRemovals int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func numNewValidators(updates []*Validator, vals *ValidatorSet) int {
	_ = "STUB: not implemented"
	return 0
}

// computeNewPriorities computes the proposer priority for the validators not present in the set based on
// 'updatedTotalVotingPower'.
// Leaves unchanged the priorities of validators that are changed.
//
// 'updates' parameter must be a list of unique validators to be added or updated.
//
// 'updatedTotalVotingPower' is the total voting power of a set where all updates would be applied but
//
//	not the removals. It must be < 2*MaxTotalVotingPower and may be close to this limit if close to
//	MaxTotalVotingPower will be removed. This is still safe from overflow since MaxTotalVotingPower is maxInt64/8.
//
// No changes are made to the validator set 'vals'.
func computeNewPriorities(updates []*Validator, vals *ValidatorSet, updatedTotalVotingPower int64) {
	_ = "STUB: not implemented"
	return
}

// add val
// Set ProposerPriority to -C*totalVotingPower (with C ~= 1.125) to make sure validators can't
// un-bond and then re-bond to reset their (potentially previously negative) ProposerPriority to zero.
//
// Contract: updatedVotingPower < 2 * MaxTotalVotingPower to ensure ProposerPriority does
// not exceed the bounds of int64.
//
// Compute ProposerPriority = -1.125*totalVotingPower == -(updatedVotingPower + (updatedVotingPower >> 3)).

// Merges the vals' validator list with the updates list.
// When two elements with same address are seen, the one from updates is selected.
// Expects updates to be a list of updates sorted by address with no duplicates or errors,
// must have been validated with verifyUpdates() and priorities computed with computeNewPriorities().
func (vals *ValidatorSet) applyUpdates(updates []*Validator) { _ = "STUB: not implemented"; return }

// unchanged validator

// Apply add or update.

// Validator is present in both, advance existing.

// Add the elements which are left.

// OR add updates which are left.

// Checks that the validators to be removed are part of the validator set.
// No changes are made to the validator set 'vals'.
func verifyRemovals(deletes []*Validator, vals *ValidatorSet) (votingPower int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Removes the validators specified in 'deletes' from validator set 'vals'.
// Should not fail as verification has been done before.
// Expects vals to be sorted by address (done by applyUpdates).
func (vals *ValidatorSet) applyRemovals(deletes []*Validator) { _ = "STUB: not implemented"; return }

// Loop over deletes until we removed all of them.

// Leave it in the resulting slice.

// Add the elements which are left.

// Main function used by UpdateWithChangeSet() and NewValidatorSet().
// If 'allowDeletes' is false then delete operations (identified by validators with voting power 0)
// are not allowed and will trigger an error if present in 'changes'.
// The 'allowDeletes' flag is set to false by NewValidatorSet() and to true by UpdateWithChangeSet().
func (vals *ValidatorSet) updateWithChangeSet(changes []*Validator, allowDeletes bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Check for duplicates within changes, split in 'updates' and 'deletes' lists (sorted).

// Check that the resulting set will not be empty.

// Verify that applying the 'deletes' against 'vals' will not result in error.
// Get the voting power that is going to be removed.

// Verify that applying the 'updates' against 'vals' will not result in error.
// Get the updated total voting power before removal. Note that this is < 2 * MaxTotalVotingPower

// Compute the priorities for updates.

// Apply updates and removals.

// Should go after additions.

// Scale and center.

// UpdateWithChangeSet attempts to update the validator set with 'changes'.
// It performs the following steps:
//   - validates the changes making sure there are no duplicates and splits them in updates and deletes
//   - verifies that applying the changes will not result in errors
//   - computes the total voting power BEFORE removals to ensure that in the next steps the priorities
//     across old and newly added validators are fair
//   - computes the priorities of new validators against the final set
//   - applies the updates against the validator set
//   - applies the removals against the validator set
//   - performs scaling and centering of priority values
//
// If an error is detected during verification steps, it is returned and the validator set
// is not changed.
func (vals *ValidatorSet) UpdateWithChangeSet(changes []*Validator) error {
	_ = "STUB: not implemented"
	return nil
}

// VerifyCommit verifies +2/3 of the set had signed the given commit and all
// other signatures are valid
func (vals *ValidatorSet) VerifyCommit(chainID string, blockID BlockID,
	height int64, commit *Commit,
) error {
	_ = "STUB: not implemented"
	return nil
}

// LIGHT CLIENT VERIFICATION METHODS

// VerifyCommitLight verifies +2/3 of the set had signed the given commit.
// It does NOT count all signatures.
func (vals *ValidatorSet) VerifyCommitLight(chainID string, blockID BlockID,
	height int64, commit *Commit,
) error {
	_ = "STUB: not implemented"
	return nil
}

// VerifyCommitLight verifies +2/3 of the set had signed the given commit.
// It DOES count all signatures.
func (vals *ValidatorSet) VerifyCommitLightAllSignatures(chainID string, blockID BlockID,
	height int64, commit *Commit,
) error {
	_ = "STUB: not implemented"
	return nil
}

// VerifyCommitLightTrusting verifies that trustLevel of the validator set signed
// this commit.
// It does NOT count all signatures.
func (vals *ValidatorSet) VerifyCommitLightTrusting(
	chainID string,
	commit *Commit,
	trustLevel cmtmath.Fraction,
) error {
	_ = "STUB: not implemented"
	return nil
}

// VerifyCommitLightTrusting verifies that trustLevel of the validator set signed
// this commit.
// It DOES count all signatures.
func (vals *ValidatorSet) VerifyCommitLightTrustingAllSignatures(
	chainID string,
	commit *Commit,
	trustLevel cmtmath.Fraction,
) error {
	_ = "STUB: not implemented"
	return nil
}

// findPreviousProposer reverses the compare proposer priority function to find the validator
// with the lowest proposer priority which would have been the previous proposer.
//
// Is used when recreating a validator set from an existing array of validators.
func (vals *ValidatorSet) findPreviousProposer() *Validator { _ = "STUB: not implemented"; return nil }

func (vals *ValidatorSet) checkAllKeysHaveSameType() { _ = "STUB: not implemented"; return }

// XXX: Should only be the case in tests.

// AllKeysHaveSameType returns true if all validators have the same type of
// public key or if the set is empty.
func (vals *ValidatorSet) AllKeysHaveSameType() bool { _ = "STUB: not implemented"; return false }

// -----------------

// IsErrNotEnoughVotingPowerSigned returns true if err is
// ErrNotEnoughVotingPowerSigned.
func IsErrNotEnoughVotingPowerSigned(err error) bool { _ = "STUB: not implemented"; return false }

// ErrNotEnoughVotingPowerSigned is returned when not enough validators signed
// a commit.
type ErrNotEnoughVotingPowerSigned struct {
	Got    int64
	Needed int64
}

func (e ErrNotEnoughVotingPowerSigned) Error() string { _ = "STUB: not implemented"; return "" }

//----------------

// String returns a string representation of ValidatorSet.
//
// See StringIndented.
func (vals *ValidatorSet) String() string { _ = "STUB: not implemented"; return "" }

// StringIndented returns an intended String.
//
// See Validator#String.
func (vals *ValidatorSet) StringIndented(indent string) string {
	_ = "STUB: not implemented"
	return ""
}

//-------------------------------------

// ValidatorsByVotingPower implements sort.Interface for []*Validator based on
// the VotingPower and Address fields.
type ValidatorsByVotingPower []*Validator

func (valz ValidatorsByVotingPower) Len() int { _ = "STUB: not implemented"; return 0 }

func (valz ValidatorsByVotingPower) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (valz ValidatorsByVotingPower) Swap(i, j int) { _ = "STUB: not implemented"; return }

// ValidatorsByAddress implements sort.Interface for []*Validator based on
// the Address field.
type ValidatorsByAddress []*Validator

func (valz ValidatorsByAddress) Len() int { _ = "STUB: not implemented"; return 0 }

func (valz ValidatorsByAddress) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (valz ValidatorsByAddress) Swap(i, j int) { _ = "STUB: not implemented"; return }

// ToProto converts ValidatorSet to protobuf
func (vals *ValidatorSet) ToProto() (*cmtproto.ValidatorSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// validator set should never be nil

// NOTE: Sometimes we use the bytes of the proto form as a hash. This means that we need to
// be consistent with cached data

// ValidatorSetFromProto sets a protobuf ValidatorSet to the given pointer.
// It returns an error if any of the validators from the set or the proposer
// is invalid
func ValidatorSetFromProto(vp *cmtproto.ValidatorSet) (*ValidatorSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// validator set should never be nil, bigger issues are at play if empty

// NOTE: We can't trust the total voting power given to us by other peers. If someone were to
// inject a non-zeo value that wasn't the correct voting power we could assume a wrong total
// power hence we need to recompute it.
// FIXME: We should look to remove TotalVotingPower from proto or add it in the validators hash
// so we don't have to do this
// NOTE: Use TotalVotingPowerSafe to return error on invalid input.

// ValidatorSetFromExistingValidators takes an existing array of validators and
// rebuilds the exact same validator set that corresponds to it without
// changing the proposer priority or power if any of the validators fail
// validate basic then an empty set is returned.
func ValidatorSetFromExistingValidators(valz []*Validator) (*ValidatorSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//----------------------------------------

// RandValidatorSet returns a randomized validator set (size: +numValidators+),
// where each validator has a voting power of +votingPower+.
//
// EXPOSED FOR TESTING.
func RandValidatorSet(numValidators int, votingPower int64) (*ValidatorSet, []PrivValidator) {
	_ = "STUB: not implemented"
	return nil, nil
}

// safe addition/subtraction/multiplication

func safeAdd(a, b int64) (int64, bool) { _ = "STUB: not implemented"; return 0, false }

func safeSub(a, b int64) (int64, bool) { _ = "STUB: not implemented"; return 0, false }

func safeAddClip(a, b int64) int64 { _ = "STUB: not implemented"; return 0 }

func safeSubClip(a, b int64) int64 { _ = "STUB: not implemented"; return 0 }

func safeMul(a, b int64) (int64, bool) { _ = "STUB: not implemented"; return 0, false }
