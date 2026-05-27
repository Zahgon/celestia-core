package server

import (
	"context"
	"io"
	"net"

	"github.com/cometbft/cometbft/abci/types"
	cmtlog "github.com/cometbft/cometbft/libs/log"
	"github.com/cometbft/cometbft/libs/service"
	cmtsync "github.com/cometbft/cometbft/libs/sync"
)

// SocketServer is the server-side implementation of the TSP (Tendermint Socket Protocol)
// for out-of-process go applications. Note, in the case of an application written in golang,
// the developer may also run both Tendermint and the application within the same process.
//
// The socket server deliver
type SocketServer struct {
	service.BaseService
	isLoggerSet bool

	proto    string
	addr     string
	listener net.Listener

	connsMtx   cmtsync.Mutex
	conns      map[int]net.Conn
	nextConnID int

	appMtx cmtsync.Mutex
	app    types.Application
}

const responseBufferSize = 1000

// NewSocketServer creates a server from a golang-based out-of-process application.
func NewSocketServer(protoAddr string, app types.Application) service.Service {
	_ = "STUB: not implemented"
	return *new(service.Service)
}

func (s *SocketServer) SetLogger(l cmtlog.Logger) { _ = "STUB: not implemented"; return }

func (s *SocketServer) OnStart() error { _ = "STUB: not implemented"; return nil }

func (s *SocketServer) OnStop() { _ = "STUB: not implemented"; return }

func (s *SocketServer) addConn(conn net.Conn) int { _ = "STUB: not implemented"; return 0 }

// deletes conn even if close errs
func (s *SocketServer) rmConn(connID int) error { _ = "STUB: not implemented"; return nil }

func (s *SocketServer) acceptConnectionsRoutine() {
	_ = "STUB: not implemented"

	// Accept a connection
	return
}

// Ignore error from listener closing.

// Push to signal connection closed
// A channel to buffer responses

// Read requests from conn and deal with them

// Pull responses from 'responses' and write them to conn.

// Wait until signal to close connection

func (s *SocketServer) waitForClose(closeConn chan error, connID int) {
	_ = "STUB: not implemented"
	return
}

// never happens

// Close the connection

// Read requests from conn and deal with them
func (s *SocketServer) handleRequests(closeConn chan error, conn io.Reader, responses chan<- *types.Response) {
	_ = "STUB: not implemented"
	return
}

// true only while appMtx is held inside the loop

// make sure to recover from any app-related panics to allow proper socket cleanup.
// In the case of a panic, we do not notify the client by passing an exception so
// presume that the client is still running and retying to connect

// any error either from the application or because of an unknown request
// throws an exception back to the client. This will stop the server and
// should also halt the client.

// handleRequests takes a request and calls the application passing the returned
func (s *SocketServer) handleRequest(ctx context.Context, req *types.Request) (*types.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Pull responses from 'responses' and write them to conn.
func (s *SocketServer) handleResponses(closeConn chan error, conn io.Writer, responses <-chan *types.Response) {
	_ = "STUB: not implemented"
	return
}

// If the application has responded with an exception, the server returns the error
// back to the client and closes the connection. The receiving Tendermint client should
// log the error and gracefully terminate
