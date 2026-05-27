package types

import (
	"github.com/cometbft/cometbft/types"
)

// UnmarshalledTx is an intermediary type that allows keeping the transaction
// metadata, its Key and the actual tx bytes. This will be used to create the
// parts from the local txs.
type UnmarshalledTx struct {
	MetaData TxMetaData
	Key      types.TxKey
	TxBytes  []byte
}

func TxsToParts(txs []UnmarshalledTx, partCount, partSize, lastPartLen uint32) ([]*types.Part, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Adjust for a short final part if necessary.

// No overlap.

// Defensive check: Ensure the computed bounds are within the slice lengths.
