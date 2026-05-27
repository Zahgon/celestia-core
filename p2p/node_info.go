package p2p

import (
	cmtbytes "github.com/cometbft/cometbft/libs/bytes"
	tmp2p "github.com/cometbft/cometbft/proto/tendermint/p2p"
	"github.com/cometbft/cometbft/version"
)

const (
	maxNodeInfoSize = 10240 // 10KB
	maxNumChannels  = 16    // plenty of room for upgrades, for now
)

// Max size of the NodeInfo struct
func MaxNodeInfoSize() int { _ = "STUB: not implemented"; return 0 }

//-------------------------------------------------------------

// NodeInfo exposes basic info of a node
// and determines if we're compatible.
type NodeInfo interface {
	ID() ID
	nodeInfoAddress
	nodeInfoTransport
}

type nodeInfoAddress interface {
	NetAddress() (*NetAddress, error)
}

// nodeInfoTransport validates a nodeInfo and checks
// our compatibility with it. It's for use in the handshake.
type nodeInfoTransport interface {
	Validate() error
	CompatibleWith(other NodeInfo) error
}

//-------------------------------------------------------------

// ProtocolVersion contains the protocol versions for the software.
type ProtocolVersion struct {
	P2P   uint64 `json:"p2p"`
	Block uint64 `json:"block"`
	App   uint64 `json:"app"`
}

// defaultProtocolVersion populates the Block and P2P versions using
// the global values, but not the App.
var defaultProtocolVersion = NewProtocolVersion(
	version.P2PProtocol,
	version.BlockProtocol,
	0,
)

// NewProtocolVersion returns a fully populated ProtocolVersion.
func NewProtocolVersion(p2p, block, app uint64) ProtocolVersion {
	_ = "STUB: not implemented"
	return *new(ProtocolVersion)
}

//-------------------------------------------------------------

// Assert DefaultNodeInfo satisfies NodeInfo
var _ NodeInfo = DefaultNodeInfo{}

// DefaultNodeInfo is the basic node information exchanged
// between two peers during the CometBFT P2P handshake.
type DefaultNodeInfo struct {
	ProtocolVersion ProtocolVersion `json:"protocol_version"`

	// Authenticate
	// TODO: replace with NetAddress
	DefaultNodeID ID     `json:"id"`          // authenticated identifier
	ListenAddr    string `json:"listen_addr"` // accepting incoming

	// Check compatibility.
	// Channels are HexBytes so easier to read as JSON
	Network  string            `json:"network"`  // network/chain ID
	Version  string            `json:"version"`  // major.minor.revision
	Channels cmtbytes.HexBytes `json:"channels"` // channels this node knows about

	// ASCIIText fields
	Moniker string               `json:"moniker"` // arbitrary moniker
	Other   DefaultNodeInfoOther `json:"other"`   // other application specific data
}

// DefaultNodeInfoOther is the misc. applcation specific data
type DefaultNodeInfoOther struct {
	TxIndex    string `json:"tx_index"`
	RPCAddress string `json:"rpc_address"`
}

// ID returns the node's peer ID.
func (info DefaultNodeInfo) ID() ID {
	_ = "STUB: not implemented"
	return *

	// Validate checks the self-reported DefaultNodeInfo is safe.
	// It returns an error if there
	// are too many Channels, if there are any duplicate Channels,
	// if the ListenAddr is malformed, or if the ListenAddr is a host name
	// that can not be resolved to some IP.
	// TODO: constraints for Moniker/Other? Or is that for the UI ?
	// JAE: It needs to be done on the client, but to prevent ambiguous
	// unicode characters, maybe it's worth sanitizing it here.
	// In the future we might want to validate these, once we have a
	// name-resolution system up.
	// International clients could then use punycode (or we could use
	// url-encoding), and we just need to be careful with how we handle that in our
	// clients. (e.g. off by default).
	new(ID)
}

func (info DefaultNodeInfo) Validate() error {
	_ = "STUB: not implemented"

	// ID is already validated.
	return nil
}

// Validate ListenAddr.

// Network is validated in CompatibleWith.

// Validate Version

// Validate Channels - ensure max and check for duplicates.

// Validate Moniker.

// Validate Other.

// XXX: Should we be more strict about address formats?

// CompatibleWith checks if two DefaultNodeInfo are compatible with eachother.
// CONTRACT: two nodes are compatible if the Block version and network match
// and they have at least one channel in common.
func (info DefaultNodeInfo) CompatibleWith(otherInfo NodeInfo) error {
	_ = "STUB: not implemented"
	return nil
}

// nodes must be on the same network

// if we have no channels, we're just testing

// for each of our channels, check if they have it

// only need one

// NetAddress returns a NetAddress derived from the DefaultNodeInfo -
// it includes the authenticated peer ID and the self-reported
// ListenAddr. Note that the ListenAddr is not authenticated and
// may not match that address actually dialed if its an outbound peer.
func (info DefaultNodeInfo) NetAddress() (*NetAddress, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (info DefaultNodeInfo) HasChannel(chID byte) bool { _ = "STUB: not implemented"; return false }

func (info DefaultNodeInfo) ToProto() *tmp2p.DefaultNodeInfo { _ = "STUB: not implemented"; return nil }

func DefaultNodeInfoFromToProto(pb *tmp2p.DefaultNodeInfo) (DefaultNodeInfo, error) {
	_ = "STUB: not implemented"
	return *new(DefaultNodeInfo), nil
}
