package event

import (
	"context"
	"time"

	"github.com/dddsphere/martello/internal/core"
)

type (
	Processor interface {
		AddEvent(name string, payload Payload)
		Events() map[string]Event
		Reset()
	}

	Event interface {
		core.ID
		Payload() Payload
		Timestamp() time.Time
	}

	Payload interface {
	}

	Base struct {
		core.Identifier
		payload   any
		timestamp time.Time
	}
)

func NewEvent(name string, payload Payload) *Base {
	return &Base{
		Identifier: core.NewIdentifier(name),
		payload:    payload,
		timestamp:  time.Time{},
	}
}

func (b *Base) Payload() Payload { // payload should be a type
	return b.payload
}

func (b *Base) Timestamp() time.Time {
	return b.timestamp
}

type (
	Manager[T Event] interface {
		Subscriber[Event]
		Publisher[Event]
	}

	Subscriber[T Event] interface {
		Subscribe(handler Handler[T], events ...string)
	}

	Publisher[T Event] interface {
		Publish(ctx context.Context, events ...T) error
	}

	Handler[T Event] interface {
		Handle(ctx context.Context, event T) error
	}
)
