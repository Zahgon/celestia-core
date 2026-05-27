package commands

import (
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"

	dbm "github.com/cometbft/cometbft-db"
)

// LightCmd represents the base command when called without any subcommands
var LightCmd = &cobra.Command{
	Use:   "light [chainID]",
	Short: "Run a light client proxy server, verifying CometBFT rpc",
	Long: `Run a light client proxy server, verifying CometBFT rpc.

All calls that can be tracked back to a block header by a proof
will be verified before passing them back to the caller. Other than
that, it will present the same interface as a full CometBFT node.

Furthermore to the chainID, a fresh instance of a light client will
need a primary RPC address, a trusted hash and height and witness RPC addresses
(if not using sequential verification). To restart the node, thereafter
only the chainID is required.

When /abci_query is called, the Merkle key path format is:

	/{store name}/{key}

Please verify with your application that this Merkle key format is used (true
for applications built w/ Cosmos SDK).
`,
	RunE: runProxy,
	Args: cobra.ExactArgs(1),
	Example: `light cosmoshub-3 -p http://52.57.29.196:26657 -w http://public-seed-node.cosmoshub.certus.one:26657
	--height 962118 --hash 28B97BE9F6DE51AC69F70E0B7BFD7E5C9CD1A595B7DC31AFF27C50D4948020CD`,
}

var (
	listenAddr         string
	primaryAddr        string
	witnessAddrsJoined string
	chainID            string
	home               string
	maxOpenConnections int

	sequential     bool
	trustingPeriod time.Duration
	trustedHeight  int64
	trustedHash    []byte
	trustLevelStr  string

	verbose bool

	primaryKey   = []byte("primary")
	witnessesKey = []byte("witnesses")
)

func init() {
	LightCmd.Flags().StringVar(&listenAddr, "laddr", "tcp://localhost:8888",
		"serve the proxy on the given address")
	LightCmd.Flags().StringVarP(&primaryAddr, "primary", "p", "",
		"connect to a CometBFT node at this address")
	LightCmd.Flags().StringVarP(&witnessAddrsJoined, "witnesses", "w", "",
		"CometBFT nodes to cross-check the primary node, comma-separated")
	LightCmd.Flags().StringVar(&home, "home-dir", os.ExpandEnv(filepath.Join("$HOME", ".cometbft-light")),
		"specify the home directory")
	LightCmd.Flags().IntVar(
		&maxOpenConnections,
		"max-open-connections",
		900,
		"maximum number of simultaneous connections (including WebSocket).")
	LightCmd.Flags().DurationVar(&trustingPeriod, "trusting-period", 168*time.Hour,
		"trusting period that headers can be verified within. Should be significantly less than the unbonding period")
	LightCmd.Flags().Int64Var(&trustedHeight, "height", 1, "Trusted header's height")
	LightCmd.Flags().BytesHexVar(&trustedHash, "hash", []byte{}, "Trusted header's hash")
	LightCmd.Flags().BoolVar(&verbose, "verbose", false, "Verbose output")
	LightCmd.Flags().StringVar(&trustLevelStr, "trust-level", "1/3",
		"trust level. Must be between 1/3 and 3/3",
	)
	LightCmd.Flags().BoolVar(&sequential, "sequential", false,
		"sequential verification. Verify all headers sequentially as opposed to using skipping verification",
	)
}

func runProxy(_ *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	// Initialize logger.
	return nil
}

// check to see if we can start from an existing state

// fresh installation

// continue from latest state

// If necessary adjust global WriteTimeout to ensure it's greater than
// TimeoutBroadcastTxCommit.
// See https://github.com/tendermint/tendermint/issues/3435

// Stop upon receiving SIGTERM or CTRL-C.

// Error starting or closing listener:

func checkForExistingProviders(db dbm.DB) (string, []string, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func saveProviders(db dbm.DB, primaryAddr, witnessesAddrs string) error {
	_ = "STUB: not implemented"
	return nil
}
