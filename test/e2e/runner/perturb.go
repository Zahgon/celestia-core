package main

import (
	"context"

	rpctypes "github.com/cometbft/cometbft/rpc/core/types"
	e2e "github.com/cometbft/cometbft/test/e2e/pkg"
)

// Perturbs a running testnet.
func Perturb(ctx context.Context, testnet *e2e.Testnet) error {
	_ = "STUB: not implemented"
	return nil
}

// give network some time to recover between each

// PerturbNode perturbs a node with a given perturbation, returning its status
// after recovering.
func PerturbNode(ctx context.Context, node *e2e.Node, perturbation e2e.Perturbation) (*rpctypes.ResultStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
