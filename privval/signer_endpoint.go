package privval

import (
	"net"
	"time"

	"github.com/cometbft/cometbft/libs/service"
	cmtsync "github.com/cometbft/cometbft/libs/sync"
	privvalproto "github.com/cometbft/cometbft/proto/tendermint/privval"
)

const (
	defaultTimeoutReadWriteSeconds = 5
)

type signerEndpoint struct {
	service.BaseService

	connMtx cmtsync.Mutex
	conn    net.Conn

	timeoutReadWrite time.Duration
}

// Close closes the underlying net.Conn.
func (se *signerEndpoint) Close() error { _ = "STUB: not implemented"; return nil }

// IsConnected indicates if there is an active connection
func (se *signerEndpoint) IsConnected() bool { _ = "STUB: not implemented"; return false }

// TryGetConnection retrieves a connection if it is already available
func (se *signerEndpoint) GetAvailableConnection(connectionAvailableCh chan net.Conn) bool {
	_ = "STUB: not implemented"
	return false
}

// Is there a connection ready?

// TryGetConnection retrieves a connection if it is already available
func (se *signerEndpoint) WaitConnection(connectionAvailableCh chan net.Conn, maxWait time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// SetConnection replaces the current connection object
func (se *signerEndpoint) SetConnection(newConnection net.Conn) { _ = "STUB: not implemented"; return }

// IsConnected indicates if there is an active connection
func (se *signerEndpoint) DropConnection() { _ = "STUB: not implemented"; return }

// ReadMessage reads a message from the endpoint
func (se *signerEndpoint) ReadMessage() (msg privvalproto.Message, err error) {
	_ = "STUB: not implemented"
	return *new(privvalproto.Message), nil
}

// Reset read deadline

// WriteMessage writes a message from the endpoint
func (se *signerEndpoint) WriteMessage(msg privvalproto.Message) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Reset read deadline

func (se *signerEndpoint) isConnected() bool { _ = "STUB: not implemented"; return false }

func (se *signerEndpoint) dropConnection() { _ = "STUB: not implemented"; return }
