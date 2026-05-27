package main

import (
	e2e "github.com/cometbft/cometbft/test/e2e/pkg"
)

// Cleanup removes the Docker Compose containers and testnet directory.
func Cleanup(testnet *e2e.Testnet) error { _ = "STUB: not implemented"; return nil }

// cleanupDocker removes all E2E resources (with label e2e=True), regardless
// of testnet.
func cleanupDocker() error { _ = "STUB: not implemented"; return nil }

// GNU xargs requires the -r flag to not run when input is empty, macOS
// does this by default. Ugly, but works.

// cleanupDir cleans up a testnet directory
func cleanupDir(dir string) error { _ = "STUB: not implemented"; return nil }

// On Linux, some local files in the volume will be owned by root since CometBFT
// runs as root inside the container, so we need to clean them up from within a
// container running as root too.

// depending on how docker is setup, this can fail. If it fails, then the user needs to remove excess data.
