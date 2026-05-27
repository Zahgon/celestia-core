// Package pubsub implements a pub-sub model with a single publisher (Server)
// and multiple subscribers (clients).
//
// Though you can have multiple publishers by sharing a pointer to a server or
// by giving the same channel to each publisher and publishing messages from
// that channel (fan-in).
//
// Clients subscribe for messages, which could be of any type, using a query.
// When some message is published, we match it with all queries. If there is a
// match, this message will be pushed to all clients, subscribed to that query.
// See query subpackage for our implementation.
//
// Example:
//
//	q, err := query.New("account.name='John'")
//	if err != nil {
//	    return err
//	}
//	ctx, cancel := context.WithTimeout(context.Background(), 1 * time.Second)
//	defer cancel()
//	subscription, err := pubsub.Subscribe(ctx, "johns-transactions", q)
//	if err != nil {
//	    return err
//	}
//
//	for {
//	    select {
//	    case msg <- subscription.Out():
//	        // handle msg.Data() and msg.Events()
//	    case <-subscription.Canceled():
//	        return subscription.Err()
//	    }
//	}
package pubsub

import (
	"context"
	"errors"

	"github.com/cometbft/cometbft/libs/service"
	cmtsync "github.com/cometbft/cometbft/libs/sync"
)

type operation int

const (
	sub operation = iota
	pub
	unsub
	shutdown
)

var (
	// ErrSubscriptionNotFound is returned when a client tries to unsubscribe
	// from not existing subscription.
	ErrSubscriptionNotFound = errors.New("subscription not found")

	// ErrAlreadySubscribed is returned when a client tries to subscribe twice or
	// more using the same query.
	ErrAlreadySubscribed = errors.New("already subscribed")
)

// Query defines an interface for a query to be used for subscribing. A query
// matches against a map of events. Each key in this map is a composite of the
// even type and an attribute key (e.g. "{eventType}.{eventAttrKey}") and the
// values are the event values that are contained under that relationship. This
// allows event types to repeat themselves with the same set of keys and
// different values.
type Query interface {
	Matches(events map[string][]string) (bool, error)
	String() string
}

type cmd struct {
	op operation

	// subscribe, unsubscribe
	query        Query
	subscription *Subscription
	clientID     string

	// publish
	msg    interface{}
	events map[string][]string
}

// Server allows clients to subscribe/unsubscribe for messages, publishing
// messages with or without events, and manages internal state.
type Server struct {
	service.BaseService

	cmds    chan cmd
	cmdsCap int

	// check if we have subscription before
	// subscribing or unsubscribing
	mtx           cmtsync.RWMutex
	subscriptions map[string]map[string]struct{} // subscriber -> query (string) -> empty struct
}

// Option sets a parameter for the server.
type Option func(*Server)

// NewServer returns a new server. See the commentary on the Option functions
// for a detailed description of how to configure buffering. If no options are
// provided, the resulting server's queue is unbuffered.
func NewServer(options ...Option) *Server { _ = "STUB: not implemented"; return nil }

// if BufferCapacity option was not set, the channel is unbuffered

// BufferCapacity allows you to specify capacity for the internal server's
// queue. Since the server, given Y subscribers, could only process X messages,
// this option could be used to survive spikes (e.g. high amount of
// transactions during peak hours).
func BufferCapacity(cap int) Option { _ = "STUB: not implemented"; return *new(Option) }

// BufferCapacity returns capacity of the internal server's queue.
func (s *Server) BufferCapacity() int {
	_ = "STUB: not implemented"

	// Subscribe creates a subscription for the given client.
	//
	// An error will be returned to the caller if the context is canceled or if
	// subscription already exist for pair clientID and query.
	//
	// outCapacity can be used to set a capacity for Subscription#Out channel (1 by
	// default). Panics if outCapacity is less than or equal to zero. If you want
	// an unbuffered channel, use SubscribeUnbuffered.
	return 0
}

func (s *Server) Subscribe(
	ctx context.Context,
	clientID string,
	query Query,
	outCapacity ...int) (*Subscription, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SubscribeUnbuffered does the same as Subscribe, except it returns a
// subscription with unbuffered channel. Use with caution as it can freeze the
// server.
func (s *Server) SubscribeUnbuffered(ctx context.Context, clientID string, query Query) (*Subscription, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Server) subscribe(ctx context.Context, clientID string, query Query, outCapacity int) (*Subscription, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Unsubscribe removes the subscription on the given query. An error will be
// returned to the caller if the context is canceled or if subscription does
// not exist.
func (s *Server) Unsubscribe(ctx context.Context, clientID string, query Query) error {
	_ = "STUB: not implemented"
	return nil
}

// UnsubscribeAll removes all client subscriptions. An error will be returned
// to the caller if the context is canceled or if subscription does not exist.
func (s *Server) UnsubscribeAll(ctx context.Context, clientID string) error {
	_ = "STUB: not implemented"
	return nil
}

// NumClients returns the number of clients.
func (s *Server) NumClients() int { _ = "STUB: not implemented"; return 0 }

// NumClientSubscriptions returns the number of subscriptions the client has.
func (s *Server) NumClientSubscriptions(clientID string) int { _ = "STUB: not implemented"; return 0 }

// Publish publishes the given message. An error will be returned to the caller
// if the context is canceled.
func (s *Server) Publish(ctx context.Context, msg interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// PublishWithEvents publishes the given message with the set of events. The set
// is matched with clients queries. If there is a match, the message is sent to
// the client.
func (s *Server) PublishWithEvents(ctx context.Context, msg interface{}, events map[string][]string) error {
	_ = "STUB: not implemented"
	return nil
}

// OnStop implements Service.OnStop by shutting down the server.
func (s *Server) OnStop() { _ = "STUB: not implemented"; return }

// NOTE: not goroutine safe
type state struct {
	// query string -> client -> subscription
	subscriptions map[string]map[string]*Subscription
	// query string -> queryPlusRefCount
	queries map[string]*queryPlusRefCount
}

// queryPlusRefCount holds a pointer to a query and reference counter. When
// refCount is zero, query will be removed.
type queryPlusRefCount struct {
	q        Query
	refCount int
}

// OnStart implements Service.OnStart by starting the server.
func (s *Server) OnStart() error { _ = "STUB: not implemented"; return nil }

// OnReset implements Service.OnReset
func (s *Server) OnReset() error { _ = "STUB: not implemented"; return nil }

func (s *Server) loop(state state) { _ = "STUB: not implemented"; return }

func (state *state) add(clientID string, q Query, subscription *Subscription) {
	_ = "STUB: not implemented"

	// initialize subscription for this client per query if needed
	return
}

// create subscription

// initialize query if needed

// increment reference counter

func (state *state) remove(clientID string, qStr string, reason error) {
	_ = "STUB: not implemented"
	return
}

// remove client from query map.
// if query has no other clients subscribed, remove it.

// decrease ref counter in queries

// remove the query if nobody else is using it

func (state *state) removeClient(clientID string, reason error) { _ = "STUB: not implemented"; return }

func (state *state) removeAll(reason error) { _ = "STUB: not implemented"; return }

func (state *state) send(msg interface{}, events map[string][]string) error {
	_ = "STUB: not implemented"
	return nil
}

// block on unbuffered channel

// don't block on buffered channels
