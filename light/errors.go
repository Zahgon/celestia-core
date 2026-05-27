package light

import (
	"errors"
	"time"

	"github.com/cometbft/cometbft/types"
)

// ErrOldHeaderExpired means the old (trusted) header has expired according to
// the given trustingPeriod and current time. If so, the light client must be
// reset subjectively.
type ErrOldHeaderExpired struct {
	At  time.Time
	Now time.Time
}

func (e ErrOldHeaderExpired) Error() string { _ = "STUB: not implemented"; return "" }

// ErrNewValSetCantBeTrusted means the new validator set cannot be trusted
// because < 1/3rd (+trustLevel+) of the old validator set has signed.
type ErrNewValSetCantBeTrusted struct {
	Reason types.ErrNotEnoughVotingPowerSigned
}

func (e ErrNewValSetCantBeTrusted) Error() string { _ = "STUB: not implemented"; return "" }

// ErrInvalidHeader means the header either failed the basic validation or
// commit is not signed by 2/3+.
type ErrInvalidHeader struct {
	Reason error
}

func (e ErrInvalidHeader) Error() string { _ = "STUB: not implemented"; return "" }

// ErrFailedHeaderCrossReferencing is returned when the detector was not able to cross reference the header
// with any of the connected witnesses.
var ErrFailedHeaderCrossReferencing = errors.New("all witnesses have either not responded, don't have the " +
	" blocks or sent invalid blocks. You should look to change your witnesses" +
	"  or review the light client's logs for more information")

// ErrVerificationFailed means either sequential or skipping verification has
// failed to verify from header #1 to header #2 due to some reason.
type ErrVerificationFailed struct {
	From   int64
	To     int64
	Reason error
}

// Unwrap returns underlying reason.
func (e ErrVerificationFailed) Unwrap() error { _ = "STUB: not implemented"; return nil }

func (e ErrVerificationFailed) Error() string { _ = "STUB: not implemented"; return "" }

// ErrLightClientAttack is returned when the light client has detected an attempt
// to verify a false header and has sent the evidence to either a witness or primary.
var ErrLightClientAttack = errors.New(`attempted attack detected.
	Light client received valid conflicting header from witness.
	Unable to verify header. Evidence has been sent to both providers.
	Check logs for full evidence and trace`,
)

// ErrNoWitnesses means that there are not enough witnesses connected to
// continue running the light client.
var ErrNoWitnesses = errors.New("no witnesses connected. please reset light client")

// ErrConflictingHeaders is thrown when two conflicting headers are discovered.
type ErrConflictingHeaders struct {
	Block        *types.LightBlock
	WitnessIndex int
}

func (e ErrConflictingHeaders) Error() string { _ = "STUB: not implemented"; return "" }

// ErrProposerPrioritiesDiverge is thrown when two conflicting headers are
// discovered, but the error is non-attributable comparing to ErrConflictingHeaders.
// The difference is in validator set proposer priorities, which may change
// with every round of consensus.
type ErrProposerPrioritiesDiverge struct {
	WitnessHash  []byte
	WitnessIndex int
	PrimaryHash  []byte
}

func (e ErrProposerPrioritiesDiverge) Error() string { _ = "STUB: not implemented"; return "" }

// ----------------------------- INTERNAL ERRORS ---------------------------------

// errBadWitness is returned when the witness either does not respond or
// responds with an invalid header.
type errBadWitness struct {
	Reason       error
	WitnessIndex int
}

func (e errBadWitness) Error() string { _ = "STUB: not implemented"; return "" }

var errNoDivergence = errors.New(
	"sanity check failed: no divergence between the original trace and the provider's new trace",
)
