package event

import "time"

// Event represents a generic event within the system.
type Event[T any] interface {
	ID() string
	Timestamp() time.Time
	Type() string
	Payload() T
}

// BaseEvent is a minimal implementation of the Event interface.
type BaseEvent[T any] struct {
	id        string
	timestamp time.Time
	eventType string
	payload   T
}

func NewBaseEvent[T any](id string, eventType string, payload T) *BaseEvent[T] {
	return &BaseEvent[T]{
		id:        id,
		timestamp: time.Now(),
		eventType: eventType,
		payload:   payload,
	}
}

func (e *BaseEvent[T]) ID() string           { return e.id }
func (e *BaseEvent[T]) Timestamp() time.Time { return e.timestamp }
func (e *BaseEvent[T]) Type() string         { return e.eventType }
func (e *BaseEvent[T]) Payload() T           { return e.payload }
