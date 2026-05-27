package mempool

import (
	"errors"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/libs/service"
	"github.com/cometbft/cometbft/p2p"
	"github.com/cometbft/cometbft/types"
)

// NopMempool is a mempool that does nothing.
//
// The ABCI app is responsible for storing, disseminating, and proposing transactions.
// See [ADR-111](../docs/architecture/adr-111-nop-mempool.md).
type NopMempool struct{}

// errNotAllowed indicates that the operation is not allowed with `nop` mempool.
var errNotAllowed = errors.New("not allowed with `nop` mempool")

var _ Mempool = &NopMempool{}

// CheckTx always returns an error.
func (*NopMempool) CheckTx(types.Tx, func(*abci.ResponseCheckTx), TxInfo) error {
	_ = "STUB: not implemented"
	return nil

	// RemoveTxByKey always returns an error.
}

func (*NopMempool) RemoveTxByKey(types.TxKey) error { _ = "STUB: not implemented"; return nil }

// ReapMaxBytesMaxGas always returns nil.
func (*NopMempool) ReapMaxBytesMaxGas(int64, int64) []*types.CachedTx {
	_ = "STUB: not implemented"

	// ReapMaxTxs always returns nil.
	return nil
}

func (*NopMempool) ReapMaxTxs(int) []*types.CachedTx {
	_ = "STUB: not implemented"

	// Lock does nothing.
	return nil
}

func (*NopMempool) Lock() {
	_ = "STUB: not implemented"

	// Unlock does nothing.
	return
}

func (*NopMempool) Unlock() {
	_ = "STUB: not implemented"

	// Update does nothing.
	return
}

func (*NopMempool) Update(
	int64,
	[]*types.CachedTx,
	[]*abci.ExecTxResult,
	PreCheckFunc,
	PostCheckFunc,
) error {
	_ = "STUB: not implemented"

	// FlushAppConn does nothing.
	return nil
}

func (*NopMempool) FlushAppConn() error {
	_ = "STUB: not implemented"

	// Flush does nothing.
	return nil
}

func (*NopMempool) Flush() {
	_ = "STUB: not implemented"

	// TxsAvailable always returns nil.
	return
}

func (*NopMempool) TxsAvailable() <-chan struct{} {
	_ = "STUB: not implemented"

	// EnableTxsAvailable does nothing.
	return nil
}

func (*NopMempool) EnableTxsAvailable() {
	_ = "STUB: not implemented"

	// SetTxRemovedCallback does nothing.
	return
}

func (*NopMempool) SetTxRemovedCallback(func(txKey types.TxKey)) {
	_ = "STUB: not implemented"

	// Size always returns 0.
	return
}

func (*NopMempool) Size() int {
	_ = "STUB: not implemented"

	// SizeBytes always returns 0.
	return 0
}

func (*NopMempool) SizeBytes() int64 {
	_ = "STUB: not implemented"

	// GetTxByKey always returns nil.
	return 0
}

func (*NopMempool) GetTxByKey(types.TxKey) (*types.CachedTx, bool) {
	_ = "STUB: not implemented"

	// WasRecentlyEvicted always returns false.
	return nil, false
}

func (*NopMempool) WasRecentlyEvicted(types.TxKey) bool {
	_ = "STUB: not implemented"

	// WasRecentlyRejected always returns false, zero and an empty string.
	return false
}

func (*NopMempool) WasRecentlyRejected(types.TxKey) (bool, uint32, string) {
	_ = "STUB: not implemented"
	return false,

		// NopMempoolReactor is a mempool reactor that does nothing.
		0, ""
}

type NopMempoolReactor struct {
	service.BaseService
}

// NewNopMempoolReactor returns a new `nop` reactor.
//
// To be used only in RPC.
func NewNopMempoolReactor() *NopMempoolReactor { _ = "STUB: not implemented"; return nil }

var _ p2p.Reactor = &NopMempoolReactor{}

// GetChannels always returns nil.
func (*NopMempoolReactor) GetChannels() []*p2p.ChannelDescriptor {
	_ = "STUB: not implemented"

	// AddPeer does nothing.
	return nil
}

func (*NopMempoolReactor) AddPeer(p2p.Peer) {
	_ = "STUB: not implemented"

	// InitPeer always returns nil.
	return
}

func (*NopMempoolReactor) InitPeer(p2p.Peer) (p2p.Peer, error) {
	_ = "STUB: not implemented"

	// RemovePeer does nothing.
	return *new(p2p.Peer), nil
}

func (*NopMempoolReactor) RemovePeer(p2p.Peer, interface{}) {
	_ = "STUB: not implemented"

	// Receive does nothing.
	return
}

func (*NopMempoolReactor) Receive(p2p.Envelope) {
	_ = "STUB: not implemented"

	// SetSwitch does nothing.
	return
}

func (*NopMempoolReactor) SetSwitch(*p2p.Switch) {
	_ = "STUB: not implemented"

	// QueueUnprocessedEnvelope does nothing.
	return
}

func (*NopMempoolReactor) QueueUnprocessedEnvelope(p2p.UnprocessedEnvelope) {
	_ = "STUB: not implemented"
	return
}
