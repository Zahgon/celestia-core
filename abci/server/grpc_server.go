package server

import (
	"context"
	"net"

	"google.golang.org/grpc"

	"github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/libs/service"
)

const mebibyte = 1024 * 1024

type GRPCServer struct {
	service.BaseService

	proto    string
	addr     string
	listener net.Listener
	server   *grpc.Server

	app types.Application
}

// NewGRPCServer returns a new gRPC ABCI server
func NewGRPCServer(protoAddr string, app types.Application) service.Service {
	_ = "STUB: not implemented"
	return *new(service.Service)
}

// OnStart starts the gRPC service.
func (s *GRPCServer) OnStart() error { _ = "STUB: not implemented"; return nil }

// OnStop stops the gRPC server.
func (s *GRPCServer) OnStop() {
	_ = "STUB: not implemented"

	// -------------------------------------------------------
	return
}

// gRPCApplication is a gRPC shim for Application
type gRPCApplication struct {
	types.Application
}

func (app *gRPCApplication) Echo(_ context.Context, req *types.RequestEcho) (*types.ResponseEcho, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (app *gRPCApplication) Flush(context.Context, *types.RequestFlush) (*types.ResponseFlush, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (app *gRPCApplication) QuerySequence(ctx context.Context, req *types.RequestQuerySequence) (*types.ResponseQuerySequence, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
