package consensus

import (
	"time"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cometbft/cometbft/libs/service"
)

var (
	tickTockBufferSize = 10
)

// TimeoutTicker is a timer that schedules timeouts
// conditional on the height/round/step in the timeoutInfo.
// The timeoutInfo.Duration may be non-positive.
type TimeoutTicker interface {
	Start() error
	Stop() error
	Chan() <-chan timeoutInfo       // on which to receive a timeout
	ScheduleTimeout(ti timeoutInfo) // reset the timer

	SetLogger(log.Logger)
}

// timeoutTicker wraps time.Timer,
// scheduling timeouts only for greater height/round/step
// than what it's already seen.
// Timeouts are scheduled along the tickChan,
// and fired on the tockChan.
type timeoutTicker struct {
	service.BaseService

	timerActive bool
	timer       *time.Timer
	tickChan    chan timeoutInfo // for scheduling timeouts
	tockChan    chan timeoutInfo // for notifying about them
}

// NewTimeoutTicker returns a new TimeoutTicker.
func NewTimeoutTicker() TimeoutTicker { _ = "STUB: not implemented"; return *new(TimeoutTicker) }

// An indicator variable to check if the timer is active or not.
// Concurrency safe because the timer is only accessed by a single goroutine.

// don't want to fire until the first scheduled timeout

// OnStart implements service.Service. It starts the timeout routine.
func (t *timeoutTicker) OnStart() error { _ = "STUB: not implemented"; return nil }

// OnStop implements service.Service. It stops the timeout routine.
func (t *timeoutTicker) OnStop() { _ = "STUB: not implemented"; return }

// Chan returns a channel on which timeouts are sent.
func (t *timeoutTicker) Chan() <-chan timeoutInfo {
	_ = "STUB: not implemented"

	// ScheduleTimeout schedules a new timeout by sending on the internal tickChan.
	// The timeoutRoutine is always available to read from tickChan, so this won't block.
	// The scheduling may fail if the timeoutRoutine has already scheduled a timeout for a later height/round/step.
	return nil
}

func (t *timeoutTicker) ScheduleTimeout(ti timeoutInfo) {
	_ = "STUB: not implemented"

	// -------------------------------------------------------------
	return
}

// if the timer is active, stop it and drain the channel.
func (t *timeoutTicker) stopTimer() { _ = "STUB: not implemented"; return }

// Stop() returns false if it was already fired or was stopped

// send on tickChan to start a new timer.
// timers are interrupted and replaced by new ticks from later steps
// timeouts of 0 on the tickChan will be immediately relayed to the tockChan.
// NOTE: timerActive is not concurrency safe, but it's only accessed in NewTimer and timeoutRoutine,
// making it single-threaded access.
func (t *timeoutTicker) timeoutRoutine() { _ = "STUB: not implemented"; return }

// ignore tickers for old height/round/step

// stop the last timer if it exists

// update timeoutInfo, reset timer, and mark timer as active
// NOTE time.Timer allows duration to be non-positive

// go routine here guarantees timeoutRoutine doesn't block.
// Determinism comes from playback in the receiveRoutine.
// We can eliminate it by merging the timeoutRoutine into receiveRoutine
//  and managing the timeouts ourselves with a millisecond ticker
