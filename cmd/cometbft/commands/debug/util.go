package debug

import (
	cfg "github.com/cometbft/cometbft/config"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
)

// dumpStatus gets node status state dump from the CometBFT RPC and writes it
// to file. It returns an error upon failure.
func dumpStatus(rpc *rpchttp.HTTP, dir, filename string) error {
	_ = "STUB: not implemented"
	return nil
}

// dumpNetInfo gets network information state dump from the CometBFT RPC and
// writes it to file. It returns an error upon failure.
func dumpNetInfo(rpc *rpchttp.HTTP, dir, filename string) error {
	_ = "STUB: not implemented"
	return nil
}

// dumpConsensusState gets consensus state dump from the CometBFT RPC and
// writes it to file. It returns an error upon failure.
func dumpConsensusState(rpc *rpchttp.HTTP, dir, filename string) error {
	_ = "STUB: not implemented"
	return nil
}

// copyWAL copies the CometBFT node's WAL file. It returns an error if the
// WAL file cannot be read or copied.
func copyWAL(conf *cfg.Config, dir string) error { _ = "STUB: not implemented"; return nil }

// copyConfig copies the CometBFT node's config file. It returns an error if
// the config file cannot be read or copied.
func copyConfig(home, dir string) error { _ = "STUB: not implemented"; return nil }

func dumpProfile(dir, addr, profile string, debug int) error { _ = "STUB: not implemented"; return nil }

//nolint:gosec,nolintlint
