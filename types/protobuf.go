package types

import (
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/crypto"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
)

//-------------------------------------------------------

// TM2PB is used for converting CometBFT ABCI to protobuf ABCI.
// UNSTABLE
var TM2PB = tm2pb{}

type tm2pb struct{}

func (tm2pb) Header(header *Header) cmtproto.Header {
	_ = "STUB: not implemented"
	return *new(cmtproto.Header)
}

func (tm2pb) Validator(val *Validator) abci.Validator {
	_ = "STUB: not implemented"
	return *new(abci.Validator)
}

func (tm2pb) BlockID(blockID BlockID) cmtproto.BlockID {
	_ = "STUB: not implemented"
	return *new(cmtproto.BlockID)
}

func (tm2pb) PartSetHeader(header PartSetHeader) cmtproto.PartSetHeader {
	_ = "STUB: not implemented"
	return *new(cmtproto.PartSetHeader)
}

// XXX: panics on unknown pubkey type
func (tm2pb) ValidatorUpdate(val *Validator) abci.ValidatorUpdate {
	_ = "STUB: not implemented"
	return *new(abci.ValidatorUpdate)
}

// XXX: panics on nil or unknown pubkey type
func (tm2pb) ValidatorUpdates(vals *ValidatorSet) []abci.ValidatorUpdate {
	_ = "STUB: not implemented"
	return nil
}

// XXX: panics on nil or unknown pubkey type
func (tm2pb) NewValidatorUpdate(pubkey crypto.PubKey, power int64) abci.ValidatorUpdate {
	_ = "STUB: not implemented"
	return *new(abci.ValidatorUpdate)
}

//----------------------------------------------------------------------------

// PB2TM is used for converting protobuf ABCI to CometBFT ABCI.
// UNSTABLE
var PB2TM = pb2tm{}

type pb2tm struct{}

func (pb2tm) ValidatorUpdates(vals []abci.ValidatorUpdate) ([]*Validator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
