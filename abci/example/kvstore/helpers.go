package kvstore

import (
	"context"

	"github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/proto/tendermint/crypto"
)

// RandVal creates one random validator, with a key derived
// from the input value
func RandVal() types.ValidatorUpdate { _ = "STUB: not implemented"; return *new(types.ValidatorUpdate) }

// RandVals returns a list of cnt validators for initializing
// the application. Note that the keys are deterministically
// derived from the index in the array, while the power is
// random (Change this if not desired)
func RandVals(cnt int) []types.ValidatorUpdate { _ = "STUB: not implemented"; return nil }

// InitKVStore initializes the kvstore app with some data,
// which allows tests to pass and is fine as long as you
// don't make any tx that modify the validator state
func InitKVStore(ctx context.Context, app *Application) error {
	_ = "STUB: not implemented"
	return nil
}

// Create a new transaction
func NewTx(key, value string) []byte { _ = "STUB: not implemented"; return nil }

func NewRandomTx(size int) []byte { _ = "STUB: not implemented"; return nil }

func NewRandomTxs(n int) [][]byte { _ = "STUB: not implemented"; return nil }

func NewTxFromID(i int) []byte { _ = "STUB: not implemented"; return nil }

// Create a transaction to add/remove/update a validator
// To remove, set power to 0.
func MakeValSetChangeTx(pubkey crypto.PublicKey, power int64) []byte {
	_ = "STUB: not implemented"
	return nil
}
