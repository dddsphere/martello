package domain

import (
	"sync"
	"time"

	"github.com/google/uuid"
)

type (
	Aggregate interface {
		*ID
		EventManager
	}

	BaseAggregate struct {
		id ID
		sync.Map
	}
)

func NewAggregate(id, name string) *BaseAggregate {
	return &BaseAggregate{
		id: NewIdentifier(id, name),
	}
}

func (ba *BaseAggregate) AddEvent(name string, payload any) {
	ba.Store(name, payload)
}

func (ba *BaseAggregate) Events() map[string]Event {
	events := make(map[string]Event)

	ba.Map.Range(func(key, value interface{}) bool {
		if strKey, ok := key.(string); ok {
			if event, ok := value.(Event); ok {
				events[strKey] = event
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
	panic("not implemented yet")
}

type (
	EventManager interface {
		AddEvent(name string, payload any)
		Events() map[string]Event
		Reset()
		Apply()
	}

	Event interface {
		ID
		Payload() any
		Timestamp() time.Time
	}
)

func (ba *BaseAggregate) Name() string {
	return ba.id.Name()
}

type (
	ID interface {
		ID() string
		Name() string
		Equals(other ID) bool
	}

	Identifier struct {
		id   string
		name string
	}
)

func NewIdentifier(id, name string) *Identifier {
	return &Identifier{
		id:   id,
		name: name,
	}
}

func (i *Identifier) ID() string {
	return i.id
}

func (i *Identifier) SetID(id string, force ...bool) {
	if !(len(force) > 0 && force[0]) {
		return
	}

	i.id = id
}

func (i *Identifier) GenID(force ...bool) (ok bool) {
	uid, err := uuid.NewUUID()
	if err != nil {
		return false
	}

	i.SetID(uid.String(), force...)
	return true
}

func (i *Identifier) Name() string {
	return i.name
}

func (i *Identifier) SetName(name string) {
	i.name = name
}

func (i *Identifier) Equals(other ID) bool {
	return i.id == other.ID()
}
