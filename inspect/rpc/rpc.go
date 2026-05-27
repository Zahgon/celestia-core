package rpc

import (
	"context"
	"net/http"

	"github.com/cometbft/cometbft/config"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cometbft/cometbft/rpc/core"
	"github.com/cometbft/cometbft/rpc/jsonrpc/server"
	"github.com/cometbft/cometbft/state"
	"github.com/cometbft/cometbft/state/indexer"
	"github.com/cometbft/cometbft/state/txindex"
)

// Server defines parameters for running an Inspector rpc server.
type Server struct {
	Addr    string // TCP address to listen on, ":http" if empty
	Handler http.Handler
	Logger  log.Logger
	Config  *config.RPCConfig
}

// Routes returns the set of routes used by the Inspector server.
func Routes(cfg config.RPCConfig, s state.Store, bs state.BlockStore, txidx txindex.TxIndexer, blkidx indexer.BlockIndexer, logger log.Logger) core.RoutesMap {
	_ = "STUB: not implemented" //nolint: lll
	return *new(core.RoutesMap)
}

//nolint:staticcheck
//nolint:staticcheck
//nolint:staticcheck

// Handler returns the http.Handler configured for use with an Inspector server. Handler
// registers the routes on the http.Handler and also registers the websocket handler
// and the CORS handler if specified by the configuration options.
func Handler(rpcConfig *config.RPCConfig, routes core.RoutesMap, logger log.Logger) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

func addCORSHandler(rpcConfig *config.RPCConfig, h http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

type waitSyncCheckerImpl struct{}

func (waitSyncCheckerImpl) WaitSync() bool {
	_ = "STUB: not implemented"

	// ListenAndServe listens on the address specified in srv.Addr and handles any
	// incoming requests over HTTP using the Inspector rpc handler specified on the server.
	return false
}

func (srv *Server) ListenAndServe(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// ListenAndServeTLS listens on the address specified in srv.Addr. ListenAndServeTLS handles
// incoming requests over HTTPS using the Inspector rpc handler specified on the server.
func (srv *Server) ListenAndServeTLS(ctx context.Context, certFile, keyFile string) error {
	_ = "STUB: not implemented"
	return nil
}

func serverRPCConfig(r *config.RPCConfig) *server.Config { _ = "STUB: not implemented"; return nil }

// If necessary adjust global WriteTimeout to ensure it's greater than
// TimeoutBroadcastTxCommit.
// See https://github.com/tendermint/tendermint/issues/3435
