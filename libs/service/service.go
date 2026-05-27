package service

import (
	"errors"

	"github.com/cometbft/cometbft/libs/log"
)

var (
	// ErrAlreadyStarted is returned when somebody tries to start an already
	// running service.
	ErrAlreadyStarted = errors.New("already started")
	// ErrAlreadyStopped is returned when somebody tries to stop an already
	// stopped service (without resetting it).
	ErrAlreadyStopped = errors.New("already stopped")
	// ErrNotStarted is returned when somebody tries to stop a not running
	// service.
	ErrNotStarted = errors.New("not started")
)

// Service defines a service that can be started, stopped, and reset.
type Service interface {
	// Start the service.
	// If it's already started or stopped, will return an error.
	// If OnStart() returns an error, it's returned by Start()
	Start() error
	OnStart() error

	// Stop the service.
	// If it's already stopped, will return an error.
	// OnStop must never error.
	Stop() error
	OnStop()

	// Reset the service.
	// Panics by default - must be overwritten to enable reset.
	Reset() error
	OnReset() error

	// Return true if the service is running
	IsRunning() bool

	// Quit returns a channel, which is closed once service is stopped.
	Quit() <-chan struct{}

	// String representation of the service
	String() string

	// SetLogger sets a logger.
	SetLogger(log.Logger)
}

/*
Classical-inheritance-style service declarations. Services can be started, then
stopped, then optionally restarted.

Users can override the OnStart/OnStop methods. In the absence of errors, these
methods are guaranteed to be called at most once. If OnStart returns an error,
service won't be marked as started, so the user can call Start again.

A call to Reset will panic, unless OnReset is overwritten, allowing
OnStart/OnStop to be called again.

The caller must ensure that Start and Stop are not called concurrently.

It is ok to call Stop without calling Start first.

Typical usage:

	type FooService struct {
		BaseService
		// private fields
	}

	func NewFooService() *FooService {
		fs := &FooService{
			// init
		}
		fs.BaseService = *NewBaseService(log, "FooService", fs)
		return fs
	}

	func (fs *FooService) OnStart() error {
		fs.BaseService.OnStart() // Always call the overridden method.
		// initialize private fields
		// start subroutines, etc.
	}

	func (fs *FooService) OnStop() error {
		fs.BaseService.OnStop() // Always call the overridden method.
		// close/destroy private fields
		// stop subroutines, etc.
	}
*/
type BaseService struct {
	Logger  log.Logger
	name    string
	started uint32 // atomic
	stopped uint32 // atomic
	quit    chan struct{}

	// The "subclass" of BaseService
	impl Service
}

// NewBaseService creates a new BaseService.
func NewBaseService(logger log.Logger, name string, impl Service) *BaseService {
	_ = "STUB: not implemented"
	return nil
}

// SetLogger implements Service by setting a logger.
func (bs *BaseService) SetLogger(l log.Logger) {
	_ = "STUB: not implemented"

	// Start implements Service by calling OnStart (if defined). An error will be
	// returned if the service is already running or stopped. Not to start the
	// stopped service, you need to call Reset.
	return
}

func (bs *BaseService) Start() error { _ = "STUB: not implemented"; return nil }

// revert flag

// revert flag

// OnStart implements Service by doing nothing.
// NOTE: Do not put anything in here,
// that way users don't need to call BaseService.OnStart()
func (bs *BaseService) OnStart() error {
	_ = "STUB: not implemented"

	// Stop implements Service by calling OnStop (if defined) and closing quit
	// channel. An error will be returned if the service is already stopped.
	return nil
}

func (bs *BaseService) Stop() error { _ = "STUB: not implemented"; return nil }

// revert flag

// OnStop implements Service by doing nothing.
// NOTE: Do not put anything in here,
// that way users don't need to call BaseService.OnStop()
func (bs *BaseService) OnStop() {
	_ = "STUB: not implemented"

	// Reset implements Service by calling OnReset callback (if defined). An error
	// will be returned if the service is running.
	return
}

func (bs *BaseService) Reset() error { _ = "STUB: not implemented"; return nil }

// whether or not we've started, we can reset

// OnReset implements Service by panicking.
func (bs *BaseService) OnReset() error { _ = "STUB: not implemented"; return nil }

// IsRunning implements Service by returning true or false depending on the
// service's state.
func (bs *BaseService) IsRunning() bool { _ = "STUB: not implemented"; return false }

// Wait blocks until the service is stopped.
func (bs *BaseService) Wait() {
	_ = "STUB: not implemented"

	// String implements Service by returning a string representation of the service.
	return
}

func (bs *BaseService) String() string {
	_ = "STUB: not implemented"

	// Quit Implements Service by returning a quit channel.
	return ""
}

func (bs *BaseService) Quit() <-chan struct{} { _ = "STUB: not implemented"; return nil }
