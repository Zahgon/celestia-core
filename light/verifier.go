package light

import (
	"time"

	cmtmath "github.com/cometbft/cometbft/libs/math"
	"github.com/cometbft/cometbft/types"
)

var (
	// DefaultTrustLevel - new header can be trusted if at least one correct
	// validator signed it.
	DefaultTrustLevel = cmtmath.Fraction{Numerator: 1, Denominator: 3}
)

// VerifyNonAdjacent verifies non-adjacent untrustedHeader against
// trustedHeader. It ensures that:
//
//		a) trustedHeader can still be trusted (if not, ErrOldHeaderExpired is returned)
//		b) untrustedHeader is valid (if not, ErrInvalidHeader is returned)
//		c) trustLevel ([1/3, 1]) of trustedHeaderVals (or trustedHeaderNextVals)
//	 signed correctly (if not, ErrNewValSetCantBeTrusted is returned)
//		d) more than 2/3 of untrustedVals have signed h2
//	   (otherwise, ErrInvalidHeader is returned)
//	 e) headers are non-adjacent.
//
// maxClockDrift defines how much untrustedHeader.Time can drift into the
// future.
func VerifyNonAdjacent(
	trustedHeader *types.SignedHeader, // height=X
	trustedVals *types.ValidatorSet, // height=X or height=X+1
	untrustedHeader *types.SignedHeader, // height=Y
	untrustedVals *types.ValidatorSet, // height=Y
	trustingPeriod time.Duration,
	now time.Time,
	maxClockDrift time.Duration,
	trustLevel cmtmath.Fraction) error {
	_ = "STUB: not implemented"
	return nil
}

// Ensure that +`trustLevel` (default 1/3) or more of last trusted validators signed correctly.

// Ensure that +2/3 of new validators signed correctly.
//
// NOTE: this should always be the last check because untrustedVals can be
// intentionally made very large to DOS the light client. not the case for
// VerifyAdjacent, where validator set is known in advance.

// VerifyAdjacent verifies directly adjacent untrustedHeader against
// trustedHeader. It ensures that:
//
//	a) trustedHeader can still be trusted (if not, ErrOldHeaderExpired is returned)
//	b) untrustedHeader is valid (if not, ErrInvalidHeader is returned)
//	c) untrustedHeader.ValidatorsHash equals trustedHeader.NextValidatorsHash
//	d) more than 2/3 of new validators (untrustedVals) have signed h2
//	  (otherwise, ErrInvalidHeader is returned)
//	e) headers are adjacent.
//
// maxClockDrift defines how much untrustedHeader.Time can drift into the
// future.
func VerifyAdjacent(
	trustedHeader *types.SignedHeader, // height=X
	untrustedHeader *types.SignedHeader, // height=X+1
	untrustedVals *types.ValidatorSet, // height=X+1
	trustingPeriod time.Duration,
	now time.Time,
	maxClockDrift time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// Check the validator hashes are the same

// Ensure that +2/3 of new validators signed correctly.

// Verify combines both VerifyAdjacent and VerifyNonAdjacent functions.
func Verify(
	trustedHeader *types.SignedHeader, // height=X
	trustedVals *types.ValidatorSet, // height=X or height=X+1
	untrustedHeader *types.SignedHeader, // height=Y
	untrustedVals *types.ValidatorSet, // height=Y
	trustingPeriod time.Duration,
	now time.Time,
	maxClockDrift time.Duration,
	trustLevel cmtmath.Fraction) error {
	_ = "STUB: not implemented"
	return nil
}

func verifyNewHeaderAndVals(
	untrustedHeader *types.SignedHeader,
	untrustedVals *types.ValidatorSet,
	trustedHeader *types.SignedHeader,
	now time.Time,
	maxClockDrift time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// ValidateTrustLevel checks that trustLevel is within the allowed range [1/3,
// 1]. If not, it returns an error. 1/3 is the minimum amount of trust needed
// which does not break the security model.
func ValidateTrustLevel(lvl cmtmath.Fraction) error { _ = "STUB: not implemented"; return nil }

// < 1/3
// > 1

// HeaderExpired return true if the given header expired.
func HeaderExpired(h *types.SignedHeader, trustingPeriod time.Duration, now time.Time) bool {
	_ = "STUB: not implemented"
	return false
}

// VerifyBackwards verifies an untrusted header with a height one less than
// that of an adjacent trusted header. It ensures that:
//
//		a) untrusted header is valid
//	 b) untrusted header has a time before the trusted header
//	 c) that the LastBlockID hash of the trusted header is the same as the hash
//	 of the trusted header
//
//	 For any of these cases ErrInvalidHeader is returned.
func VerifyBackwards(untrustedHeader, trustedHeader *types.Header) error {
	_ = "STUB: not implemented"
	return nil
}
