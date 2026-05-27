package mock

import (
	"context"
	"sync"

	"github.com/cometbft/cometbft/light/provider"
	"github.com/cometbft/cometbft/types"
)

type Mock struct {
	chainID string

	mtx              sync.Mutex
	headers          map[int64]*types.SignedHeader
	vals             map[int64]*types.ValidatorSet
	evidenceToReport map[string]types.Evidence // hash => evidence
	latestHeight     int64
}

var _ provider.Provider = (*Mock)(nil)

// New creates a mock provider with the given set of headers and validator
// sets.
func New(chainID string, headers map[int64]*types.SignedHeader, vals map[int64]*types.ValidatorSet) *Mock {
	_ = "STUB: not implemented"
	return nil
}

// ChainID returns the blockchain ID.
func (p *Mock) ChainID() string { _ = "STUB: not implemented"; return "" }

func (p *Mock) String() string { _ = "STUB: not implemented"; return "" }

func (p *Mock) LightBlock(ctx context.Context, height int64) (*types.LightBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// allocate a window of time for contexts to be canceled

func (p *Mock) ReportEvidence(_ context.Context, ev types.Evidence) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Mock) HasEvidence(ev types.Evidence) bool { _ = "STUB: not implemented"; return false }

func (p *Mock) AddLightBlock(lb *types.LightBlock) { _ = "STUB: not implemented"; return }

func (p *Mock) Copy(id string) *Mock { _ = "STUB: not implemented"; return nil }
