package p2p

import (
	"net"

	"github.com/cometbft/cometbft/crypto"

	"github.com/cometbft/cometbft/config"
)

const testCh = 0x01

//------------------------------------------------

type mockNodeInfo struct {
	addr *NetAddress
}

func (ni mockNodeInfo) ID() ID { _ = "STUB: not implemented"; return *new(ID) }
func (ni mockNodeInfo) NetAddress() (*NetAddress, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
func (ni mockNodeInfo) Validate() error               { _ = "STUB: not implemented"; return nil }
func (ni mockNodeInfo) CompatibleWith(NodeInfo) error { _ = "STUB: not implemented"; return nil }

func AddPeerToSwitchPeerSet(sw *Switch, peer Peer) {
	_ = "STUB: not implemented"
	//nolint:errcheck // ignore error
	return
}

func CreateRandomPeer(outbound bool) Peer { _ = "STUB: not implemented"; return *new(Peer) }

func CreateRoutableAddr() (addr string, netAddr *NetAddress) {
	_ = "STUB: not implemented"
	return "", nil
}

//------------------------------------------------------------------
// Connects switches via arbitrary net.Conn. Used for testing.

const TestHost = "localhost"

// MakeConnectedSwitches returns n switches, connected according to the connect func.
// If connect==Connect2Switches, the switches will be fully connected.
// initSwitch defines how the i'th switch should be initialized (ie. with what reactors).
// NOTE: panics if any switch fails to start.
func MakeConnectedSwitches(
	cfg *config.P2PConfig,
	n int,
	initSwitch func(int, *Switch) *Switch,
	connect func([]*Switch, int, int),
) []*Switch {
	_ = "STUB: not implemented"
	return nil
}

// Connect2Switches will connect switches i and j via net.Pipe().
// Blocks until a connection is established.
// NOTE: caller ensures i and j are within bounds.
func Connect2Switches(switches []*Switch, i, j int) { _ = "STUB: not implemented"; return }

func (sw *Switch) addPeerWithConnection(conn net.Conn) error { _ = "STUB: not implemented"; return nil }

// StartSwitches calls sw.Start() for each given switch.
// It returns the first encountered error.
func StartSwitches(switches []*Switch) error { _ = "STUB: not implemented"; return nil }

// start switch and reactors

func MakeSwitch(
	cfg *config.P2PConfig,
	i int,
	initSwitch func(int, *Switch) *Switch,
	opts ...SwitchOption,
) *Switch {
	_ = "STUB: not implemented"
	return nil
}

// TODO: let the config be passed in?

// TODO: We need to setup reactors ahead of time so the NodeInfo is properly
// populated and we don't have to do those awkward overrides and setters.

func testInboundPeerConn(
	conn net.Conn,
	config *config.P2PConfig,
	ourNodePrivKey crypto.PrivKey,
) (peerConn, error) {
	_ = "STUB: not implemented"
	return *new(peerConn), nil
}

func testPeerConn(
	rawConn net.Conn,
	cfg *config.P2PConfig,
	outbound, persistent bool,
	ourNodePrivKey crypto.PrivKey,
	socketAddr *NetAddress,
) (pc peerConn, err error) {
	_ = "STUB: not implemented"

	// Fuzz connection
	return *new(peerConn), nil
}

// so we have time to do peer handshakes and get set up

// Encrypt connection

// Only the information we already have

//----------------------------------------------------------------
// rand node info

func testNodeInfo(id ID, name string) NodeInfo { _ = "STUB: not implemented"; return *new(NodeInfo) }

func testNodeInfoWithNetwork(id ID, name, network string) NodeInfo {
	_ = "STUB: not implemented"
	return *new(NodeInfo)
}

func getFreePort() int { _ = "STUB: not implemented"; return 0 }

type AddrBookMock struct {
	Addrs        map[string]struct{}
	OurAddrs     map[string]struct{}
	PrivateAddrs map[string]struct{}
}

var _ AddrBook = (*AddrBookMock)(nil)

func (book *AddrBookMock) AddAddress(addr *NetAddress, _ *NetAddress) error {
	_ = "STUB: not implemented"
	return nil
}

func (book *AddrBookMock) AddOurAddress(addr *NetAddress) { _ = "STUB: not implemented"; return }
func (book *AddrBookMock) OurAddress(addr *NetAddress) bool {
	_ = "STUB: not implemented"
	return false
}

func (book *AddrBookMock) MarkGood(ID) { _ = "STUB: not implemented"; return }
func (book *AddrBookMock) HasAddress(addr *NetAddress) bool {
	_ = "STUB: not implemented"
	return false
}

func (book *AddrBookMock) RemoveAddress(addr *NetAddress) { _ = "STUB: not implemented"; return }

func (book *AddrBookMock) Save()                        { _ = "STUB: not implemented"; return }
func (book *AddrBookMock) AddPrivateIDs(addrs []string) { _ = "STUB: not implemented"; return }
