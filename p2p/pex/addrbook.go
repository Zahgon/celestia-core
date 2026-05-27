// Modified for CometBFT
// Originally Copyright (c) 2013-2014 Conformal Systems LLC.
// https://github.com/conformal/btcd/blob/master/LICENSE

package pex

import (
	"hash"
	"sync"
	"time"

	"github.com/cometbft/cometbft/crypto"
	cmtrand "github.com/cometbft/cometbft/libs/rand"
	"github.com/cometbft/cometbft/libs/service"
	cmtsync "github.com/cometbft/cometbft/libs/sync"
	"github.com/cometbft/cometbft/p2p"
)

const (
	bucketTypeNew = 0x01
	bucketTypeOld = 0x02
)

// AddrBook is an address book used for tracking peers
// so we can gossip about them to others and select
// peers to dial.
// TODO: break this up?
type AddrBook interface {
	service.Service

	// Add our own addresses so we don't later add ourselves
	AddOurAddress(*p2p.NetAddress)
	// Check if it is our address
	OurAddress(*p2p.NetAddress) bool

	AddPrivateIDs([]string)

	// Add and remove an address
	AddAddress(addr *p2p.NetAddress, src *p2p.NetAddress) error
	RemoveAddress(*p2p.NetAddress)

	// Check if the address is in the book
	HasAddress(*p2p.NetAddress) bool

	// Do we need more peers?
	NeedMoreAddrs() bool
	// Is Address Book Empty? Answer should not depend on being in your own
	// address book, or private peers
	Empty() bool

	// Mark address
	MarkGood(p2p.ID)
	MarkAttempt(*p2p.NetAddress)
	MarkBad(*p2p.NetAddress, time.Duration) // Move peer to bad peers list
	// Add bad peers back to addrBook
	ReinstateBadPeers()

	IsGood(*p2p.NetAddress) bool
	IsBanned(*p2p.NetAddress) bool

	// Send a selection of addresses to peers
	GetSelection() []*p2p.NetAddress
	// Send a selection of addresses with bias
	GetSelectionWithBias(biasTowardsNewAddrs int) []*p2p.NetAddress

	Size() int

	// Persist to disk
	Save()
}

var _ AddrBook = (*addrBook)(nil)

// addrBook - concurrency safe peer address manager.
// Implements AddrBook.
type addrBook struct {
	service.BaseService

	// accessed concurrently
	mtx        cmtsync.Mutex
	rand       *cmtrand.Rand
	ourAddrs   map[string]struct{}
	privateIDs map[p2p.ID]struct{}
	addrLookup map[p2p.ID]*knownAddress // new & old
	badPeers   map[p2p.ID]*knownAddress // blacklisted peers
	bucketsOld []map[string]*knownAddress
	bucketsNew []map[string]*knownAddress
	nOld       int
	nNew       int

	// immutable after creation
	filePath          string
	key               string // random prefix for bucket placement
	routabilityStrict bool
	hasher            hash.Hash64

	wg sync.WaitGroup
}

func mustNewHasher() hash.Hash64 { _ = "STUB: not implemented"; return *new(hash.Hash64) }

// NewAddrBook creates a new address book.
// Use Start to begin processing asynchronous address updates.
func NewAddrBook(filePath string, routabilityStrict bool) AddrBook {
	_ = "STUB: not implemented"
	return *new(AddrBook)
}

// Initialize the buckets.
// When modifying this, don't forget to update loadFromFile()
func (a *addrBook) init() {
	a.key = crypto.CRandHex(24) // 24/2 * 8 = 96 bits
	// New addr buckets
	a.bucketsNew = make([]map[string]*knownAddress, newBucketCount)
	for i := range a.bucketsNew {
		a.bucketsNew[i] = make(map[string]*knownAddress)
	}
	// Old addr buckets
	a.bucketsOld = make([]map[string]*knownAddress, oldBucketCount)
	for i := range a.bucketsOld {
		a.bucketsOld[i] = make(map[string]*knownAddress)
	}
	a.hasher = mustNewHasher()
}

// OnStart implements Service.
func (a *addrBook) OnStart() error { _ = "STUB: not implemented"; return nil }

// wg.Add to ensure that any invocation of .Wait()
// later on will wait for saveRoutine to terminate.

// OnStop implements Service.
func (a *addrBook) OnStop() { _ = "STUB: not implemented"; return }

func (a *addrBook) Wait() { _ = "STUB: not implemented"; return }

func (a *addrBook) FilePath() string {
	_ = "STUB: not implemented"

	// -------------------------------------------------------
	return ""
}

// AddOurAddress one of our addresses.
func (a *addrBook) AddOurAddress(addr *p2p.NetAddress) { _ = "STUB: not implemented"; return }

// OurAddress returns true if it is our address.
func (a *addrBook) OurAddress(addr *p2p.NetAddress) bool { _ = "STUB: not implemented"; return false }

func (a *addrBook) AddPrivateIDs(ids []string) { _ = "STUB: not implemented"; return }

// AddAddress implements AddrBook
// Add address to a "new" bucket. If it's already in one, only add it probabilistically.
// Returns error if the addr is non-routable. Does not add self.
// NOTE: addr must not be nil
func (a *addrBook) AddAddress(addr *p2p.NetAddress, src *p2p.NetAddress) error {
	_ = "STUB: not implemented"
	return nil
}

// RemoveAddress implements AddrBook - removes the address from the book.
func (a *addrBook) RemoveAddress(addr *p2p.NetAddress) { _ = "STUB: not implemented"; return }

// IsGood returns true if peer was ever marked as good and haven't
// done anything wrong since then.
func (a *addrBook) IsGood(addr *p2p.NetAddress) bool { _ = "STUB: not implemented"; return false }

// IsBanned returns true if the peer is currently banned
func (a *addrBook) IsBanned(addr *p2p.NetAddress) bool { _ = "STUB: not implemented"; return false }

// HasAddress returns true if the address is in the book.
func (a *addrBook) HasAddress(addr *p2p.NetAddress) bool { _ = "STUB: not implemented"; return false }

// NeedMoreAddrs implements AddrBook - returns true if there are not have enough addresses in the book.
func (a *addrBook) NeedMoreAddrs() bool { _ = "STUB: not implemented"; return false }

// Empty implements AddrBook - returns true if there are no addresses in the address book.
// Does not count the peer appearing in its own address book, or private peers.
func (a *addrBook) Empty() bool { _ = "STUB: not implemented"; return false }

// MarkGood implements AddrBook - it marks the peer as good and
// moves it into an "old" bucket.
func (a *addrBook) MarkGood(id p2p.ID) { _ = "STUB: not implemented"; return }

// MarkAttempt implements AddrBook - it marks that an attempt was made to connect to the address.
func (a *addrBook) MarkAttempt(addr *p2p.NetAddress) { _ = "STUB: not implemented"; return }

// MarkBad implements AddrBook. Kicks address out from book, places
// the address in the badPeers pool.
func (a *addrBook) MarkBad(addr *p2p.NetAddress, banTime time.Duration) {
	_ = "STUB: not implemented"
	return
}

// ReinstateBadPeers removes bad peers from ban list and places them into a new
// bucket.
func (a *addrBook) ReinstateBadPeers() { _ = "STUB: not implemented"; return }

// GetSelection implements AddrBook.
// It returns all addresses (old & new). Suitable for peer-exchange protocols.
// Must never return a nil address.
func (a *addrBook) GetSelection() []*p2p.NetAddress { _ = "STUB: not implemented"; return nil }

// Get all addresses from addrLookup

// Fisher-Yates shuffle the addresses

func percentageOfNum(p, n int) int { _ = "STUB: not implemented"; return 0 }

// GetSelectionWithBias implements AddrBook.
// It randomly selects some addresses (old & new). Suitable for peer-exchange protocols.
// Must never return a nil address.
//
// Each address is picked randomly from an old or new bucket according to the
// biasTowardsNewAddrs argument, which must be between [0, 100] (or else is truncated to
// that range) and determines how biased we are to pick an address from a new
// bucket.
func (a *addrBook) GetSelectionWithBias(biasTowardsNewAddrs int) []*p2p.NetAddress {
	_ = "STUB: not implemented"
	return nil
}

// number of new addresses that, if possible, should be in the beginning of the selection
// if there are no enough old addrs, will choose new addr instead.

//------------------------------------------------

// Size returns the number of addresses in the book.
func (a *addrBook) Size() int { _ = "STUB: not implemented"; return 0 }

func (a *addrBook) size() int { _ = "STUB: not implemented"; return 0 }

//----------------------------------------------------------

// Save persists the address book to disk.
func (a *addrBook) Save() { _ = "STUB: not implemented"; return }

// thread safe

func (a *addrBook) saveRoutine() { _ = "STUB: not implemented"; return }

//----------------------------------------------------------

func (a *addrBook) getBucket(bucketType byte, bucketIdx int) map[string]*knownAddress {
	_ = "STUB: not implemented"
	return nil
}

// Adds ka to new bucket. Returns false if it couldn't do it cuz buckets full.
// NOTE: currently it always returns true.
func (a *addrBook) addToNewBucket(ka *knownAddress, bucketIdx int) error {
	_ = "STUB: not implemented"
	// Consistency check to ensure we don't add an already known address
	return nil
}

// Already exists?

// Enforce max addresses.

// Add to bucket.

// increment nNew if the peer doesnt already exist in a bucket

// Add it to addrLookup

// Adds ka to old bucket. Returns false if it couldn't do it cuz buckets full.
func (a *addrBook) addToOldBucket(ka *knownAddress, bucketIdx int) bool {
	_ = "STUB: not implemented"
	// Sanity check
	return false
}

// Already exists?

// Enforce max addresses.

// Add to bucket.

// Ensure in addrLookup

func (a *addrBook) removeFromBucket(ka *knownAddress, bucketType byte, bucketIdx int) {
	_ = "STUB: not implemented"
	return
}

func (a *addrBook) removeFromAllBuckets(ka *knownAddress) { _ = "STUB: not implemented"; return }

//----------------------------------------------------------

func (a *addrBook) pickOldest(bucketType byte, bucketIdx int) *knownAddress {
	_ = "STUB: not implemented"
	return nil
}

// adds the address to a "new" bucket. if its already in one,
// it only adds it probabilistically
func (a *addrBook) addAddress(addr, src *p2p.NetAddress) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: we should track ourAddrs by ID and by IP:PORT and refuse both.

// If its already old and the address ID's are the same, ignore it.
// Thereby avoiding issues with a node on the network attempting to change
// the IP of a known node ID. (Which could yield an eclipse attack on the node)

// Already in max new buckets.

// The more entries we have, the less likely we are to add more.

func (a *addrBook) randomPickAddresses(bucketType byte, num int) []*p2p.NetAddress {
	_ = "STUB: not implemented"
	return nil
}

// Make space in the new buckets by expiring the really bad entries.
// If no bad entries are available we remove the oldest.
func (a *addrBook) expireNew(bucketIdx int) { _ = "STUB: not implemented"; return }

// If an entry is bad, throw it away

// If we haven't thrown out a bad entry, throw out the oldest entry

// Promotes an address from new to old. If the destination bucket is full,
// demote the oldest one to a "new" bucket.
// TODO: Demote more probabilistically?
func (a *addrBook) moveToOld(ka *knownAddress) error {
	_ = "STUB: not implemented"
	// Sanity check
	return nil
}

// Remove from all (new) buckets.

// It's officially old now.

// Try to add it to its oldBucket destination.

// No room; move the oldest to a new bucket

// Finally, add our ka to old bucket again.

func (a *addrBook) removeAddress(addr *p2p.NetAddress) { _ = "STUB: not implemented"; return }

func (a *addrBook) addBadPeer(addr *p2p.NetAddress, banTime time.Duration) bool {
	_ = "STUB: not implemented"
	// check it exists in addrbook
	return false
}

// check address is not already there

// add to bad peer list

//---------------------------------------------------------------------
// calculate bucket placements

// hash(key + sourcegroup + int64(hash(key + group + sourcegroup)) % bucket_per_group) % num_new_buckets
func (a *addrBook) calcNewBucket(addr, src *p2p.NetAddress) (int, error) {
	_ = "STUB: not implemented"
	//nolint:prealloc
	return 0, nil
}

//nolint:prealloc

// hash(key + group + int64(hash(key + addr)) % buckets_per_group) % num_old_buckets
func (a *addrBook) calcOldBucket(addr *p2p.NetAddress) (int, error) {
	_ = "STUB: not implemented"
	//nolint:prealloc
	return 0, nil
}

//nolint:prealloc

// Return a string representing the network group of this address.
// This is the /16 for IPv4 (e.g. 1.2.0.0), the /32 (/36 for he.net) for IPv6, the string
// "local" for a local address and the string "unroutable" for an unroutable
// address.
func (a *addrBook) groupKey(na *p2p.NetAddress) string { _ = "STUB: not implemented"; return "" }

func groupKeyFor(na *p2p.NetAddress, routabilityStrict bool) string {
	_ = "STUB: not implemented"
	return ""
}

// last four bytes are the ip address

// teredo tunnels have the last 4 bytes as the v4 address XOR
// 0xff.

// group is keyed off the first 4 bits of the actual onion key.

// OK, so now we know ourselves to be a IPv6 address.
// bitcoind uses /32 for everything, except for Hurricane Electric's
// (he.net) IP range, which it uses /36 for.

func (a *addrBook) hash(b []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
