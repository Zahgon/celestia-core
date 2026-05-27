package node

import (
	"context"
	"net"
	"net/http"

	"github.com/cometbft/cometbft/consensus/propagation"

	"github.com/grafana/pyroscope-go"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"google.golang.org/grpc"

	cfg "github.com/cometbft/cometbft/config"
	cs "github.com/cometbft/cometbft/consensus"
	"github.com/cometbft/cometbft/evidence"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cometbft/cometbft/libs/service"
	"github.com/cometbft/cometbft/libs/trace"
	mempl "github.com/cometbft/cometbft/mempool"
	"github.com/cometbft/cometbft/p2p"
	"github.com/cometbft/cometbft/p2p/pex"
	"github.com/cometbft/cometbft/proxy"
	rpccore "github.com/cometbft/cometbft/rpc/core"
	sm "github.com/cometbft/cometbft/state"
	"github.com/cometbft/cometbft/state/indexer"
	"github.com/cometbft/cometbft/state/txindex"
	"github.com/cometbft/cometbft/statesync"
	"github.com/cometbft/cometbft/store"
	"github.com/cometbft/cometbft/types"

	_ "net/http/pprof" //nolint: gosec
)

// Node is the highest level interface to a full CometBFT node.
// It includes all configuration information and running services.
type Node struct {
	service.BaseService

	// config
	config        *cfg.Config
	genesisDoc    *types.GenesisDoc   // initial validator set
	privValidator types.PrivValidator // local node's validator key

	// network
	transport   *p2p.MultiplexTransport
	sw          *p2p.Switch  // p2p connections
	addrBook    pex.AddrBook // known peers
	nodeInfo    p2p.NodeInfo
	nodeKey     *p2p.NodeKey // our node privkey
	isListening bool

	// services
	eventBus          *types.EventBus // pub/sub for services
	stateStore        sm.Store
	blockStore        *store.BlockStore // store the blockchain to disk
	bcReactor         p2p.Reactor       // for block-syncing
	mempoolReactor    p2p.Reactor       // for gossipping transactions
	mempool           mempl.Mempool
	stateSync         bool                    // whether the node should state sync on startup
	stateSyncReactor  *statesync.Reactor      // for hosting and restoring state sync snapshots
	stateSyncProvider statesync.StateProvider // provides state data for bootstrapping a node
	stateSyncGenesis  sm.State                // provides the genesis state for state sync
	consensusState    *cs.State               // latest consensus state
	consensusReactor  *cs.Reactor             // for participating in the consensus
	pexReactor        *pex.Reactor            // for exchanging peer addresses
	blockPropReactor  *propagation.Reactor    // the block propagation reactor. potentially nil is disabled.
	evidencePool      *evidence.Pool          // tracking evidence
	proxyApp          proxy.AppConns          // connection to the application
	rpcListeners      []net.Listener          // rpc servers
	txIndexer         txindex.TxIndexer
	blockIndexer      indexer.BlockIndexer
	indexerService    *txindex.IndexerService
	prometheusSrv     *http.Server
	pprofSrv          *http.Server

	// Celestia specific fields
	tracer            trace.Tracer
	pyroscopeProfiler *pyroscope.Profiler
	pyroscopeTracer   *sdktrace.TracerProvider
	privvalGRPCServer *grpc.Server
}

// Option sets a parameter for the node.
type Option func(*Node)

// CustomReactors allows you to add custom reactors (name -> p2p.Reactor) to
// the node's Switch.
//
// WARNING: using any name from the below list of the existing reactors will
// result in replacing it with the custom one.
//
//   - MEMPOOL
//   - BLOCKSYNC
//   - CONSENSUS
//   - EVIDENCE
//   - PEX
//   - STATESYNC
func CustomReactors(reactors map[string]p2p.Reactor) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// register the new channels to the nodeInfo
// NOTE: This is a bit messy now with the type casting but is
// cleaned up in the following version when NodeInfo is changed from
// and interface to a concrete type

// StateProvider overrides the state provider used by state sync to retrieve trusted app hashes and
// build a State object for bootstrapping the node.
// WARNING: this interface is considered unstable and subject to change.
func StateProvider(stateProvider statesync.StateProvider) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// BootstrapState synchronizes the stores with the application after state sync
// has been performed offline. It is expected that the block store and state
// store are empty at the time the function is called.
//
// If the block store is not empty, the function returns an error.
func BootstrapState(ctx context.Context, config *cfg.Config, dbProvider cfg.DBProvider, height uint64, appHash []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// BootstrapStateWithGenProvider synchronizes the stores with the application after state sync
// has been performed offline. It is expected that the block store and state
// store are empty at the time the function is called.
//
// If the block store is not empty, the function returns an error.
func BootstrapStateWithGenProvider(ctx context.Context, config *cfg.Config, dbProvider cfg.DBProvider, genProvider GenesisDocProvider, height uint64, appHash []byte) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Set the return value

// Set the return value

// Once the stores are bootstrapped, we need to set the height at which the node has finished
// statesyncing. This will allow the blocksync reactor to fetch blocks at a proper height.
// In case this operation fails, it is equivalent to a failure in  online state sync where the operator
// needs to manually delete the state and blockstores and rerun the bootstrapping process.

//------------------------------------------------------------------------------

// NewNode returns a new, ready to go, CometBFT Node.
func NewNode(config *cfg.Config,
	privValidator types.PrivValidator,
	nodeKey *p2p.NodeKey,
	clientCreator proxy.ClientCreator,
	genesisDocProvider GenesisDocProvider,
	dbProvider cfg.DBProvider,
	metricsProvider MetricsProvider,
	logger log.Logger,
	options ...Option,
) (*Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewNodeWithContext is cancellable version of NewNode.
func NewNodeWithContext(ctx context.Context,
	config *cfg.Config,
	privValidator types.PrivValidator,
	nodeKey *p2p.NodeKey,
	clientCreator proxy.ClientCreator,
	genesisDocProvider GenesisDocProvider,
	dbProvider cfg.DBProvider,
	metricsProvider MetricsProvider,
	logger log.Logger,
	options ...Option,
) (*Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create the proxyApp and establish connections to the ABCI app (consensus, mempool, query).

// EventBus and IndexerService must be started before the handshake because
// we might need to index the txs of the replayed block as this might not have happened
// when the node stopped last time (i.e. the node stopped after it saved the block
// but before it indexed the txs)

// create an optional tracer client to collect trace data.

// If an address is provided, listen on the socket for a connection from an
// external signing process.

// FIXME: we should start services inside OnStart

// Determine whether we should attempt state sync.

// Create the handshaker, which calls RequestInfo, sets the AppVersion on the state,
// and replays any blocks as necessary to sync CometBFT with the app.

// Reload the state. It will have the Version.Consensus.App set by the
// Handshake, and may have other modifications as well (ie. depending on
// what happened during block replay).

// Determine whether we should do block sync. This must happen after the handshake, since the
// app may modify the validator set, specifying ourself as the only validator.

// make block executor for consensus and blocksync reactors to execute blocks

// Don't start block sync if we're doing a state sync first.

// reduce the max block size to avoid overloading the legacy block prop mechanism

// Set up state sync reactor, and schedule a sync if requested.
// FIXME The way we do phased startups (e.g. replay -> block sync -> consensus) is very messy,
// we should clean this whole thing up. See:
// https://github.com/tendermint/tendermint/issues/4644

// Optionally, start the pex reactor
//
// TODO:
//
// We need to set Seeds and PersistentPeers on the switch,
// since it needs to be able to use these (and their DNS names)
// even if the PEX is off. We can include the DNS name in the NetAddress,
// but it would still be nice to have a clear list of the current "PersistentPeers"
// somewhere that we can return with net_info.
//
// If PEX is on, it should handle dialing the seeds. Otherwise the switch does it.
// Note we currently use the addrBook regardless at least for AddOurAddress

// Add private IDs to addrbook to block those peers being added

// Shouldn't be necessary, but need a way to pass the genesis state

// OnStart starts the Node. It implements service.Service.
func (n *Node) OnStart() error { _ = "STUB: not implemented"; return nil }

// run pprof server if it is enabled

// begin prometheus metrics gathering if it is enabled

// Start the RPC server before the P2P server
// so we can eg. receive txs for the first block

// Start the gRPC PrivValidator server if configured.

// Start the transport.

// Start the switch (the P2P server).

// Always connect to persistent peers

// Run state sync

// OnStop stops the Node. It implements service.Service.
func (n *Node) OnStop() { _ = "STUB: not implemented"; return }

// first stop the non-reactor services

// now stop the reactors

// finally stop the listeners / external services

// Error from closing listeners, or context timeout:

// ConfigureRPC makes sure RPC has all the objects it needs to operate.
func (n *Node) ConfigureRPC() (*rpccore.Environment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (n *Node) startRPC() ([]net.Listener, error) { _ = "STUB: not implemented"; return nil, nil }

// If necessary adjust global WriteTimeout to ensure it's greater than
// TimeoutBroadcastTxCommit.
// See https://github.com/tendermint/tendermint/issues/3435

// we may expose the rpc over both a unix and tcp socket

// we expose a simplified api over grpc for convenience to app devs

// NOTE: GRPCMaxOpenConnections is used, not MaxOpenConnections

// If necessary adjust global WriteTimeout to ensure it's greater than
// TimeoutBroadcastTxCommit.
// See https://github.com/tendermint/tendermint/issues/3435

//nolint:staticcheck // SA1019: core_grpc.StartGRPCClient is deprecated: A new gRPC API will be introduced after v0.38.

// startPrometheusServer starts a Prometheus HTTP server, listening for metrics
// collectors on addr.
func (n *Node) startPrometheusServer() *http.Server { _ = "STUB: not implemented"; return nil }

// Error starting or closing listener:

// starts a ppro
func (n *Node) startPprofServer() *http.Server { _ = "STUB: not implemented"; return nil }

// Error starting or closing listener:

// Switch returns the Node's Switch.
func (n *Node) Switch() *p2p.Switch {
	_ = "STUB: not implemented"

	// BlockStore returns the Node's BlockStore.
	return nil
}

func (n *Node) BlockStore() *store.BlockStore { _ = "STUB: not implemented"; return nil }

// ConsensusReactor returns the Node's ConsensusReactor.
func (n *Node) ConsensusReactor() *cs.Reactor { _ = "STUB: not implemented"; return nil }

// MempoolReactor returns the Node's mempool reactor.
func (n *Node) MempoolReactor() p2p.Reactor {
	_ = "STUB: not implemented"
	return *

	// Mempool returns the Node's mempool.
	new(p2p.Reactor)
}

func (n *Node) Mempool() mempl.Mempool {
	_ = "STUB: not implemented"

	// PEXReactor returns the Node's PEXReactor. It returns nil if PEX is disabled.
	return *new(mempl.Mempool)
}

func (n *Node) PEXReactor() *pex.Reactor { _ = "STUB: not implemented"; return nil }

// EvidencePool returns the Node's EvidencePool.
func (n *Node) EvidencePool() *evidence.Pool { _ = "STUB: not implemented"; return nil }

// EventBus returns the Node's EventBus.
func (n *Node) EventBus() *types.EventBus {
	_ = "STUB: not implemented"

	// PrivValidator returns the Node's PrivValidator.
	// XXX: for convenience only!
	return nil
}

func (n *Node) PrivValidator() types.PrivValidator {
	_ = "STUB: not implemented"
	return *

	// GenesisDoc returns the Node's GenesisDoc.
	new(types.PrivValidator)
}

func (n *Node) GenesisDoc() *types.GenesisDoc { _ = "STUB: not implemented"; return nil }

// ProxyApp returns the Node's AppConns, representing its connections to the ABCI application.
func (n *Node) ProxyApp() proxy.AppConns {
	_ = "STUB: not implemented"

	// Config returns the Node's config.
	return *new(proxy.AppConns)
}

func (n *Node) Config() *cfg.Config {
	_ = "STUB: not implemented"

	// ------------------------------------------------------------------------------
	return nil
}

func (n *Node) Listeners() []string { _ = "STUB: not implemented"; return nil }

func (n *Node) IsListening() bool { _ = "STUB: not implemented"; return false }

// NodeInfo returns the Node's Info from the Switch.
func (n *Node) NodeInfo() p2p.NodeInfo { _ = "STUB: not implemented"; return *new(p2p.NodeInfo) }

func makeNodeInfo(
	config *cfg.Config,
	nodeKey *p2p.NodeKey,
	txIndexer txindex.TxIndexer,
	genDoc *types.GenesisDoc,
	state sm.State,
	softwareVersion string,
) (p2p.DefaultNodeInfo, error) {
	_ = "STUB: not implemented"
	return *new(p2p.DefaultNodeInfo), nil
}

// global
