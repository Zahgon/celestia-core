package evidence

import (
	"github.com/cosmos/gogoproto/proto"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cometbft/cometbft/p2p"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cometbft/cometbft/types"
)

const (
	EvidenceChannel = byte(0x38)

	maxMsgSize = 1048576 // 1MB TODO make it configurable

	// broadcast all uncommitted evidence this often. This sets when the reactor
	// goes back to the start of the list and begins sending the evidence again.
	// Most evidence should be committed in the very next block that is why we wait
	// just over the block production rate before sending evidence again.
	broadcastEvidenceIntervalS = 10
	// If a message fails wait this much before sending it again
	peerRetryMessageIntervalMS = 100
	// ReactorIncomingMessageQueueSize the size of the reactor's message queue.
	ReactorIncomingMessageQueueSize = 1
)

// Reactor handles evpool evidence broadcasting amongst peers.
type Reactor struct {
	p2p.BaseReactor
	evpool   *Pool
	eventBus *types.EventBus
}

// NewReactor returns a new Reactor with the given config and evpool.
func NewReactor(evpool *Pool) *Reactor { _ = "STUB: not implemented"; return nil }

// SetLogger sets the Logger on the reactor and the underlying Evidence.
func (evR *Reactor) SetLogger(l log.Logger) { _ = "STUB: not implemented"; return }

// GetChannels implements Reactor.
// It returns the list of channels for this reactor.
func (evR *Reactor) GetChannels() []*p2p.ChannelDescriptor { _ = "STUB: not implemented"; return nil }

// AddPeer implements Reactor.
func (evR *Reactor) AddPeer(peer p2p.Peer) { _ = "STUB: not implemented"; return }

// Receive implements Reactor.
// It adds any received evidence to the evpool.
func (evR *Reactor) Receive(e p2p.Envelope) { _ = "STUB: not implemented"; return }

// punish peer

// continue to the next piece of evidence

// SetEventBus implements events.Eventable.
func (evR *Reactor) SetEventBus(b *types.EventBus) {
	_ = "STUB: not implemented"

	// Modeled after the mempool routine.
	// - Evidence accumulates in a clist.
	// - Each peer has a routine that iterates through the clist,
	// sending available evidence to the peer.
	// - If we're waiting for new evidence and the list is not empty,
	// start iterating from the beginning again.
	return
}

func (evR *Reactor) broadcastEvidenceRoutine(peer p2p.Peer) { _ = "STUB: not implemented"; return }

// This happens because the CElement we were looking at got garbage
// collected (removed). That is, .NextWait() returned nil. Go ahead and
// start from the beginning.

// Wait until evidence is available

// start from the beginning every tick.
// TODO: only do this if we're at the end of the list!

// see the start of the for loop for nil check

// Returns the message to send to the peer, or nil if the evidence is invalid for the peer.
// If message is nil, we should sleep and try again.
func (evR Reactor) prepareEvidenceMessage(
	peer p2p.Peer,
	ev types.Evidence,
) (evis []types.Evidence) {
	_ = "STUB: not implemented"

	// make sure the peer is up to date
	return nil
}

// Peer does not have a state yet. We set it in the consensus reactor, but
// when we add peer in Switch, the order we call reactors#AddPeer is
// different every time due to us using a map. Sometimes other reactors
// will be initialized before the consensus reactor. We should wait a few
// milliseconds and retry.

// NOTE: We only send evidence to peers where
// peerHeight - maxAge < evidenceHeight < peerHeight

// peer is behind. sleep while he catches up

// evidence is too old relative to the peer, skip

// NOTE: if evidence is too old for an honest peer, then we're behind and
// either it already got committed or it never will!

// send evidence

// PeerState describes the state of a peer.
type PeerState interface {
	GetHeight() int64
}

// encodemsg takes a array of evidence
// returns the byte encoding of the List Message
func evidenceListToProto(evis []types.Evidence) (*cmtproto.EvidenceList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func evidenceListFromProto(m proto.Message) ([]types.Evidence, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (evR *Reactor) OnStop() { _ = "STUB: not implemented"; return }
