package pubsub

import (
	"errors"

	cmtsync "github.com/cometbft/cometbft/libs/sync"
)

var (
	// ErrUnsubscribed is returned by Err when a client unsubscribes.
	ErrUnsubscribed = errors.New("client unsubscribed")

	// ErrOutOfCapacity is returned by Err when a client is not pulling messages
	// fast enough. Note the client's subscription will be terminated.
	ErrOutOfCapacity = errors.New("internal subscription event buffer is out of capacity")
)

// A Subscription represents a client subscription for a particular query and
// consists of three things:
// 1) channel onto which messages and events are published
// 2) channel which is closed if a client is too slow or choose to unsubscribe
// 3) err indicating the reason for (2)
type Subscription struct {
	out chan Message

	canceled chan struct{}
	mtx      cmtsync.RWMutex
	err      error
}

// NewSubscription returns a new subscription with the given outCapacity.
func NewSubscription(outCapacity int) *Subscription { _ = "STUB: not implemented"; return nil }

// Out returns a channel onto which messages and events are published.
// Unsubscribe/UnsubscribeAll does not close the channel to avoid clients from
// receiving a nil message.
func (s *Subscription) Out() <-chan Message {
	_ = "STUB: not implemented"

	// Canceled returns a channel that's closed when the subscription is
	// terminated and supposed to be used in a select statement.
	return nil
}

func (s *Subscription) Canceled() <-chan struct{} {
	_ = "STUB: not implemented"

	// Err returns nil if the channel returned is not yet closed.
	// If the channel is closed, Err returns a non-nil error explaining why:
	//   - ErrUnsubscribed if the subscriber choose to unsubscribe,
	//   - ErrOutOfCapacity if the subscriber is not pulling messages fast enough
	//     and the channel returned by Out became full,
	//
	// After Err returns a non-nil error, successive calls to Err return the same
	// error.
	return nil
}

func (s *Subscription) Err() error { _ = "STUB: not implemented"; return nil }

func (s *Subscription) cancel(err error) { _ = "STUB: not implemented"; return }

// Message glues data and events together.
type Message struct {
	data   interface{}
	events map[string][]string
}

func NewMessage(data interface{}, events map[string][]string) Message {
	_ = "STUB: not implemented"
	return *new(Message)
}

// Data returns an original data published.
func (msg Message) Data() interface{} {
	_ = "STUB: not implemented"

	// Events returns events, which matched the client's query.
	return nil
}

func (msg Message) Events() map[string][]string { _ = "STUB: not implemented"; return nil }
