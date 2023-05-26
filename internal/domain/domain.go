package domain

import (
	"sync"

	"github.com/dddsphere/martello/internal/core"
	"github.com/dddsphere/martello/internal/event"
)

type (
	Aggregate interface {
		*core.ID
		event.Processor
	}

	BaseAggregate struct {
		id core.ID
		sync.Map
	}
)

func (ba *BaseAggregate) Name() string {
	return ba.id.Name()
}

func NewAggregate(id, name string) *BaseAggregate {
	return &BaseAggregate{
		id: core.NewIdentifier(name),
	}
}

func (ba *BaseAggregate) AddEvent(name string, payload any) {
	ba.Store(name, payload)
}

func (ba *BaseAggregate) Events() map[string]event.Event {
	events := make(map[string]event.Event)

	ba.Map.Range(func(key, value interface{}) bool {
		if strKey, ok := key.(string); ok {
			if e, ok := value.(event.Event); ok {
				events[strKey] = e
			}
		}
		return true
	})

	return events
}

func (ba *BaseAggregate) Reset() {
	ba.Map.Range(func(key, value interface{}) bool {
		ba.Map.Delete(key)
		return true
	})
}

func (ba *BaseAggregate) Apply() {
	ba.Map.Range(func(key, value interface{}) bool {
		ba.Map.Delete(key)
		return true
	})
}
