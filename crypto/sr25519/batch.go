package sr25519

import (
	"github.com/oasisprotocol/curve25519-voi/primitives/sr25519"

	"github.com/cometbft/cometbft/crypto"
)

var _ crypto.BatchVerifier = &BatchVerifier{}

// BatchVerifier implements batch verification for sr25519.
type BatchVerifier struct {
	*sr25519.BatchVerifier
}

func NewBatchVerifier() crypto.BatchVerifier {
	_ = "STUB: not implemented"
	return *new(crypto.BatchVerifier)
}

func (b *BatchVerifier) Add(key crypto.PubKey, msg, signature []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *BatchVerifier) Verify() (bool, []bool) { _ = "STUB: not implemented"; return false, nil }
