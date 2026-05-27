package p2p

import (
	"net"
	"sync/atomic"
	"time"

	"github.com/cosmos/gogoproto/proto"

	"github.com/cometbft/cometbft/libs/cmap"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cometbft/cometbft/libs/service"
	"github.com/cometbft/cometbft/libs/trace"

	cmtconn "github.com/cometbft/cometbft/p2p/conn"
)

//go:generate ../scripts/mockery_generate.sh Peer

const metricsTickerDuration = 10 * time.Second

// Peer is an interface representing a peer connected on a reactor.
type Peer interface {
	service.Service
	FlushStop()

	ID() ID               // peer's cryptographic ID
	RemoteIP() net.IP     // remote IP of the connection
	RemoteAddr() net.Addr // remote address of the connection

	IsOutbound() bool   // did we dial the peer
	IsPersistent() bool // do we redial this peer when we disconnect

	CloseConn() error // close original connection

	NodeInfo() NodeInfo // peer's info
	Status() cmtconn.ConnectionStatus
	SocketAddr() *NetAddress // actual address of the socket

	Send(Envelope) bool
	TrySend(Envelope) bool

	Set(string, interface{})
	Get(string) interface{}

	SetRemovalFailed()
	GetRemovalFailed() bool

	HasIPChanged() bool // has the peer's IP changed
}

type IntrospectivePeer interface {
	Peer
	Metrics() *Metrics
	ValueToMetricLabel(i any) string
	TraceClient() trace.Tracer
}

//----------------------------------------------------------

// peerConn contains the raw connection and its config.
type peerConn struct {
	outbound   bool
	persistent bool
	conn       net.Conn // source connection

	socketAddr *NetAddress
}

func newPeerConn(
	outbound, persistent bool,
	conn net.Conn,
	socketAddr *NetAddress,
) peerConn {
	_ = "STUB: not implemented"
	return *new(peerConn)
}

// ID only exists for SecretConnection.
// NOTE: Will panic if conn is not *SecretConnection.
func (pc peerConn) ID() ID { _ = "STUB: not implemented"; return *new(ID) }

// Return the IP from the connection RemoteAddr
func (pc peerConn) RemoteIP() net.IP { _ = "STUB: not implemented"; return *new(net.IP) }

// peer implements Peer.
//
// Before using a peer, you will need to perform a handshake on connection.
type peer struct {
	service.BaseService

	// raw peerConn and the multiplex connection
	peerConn
	mconn *cmtconn.MConnection

	// peer's node info and the channel it knows about
	// channels = nodeInfo.Channels
	// cached to avoid copying nodeInfo in hasChannel
	nodeInfo NodeInfo
	channels []byte

	// User data
	Data *cmap.CMap

	metrics     *Metrics
	mlc         *metricsLabelCache
	traceClient trace.Tracer

	// Atomic fields for thread-safe concurrent access
	removalAttemptFailed atomic.Bool
	cachedIP             atomic.Pointer[net.IP]

	// Lifecycle coordination: stopped protects metricsTicker from concurrent access
	// Once stopped transitions from false->true, no further lifecycle operations occur
	stopped       atomic.Bool
	metricsTicker *time.Ticker
}

type PeerOption func(*peer)

func WithPeerTracer(t trace.Tracer) PeerOption { _ = "STUB: not implemented"; return *new(PeerOption) }

func newPeer(
	pc peerConn,
	mConfig cmtconn.MConnConfig,
	nodeInfo NodeInfo,
	reactorsByCh map[byte]Reactor,
	_ map[byte]proto.Message,
	chDescs []*cmtconn.ChannelDescriptor,
	onPeerError func(Peer, interface{}, string),
	mlc *metricsLabelCache,
	options ...PeerOption,
) *peer {
	_ = "STUB: not implemented"
	return nil
}

// String representation.
func (p *peer) String() string { _ = "STUB: not implemented"; return "" }

//---------------------------------------------------
// Implements service.Service

// SetLogger implements BaseService.
func (p *peer) SetLogger(l log.Logger) { _ = "STUB: not implemented"; return }

// OnStart implements BaseService.
func (p *peer) OnStart() error { _ = "STUB: not implemented"; return nil }

// FlushStop mimics OnStop but additionally ensures that all successful
// .Send() calls will get flushed before closing the connection.
// Thread-safe and idempotent - can be called multiple times safely.
func (p *peer) FlushStop() { _ = "STUB: not implemented"; return }

// Already stopped

// stop everything and close the conn

func (p *peer) Metrics() *Metrics { _ = "STUB: not implemented"; return nil }

func (p *peer) ValueToMetricLabel(i any) string { _ = "STUB: not implemented"; return "" }

func (p *peer) TraceClient() trace.Tracer {
	_ = "STUB: not implemented"
	return *

	// OnStop implements BaseService.
	// Thread-safe and idempotent - can be called multiple times safely.
	new(trace.Tracer)
}

func (p *peer) OnStop() { _ = "STUB: not implemented"; return }

// Already stopped

// stop everything and close the conn

//---------------------------------------------------
// Implements Peer

// ID returns the peer's ID - the hex encoded hash of its pubkey.
func (p *peer) ID() ID {
	_ = "STUB: not implemented"
	return *

	// IsOutbound returns true if the connection is outbound, false otherwise.
	new(ID)
}

func (p *peer) IsOutbound() bool { _ = "STUB: not implemented"; return false }

//nolint:staticcheck

// IsPersistent returns true if the peer is persitent, false otherwise.
func (p *peer) IsPersistent() bool { _ = "STUB: not implemented"; return false }

//nolint:staticcheck

// NodeInfo returns a copy of the peer's NodeInfo.
func (p *peer) NodeInfo() NodeInfo {
	_ = "STUB: not implemented"

	// RemoteIP returns the IP from the connection RemoteAddr with atomic caching
	return *new(NodeInfo)
}

func (p *peer) RemoteIP() net.IP {
	_ = "STUB: not implemented"
	// Fast path: return cached IP if available
	return *new(net.IP)
}

// Slow path: perform DNS lookup and cache result

// HasIPChanged returns true if the peer's IP has changed.
// This method clears the cached IP and compares with a fresh lookup.
func (p *peer) HasIPChanged() bool {
	_ = "STUB: not implemented"
	// Get the currently cached IP
	return false
}

// No cached IP, so no change detected

// Clear the cached IP to force a fresh lookup

// Get the current IP (will perform fresh DNS lookup)

// Compare the IPs

// SocketAddr returns the address of the socket.
// For outbound peers, it's the address dialed (after DNS resolution).
// For inbound peers, it's the address returned by the underlying connection
// (not what's reported in the peer's NodeInfo).
func (p *peer) SocketAddr() *NetAddress { _ = "STUB: not implemented"; return nil }

//nolint:staticcheck

// Status returns the peer's ConnectionStatus.
func (p *peer) Status() cmtconn.ConnectionStatus {
	_ = "STUB: not implemented"
	return *

	// Send msg bytes to the channel identified by chID byte. Returns false if the
	// send queue is full after timeout, specified by MConnection.
	new(cmtconn.ConnectionStatus)
}

func (p *peer) Send(e Envelope) bool { _ = "STUB: not implemented"; return false }

// TrySend msg bytes to the channel identified by chID byte. Immediately returns
// false if the send queue is full.
func (p *peer) TrySend(e Envelope) bool { _ = "STUB: not implemented"; return false }

func (p *peer) send(chID byte, msg proto.Message, sendFunc func(byte, []byte) bool) bool {
	_ = "STUB: not implemented"
	return false
}

//nolint:prealloc

// Get the data for a given key.
func (p *peer) Get(key string) interface{} { _ = "STUB: not implemented"; return nil }

// Set sets the data for the given key.
func (p *peer) Set(key string, data interface{}) { _ = "STUB: not implemented"; return }

// hasChannel returns true if the peer reported
// knowing about the given chID.
func (p *peer) hasChannel(chID byte) bool { _ = "STUB: not implemented"; return false }

// NOTE: probably will want to remove this
// but could be helpful while the feature is new

// CloseConn closes original connection. Used for cleaning up in cases where the peer had not been started at all.
func (p *peer) CloseConn() error { _ = "STUB: not implemented"; return nil }

//nolint:staticcheck

func (p *peer) SetRemovalFailed() { _ = "STUB: not implemented"; return }

func (p *peer) GetRemovalFailed() bool { _ = "STUB: not implemented"; return false }

//---------------------------------------------------
// methods only used for testing
// TODO: can we remove these?

// CloseConn closes the underlying connection
func (pc *peerConn) CloseConn() {
	_ = "STUB: not implemented"

	// RemoteAddr returns peer's remote network address.
	return
}

func (p *peer) RemoteAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

//nolint:staticcheck

// CanSend returns true if the send queue is not full, false otherwise.
func (p *peer) CanSend(chID byte) bool { _ = "STUB: not implemented"; return false }

//---------------------------------------------------

func PeerMetrics(metrics *Metrics) PeerOption { _ = "STUB: not implemented"; return *new(PeerOption) }

func (p *peer) metricsReporter() { _ = "STUB: not implemented"; return }

//------------------------------------------------------------------
// helper funcs

func createMConnection(
	conn net.Conn,
	p *peer,
	reactorsByCh map[byte]Reactor,
	chDescs []*cmtconn.ChannelDescriptor,
	onPeerError func(Peer, interface{}, string),
	config cmtconn.MConnConfig,
) *cmtconn.MConnection {
	_ = "STUB: not implemented"
	return nil
}

// Note that its ok to panic here as it's caught in the conn._recover,
// which does onPeerError.
