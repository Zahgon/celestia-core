package main

import (
	"fmt"
	"os"
	"time"

	"github.com/cometbft/cometbft/config"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cometbft/cometbft/p2p"
)

const defaultE2EDelayedPreCommitTimeout = 20 * time.Millisecond

var logger = log.NewTMLogger(log.NewSyncWriter(os.Stdout))

// main is the binary entrypoint.
func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %v <configfile>", os.Args[0])
		return
	}
	configFile := ""
	if len(os.Args) == 2 {
		configFile = os.Args[1]
	}

	if err := run(configFile); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}

// run runs the application - basically like main() with error handling.
func run(configFile string) error { _ = "STUB: not implemented"; return nil }

// Start remote signer (must start before node if running builtin).

// Start app server.

// Apparently there's no way to wait for the server, so we just sleep

// startApp starts the application server, listening for connections from CometBFT.
func startApp(cfg *Config) error { _ = "STUB: not implemented"; return nil }

// startNode starts a CometBFT node running the application directly. It assumes the CometBFT
// configuration is in $CMTHOME/config/cometbft.toml.
//
// FIXME There is no way to simply load the configuration from a file, so we need to pull in Viper.
func startNode(cfg *Config) error { _ = "STUB: not implemented"; return nil }

func startLightClient(cfg *Config) error { _ = "STUB: not implemented"; return nil }

// If necessary adjust global WriteTimeout to ensure it's greater than
// TimeoutBroadcastTxCommit.
// See https://github.com/tendermint/tendermint/issues/3435

// Error starting or closing listener:

// startSigner starts a signer server connecting to the given endpoint.
func startSigner(cfg *Config) error { _ = "STUB: not implemented"; return nil }

func setupNode() (*config.Config, log.Logger, *p2p.NodeKey, error) {
	_ = "STUB: not implemented"
	return nil, *new(log.Logger), nil, nil
}

// rpcEndpoints takes a list of persistent peers and splits them into a list of rpc endpoints
// using 26657 as the port number
func rpcEndpoints(peers string) []string { _ = "STUB: not implemented"; return nil }

// use RPC port instead
