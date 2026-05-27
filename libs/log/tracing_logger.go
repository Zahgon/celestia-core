package log

import (
	"github.com/pkg/errors"
)

// NewTracingLogger enables tracing by wrapping all errors (if they
// implement stackTracer interface) in tracedError.
//
// All errors returned by https://github.com/pkg/errors implement stackTracer
// interface.
//
// For debugging purposes only as it doubles the amount of allocations.
func NewTracingLogger(next Logger) Logger { _ = "STUB: not implemented"; return *new(Logger) }

type stackTracer interface {
	error
	StackTrace() errors.StackTrace
}

type tracingLogger struct {
	next Logger
}

func (l *tracingLogger) Trace(msg string, keyvals ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (l *tracingLogger) Info(msg string, keyvals ...interface{}) { _ = "STUB: not implemented"; return }

func (l *tracingLogger) Debug(msg string, keyvals ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (l *tracingLogger) Error(msg string, keyvals ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (l *tracingLogger) With(keyvals ...interface{}) Logger {
	_ = "STUB: not implemented"
	return *new(Logger)
}

func formatErrors(keyvals []interface{}) []interface{} { _ = "STUB: not implemented"; return nil }

// tracedError wraps a stackTracer and just makes the Error() result
// always return a full stack trace.
type tracedError struct {
	wrapped stackTracer
}

var _ stackTracer = tracedError{}

func (t tracedError) StackTrace() errors.StackTrace {
	_ = "STUB: not implemented"
	return *new(errors.StackTrace)
}

func (t tracedError) Cause() error { _ = "STUB: not implemented"; return nil }

func (t tracedError) Error() string { _ = "STUB: not implemented"; return "" }
