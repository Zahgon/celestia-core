package trace

import (
	"os"

	"github.com/cometbft/cometbft/config"
	"github.com/cometbft/cometbft/libs/log"
)

// Entry is an interface for all structs that are used to define the schema for
// traces.
type Entry interface {
	// Table defines which table the struct belongs to.
	Table() string
}

// Tracer defines the methods for a client that can write and read trace data.
type Tracer interface {
	Write(Entry)
	IsCollecting(table string) bool
	Stop()
}

func NewTracer(cfg *config.Config, logger log.Logger, chainID, nodeID string) (Tracer, error) {
	_ = "STUB: not implemented"
	return *new(Tracer), nil
}

func NoOpTracer() Tracer { _ = "STUB: not implemented"; return *new(Tracer) }

type noOpTracer struct{}

func (*noOpTracer) Write(_ Entry) { _ = "STUB: not implemented"; return }
func (*noOpTracer) ReadTable(_ string) (*os.File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*noOpTracer) IsCollecting(_ string) bool { _ = "STUB: not implemented"; return false }
func (*noOpTracer) Stop()                      { _ = "STUB: not implemented"; return }
