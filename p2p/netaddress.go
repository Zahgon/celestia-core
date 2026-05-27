// Modified for CometBFT
// Originally Copyright (c) 2013-2014 Conformal Systems LLC.
// https://github.com/conformal/btcd/blob/master/LICENSE

package p2p

import (
	"net"
	"time"

	tmp2p "github.com/cometbft/cometbft/proto/tendermint/p2p"
)

// EmptyNetAddress defines the string representation of an empty NetAddress
const EmptyNetAddress = "<nil-NetAddress>"

// NetAddress defines information about a peer on the network
// including its ID, IP address, and port.
type NetAddress struct {
	ID   ID     `json:"id"`
	IP   net.IP `json:"ip"`
	Port uint16 `json:"port"`
}

// IDAddressString returns id@hostPort. It strips the leading
// protocol from protocolHostPort if it exists.
func IDAddressString(id ID, protocolHostPort string) string { _ = "STUB: not implemented"; return "" }

// NewNetAddress returns a new NetAddress using the provided TCP
// address. When testing, other net.Addr (except TCP) will result in
// using 0.0.0.0:0. When normal run, other net.Addr (except TCP) will
// panic. Panics if ID is invalid.
// TODO: socks proxies?
func NewNetAddress(id ID, addr net.Addr) *NetAddress { _ = "STUB: not implemented"; return nil }

// normal run

// in testing

// NewNetAddressString returns a new NetAddress using the provided address in
// the form of "ID@IP:Port".
// Also resolves the host if host is not an IP.
// Errors are of type ErrNetAddressXxx where Xxx is in (NoID, Invalid, Lookup)
func NewNetAddressString(addr string) (*NetAddress, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// get ID

// get host and port

// NewNetAddressStrings returns an array of NetAddress'es build using
// the provided strings.
func NewNetAddressStrings(addrs []string) ([]*NetAddress, []error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewNetAddressIPPort returns a new NetAddress using the provided IP
// and port number.
func NewNetAddressIPPort(ip net.IP, port uint16) *NetAddress { _ = "STUB: not implemented"; return nil }

// NetAddressFromProto converts a Protobuf NetAddress into a native struct.
func NetAddressFromProto(pb tmp2p.NetAddress) (*NetAddress, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NetAddressesFromProto converts a slice of Protobuf NetAddresses into a native slice.
func NetAddressesFromProto(pbs []tmp2p.NetAddress) ([]*NetAddress, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NetAddressesToProto converts a slice of NetAddresses into a Protobuf slice.
func NetAddressesToProto(nas []*NetAddress) []tmp2p.NetAddress {
	_ = "STUB: not implemented"
	return nil
}

// ToProto converts a NetAddress to Protobuf.
func (na *NetAddress) ToProto() tmp2p.NetAddress {
	_ = "STUB: not implemented"
	return *new(tmp2p.NetAddress)
}

// Equals reports whether na and other are the same addresses,
// including their ID, IP, and Port.
func (na *NetAddress) Equals(other interface{}) bool { _ = "STUB: not implemented"; return false }

// Same returns true is na has the same non-empty ID or DialString as other.
func (na *NetAddress) Same(other interface{}) bool { _ = "STUB: not implemented"; return false }

// String representation: <ID>@<IP>:<PORT>
func (na *NetAddress) String() string { _ = "STUB: not implemented"; return "" }

func (na *NetAddress) DialString() string { _ = "STUB: not implemented"; return "" }

// Dial calls net.Dial on the address.
func (na *NetAddress) Dial() (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

// DialTimeout calls net.DialTimeout on the address.
func (na *NetAddress) DialTimeout(timeout time.Duration) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

// Routable returns true if the address is routable.
func (na *NetAddress) Routable() bool { _ = "STUB: not implemented"; return false }

// TODO(oga) bitcoind doesn't include RFC3849 here, but should we?
//nolint:staticcheck

// For IPv4 these are either a 0 or all bits set address. For IPv6 a zero
// address or one that matches the RFC3849 documentation address format.
func (na *NetAddress) Valid() error { _ = "STUB: not implemented"; return nil }

// HasID returns true if the address has an ID.
// NOTE: It does not check whether the ID is valid or not.
func (na *NetAddress) HasID() bool { _ = "STUB: not implemented"; return false }

// Local returns true if it is a local address.
func (na *NetAddress) Local() bool { _ = "STUB: not implemented"; return false }

// ReachabilityTo checks whenever o can be reached from na.
func (na *NetAddress) ReachabilityTo(o *NetAddress) int { _ = "STUB: not implemented"; return 0 }

// ipv6

/* ipv6 */

// Is our v6 is tunneled?

// only prioritize ipv6 if we aren't tunneling it.

// RFC1918: IPv4 Private networks (10.0.0.0/8, 192.168.0.0/16, 172.16.0.0/12)
// RFC3849: IPv6 Documentation address  (2001:0DB8::/32)
// RFC3927: IPv4 Autoconfig (169.254.0.0/16)
// RFC3964: IPv6 6to4 (2002::/16)
// RFC4193: IPv6 unique local (FC00::/7)
// RFC4380: IPv6 Teredo tunneling (2001::/32)
// RFC4843: IPv6 ORCHID: (2001:10::/28)
// RFC4862: IPv6 Autoconfig (FE80::/64)
// RFC6052: IPv6 well known prefix (64:FF9B::/96)
// RFC6145: IPv6 IPv4 translated address ::FFFF:0:0:0/96
var rfc1918_10 = net.IPNet{IP: net.ParseIP("10.0.0.0"), Mask: net.CIDRMask(8, 32)}
var rfc1918_192 = net.IPNet{IP: net.ParseIP("192.168.0.0"), Mask: net.CIDRMask(16, 32)}
var rfc1918_172 = net.IPNet{IP: net.ParseIP("172.16.0.0"), Mask: net.CIDRMask(12, 32)}
var rfc3849 = net.IPNet{IP: net.ParseIP("2001:0DB8::"), Mask: net.CIDRMask(32, 128)}
var rfc3927 = net.IPNet{IP: net.ParseIP("169.254.0.0"), Mask: net.CIDRMask(16, 32)}
var rfc3964 = net.IPNet{IP: net.ParseIP("2002::"), Mask: net.CIDRMask(16, 128)}
var rfc4193 = net.IPNet{IP: net.ParseIP("FC00::"), Mask: net.CIDRMask(7, 128)}
var rfc4380 = net.IPNet{IP: net.ParseIP("2001::"), Mask: net.CIDRMask(32, 128)}
var rfc4843 = net.IPNet{IP: net.ParseIP("2001:10::"), Mask: net.CIDRMask(28, 128)}
var rfc4862 = net.IPNet{IP: net.ParseIP("FE80::"), Mask: net.CIDRMask(64, 128)}
var rfc6052 = net.IPNet{IP: net.ParseIP("64:FF9B::"), Mask: net.CIDRMask(96, 128)}
var rfc6145 = net.IPNet{IP: net.ParseIP("::FFFF:0:0:0"), Mask: net.CIDRMask(96, 128)}
var zero4 = net.IPNet{IP: net.ParseIP("0.0.0.0"), Mask: net.CIDRMask(8, 32)}
var (
	// onionCatNet defines the IPv6 address block used to support Tor.
	// bitcoind encodes a .onion address as a 16 byte number by decoding the
	// address prior to the .onion (i.e. the key hash) base32 into a ten
	// byte number. It then stores the first 6 bytes of the address as
	// 0xfd, 0x87, 0xd8, 0x7e, 0xeb, 0x43.
	//
	// This is the same range used by OnionCat, which is part part of the
	// RFC4193 unique local IPv6 range.
	//
	// In summary the format is:
	// { magic 6 bytes, 10 bytes base32 decode of key hash }
	onionCatNet = ipNet("fd87:d87e:eb43::", 48, 128)
)

// ipNet returns a net.IPNet struct given the passed IP address string, number
// of one bits to include at the start of the mask, and the total number of bits
// for the mask.
func ipNet(ip string, ones, bits int) net.IPNet { _ = "STUB: not implemented"; return *new(net.IPNet) }

func (na *NetAddress) RFC1918() bool { _ = "STUB: not implemented"; return false }

func (na *NetAddress) RFC3849() bool     { _ = "STUB: not implemented"; return false }
func (na *NetAddress) RFC3927() bool     { _ = "STUB: not implemented"; return false }
func (na *NetAddress) RFC3964() bool     { _ = "STUB: not implemented"; return false }
func (na *NetAddress) RFC4193() bool     { _ = "STUB: not implemented"; return false }
func (na *NetAddress) RFC4380() bool     { _ = "STUB: not implemented"; return false }
func (na *NetAddress) RFC4843() bool     { _ = "STUB: not implemented"; return false }
func (na *NetAddress) RFC4862() bool     { _ = "STUB: not implemented"; return false }
func (na *NetAddress) RFC6052() bool     { _ = "STUB: not implemented"; return false }
func (na *NetAddress) RFC6145() bool     { _ = "STUB: not implemented"; return false }
func (na *NetAddress) OnionCatTor() bool { _ = "STUB: not implemented"; return false }

func removeProtocolIfDefined(addr string) string { _ = "STUB: not implemented"; return "" }

func validateID(id ID) error { _ = "STUB: not implemented"; return nil }
