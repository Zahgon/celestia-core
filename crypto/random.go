package crypto

import (
	"io"
)

// This only uses the OS's randomness
func randBytes(numBytes int) []byte { _ = "STUB: not implemented"; return nil }

// This only uses the OS's randomness
func CRandBytes(numBytes int) []byte { _ = "STUB: not implemented"; return nil }

// CRandHex returns a hex encoded string that's floor(numDigits/2) * 2 long.
//
// Note: CRandHex(24) gives 96 bits of randomness that
// are usually strong enough for most purposes.
func CRandHex(numDigits int) string { _ = "STUB: not implemented"; return "" }

// Returns a crand.Reader.
func CReader() io.Reader { _ = "STUB: not implemented"; return *new(io.Reader) }
