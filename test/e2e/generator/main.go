package main

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/cometbft/cometbft/libs/log"
)

const (
	randomSeed int64 = 4827085738
)

var logger = log.NewTMLogger(log.NewSyncWriter(os.Stdout))

func main() {
	NewCLI().Run()
}

// CLI is the Cobra-based command-line interface.
type CLI struct {
	root *cobra.Command
}

// NewCLI sets up the CLI.
func NewCLI() *CLI { _ = "STUB: not implemented"; return nil }

// we'll output them ourselves in Run()

// generate generates manifests in a directory.
func (cli *CLI) generate(dir string, groups int, multiVersion string, prometheus bool) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gosec

// Run runs the CLI.
func (cli *CLI) Run() { _ = "STUB: not implemented"; return }
