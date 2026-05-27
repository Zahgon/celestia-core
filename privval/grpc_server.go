package privval

import (
	"context"

	"github.com/cometbft/cometbft/libs/log"
	privvalproto "github.com/cometbft/cometbft/proto/tendermint/privval"
	"github.com/cometbft/cometbft/types"
)

// PrivValidatorGRPCServer implements the PrivValidatorAPIServer gRPC interface
// by forwarding signing requests to an underlying types.PrivValidator.
type PrivValidatorGRPCServer struct {
	privVal types.PrivValidator
	logger  log.Logger
}

// NewPrivValidatorGRPCServer returns a new gRPC server that wraps the given PrivValidator.
func NewPrivValidatorGRPCServer(
	privVal types.PrivValidator,
	logger log.Logger,
) *PrivValidatorGRPCServer {
	_ = "STUB: not implemented"
	return nil
}

// SignRawBytes forwards a raw bytes signing request to the underlying PrivValidator.
func (s *PrivValidatorGRPCServer) SignRawBytes(
	_ context.Context,
	req *privvalproto.SignRawBytesRequest,
) (*privvalproto.SignedRawBytesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *PrivValidatorGRPCServer) GetPubKey(
	_ context.Context,
	req *privvalproto.PubKeyRequest,
) (*privvalproto.PubKeyResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
