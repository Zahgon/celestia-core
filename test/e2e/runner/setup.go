package main

import (
	"github.com/cometbft/cometbft/config"
	e2e "github.com/cometbft/cometbft/test/e2e/pkg"
	"github.com/cometbft/cometbft/test/e2e/pkg/infra"
	"github.com/cometbft/cometbft/types"
)

const (
	AppAddressTCP  = "tcp://127.0.0.1:30000"
	AppAddressUNIX = "unix:///var/run/app.sock"

	PrivvalAddressTCP     = "tcp://0.0.0.0:27559"
	PrivvalAddressUNIX    = "unix:///var/run/privval.sock"
	PrivvalKeyFile        = "config/priv_validator_key.json"
	PrivvalStateFile      = "data/priv_validator_state.json"
	PrivvalDummyKeyFile   = "config/dummy_validator_key.json"
	PrivvalDummyStateFile = "data/dummy_validator_state.json"
)

// Setup sets up the testnet configuration.
func Setup(testnet *e2e.Testnet, infp infra.Provider) error { _ = "STUB: not implemented"; return nil }

// light clients don't need an app directory

// panics

//nolint:gosec

// stop early if a light client

// Set up a dummy validator. CometBFT requires a file PV even when not used, so we
// give it a dummy such that it will fail if it actually tries to use it.

// MakeGenesis generates a genesis document.
func MakeGenesis(testnet *e2e.Testnet) (types.GenesisDoc, error) {
	_ = "STUB: not implemented"
	return *new(types.GenesisDoc), nil
}

// set the app version to 1

// The validator set will be sorted internally by CometBFT ranked by power,
// but we sort it here as well so that all genesis files are identical.

// MakeConfig generates a CometBFT config for a node.
func MakeConfig(node *e2e.Node) (*config.Config, error) { _ = "STUB: not implemented"; return nil, nil }

// The default tx indexer is "null"; override to "kv" so e2e tests that
// query the deprecated tx/tx_search endpoints continue to work.

// CometBFT errors if it does not have a privval key set up, regardless of whether
// it's actually needed (e.g. for remote KMS or non-validators). We set up a dummy
// key here by default, and use the real key for actual validators that should use
// the file privval.

// Don't need to do anything, since we're using a dummy privval key by default.

// Set consensus configuration for propagation reactor

// MakeAppConfig generates an ABCI application config for a node.
func MakeAppConfig(node *e2e.Node) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UpdateConfigStateSync updates the state sync config for a node.
func UpdateConfigStateSync(node *e2e.Node, height int64, hash []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// FIXME Apparently there's no function to simply load a config file without
// involving the entire Viper apparatus, so we'll just resort to regexps.

//nolint:gosec
