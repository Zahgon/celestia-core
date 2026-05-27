package batch

import (
	"github.com/cometbft/cometbft/crypto"
)

// CreateBatchVerifier checks if a key type implements the batch verifier interface.
// Currently only ed25519 & sr25519 supports batch verification.
func CreateBatchVerifier(pk crypto.PubKey) (crypto.BatchVerifier, bool) {
	_ = "STUB: not implemented"
	return *new(crypto.BatchVerifier), false
}

// case where the key does not support batch verification

// SupportsBatchVerifier checks if a key type implements the batch verifier
// interface.
func SupportsBatchVerifier(pk crypto.PubKey) bool { _ = "STUB: not implemented"; return false }
