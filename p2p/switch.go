package p2p

import (
	"time"

	"github.com/cosmos/gogoproto/proto"

	"github.com/cometbft/cometbft/config"
	"github.com/cometbft/cometbft/libs/cmap"
	"github.com/cometbft/cometbft/libs/rand"
	"github.com/cometbft/cometbft/libs/service"
	"github.com/cometbft/cometbft/libs/trace"
	"github.com/cometbft/cometbft/p2p/conn"
)

const (
	// wait a random amount of time from this interval
	// before dialing peers or reconnecting to help prevent DoS
	dialRandomizerIntervalMilliseconds = 3000

	// repeatedly try to reconnect for a few minutes
	// ie. 5 * 20 = 100s
	reconnectAttempts = 20
	reconnectInterval = 5 * time.Second

	// then move into exponential backoff mode for ~1day
	// ie. 3**10 = 16hrs
	reconnectBackOffAttempts    = 10
	reconnectBackOffBaseSeconds = 3
)

// MConnConfig returns an MConnConfig with fields updated
// from the P2PConfig.
func MConnConfig(cfg *config.P2PConfig) conn.MConnConfig {
	_ = "STUB: not implemented"
	return *new(conn.MConnConfig)
}

//-----------------------------------------------------------------------------

// An AddrBook represents an address book from the pex package, which is used
// to store peer addresses.
type AddrBook interface {
	AddAddress(addr *NetAddress, src *NetAddress) error
	AddPrivateIDs([]string)
	AddOurAddress(*NetAddress)
	OurAddress(*NetAddress) bool
	MarkGood(ID)
	RemoveAddress(*NetAddress)
	HasAddress(*NetAddress) bool
	Save()
}

// PeerFilterFunc to be implemented by filter hooks after a new Peer has been
// fully setup.
type PeerFilterFunc func(IPeerSet, Peer) error

//-----------------------------------------------------------------------------

// Switch handles peer connections and exposes an API to receive incoming messages
// on `Reactors`.  Each `Reactor` is responsible for handling incoming messages of one
// or more `Channels`.  So while sending outgoing messages is typically performed on the peer,
// incoming messages are received on the reactor.
type Switch struct {
	service.BaseService

	config           *config.P2PConfig
	reactors         map[string]Reactor
	chDescs          []*conn.ChannelDescriptor
	reactorsByCh     map[byte]Reactor
	msgTypeByChID    map[byte]proto.Message
	peerSetByReactor map[string]*PeerSet
	peers            *PeerSet
	dialing          *cmap.CMap
	reconnecting     *cmap.CMap
	nodeInfo         NodeInfo // our node info
	nodeKey          *NodeKey // our node privkey
	addrBook         AddrBook
	// peers addresses with whom we'll maintain constant connection
	persistentPeersAddrs []*NetAddress
	unconditionalPeerIDs map[ID]struct{}

	transport Transport

	filterTimeout time.Duration
	peerFilters   []PeerFilterFunc

	rng *rand.Rand // seed for randomizing dial times and orders

	metrics     *Metrics
	mlc         *metricsLabelCache
	traceClient trace.Tracer
}

// NetAddress returns the address the switch is listening on.
func (sw *Switch) NetAddress() *NetAddress { _ = "STUB: not implemented"; return nil }

// SwitchOption sets an optional parameter on the Switch.
type SwitchOption func(*Switch)

// NewSwitch creates a new Switch with the given config.
func NewSwitch(
	cfg *config.P2PConfig,
	transport Transport,
	options ...SwitchOption,
) *Switch {
	_ = "STUB: not implemented"
	return nil
}

// Ensure we have a completely undeterministic PRNG.

// SwitchFilterTimeout sets the timeout used for peer filters.
func SwitchFilterTimeout(timeout time.Duration) SwitchOption {
	_ = "STUB: not implemented"
	return *new(SwitchOption)
}

// SwitchPeerFilters sets the filters for rejection of new peers.
func SwitchPeerFilters(filters ...PeerFilterFunc) SwitchOption {
	_ = "STUB: not implemented"
	return *new(SwitchOption)
}

// WithMetrics sets the metrics.
func WithMetrics(metrics *Metrics) SwitchOption {
	_ = "STUB: not implemented"
	return *new(SwitchOption)
}

// WithTracer sets the tracer.
func WithTracer(tracer trace.Tracer) SwitchOption {
	_ = "STUB: not implemented"
	return *new(SwitchOption)
}

//---------------------------------------------------------------------
// Switch setup

// AddReactor adds the given reactor to the switch.
// NOTE: Not goroutine safe.
func (sw *Switch) AddReactor(name string, reactor Reactor) Reactor {
	_ = "STUB: not implemented"
	return *new(Reactor)
}

// No two reactors can share the same channel.

// RemoveReactor removes the given Reactor from the Switch.
// NOTE: Not goroutine safe.
func (sw *Switch) RemoveReactor(name string, reactor Reactor) { _ = "STUB: not implemented"; return }

// remove channel description

// Reactors returns a map of reactors registered on the switch.
// NOTE: Not goroutine safe.
func (sw *Switch) Reactors() map[string]Reactor {
	_ = "STUB: not implemented"

	// Reactor returns the reactor with the given name.
	// NOTE: Not goroutine safe.
	return nil
}

func (sw *Switch) Reactor(name string) Reactor {
	_ = "STUB: not implemented"
	return *

	// SetNodeInfo sets the switch's NodeInfo for checking compatibility and handshaking with other nodes.
	// NOTE: Not goroutine safe.
	new(Reactor)
}

func (sw *Switch) SetNodeInfo(nodeInfo NodeInfo) { _ = "STUB: not implemented"; return }

// NodeInfo returns the switch's NodeInfo.
// NOTE: Not goroutine safe.
func (sw *Switch) NodeInfo() NodeInfo {
	_ = "STUB: not implemented"

	// SetNodeKey sets the switch's private key for authenticated encryption.
	// NOTE: Not goroutine safe.
	return *new(NodeInfo)
}

func (sw *Switch) SetNodeKey(nodeKey *NodeKey) { _ = "STUB: not implemented"; return }

//---------------------------------------------------------------------
// Service start/stop

// OnStart implements BaseService. It starts all the reactors and peers.
func (sw *Switch) OnStart() error {
	_ = "STUB: not implemented"
	// Start reactors
	return nil
}

// Start accepting Peers.

// OnStop implements BaseService. It stops all peers and reactors.
func (sw *Switch) OnStop() {
	_ = "STUB: not implemented"
	// Stop peers
	return
}

// Stop reactors

//---------------------------------------------------------------------
// Peers

// Broadcast runs a go routine for each attempted send, which will block trying
// to send for defaultSendTimeoutSeconds. Returns a channel which receives
// success values for each attempted send (false if times out). Channel will be
// closed once msg bytes are sent to all peers (or time out).
//
// NOTE: Broadcast uses goroutines, so order of broadcast may not be preserved.
func (sw *Switch) Broadcast(e Envelope) chan bool { _ = "STUB: not implemented"; return nil }

func (sw *Switch) peersForEnvelope(e Envelope) []Peer { _ = "STUB: not implemented"; return nil }

// NumPeers returns the count of outbound/inbound and outbound-dialing peers.
// unconditional peers are not counted here.
func (sw *Switch) NumPeers() (outbound, inbound, dialing int) {
	_ = "STUB: not implemented"
	return 0, 0, 0
}

func (sw *Switch) IsPeerUnconditional(id ID) bool { _ = "STUB: not implemented"; return false }

// MaxNumOutboundPeers returns a maximum number of outbound peers.
func (sw *Switch) MaxNumOutboundPeers() int { _ = "STUB: not implemented"; return 0 }

// Peers returns the set of peers that are connected to the switch.
func (sw *Switch) Peers() IPeerSet {
	_ = "STUB: not implemented"

	// StopPeerForError disconnects from a peer due to external error.
	// If the peer is persistent, it will attempt to reconnect.
	// TODO: make record depending on reason.
	return *new(IPeerSet)
}

func (sw *Switch) StopPeerForError(peer Peer, reason interface{}, reactorName string) {
	_ = "STUB: not implemented"
	return
}

// getPeerAddress returns the appropriate NetAddress for a given peer,
// handling both outbound and inbound peers.
func (sw *Switch) getPeerAddress(peer Peer) (*NetAddress, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// For inbound peers, get the self-reported address

// StopPeerGracefully disconnects from a peer gracefully.
// TODO: handle graceful disconnects.
func (sw *Switch) StopPeerGracefully(peer Peer, reactorName string) {
	_ = "STUB: not implemented"
	return
}

func (sw *Switch) stopAndRemovePeer(peer Peer, reason interface{}) {
	_ = "STUB: not implemented"
	return
}

// TODO: should return error to be handled accordingly

// Removing a peer should go last to avoid a situation where a peer
// reconnect to our node and the switch calls InitPeer before
// RemovePeer is finished.
// https://github.com/tendermint/tendermint/issues/3338

// Removal of the peer has failed. The function above sets a flag within the peer to mark this.
// We keep this message here as information to the developer.

// reconnectToPeer tries to reconnect to the addr, first repeatedly
// with a fixed interval (approximately 2 minutes), then with
// exponential backoff (approximately close to 24 hours).
// If no success after all that, it stops trying, and leaves it
// to the PEX/Addrbook to find the peer with the addr again
// NOTE: this will keep trying even if the handshake or auth fails.
// TODO: be more explicit with error types so we only retry on certain failures
//   - ie. if we're getting ErrDuplicatePeer we can stop
//     because the addrbook got us the peer back already
func (sw *Switch) reconnectToPeer(addr *NetAddress) { _ = "STUB: not implemented"; return }

// success

// sleep a set amount

// sleep an exponentially increasing amount

// success

// SetAddrBook allows to set address book on Switch.
func (sw *Switch) SetAddrBook(addrBook AddrBook) { _ = "STUB: not implemented"; return }

// MarkPeerAsGood marks the given peer as good when it did something useful
// like contributed to consensus.
func (sw *Switch) MarkPeerAsGood(peer Peer) { _ = "STUB: not implemented"; return }

//---------------------------------------------------------------------
// Dialing

type privateAddr interface {
	PrivateAddr() bool
}

func isPrivateAddr(err error) bool { _ = "STUB: not implemented"; return false }

// DialPeersAsync dials a list of peers asynchronously in random order.
// Used to dial peers from config on startup or from unsafe-RPC (trusted sources).
// It ignores ErrNetAddressLookup. However, if there are other errors, first
// encounter is returned.
// Nop if there are no peers.
func (sw *Switch) DialPeersAsync(peers []string) error { _ = "STUB: not implemented"; return nil }

// report all the errors

// return first non-ErrNetAddressLookup error

func (sw *Switch) dialPeersAsync(netAddrs []*NetAddress) { _ = "STUB: not implemented"; return }

// TODO: this code feels like it's in the wrong place.
// The integration tests depend on the addrBook being saved
// right away but maybe we can change that. Recall that
// the addrBook is only written to disk every 2min

// add peers to `addrBook`

// do not add our address or ID

// Persist some peers to disk right away.
// NOTE: integration tests depend on this

// permute the list, dial them in random order.

// DialPeerWithAddress dials the given peer and runs sw.addPeer if it connects
// and authenticates successfully.
// If we're currently dialing this address or it belongs to an existing peer,
// ErrCurrentlyDialingOrExistingAddress is returned.
func (sw *Switch) DialPeerWithAddress(addr *NetAddress) error {
	_ = "STUB: not implemented"
	return nil
}

// sleep for interval plus some random amount of ms on [0, dialRandomizerIntervalMilliseconds]
func (sw *Switch) randomSleep(interval time.Duration) { _ = "STUB: not implemented"; return }

// IsDialingOrExistingAddress returns true if switch has a peer with the given
// address or dialing it at the moment.
func (sw *Switch) IsDialingOrExistingAddress(addr *NetAddress) bool {
	_ = "STUB: not implemented"
	return false
}

// AddPersistentPeers allows you to set persistent peers. It ignores
// ErrNetAddressLookup. However, if there are other errors, first encounter is
// returned.
func (sw *Switch) AddPersistentPeers(addrs []string) error { _ = "STUB: not implemented"; return nil }

// report all the errors

// return first non-ErrNetAddressLookup error

func (sw *Switch) AddUnconditionalPeerIDs(ids []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (sw *Switch) AddPrivatePeerIDs(ids []string) error { _ = "STUB: not implemented"; return nil }

func (sw *Switch) IsPeerPersistent(na *NetAddress) bool { _ = "STUB: not implemented"; return false }

func (sw *Switch) acceptRoutine() { _ = "STUB: not implemented"; return }

// Remove the given address from the address book and add to our addresses
// to avoid dialing in the future.

// We could instead have a retry loop around the acceptRoutine,
// but that would need to stop and let the node shutdown eventually.
// So might as well panic and let process managers restart the node.
// There's no point in letting the node run without the acceptRoutine,
// since it won't be able to accept new connections.

// Ignore connection if we already have enough peers.

// dial the peer; make secret connection; authenticate against the dialed ID;
// add the peer.
// if dialing fails, start the reconnect loop. If handshake fails, it's over.
// If peer is started successfully, reconnectLoop will start when
// StopPeerForError is called.
func (sw *Switch) addOutboundPeerWithConfig(
	addr *NetAddress,
	cfg *config.P2PConfig,
) error {
	_ = "STUB: not implemented"
	return nil
}

// XXX(xla): Remove the leakage of test concerns in implementation.

// Remove the given address from the address book and add to our addresses
// to avoid dialing in the future.

// retry persistent peers after
// any dial error besides IsSelf()

func (sw *Switch) filterPeer(p Peer) error {
	_ = "STUB: not implemented"
	// Avoid duplicate
	return nil
}

// addPeer starts up the Peer and adds it to the Switch. Error is returned if
// the peer is filtered out or failed to start or can't be added.
func (sw *Switch) addPeer(p Peer) error { _ = "STUB: not implemented"; return nil }

// Handle the shut down case where the switch has stopped but we're
// concurrently trying to add a peer.

// XXX should this return an error or just log and terminate?

// Add some data to the peer, which is required by reactors.

// only update peer if the reactor accepted it

// Start the peer's send/recv routines.
// Must start it before adding it to the peer set
// to prevent Start and Stop from being called concurrently.

// Should never happen

// Add the peer to PeerSet. Do this before starting the reactors
// so that if Receive errors, we will find the peer and remove it.
// Add should not err since we already checked peers.Has().

// Start all the reactor protocols on the peer.

// peerSetForReactor retrieves the PeerSet associated with the given Reactor.
// Returns nil if the reactor is not registered.
func (sw *Switch) peerSetForReactor(r Reactor) *PeerSet { _ = "STUB: not implemented"; return nil }

// removePeerFromAllReactors removes the given peer from all reactors
func (sw *Switch) removePeerFromAllReactors(peer Peer, reason interface{}) {
	_ = "STUB: not implemented"
	return
}

// removePeerFromReactor removes the peer from the specified reactor
func (sw *Switch) removePeerFromReactor(peer Peer, reactorName string) {
	_ = "STUB: not implemented"
	return
}

// doRemovePeer removes the specified peer from the given reactor and its corresponding peer set, logging any issues.
func (sw *Switch) doRemovePeer(peer Peer, reactor Reactor, reason interface{}) {
	_ = "STUB: not implemented"
	return
}

// countActivePeerConnections returns the number of reactors that have this peer
func (sw *Switch) countActivePeerConnections(peer Peer) int { _ = "STUB: not implemented"; return 0 }
