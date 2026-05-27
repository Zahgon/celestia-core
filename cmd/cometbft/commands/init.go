package commands

import (
	"github.com/spf13/cobra"

	cfg "github.com/cometbft/cometbft/config"
)

// InitFilesCmd initializes a fresh CometBFT instance.
var InitFilesCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize CometBFT",
	RunE:  initFiles,
}

func initFiles(*cobra.Command, []string) error { _ = "STUB: not implemented"; return nil }

func initFilesWithConfig(config *cfg.Config) error {
	_ = "STUB: not implemented"
	// private validator
	return nil
}

// genesis file
