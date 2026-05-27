package p2p

import (
	"context"
	"net"
	"time"

	"github.com/cosmos/gogoproto/proto"

	"github.com/cometbft/cometbft/crypto"
	"github.com/cometbft/cometbft/libs/trace"
	"github.com/cometbft/cometbft/p2p/conn"
)

const (
	defaultDialTimeout      = time.Second
	defaultFilterTimeout    = 5 * time.Second
	defaultHandshakeTimeout = 3 * time.Second
)

// IPResolver is a behavior subset of net.Resolver.
type IPResolver interface {
	LookupIPAddr(context.Context, string) ([]net.IPAddr, error)
}

// accept is the container to carry the upgraded connection and NodeInfo from an
// asynchronously running routine to the Accept method.
type accept struct {
	netAddr  *NetAddress
	conn     net.Conn
	nodeInfo NodeInfo
	err      error
}

// peerConfig is used to bundle data we need to fully setup a Peer with an
// MConn, provided by the caller of Accept and Dial (currently the Switch). This
// a temporary measure until reactor setup is less dynamic and we introduce the
// concept of PeerBehaviour to communicate about significant Peer lifecycle
// events.
// TODO(xla): Refactor out with more static Reactor setup and PeerBehaviour.
type peerConfig struct {
	chDescs     []*conn.ChannelDescriptor
	onPeerError func(Peer, interface{}, string)
	outbound    bool
	// isPersistent allows you to set a function, which, given socket address
	// (for outbound peers) OR self-reported address (for inbound peers), tells
	// if the peer is persistent or not.
	isPersistent  func(*NetAddress) bool
	reactorsByCh  map[byte]Reactor
	msgTypeByChID map[byte]proto.Message
	metrics       *Metrics
	mlc           *metricsLabelCache
}

// Transport emits and connects to Peers. The implementation of Peer is left to
// the transport. Each transport is also responsible to filter establishing
// peers specific to its domain.
type Transport interface {
	// Listening address.
	NetAddress() NetAddress

	// Accept returns a newly connected Peer.
	Accept(peerConfig) (Peer, error)

	// Dial connects to the Peer for the address.
	Dial(NetAddress, peerConfig) (Peer, error)

	// Cleanup any resources associated with Peer.
	Cleanup(Peer)
}

// transportLifecycle bundles the methods for callers to control start and stop
// behavior.
type transportLifecycle interface {
	Close() error
	Listen(NetAddress) error
}

// ConnFilterFunc to be implemented by filter hooks after a new connection has
// been established. The set of exisiting connections is passed along together
// with all resolved IPs for the new connection.
type ConnFilterFunc func(ConnSet, net.Conn, []net.IP) error

func generateTraceID() string { _ = "STUB: not implemented"; return "" }

// ConnDuplicateIPFilter resolves and keeps all ips for an incoming connection
// and refuses new ones if they come from a known ip.
func ConnDuplicateIPFilter() ConnFilterFunc { _ = "STUB: not implemented"; return *new(ConnFilterFunc) }

// MultiplexTransportOption sets an optional parameter on the
// MultiplexTransport.
type MultiplexTransportOption func(*MultiplexTransport)

// MultiplexTransportConnFilters sets the filters for rejection new connections.
func MultiplexTransportConnFilters(
	filters ...ConnFilterFunc,
) MultiplexTransportOption {
	_ = "STUB: not implemented"
	return *new(MultiplexTransportOption)
}

// MultiplexTransportFilterTimeout sets the timeout waited for filter calls to
// return.
func MultiplexTransportFilterTimeout(
	timeout time.Duration,
) MultiplexTransportOption {
	_ = "STUB: not implemented"
	return *new(MultiplexTransportOption)
}

// MultiplexTransportResolver sets the Resolver used for ip lokkups, defaults to
// net.DefaultResolver.
func MultiplexTransportResolver(resolver IPResolver) MultiplexTransportOption {
	_ = "STUB: not implemented"
	return *new(MultiplexTransportOption)
}

// MultiplexTransportMaxIncomingConnections sets the maximum number of
// simultaneous connections (incoming). Default: 0 (unlimited)
func MultiplexTransportMaxIncomingConnections(n int) MultiplexTransportOption {
	_ = "STUB: not implemented"
	return *new(MultiplexTransportOption)
}

// MultiplexTransport accepts and dials tcp connections and upgrades them to
// multiplexed peers.
type MultiplexTransport struct {
	netAddr                NetAddress
	listener               net.Listener
	maxIncomingConnections int // see MaxIncomingConnections

	acceptc chan accept
	closec  chan struct{}

	// Lookup table for duplicate ip and id checks.
	conns       ConnSet
	connFilters []ConnFilterFunc

	dialTimeout      time.Duration
	filterTimeout    time.Duration
	handshakeTimeout time.Duration
	nodeInfo         NodeInfo
	nodeKey          NodeKey
	resolver         IPResolver

	// TODO(xla): This config is still needed as we parameterise peerConn and
	// peer currently. All relevant configuration should be refactored into options
	// with sane defaults.
	mConfig conn.MConnConfig
	// the tracer is passed to peers for collecting trace data
	tracer trace.Tracer
}

// Test multiplexTransport for interface completeness.
var _ Transport = (*MultiplexTransport)(nil)
var _ transportLifecycle = (*MultiplexTransport)(nil)

// NewMultiplexTransport returns a tcp connected multiplexed peer.
func NewMultiplexTransport(
	nodeInfo NodeInfo,
	nodeKey NodeKey,
	mConfig conn.MConnConfig,
	tracer trace.Tracer,
) *MultiplexTransport {
	_ = "STUB: not implemented"
	return nil
}

// NetAddress implements Transport.
func (mt *MultiplexTransport) NetAddress() NetAddress {
	_ = "STUB: not implemented"

	// Accept implements Transport.
	return *new(NetAddress)
}

func (mt *MultiplexTransport) Accept(cfg peerConfig) (Peer, error) {
	_ = "STUB: not implemented"

	// This case should never have any side-effectful/blocking operations to
	// ensure that quality peers are ready to be used.
	return *new(Peer), nil
}

// Dial implements Transport.
func (mt *MultiplexTransport) Dial(
	addr NetAddress,
	cfg peerConfig,
) (Peer, error) {
	_ = "STUB: not implemented"
	return *new(Peer), nil
}

// so we have time to do peer handshakes and get set up.

// TODO(xla): Evaluate if we should apply filters if we explicitly dial.

// Close implements transportLifecycle.
func (mt *MultiplexTransport) Close() error { _ = "STUB: not implemented"; return nil }

// Listen implements transportLifecycle.
func (mt *MultiplexTransport) Listen(addr NetAddress) error { _ = "STUB: not implemented"; return nil }

// AddChannel registers a channel to nodeInfo.
// NOTE: NodeInfo must be of type DefaultNodeInfo else channels won't be updated
// This is a bit messy at the moment but is cleaned up in the following version
// when NodeInfo changes from an interface to a concrete type
func (mt *MultiplexTransport) AddChannel(chID byte) { _ = "STUB: not implemented"; return }

func (mt *MultiplexTransport) acceptPeers() { _ = "STUB: not implemented"; return }

// If Close() has been called, silently exit.

// Transport is not closed

// Connection upgrade and filtering should be asynchronous to avoid
// Head-of-line blocking[0].
// Reference:  https://github.com/tendermint/tendermint/issues/2047
//
// [0] https://en.wikipedia.org/wiki/Head-of-line_blocking

// Give up if the transport was closed.

// Make the upgraded peer available.

// Give up if the transport was closed.

// Cleanup removes the given address from the connections set and
// closes the connection.
func (mt *MultiplexTransport) Cleanup(p Peer) { _ = "STUB: not implemented"; return }

func (mt *MultiplexTransport) cleanup(c net.Conn) error { _ = "STUB: not implemented"; return nil }

func (mt *MultiplexTransport) filterConn(c net.Conn) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Reject if connection is already present.

// Resolve ips for incoming conn.

func (mt *MultiplexTransport) upgrade(
	c net.Conn,
	dialedAddr *NetAddress,
) (secretConn *conn.SecretConnection, nodeInfo NodeInfo, err error) {
	_ = "STUB: not implemented"
	return nil, *new(NodeInfo), nil
}

func (mt *MultiplexTransport) wrapPeer(
	c net.Conn,
	ni NodeInfo,
	cfg peerConfig,
	socketAddr *NetAddress,
) Peer {
	_ = "STUB: not implemented"
	return *new(Peer)
}

func handshake(
	c net.Conn,
	timeout time.Duration,
	nodeInfo NodeInfo,
) (NodeInfo, error) {
	_ = "STUB: not implemented"
	return *new(NodeInfo), nil
}

func upgradeSecretConn(
	c net.Conn,
	timeout time.Duration,
	privKey crypto.PrivKey,
) (*conn.SecretConnection, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func resolveIPs(resolver IPResolver, c net.Conn) ([]net.IP, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
