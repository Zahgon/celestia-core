package node

import (
	"context"
	"time"

	"github.com/cometbft/cometbft/consensus/propagation"

	_ "net/http/pprof" //nolint: gosec // securely exposed on separate, optional port

	dbm "github.com/cometbft/cometbft-db"

	"github.com/cometbft/cometbft/blocksync"
	cfg "github.com/cometbft/cometbft/config"
	cs "github.com/cometbft/cometbft/consensus"
	"github.com/cometbft/cometbft/crypto"
	"github.com/cometbft/cometbft/evidence"
	"github.com/cometbft/cometbft/statesync"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cometbft/cometbft/libs/trace"
	mempl "github.com/cometbft/cometbft/mempool"
	"github.com/cometbft/cometbft/p2p"
	"github.com/cometbft/cometbft/p2p/pex"
	"github.com/cometbft/cometbft/proxy"
	sm "github.com/cometbft/cometbft/state"
	"github.com/cometbft/cometbft/state/indexer"
	"github.com/cometbft/cometbft/state/txindex"
	"github.com/cometbft/cometbft/store"
	"github.com/cometbft/cometbft/types"

	_ "github.com/lib/pq" // provide the psql db driver
)

const readHeaderTimeout = 10 * time.Second

// GenesisDocProvider returns a GenesisDoc.
// It allows the GenesisDoc to be pulled from sources other than the
// filesystem, for instance from a distributed key-value store cluster.
type GenesisDocProvider func() (*types.GenesisDoc, error)

// DefaultGenesisDocProviderFunc returns a GenesisDocProvider that loads
// the GenesisDoc from the config.GenesisFile() on the filesystem.
func DefaultGenesisDocProviderFunc(config *cfg.Config) GenesisDocProvider {
	_ = "STUB: not implemented"
	return *new(GenesisDocProvider)
}

// Provider takes a config and a logger and returns a ready to go Node.
type Provider func(*cfg.Config, log.Logger) (*Node, error)

// DefaultNewNode returns a CometBFT node with default settings for the
// PrivValidator, ClientCreator, GenesisDoc, and DBProvider.
// It implements NodeProvider.
func DefaultNewNode(config *cfg.Config, logger log.Logger) (*Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MetricsProvider returns a consensus, p2p and mempool Metrics.
type MetricsProvider func(chainID string) (*cs.Metrics, *p2p.Metrics, *mempl.Metrics, *sm.Metrics, *proxy.Metrics, *blocksync.Metrics, *statesync.Metrics)

// DefaultMetricsProvider returns Metrics build using Prometheus client library
// if Prometheus is enabled. Otherwise, it returns no-op Metrics.
func DefaultMetricsProvider(config *cfg.InstrumentationConfig) MetricsProvider {
	_ = "STUB: not implemented"
	return *new(MetricsProvider)
}

type blockSyncReactor interface {
	SwitchToBlockSync(sm.State) error
}

//------------------------------------------------------------------------------

// initDBs opens or creates the blockstore and state databases.
// If config.BlockstoreDir() differs from config.DBDir(), users must manually
// migrate their existing blockstore data before changing this configuration.
// No automatic migration is performed to prevent accidental data inconsistency.
func initDBs(config *cfg.Config, dbProvider cfg.DBProvider, logger log.Logger) (blockStore *store.BlockStore, stateDB dbm.DB, err error) {
	_ = "STUB: not implemented"
	return nil, *new(dbm.DB), nil
}

func createAndStartProxyAppConns(clientCreator proxy.ClientCreator, logger log.Logger, metrics *proxy.Metrics) (proxy.AppConns, error) {
	_ = "STUB: not implemented"
	return *new(proxy.AppConns), nil
}

func createAndStartEventBus(logger log.Logger) (*types.EventBus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createAndStartIndexerService(
	config *cfg.Config,
	chainID string,
	dbProvider cfg.DBProvider,
	eventBus *types.EventBus,
	logger log.Logger,
) (*txindex.IndexerService, txindex.TxIndexer, indexer.BlockIndexer, error) {
	_ = "STUB: not implemented"
	return nil, *new(txindex.TxIndexer), *new(indexer.BlockIndexer), nil
}

func doHandshake(
	ctx context.Context,
	stateStore sm.Store,
	state sm.State,
	blockStore sm.BlockStore,
	genDoc *types.GenesisDoc,
	eventBus types.BlockEventPublisher,
	proxyApp proxy.AppConns,
	consensusLogger log.Logger,
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func logNodeStartupInfo(state sm.State, pubKey crypto.PubKey, logger, consensusLogger log.Logger) {
	_ = "STUB: not implemented"
	// Log the version info.
	return
}

// If the state and software differ in block version, at least log it.

// Log whether this node is a validator or an observer

func onlyValidatorIsUs(state sm.State, localAddr crypto.Address) bool {
	_ = "STUB: not implemented"
	return false
}

// createMempoolAndMempoolReactor creates a mempool and a mempool reactor based on the config.
func createMempoolAndMempoolReactor(
	config *cfg.Config,
	proxyApp proxy.AppConns,
	state sm.State,
	memplMetrics *mempl.Metrics,
	logger log.Logger,
	traceClient trace.Tracer,
) (mempl.Mempool, p2p.Reactor) {
	_ = "STUB: not implemented"
	return *new(mempl.Mempool), *new(p2p.Reactor)
}

// Strictly speaking, there's no need to have a `mempl.NopMempoolReactor`, but
// adding it leads to a cleaner code.

// TODO: find a more polite way of handling this error

func createEvidenceReactor(config *cfg.Config, dbProvider cfg.DBProvider,
	stateStore sm.Store, blockStore *store.BlockStore, logger log.Logger,
) (*evidence.Reactor, *evidence.Pool, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func createBlocksyncReactor(config *cfg.Config,
	state sm.State,
	blockExec *sm.BlockExecutor,
	blockStore *store.BlockStore,
	blockSync bool,
	localAddr crypto.Address,
	logger log.Logger,
	metrics *blocksync.Metrics,
	offlineStateSyncHeight int64,
	tracer trace.Tracer,
) (bcReactor p2p.Reactor, err error) {
	_ = "STUB: not implemented"
	return *new(p2p.Reactor), nil
}

func createConsensusReactor(config *cfg.Config,
	state sm.State,
	blockExec *sm.BlockExecutor,
	blockStore sm.BlockStore,
	mempool mempl.Mempool,
	evidencePool *evidence.Pool,
	privValidator types.PrivValidator,
	csMetrics *cs.Metrics,
	propagator propagation.Propagator,
	waitSync bool,
	eventBus *types.EventBus,
	consensusLogger log.Logger,
	offlineStateSyncHeight int64,
	traceClient trace.Tracer,
) (*cs.Reactor, *cs.State) {
	_ = "STUB: not implemented"
	return nil, nil
}

// services which will be publishing and/or subscribing for messages (events)
// consensusReactor will set it on consensusState and blockExecutor

func createTransport(
	config *cfg.Config,
	nodeInfo p2p.NodeInfo,
	nodeKey *p2p.NodeKey,
	proxyApp proxy.AppConns,
	traceClient trace.Tracer,
) (
	*p2p.MultiplexTransport,
	[]p2p.PeerFilterFunc,
) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Filter peers by addr or pubkey with an ABCI query.
// If the query return code is OK, add peer.

// ABCI query for address filtering.

// ABCI query for ID filtering.

// Filter peers that don't support the new block propagation or CAT mempool.

// Check for legacy block propagation

// Check for CAT mempool

// Limit the number of incoming connections.

func createSwitch(config *cfg.Config,
	transport p2p.Transport,
	p2pMetrics *p2p.Metrics,
	peerFilters []p2p.PeerFilterFunc,
	mempoolReactor p2p.Reactor,
	bcReactor p2p.Reactor,
	stateSyncReactor *statesync.Reactor,
	consensusReactor *cs.Reactor,
	evidenceReactor *evidence.Reactor,
	propagationReactor *propagation.Reactor,
	nodeInfo p2p.NodeInfo,
	nodeKey *p2p.NodeKey,
	p2pLogger log.Logger,
	traceClient trace.Tracer,
) *p2p.Switch {
	_ = "STUB: not implemented"
	return nil
}

func createAddrBookAndSetOnSwitch(config *cfg.Config, sw *p2p.Switch,
	p2pLogger log.Logger, nodeKey *p2p.NodeKey,
) (pex.AddrBook, error) {
	_ = "STUB: not implemented"
	return *new(pex.AddrBook), nil
}

// Add ourselves to addrbook to prevent dialing ourselves

func createPEXReactorAndAddToSwitch(addrBook pex.AddrBook, config *cfg.Config,
	sw *p2p.Switch, logger log.Logger,
) *pex.Reactor {
	_ = "STUB: not implemented"
	// TODO persistent peers ? so we can have their DNS addrs saved
	return nil
}

// See consensus/reactor.go: blocksToContributeToBecomeGoodPeer 10000
// blocks assuming 10s blocks ~ 28 hours.
// TODO (melekes): make it dynamic based on the actual block latencies
// from the live network.
// https://github.com/tendermint/tendermint/issues/3523

// startStateSync starts an asynchronous state sync process, then switches to block sync mode.
func startStateSync(
	ssR *statesync.Reactor,
	bcR blockSyncReactor,
	stateProvider statesync.StateProvider,
	config *cfg.StateSyncConfig,
	stateStore sm.Store,
	blockStore *store.BlockStore,
	state sm.State,
) error {
	_ = "STUB: not implemented"
	return nil
}

//------------------------------------------------------------------------------

var genesisDocKey = []byte("genesisDoc")

// LoadStateFromDBOrGenesisDocProvider attempts to load the state from the
// database, or creates one using the given genesisDocProvider. On success this also
// returns the genesis doc loaded through the given provider.
func LoadStateFromDBOrGenesisDocProvider(
	stateDB dbm.DB,
	genesisDocProvider GenesisDocProvider,
) (sm.State, *types.GenesisDoc, error) {
	_ = "STUB: not implemented"
	// Get genesis doc
	return *new(sm.State), nil, nil
}

// save genesis doc to prevent a certain class of user errors (e.g. when it
// was changed, accidentally or not). Also good for audit trail.

// panics if failed to unmarshal bytes
func loadGenesisDoc(db dbm.DB) (*types.GenesisDoc, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// panics if failed to marshal the given genesis document
func saveGenesisDoc(db dbm.DB, genDoc *types.GenesisDoc) error {
	_ = "STUB: not implemented"
	return nil
}

func createAndStartPrivValidatorSocketClient(
	listenAddr,
	chainID string,
	logger log.Logger,
	tracer trace.Tracer,
) (types.PrivValidator, error) {
	_ = "STUB: not implemented"
	return *new(types.PrivValidator), nil
}

// try to get a pubkey from private validate first time

// 50 * 100ms = 5s total

// splitAndTrimEmpty slices s into all subslices separated by sep and returns a
// slice of the string s with all leading and trailing Unicode code points
// contained in cutset removed. If sep is empty, SplitAndTrim splits after each
// UTF-8 sequence. First part is equivalent to strings.SplitN with a count of
// -1.  also filter out empty strings, only return non-empty strings.
func splitAndTrimEmpty(s, sep, cutset string) []string { _ = "STUB: not implemented"; return nil }
