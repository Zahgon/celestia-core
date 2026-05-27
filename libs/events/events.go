// Package events - Pub-Sub in go with event caching
package events

import (
	"github.com/cometbft/cometbft/libs/service"
	cmtsync "github.com/cometbft/cometbft/libs/sync"
)

// ErrListenerWasRemoved is returned by AddEvent if the listener was removed.
type ErrListenerWasRemoved struct {
	listenerID string
}

// Error implements the error interface.
func (e ErrListenerWasRemoved) Error() string { _ = "STUB: not implemented"; return "" }

// EventData is a generic event data can be typed and registered with
// tendermint/go-amino via concrete implementation of this interface.
type EventData interface{}

// Eventable is the interface reactors and other modules must export to become
// eventable.
type Eventable interface {
	SetEventSwitch(evsw EventSwitch)
}

// Fireable is the interface that wraps the FireEvent method.
//
// FireEvent fires an event with the given name and data.
type Fireable interface {
	FireEvent(event string, data EventData)
}

// EventSwitch is the interface for synchronous pubsub, where listeners
// subscribe to certain events and, when an event is fired (see Fireable),
// notified via a callback function.
//
// Listeners are added by calling AddListenerForEvent function.
// They can be removed by calling either RemoveListenerForEvent or
// RemoveListener (for all events).
type EventSwitch interface {
	service.Service
	Fireable

	AddListenerForEvent(listenerID, event string, cb EventCallback) error
	RemoveListenerForEvent(event string, listenerID string)
	RemoveListener(listenerID string)
}

type eventSwitch struct {
	service.BaseService

	mtx        cmtsync.RWMutex
	eventCells map[string]*eventCell
	listeners  map[string]*eventListener
}

func NewEventSwitch() EventSwitch { _ = "STUB: not implemented"; return *new(EventSwitch) }

func (evsw *eventSwitch) OnStart() error { _ = "STUB: not implemented"; return nil }

func (evsw *eventSwitch) OnStop() { _ = "STUB: not implemented"; return }

func (evsw *eventSwitch) AddListenerForEvent(listenerID, event string, cb EventCallback) error {
	_ = "STUB: not implemented"
	// Get/Create eventCell and listener.
	return nil
}

// Add event and listener.

func (evsw *eventSwitch) RemoveListener(listenerID string) {
	_ = "STUB: not implemented"
	// Get and remove listener.
	return
}

// Remove callback for each event.

func (evsw *eventSwitch) RemoveListenerForEvent(event string, listenerID string) {
	_ = "STUB: not implemented"
	// Get eventCell
	return
}

// Remove listenerID from eventCell

// Maybe garbage collect eventCell.

// Lock again and double check.
// OUTER LOCK
// INNER LOCK

// INNER LOCK
// OUTER LOCK

func (evsw *eventSwitch) FireEvent(event string, data EventData) {
	_ = "STUB: not implemented"
	// Get the eventCell
	return
}

// Fire event for all listeners in eventCell

//-----------------------------------------------------------------------------

// eventCell handles keeping track of listener callbacks for a given event.
type eventCell struct {
	mtx       cmtsync.RWMutex
	listeners map[string]EventCallback
}

func newEventCell() *eventCell { _ = "STUB: not implemented"; return nil }

func (cell *eventCell) AddListener(listenerID string, cb EventCallback) {
	_ = "STUB: not implemented"
	return
}

func (cell *eventCell) RemoveListener(listenerID string) int { _ = "STUB: not implemented"; return 0 }

func (cell *eventCell) FireEvent(data EventData) { _ = "STUB: not implemented"; return }

//-----------------------------------------------------------------------------

type EventCallback func(data EventData)

type eventListener struct {
	id string

	mtx     cmtsync.RWMutex
	removed bool
	events  []string
}

func newEventListener(id string) *eventListener { _ = "STUB: not implemented"; return nil }

func (evl *eventListener) AddEvent(event string) error { _ = "STUB: not implemented"; return nil }

func (evl *eventListener) GetEvents() []string { _ = "STUB: not implemented"; return nil }

func (evl *eventListener) SetRemoved() { _ = "STUB: not implemented"; return }
