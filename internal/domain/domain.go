package domain

import "github.com/google/uuid"

type (
	ID struct {
		id   string
		name string
	}
)

type (
	Aggregate struct {
		id   ID
		name string
	}
)

func (i ID) ID() string {
	return i.id
}

func (i ID) SetID(id string, force ...bool) {
	if !(len(force) > 0 && force[0]) {
		return
	}

	i.id = id
}

func (i ID) GenID(force ...bool) (ok bool) {
	uid, err := uuid.NewUUID()
	if err != nil {
		return false
	}

	i.SetID(uid.String(), force...)
	return true
}

func (a *Aggregate) Name() string {
	return a.name
}

func (a *Aggregate) SetName(name string) {
	a.name = name
}
