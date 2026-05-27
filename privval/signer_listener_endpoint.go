package privval

import (
	"net"
	"sync/atomic"
	"time"

	"github.com/cometbft/cometbft/libs/log"
	cmtsync "github.com/cometbft/cometbft/libs/sync"
	privvalproto "github.com/cometbft/cometbft/proto/tendermint/privval"
)

// SignerListenerEndpointOption sets an optional parameter on the SignerListenerEndpoint.
type SignerListenerEndpointOption func(*SignerListenerEndpoint)

// SignerListenerEndpointTimeoutReadWrite sets the read and write timeout for
// connections from external signing processes.
//
// Default: 5s
func SignerListenerEndpointTimeoutReadWrite(timeout time.Duration) SignerListenerEndpointOption {
	_ = "STUB: not implemented"
	return *new(SignerListenerEndpointOption)
}

//nolint:staticcheck

// SignerListenerEndpoint listens for an external process to dial in and keeps
// the connection alive by dropping and reconnecting.
//
// The process will send pings every ~3s (read/write timeout * 2/3) to keep the
// connection alive.
type SignerListenerEndpoint struct {
	signerEndpoint

	listener              net.Listener
	connectRequestCh      chan struct{}
	connectionAvailableCh chan net.Conn

	timeoutAccept   time.Duration
	acceptFailCount atomic.Uint32
	pingTimer       *time.Ticker
	pingInterval    time.Duration

	instanceMtx cmtsync.Mutex // Ensures instance public methods access, i.e. SendRequest
}

// NewSignerListenerEndpoint returns an instance of SignerListenerEndpoint.
func NewSignerListenerEndpoint(
	logger log.Logger,
	listener net.Listener,
	options ...SignerListenerEndpointOption,
) *SignerListenerEndpoint {
	_ = "STUB: not implemented"
	return nil
}

//nolint:staticcheck

// OnStart implements service.Service.
func (sl *SignerListenerEndpoint) OnStart() error { _ = "STUB: not implemented"; return nil }

// Buffer of 1 to allow `serviceLoop` to re-trigger itself.

// NOTE: ping timeout must be less than read/write timeout.
//nolint:staticcheck

// OnStop implements service.Service
func (sl *SignerListenerEndpoint) OnStop() { _ = "STUB: not implemented"; return }

// Stop listening

// WaitForConnection waits maxWait for a connection or returns a timeout error
func (sl *SignerListenerEndpoint) WaitForConnection(maxWait time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// SendRequest ensures there is a connection, sends a request and waits for a response
func (sl *SignerListenerEndpoint) SendRequest(request privvalproto.Message) (*privvalproto.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Reset pingTimer to avoid sending unnecessary pings.

func (sl *SignerListenerEndpoint) ensureConnection(maxWait time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// Is there a connection ready? then use it

// block until connected or timeout

func (sl *SignerListenerEndpoint) acceptNewConnection() (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

// wait for a new conn

func (sl *SignerListenerEndpoint) triggerConnect() { _ = "STUB: not implemented"; return }

func (sl *SignerListenerEndpoint) triggerReconnect() { _ = "STUB: not implemented"; return }

func (sl *SignerListenerEndpoint) serviceLoop() { _ = "STUB: not implemented"; return }

// On start, listen timeouts can queue a duplicate connect request to queue
// while the first request connects.  Drop duplicate request.

// Listen for remote signer

// We have a good connection, wait for someone that needs one otherwise cancellation

func (sl *SignerListenerEndpoint) pingLoop() { _ = "STUB: not implemented"; return }
