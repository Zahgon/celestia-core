package log

type level byte

const (
	levelTrace level = 1 << iota
	levelDebug
	levelInfo
	levelError
)

type filter struct {
	next             Logger
	allowed          level            // XOR'd levels for default case
	initiallyAllowed level            // XOR'd levels for initial case
	allowedKeyvals   map[keyval]level // When key-value match, use this level
}

type keyval struct {
	key   interface{}
	value interface{}
}

// NewFilter wraps next and implements filtering. See the commentary on the
// Option functions for a detailed description of how to configure levels. If
// no options are provided, all leveled log events created with Trace, Debug,
// Info or Error helper methods are squelched.
func NewFilter(next Logger, options ...Option) Logger {
	_ = "STUB: not implemented"
	return *new(Logger)
}

func (l *filter) Trace(msg string, keyvals ...interface{}) { _ = "STUB: not implemented"; return }

func (l *filter) Info(msg string, keyvals ...interface{}) { _ = "STUB: not implemented"; return }

func (l *filter) Debug(msg string, keyvals ...interface{}) { _ = "STUB: not implemented"; return }

func (l *filter) Error(msg string, keyvals ...interface{}) { _ = "STUB: not implemented"; return }

// With implements Logger by constructing a new filter with a keyvals appended
// to the logger.
//
// If custom level was set for a keyval pair using one of the
// Allow*With methods, it is used as the logger's level.
//
// Examples:
//
//	    logger = log.NewFilter(logger, log.AllowError(), log.AllowInfoWith("module", "crypto"))
//			 logger.With("module", "crypto").Info("Hello") # produces "I... Hello module=crypto"
//
//	    logger = log.NewFilter(logger, log.AllowError(),
//					log.AllowInfoWith("module", "crypto"),
//					log.AllowNoneWith("user", "Sam"))
//			 logger.With("module", "crypto", "user", "Sam").Info("Hello") # returns nil
//
//	    logger = log.NewFilter(logger,
//					log.AllowError(),
//					log.AllowInfoWith("module", "crypto"), log.AllowNoneWith("user", "Sam"))
//			 logger.With("user", "Sam").With("module", "crypto").Info("Hello") # produces "I... Hello module=crypto user=Sam"
func (l *filter) With(keyvals ...interface{}) Logger {
	_ = "STUB: not implemented"
	return *new(Logger)
}

// Example:
//		logger = log.NewFilter(logger, log.AllowError(), log.AllowInfoWith("module", "crypto"))
//		logger.With("module", "crypto")

// set the desired level

// Example:
//		logger = log.NewFilter(logger, log.AllowError(), log.AllowInfoWith("module", "crypto"))
//		logger.With("module", "main")

// return back to initially allowed

// simply continue with the current level

//--------------------------------------------------------------------------------

// Option sets a parameter for the filter.
type Option func(*filter)

// AllowLevel returns an option for the given level or error if no option exist
// for such level.
func AllowLevel(lvl string) (Option, error) { _ = "STUB: not implemented"; return *new(Option), nil }

// AllowAll is an alias for AllowTrace.
func AllowAll() Option {
	_ = "STUB: not implemented"
	return *

	// AllowTrace allows error, info, debug and trace level log events to pass.
	new(Option)
}

func AllowTrace() Option { _ = "STUB: not implemented"; return *new(Option) }

// AllowDebug allows error, info and debug level log events to pass.
func AllowDebug() Option { _ = "STUB: not implemented"; return *new(Option) }

// AllowInfo allows error and info level log events to pass.
func AllowInfo() Option { _ = "STUB: not implemented"; return *new(Option) }

// AllowError allows only error level log events to pass.
func AllowError() Option { _ = "STUB: not implemented"; return *new(Option) }

// AllowNone allows no leveled log events to pass.
func AllowNone() Option { _ = "STUB: not implemented"; return *new(Option) }

func allowed(allowed level) Option { _ = "STUB: not implemented"; return *new(Option) }

// AllowTraceWith allows error, info, debug and trace level log events to pass for a specific key value pair.
func AllowTraceWith(key interface{}, value interface{}) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// AllowDebugWith allows error, info and debug level log events to pass for a specific key value pair.
func AllowDebugWith(key interface{}, value interface{}) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// AllowInfoWith allows error and info level log events to pass for a specific key value pair.
func AllowInfoWith(key interface{}, value interface{}) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// AllowErrorWith allows only error level log events to pass for a specific key value pair.
func AllowErrorWith(key interface{}, value interface{}) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// AllowNoneWith allows no leveled log events to pass for a specific key value pair.
func AllowNoneWith(key interface{}, value interface{}) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
