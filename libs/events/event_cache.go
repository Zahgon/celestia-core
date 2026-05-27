package events

// An EventCache buffers events for a Fireable
// All events are cached. Filtering happens on Flush
type EventCache struct {
	evsw   Fireable
	events []eventInfo
}

// Create a new EventCache with an EventSwitch as backend
func NewEventCache(evsw Fireable) *EventCache { _ = "STUB: not implemented"; return nil }

// a cached event
type eventInfo struct {
	event string
	data  EventData
}

// Cache an event to be fired upon finality.
func (evc *EventCache) FireEvent(event string, data EventData) {
	_ = "STUB: not implemented"
	// append to list (go will grow our backing array exponentially)
	return
}

// Fire events by running evsw.FireEvent on all cached events. Blocks.
// Clears cached events
func (evc *EventCache) Flush() { _ = "STUB: not implemented"; return }

// Clear the buffer, since we only add to it with append it's safe to just set it to nil and maybe safe an allocation
