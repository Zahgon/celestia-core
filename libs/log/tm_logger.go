package log

import (
	"io"

	kitlog "github.com/go-kit/log"
	"github.com/go-kit/log/term"
)

const (
	msgKey    = "_msg" // "_" prefixed to avoid collisions
	moduleKey = "module"
	levelKey  = "level"
)

type tmLogger struct {
	srcLogger kitlog.Logger
}

// Interface assertions
var _ Logger = (*tmLogger)(nil)

// NewTMLogger returns a logger that encodes msg and keyvals to the Writer
// using go-kit's log as an underlying logger and our custom formatter. Note
// that underlying logger could be swapped with something else.
func NewTMLogger(w io.Writer) Logger { _ = "STUB: not implemented"; return *new(Logger) }

// Handle custom trace level

// Safely handle this case

// Handle standard go-kit levels

// Type-safe extraction for current go-kit/log (pointer and value types)

// Never panic: just return default color for unknown key types

// NewTMLoggerWithColorFn allows you to provide your own color function. See
// NewTMLogger for documentation.
func NewTMLoggerWithColorFn(w io.Writer, colorFn func(keyvals ...interface{}) term.FgBgColor) Logger {
	_ = "STUB: not implemented"
	return *new(Logger)
}

// Trace logs a message at level Trace.
func (l *tmLogger) Trace(msg string, keyvals ...interface{}) { _ = "STUB: not implemented"; return }

//nolint:errcheck // no need to check error again

// Info logs a message at level Info.
func (l *tmLogger) Info(msg string, keyvals ...interface{}) { _ = "STUB: not implemented"; return }

//nolint:errcheck // no need to check error again

// Debug logs a message at level Debug.
func (l *tmLogger) Debug(msg string, keyvals ...interface{}) { _ = "STUB: not implemented"; return }

//nolint:errcheck // no need to check error again

// Error logs a message at level Error.
func (l *tmLogger) Error(msg string, keyvals ...interface{}) { _ = "STUB: not implemented"; return }

//nolint:errcheck // no need to check error again

// With returns a new contextual logger with keyvals prepended to those passed
// to calls to Trace, Info, Debug or Error.
func (l *tmLogger) With(keyvals ...interface{}) Logger {
	_ = "STUB: not implemented"
	return *new(Logger)
}
