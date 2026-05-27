package node

import (
	"github.com/grafana/pyroscope-go"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"

	"github.com/cometbft/cometbft/config"
)

// setupPyroscope sets up pyroscope profiler and optionally tracing.
func setupPyroscope(instCfg *config.InstrumentationConfig, nodeID string) (*pyroscope.Profiler, *sdktrace.TracerProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// use the noop logger by passing nil

func setupTracing(addr string, labels map[string]string) (tp *sdktrace.TracerProvider, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Set the Tracer Provider and the W3C Trace Context propagator as globals.
// We wrap the tracer provider to also annotate goroutines with Span ID so
// that pprof would add corresponding labels to profiling samples.

// Register the trace context and baggage propagators so data is propagated across services/processes.

func tracerProviderDebug() (*sdktrace.TracerProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func toPyroscopeProfiles(profiles []string) []pyroscope.ProfileType {
	_ = "STUB: not implemented"
	return nil
}
