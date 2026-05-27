package commands

import (
	"github.com/spf13/cobra"
)

// GenNodeKeyCmd allows the generation of a node key. It prints node's ID to
// the standard output.
var GenNodeKeyCmd = &cobra.Command{
	Use:     "gen-node-key",
	Aliases: []string{"gen_node_key"},
	Short:   "Generate a node key for this node and print its ID",
	RunE:    genNodeKey,
}

func genNodeKey(*cobra.Command, []string) error { _ = "STUB: not implemented"; return nil }
