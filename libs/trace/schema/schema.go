package schema

import (
	"strings"

	"github.com/cometbft/cometbft/config"
)

func init() {
	config.DefaultTracingTables = strings.Join(AllTables(), ",")
}

func AllTables() []string {
	_ = "STUB: not implemented"
	//nolint:prealloc
	return nil
}

const (
	Broadcast = "broadcast"
)

type TransferType int

const (
	Download TransferType = iota
	Upload
	Haves
)

func (t TransferType) String() string { _ = "STUB: not implemented"; return "" }
