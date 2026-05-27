package evidence

import (
	"time"

	"github.com/cometbft/cometbft/types"
)

// verify verifies the evidence fully by checking:
// - It has not already been committed
// - it is sufficiently recent (MaxAge)
// - it is from a key who was a validator at the given height
// - it is internally consistent with state
// - it was properly signed by the alleged equivocator and meets the individual evidence verification requirements
func (evpool *Pool) verify(evidence types.Evidence) error { _ = "STUB: not implemented"; return nil }

// verify the time of the evidence

// checking if evidence is expired calculated using the block evidence time and height

// apply the evidence-specific verification logic

// in the case of lunatic the trusted header is different to the common header

// FIXME: This multi step process is a bit unergonomic. We may want to consider a more efficient process
// that doesn't require as much io and is atomic.

// If the node doesn't have a block at the height of the conflicting block, then this could be
// a forward lunatic attack. Thus the node must get the latest height it has

// VerifyLightClientAttack verifies LightClientAttackEvidence against the state of the full node. This involves
// the following checks:
//   - the common header from the full node has at least 1/3 voting power which is also present in
//     the conflicting header's commit
//   - 2/3+ of the conflicting validator set correctly signed the conflicting block
//   - the nodes trusted header at the same height as the conflicting header has a different hash
//   - all signatures must be checked as this will be used as evidence
//
// CONTRACT: must run ValidateBasic() on the evidence before verifying
//
//	must check that the evidence has not expired (i.e. is outside the maximum age threshold)
func VerifyLightClientAttack(
	e *types.LightClientAttackEvidence,
	commonHeader, trustedHeader *types.SignedHeader,
	commonVals *types.ValidatorSet,
	now time.Time, //nolint:revive
	trustPeriod time.Duration, //nolint:revive
) error {
	_ = "STUB: not implemented"
	// TODO: Should the current time and trust period be used in this method?
	// If not, why were the parameters present?
	return nil
}

// In the case of lunatic attack there will be a different commonHeader height. Therefore the node perform a single
// verification jump between the common header and the conflicting one

// In the case of equivocation and amnesia we expect all header hashes to be correctly derived

// Verify that the 2/3+ commits from the conflicting validator set were for the conflicting header

// Assert the correct amount of voting power of the validator set

// check in the case of a forward lunatic attack that monotonically increasing time has been violated

// In all other cases check that the hashes of the conflicting header and the trusted header are different

// VerifyDuplicateVote verifies DuplicateVoteEvidence against the state of full node. This involves the
// following checks:
//   - the validator is in the validator set at the height of the evidence
//   - the height, round, type and validator address of the votes must be the same
//   - the block ID's must be different
//   - The signatures must both be valid
func VerifyDuplicateVote(e *types.DuplicateVoteEvidence, chainID string, valSet *types.ValidatorSet) error {
	_ = "STUB: not implemented"
	return nil
}

// H/R/S must be the same

// Address must be the same

// BlockIDs must be different

// pubkey must match address (this should already be true, sanity check)

// validator voting power and total voting power must match

// Signatures must be valid

// validateABCIEvidence validates the ABCI component of the light client attack
// evidence i.e voting power and byzantine validators
func validateABCIEvidence(
	ev *types.LightClientAttackEvidence,
	commonVals *types.ValidatorSet,
	trustedHeader *types.SignedHeader,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Find out what type of attack this was and thus extract the malicious
// validators. Note, in the case of an Amnesia attack we don't have any
// malicious validators.

// Ensure this matches the validators that are listed in the evidence. They
// should be ordered based on power.

func getSignedHeader(blockStore BlockStore, height int64) (*types.SignedHeader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// check that the evidence hasn't expired
func IsEvidenceExpired(heightNow int64, timeNow time.Time, heightEv int64, timeEv time.Time, evidenceParams types.EvidenceParams) bool {
	_ = "STUB: not implemented"
	return false
}
