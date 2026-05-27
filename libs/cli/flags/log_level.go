package flags

import (
	"github.com/cometbft/cometbft/libs/log"
)

const (
	defaultLogLevelKey = "*"
)

// ParseLogLevel parses complex log level - comma-separated
// list of module:level pairs with an optional *:level pair (* means
// all other modules).
//
// Example:
//
//	ParseLogLevel("consensus:debug,mempool:debug,*:error", log.NewTMLogger(os.Stdout), "info")
func ParseLogLevel(lvl string, logger log.Logger, defaultLogLevelValue string) (log.Logger, error) {
	_ = "STUB: not implemented"
	return *new(log.Logger), nil
}

// prefix simple one word levels (e.g. "info") with "*"

// if "*" is not provided, set default global level
