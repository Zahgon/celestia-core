package privval

import (
	"github.com/cometbft/cometbft/libs/service"
	cmtsync "github.com/cometbft/cometbft/libs/sync"
	privvalproto "github.com/cometbft/cometbft/proto/tendermint/privval"
	"github.com/cometbft/cometbft/types"
)

// ValidationRequestHandlerFunc handles different remoteSigner requests
type ValidationRequestHandlerFunc func(
	privVal types.PrivValidator,
	requestMessage privvalproto.Message,
	chainID string) (privvalproto.Message, error)

type SignerServer struct {
	service.BaseService

	endpoint *SignerDialerEndpoint
	chainID  string
	privVal  types.PrivValidator

	handlerMtx               cmtsync.Mutex
	validationRequestHandler ValidationRequestHandlerFunc
}

func NewSignerServer(endpoint *SignerDialerEndpoint, chainID string, privVal types.PrivValidator) *SignerServer {
	_ = "STUB: not implemented"
	return nil
}

// OnStart implements service.Service.
func (ss *SignerServer) OnStart() error { _ = "STUB: not implemented"; return nil }

// OnStop implements service.Service.
func (ss *SignerServer) OnStop() { _ = "STUB: not implemented"; return }

// SetRequestHandler override the default function that is used to service requests
func (ss *SignerServer) SetRequestHandler(validationRequestHandler ValidationRequestHandlerFunc) {
	_ = "STUB: not implemented"
	return
}

func (ss *SignerServer) servicePendingRequest() { _ = "STUB: not implemented"; return }

// Ignore error from closing.

// limit the scope of the lock

// only log the error; we'll reply with an error in res

func (ss *SignerServer) serviceLoop() { _ = "STUB: not implemented"; return }
