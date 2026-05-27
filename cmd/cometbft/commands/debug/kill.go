package debug

import (
	"github.com/spf13/cobra"
)

var killCmd = &cobra.Command{
	Use:   "kill [pid] [compressed-output-file]",
	Short: "Kill a CometBFT process while aggregating and packaging debugging data",
	Long: `Kill a CometBFT process while also aggregating CometBFT process data
such as the latest node state, including consensus and networking state,
go-routine state, and the node's WAL and config information. This aggregated data
is packaged into a compressed archive.

Example:
$ cometbft debug 34255 /path/to/cmt-debug.zip`,
	Args: cobra.ExactArgs(2),
	RunE: killCmdHandler,
}

func killCmdHandler(_ *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }

// Create a temporary directory which will contain all the state dumps and
// relevant files and directories that will be compressed into a file.

// killProc attempts to kill the CometBFT process with a given PID with an
// ABORT signal which should result in a goroutine stacktrace. The PID's STDERR
// is tailed and piped to a file under the directory dir. An error is returned
// if the output file cannot be created or the tail command cannot be started.
// An error is not returned if any subsequent syscall fails.
func killProc(pid int, dir string) error {
	_ = "STUB: not implemented"
	// pipe STDERR output from tailing the CometBFT process to a file
	//
	// NOTE: This will only work on UNIX systems.
	return nil
}

//nolint: gosec

// kill the underlying CometBFT process and subsequent tailing process

// Killing the CometBFT process with the '-ABRT|-6' signal will result in
// a goroutine stacktrace.

// allow some time to allow the CometBFT process to be killed
//
// TODO: We should 'wait' for a kill to succeed (e.g. poll for PID until it
// cannot be found). Regardless, this should be ample time.

// only return an error not invoked by a manual kill
