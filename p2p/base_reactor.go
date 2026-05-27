package p2p

import (
	"context"

	"github.com/cometbft/cometbft/libs/trace"

	"github.com/cometbft/cometbft/libs/service"
	"github.com/cometbft/cometbft/p2p/conn"
	"github.com/cosmos/gogoproto/proto"
)

// ProcessorFunc is the message processor function type.
type ProcessorFunc func(context.Context, <-chan UnmarshalResult)

// Reactor is responsible for handling incoming messages on one or more
// Channel. Switch calls GetChannels when reactor is added to it. When a new
// peer joins our node, InitPeer and AddPeer are called. RemovePeer is called
// when the peer is stopped. Receive is called when a message is received on a
// channel associated with this reactor.
//
// Peer#Send or Peer#TrySend should be used to send the message to a peer.
type Reactor interface {
	service.Service // Start, Stop

	// SetSwitch allows setting a switch.
	SetSwitch(*Switch)

	// GetChannels returns the list of MConnection.ChannelDescriptor. Make sure
	// that each ID is unique across all the reactors added to the switch.
	GetChannels() []*conn.ChannelDescriptor

	// InitPeer is called by the switch before the peer is started. Use it to
	// initialize data for the peer (e.g. peer state).
	//
	// NOTE: The switch won't call AddPeer nor RemovePeer if it fails to start
	// the peer. Do not store any data associated with the peer in the reactor
	// itself unless you don't want to have a state, which is never cleaned up.
	InitPeer(peer Peer) (Peer, error)

	// AddPeer is called by the switch after the peer is added and successfully
	// started. Use it to start goroutines communicating with the peer.
	AddPeer(peer Peer)

	// RemovePeer is called by the switch when the peer is stopped (due to error
	// or other reason).
	RemovePeer(peer Peer, reason interface{})

	// Receive is called by the switch when an envelope is received from any connected
	// peer on any of the channels registered by the reactor
	Receive(Envelope)

	// QueueUnprocessedEnvelope is called by the switch when an unprocessed
	// envelope is received. Unprocessed envelopes are immediately buffered in a
	// queue to avoid blocking. Incoming messages are then passed to a
	// processing function. The default processing function unmarshals the
	// messages in the order the sender sent them and then calls Receive on the
	// reactor. The queue size and the processing function can be changed via
	// passing options to the base reactor.
	QueueUnprocessedEnvelope(e UnprocessedEnvelope)
}

//--------------------------------------

type UnmarshalResult struct {
	Src           IntrospectivePeer
	Msg           proto.Message
	BytesReceived int
	ChannelID     byte
	Err           error
}

type BaseReactor struct {
	service.BaseService // Provides Start, Stop, .Quit
	Switch              *Switch

	incoming     chan UnmarshalResult
	queueingFunc func(UnprocessedEnvelope)

	ctx    context.Context
	cancel context.CancelFunc
	chIDs  map[byte]proto.Message
	// processor is called with the incoming channel and is responsible for
	// unmarshalling the messages and calling Receive on the reactor.
	processor   ProcessorFunc
	name        string
	traceClient trace.Tracer
}

type ReactorOptions func(*BaseReactor)

func NewBaseReactor(name string, impl Reactor, opts ...ReactorOptions) *BaseReactor {
	_ = "STUB: not implemented"
	return nil
}

// set by the switch later

// Will be set after base is created

// Set the processor after base is created, only if it hasn't been set by options

// Try to stop the reactor gracefully only if it's running

// WithProcessor sets the processor function for the reactor. The processor
// function is called with the incoming channel and is responsible for
// calling Receive on the reactor.
func WithProcessor(processor ProcessorFunc) ReactorOptions {
	_ = "STUB: not implemented"
	return *new(ReactorOptions)
}

// WithTraceClient sets the tracing client using options
func WithTraceClient(traceClient trace.Tracer) ReactorOptions {
	_ = "STUB: not implemented"
	return *new(ReactorOptions)
}

// SetTraceClient sets the tracing client.
func (br *BaseReactor) SetTraceClient(traceClient trace.Tracer) { _ = "STUB: not implemented"; return }

// WithQueueingFunc sets the queuing function to use when receiving a message.
func WithQueueingFunc(queuingFunc func(UnprocessedEnvelope)) ReactorOptions {
	_ = "STUB: not implemented"
	return *new(ReactorOptions)
}

// WithIncomingQueueSize sets the size of the incoming message queue for a
// reactor.
func WithIncomingQueueSize(size int) ReactorOptions {
	_ = "STUB: not implemented"
	return *new(ReactorOptions)
}

// QueueUnprocessedEnvelope is called by the switch when an unprocessed
// envelope is received. Unprocessed envelopes are immediately buffered in a
// queue to avoid blocking. The size of the queue can be changed by passing
// options to the base reactor.
func (br *BaseReactor) QueueUnprocessedEnvelope(e UnprocessedEnvelope) {
	_ = "STUB: not implemented"
	return
}

// if the context is done, do nothing.

// if not, add the item to the channel.

// TryQueueUnprocessedEnvelope an alternative to QueueUnprocessedEnvelope that attempts to queue an unprocessed envelope.
// If the queue is full, it drops the envelope.
func (br *BaseReactor) TryQueueUnprocessedEnvelope(e UnprocessedEnvelope) {
	_ = "STUB: not implemented"
	return
}

func (br *BaseReactor) unmarshalEnvelope(e UnprocessedEnvelope) UnmarshalResult {
	_ = "STUB: not implemented"
	return *new(UnmarshalResult)
}

func (br *BaseReactor) OnStop() {
	_ = "STUB: not implemented"

	// ProcessorWithReactor unmarshalls the message and calls Receive on the reactor.
	// This preserves the sender's original order for all messages and supports panic recovery with peer disconnection.
	return
}

func ProcessorWithReactor(impl Reactor, baseReactor *BaseReactor) ProcessorFunc {
	_ = "STUB: not implemented"
	return *new(ProcessorFunc)
}

// this means the channel was closed.

// Process message with panic recovery for individual peer

// ProtectPanic provides panic recovery for reactor operations involving a specific peer.
// If a panic occurs, it will disconnect the peer with an appropriate error message.
// Usage: defer baseReactor.ProtectPanic(peer)
func (br *BaseReactor) ProtectPanic(peer Peer) { _ = "STUB: not implemented"; return }

func disconnectPeer(baseReactor *BaseReactor, peer Peer, reason, reactor string) {
	_ = "STUB: not implemented"
	// the switch is added for all reactors so should be here. the worst case if not is
	// that the peer doesn't get disconnected.
	return
}

func (br *BaseReactor) SetSwitch(sw *Switch) { _ = "STUB: not implemented"; return }

func (*BaseReactor) GetChannels() []*conn.ChannelDescriptor { _ = "STUB: not implemented"; return nil }
func (*BaseReactor) AddPeer(Peer)                           { _ = "STUB: not implemented"; return }
func (*BaseReactor) RemovePeer(Peer, interface{})           { _ = "STUB: not implemented"; return }
func (*BaseReactor) Receive(Envelope)                       { _ = "STUB: not implemented"; return }
func (*BaseReactor) InitPeer(peer Peer) (Peer, error) {
	_ = "STUB: not implemented"
	return *new(Peer), nil
}
