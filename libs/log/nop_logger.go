package log

type nopLogger struct{}

// Interface assertions
var _ Logger = (*nopLogger)(nil)

// NewNopLogger returns a logger that doesn't do anything.
func NewNopLogger() Logger { _ = "STUB: not implemented"; return *new(Logger) }

func (nopLogger) Trace(string, ...interface{}) { _ = "STUB: not implemented"; return }
func (nopLogger) Info(string, ...interface{})  { _ = "STUB: not implemented"; return }
func (nopLogger) Debug(string, ...interface{}) { _ = "STUB: not implemented"; return }
func (nopLogger) Error(string, ...interface{}) { _ = "STUB: not implemented"; return }

func (l *nopLogger) With(...interface{}) Logger { _ = "STUB: not implemented"; return *new(Logger) }
