package config

import (
	"path/filepath"
	"regexp"
	"time"
)

const (
	// FuzzModeDrop is a mode in which we randomly drop reads/writes, connections or sleep
	FuzzModeDrop = iota
	// FuzzModeDelay is a mode in which we randomly sleep
	FuzzModeDelay

	// LogFormatPlain is a format for colored text
	LogFormatPlain = "plain"
	// LogFormatJSON is a format for json output
	LogFormatJSON = "json"

	// DefaultLogLevel defines a default log level as INFO.
	DefaultLogLevel = "info"

	DefaultTendermintDir = ".cometbft"
	DefaultConfigDir     = "config"
	DefaultDataDir       = "data"

	DefaultConfigFileName  = "config.toml"
	DefaultGenesisJSONName = "genesis.json"

	DefaultPrivValKeyName   = "priv_validator_key.json"
	DefaultPrivValStateName = "priv_validator_state.json"

	DefaultNodeKeyName  = "node_key.json"
	DefaultAddrBookName = "addrbook.json"

	MempoolTypeNop = "nop"
	MempoolTypeCAT = "cat"

	// DefaultMaxPersistentStickyPeers caps how many persistent peers are guaranteed
	// in the SeenTx broadcast set per signer when the config field is unset.
	DefaultMaxPersistentStickyPeers = 4
)

// NOTE: Most of the structs & relevant comments + the
// default configuration options were used to manually
// generate the config.toml. Please reflect any changes
// made here in the defaultConfigTemplate constant in
// config/toml.go
// NOTE: libs/cli must know to look in the config dir!
var (
	defaultConfigFilePath   = filepath.Join(DefaultConfigDir, DefaultConfigFileName)
	defaultGenesisJSONPath  = filepath.Join(DefaultConfigDir, DefaultGenesisJSONName)
	defaultPrivValKeyPath   = filepath.Join(DefaultConfigDir, DefaultPrivValKeyName)
	defaultPrivValStatePath = filepath.Join(DefaultDataDir, DefaultPrivValStateName)

	defaultNodeKeyPath  = filepath.Join(DefaultConfigDir, DefaultNodeKeyName)
	defaultAddrBookPath = filepath.Join(DefaultConfigDir, DefaultAddrBookName)

	minSubscriptionBufferSize     = 100
	defaultSubscriptionBufferSize = 200

	// taken from https://semver.org/
	semverRegexp = regexp.MustCompile(`^(?P<major>0|[1-9]\d*)\.(?P<minor>0|[1-9]\d*)\.(?P<patch>0|[1-9]\d*)(?:-(?P<prerelease>(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+(?P<buildmetadata>[0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`)

	// DefaultTracingTables is a list of tables that are used for storing traces.
	// This global var is filled by an init function in the schema package. This
	// allows for the schema package to contain all the relevant logic while
	// avoiding import cycles.
	DefaultTracingTables = ""
)

// Config defines the top level configuration for a CometBFT node
type Config struct {
	// Top level options use an anonymous struct
	BaseConfig `mapstructure:",squash"`

	// Options for services
	RPC             *RPCConfig             `mapstructure:"rpc"`
	P2P             *P2PConfig             `mapstructure:"p2p"`
	Mempool         *MempoolConfig         `mapstructure:"mempool"`
	StateSync       *StateSyncConfig       `mapstructure:"statesync"`
	BlockSync       *BlockSyncConfig       `mapstructure:"blocksync"`
	Consensus       *ConsensusConfig       `mapstructure:"consensus"`
	Storage         *StorageConfig         `mapstructure:"storage"`
	TxIndex         *TxIndexConfig         `mapstructure:"tx_index"`
	Instrumentation *InstrumentationConfig `mapstructure:"instrumentation"`
}

// DefaultConfig returns a default configuration for a CometBFT node
func DefaultConfig() *Config { _ = "STUB: not implemented"; return nil }

// TestConfig returns a configuration that can be used for testing
func TestConfig() *Config { _ = "STUB: not implemented"; return nil }

// SetRoot sets the RootDir for all Config structs
func (cfg *Config) SetRoot(root string) *Config { _ = "STUB: not implemented"; return nil }

//nolint:staticcheck

// ValidateBasic performs basic validation (checking param bounds, etc.) and
// returns an error if any check fails.
func (cfg *Config) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// CheckDeprecated returns any deprecation warnings. These are printed to the operator on startup
func (cfg *Config) CheckDeprecated() []string { _ = "STUB: not implemented"; return nil }

// Deprecated Mempool configs

// Deprecated Consensus configs

// CreateEmptyBlocks defaults to true, so we check if it's explicitly set to any value

//-----------------------------------------------------------------------------
// BaseConfig

// BaseConfig defines the base configuration for a CometBFT node
type BaseConfig struct {

	// The version of the CometBFT binary that created
	// or last modified the config file
	Version string `mapstructure:"version"`

	// The root directory for all data.
	// This should be set in viper so it can unmarshal into this struct
	RootDir string `mapstructure:"home"`

	// TCP or UNIX socket address of the ABCI application,
	// or the name of an ABCI application compiled in with the CometBFT binary
	ProxyApp string `mapstructure:"proxy_app"`

	// A custom human readable name for this node
	Moniker string `mapstructure:"moniker"`

	// Database backend: pebbledb
	// * pebbledb (github.com/cockroachdb/pebble)
	//   - pure go
	//   - fast and stable
	DBBackend string `mapstructure:"db_backend"`

	// Database directory
	DBPath string `mapstructure:"db_dir"`

	// Blockstore directory. If not set, defaults to DBPath
	BlockstorePath string `mapstructure:"blockstore_dir"`

	// Output level for logging
	LogLevel string `mapstructure:"log_level"`

	// Output format: 'plain' (colored text) or 'json'
	LogFormat string `mapstructure:"log_format"`

	// Path to the JSON file containing the initial validator set and other meta data
	Genesis string `mapstructure:"genesis_file"`

	// Path to the JSON file containing the private key to use as a validator in the consensus protocol
	PrivValidatorKey string `mapstructure:"priv_validator_key_file"`

	// Path to the JSON file containing the last sign state of a validator
	PrivValidatorState string `mapstructure:"priv_validator_state_file"`

	// TCP or UNIX socket address for CometBFT to listen on for
	// connections from an external PrivValidator process
	PrivValidatorListenAddr string `mapstructure:"priv_validator_laddr"`

	// gRPC address for the PrivValidator server.
	// If set, the node exposes its PrivValidator over gRPC on this address,
	// allowing external services (fiber server) to request signatures.
	PrivValidatorGRPCListenAddr string `mapstructure:"priv_validator_grpc_laddr"`

	// A JSON file containing the private key to use for p2p authenticated encryption
	NodeKey string `mapstructure:"node_key_file"`

	// Mechanism to connect to the ABCI application: socket | grpc
	ABCI string `mapstructure:"abci"`

	// If true, query the ABCI app on connecting to a new peer
	// so the app can decide if we should keep the connection or not
	FilterPeers bool `mapstructure:"filter_peers"` // false
}

// DefaultBaseConfig returns a default base configuration for a CometBFT node
func DefaultBaseConfig() BaseConfig { _ = "STUB: not implemented"; return *new(BaseConfig) }

// TestBaseConfig returns a base configuration for testing a CometBFT node
func TestBaseConfig() BaseConfig { _ = "STUB: not implemented"; return *new(BaseConfig) }

// GenesisFile returns the full path to the genesis.json file
func (cfg BaseConfig) GenesisFile() string { _ = "STUB: not implemented"; return "" }

// PrivValidatorKeyFile returns the full path to the priv_validator_key.json file
func (cfg BaseConfig) PrivValidatorKeyFile() string { _ = "STUB: not implemented"; return "" }

// PrivValidatorFile returns the full path to the priv_validator_state.json file
func (cfg BaseConfig) PrivValidatorStateFile() string { _ = "STUB: not implemented"; return "" }

// NodeKeyFile returns the full path to the node_key.json file
func (cfg BaseConfig) NodeKeyFile() string { _ = "STUB: not implemented"; return "" }

// DBDir returns the full path to the database directory
func (cfg BaseConfig) DBDir() string { _ = "STUB: not implemented"; return "" }

// BlockstoreDir returns the full path to the blockstore directory
func (cfg BaseConfig) BlockstoreDir() string { _ = "STUB: not implemented"; return "" }

// ValidateBasic performs basic validation (checking param bounds, etc.) and
// returns an error if any check fails.
func (cfg BaseConfig) ValidateBasic() error {
	_ = "STUB: not implemented"
	// version on old config files aren't set so we can't expect it
	// always to exist
	return nil
}

//-----------------------------------------------------------------------------
// RPCConfig

// RPCConfig defines the configuration options for the CometBFT RPC server
type RPCConfig struct {
	RootDir string `mapstructure:"home"`

	// TCP or UNIX socket address for the RPC server to listen on
	ListenAddress string `mapstructure:"laddr"`

	// A list of origins a cross-domain request can be executed from.
	// If the special '*' value is present in the list, all origins will be allowed.
	// An origin may contain a wildcard (*) to replace 0 or more characters (i.e.: http://*.domain.com).
	// Only one wildcard can be used per origin.
	CORSAllowedOrigins []string `mapstructure:"cors_allowed_origins"`

	// A list of methods the client is allowed to use with cross-domain requests.
	CORSAllowedMethods []string `mapstructure:"cors_allowed_methods"`

	// A list of non simple headers the client is allowed to use with cross-domain requests.
	CORSAllowedHeaders []string `mapstructure:"cors_allowed_headers"`

	// TCP or UNIX socket address for the gRPC server to listen on
	GRPCListenAddress string `mapstructure:"grpc_laddr"`

	// Maximum number of simultaneous connections.
	// Does not include RPC (HTTP&WebSocket) connections. See max_open_connections
	// If you want to accept a larger number than the default, make sure
	// you increase your OS limits.
	// 0 - unlimited.
	GRPCMaxOpenConnections int `mapstructure:"grpc_max_open_connections"`

	// Activate unsafe RPC commands like /dial_persistent_peers and /unsafe_flush_mempool
	Unsafe bool `mapstructure:"unsafe"`

	// Maximum number of simultaneous connections (including WebSocket).
	// Does not include gRPC connections. See grpc_max_open_connections
	// If you want to accept a larger number than the default, make sure
	// you increase your OS limits.
	// 0 - unlimited.
	// Should be < {ulimit -Sn} - {MaxNumInboundPeers} - {MaxNumOutboundPeers} - {N of wal, db and other open files}
	// 1024 - 40 - 10 - 50 = 924 = ~900
	MaxOpenConnections int `mapstructure:"max_open_connections"`

	// Maximum number of unique clientIDs that can /subscribe
	// If you're using /broadcast_tx_commit, set to the estimated maximum number
	// of broadcast_tx_commit calls per block.
	MaxSubscriptionClients int `mapstructure:"max_subscription_clients"`

	// Maximum number of unique queries a given client can /subscribe to
	// If you're using GRPC (or Local RPC client) and /broadcast_tx_commit, set
	// to the estimated maximum number of broadcast_tx_commit calls per block.
	MaxSubscriptionsPerClient int `mapstructure:"max_subscriptions_per_client"`

	// The number of events that can be buffered per subscription before
	// returning `ErrOutOfCapacity`.
	SubscriptionBufferSize int `mapstructure:"experimental_subscription_buffer_size"`

	// The maximum number of responses that can be buffered per WebSocket
	// client. If clients cannot read from the WebSocket endpoint fast enough,
	// they will be disconnected, so increasing this parameter may reduce the
	// chances of them being disconnected (but will cause the node to use more
	// memory).
	//
	// Must be at least the same as `SubscriptionBufferSize`, otherwise
	// connections may be dropped unnecessarily.
	WebSocketWriteBufferSize int `mapstructure:"experimental_websocket_write_buffer_size"`

	// If a WebSocket client cannot read fast enough, at present we may
	// silently drop events instead of generating an error or disconnecting the
	// client.
	//
	// Enabling this parameter will cause the WebSocket connection to be closed
	// instead if it cannot read fast enough, allowing for greater
	// predictability in subscription behavior.
	CloseOnSlowClient bool `mapstructure:"experimental_close_on_slow_client"`

	// How long to wait for a tx to be committed during /broadcast_tx_commit
	// WARNING: Using a value larger than 10s will result in increasing the
	// global HTTP write timeout, which applies to all connections and endpoints.
	// See https://github.com/tendermint/tendermint/issues/3435
	TimeoutBroadcastTxCommit time.Duration `mapstructure:"timeout_broadcast_tx_commit"`

	// Maximum number of requests that can be sent in a batch
	// https://www.jsonrpc.org/specification#batch
	MaxRequestBatchSize int `mapstructure:"max_request_batch_size"`

	// Maximum size of request body, in bytes
	MaxBodyBytes int64 `mapstructure:"max_body_bytes"`

	// Maximum size of request header, in bytes
	MaxHeaderBytes int `mapstructure:"max_header_bytes"`

	// The path to a file containing certificate that is used to create the HTTPS server.
	// Might be either absolute path or path related to CometBFT's config directory.
	//
	// If the certificate is signed by a certificate authority,
	// the certFile should be the concatenation of the server's certificate, any intermediates,
	// and the CA's certificate.
	//
	// NOTE: both tls_cert_file and tls_key_file must be present for CometBFT to create HTTPS server.
	// Otherwise, HTTP server is run.
	TLSCertFile string `mapstructure:"tls_cert_file"`

	// The path to a file containing matching private key that is used to create the HTTPS server.
	// Might be either absolute path or path related to CometBFT's config directory.
	//
	// NOTE: both tls_cert_file and tls_key_file must be present for CometBFT to create HTTPS server.
	// Otherwise, HTTP server is run.
	TLSKeyFile string `mapstructure:"tls_key_file"`

	// pprof listen address (https://golang.org/pkg/net/http/pprof)
	// FIXME: This should be moved under the instrumentation section
	PprofListenAddress string `mapstructure:"pprof_laddr"`
}

// DefaultRPCConfig returns a default configuration for the RPC server
func DefaultRPCConfig() *RPCConfig { _ = "STUB: not implemented"; return nil }

// maximum requests in a JSON-RPC batch request
// 1MB
// same as the net/http default

// TestRPCConfig returns a configuration for testing the RPC server
func TestRPCConfig() *RPCConfig { _ = "STUB: not implemented"; return nil }

// ValidateBasic performs basic validation (checking param bounds, etc.) and
// returns an error if any check fails.
func (cfg *RPCConfig) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// IsCorsEnabled returns true if cross-origin resource sharing is enabled.
func (cfg *RPCConfig) IsCorsEnabled() bool { _ = "STUB: not implemented"; return false }

func (cfg *RPCConfig) IsPprofEnabled() bool { _ = "STUB: not implemented"; return false }

func (cfg RPCConfig) KeyFile() string { _ = "STUB: not implemented"; return "" }

func (cfg RPCConfig) CertFile() string { _ = "STUB: not implemented"; return "" }

func (cfg RPCConfig) IsTLSEnabled() bool { _ = "STUB: not implemented"; return false }

//-----------------------------------------------------------------------------
// P2PConfig

// P2PConfig defines the configuration options for the CometBFT peer-to-peer networking layer
type P2PConfig struct {
	RootDir string `mapstructure:"home"`

	// Address to listen for incoming connections
	ListenAddress string `mapstructure:"laddr"`

	// Address to advertise to peers for them to dial
	ExternalAddress string `mapstructure:"external_address"`

	// Comma separated list of seed nodes to connect to
	// We only use these if we can’t connect to peers in the addrbook
	Seeds string `mapstructure:"seeds"`

	// Comma separated list of nodes to keep persistent connections to
	PersistentPeers string `mapstructure:"persistent_peers"`

	// Path to address book
	AddrBook string `mapstructure:"addr_book_file"`

	// Set true for strict address routability rules
	// Set false for private or local networks
	AddrBookStrict bool `mapstructure:"addr_book_strict"`

	// Maximum number of inbound peers
	MaxNumInboundPeers int `mapstructure:"max_num_inbound_peers"`

	// Maximum number of outbound peers to connect to, excluding persistent peers
	MaxNumOutboundPeers int `mapstructure:"max_num_outbound_peers"`

	// List of node IDs, to which a connection will be (re)established ignoring any existing limits
	UnconditionalPeerIDs string `mapstructure:"unconditional_peer_ids"`

	// Maximum pause when redialing a persistent peer (if zero, exponential backoff is used)
	PersistentPeersMaxDialPeriod time.Duration `mapstructure:"persistent_peers_max_dial_period"`

	// Time to wait before flushing messages out on the connection
	FlushThrottleTimeout time.Duration `mapstructure:"flush_throttle_timeout"`

	// Maximum size of a message packet payload, in bytes
	MaxPacketMsgPayloadSize int `mapstructure:"max_packet_msg_payload_size"`

	// Rate at which packets can be sent, in bytes/second
	SendRate int64 `mapstructure:"send_rate"`

	// Rate at which packets can be received, in bytes/second
	RecvRate int64 `mapstructure:"recv_rate"`

	// Set true to enable the peer-exchange reactor
	PexReactor bool `mapstructure:"pex"`

	// Seed mode, in which node constantly crawls the network and looks for
	// peers. If another node asks it for addresses, it responds and disconnects.
	//
	// Does not work if the peer-exchange reactor is disabled.
	SeedMode bool `mapstructure:"seed_mode"`

	// Comma separated list of peer IDs to keep private (will not be gossiped to
	// other peers)
	PrivatePeerIDs string `mapstructure:"private_peer_ids"`

	// Toggle to disable guard against peers connecting from the same ip.
	AllowDuplicateIP bool `mapstructure:"allow_duplicate_ip"`

	// Peer connection configuration.
	HandshakeTimeout time.Duration `mapstructure:"handshake_timeout"`
	DialTimeout      time.Duration `mapstructure:"dial_timeout"`

	// Testing params.
	// Force dial to fail
	TestDialFail bool `mapstructure:"test_dial_fail"`
	// Fuzz connection
	TestFuzz       bool            `mapstructure:"test_fuzz"`
	TestFuzzConfig *FuzzConnConfig `mapstructure:"test_fuzz_config"`
}

// DefaultP2PConfig returns a default configuration for the peer-to-peer layer
func DefaultP2PConfig() *P2PConfig { _ = "STUB: not implemented"; return nil }

// 1 kB
// 5 mB/s
// 5 mB/s

// TestP2PConfig returns a configuration for testing the peer-to-peer layer
func TestP2PConfig() *P2PConfig { _ = "STUB: not implemented"; return nil }

// AddrBookFile returns the full path to the address book
func (cfg *P2PConfig) AddrBookFile() string { _ = "STUB: not implemented"; return "" }

// ValidateBasic performs basic validation (checking param bounds, etc.) and
// returns an error if any check fails.
func (cfg *P2PConfig) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// FuzzConnConfig is a FuzzedConnection configuration.
type FuzzConnConfig struct {
	Mode         int
	MaxDelay     time.Duration
	ProbDropRW   float64
	ProbDropConn float64
	ProbSleep    float64
}

// DefaultFuzzConnConfig returns the default config.
func DefaultFuzzConnConfig() *FuzzConnConfig { _ = "STUB: not implemented"; return nil }

//-----------------------------------------------------------------------------
// MempoolConfig

// MempoolConfig defines the configuration options for the CometBFT mempool.
type MempoolConfig struct {
	// The type of mempool for this node to use.
	//
	//  Possible types:
	//  - "cat"   : content addressable mempool (default)
	//  - "nop"   : nop-mempool (short for no operation; the ABCI app is
	//  responsible for storing, disseminating and proposing txs).
	//  "create_empty_blocks=false" is not supported.
	Type string `mapstructure:"type"`
	// RootDir is the root directory for all data. This should be configured via
	// the $CMTHOME env variable or --home cmd flag rather than overriding this
	// struct field.
	RootDir string `mapstructure:"home"`
	// Recheck (default: true) defines whether CometBFT should recheck the
	// validity for all remaining transaction in the mempool after a block.
	// Since a block affects the application state, some transactions in the
	// mempool may become invalid. If this does not apply to your application,
	// you can disable rechecking.
	Recheck bool `mapstructure:"recheck"`
	// RecheckTimeout is the time the application has during the rechecking process
	// to return CheckTx responses, once all requests have been sent. Responses that
	// arrive after the timeout expires are discarded. It only applies to
	// non-local ABCI clients and when recheck is enabled.
	//
	// The ideal value will strongly depend on the application. It could roughly be estimated as the
	// average size of the mempool multiplied by the average time it takes the application to validate one
	// transaction. We consider that the ABCI application runs in the same location as the CometBFT binary
	// so that the recheck duration is not affected by network delays when making requests and receiving responses.
	RecheckTimeout time.Duration `mapstructure:"recheck_timeout"`
	// Broadcast (default: true) defines whether the mempool should relay
	// transactions to other peers. Setting this to false will stop the mempool
	// from relaying transactions to other peers until they are included in a
	// block. In other words, if Broadcast is disabled, only the peer you send
	// the tx to will see it until it is included in a block.
	Broadcast bool `mapstructure:"broadcast"`
	// WalPath (default: "") configures the location of the Write Ahead Log
	// (WAL) for the mempool. The WAL is disabled by default. To enable, set
	// WalPath to where you want the WAL to be written (e.g.
	// "data/mempool.wal").
	WalPath string `mapstructure:"wal_dir"`
	// Maximum number of transactions in the mempool
	Size int `mapstructure:"size"`
	// Limit the total size of all txs in the mempool.
	// This only accounts for raw transactions (e.g. given 1MB transactions and
	// max_txs_bytes=5MB, mempool will only accept 5 transactions).
	MaxTxsBytes int64 `mapstructure:"max_txs_bytes"`
	// Size of the cache (used to filter transactions we saw earlier) in transactions
	CacheSize int `mapstructure:"cache_size"`
	// Do not remove invalid transactions from the cache (default: false)
	// Set to true if it's not possible for any invalid transaction to become
	// valid again in the future.
	// Deprecated: KeepInvalidTxsInCache is deprecated and will be removed in a future version.
	KeepInvalidTxsInCache bool `mapstructure:"keep-invalid-txs-in-cache"`
	// Maximum size of a single transaction
	// NOTE: the max size of a tx transmitted over the network is {max_tx_bytes}.
	// Deprecated: MaxTxBytes is deprecated and will be removed in a future version.
	MaxTxBytes int `mapstructure:"max_tx_bytes"`
	// Maximum size of a batch of transactions to send to a peer
	// Including space needed by encoding (one varint per transaction).
	// XXX: Unused due to https://github.com/tendermint/tendermint/issues/5796
	MaxBatchBytes int `mapstructure:"max_batch_bytes"`
	// Experimental parameters to limit gossiping txs to up to the specified number of peers.
	// We use two independent upper values for persistent and non-persistent peers.
	// Unconditional peers are not affected by this feature.
	// If we are connected to more than the specified number of persistent peers, only send txs to
	// ExperimentalMaxGossipConnectionsToPersistentPeers of them. If one of those
	// persistent peers disconnects, activate another persistent peer.
	// Similarly for non-persistent peers, with an upper limit of
	// ExperimentalMaxGossipConnectionsToNonPersistentPeers.
	// If set to 0, the feature is disabled for the corresponding group of peers, that is, the
	// number of active connections to that group of peers is not bounded.
	// For non-persistent peers, if enabled, a value of 10 is recommended based on experimental
	// performance results using the default P2P configuration.
	ExperimentalMaxGossipConnectionsToPersistentPeers    int `mapstructure:"experimental_max_gossip_connections_to_persistent_peers"`
	ExperimentalMaxGossipConnectionsToNonPersistentPeers int `mapstructure:"experimental_max_gossip_connections_to_non_persistent_peers"`

	// MaxGossipDelay is the maximum allotted time that the reactor expects a transaction to
	// arrive before issuing a new request to a different peer
	// Only applicable to the v2 / CAT mempool
	// Default is 200ms
	// Deprecated: MaxGossipDelay is deprecated and will be removed in a future version.
	MaxGossipDelay time.Duration `mapstructure:"max-gossip-delay"`

	// TTLDuration, if non-zero, defines the maximum amount of time a transaction
	// can exist for in the mempool.
	//
	// Note, if TTLNumBlocks is also defined, a transaction will be removed if it
	// has existed in the mempool at least TTLNumBlocks number of blocks or if it's
	// insertion time into the mempool is beyond TTLDuration.
	// Deprecated: TTLDuration is deprecated and will be removed in a future version.
	TTLDuration time.Duration `mapstructure:"ttl-duration"`

	// TTLNumBlocks, if non-zero, defines the maximum number of blocks a transaction
	// can exist for in the mempool.
	//
	// Note, if TTLDuration is also defined, a transaction will be removed if it
	// has existed in the mempool at least TTLNumBlocks number of blocks or if
	// it's insertion time into the mempool is beyond TTLDuration.
	// Deprecated: TTLNumBlocks is deprecated and will be removed in a future version.
	TTLNumBlocks int64 `mapstructure:"ttl-num-blocks"`

	// MaxPersistentStickyPeers is the upper bound on persistent peers guaranteed
	// to receive SeenTx broadcasts per signer (added on top of the natural sticky
	// set, never displacing it). 0 falls back to the default. This key is omitted
	// from the generated config.toml; set it explicitly to override.
	MaxPersistentStickyPeers int `mapstructure:"max_persistent_sticky_peers"`
}

// DefaultMempoolConfig returns a default configuration for the CometBFT mempool
func DefaultMempoolConfig() *MempoolConfig { _ = "STUB: not implemented"; return nil }

// Each signature verification takes .5ms, Size reduced until we implement
// ABCI Recheck

// 1GB

// 1MB

// TestMempoolConfig returns a configuration for testing the CometBFT mempool
func TestMempoolConfig() *MempoolConfig { _ = "STUB: not implemented"; return nil }

// WalDir returns the full path to the mempool's write-ahead log
func (cfg *MempoolConfig) WalDir() string { _ = "STUB: not implemented"; return "" }

// WalEnabled returns true if the WAL is enabled.
func (cfg *MempoolConfig) WalEnabled() bool { _ = "STUB: not implemented"; return false }

// ValidateBasic performs basic validation (checking param bounds, etc.) and
// returns an error if any check fails.
func (cfg *MempoolConfig) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

//-----------------------------------------------------------------------------
// StateSyncConfig

// StateSyncConfig defines the configuration for the CometBFT state sync service
type StateSyncConfig struct {
	Enable              bool          `mapstructure:"enable"`
	TempDir             string        `mapstructure:"temp_dir"`
	RPCServers          []string      `mapstructure:"rpc_servers"`
	TrustPeriod         time.Duration `mapstructure:"trust_period"`
	TrustHeight         int64         `mapstructure:"trust_height"`
	TrustHash           string        `mapstructure:"trust_hash"`
	DiscoveryTime       time.Duration `mapstructure:"discovery_time"`
	ChunkRequestTimeout time.Duration `mapstructure:"chunk_request_timeout"`
	ChunkFetchers       int32         `mapstructure:"chunk_fetchers"`
	MaxSnapshotChunks   uint32        `mapstructure:"max_snapshot_chunks"`
}

func (cfg *StateSyncConfig) TrustHashBytes() []byte {
	_ = "STUB: not implemented"
	// validated in ValidateBasic, so we can safely panic here
	return nil
}

// DefaultStateSyncConfig returns a default configuration for the state sync service
func DefaultStateSyncConfig() *StateSyncConfig { _ = "STUB: not implemented"; return nil }

// TestStateSyncConfig returns a default configuration for the state sync service
func TestStateSyncConfig() *StateSyncConfig { _ = "STUB: not implemented"; return nil }

// ValidateBasic performs basic validation.
func (cfg *StateSyncConfig) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

//-----------------------------------------------------------------------------
// BlockSyncConfig

// BlockSyncConfig (formerly known as FastSync) defines the configuration for the CometBFT block sync service
type BlockSyncConfig struct {
	Version    string `mapstructure:"version"`
	VerifyData bool   `mapstructure:"verify_data"`
}

// DefaultBlockSyncConfig returns a default configuration for the block sync service
func DefaultBlockSyncConfig() *BlockSyncConfig { _ = "STUB: not implemented"; return nil }

// TestBlockSyncConfig returns a default configuration for the block sync.
func TestBlockSyncConfig() *BlockSyncConfig { _ = "STUB: not implemented"; return nil }

// ValidateBasic performs basic validation.
func (cfg *BlockSyncConfig) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

//-----------------------------------------------------------------------------
// ConsensusConfig

// ConsensusConfig defines the configuration for the Tendermint consensus algorithm, adopted by CometBFT,
// including timeouts and details about the WAL and the block structure.
type ConsensusConfig struct {
	RootDir string `mapstructure:"home"`
	// If set to true, only internal messages will be written
	// to the WAL. External messages like votes, proposals
	// block parts, will not be written
	// Default: true
	OnlyInternalWal bool   `mapstructure:"only_internal_wal"`
	WalPath         string `mapstructure:"wal_file"`
	walFile         string // overrides WalPath if set

	// How long we wait for a proposal block before prevoting nil
	// Deprecated: TimeoutPropose is deprecated and will be removed in a future version.
	TimeoutPropose time.Duration `mapstructure:"timeout_propose"`
	// How much timeout_propose increases with each round
	// Deprecated: TimeoutProposeDelta is deprecated and will be removed in a future version.
	TimeoutProposeDelta time.Duration `mapstructure:"timeout_propose_delta"`
	// How long we wait after receiving +2/3 prevotes for “anything” (ie. not a single block or nil)
	// Deprecated: TimeoutPrevote is deprecated and will be removed in a future version.
	TimeoutPrevote time.Duration `mapstructure:"timeout_prevote"`
	// How much the timeout_prevote increases with each round
	// Deprecated: TimeoutPrevoteDelta is deprecated and will be removed in a future version.
	TimeoutPrevoteDelta time.Duration `mapstructure:"timeout_prevote_delta"`
	// How long we wait after receiving +2/3 precommits for “anything” (ie. not a single block or nil)
	// Deprecated: TimeoutPrecommit is deprecated and will be removed in a future version.
	TimeoutPrecommit time.Duration `mapstructure:"timeout_precommit"`
	// How much the timeout_precommit increases with each round
	// Deprecated: TimeoutPrecommitDelta is deprecated and will be removed in a future version.
	TimeoutPrecommitDelta time.Duration `mapstructure:"timeout_precommit_delta"`
	// DelayedPrecommitTimeout ensures a minimum block time by waiting during the pre-commit time.
	// The new pre-commit vote time starts at: StartTime (the new block start time) + the delayed pre-commit time.
	//  Note that this change doesn't affect the rounds >= 1.
	DelayedPrecommitTimeout time.Duration // TODO dynamically change using the app
	// How long we wait after committing a block, before starting on the new
	// height (this gives us a chance to receive some more precommits, even
	// though we already have +2/3).
	// NOTE: when modifying, make sure to update time_iota_ms genesis parameter
	TimeoutCommit time.Duration `mapstructure:"timeout_commit"`

	// Make progress as soon as we have all the precommits (as if TimeoutCommit = 0)
	SkipTimeoutCommit bool `mapstructure:"skip_timeout_commit"`

	// EmptyBlocks mode and possible interval between empty blocks
	// Deprecated: CreateEmptyBlocks is deprecated and will be removed in a future version.
	CreateEmptyBlocks         bool          `mapstructure:"create_empty_blocks"`
	CreateEmptyBlocksInterval time.Duration `mapstructure:"create_empty_blocks_interval"`

	// Reactor sleep duration parameters
	PeerGossipSleepDuration     time.Duration `mapstructure:"peer_gossip_sleep_duration"`
	PeerQueryMaj23SleepDuration time.Duration `mapstructure:"peer_query_maj23_sleep_duration"`

	DoubleSignCheckHeight int64 `mapstructure:"double_sign_check_height"`

	// Disable the propagation reactor for block and proposal recovery
	DisablePropagationReactor bool `mapstructure:"disable_propagation_reactor"`
	EnableLegacyBlockProp     bool `mapstructure:"enable_legacy_block_prop"`
}

// DefaultConsensusConfig returns a default configuration for the consensus service
func DefaultConsensusConfig() *ConsensusConfig { _ = "STUB: not implemented"; return nil }

// TestConsensusConfig returns a configuration for testing the consensus service
func TestConsensusConfig() *ConsensusConfig { _ = "STUB: not implemented"; return nil }

// NOTE: when modifying, make sure to update time_iota_ms (testGenesisFmt) in toml.go

// WaitForTxs returns true if the consensus should wait for transactions before entering the propose step
func (cfg *ConsensusConfig) WaitForTxs() bool { _ = "STUB: not implemented"; return false }

// Propose returns the amount of time to wait for a proposal
func (cfg *ConsensusConfig) Propose(round int32) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// Prevote returns the amount of time to wait for straggler votes after receiving any +2/3 prevotes
func (cfg *ConsensusConfig) Prevote(round int32) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// Precommit returns the amount of time to wait for straggler votes after receiving any +2/3 precommits
func (cfg *ConsensusConfig) Precommit(round int32) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// Commit returns the amount of time to wait for straggler votes after receiving +2/3 precommits
// for a single block (ie. a commit).
func (cfg *ConsensusConfig) Commit(t time.Time) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

// ProposeWithCustomTimeout is identical to Propose. However,
// it calculates the amount of time to wait for a proposal using the supplied
// customTimeout.
// If customTimeout is 0, the TimeoutPropose from cfg is used.
func (cfg *ConsensusConfig) ProposeWithCustomTimeout(round int32, customTimeout time.Duration) time.Duration {
	_ = "STUB: not implemented"
	// this is to capture any unforeseen cases where the customTimeout is 0
	return *new(time.Duration)
}

// falling back to default timeout

// CommitWithCustomTimeout is identical to Commit. However, it calculates the time for commit using the supplied customTimeout.
// If customTimeout is 0, the TimeoutCommit from cfg is used.
func (cfg *ConsensusConfig) CommitWithCustomTimeout(t time.Time, customTimeout time.Duration) time.Time {
	_ = "STUB: not implemented"
	// this is to capture any unforeseen cases where the customTimeout is 0
	return *new(time.Time)
}

// falling back to default timeout

// WalFile returns the full path to the write-ahead log file
func (cfg *ConsensusConfig) WalFile() string { _ = "STUB: not implemented"; return "" }

// SetWalFile sets the path to the write-ahead log file
func (cfg *ConsensusConfig) SetWalFile(walFile string) { _ = "STUB: not implemented"; return }

// ValidateBasic performs basic validation (checking param bounds, etc.) and
// returns an error if any check fails.
func (cfg *ConsensusConfig) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

//-----------------------------------------------------------------------------
// StorageConfig

// StorageConfig allows more fine-grained control over certain storage-related
// behavior.
type StorageConfig struct {
	// Set to false to ensure ABCI responses are persisted. ABCI responses are
	// required for `/block_results` RPC queries, and to reindex events in the
	// command-line tool.
	DiscardABCIResponses bool `mapstructure:"discard_abci_responses"`
	// Compaction on pruning - enable or disable in-process compaction.
	// If the DB backend supports it, this will force the DB to compact
	// the database levels and save on storage space. Setting this to true
	// is most beneficial when used in combination with pruning as it will
	// physically delete the entries marked for deletion.
	// false by default (forcing compaction is disabled).
	Compact bool `mapstructure:"compact"`
	// Compaction interval - number of blocks to try explicit compaction on.
	// This parameter should be tuned depending on the number of items
	// you expect to delete between two calls to forced compaction.
	// If your retain height is 1 block, it is too much of an overhead
	// to try compaction every block. But it should also not be a very
	// large multiple of your retain height as it might occur bigger overheads.
	// 10000 by default.
	CompactionInterval int64 `mapstructure:"compaction_interval"`
}

// DefaultStorageConfig returns the default configuration options relating to
// CometBFT storage optimization.
func DefaultStorageConfig() *StorageConfig { _ = "STUB: not implemented"; return nil }

// TestStorageConfig returns storage configuration that can be used for
// testing.
func TestStorageConfig() *StorageConfig { _ = "STUB: not implemented"; return nil }

// ValidateBasic performs basic validation, returning an error if any check
// fails.
func (cfg *StorageConfig) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// -----------------------------------------------------------------------------
// TxIndexConfig
// Remember that Event has the following structure:
// type: [
//
//	key: value,
//	...
//
// ]
//
// CompositeKeys are constructed by `type.key`
// TxIndexConfig defines the configuration for the transaction indexer,
// including composite keys to index.
type TxIndexConfig struct {
	// Indexer is the indexer to use for transactions
	Indexer string `mapstructure:"indexer"`

	// The PostgreSQL connection configuration, the connection format:
	// postgresql://<user>:<password>@<host>:<port>/<db>?<opts>
	PsqlConn string `mapstructure:"psql-conn"`
}

// DefaultTxIndexConfig returns a default configuration for the transaction indexer.
func DefaultTxIndexConfig() *TxIndexConfig { _ = "STUB: not implemented"; return nil }

// TestTxIndexConfig returns a test configuration for the transaction indexer
// with "kv" enabled so that RPC tests can still use indexer-dependent endpoints.
func TestTxIndexConfig() *TxIndexConfig { _ = "STUB: not implemented"; return nil }

//-----------------------------------------------------------------------------
// InstrumentationConfig

// InstrumentationConfig defines the configuration for metrics reporting.
type InstrumentationConfig struct {
	// When true, Prometheus metrics are served under /metrics on
	// PrometheusListenAddr.
	// Check out the documentation for the list of available metrics.
	Prometheus bool `mapstructure:"prometheus"`

	// Address to listen for Prometheus collector(s) connections.
	PrometheusListenAddr string `mapstructure:"prometheus_listen_addr"`

	// Maximum number of simultaneous connections.
	// If you want to accept a larger number than the default, make sure
	// you increase your OS limits.
	// 0 - unlimited.
	MaxOpenConnections int `mapstructure:"max_open_connections"`

	// Instrumentation namespace.
	Namespace string `mapstructure:"namespace"`

	// TracePushConfig is the relative path of the push config. This second
	// config contains credentials for where and how often to.
	TracePushConfig string `mapstructure:"trace_push_config"`

	// TracePullAddress is the address that the trace server will listen on for
	// pulling data.
	TracePullAddress string `mapstructure:"trace_pull_address"`

	// TraceType is the type of tracer used. Options are "local" and "noop".
	TraceType string `mapstructure:"trace_type"`

	// TraceBufferSize is the number of traces to write in a single batch.
	TraceBufferSize int `mapstructure:"trace_push_batch_size"`

	// TracingTables is the list of tables that will be traced. See the
	// pkg/trace/schema for a complete list of tables. It is represented as a
	// comma separate string. For example: "consensus_round_state,mempool_tx".
	TracingTables string `mapstructure:"tracing_tables"`

	// PyroscopeURL is the pyroscope url used to establish a connection with a
	// pyroscope continuous profiling server.
	PyroscopeURL string `mapstructure:"pyroscope_url"`

	// PyroscopeProfile is a flag that enables tracing with pyroscope.
	PyroscopeTrace bool `mapstructure:"pyroscope_trace"`

	// PyroscopeProfileTypes is a list of profile types to be traced with
	// pyroscope. Available profile types are: cpu, alloc_objects, alloc_space,
	// inuse_objects, inuse_space, goroutines, mutex_count, mutex_duration,
	// block_count, block_duration.
	PyroscopeProfileTypes []string `mapstructure:"pyroscope_profile_types"`
}

// DefaultInstrumentationConfig returns a default configuration for metrics
// reporting.
func DefaultInstrumentationConfig() *InstrumentationConfig { _ = "STUB: not implemented"; return nil }

// TestInstrumentationConfig returns a default configuration for metrics
// reporting.
func TestInstrumentationConfig() *InstrumentationConfig { _ = "STUB: not implemented"; return nil }

// ValidateBasic performs basic validation (checking param bounds, etc.) and
// returns an error if any check fails.
func (cfg *InstrumentationConfig) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// if there is not TracePushConfig configured, then we do not need to validate the rest
// of the config because we are not connecting.

func (cfg *InstrumentationConfig) IsPrometheusEnabled() bool {
	_ = "STUB: not implemented"
	return false
}

//-----------------------------------------------------------------------------
// Utils

// helper function to make config creation independent of root dir
func rootify(path, root string) string { _ = "STUB: not implemented"; return "" }

//-----------------------------------------------------------------------------
// Moniker

var defaultMoniker = getDefaultMoniker()

// getDefaultMoniker returns a default moniker, which is the host name. If runtime
// fails to get the host name, "anonymous" will be returned.
func getDefaultMoniker() string { _ = "STUB: not implemented"; return "" }
