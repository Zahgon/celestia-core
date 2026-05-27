package cli

import (
	"github.com/spf13/cobra"
)

const (
	HomeFlag     = "home"
	TraceFlag    = "trace"
	OutputFlag   = "output"
	EncodingFlag = "encoding"
)

// Executable is the minimal interface to *corba.Command, so we can
// wrap if desired before the test
type Executable interface {
	Execute() error
}

// PrepareBaseCmd is meant for CometBFT and other servers
func PrepareBaseCmd(cmd *cobra.Command, envPrefix, defaultHome string) Executor {
	_ = "STUB: not implemented"
	return *new(Executor)
}

// PrepareMainCmd is meant for client side libs that want some more flags
//
// This adds --encoding (hex, btc, base64) and --output (text, json) to
// the command.  These only really make sense in interactive commands.
func PrepareMainCmd(cmd *cobra.Command, envPrefix, defaultHome string) Executor {
	_ = "STUB: not implemented"
	return *new(Executor)
}

// initEnv sets to use ENV variables if set.
func initEnv(prefix string) {
	_ = "STUB: not implemented"

	// env variables with TM prefix (eg. TM_ROOT)
	return
}

// This copies all variables like TMROOT to TM_ROOT,
// so we can support both formats for the user
func copyEnvVars(prefix string) { _ = "STUB: not implemented"; return }

// Executor wraps the cobra Command with a nicer Execute method
type Executor struct {
	*cobra.Command
	Exit func(int) // this is os.Exit by default, override in tests
}

type ExitCoder interface {
	ExitCode() int
}

// execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func (e Executor) Execute() error { _ = "STUB: not implemented"; return nil }

// return error code 1 by default, can override it with a special error type

type cobraCmdFunc func(cmd *cobra.Command, args []string) error

// Returns a single function that calls each argument function in sequence
// RunE, PreRunE, PersistentPreRunE, etc. all have this same signature
func concatCobraCmdFuncs(fs ...cobraCmdFunc) cobraCmdFunc {
	_ = "STUB: not implemented"
	return *new(cobraCmdFunc)
}

// Bind all flags and read the config into viper
func bindFlagsLoadViper(cmd *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	// cmd.Flags() includes flags from this command and all persistent flags from the parent
	return nil
}

// name of config file (without extension)
// search root directory
// search root directory /config

// If a config file is found, read it in.

// ignore not found error, return other errors

func validateOutput(_ *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	// validate output format
	return nil
}
