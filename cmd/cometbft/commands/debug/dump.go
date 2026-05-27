package debug

import (
	"github.com/spf13/cobra"

	cfg "github.com/cometbft/cometbft/config"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
)

var dumpCmd = &cobra.Command{
	Use:   "dump [output-directory]",
	Short: "Continuously poll a CometBFT process and dump debugging data into a single location",
	Long: `Continuously poll a CometBFT process and dump debugging data into a single
location at a specified frequency. At each frequency interval, an archived and compressed
file will contain node debugging information including the goroutine and heap profiles
if enabled.`,
	Args: cobra.ExactArgs(1),
	RunE: dumpCmdHandler,
}

func init() {
	dumpCmd.Flags().UintVar(
		&frequency,
		flagFrequency,
		30,
		"the frequency (seconds) in which to poll, aggregate and dump CometBFT debug data",
	)

	dumpCmd.Flags().StringVar(
		&profAddr,
		flagProfAddr,
		"",
		"the profiling server address (<host>:<port>)",
	)
}

func dumpCmdHandler(_ *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }

func dumpDebugData(outDir string, conf *cfg.Config, rpc *rpchttp.HTTP) {
	_ = "STUB: not implemented"
	return
}
