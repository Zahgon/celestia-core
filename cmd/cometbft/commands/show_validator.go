package commands

import (
	"github.com/spf13/cobra"
)

// ShowValidatorCmd adds capabilities for showing the validator info.
var ShowValidatorCmd = &cobra.Command{
	Use:     "show-validator",
	Aliases: []string{"show_validator"},
	Short:   "Show this node's validator info",
	RunE:    showValidator,
}

func showValidator(*cobra.Command, []string) error { _ = "STUB: not implemented"; return nil }
