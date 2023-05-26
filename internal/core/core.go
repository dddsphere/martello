package core

import (
	"github.com/google/uuid"
)

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

func NewIdentifier(name string) Identifier {
	return Identifier{
		id:   genID(),
		name: name,
	}
}

func genID() string {
	return uuid.New().String()
}

func (i Identifier) ID() string {
	return i.id
}

func (i Identifier) SetID(id string, force ...bool) {
	if !(len(force) > 0 && force[0]) {
		return
	}

	i.id = id
}

func (i Identifier) GenID(force ...bool) (ok bool) {
	uid, err := uuid.NewUUID()
	if err != nil {
		return false
	}

	i.SetID(uid.String(), force...)
	return true
}

func (i Identifier) Name() string {
	return i.name
}

func (i Identifier) SetName(name string) {
	i.name = name
}

func (i Identifier) Equals(other ID) bool {
	return i.id == other.ID()
}
