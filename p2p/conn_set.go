package p2p

import (
	"net"

	cmtsync "github.com/cometbft/cometbft/libs/sync"
)

// ConnSet is a lookup table for connections and all their ips.
type ConnSet interface {
	Has(net.Conn) bool
	HasIP(net.IP) bool
	Set(net.Conn, []net.IP)
	Remove(net.Conn)
	RemoveAddr(net.Addr)
}

type connSetItem struct {
	conn net.Conn
	ips  []net.IP
}

type connSet struct {
	cmtsync.RWMutex

	conns map[string]connSetItem
}

// NewConnSet returns a ConnSet implementation.
func NewConnSet() ConnSet { _ = "STUB: not implemented"; return *new(ConnSet) }

func (cs *connSet) Has(c net.Conn) bool { _ = "STUB: not implemented"; return false }

func (cs *connSet) HasIP(ip net.IP) bool { _ = "STUB: not implemented"; return false }

func (cs *connSet) Remove(c net.Conn) { _ = "STUB: not implemented"; return }

func (cs *connSet) RemoveAddr(addr net.Addr) { _ = "STUB: not implemented"; return }

func (cs *connSet) Set(c net.Conn, ips []net.IP) { _ = "STUB: not implemented"; return }
