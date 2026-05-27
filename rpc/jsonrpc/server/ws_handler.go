package server

import (
	"context"
	"net/http"
	"time"

	"github.com/gorilla/websocket"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cometbft/cometbft/libs/service"
	types "github.com/cometbft/cometbft/rpc/jsonrpc/types"
)

// WebSocket handler

const (
	defaultWSWriteChanCapacity = 100
	defaultWSWriteWait         = 10 * time.Second
	defaultWSReadWait          = 30 * time.Second
	defaultWSPingPeriod        = (defaultWSReadWait * 9) / 10
)

// WebsocketManager provides a WS handler for incoming connections and passes a
// map of functions along with any additional params to new connections.
// NOTE: The websocket path is defined externally, e.g. in node/node.go
type WebsocketManager struct {
	websocket.Upgrader

	funcMap       map[string]*RPCFunc
	logger        log.Logger
	wsConnOptions []func(*wsConnection)
}

// NewWebsocketManager returns a new WebsocketManager that passes a map of
// functions, connection options and logger to new WS connections.
func NewWebsocketManager(
	funcMap map[string]*RPCFunc,
	wsConnOptions ...func(*wsConnection),
) *WebsocketManager {
	_ = "STUB: not implemented"
	return nil
}

// TODO ???
//
// The default behavior would be relevant to browser-based clients,
// afaik. I suppose having a pass-through is a workaround for allowing
// for more complex security schemes, shifting the burden of
// AuthN/AuthZ outside the CometBFT RPC.
// I can't think of other uses right now that would warrant a TODO
// though. The real backstory of this TODO shall remain shrouded in
// mystery

// SetLogger sets the logger.
func (wm *WebsocketManager) SetLogger(l log.Logger) {
	_ = "STUB: not implemented"

	// WebsocketHandler upgrades the request/response (via http.Hijack) and starts
	// the wsConnection.
	return
}

func (wm *WebsocketManager) WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// TODO - return http error

// register connection

// BLOCKING

// WebSocket connection

// A single websocket connection contains listener id, underlying ws
// connection, and the event switch for subscribing to events.
//
// In case of an error, the connection is stopped.
type wsConnection struct {
	service.BaseService

	remoteAddr string
	baseConn   *websocket.Conn
	// writeChan is never closed, to allow WriteRPCResponse() to fail.
	writeChan chan types.RPCResponse

	// chan, which is closed when/if readRoutine errors
	// used to abort writeRoutine
	readRoutineQuit chan struct{}

	funcMap map[string]*RPCFunc

	// write channel capacity
	writeChanCapacity int

	// each write times out after this.
	writeWait time.Duration

	// Connection times out if we haven't received *anything* in this long, not even pings.
	readWait time.Duration

	// Send pings to server with this period. Must be less than readWait, but greater than zero.
	pingPeriod time.Duration

	// Maximum message size.
	readLimit int64

	// callback which is called upon disconnect
	onDisconnect func(remoteAddr string)

	ctx    context.Context
	cancel context.CancelFunc
}

// NewWSConnection wraps websocket.Conn.
//
// See the commentary on the func(*wsConnection) functions for a detailed
// description of how to configure ping period and pong wait time. NOTE: if the
// write buffer is full, pongs may be dropped, which may cause clients to
// disconnect. see https://github.com/gorilla/websocket/issues/97
func newWSConnection(
	baseConn *websocket.Conn,
	funcMap map[string]*RPCFunc,
	options ...func(*wsConnection),
) *wsConnection {
	_ = "STUB: not implemented"
	return nil
}

// OnDisconnect sets a callback which is used upon disconnect - not
// Goroutine-safe. Nop by default.
func OnDisconnect(onDisconnect func(remoteAddr string)) func(*wsConnection) {
	_ = "STUB: not implemented"
	return nil
}

// WriteWait sets the amount of time to wait before a websocket write times out.
// It should only be used in the constructor - not Goroutine-safe.
func WriteWait(writeWait time.Duration) func(*wsConnection) { _ = "STUB: not implemented"; return nil }

// WriteChanCapacity sets the capacity of the websocket write channel.
// It should only be used in the constructor - not Goroutine-safe.
func WriteChanCapacity(cap int) func(*wsConnection) { _ = "STUB: not implemented"; return nil }

// ReadWait sets the amount of time to wait before a websocket read times out.
// It should only be used in the constructor - not Goroutine-safe.
func ReadWait(readWait time.Duration) func(*wsConnection) { _ = "STUB: not implemented"; return nil }

// PingPeriod sets the duration for sending websocket pings.
// It should only be used in the constructor - not Goroutine-safe.
func PingPeriod(pingPeriod time.Duration) func(*wsConnection) {
	_ = "STUB: not implemented"
	return nil
}

// ReadLimit sets the maximum size for reading message.
// It should only be used in the constructor - not Goroutine-safe.
func ReadLimit(readLimit int64) func(*wsConnection) { _ = "STUB: not implemented"; return nil }

// OnStart implements service.Service by starting the read and write routines. It
// blocks until there's some error.
func (wsc *wsConnection) OnStart() error { _ = "STUB: not implemented"; return nil }

// Read subscriptions/unsubscriptions to events

// Write responses, BLOCKING.

// OnStop implements service.Service by unsubscribing remoteAddr from all
// subscriptions.
func (wsc *wsConnection) OnStop() { _ = "STUB: not implemented"; return }

// GetRemoteAddr returns the remote address of the underlying connection.
// It implements WSRPCConnection
func (wsc *wsConnection) GetRemoteAddr() string { _ = "STUB: not implemented"; return "" }

// WriteRPCResponse pushes a response to the writeChan, and blocks until it is
// accepted.
// It implements WSRPCConnection. It is Goroutine-safe.
func (wsc *wsConnection) WriteRPCResponse(ctx context.Context, resp types.RPCResponse) error {
	_ = "STUB: not implemented"
	return nil
}

// TryWriteRPCResponse attempts to push a response to the writeChan, but does
// not block.
// It implements WSRPCConnection. It is Goroutine-safe
func (wsc *wsConnection) TryWriteRPCResponse(resp types.RPCResponse) bool {
	_ = "STUB: not implemented"
	return false
}

// Context returns the connection's context.
// The context is canceled when the client's connection closes.
func (wsc *wsConnection) Context() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// Read from the socket and subscribe to or unsubscribe from events
func (wsc *wsConnection) readRoutine() {
	_ = "STUB: not implemented"
	// readRoutine will block until response is written or WS connection is closed
	return
}

// reset deadline for every type of message (control or data)

// A Notification is a Request object without an "id" member.
// The Server MUST NOT reply to a Notification, including those that are within a batch request.

// Now, fetch the RPCFunc and execute it.

// TODO: Need to encode args/returns to string if we want to log them

// receives on a write channel and writes out on the socket
func (wsc *wsConnection) writeRoutine() { _ = "STUB: not implemented"; return }

// https://github.com/gorilla/websocket/issues/97

// error in readRoutine

// Use json.MarshalIndent instead of Marshal for pretty output.
// Pretty output not necessary, since most consumers of WS events are
// automated processes, not humans.

// All writes to the websocket must (re)set the write deadline.
// If some writes don't set it while others do, they may timeout incorrectly
// (https://github.com/tendermint/tendermint/issues/553)
func (wsc *wsConnection) writeMessageWithDeadline(msgType int, msg []byte) error {
	_ = "STUB: not implemented"
	return nil
}
