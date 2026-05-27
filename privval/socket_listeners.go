package privval

import (
	"net"
	"time"

	"github.com/cometbft/cometbft/crypto/ed25519"
)

const (
	defaultTimeoutAcceptSeconds = 3
)

// timeoutError can be used to check if an error returned from the netp package
// was due to a timeout.
type timeoutError interface {
	Timeout() bool
}

//------------------------------------------------------------------
// TCP Listener

// TCPListenerOption sets an optional parameter on the tcpListener.
type TCPListenerOption func(*TCPListener)

// TCPListenerTimeoutAccept sets the timeout for the listener.
// A zero time value disables the timeout.
func TCPListenerTimeoutAccept(timeout time.Duration) TCPListenerOption {
	_ = "STUB: not implemented"
	return *new(TCPListenerOption)
}

// TCPListenerTimeoutReadWrite sets the read and write timeout for connections
// from external signing processes.
func TCPListenerTimeoutReadWrite(timeout time.Duration) TCPListenerOption {
	_ = "STUB: not implemented"
	return *new(TCPListenerOption)
}

// tcpListener implements net.Listener.
var _ net.Listener = (*TCPListener)(nil)

// TCPListener wraps a *net.TCPListener to standardize protocol timeouts
// and potentially other tuning parameters. It also returns encrypted connections.
type TCPListener struct {
	*net.TCPListener

	secretConnKey ed25519.PrivKey

	timeoutAccept    time.Duration
	timeoutReadWrite time.Duration
}

// NewTCPListener returns a listener that accepts authenticated encrypted connections
// using the given secretConnKey and the default timeout values.
func NewTCPListener(ln net.Listener, secretConnKey ed25519.PrivKey) *TCPListener {
	_ = "STUB: not implemented"
	return nil
}

// Accept implements net.Listener.
func (ln *TCPListener) Accept() (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

// Wrap the conn in our timeout and encryption wrappers

//------------------------------------------------------------------
// Unix Listener

// unixListener implements net.Listener.
var _ net.Listener = (*UnixListener)(nil)

type UnixListenerOption func(*UnixListener)

// UnixListenerTimeoutAccept sets the timeout for the listener.
// A zero time value disables the timeout.
func UnixListenerTimeoutAccept(timeout time.Duration) UnixListenerOption {
	_ = "STUB: not implemented"
	return *new(UnixListenerOption)
}

// UnixListenerTimeoutReadWrite sets the read and write timeout for connections
// from external signing processes.
func UnixListenerTimeoutReadWrite(timeout time.Duration) UnixListenerOption {
	_ = "STUB: not implemented"
	return *new(UnixListenerOption)
}

// UnixListener wraps a *net.UnixListener to standardize protocol timeouts
// and potentially other tuning parameters. It returns unencrypted connections.
type UnixListener struct {
	*net.UnixListener

	timeoutAccept    time.Duration
	timeoutReadWrite time.Duration
}

// NewUnixListener returns a listener that accepts unencrypted connections
// using the default timeout values.
func NewUnixListener(ln net.Listener) *UnixListener { _ = "STUB: not implemented"; return nil }

// Accept implements net.Listener.
func (ln *UnixListener) Accept() (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

// Wrap the conn in our timeout wrapper

// TODO: wrap in something that authenticates
// with a MAC - https://github.com/tendermint/tendermint/issues/3099

//------------------------------------------------------------------
// Connection

// timeoutConn implements net.Conn.
var _ net.Conn = (*timeoutConn)(nil)

// timeoutConn wraps a net.Conn to standardize protocol timeouts / deadline resets.
type timeoutConn struct {
	net.Conn
	timeout time.Duration
}

// newTimeoutConn returns an instance of timeoutConn.
func newTimeoutConn(conn net.Conn, timeout time.Duration) *timeoutConn {
	_ = "STUB: not implemented"
	return nil
}

// Read implements net.Conn.
func (c timeoutConn) Read(b []byte) (n int, err error) {
	_ = "STUB: not implemented"
	// Reset deadline
	return 0, nil
}

//nolint:staticcheck

// Write implements net.Conn.
func (c timeoutConn) Write(b []byte) (n int, err error) {
	_ = "STUB: not implemented"
	// Reset deadline
	return 0, nil
}

//nolint:staticcheck
