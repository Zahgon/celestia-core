package proxy

import (
	abcicli "github.com/cometbft/cometbft/abci/client"
	"github.com/cometbft/cometbft/libs/service"
)

const (
	connConsensus = "consensus"
	connMempool   = "mempool"
	connQuery     = "query"
	connSnapshot  = "snapshot"
)

// AppConns is the CometBFT's interface to the application that consists of
// multiple connections.
type AppConns interface {
	service.Service

	// Mempool connection
	Mempool() AppConnMempool
	// Consensus connection
	Consensus() AppConnConsensus
	// Query connection
	Query() AppConnQuery
	// Snapshot connection
	Snapshot() AppConnSnapshot
}

// NewAppConns calls NewMultiAppConn.
func NewAppConns(clientCreator ClientCreator, metrics *Metrics) AppConns {
	_ = "STUB: not implemented"
	return *new(AppConns)
}

// multiAppConn implements AppConns.
//
// A multiAppConn is made of a few appConns and manages their underlying abci
// clients.
// TODO: on app restart, clients must reboot together
type multiAppConn struct {
	service.BaseService

	metrics       *Metrics
	consensusConn AppConnConsensus
	mempoolConn   AppConnMempool
	queryConn     AppConnQuery
	snapshotConn  AppConnSnapshot

	consensusConnClient abcicli.Client
	mempoolConnClient   abcicli.Client
	queryConnClient     abcicli.Client
	snapshotConnClient  abcicli.Client

	clientCreator ClientCreator
}

// NewMultiAppConn makes all necessary abci connections to the application.
func NewMultiAppConn(clientCreator ClientCreator, metrics *Metrics) AppConns {
	_ = "STUB: not implemented"
	return *new(AppConns)
}

func (app *multiAppConn) Mempool() AppConnMempool {
	_ = "STUB: not implemented"
	return *new(AppConnMempool)
}

func (app *multiAppConn) Consensus() AppConnConsensus {
	_ = "STUB: not implemented"
	return *new(AppConnConsensus)
}

func (app *multiAppConn) Query() AppConnQuery { _ = "STUB: not implemented"; return *new(AppConnQuery) }

func (app *multiAppConn) Snapshot() AppConnSnapshot {
	_ = "STUB: not implemented"
	return *new(AppConnSnapshot)
}

func (app *multiAppConn) OnStart() error { _ = "STUB: not implemented"; return nil }

// Kill CometBFT if the ABCI application crashes.

func (app *multiAppConn) OnStop() { _ = "STUB: not implemented"; return }

func (app *multiAppConn) killTMOnClientError() { _ = "STUB: not implemented"; return }

func (app *multiAppConn) stopAllClients() { _ = "STUB: not implemented"; return }

func (app *multiAppConn) abciClientFor(conn string) (abcicli.Client, error) {
	_ = "STUB: not implemented"
	return *new(abcicli.Client), nil
}
