package main

import (
	"math/rand"

	e2e "github.com/cometbft/cometbft/test/e2e/pkg"
)

var (
	// testnetCombinations defines global testnet options, where we generate a
	// separate testnet for each combination (Cartesian product) of options.
	testnetCombinations = map[string][]interface{}{
		"topology":      {"single", "quad", "large"},
		"initialHeight": {0, 1000},
		"initialState": {
			map[string]string{},
			map[string]string{"initial01": "a", "initial02": "b", "initial03": "c"},
		},
		"validators": {"genesis", "initchain"},
	}
	nodeVersions = weightedChoice{
		"": 2,
	}

	// The following specify randomly chosen values for testnet nodes.
	nodeDatabases = uniformChoice{"pebbledb"}
	ipv6          = uniformChoice{false, true}
	// FIXME: grpc disabled due to https://github.com/tendermint/tendermint/issues/5439
	nodeABCIProtocols     = uniformChoice{"unix", "tcp", "builtin", "builtin_connsync"} // "grpc"
	nodePrivvalProtocols  = uniformChoice{"file", "unix", "tcp"}
	nodeBlockSyncs        = uniformChoice{"v0"} // "v2"
	nodeStateSyncs        = uniformChoice{false, true}
	nodePersistIntervals  = uniformChoice{0, 1, 5}
	nodeSnapshotIntervals = uniformChoice{0, 3}
	nodeRetainBlocks      = uniformChoice{
		0,
		2 * int(e2e.EvidenceAgeHeight),
		4 * int(e2e.EvidenceAgeHeight),
	}
	evidence          = uniformChoice{0, 1, 10, 20, 200}
	abciDelays        = uniformChoice{"none", "small", "large"}
	nodePerturbations = probSetChoice{
		"disconnect": 0.1,
		"pause":      0.1,
		"kill":       0.1,
		"restart":    0.1,
		"upgrade":    0.3,
	}
	lightNodePerturbations = probSetChoice{
		"upgrade": 0.3,
	}
	voteExtensionUpdateHeight = uniformChoice{int64(-1), int64(0), int64(1)} // -1: genesis, 0: InitChain, 1: (use offset)
	voteExtensionEnabled      = weightedChoice{true: 3, false: 1}
	voteExtensionHeightOffset = uniformChoice{int64(0), int64(10), int64(100)}
	mempoolVersion            = "cat"
)

type generateConfig struct {
	randSource   *rand.Rand
	outputDir    string
	multiVersion string
	prometheus   bool
}

// Generate generates random testnets using the given RNG.
func Generate(cfg *generateConfig) ([]e2e.Manifest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// generateTestnet generates a single testnet with the given options.
func generateTestnet(r *rand.Rand, opt map[string]interface{}, upgradeVersion string, prometheus bool) (e2e.Manifest, error) {
	_ = "STUB: not implemented"
	return *new(e2e.Manifest), nil
}

// FIXME Networks are kept small since large ones use too much CPU.

// First we generate seed nodes, starting at the initial height.

// Next, we generate validators. We make sure a BFT quorum of validators start
// at the initial height, and that we have two archive nodes. We also set up
// the initial validator set, and validator set updates for delayed nodes.

// Move validators to InitChain if specified.

// Finally, we generate random full nodes.

// We now set up peer discovery for nodes. Seed nodes are fully meshed with
// each other, while non-seed nodes either use a set of random seeds or a
// set of random peers that start before themselves.

// if the full node or validator is an ideal candidate, it is added as a light provider.
// There are at least two archive nodes so there should be at least two ideal candidates

// lastly, set up the light clients

// generateNode randomly generates a node, with some constraints to avoid
// generating invalid configurations. We do not set Seeds or PersistentPeers
// here, since we need to know the overall network topology and startup
// sequencing.
func generateNode(
	r *rand.Rand, mode e2e.Mode, startAt int64, forceArchive bool,
) *e2e.ManifestNode {
	_ = "STUB: not implemented"
	return nil
}

// If this node is forced to be an archive node, retain all blocks and
// enable state sync snapshotting.

// If a node which does not persist state also does not retain blocks, randomly
// choose to either persist state or retain all blocks.

// If either PersistInterval or SnapshotInterval are greater than RetainBlocks,
// expand the block retention time.

func generateLightNode(r *rand.Rand, startAt int64, providers []string) *e2e.ManifestNode {
	_ = "STUB: not implemented"
	return nil
}

func ptrUint64(i uint64) *uint64 {
	_ = "STUB: not implemented"

	// Parses strings like "v0.34.21:1,v0.34.22:2" to represent two versions
	// ("v0.34.21" and "v0.34.22") with weights of 1 and 2 respectively.
	// Versions may be specified as cometbft/e2e-node:v0.34.27-alpha.1:1 or
	// ghcr.io/informalsystems/tendermint:v0.34.26:1.
	// If only the tag and weight are specified, cometbft/e2e-node is assumed.
	// Also returns the last version in the list, which will be used for updates.
	return nil
}

func parseWeightedVersions(s string) (weightedChoice, string, error) {
	_ = "STUB: not implemented"
	return *new(weightedChoice), "", nil
}

// Extracts the latest release version from the given Git repository. Uses the
// current version of CometBFT to establish the "major" version
// currently in use.
func gitRepoLatestReleaseVersion(gitRepoDir string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func findLatestReleaseTag(baseVer string, tags []string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Build our version comparison string
// See https://github.com/Masterminds/semver#caret-range-comparisons-major for details

// Skip tags that are not valid semantic versions

// Skip pre-releases

// Skip versions that don't match our constraints

// No relevant latest version (will cause the generator to only use the tip
// of the current branch)

// Ensure the version string has a "v" prefix, because all CometBFT E2E
// node Docker images' versions have a "v" prefix.
