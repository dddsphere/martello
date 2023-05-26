package eventsource

import (
	"github.com/dddsphere/martello/internal/domain"
	"github.com/dddsphere/martello/internal/event"
)

type (
	Aggregate interface {
		domain.Aggregate
		Seal()
		Current() int
		SetVersion(int)
		Awaiting() int
	}

	BaseAggregate struct {
		domain.Aggregate
		version int
	}
)

func NewAggregate(name string) BaseAggregate {
	return BaseAggregate{
		Aggregate: domain.NewAggregate(name),
		version:   0,
	}
}

func (ba BaseAggregate) AddEvent(name string, payload event.Payload) {
	ba.Aggregate.AddEvent(name, payload)
}

func (ba BaseAggregate) Seal() {
	ba.version = ba.newVersion()
	ba.Reset()
}

func (ba BaseAggregate) newVersion() int {
	return len(ba.Events()) + 1
}

func (ba BaseAggregate) Current() int {
	return ba.version
}

func (ba BaseAggregate) SetVersion(version int) {
	ba.version = version
}

func (ba BaseAggregate) Awaiting() int {
	return ba.Current() + len(ba.Events())
}
