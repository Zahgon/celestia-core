package p2p

import (
	"net"
	"time"

	"github.com/cometbft/cometbft/config"
	cmtsync "github.com/cometbft/cometbft/libs/sync"
)

// FuzzedConnection wraps any net.Conn and depending on the mode either delays
// reads/writes or randomly drops reads/writes/connections.
type FuzzedConnection struct {
	conn net.Conn

	mtx    cmtsync.Mutex
	start  <-chan time.Time
	active bool

	config *config.FuzzConnConfig
}

// FuzzConn creates a new FuzzedConnection. Fuzzing starts immediately.
func FuzzConn(conn net.Conn) net.Conn { _ = "STUB: not implemented"; return *new(net.Conn) }

// FuzzConnFromConfig creates a new FuzzedConnection from a config. Fuzzing
// starts immediately.
func FuzzConnFromConfig(conn net.Conn, config *config.FuzzConnConfig) net.Conn {
	_ = "STUB: not implemented"
	return *new(net.Conn)
}

// FuzzConnAfter creates a new FuzzedConnection. Fuzzing starts when the
// duration elapses.
func FuzzConnAfter(conn net.Conn, d time.Duration) net.Conn {
	_ = "STUB: not implemented"
	return *new(net.Conn)
}

// FuzzConnAfterFromConfig creates a new FuzzedConnection from a config.
// Fuzzing starts when the duration elapses.
func FuzzConnAfterFromConfig(
	conn net.Conn,
	d time.Duration,
	config *config.FuzzConnConfig,
) net.Conn {
	_ = "STUB: not implemented"
	return *new(net.Conn)
}

// Config returns the connection's config.
func (fc *FuzzedConnection) Config() *config.FuzzConnConfig {
	_ = "STUB: not implemented"

	// Read implements net.Conn.
	return nil
}

func (fc *FuzzedConnection) Read(data []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Write implements net.Conn.
func (fc *FuzzedConnection) Write(data []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Close implements net.Conn.
func (fc *FuzzedConnection) Close() error { _ = "STUB: not implemented"; return nil }

// LocalAddr implements net.Conn.
func (fc *FuzzedConnection) LocalAddr() net.Addr {
	_ = "STUB: not implemented"
	return *

	// RemoteAddr implements net.Conn.
	new(net.Addr)
}

func (fc *FuzzedConnection) RemoteAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

// SetDeadline implements net.Conn.
func (fc *FuzzedConnection) SetDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

// SetReadDeadline implements net.Conn.
func (fc *FuzzedConnection) SetReadDeadline(t time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

// SetWriteDeadline implements net.Conn.
func (fc *FuzzedConnection) SetWriteDeadline(t time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (fc *FuzzedConnection) randomDuration() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

//nolint: gas

// implements the fuzz (delay, kill conn)
// and returns whether or not the read/write should be ignored
func (fc *FuzzedConnection) fuzz() bool { _ = "STUB: not implemented"; return false }

// randomly drop the r/w, drop the conn, or sleep

// XXX: can't this fail because machine precision?
// XXX: do we need an error?

// sleep a bit

func (fc *FuzzedConnection) shouldFuzz() bool { _ = "STUB: not implemented"; return false }
