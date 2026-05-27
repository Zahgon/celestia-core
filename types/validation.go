package types

import (
	cmtmath "github.com/cometbft/cometbft/libs/math"
)

const batchVerifyThreshold = 2

func shouldBatchVerify(vals *ValidatorSet, commit *Commit) bool {
	_ = "STUB: not implemented"
	return false
}

// VerifyCommit verifies +2/3 of the set had signed the given commit.
//
// It checks all the signatures! While it's safe to exit as soon as we have
// 2/3+ signatures, doing so would impact incentivization logic in the ABCI
// application that depends on the LastCommitInfo sent in FinalizeBlock, which
// includes which validators signed. For instance, Gaia incentivizes proposers
// with a bonus for including more than +2/3 of the signatures.
func VerifyCommit(chainID string, vals *ValidatorSet, blockID BlockID,
	height int64, commit *Commit) error {
	_ = "STUB: not implemented"
	// run a basic validation of the arguments
	return nil
}

// calculate voting power needed. Note that total voting power is capped to
// 1/8th of max int64 so this operation should never overflow

// ignore all absent signatures

// only count the signatures that are for the block

// attempt to batch verify

// if verification failed or is not supported then fallback to single verification

// LIGHT CLIENT VERIFICATION METHODS

// VerifyCommitLight verifies +2/3 of the set had signed the given commit.
//
// This method is primarily used by the light client and does NOT check all the
// signatures.
func VerifyCommitLight(
	chainID string,
	vals *ValidatorSet,
	blockID BlockID,
	height int64,
	commit *Commit,
) error {
	_ = "STUB: not implemented"
	return nil
}

// VerifyCommitLightAllSignatures verifies +2/3 of the set had signed the given commit.
//
// This method DOES check all the signatures.
func VerifyCommitLightAllSignatures(
	chainID string,
	vals *ValidatorSet,
	blockID BlockID,
	height int64,
	commit *Commit,
) error {
	_ = "STUB: not implemented"
	return nil
}

func verifyCommitLightInternal(
	chainID string,
	vals *ValidatorSet,
	blockID BlockID,
	height int64,
	commit *Commit,
	countAllSignatures bool,
) error {
	_ = "STUB: not implemented"
	// run a basic validation of the arguments
	return nil
}

// calculate voting power needed

// ignore all commit signatures that are not for the block

// count all the remaining signatures

// attempt to batch verify

// if verification failed or is not supported then fallback to single verification

// VerifyCommitLightTrusting verifies that trustLevel of the validator set signed
// this commit.
//
// NOTE the given validators do not necessarily correspond to the validator set
// for this commit, but there may be some intersection.
//
// This method is primarily used by the light client and does NOT check all the
// signatures.
func VerifyCommitLightTrusting(
	chainID string,
	vals *ValidatorSet,
	commit *Commit,
	trustLevel cmtmath.Fraction,
) error {
	_ = "STUB: not implemented"
	return nil
}

// VerifyCommitLightTrustingAllSignatures verifies that trustLevel of the validator
// set signed this commit.
//
// NOTE the given validators do not necessarily correspond to the validator set
// for this commit, but there may be some intersection.
//
// This method DOES check all the signatures.
func VerifyCommitLightTrustingAllSignatures(
	chainID string,
	vals *ValidatorSet,
	commit *Commit,
	trustLevel cmtmath.Fraction,
) error {
	_ = "STUB: not implemented"
	return nil
}

func verifyCommitLightTrustingInternal(
	chainID string,
	vals *ValidatorSet,
	commit *Commit,
	trustLevel cmtmath.Fraction,
	countAllSignatures bool,
) error {
	_ = "STUB: not implemented"
	// sanity checks
	return nil
}

// safely calculate voting power needed.

// ignore all commit signatures that are not for the block

// count all the remaining signatures

// attempt to batch verify commit. As the validator set doesn't necessarily
// correspond with the validator set that signed the block we need to look
// up by address rather than index.

// attempt with single verification

// ValidateHash returns an error if the hash is not empty, but its
// size != tmhash.Size.
func ValidateHash(h []byte) error { _ = "STUB: not implemented"; return nil }

// Batch verification

// verifyCommitBatch batch verifies commits.  This routine is equivalent
// to verifyCommitSingle in behavior, just faster iff every signature in the
// batch is valid.
//
// Note: The caller is responsible for checking to see if this routine is
// usable via `shouldVerifyBatch(vals, commit)`.
func verifyCommitBatch(
	chainID string,
	vals *ValidatorSet,
	commit *Commit,
	votingPowerNeeded int64,
	ignoreSig func(CommitSig) bool,
	countSig func(CommitSig) bool,
	countAllSignatures bool,
	lookUpByIndex bool,
) error {
	_ = "STUB: not implemented"
	return nil
}

// attempt to create a batch verifier

// re-check if batch verification is supported

// This should *NEVER* happen.

// skip over signatures that should be ignored

// If the vals and commit have a 1-to-1 correspondance we can retrieve
// them by index else we need to retrieve them by address

// if the signature doesn't belong to anyone in the validator set
// then we just skip over it

// because we are getting validators by address we need to make sure
// that the same validator doesn't commit twice

// Validate signature.

// add the key, sig and message to the verifier

// If this signature counts then add the voting power of the validator
// to the tally

// if we don't need to verify all signatures and already have sufficient
// voting power we can break from batching and verify all the signatures

// ensure that we have batched together enough signatures to exceed the
// voting power needed else there is no need to even verify

// attempt to verify the batch.

// success

// one or more of the signatures is invalid, find and return the first
// invalid signature.

// go back from the batch index to the commit.Signatures index

// execution reaching here is a bug, and one of the following has
// happened:
//  * non-zero tallied voting power, empty batch (impossible?)
//  * bv.Verify() returned `false, []bool{true, ..., true}` (BUG)

// Single Verification

// verifyCommitSingle single verifies commits.
// If a key does not support batch verification, or batch verification fails this will be used
// This method is used to check all the signatures included in a commit.
// It is used in consensus for validating a block LastCommit.
// CONTRACT: both commit and validator set should have passed validate basic
func verifyCommitSingle(
	chainID string,
	vals *ValidatorSet,
	commit *Commit,
	votingPowerNeeded int64,
	ignoreSig func(CommitSig) bool,
	countSig func(CommitSig) bool,
	countAllSignatures bool,
	lookUpByIndex bool,
) error {
	_ = "STUB: not implemented"
	return nil
}

// If the vals and commit have a 1-to-1 correspondence we can retrieve
// them by index else we need to retrieve them by address

// if the signature doesn't belong to anyone in the validator set
// then we just skip over it

// because we are getting validators by address we need to make sure
// that the same validator doesn't commit twice

// If this signature counts then add the voting power of the validator
// to the tally

// check if we have enough signatures and can thus exit early

func verifyBasicValsAndCommit(vals *ValidatorSet, commit *Commit, height int64, blockID BlockID) error {
	_ = "STUB: not implemented"
	return nil
}

// Validate Height and BlockID.
