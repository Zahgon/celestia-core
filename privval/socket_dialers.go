package privval

import (
	"errors"
	"net"
	"time"

	"github.com/cometbft/cometbft/crypto"
)

// Socket errors.
var (
	ErrDialRetryMax = errors.New("dialed maximum retries")
)

// SocketDialer dials a remote address and returns a net.Conn or an error.
type SocketDialer func() (net.Conn, error)

// DialTCPFn dials the given tcp addr, using the given timeoutReadWrite and
// privKey for the authenticated encryption handshake.
func DialTCPFn(addr string, timeoutReadWrite time.Duration, privKey crypto.PrivKey) SocketDialer {
	_ = "STUB: not implemented"
	return *new(SocketDialer)
}

// DialUnixFn dials the given unix socket.
func DialUnixFn(addr string) SocketDialer { _ = "STUB: not implemented"; return *new(SocketDialer) }
