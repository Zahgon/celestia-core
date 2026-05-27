package consensus

import (
	"github.com/cosmos/gogoproto/proto"

	cmtcons "github.com/cometbft/cometbft/proto/tendermint/consensus"
)

// MsgToProto takes a consensus message type and returns the proto defined consensus message.
//
// TODO: This needs to be removed, but WALToProto depends on this.
func MsgToProto(msg Message) (proto.Message, error) {
	_ = "STUB: not implemented"
	return *new(proto.Message), nil
}

// MsgFromProto takes a consensus proto message and returns the native go type
func MsgFromProto(p proto.Message) (Message, error) {
	_ = "STUB: not implemented"
	return *new(Message), nil
}

// deny message based on possible overflow

// Vote validation will be handled in the vote message ValidateBasic
// call below.

// WALToProto takes a WAL message and return a proto walMessage and error
func WALToProto(msg WALMessage) (*cmtcons.WALMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// WALFromProto takes a proto wal message and return a consensus walMessage and error
func WALFromProto(msg *cmtcons.WALMessage) (WALMessage, error) {
	_ = "STUB: not implemented"
	return *new(WALMessage), nil
}

// deny message based on possible overflow
