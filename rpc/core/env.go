package core

import (
	"time"

	cfg "github.com/cometbft/cometbft/config"
	"github.com/cometbft/cometbft/crypto"
	"github.com/cometbft/cometbft/libs/log"
	mempl "github.com/cometbft/cometbft/mempool"
	"github.com/cometbft/cometbft/p2p"
	"github.com/cometbft/cometbft/proxy"
	sm "github.com/cometbft/cometbft/state"
	"github.com/cometbft/cometbft/state/indexer"
	"github.com/cometbft/cometbft/state/txindex"
	"github.com/cometbft/cometbft/types"
)

const (
	// see README
	defaultPerPage = 30
	maxPerPage     = 100

	// SubscribeTimeout is the maximum time we wait to subscribe for an event.
	// must be less than the server's write timeout (see rpcserver.DefaultConfig)
	SubscribeTimeout = 5 * time.Second

	// genesisChunkSize is the maximum size, in bytes, of each
	// chunk in the genesis structure for the chunked API
	genesisChunkSize = 16 * 1024 * 1024 // 16
)

//----------------------------------------------
// These interfaces are used by RPC and must be thread safe

type Consensus interface {
	GetState() sm.State
	GetValidators() (int64, []*types.Validator)
	GetLastHeight() int64
	GetRoundStateJSON() ([]byte, error)
	GetRoundStateSimpleJSON() ([]byte, error)
}

type transport interface {
	Listeners() []string
	IsListening() bool
	NodeInfo() p2p.NodeInfo
}

type peers interface {
	AddPersistentPeers([]string) error
	AddUnconditionalPeerIDs([]string) error
	AddPrivatePeerIDs([]string) error
	DialPeersAsync([]string) error
	Peers() p2p.IPeerSet
}

type consensusReactor interface {
	WaitSync() bool
}

// ----------------------------------------------
// Environment contains objects and interfaces used by the RPC. It is expected
// to be setup once during startup.
type Environment struct {
	// external, thread safe interfaces
	ProxyAppQuery   proxy.AppConnQuery
	ProxyAppMempool proxy.AppConnMempool

	// interfaces defined in types and above
	StateStore       sm.Store
	BlockStore       sm.BlockStore
	EvidencePool     sm.EvidencePool
	ConsensusState   Consensus
	ConsensusReactor consensusReactor
	P2PPeers         peers
	P2PTransport     transport

	// objects
	PubKey       crypto.PubKey
	GenDoc       *types.GenesisDoc // cache the genesis structure
	TxIndexer    txindex.TxIndexer
	BlockIndexer indexer.BlockIndexer
	EventBus     *types.EventBus // thread safe
	Mempool      mempl.Mempool

	Logger log.Logger

	Config cfg.RPCConfig

	// cache of chunked genesis data.
	genChunks []string
}

//----------------------------------------------

func validatePage(pagePtr *int, perPage, totalCount int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// no page parameter

// one page (even if it's empty)

func (env *Environment) validatePerPage(perPagePtr *int) int { _ = "STUB: not implemented"; return 0 }

// no per_page parameter

func (env *Environment) validateUnconfirmedTxsPerPage(perPagePtr *int) int {
	_ = "STUB: not implemented"
	return 0
	// no per_page parameter
}

// InitGenesisChunks configures the environment and should be called on service
// startup.
func (env *Environment) InitGenesisChunks() error { _ = "STUB: not implemented"; return nil }

func validateSkipCount(page, perPage int) int { _ = "STUB: not implemented"; return 0 }

// latestHeight can be either latest committed or uncommitted (+1) height.
func (env *Environment) getHeight(latestHeight int64, heightPtr *int64) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (env *Environment) latestUncommittedHeight() int64 { _ = "STUB: not implemented"; return 0 }
