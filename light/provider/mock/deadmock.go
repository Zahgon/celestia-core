package mock

import (
	"context"

	"github.com/cometbft/cometbft/light/provider"
	"github.com/cometbft/cometbft/types"
)

type deadMock struct {
	chainID string
}

// NewDeadMock creates a mock provider that always errors.
func NewDeadMock(chainID string) provider.Provider {
	_ = "STUB: not implemented"
	return *new(provider.Provider)
}

func (p *deadMock) ChainID() string { _ = "STUB: not implemented"; return "" }

func (p *deadMock) String() string { _ = "STUB: not implemented"; return "" }

func (p *deadMock) LightBlock(context.Context, int64) (*types.LightBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *deadMock) ReportEvidence(context.Context, types.Evidence) error {
	_ = "STUB: not implemented"
	return nil
}
