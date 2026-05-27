package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/cometbft/cometbft/libs/log"

	abcicli "github.com/cometbft/cometbft/abci/client"
	"github.com/cometbft/cometbft/abci/version"
	"github.com/cometbft/cometbft/proto/tendermint/crypto"
)

// client is a global variable so it can be reused by the console
var (
	client abcicli.Client
	logger log.Logger
)

// flags
var (
	// global
	flagAddress  string
	flagAbci     string
	flagVerbose  bool   // for the println output
	flagLogLevel string // for the logger

	// query
	flagPath   string
	flagHeight int
	flagProve  bool

	// kvstore
	flagPersist string
)

var RootCmd = &cobra.Command{
	Use:   "abci-cli",
	Short: "the ABCI CLI tool wraps an ABCI client",
	Long:  "the ABCI CLI tool wraps an ABCI client and is used for testing ABCI servers",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		switch cmd.Use {
		case "kvstore", "version", "help [command]":
			return nil
		}

		if logger == nil {
			allowLevel, err := log.AllowLevel(flagLogLevel)
			if err != nil {
				return err
			}
			logger = log.NewFilter(log.NewTMLogger(log.NewSyncWriter(os.Stdout)), allowLevel)
		}
		if client == nil {
			var err error
			client, err = abcicli.NewClient(flagAddress, flagAbci, false)
			if err != nil {
				return err
			}
			client.SetLogger(logger.With("module", "abci-client"))
			if err := client.Start(); err != nil {
				return err
			}
		}
		return nil
	},
}

// Structure for data passed to print response.
type response struct {
	// generic abci response
	Data   []byte
	Code   uint32
	Info   string
	Log    string
	Status int32

	Query *queryResponse
}

type queryResponse struct {
	Key      []byte
	Value    []byte
	Height   int64
	ProofOps *crypto.ProofOps
}

func Execute() error { _ = "STUB: not implemented"; return nil }

func addGlobalFlags() { _ = "STUB: not implemented"; return }

func addQueryFlags() { _ = "STUB: not implemented"; return }

func addKVStoreFlags() { _ = "STUB: not implemented"; return }

func addCommands() { _ = "STUB: not implemented"; return }

// examples

var batchCmd = &cobra.Command{
	Use:   "batch",
	Short: "run a batch of abci commands against an application",
	Long: `run a batch of abci commands against an application

This command is run by piping in a file containing a series of commands
you'd like to run:

    abci-cli batch < example.file

where example.file looks something like:

    check_tx 0x00
    check_tx 0xff
    finalize_block 0x00
    check_tx 0x00
    finalize_block 0x01 0x04 0xff
    info
`,
	Args: cobra.ExactArgs(0),
	RunE: cmdBatch,
}

var consoleCmd = &cobra.Command{
	Use:   "console",
	Short: "start an interactive ABCI console for multiple commands",
	Long: `start an interactive ABCI console for multiple commands

This command opens an interactive console for running any of the other commands
without opening a new connection each time
`,
	Args:      cobra.ExactArgs(0),
	ValidArgs: []string{"echo", "info", "finalize_block", "check_tx", "prepare_proposal", "process_proposal", "commit", "query"},
	RunE:      cmdConsole,
}

var echoCmd = &cobra.Command{
	Use:   "echo",
	Short: "have the application echo a message",
	Long:  "have the application echo a message",
	Args:  cobra.ExactArgs(1),
	RunE:  cmdEcho,
}

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "get some info about the application",
	Long:  "get some info about the application",
	Args:  cobra.ExactArgs(0),
	RunE:  cmdInfo,
}

var finalizeBlockCmd = &cobra.Command{
	Use:   "finalize_block",
	Short: "deliver a block of transactions to the application",
	Long:  "deliver a block of transactions to the application",
	Args:  cobra.MinimumNArgs(1),
	RunE:  cmdFinalizeBlock,
}

var checkTxCmd = &cobra.Command{
	Use:   "check_tx",
	Short: "validate a transaction",
	Long:  "validate a transaction",
	Args:  cobra.ExactArgs(1),
	RunE:  cmdCheckTx,
}

var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "commit the application state and return the Merkle root hash",
	Long:  "commit the application state and return the Merkle root hash",
	Args:  cobra.ExactArgs(0),
	RunE:  cmdCommit,
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print ABCI console version",
	Long:  "print ABCI console version",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println(version.Version)
		return nil
	},
}

var prepareProposalCmd = &cobra.Command{
	Use:   "prepare_proposal",
	Short: "prepare proposal",
	Long:  "prepare proposal",
	Args:  cobra.MinimumNArgs(0),
	RunE:  cmdPrepareProposal,
}

var processProposalCmd = &cobra.Command{
	Use:   "process_proposal",
	Short: "process proposal",
	Long:  "process proposal",
	Args:  cobra.MinimumNArgs(0),
	RunE:  cmdProcessProposal,
}

var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "query the application state",
	Long:  "query the application state",
	Args:  cobra.ExactArgs(1),
	RunE:  cmdQuery,
}

var kvstoreCmd = &cobra.Command{
	Use:   "kvstore",
	Short: "ABCI demo example",
	Long:  "ABCI demo example",
	Args:  cobra.ExactArgs(0),
	RunE:  cmdKVStore,
}

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "run integration tests",
	Long:  "run integration tests",
	Args:  cobra.ExactArgs(0),
	RunE:  cmdTest,
}

// Generates new Args array based off of previous call args to maintain flag persistence
func persistentArgs(line []byte) []string {
	_ = "STUB: not implemented"
	// generate the arguments to run from original os.Args
	// to maintain flag arguments
	return nil
}

// remove the previous command argument

// prevents introduction of extra space leading to argument parse errors

//--------------------------------------------------------------------------------

func compose(fs []func() error) error { _ = "STUB: not implemented"; return nil }

func cmdTest(cmd *cobra.Command, _ []string) error { _ = "STUB: not implemented"; return nil }

func cmdBatch(cmd *cobra.Command, _ []string) error { _ = "STUB: not implemented"; return nil }

func cmdConsole(cmd *cobra.Command, _ []string) error { _ = "STUB: not implemented"; return nil }

func muxOnCommands(cmd *cobra.Command, pArgs []string) error { _ = "STUB: not implemented"; return nil }

// TODO: this parsing is fragile

// check for flags

// if it has an equal, we can just skip

// if its a boolean, we can just skip

// otherwise, we need to skip the next one too

// append the actual arg

// for later print statements ...

func cmdUnimplemented(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

// Have the application echo a message
func cmdEcho(cmd *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }

// Get some info from the application
func cmdInfo(cmd *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }

const codeBad uint32 = 10

// Append new txs to application
func cmdFinalizeBlock(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

// Validate a tx
func cmdCheckTx(cmd *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }

// Get application Merkle root hash
func cmdCommit(cmd *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }

// Query application state
func cmdQuery(cmd *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }

func cmdPrepareProposal(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

// kvstore has to have this parameter in order not to reject a tx as the default value is 0

// CodeOK

func cmdProcessProposal(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

func cmdKVStore(*cobra.Command, []string) error { _ = "STUB: not implemented"; return nil }

// Create the application - in memory or persisted to disk

// Start the listener

// Stop upon receiving SIGTERM or CTRL-C.

// Cleanup

// Run forever.

//--------------------------------------------------------------------------------

func printResponse(cmd *cobra.Command, args []string, rsps ...response) {
	_ = "STUB: not implemented"
	return
}

// Always print the status code.

// Do no print this line when using the finalize_block command
// because the string comes out as gibberish

// NOTE: s is interpreted as a string unless prefixed with 0x
func stringOrHexToBytes(s string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
