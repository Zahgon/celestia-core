package pex

import (
	"sync"
	"time"

	"github.com/cometbft/cometbft/libs/cmap"
	"github.com/cometbft/cometbft/p2p"
	"github.com/cometbft/cometbft/p2p/conn"
	tmp2p "github.com/cometbft/cometbft/proto/tendermint/p2p"
)

type Peer = p2p.Peer

const (
	// PexChannel is a channel for PEX messages
	PexChannel = byte(0x00)

	// over-estimate of max NetAddress size
	// hexID (40) + IP (16) + Port (2) + Name (100) ...
	// NOTE: dont use massive DNS name ..
	maxAddressSize = 256

	// NOTE: amplificaiton factor!
	// small request results in up to maxMsgSize response
	maxMsgSize = maxAddressSize * maxGetSelection

	// ensure we have enough peers
	defaultEnsurePeersPeriod = 10 * time.Second

	// Seed/Crawler constants

	// minTimeBetweenCrawls is a minimum time between attempts to crawl a peer.
	minTimeBetweenCrawls = 2 * time.Minute

	// check some peers every this
	crawlPeerPeriod = 30 * time.Second

	maxAttemptsToDial = 8 // 8 attempts with 10s interval

	// if node connects to seed, it does not have any trusted peers.
	// Especially in the beginning, node should have more trusted peers than
	// untrusted.
	biasToSelectNewPeers = 30 // 70 to select good peers

	// if a peer is marked bad, it will be banned for at least this time period
	defaultBanTime = 24 * time.Hour

	// ReactorIncomingMessageQueueSize the size of the reactor's message queue.
	ReactorIncomingMessageQueueSize = 10
)

type errMaxAttemptsToDial struct{}

func (e errMaxAttemptsToDial) Error() string { _ = "STUB: not implemented"; return "" }

type errTooEarlyToDial struct {
	backoffDuration time.Duration
	lastDialed      time.Time
}

func (e errTooEarlyToDial) Error() string { _ = "STUB: not implemented"; return "" }

// Reactor handles PEX (peer exchange) and ensures that an
// adequate number of peers are connected to the switch.
//
// It uses `AddrBook` (address book) to store `NetAddress`es of the peers.
//
// ## Preventing abuse
//
// Only accept pexAddrsMsg from peers we sent a corresponding pexRequestMsg too.
// Only accept one pexRequestMsg every ~defaultEnsurePeersPeriod.
type Reactor struct {
	p2p.BaseReactor

	book              AddrBook
	config            *ReactorConfig
	ensurePeersCh     chan struct{} // Wakes up ensurePeersRoutine()
	ensurePeersPeriod time.Duration // TODO: should go in the config

	// maps to prevent abuse
	requestsSent         *cmap.CMap // ID->struct{}: unanswered send requests
	lastReceivedRequests *cmap.CMap // ID->time.Time: last time peer requested from us

	seedAddrs []*p2p.NetAddress

	attemptsToDial sync.Map // address (string) -> {number of attempts (int), last time dialed (time.Time)}

	// seed/crawled mode fields
	crawlPeerInfos map[p2p.ID]crawlPeerInfo
}

func (r *Reactor) minReceiveRequestInterval() time.Duration {
	_ = "STUB: not implemented"
	// NOTE: must be less than ensurePeersPeriod, otherwise we'll request
	// peers too quickly from others and they'll think we're bad!
	// According to the spec, the minimum accepted interval should be
	// ensurePeersPeriod / 3 to allow for timing variations while still
	// preventing abuse.
	return *new(time.Duration)
}

// ReactorConfig holds reactor specific configuration data.
type ReactorConfig struct {
	// Seed/Crawler mode
	SeedMode bool

	// We want seeds to only advertise good peers. Therefore they should wait at
	// least as long as we expect it to take for a peer to become good before
	// disconnecting.
	SeedDisconnectWaitPeriod time.Duration

	// Maximum pause when redialing a persistent peer (if zero, exponential backoff is used)
	PersistentPeersMaxDialPeriod time.Duration

	// Seeds is a list of addresses reactor may use
	// if it can't connect to peers in the addrbook.
	Seeds []string
}

type _attemptsToDial struct {
	number     int
	lastDialed time.Time
}

// NewReactor creates new PEX reactor.
func NewReactor(b AddrBook, config *ReactorConfig) *Reactor { _ = "STUB: not implemented"; return nil }

// OnStart implements BaseService
func (r *Reactor) OnStart() error { _ = "STUB: not implemented"; return nil }

// Check if this node should run
// in seed/crawler mode

// OnStop implements BaseService
func (r *Reactor) OnStop() { _ = "STUB: not implemented"; return }

// GetChannels implements Reactor
func (r *Reactor) GetChannels() []*conn.ChannelDescriptor { _ = "STUB: not implemented"; return nil }

// AddPeer implements Reactor by adding peer to the address book (if inbound)
// or by requesting more addresses (if outbound).
func (r *Reactor) AddPeer(p Peer) {
	_ = "STUB: not implemented"

	// For outbound peers, the address is already in the books -
	// either via DialPeersAsync or r.Receive.
	// Ask it for more peers if we need.
	return
}

// inbound peer is its own source

// Make it explicit that addr and src are the same for an inbound peer.

// add to book. dont RequestAddrs right away because
// we don't trust inbound as much - let ensurePeersRoutine handle it.

// RemovePeer implements Reactor by resetting peer's requests info.
func (r *Reactor) RemovePeer(p Peer, _ interface{}) { _ = "STUB: not implemented"; return }

func (r *Reactor) logErrAddrBook(err error) { _ = "STUB: not implemented"; return }

// non-routable, self, full book, private, etc.

// Receive implements Reactor by handling incoming PEX messages.
func (r *Reactor) Receive(e p2p.Envelope) { _ = "STUB: not implemented"; return }

// NOTE: this is a prime candidate for amplification attacks,
// so it's important we
// 1) restrict how frequently peers can request
// 2) limit the output size

// If we're a seed and this is an inbound peer,
// respond once and disconnect.

// FlushStop/StopPeer are already
// running in a go-routine.

// Send addrs and disconnect

// In a go-routine so it doesn't block .Receive.

// Check we're not receiving requests too frequently.

// If we asked for addresses, add them to the book

// enforces a minimum amount of time between requests
func (r *Reactor) receiveRequest(src Peer) error { _ = "STUB: not implemented"; return nil }

// initialize with empty time

// first time gets a free pass. then we start tracking the time

// RequestAddrs asks peer for more addresses if we do not already have a
// request out for this peer.
func (r *Reactor) RequestAddrs(p Peer) { _ = "STUB: not implemented"; return }

// ReceiveAddrs adds the given addrs to the addrbook if theres an open
// request for this peer and deletes the open request.
// If there's no open request for the src peer, it returns an error.
func (r *Reactor) ReceiveAddrs(addrs []*p2p.NetAddress, src Peer) error {
	_ = "STUB: not implemented"
	return nil
}

// NOTE: we check netAddr validity and routability in book#AddAddress.

// XXX: should we be strict about incoming data and disconnect from a
// peer here too?

// Try to connect to addresses coming from a seed node without waiting (#2093)

// SendAddrs sends addrs to the peer.
func (r *Reactor) SendAddrs(p Peer, netAddrs []*p2p.NetAddress) { _ = "STUB: not implemented"; return }

// safeMaxAddressEstimationMargin is a factor used to estimate the maximum number of addresses fitting into a size limit.
const safeMaxAddressEstimationMargin = 0.9

// capAddresses transforms a slice of NetAddress into PexAddrs while ensuring the resulting message size doesn't exceed maxMsgSize.
func capAddresses(netAddrs []*p2p.NetAddress) *tmp2p.PexAddrs {
	_ = "STUB: not implemented"
	return nil
}

// Calculate how many addresses we can fit based on the ratio

// SetEnsurePeersPeriod sets period to ensure peers connected.
func (r *Reactor) SetEnsurePeersPeriod(d time.Duration) { _ = "STUB: not implemented"; return }

// Ensures that sufficient peers are connected. (continuous)
func (r *Reactor) ensurePeersRoutine() { _ = "STUB: not implemented"; return }

// Randomize first round of communication to avoid thundering herd.
// If no peers are present directly start connecting so we guarantee swift
// setup with the help of configured seeds.

// fire once immediately.
// ensures we dial the seeds right away if the book is empty

// fire periodically

// ensurePeers ensures that sufficient peers are connected. (once)
//
// heuristic that we haven't perfected yet, or, perhaps is manually edited by
// the node operator. It should not be used to compute what addresses are
// already connected or not.
func (r *Reactor) ensurePeers(ensurePeersPeriodElapsed bool) { _ = "STUB: not implemented"; return }

// check if the addressbook is smaller than maxDials

// We don't need to randomize the addresses since the addressbook is already shuffled

// Check if banned nodes can be reinstated

// 1) Pick a random peer and ask for more.

// Get updated address book and if empty, dial seeds

func (r *Reactor) dialAttemptsInfo(addr *p2p.NetAddress) (attempts int, lastDialed time.Time) {
	_ = "STUB: not implemented"
	return 0, *new(time.Time)
}

// isDuplicatePeerIDRejection reports whether err is a rejection caused by a
// peer with id wantedID already being connected. Used by dialPeer to skip
// backoff when an outbound dial loses a race to an inbound connection from
// the same peer. Duplicate-IP rejections (no peer ID recorded) and
// duplicate-ID rejections naming a *different* peer return false so they
// follow the normal failure path — preventing an unrelated peer sharing an
// IP, or a peer presenting an unexpected ID, from suppressing backoff for
// the address we actually wanted to reach.
func isDuplicatePeerIDRejection(err error, wantedID p2p.ID) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *Reactor) dialPeer(addr *p2p.NetAddress) error { _ = "STUB: not implemented"; return nil }

// If the rejection came from a duplicate *peer ID* matching the peer
// we tried to dial, treat the dial as success-equivalent: the peer is
// already connected (typically via an inbound connection raced against
// our outbound dial), so we shouldn't penalize this address with the
// 30s dial backoff or increment its attempt counter. Duplicate-IP
// rejections take the normal failure path, since the colliding peer
// may not be the one we wanted (e.g. an unrelated peer sharing an IP).

// NOTE: addr is removed from addrbook in markAddrInBookBasedOnErr

// cleanup any history

// checkSeeds checks that addresses are well formed.
// Returns number of seeds we can connect to, along with all seeds addrs.
// return err if user provided any badly formatted seed addresses.
// Doesn't error if the seed node can't be reached.
// numOnline returns -1 if no seed nodes were in the initial configuration.
func (r *Reactor) checkSeeds() (numOnline int, netAddrs []*p2p.NetAddress, err error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

// randomly dial seeds until we connect to one or exhaust them
func (r *Reactor) dialSeeds() { _ = "STUB: not implemented"; return }

// perm := r.Switch.rng.Perm(lSeeds)

// dial a random seed

// do not write error message if there were no seeds specified in config

// AttemptsToDial returns the number of attempts to dial specific address. It
// returns 0 if never attempted or successfully connected.
func (r *Reactor) AttemptsToDial(addr *p2p.NetAddress) int { _ = "STUB: not implemented"; return 0 }

//----------------------------------------------------------

// Explores the network searching for more peers. (continuous)
// Seed/Crawler Mode causes this node to quickly disconnect
// from peers, except other seed nodes.
func (r *Reactor) crawlPeersRoutine() {
	_ = "STUB: not implemented"
	// If we have any seed nodes, consult them first
	return
}

// Do an initial crawl

// Fire periodically

// nodeHasSomePeersOrDialingAny returns true if the node is connected to some
// peers or dialing them currently.
func (r *Reactor) nodeHasSomePeersOrDialingAny() bool { _ = "STUB: not implemented"; return false }

// crawlPeerInfo handles temporary data needed for the network crawling
// performed during seed/crawler mode.
type crawlPeerInfo struct {
	Addr *p2p.NetAddress `json:"addr"`
	// The last time we crawled the peer or attempted to do so.
	LastCrawled time.Time `json:"last_crawled"`
}

// crawlPeers will crawl the network looking for new peer addresses.
func (r *Reactor) crawlPeers(addrs []*p2p.NetAddress) { _ = "STUB: not implemented"; return }

// Do not attempt to connect with peers we recently crawled.

// Record crawling attempt.

func (r *Reactor) cleanupCrawlPeerInfos() { _ = "STUB: not implemented"; return }

// If we did not crawl a peer for 24 hours, it means the peer was removed
// from the addrbook => remove
//
// 10000 addresses / maxGetSelection = 40 cycles to get all addresses in
// the ideal case,
// 40 * crawlPeerPeriod ~ 20 minutes

// attemptDisconnects checks if we've been with each peer long enough to disconnect
func (r *Reactor) attemptDisconnects() { _ = "STUB: not implemented"; return }

func markAddrInBookBasedOnErr(addr *p2p.NetAddress, book AddrBook, err error) {
	_ = "STUB: not implemented"
	// TODO: detect more "bad peer" scenarios
	return
}
