package commands

import (
	"github.com/spf13/cobra"
)

// InspectCmd is the command for starting an inspect server.
var InspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "Run an inspect server for investigating CometBFT state",
	Long: `
	inspect runs a subset of CometBFT's RPC endpoints that are useful for debugging
	issues with CometBFT.

	When the CometBFT detects inconsistent state, it will crash the
	CometBFT process. CometBFT will not start up while in this inconsistent state.
	The inspect command can be used to query the block and state store using CometBFT
	RPC calls to debug issues of inconsistent state.
	`,

	RunE: runInspect,
}

func init() {
	InspectCmd.Flags().
		String("rpc.laddr",
			config.RPC.ListenAddress, "RPC listenener address. Port required")
	InspectCmd.Flags().
		String("db-backend",
			config.DBBackend, "database backend: pebbledb")
	InspectCmd.Flags().
		String("db-dir", config.DBPath, "database directory")
}

func runInspect(cmd *cobra.Command, _ []string) error { _ = "STUB: not implemented"; return nil }
