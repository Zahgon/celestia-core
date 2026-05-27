package rpctest

import (
	abci "github.com/cometbft/cometbft/abci/types"

	cfg "github.com/cometbft/cometbft/config"
	nm "github.com/cometbft/cometbft/node"
	core_grpc "github.com/cometbft/cometbft/rpc/grpc"
)

// Options helps with specifying some parameters for our RPC testing for greater
// control.
type Options struct {
	suppressStdout  bool
	recreateConfig  bool
	maxReqBatchSize int

	// SpecificConfig will replace the global config if not nil
	SpecificConfig *cfg.Config
}

var (
	globalConfig   *cfg.Config
	defaultOptions = Options{
		suppressStdout: false,
		recreateConfig: false,
	}
)

func waitForRPC() { _ = "STUB: not implemented"; return }

func waitForGRPC() { _ = "STUB: not implemented"; return }

// f**ing long, but unique for each test
func makePathname() string {
	_ = "STUB: not implemented"
	// get path
	return ""
}

// fmt.Println(p)

func randPort() int { _ = "STUB: not implemented"; return 0 }

func makeAddrs() (string, string, string) { _ = "STUB: not implemented"; return "", "", "" }

func createConfig() *cfg.Config { _ = "STUB: not implemented"; return nil }

// and we use random ports to run in parallel

// GetConfig returns a config for the test cases as a singleton
func GetConfig(forceCreate ...bool) *cfg.Config { _ = "STUB: not implemented"; return nil }

func GetGRPCClient() core_grpc.BroadcastAPIClient {
	_ = "STUB: not implemented"
	return *new(core_grpc.BroadcastAPIClient)
}

//nolint:staticcheck

// StartTendermint starts a test CometBFT server in a go routine and returns when it is initialized
func StartTendermint(app abci.Application, opts ...func(*Options)) *nm.Node {
	_ = "STUB: not implemented"
	return nil
}

// wait for rpc

// StopTendermint stops a test CometBFT server, waits until it's stopped and
// cleans up test/config files.
func StopTendermint(node *nm.Node) { _ = "STUB: not implemented"; return }

// NewTendermint creates a new CometBFT server and sleeps forever
func NewTendermint(app abci.Application, opts *Options) *nm.Node {
	_ = "STUB: not implemented"
	// Create & start node
	return nil
}

// SuppressStdout is an option that tries to make sure the RPC test CometBFT
// node doesn't log anything to stdout.
func SuppressStdout(o *Options) { _ = "STUB: not implemented"; return }

// RecreateConfig instructs the RPC test to recreate the configuration each
// time, instead of treating it as a global singleton.
func RecreateConfig(o *Options) { _ = "STUB: not implemented"; return }

// MaxReqBatchSize is an option to limit the maximum number of requests per batch.
func MaxReqBatchSize(o *Options) { _ = "STUB: not implemented"; return }

func GetBlockAPIClient() (core_grpc.BlockAPIClient, error) {
	_ = "STUB: not implemented"
	return *new(core_grpc.BlockAPIClient), nil
}

func GetBlobstreamAPIClient() (core_grpc.BlobstreamAPIClient, error) {
	_ = "STUB: not implemented"
	return *new(core_grpc.BlobstreamAPIClient), nil
}
