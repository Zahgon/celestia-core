package client

import (
	"context"
	"net"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	metrics "github.com/rcrowley/go-metrics"

	"github.com/cometbft/cometbft/libs/service"
	cmtsync "github.com/cometbft/cometbft/libs/sync"
	types "github.com/cometbft/cometbft/rpc/jsonrpc/types"
)

const (
	defaultMaxReconnectAttempts = 25
	defaultWriteWait            = 0
	defaultReadWait             = 0
	defaultPingPeriod           = 0
)

// WSClient is a JSON-RPC client, which uses WebSocket for communication with
// the remote server.
//
// WSClient is safe for concurrent use by multiple goroutines.
type WSClient struct {
	conn *websocket.Conn

	Address  string // IP:PORT or /path/to/socket
	Endpoint string // /websocket/url/endpoint
	Username string
	Password string

	Dialer func(string, string) (net.Conn, error)

	// Single user facing channel to read RPCResponses from, closed only when the
	// client is being stopped.
	ResponsesCh chan types.RPCResponse

	// Callback, which will be called each time after successful reconnect.
	onReconnect func()

	// internal channels
	send            chan types.RPCRequest // user requests
	backlog         chan types.RPCRequest // stores a single user request received during a conn failure
	reconnectAfter  chan error            // reconnect requests
	readRoutineQuit chan struct{}         // a way for readRoutine to close writeRoutine

	// Maximum reconnect attempts (0 or greater; default: 25).
	maxReconnectAttempts int

	// Support both ws and wss protocols
	protocol string

	wg sync.WaitGroup

	mtx            cmtsync.RWMutex
	sentLastPingAt time.Time
	reconnecting   bool
	nextReqID      int
	// sentIDs        map[types.JSONRPCIntID]bool // IDs of the requests currently in flight

	// Time allowed to write a message to the server. 0 means block until operation succeeds.
	writeWait time.Duration

	// Time allowed to read the next message from the server. 0 means block until operation succeeds.
	readWait time.Duration

	// Send pings to server with this period. Must be less than readWait. If 0, no pings will be sent.
	pingPeriod time.Duration

	service.BaseService

	// Time between sending a ping and receiving a pong. See
	// https://godoc.org/github.com/rcrowley/go-metrics#Timer.
	PingPongLatencyTimer metrics.Timer
}

// NewWS returns a new client. See the commentary on the func(*WSClient)
// functions for a detailed description of how to configure ping period and
// pong wait time. The endpoint argument must begin with a `/`.
// An error is returned on invalid remote. The function panics when remote is nil.
func NewWS(remoteAddr, endpoint string, options ...func(*WSClient)) (*WSClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// default to ws protocol, unless wss or https is specified

// extract username and password from URL if any

// sentIDs: make(map[types.JSONRPCIntID]bool),

// MaxReconnectAttempts sets the maximum number of reconnect attempts before returning an error.
// It should only be used in the constructor and is not Goroutine-safe.
func MaxReconnectAttempts(max int) func(*WSClient) { _ = "STUB: not implemented"; return nil }

// ReadWait sets the amount of time to wait before a websocket read times out.
// It should only be used in the constructor and is not Goroutine-safe.
func ReadWait(readWait time.Duration) func(*WSClient) { _ = "STUB: not implemented"; return nil }

// WriteWait sets the amount of time to wait before a websocket write times out.
// It should only be used in the constructor and is not Goroutine-safe.
func WriteWait(writeWait time.Duration) func(*WSClient) { _ = "STUB: not implemented"; return nil }

// PingPeriod sets the duration for sending websocket pings.
// It should only be used in the constructor - not Goroutine-safe.
func PingPeriod(pingPeriod time.Duration) func(*WSClient) { _ = "STUB: not implemented"; return nil }

// OnReconnect sets the callback, which will be called every time after
// successful reconnect.
func OnReconnect(cb func()) func(*WSClient) { _ = "STUB: not implemented"; return nil }

// String returns WS client full address.
func (c *WSClient) String() string { _ = "STUB: not implemented"; return "" }

// OnStart implements service.Service by dialing a server and creating read and
// write routines.
func (c *WSClient) OnStart() error { _ = "STUB: not implemented"; return nil }

// 1 additional error may come from the read/write
// goroutine depending on which failed first.

// capacity for 1 request. a user won't be able to send more because the send
// channel is unbuffered.

// Stop overrides service.Service#Stop. There is no other way to wait until Quit
// channel is closed.
func (c *WSClient) Stop() error { _ = "STUB: not implemented"; return nil }

// only close user-facing channels when we can't write to them

// IsReconnecting returns true if the client is reconnecting right now.
func (c *WSClient) IsReconnecting() bool { _ = "STUB: not implemented"; return false }

// IsActive returns true if the client is running and not reconnecting.
func (c *WSClient) IsActive() bool { _ = "STUB: not implemented"; return false }

// Send the given RPC request to the server. Results will be available on
// ResponsesCh, errors, if any, on ErrorsCh. Will block until send succeeds or
// ctx.Done is closed.
func (c *WSClient) Send(ctx context.Context, request types.RPCRequest) error {
	_ = "STUB: not implemented"
	return nil
}

// c.mtx.Lock()
// c.sentIDs[request.ID.(types.JSONRPCIntID)] = true
// c.mtx.Unlock()

// Call enqueues a call request onto the Send queue. Requests are JSON encoded.
func (c *WSClient) Call(ctx context.Context, method string, params map[string]interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// CallWithArrayParams enqueues a call request onto the Send queue. Params are
// in a form of array (e.g. []interface{}{"abcd"}). Requests are JSON encoded.
func (c *WSClient) CallWithArrayParams(ctx context.Context, method string, params []interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// Private methods

func (c *WSClient) nextRequestID() types.JSONRPCIntID {
	_ = "STUB: not implemented"
	return *new(types.JSONRPCIntID)
}

func (c *WSClient) dial() error { _ = "STUB: not implemented"; return nil }

// Set basic auth header if username and password are provided

//nolint:bodyclose

// reconnect tries to redial up to maxReconnectAttempts with exponential
// backoff.
func (c *WSClient) reconnect() error { _ = "STUB: not implemented"; return nil }

// 1s == (1e9 ns)

func (c *WSClient) startReadWriteRoutines() { _ = "STUB: not implemented"; return }

func (c *WSClient) processBacklog() error { _ = "STUB: not implemented"; return nil }

// requeue request

func (c *WSClient) reconnectRoutine() { _ = "STUB: not implemented"; return }

// wait until writeRoutine and readRoutine finish

// drain reconnectAfter

// The client ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *WSClient) writeRoutine() { _ = "STUB: not implemented"; return }

// ticker with a predefined period

// ticker that never fires

// err != nil {
// ignore error; it will trigger in tests
// likely because it's closing an already closed connection
// }

// add request to the backlog, so we don't lose it

// The client ensures that there is at most one reader to a connection by
// executing all reads from this goroutine.
func (c *WSClient) readRoutine() { _ = "STUB: not implemented"; return }

// err != nil {
// ignore error; it will trigger in tests
// likely because it's closing an already closed connection
// }

// gather latency stats

// reset deadline for every message type (control or data)

// TODO: events resulting from /subscribe do not work with ->
// because they are implemented as responses with the subscribe request's
// ID. According to the spec, they should be notifications (requests
// without IDs).
// https://github.com/tendermint/tendermint/issues/2949
// c.mtx.Lock()
// if _, ok := c.sentIDs[response.ID.(types.JSONRPCIntID)]; !ok {
// 	c.Logger.Error("unsolicited response ID", "id", response.ID, "expected", c.sentIDs)
// 	c.mtx.Unlock()
// 	continue
// }
// delete(c.sentIDs, response.ID.(types.JSONRPCIntID))
// c.mtx.Unlock()
// Combine a non-blocking read on BaseService.Quit with a non-blocking write on ResponsesCh to avoid blocking
// c.wg.Wait() in c.Stop(). Note we rely on Quit being closed so that it sends unlimited Quit signals to stop
// both readRoutine and writeRoutine

// Predefined methods

// Subscribe to a query. Note the server must have a "subscribe" route
// defined.
func (c *WSClient) Subscribe(ctx context.Context, query string) error {
	_ = "STUB: not implemented"
	return nil
}

// Unsubscribe from a query. Note the server must have a "unsubscribe" route
// defined.
func (c *WSClient) Unsubscribe(ctx context.Context, query string) error {
	_ = "STUB: not implemented"
	return nil
}

// UnsubscribeAll from all. Note the server must have a "unsubscribe_all" route
// defined.
func (c *WSClient) UnsubscribeAll(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
