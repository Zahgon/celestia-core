package commands

import (
	"github.com/spf13/cobra"

	cfg "github.com/cometbft/cometbft/config"
	nm "github.com/cometbft/cometbft/node"
)

var (
	genesisHash []byte
)

// AddNodeFlags exposes some common configuration options on the command-line
// These are exposed for convenience of commands embedding a CometBFT node
func AddNodeFlags(cmd *cobra.Command) {
	_ = "STUB: not implemented"
	// bind flags
	return
}

// priv val flags

// node flags

// abci flags

// rpc flags

// p2p flags

// consensus flags

// db flags

// NewRunNodeCmd returns the command that allows the CLI to start a node.
// It can be used with a custom PrivValidator and in-process ABCI application.
func NewRunNodeCmd(nodeProvider nm.Provider) *cobra.Command { _ = "STUB: not implemented"; return nil }

// Stop upon receiving SIGTERM or CTRL-C.

// Run forever.

func checkGenesisHash(config *cfg.Config) error { _ = "STUB: not implemented"; return nil }

// Calculate SHA-256 hash of the genesis file.

// Compare with the flag.
