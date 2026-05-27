package main

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/cometbft/cometbft/libs/log"
	e2e "github.com/cometbft/cometbft/test/e2e/pkg"
	"github.com/cometbft/cometbft/test/e2e/pkg/infra"
)

const randomSeed = 2308084734268

var logger = log.NewTMLogger(log.NewSyncWriter(os.Stdout))

func main() {
	NewCLI().Run()
}

// CLI is the Cobra-based command-line interface.
type CLI struct {
	root     *cobra.Command
	testnet  *e2e.Testnet
	preserve bool
	infp     infra.Provider
}

// NewCLI sets up the CLI.
func NewCLI() *CLI { _ = "STUB: not implemented"; return nil }

// we'll output them ourselves in Run()

//nolint: gosec

// allow some txs to go through

// allow some txs to go through

// ensure chain progress

// wait for network to settle before tests

//nolint: gosec

// allow some txs to go through

// we benchmark performance over the next 100 blocks

// Run runs the CLI.
func (cli *CLI) Run() { _ = "STUB: not implemented"; return }
