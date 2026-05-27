package cat

import (
	"errors"
	"sync/atomic"
	"time"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cometbft/cometbft/libs/trace"
	"github.com/cometbft/cometbft/mempool"
	"github.com/cometbft/cometbft/p2p"
	"github.com/cometbft/cometbft/types"
)

const (
	// default duration to wait before considering a peer non-responsive
	// and searching for the tx from a new peer
	DefaultGossipDelay = 60 * time.Second

	// MempoolDataChannel channel for SeenTx and blob messages.
	MempoolDataChannel = byte(0x31)

	// MempoolWantsChannel channel for wantTx messages.
	MempoolWantsChannel = byte(0x32)

	// peerHeightDiff signifies the tolerance in difference in height between the peer and the height
	// the node received the tx
	peerHeightDiff = 10

	// ReactorIncomingMessageQueueSize the size of the reactor's message queue.
	ReactorIncomingMessageQueueSize = 5000

	// maxSeenTxBroadcast defines the maximum number of peers to which a SeenTx message should be broadcasted.
	maxSeenTxBroadcast = 15

	// defaultMaxPersistentStickyPeers caps how many persistent peers are guaranteed
	// to receive SeenTx broadcasts per signer (added on top of the natural sticky
	// set, never displacing it). Used when ReactorOptions.MaxPersistentStickyPeers
	// is unset.
	defaultMaxPersistentStickyPeers = 4

	// maxReceivedBufferSize limits how far ahead of the expected sequence we will
	// request/buffer transactions. Txs with sequence > expected + maxReceivedBufferSize are rejected.
	maxReceivedBufferSize = 30

	// maxRequestsPerPeer limits the number of concurrent outstanding requests to a single peer.
	// When a peer reaches this limit, requests will be sent to alternative peers instead.
	maxRequestsPerPeer = 30

	// maxSignerLength is the maximum allowed length for a signer field in SeenTx messages.
	maxSignerLength = 64
)

var (
	errSignerTooLong = errors.New("signer field exceeds maximum length")
	errTooManyTxs    = errors.New("txs message contains too many transactions")
	errEmptyTx       = errors.New("txs message contains an empty transaction")
)

// Reactor handles mempool tx broadcasting logic amongst peers. For the main
// logic behind the protocol, refer to `ReceiveEnvelope` or to the english
// spec under /.spec.md
type Reactor struct {
	p2p.BaseReactor
	opts           *ReactorOptions
	mempool        *TxPool
	ids            *mempoolIDs
	requests       *requestScheduler
	pendingSeen    *pendingSeenTracker
	receivedBuffer *receivedTxBuffer // buffer for out-of-order tx arrivals
	traceClient    trace.Tracer
	// stickySalt stores []byte rendezvous salt for sticky peer selection; nil/empty keeps default ordering.
	stickySalt atomic.Value
}

type ReactorOptions struct {
	// ListenOnly means that the node will never broadcast any of the transactions that
	// it receives. This is useful for keeping transactions private
	ListenOnly bool

	// MaxTxSize is the maximum size of a transaction that can be received
	MaxTxSize int

	// MaxGossipDelay is the maximum allotted time that the reactor expects a transaction to
	// arrive before issuing a new request to a different peer
	MaxGossipDelay time.Duration

	// TraceClient is the trace client for collecting trace level events
	TraceClient trace.Tracer

	// StickyPeerSalt is used to derive the rendezvous hash for sticky peer selection.
	// If unset, all nodes will derive the same peer ordering.
	StickyPeerSalt []byte

	// MaxPersistentStickyPeers caps how many persistent peers are guaranteed
	// to receive SeenTx broadcasts per signer, added on top of the natural sticky
	// set without displacing it. <= 0 falls back to defaultMaxPersistentStickyPeers.
	MaxPersistentStickyPeers int
}

func (opts *ReactorOptions) VerifyAndComplete() error { _ = "STUB: not implemented"; return nil }

// NewReactor returns a new Reactor with the given config and mempool.
func NewReactor(mempool *TxPool, opts *ReactorOptions) (*Reactor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// establish concrete []byte type for atomic.Value

// SetLogger sets the Logger on the reactor and the underlying mempool.
func (memR *Reactor) SetLogger(l log.Logger) {
	_ = "STUB: not implemented"

	// SetStickySalt configures the salt used for sticky peer selection. Passing nil resets it.
	return
}

func (memR *Reactor) SetStickySalt(salt []byte) { _ = "STUB: not implemented"; return }

func (memR *Reactor) currentStickyPeerSalt() []byte { _ = "STUB: not implemented"; return nil }

// OnStart implements Service.
func (memR *Reactor) OnStart() error { _ = "STUB: not implemented"; return nil }

// listen in for any newly verified tx via RPC, then immediately
// broadcast it to all connected peers.

// OnStop implements Service
func (memR *Reactor) OnStop() {
	_ = "STUB: not implemented"
	// stop all the timers tracking outbound requests
	return
}

// GetChannels implements Reactor by returning the list of channels for this
// reactor.
func (memR *Reactor) GetChannels() []*p2p.ChannelDescriptor { _ = "STUB: not implemented"; return nil }

// InitPeer implements Reactor by creating a state for the peer.
func (memR *Reactor) InitPeer(peer p2p.Peer) (p2p.Peer, error) {
	_ = "STUB: not implemented"
	return *new(p2p.Peer), nil
}

// RemovePeer implements Reactor. For all current outbound requests to this
// peer it will find a new peer to rerequest the same transactions.
func (memR *Reactor) RemovePeer(peer p2p.Peer, reason interface{}) {
	_ = "STUB: not implemented"
	return
}

// clear all memory of seen txs by that peer

// remove and rerequest all pending outbound requests to that peer since we know
// we won't receive any responses from them.

// ReceiveEnvelope implements Reactor.
// It processes one of three messages: Txs, SeenTx, WantTx.
func (memR *Reactor) Receive(e p2p.Envelope) { _ = "STUB: not implemented"; return }

// A peer has sent us one or more transactions. This could be either because we requested them
// or because the peer received a new transaction and is broadcasting it to us.
// NOTE: This setup also means that we can support older mempool implementations that simply
// flooded the network with transactions.

// If we requested the transaction we mark it as received.

// If we didn't request the transaction we simply mark the peer as having the
// tx (we'd have already done it if we were requesting the tx).

// Look up signer/sequence from pending tracker

// We have sequence info - check if we should buffer or process

// Future sequence within lookahead - buffer it for later

// Process this tx through CheckTx without putting into buffer

// A peer has indicated to us that it has a transaction. We first verify the txkey and
// mark that peer as having the transaction. Then we proceed with the following logic:
//
// 1. If we have the transaction, we do nothing.
// 2. If we don't yet have the tx but have an outgoing request for it, we do nothing.
// 3. If we recently evicted the tx and still don't have space for it, we do nothing.
// 4. Else, we request the transaction from that peer.

// Check if we don't already have the transaction

// If we are already requesting that tx, then we don't need to go any further.

// fall through and request immediately when sequence info is missing

// fall through and request immediately if we cannot query the application

// fall through and request immediately for the expected sequence

// TODO: add per-peer limits or something similar to pendingSeen to prevent overflowing

// We don't have the transaction, nor are we requesting it so we send the node
// a want msg. Enforce the per-peer request limit so a single peer cannot
// drive unbounded outstanding requests via the direct SeenTx path.

// A peer is requesting a transaction that we have claimed to have. Find the specified
// transaction and broadcast it to the peer. We may no longer have the transaction

// PeerState describes the state of a peer.
type PeerState interface {
	GetHeight() int64
}

// broadcastSeenTx broadcasts a SeenTx message to limited peers unless we
// know they have already seen the transaction
func (memR *Reactor) broadcastSeenTx(txKey types.TxKey, signer []byte, sequence uint64) {
	_ = "STUB: not implemented"
	return
}

// broadcastNewTx broadcast new transaction to limited peers unless we are already sure they have seen the tx.
func (memR *Reactor) broadcastNewTx(wtx *wrappedTx) { _ = "STUB: not implemented"; return }

// broadcastSeenTxWithHeight is a helper that broadcasts a SeenTx message with height checking.
func (memR *Reactor) broadcastSeenTxWithHeight(txKey types.TxKey, height int64, signer []byte, sequence uint64) {
	_ = "STUB: not implemented"
	return
}

// Send to the natural top maxSeenTxBroadcast peers; additionally guarantee
// up to maxPersistent persistent peers receive the SeenTx, even if they
// rank outside the top maxSeenTxBroadcast. Persistent peers are appended,
// never displacing the natural set.

// make sure peer isn't too far behind. This can happen
// if the peer is blocksyncing still and catching up
// in which case we just skip sending the transaction

// requesting it from another peer if the first peer does not respond.
func (memR *Reactor) requestTx(txKey types.TxKey, peer p2p.Peer) bool {
	_ = "STUB: not implemented"

	// we have disconnected from the peer
	return false
}

// tryAddNewTx attempts to add a tx to the mempool and traces the result.
// Returns the response and true if processing should continue (success or already in mempool).
func (memR *Reactor) tryAddNewTx(cachedTx *types.CachedTx, key types.TxKey, txInfo mempool.TxInfo, peerID string) (*abci.ResponseCheckTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// processReceivedTx handles a received transaction by running CheckTx and then
// draining any buffered transactions for the same signer.
func (memR *Reactor) processReceivedTx(cachedTx *types.CachedTx, key types.TxKey, txInfo mempool.TxInfo, src p2p.Peer) {
	_ = "STUB: not implemented"
	return
}

// processReceivedBuffer drains buffered transactions for a signer in sequence order.
// It processes buffered txs as long as they match the next expected sequence.
func (memR *Reactor) processReceivedBuffer(signer []byte) { _ = "STUB: not implemented"; return }

// processPendingSeenForSigner tries to advance the pipeline of queued transactions for a signer.
// It requests consecutive sequences in parallel from different peers whenever we have seen a consecutive sequence numbers,
// buffering out-of-order arrivals for later processing. This allows fast catch-up even when tx sources are distributed.
func (memR *Reactor) processPendingSeenForSigner(signer []byte) { _ = "STUB: not implemented"; return }

// Clean up old entries and request consecutive sequences in parallel

// Clean up entries that are already processed

// Check limits

// Skip if already in mempool

// Skip if already being requested, but count it

// Request from first available peer

func (memR *Reactor) tryRequestQueuedTx(entry *pendingSeenTx) bool {
	_ = "STUB: not implemented"
	// Try each peer that has seen this tx, skipping those at capacity
	return false
}

// No known peer available, try to find a new one

func (memR *Reactor) onRequestTimeout(txKey types.TxKey, peerID uint16) {
	_ = "STUB: not implemented"
	return
}

func (memR *Reactor) heightSignalLoop() { _ = "STUB: not implemented"; return }

func (memR *Reactor) refreshPendingSeenQueues() {
	_ = "STUB: not implemented"
	// Collect signers from both pending seen and buffer
	return
}

// First drain any buffered txs that can now be processed
// (their expected sequence may have advanced due to block commit)

// Then request more pending txs

func (memR *Reactor) querySequenceFromApplication(signer []byte) (uint64, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

// If the response is 0, treat it as "sequence tracking not available"
// to maintain backward compatibility with apps that don't implement QuerySequence.
// This prevents transactions from being stuck in pendingSeen when the app returns 0.

// findNewPeerToSendTx finds a new peer that has already seen the transaction to
// request a transaction from.
func (memR *Reactor) findNewPeerToRequestTx(txKey types.TxKey) {
	_ = "STUB: not implemented"
	// ensure that we are connected to peers
	return
}

// No other free peer has the transaction we are looking for.
// We give up 🤷‍♂️ and hope either a peer responds late or the tx
// is gossiped again
