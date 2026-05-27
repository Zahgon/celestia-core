package cli

import (
	"github.com/spf13/cobra"
)

// WriteConfigVals writes a toml file with the given values.
// It returns an error if writing was impossible.
func WriteConfigVals(dir string, vals map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

// RunWithArgs executes the given command with the specified command line args
// and environmental variables set. It returns any error returned from cmd.Execute()
func RunWithArgs(cmd Executable, args []string, env map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

// defer returns the environment back to normal

// set the args and env how we want them

// backup old value if there, to restore at end

// and finally run the command

// RunCaptureWithArgs executes the given command with the specified command
// line args and environmental variables set. It returns string fields
// representing output written to stdout and stderr, additionally any error
// from cmd.Execute() is also returned
func RunCaptureWithArgs(cmd Executable, args []string, env map[string]string) (stdout, stderr string, err error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

// keep backup of the real stdout

// restoring the real stdout

// copy the output in a separate goroutine so printing can't block indefinitely

// io.Copy will end when we call reader.Close() below
//nolint:errcheck //ignore error

// now run the command

// and grab the stdout to return

// NewCompletionCmd returns a cobra.Command that generates bash and zsh
// completion scripts for the given root command. If hidden is true, the
// command will not show up in the root command's list of available commands.
func NewCompletionCmd(rootCmd *cobra.Command, hidden bool) *cobra.Command {
	_ = "STUB: not implemented"
	return nil
}
