package domain

import (
	"time"

	"github.com/google/uuid"
)

type (
	Aggregate interface {
		ID
		EventManager
	}

	BaseAggregate struct {
		id   ID
		name string
	}
)

type (
	EventManager interface {
		AddEvent(name string, payload any)
		Events() []Event
		Reset()
		Apply()
	}

	Event interface {
		ID
		Payload() any
		Timestamp() time.Time
	}
)

func (a *BaseAggregate) Name() string {
	return a.name
}

type (
	ID interface {
		ID()
		Name()
		Equals(other ID) bool
	}

	Identifier struct {
		id   string
		name string
	}
)

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

func (i *Identifier) Equals(other *Identifier) bool {
	return i.id == other.id
}
