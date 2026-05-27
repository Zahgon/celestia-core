package test

import (
	"context"
	"testing"

	"github.com/cometbft/cometbft/types"
)

func Validator(_ context.Context, votingPower int64) (*types.Validator, types.PrivValidator, error) {
	_ = "STUB: not implemented"
	return nil, *new(types.PrivValidator), nil
}

func ValidatorSet(ctx context.Context, t *testing.T, numValidators int, votingPower int64) (*types.ValidatorSet, []types.PrivValidator) {
	_ = "STUB: not implemented"
	return nil, nil
}
