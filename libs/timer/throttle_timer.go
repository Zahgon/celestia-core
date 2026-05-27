package timer

import (
	"time"

	cmtsync "github.com/cometbft/cometbft/libs/sync"
)

/*
ThrottleTimer fires an event at most "dur" after each .Set() call.
If a short burst of .Set() calls happens, ThrottleTimer fires once.
If a long continuous burst of .Set() calls happens, ThrottleTimer fires
at most once every "dur".
*/
type ThrottleTimer struct {
	Name string
	Ch   chan struct{}
	quit chan struct{}
	dur  time.Duration

	mtx   cmtsync.Mutex
	timer *time.Timer
	isSet bool
}

func NewThrottleTimer(name string, dur time.Duration) *ThrottleTimer {
	_ = "STUB: not implemented"
	return nil
}

func (t *ThrottleTimer) fireRoutine() { _ = "STUB: not implemented"; return }

// do nothing

func (t *ThrottleTimer) Set() { _ = "STUB: not implemented"; return }

func (t *ThrottleTimer) Unset() { _ = "STUB: not implemented"; return }

// For ease of .Stop()'ing services before .Start()'ing them,
// we ignore .Stop()'s on nil ThrottleTimers
func (t *ThrottleTimer) Stop() bool { _ = "STUB: not implemented"; return false }
