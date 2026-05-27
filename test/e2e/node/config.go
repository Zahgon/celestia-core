package main

import (
	"github.com/cometbft/cometbft/test/e2e/app"
)

// Config is the application configuration.
type Config struct {
	ChainID                    string                      `toml:"chain_id"`
	Listen                     string                      `toml:"listen"`
	Protocol                   string                      `toml:"protocol"`
	Dir                        string                      `toml:"dir"`
	Mode                       string                      `toml:"mode"`
	PersistInterval            uint64                      `toml:"persist_interval"`
	SnapshotInterval           uint64                      `toml:"snapshot_interval"`
	RetainBlocks               uint64                      `toml:"retain_blocks"`
	ValidatorUpdates           map[string]map[string]uint8 `toml:"validator_update"`
	PrivValServer              string                      `toml:"privval_server"`
	PrivValKey                 string                      `toml:"privval_key"`
	PrivValState               string                      `toml:"privval_state"`
	KeyType                    string                      `toml:"key_type"`
	VoteExtensionsEnableHeight int64                       `toml:"vote_extensions_enable_height"`
	VoteExtensionsUpdateHeight int64                       `toml:"vote_extensions_update_height"`
}

// App extracts out the application specific configuration parameters
func (cfg *Config) App() *app.Config { _ = "STUB: not implemented"; return nil }

// LoadConfig loads the configuration from disk.
func LoadConfig(file string) (*Config, error) { _ = "STUB: not implemented"; return nil, nil }

// Validate validates the configuration. We don't do exhaustive config
// validation here, instead relying on Testnet.Validate() to handle it.
func (cfg Config) Validate() error { _ = "STUB: not implemented"; return nil }
