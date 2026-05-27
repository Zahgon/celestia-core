package mock

import (
	"github.com/cometbft/cometbft/p2p"
	"github.com/cometbft/cometbft/p2p/conn"
)

type Reactor struct {
	p2p.BaseReactor

	Channels []*conn.ChannelDescriptor
}

func NewReactor() *Reactor { _ = "STUB: not implemented"; return nil }

func (r *Reactor) GetChannels() []*conn.ChannelDescriptor { _ = "STUB: not implemented"; return nil }
func (r *Reactor) AddPeer(peer p2p.Peer)                  { _ = "STUB: not implemented"; return }
func (r *Reactor) RemovePeer(_ p2p.Peer, _ interface{})   { _ = "STUB: not implemented"; return }
func (r *Reactor) Receive(_ p2p.Envelope)                 { _ = "STUB: not implemented"; return }
func (r *Reactor) InitPeer(peer p2p.Peer) (p2p.Peer, error) {
	_ = "STUB: not implemented"
	return *new(p2p.Peer), nil
}
