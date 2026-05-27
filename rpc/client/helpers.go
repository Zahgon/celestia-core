package client

import (
	"time"

	"github.com/cometbft/cometbft/types"
)

// Waiter is informed of current height, decided whether to quit early
type Waiter func(delta int64) (abort error)

// DefaultWaitStrategy is the standard backoff algorithm,
// but you can plug in another one
func DefaultWaitStrategy(delta int64) (abort error) { _ = "STUB: not implemented"; return nil }

// estimate of wait time....
// wait half a second for the next block (in progress)
// plus one second for every full block

// Wait for height will poll status at reasonable intervals until
// the block at the given height is available.
//
// If waiter is nil, we use DefaultWaitStrategy, but you can also
// provide your own implementation
func WaitForHeight(c StatusClient, h int64, waiter Waiter) error {
	_ = "STUB: not implemented"
	return nil
}

// wait for the time, or abort early

// WaitForOneEvent subscribes to a websocket event for the given
// event time and returns upon receiving it one time, or
// when the timeout duration has expired.
//
// This handles subscribing and unsubscribing under the hood
func WaitForOneEvent(c EventsClient, evtTyp string, timeout time.Duration) (types.TMEventData, error) {
	_ = "STUB: not implemented"
	return *new(types.TMEventData), nil
}

// register for the next event of this type

// make sure to unregister after the test is over
