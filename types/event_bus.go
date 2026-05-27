package types

import (
	"context"

	"github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/libs/log"
	cmtpubsub "github.com/cometbft/cometbft/libs/pubsub"
	"github.com/cometbft/cometbft/libs/service"
)

const defaultCapacity = 1000

type EventBusSubscriber interface {
	Subscribe(ctx context.Context, subscriber string, query cmtpubsub.Query, outCapacity ...int) (Subscription, error)
	Unsubscribe(ctx context.Context, subscriber string, query cmtpubsub.Query) error
	UnsubscribeAll(ctx context.Context, subscriber string) error

	NumClients() int
	NumClientSubscriptions(clientID string) int
}

type Subscription interface {
	Out() <-chan cmtpubsub.Message
	Canceled() <-chan struct{}
	Err() error
}

// EventBus is a common bus for all events going through the system. All calls
// are proxied to underlying pubsub server. All events must be published using
// EventBus to ensure correct data types.
type EventBus struct {
	service.BaseService
	pubsub *cmtpubsub.Server
}

// NewEventBus returns a new event bus.
func NewEventBus() *EventBus { _ = "STUB: not implemented"; return nil }

// NewEventBusWithBufferCapacity returns a new event bus with the given buffer capacity.
func NewEventBusWithBufferCapacity(cap int) *EventBus {
	_ = "STUB: not implemented"
	// capacity could be exposed later if needed
	return nil
}

func (b *EventBus) SetLogger(l log.Logger) { _ = "STUB: not implemented"; return }

func (b *EventBus) OnStart() error { _ = "STUB: not implemented"; return nil }

func (b *EventBus) OnStop() { _ = "STUB: not implemented"; return }

func (b *EventBus) NumClients() int { _ = "STUB: not implemented"; return 0 }

func (b *EventBus) NumClientSubscriptions(clientID string) int { _ = "STUB: not implemented"; return 0 }

func (b *EventBus) Subscribe(
	ctx context.Context,
	subscriber string,
	query cmtpubsub.Query,
	outCapacity ...int,
) (Subscription, error) {
	_ = "STUB: not implemented"
	return *new(Subscription), nil
}

// This method can be used for a local consensus explorer and synchronous
// testing. Do not use for for public facing / untrusted subscriptions!
func (b *EventBus) SubscribeUnbuffered(
	ctx context.Context,
	subscriber string,
	query cmtpubsub.Query,
) (Subscription, error) {
	_ = "STUB: not implemented"
	return *new(Subscription), nil
}

func (b *EventBus) Unsubscribe(ctx context.Context, subscriber string, query cmtpubsub.Query) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *EventBus) UnsubscribeAll(ctx context.Context, subscriber string) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *EventBus) Publish(eventType string, eventData TMEventData) error {
	_ = "STUB: not implemented"
	// no explicit deadline for publishing events
	return nil
}

// validateAndStringifyEvents takes a slice of event objects and creates a
// map of stringified events where each key is composed of the event
// type and each of the event's attributes keys in the form of
// "{event.Type}.{attribute.Key}" and the value is each attribute's value.
func (*EventBus) validateAndStringifyEvents(events []types.Event) map[string][]string {
	_ = "STUB: not implemented"
	return nil
}

func (b *EventBus) PublishEventNewBlock(data EventDataNewBlock) error {
	_ = "STUB: not implemented"
	// no explicit deadline for publishing events
	return nil
}

// add predefined new block event

func (b *EventBus) PublishEventSignedBlock(data EventDataSignedBlock) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *EventBus) PublishEventNewBlockEvents(data EventDataNewBlockEvents) error {
	_ = "STUB: not implemented"
	// no explicit deadline for publishing events
	return nil
}

// add predefined new block event

func (b *EventBus) PublishEventNewBlockHeader(data EventDataNewBlockHeader) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *EventBus) PublishEventNewEvidence(evidence EventDataNewEvidence) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *EventBus) PublishEventVote(data EventDataVote) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *EventBus) PublishEventValidBlock(data EventDataRoundState) error {
	_ = "STUB: not implemented"
	return nil
}

// PublishEventTx publishes tx event with events from Result. Note it will add
// predefined keys (EventTypeKey, TxHashKey). Existing events with the same keys
// will be overwritten.
func (b *EventBus) PublishEventTx(data EventDataTx) error {
	_ = "STUB: not implemented"
	// no explicit deadline for publishing events
	return nil
}

// add predefined compositeKeys

func (b *EventBus) PublishEventNewRoundStep(data EventDataRoundState) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *EventBus) PublishEventTimeoutPropose(data EventDataRoundState) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *EventBus) PublishEventTimeoutWait(data EventDataRoundState) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *EventBus) PublishEventNewRound(data EventDataNewRound) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *EventBus) PublishEventCompleteProposal(data EventDataCompleteProposal) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *EventBus) PublishEventPolka(data EventDataRoundState) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *EventBus) PublishEventUnlock(data EventDataRoundState) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *EventBus) PublishEventRelock(data EventDataRoundState) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *EventBus) PublishEventLock(data EventDataRoundState) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *EventBus) PublishEventValidatorSetUpdates(data EventDataValidatorSetUpdates) error {
	_ = "STUB: not implemented"
	return nil
}

// -----------------------------------------------------------------------------
type NopEventBus struct{}

func (NopEventBus) Subscribe(
	context.Context,
	string,
	cmtpubsub.Query,
	chan<- interface{},
) error {
	_ = "STUB: not implemented"
	return nil
}

func (NopEventBus) Unsubscribe(context.Context, string, cmtpubsub.Query) error {
	_ = "STUB: not implemented"
	return nil
}

func (NopEventBus) UnsubscribeAll(context.Context, string) error {
	_ = "STUB: not implemented"
	return nil
}

func (NopEventBus) PublishEventNewBlock(EventDataNewBlock) error {
	_ = "STUB: not implemented"
	return nil
}

func (NopEventBus) PublishEventNewBlockHeader(EventDataNewBlockHeader) error {
	_ = "STUB: not implemented"
	return nil
}

func (NopEventBus) PublishEventNewBlockEvents(EventDataNewBlockEvents) error {
	_ = "STUB: not implemented"
	return nil
}

func (NopEventBus) PublishEventSignedBlock(EventDataSignedBlock) error {
	_ = "STUB: not implemented"
	return nil
}

func (NopEventBus) PublishEventNewEvidence(EventDataNewEvidence) error {
	_ = "STUB: not implemented"
	return nil
}

func (NopEventBus) PublishEventVote(EventDataVote) error { _ = "STUB: not implemented"; return nil }

func (NopEventBus) PublishEventTx(EventDataTx) error { _ = "STUB: not implemented"; return nil }

func (NopEventBus) PublishEventNewRoundStep(EventDataRoundState) error {
	_ = "STUB: not implemented"
	return nil
}

func (NopEventBus) PublishEventTimeoutPropose(EventDataRoundState) error {
	_ = "STUB: not implemented"
	return nil
}

func (NopEventBus) PublishEventTimeoutWait(EventDataRoundState) error {
	_ = "STUB: not implemented"
	return nil
}

func (NopEventBus) PublishEventNewRound(EventDataRoundState) error {
	_ = "STUB: not implemented"
	return nil
}

func (NopEventBus) PublishEventCompleteProposal(EventDataRoundState) error {
	_ = "STUB: not implemented"
	return nil
}

func (NopEventBus) PublishEventPolka(EventDataRoundState) error {
	_ = "STUB: not implemented"
	return nil
}

func (NopEventBus) PublishEventUnlock(EventDataRoundState) error {
	_ = "STUB: not implemented"
	return nil
}

func (NopEventBus) PublishEventRelock(EventDataRoundState) error {
	_ = "STUB: not implemented"
	return nil
}

func (NopEventBus) PublishEventLock(EventDataRoundState) error {
	_ = "STUB: not implemented"
	return nil
}

func (NopEventBus) PublishEventValidatorSetUpdates(EventDataValidatorSetUpdates) error {
	_ = "STUB: not implemented"
	return nil
}
