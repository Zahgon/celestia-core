package payload

const keyPrefix = "a="
const maxPayloadSize = 4 * 1024 * 1024

// NewBytes generates a new payload and returns the encoded representation of
// the payload as a slice of bytes. NewBytes uses the fields on the Options
// to create the payload.
func NewBytes(p *Payload) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// #nosec -- The "if" above makes this cast safe

// We halve the padding size because we transform the TX to hex

// prepend a single key so that the kv store only ever stores a single
// transaction instead of storing all tx and ballooning in size.

// FromBytes extracts a paylod from the byte representation of the payload.
// FromBytes leaves the padding untouched, returning it to the caller to handle
// or discard per their preference.
func FromBytes(b []byte) (*Payload, error) { _ = "STUB: not implemented"; return nil, nil }

// MaxUnpaddedSize returns the maximum size that a payload may be if no padding
// is included.
func MaxUnpaddedSize() (int, error) { _ = "STUB: not implemented"; return 0, nil }

// CalculateUnpaddedSize calculates the size of the passed in payload for the
// purpose of determining how much padding to add to add to reach the target size.
// CalculateUnpaddedSize returns an error if the payload Padding field is longer than 1.
func CalculateUnpaddedSize(p *Payload) (int, error) { _ = "STUB: not implemented"; return 0, nil }
