package pex

import (
	"errors"

	"github.com/cometbft/cometbft/p2p"
)

type ErrAddrBookNonRoutable struct {
	Addr *p2p.NetAddress
}

func (err ErrAddrBookNonRoutable) Error() string { _ = "STUB: not implemented"; return "" }

type errAddrBookOldAddressNewBucket struct {
	Addr     *p2p.NetAddress
	BucketID int
}

func (err errAddrBookOldAddressNewBucket) Error() string { _ = "STUB: not implemented"; return "" }

type ErrAddrBookSelf struct {
	Addr *p2p.NetAddress
}

func (err ErrAddrBookSelf) Error() string { _ = "STUB: not implemented"; return "" }

type ErrAddrBookPrivate struct {
	Addr *p2p.NetAddress
}

func (err ErrAddrBookPrivate) Error() string { _ = "STUB: not implemented"; return "" }

func (err ErrAddrBookPrivate) PrivateAddr() bool { _ = "STUB: not implemented"; return false }

type ErrAddrBookPrivateSrc struct {
	Src *p2p.NetAddress
}

func (err ErrAddrBookPrivateSrc) Error() string { _ = "STUB: not implemented"; return "" }

func (err ErrAddrBookPrivateSrc) PrivateAddr() bool { _ = "STUB: not implemented"; return false }

type ErrAddrBookNilAddr struct {
	Addr *p2p.NetAddress
	Src  *p2p.NetAddress
}

func (err ErrAddrBookNilAddr) Error() string { _ = "STUB: not implemented"; return "" }

type ErrAddrBookInvalidAddr struct {
	Addr    *p2p.NetAddress
	AddrErr error
}

func (err ErrAddrBookInvalidAddr) Error() string { _ = "STUB: not implemented"; return "" }

// ErrAddressBanned is thrown when the address has been banned and therefore cannot be used
type ErrAddressBanned struct {
	Addr *p2p.NetAddress
}

func (err ErrAddressBanned) Error() string { _ = "STUB: not implemented"; return "" }

// ErrUnsolicitedList is thrown when a peer provides a list of addresses that have not been asked for.
var ErrUnsolicitedList = errors.New("unsolicited pexAddrsMessage")
