package docker

import (
	"context"

	e2e "github.com/cometbft/cometbft/test/e2e/pkg"
	"github.com/cometbft/cometbft/test/e2e/pkg/infra"
)

var _ infra.Provider = (*Provider)(nil)

// Provider implements a docker-compose backed infrastructure provider.
type Provider struct {
	infra.ProviderData
}

// Setup generates the docker-compose file and write it to disk, erroring if
// any of these operations fail.
func (p *Provider) Setup() error { _ = "STUB: not implemented"; return nil }

//nolint: gosec
// G306: Expect WriteFile permissions to be 0600 or less

func (p Provider) StartNodes(ctx context.Context, nodes ...*e2e.Node) error {
	_ = "STUB: not implemented"
	return nil
}

func (p Provider) StopTestnet(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// dockerComposeBytes generates a Docker Compose config file for a testnet and returns the
// file as bytes to be written out to disk.
func dockerComposeBytes(testnet *e2e.Testnet) ([]byte, error) {
	_ = "STUB: not implemented"
	// Must use version 2 Docker Compose format, to support IPv6.
	return nil, nil
}

// ExecCompose runs a Docker Compose command for a testnet.
func ExecCompose(ctx context.Context, dir string, args ...string) error {
	_ = "STUB: not implemented"
	return nil
}

// ExecCompose runs a Docker Compose command for a testnet and returns the command's output.
func ExecComposeOutput(ctx context.Context, dir string, args ...string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ExecComposeVerbose runs a Docker Compose command for a testnet and displays its output.
func ExecComposeVerbose(ctx context.Context, dir string, args ...string) error {
	_ = "STUB: not implemented"
	return nil
}

// Exec runs a Docker command.
func Exec(ctx context.Context, args ...string) error { _ = "STUB: not implemented"; return nil }
