package mock

import (
	"net"

	"github.com/cometbft/cometbft/libs/service"
	"github.com/cometbft/cometbft/p2p"
	"github.com/cometbft/cometbft/p2p/conn"
)

type Peer struct {
	*service.BaseService
	ip                   net.IP
	id                   p2p.ID
	addr                 *p2p.NetAddress
	kv                   map[string]interface{}
	Outbound, Persistent bool
}

// NewPeer creates and starts a new mock peer. If the ip
// is nil, random routable address is used.
func NewPeer(ip net.IP) *Peer { _ = "STUB: not implemented"; return nil }

func (mp *Peer) FlushStop()                  { _ = "STUB: not implemented"; return } //nolint:errcheck //ignore error
func (mp *Peer) TrySend(_ p2p.Envelope) bool { _ = "STUB: not implemented"; return false }
func (mp *Peer) Send(_ p2p.Envelope) bool    { _ = "STUB: not implemented"; return false }
func (mp *Peer) NodeInfo() p2p.NodeInfo      { _ = "STUB: not implemented"; return *new(p2p.NodeInfo) }

func (mp *Peer) Status() conn.ConnectionStatus {
	_ = "STUB: not implemented"
	return *new(conn.ConnectionStatus)
}
func (mp *Peer) ID() p2p.ID                 { _ = "STUB: not implemented"; return *new(p2p.ID) }
func (mp *Peer) IsOutbound() bool           { _ = "STUB: not implemented"; return false }
func (mp *Peer) IsPersistent() bool         { _ = "STUB: not implemented"; return false }
func (mp *Peer) Get(key string) interface{} { _ = "STUB: not implemented"; return nil }

func (mp *Peer) Set(key string, value interface{}) { _ = "STUB: not implemented"; return }

func (mp *Peer) RemoteIP() net.IP            { _ = "STUB: not implemented"; return *new(net.IP) }
func (mp *Peer) SocketAddr() *p2p.NetAddress { _ = "STUB: not implemented"; return nil }
func (mp *Peer) RemoteAddr() net.Addr        { _ = "STUB: not implemented"; return *new(net.Addr) }
func (mp *Peer) CloseConn() error            { _ = "STUB: not implemented"; return nil }
func (mp *Peer) SetRemovalFailed()           { _ = "STUB: not implemented"; return }
func (mp *Peer) GetRemovalFailed() bool      { _ = "STUB: not implemented"; return false }
func (*Peer) HasIPChanged() bool             { _ = "STUB: not implemented"; return false }
