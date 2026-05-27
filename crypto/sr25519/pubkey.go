package sr25519

import (
	"github.com/cometbft/cometbft/crypto"
)

var _ crypto.PubKey = PubKey{}

const (
	// PubKeySize is the number of bytes in an Sr25519 public key.
	PubKeySize = 32

	// SignatureSize is the size of a Sr25519 signature in bytes.
	SignatureSize = 64
)

// PubKey implements crypto.PubKey for the Sr25519 signature scheme.
type PubKey []byte

// Address is the SHA256-20 of the raw pubkey bytes.
func (pubKey PubKey) Address() crypto.Address {
	_ = "STUB: not implemented"
	return *new(crypto.Address)
}

// Bytes returns the byte representation of the PubKey.
func (pubKey PubKey) Bytes() []byte { _ = "STUB: not implemented"; return nil }

// Equals - checks that two public keys are the same time
// Runs in constant time based on length of the keys.
func (pubKey PubKey) Equals(other crypto.PubKey) bool { _ = "STUB: not implemented"; return false }

func (pubKey PubKey) VerifySignature(msg []byte, sigBytes []byte) bool {
	_ = "STUB: not implemented"
	return false
}

func (pubKey PubKey) String() string { _ = "STUB: not implemented"; return "" }

func (pubKey PubKey) Type() string { _ = "STUB: not implemented"; return "" }
