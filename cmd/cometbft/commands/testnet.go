package commands

import (
	"github.com/spf13/cobra"

	cfg "github.com/cometbft/cometbft/config"
)

var (
	nValidators    int
	nNonValidators int
	initialHeight  int64
	configFile     string
	outputDir      string
	nodeDirPrefix  string

	populatePersistentPeers bool
	hostnamePrefix          string
	hostnameSuffix          string
	startingIPAddress       string
	hostnames               []string
	p2pPort                 int
	randomMonikers          bool
)

const (
	nodeDirPerm = 0o755
)

func init() {
	TestnetFilesCmd.Flags().IntVar(&nValidators, "v", 4,
		"number of validators to initialize the testnet with")
	TestnetFilesCmd.Flags().StringVar(&configFile, "config", "",
		"config file to use (note some options may be overwritten)")
	TestnetFilesCmd.Flags().IntVar(&nNonValidators, "n", 0,
		"number of non-validators to initialize the testnet with")
	TestnetFilesCmd.Flags().StringVar(&outputDir, "o", "./mytestnet",
		"directory to store initialization data for the testnet")
	TestnetFilesCmd.Flags().StringVar(&nodeDirPrefix, "node-dir-prefix", "node",
		"prefix the directory name for each node with (node results in node0, node1, ...)")
	TestnetFilesCmd.Flags().Int64Var(&initialHeight, "initial-height", 0,
		"initial height of the first block")

	TestnetFilesCmd.Flags().BoolVar(&populatePersistentPeers, "populate-persistent-peers", true,
		"update config of each node with the list of persistent peers build using either"+
			" hostname-prefix or"+
			" starting-ip-address")
	TestnetFilesCmd.Flags().StringVar(&hostnamePrefix, "hostname-prefix", "node",
		"hostname prefix (\"node\" results in persistent peers list ID0@node0:26656, ID1@node1:26656, ...)")
	TestnetFilesCmd.Flags().StringVar(&hostnameSuffix, "hostname-suffix", "",
		"hostname suffix ("+
			"\".xyz.com\""+
			" results in persistent peers list ID0@node0.xyz.com:26656, ID1@node1.xyz.com:26656, ...)")
	TestnetFilesCmd.Flags().StringVar(&startingIPAddress, "starting-ip-address", "",
		"starting IP address ("+
			"\"192.168.0.1\""+
			" results in persistent peers list ID0@192.168.0.1:26656, ID1@192.168.0.2:26656, ...)")
	TestnetFilesCmd.Flags().StringArrayVar(&hostnames, "hostname", []string{},
		"manually override all hostnames of validators and non-validators (use --hostname multiple times for multiple hosts)")
	TestnetFilesCmd.Flags().IntVar(&p2pPort, "p2p-port", 26656,
		"P2P Port")
	TestnetFilesCmd.Flags().BoolVar(&randomMonikers, "random-monikers", false,
		"randomize the moniker for each generated node")
}

// TestnetFilesCmd allows initialisation of files for a CometBFT testnet.
var TestnetFilesCmd = &cobra.Command{
	Use:   "testnet",
	Short: "Initialize files for a CometBFT testnet",
	Long: `testnet will create "v" + "n" number of directories and populate each with
necessary files (private validator, genesis, config, etc.).

Note, strict routability for addresses is turned off in the config file.

Optionally, it will fill in persistent_peers list in config file using either hostnames or IPs.

Example:

	cometbft testnet --v 4 --o ./output --populate-persistent-peers --starting-ip-address 192.168.10.2
	`,
	RunE: testnetFiles,
}

func testnetFiles(*cobra.Command, []string) error { _ = "STUB: not implemented"; return nil }

// overwrite default config if set and valid

//nolint:staticcheck
//nolint:staticcheck

// Generate genesis doc from generated validators

// Write genesis file.

//nolint:staticcheck

// Gather persistent peer addresses.

// Overwrite default config.

func hostnameOrIP(i int) string { _ = "STUB: not implemented"; return "" }

func persistentPeersString(config *cfg.Config) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func moniker(i int) string { _ = "STUB: not implemented"; return "" }

func randomMoniker() string { _ = "STUB: not implemented"; return "" }
