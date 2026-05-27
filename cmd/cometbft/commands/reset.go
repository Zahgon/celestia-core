package commands

import (
	"github.com/spf13/cobra"

	"github.com/cometbft/cometbft/libs/log"
)

// ResetAllCmd removes the database of this CometBFT core
// instance.
var ResetAllCmd = &cobra.Command{
	Use:     "unsafe-reset-all",
	Aliases: []string{"unsafe_reset_all"},
	Short:   "(unsafe) Remove all the data and WAL, reset this node's validator to genesis state",
	RunE:    resetAllCmd,
}

var keepAddrBook bool

// ResetStateCmd removes the database of the specified CometBFT core instance.
var ResetStateCmd = &cobra.Command{
	Use:     "reset-state",
	Aliases: []string{"reset_state"},
	Short:   "Remove all the data and WAL",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		config, err = ParseConfig(cmd)
		if err != nil {
			return err
		}

		return resetState(config.DBDir(), logger)
	},
}

func init() {
	ResetAllCmd.Flags().BoolVar(&keepAddrBook, "keep-addr-book", false, "keep the address book intact")
}

// ResetPrivValidatorCmd resets the private validator files.
var ResetPrivValidatorCmd = &cobra.Command{
	Use:     "unsafe-reset-priv-validator",
	Aliases: []string{"unsafe_reset_priv_validator"},
	Short:   "(unsafe) Reset this node's validator to genesis state",
	RunE:    resetPrivValidator,
}

// XXX: this is totally unsafe.
// it's only suitable for testnets.
func resetAllCmd(cmd *cobra.Command, _ []string) (err error) { _ = "STUB: not implemented"; return nil }

// XXX: this is totally unsafe.
// it's only suitable for testnets.
func resetPrivValidator(cmd *cobra.Command, _ []string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// resetAll removes address book files plus all data, and resets the privValdiator data.
func resetAll(dbDir, addrBookFile, privValKeyFile, privValStateFile string, logger log.Logger) error {
	_ = "STUB: not implemented"
	return nil
}

// recreate the dbDir since the privVal state needs to live there

// resetState removes address book files plus all databases.
func resetState(dbDir string, logger log.Logger) error { _ = "STUB: not implemented"; return nil }

func resetFilePV(privValKeyFile, privValStateFile string, logger log.Logger) {
	_ = "STUB: not implemented"
	return
}

func removeAddrBook(addrBookFile string, logger log.Logger) { _ = "STUB: not implemented"; return }
