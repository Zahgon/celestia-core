package consensus

import (
	"context"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/libs/clist"
	mempl "github.com/cometbft/cometbft/mempool"
	"github.com/cometbft/cometbft/proxy"
	"github.com/cometbft/cometbft/types"
)

//-----------------------------------------------------------------------------

type emptyMempool struct{}

var _ mempl.Mempool = emptyMempool{}

func (emptyMempool) Lock()            { _ = "STUB: not implemented"; return }
func (emptyMempool) Unlock()          { _ = "STUB: not implemented"; return }
func (emptyMempool) Size() int        { _ = "STUB: not implemented"; return 0 }
func (emptyMempool) SizeBytes() int64 { _ = "STUB: not implemented"; return 0 }
func (emptyMempool) CheckTx(types.Tx, func(*abci.ResponseCheckTx), mempl.TxInfo) error {
	_ = "STUB: not implemented"
	return nil
}

func (emptyMempool) GetTxByKey(types.TxKey) (*types.CachedTx, bool) {
	_ = "STUB: not implemented"
	return nil, false
}
func (emptyMempool) WasRecentlyEvicted(types.TxKey) bool { _ = "STUB: not implemented"; return false }
func (emptyMempool) WasRecentlyRejected(types.TxKey) (bool, uint32, string) {
	_ = "STUB: not implemented"
	return false, 0, ""
}

func (txmp emptyMempool) RemoveTxByKey(types.TxKey) error { _ = "STUB: not implemented"; return nil }

func (emptyMempool) ReapMaxBytesMaxGas(int64, int64) []*types.CachedTx {
	_ = "STUB: not implemented"
	return nil
}

func (emptyMempool) ReapMaxTxs(int) []*types.CachedTx { _ = "STUB: not implemented"; return nil }
func (emptyMempool) Update(
	int64,
	[]*types.CachedTx,
	[]*abci.ExecTxResult,
	mempl.PreCheckFunc,
	mempl.PostCheckFunc,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (emptyMempool) Flush()                        { _ = "STUB: not implemented"; return }
func (emptyMempool) FlushAppConn() error           { _ = "STUB: not implemented"; return nil }
func (emptyMempool) TxsAvailable() <-chan struct{} { _ = "STUB: not implemented"; return nil }
func (emptyMempool) EnableTxsAvailable()           { _ = "STUB: not implemented"; return }
func (emptyMempool) TxsBytes() int64               { _ = "STUB: not implemented"; return 0 }

func (emptyMempool) TxsFront() *clist.CElement    { _ = "STUB: not implemented"; return nil }
func (emptyMempool) TxsWaitChan() <-chan struct{} { _ = "STUB: not implemented"; return nil }

func (emptyMempool) InitWAL() error { _ = "STUB: not implemented"; return nil }
func (emptyMempool) CloseWAL() {
	_ = "STUB: not implemented"

	// -----------------------------------------------------------------------------
	// mockProxyApp uses ABCIResponses to give the right results.
	//
	// Useful because we don't want to call Commit() twice for the same block on
	// the real app.
	return
}

func newMockProxyApp(finalizeBlockResponse *abci.ResponseFinalizeBlock) proxy.AppConnConsensus {
	_ = "STUB: not implemented"
	return *new(proxy.AppConnConsensus)
}

type mockProxyApp struct {
	abci.BaseApplication
	finalizeBlockResponse *abci.ResponseFinalizeBlock
}

func (mock *mockProxyApp) FinalizeBlock(context.Context, *abci.RequestFinalizeBlock) (*abci.ResponseFinalizeBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
