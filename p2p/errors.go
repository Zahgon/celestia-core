package p2p

import (
	"net"
)

// ErrFilterTimeout indicates that a filter operation timed out.
type ErrFilterTimeout struct{}

func (e ErrFilterTimeout) Error() string { _ = "STUB: not implemented"; return "" }

// ErrRejected indicates that a Peer was rejected carrying additional
// information as to the reason.
type ErrRejected struct {
	addr               NetAddress
	conn               net.Conn
	err                error
	id                 ID
	isAuthFailure      bool
	isDuplicate        bool
	isFiltered         bool
	isIncompatible     bool
	isNodeInfoInvalid  bool
	isSelf             bool
	localNodeID        string
	remoteNodeID       string
	localAddr          string
	remoteAddr         string
	handshakeStage     string
	traceID            string
	chainID            string
	peerChainID        string
	malformedHandshake bool
}

// Addr returns the NetAddress for the rejected Peer.
func (e ErrRejected) Addr() NetAddress { _ = "STUB: not implemented"; return *new(NetAddress) }

func (e ErrRejected) Error() string { _ = "STUB: not implemented"; return "" }

// IsAuthFailure when Peer authentication was unsuccessful.
func (e ErrRejected) IsAuthFailure() bool { _ = "STUB: not implemented"; return false }

// IsDuplicate when Peer ID or IP are present already.
func (e ErrRejected) IsDuplicate() bool { _ = "STUB: not implemented"; return false }

// DuplicatePeerID returns the peer ID and true when the rejection was caused
// by a duplicate peer ID. Returns "", false for duplicate-IP rejections or
// non-duplicate errors. Callers that need to know *why* a peer was deemed a
// duplicate (e.g. to decide whether to penalize the dialed address) should
// use this in preference to IsDuplicate.
func (e ErrRejected) DuplicatePeerID() (ID, bool) {
	_ = "STUB: not implemented"
	return *new(ID), false
}

// NewErrRejectedDuplicateID builds a duplicate-by-peer-ID rejection. The
// returned error reports DuplicatePeerID() == (id, true). Exposed primarily
// so tests outside the p2p package can construct rejections without
// reaching into unexported fields.
func NewErrRejectedDuplicateID(id ID) ErrRejected {
	_ = "STUB: not implemented"
	return *new(ErrRejected)
}

// NewErrRejectedDuplicateIP builds a duplicate-by-IP rejection (no peer ID
// recorded). The returned error reports IsDuplicate() == true but
// DuplicatePeerID() == ("", false). Mirrors the rejection produced by the
// duplicate-IP filter in the transport layer.
func NewErrRejectedDuplicateIP() ErrRejected { _ = "STUB: not implemented"; return *new(ErrRejected) }

// IsFiltered when Peer ID or IP was filtered.
func (e ErrRejected) IsFiltered() bool {
	_ = "STUB: not implemented"

	// IsIncompatible when Peer NodeInfo is not compatible with our own.
	return false
}

func (e ErrRejected) IsIncompatible() bool { _ = "STUB: not implemented"; return false }

// IsNodeInfoInvalid when the sent NodeInfo is not valid.
func (e ErrRejected) IsNodeInfoInvalid() bool { _ = "STUB: not implemented"; return false }

// IsSelf when Peer is our own node.
func (e ErrRejected) IsSelf() bool {
	_ = "STUB: not implemented"

	// ErrSwitchDuplicatePeerID to be raised when a peer is connecting with a known
	// ID.
	return false
}

type ErrSwitchDuplicatePeerID struct {
	ID ID
}

func (e ErrSwitchDuplicatePeerID) Error() string { _ = "STUB: not implemented"; return "" }

// ErrSwitchDuplicatePeerIP to be raised whena a peer is connecting with a known
// IP.
type ErrSwitchDuplicatePeerIP struct {
	IP net.IP
}

func (e ErrSwitchDuplicatePeerIP) Error() string { _ = "STUB: not implemented"; return "" }

// ErrSwitchConnectToSelf to be raised when trying to connect to itself.
type ErrSwitchConnectToSelf struct {
	Addr *NetAddress
}

func (e ErrSwitchConnectToSelf) Error() string { _ = "STUB: not implemented"; return "" }

type ErrSwitchAuthenticationFailure struct {
	Dialed *NetAddress
	Got    ID
}

func (e ErrSwitchAuthenticationFailure) Error() string { _ = "STUB: not implemented"; return "" }

// ErrTransportClosed is raised when the Transport has been closed.
type ErrTransportClosed struct{}

func (e ErrTransportClosed) Error() string { _ = "STUB: not implemented"; return "" }

// ErrPeerRemoval is raised when attempting to remove a peer results in an error.
type ErrPeerRemoval struct{}

func (e ErrPeerRemoval) Error() string { _ = "STUB: not implemented"; return "" }

//-------------------------------------------------------------------

type ErrNetAddressNoID struct {
	Addr string
}

func (e ErrNetAddressNoID) Error() string { _ = "STUB: not implemented"; return "" }

type ErrNetAddressInvalid struct {
	Addr string
	Err  error
}

func (e ErrNetAddressInvalid) Error() string { _ = "STUB: not implemented"; return "" }

type ErrNetAddressLookup struct {
	Addr string
	Err  error
}

func (e ErrNetAddressLookup) Error() string { _ = "STUB: not implemented"; return "" }

// ErrCurrentlyDialingOrExistingAddress indicates that we're currently
// dialing this address or it belongs to an existing peer.
type ErrCurrentlyDialingOrExistingAddress struct {
	Addr string
}

func (e ErrCurrentlyDialingOrExistingAddress) Error() string { _ = "STUB: not implemented"; return "" }
