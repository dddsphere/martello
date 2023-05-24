package domain

import (
	"time"

	"github.com/google/uuid"
)

type (
	Catalog struct {
		id             ID
		name           string
		description    string
		active         bool
		releaseDate    time.Time
		expirationDate time.Time
	}
)

func (c *Catalog) Name() string {
	return c.name
}

func (c *Catalog) SetName(name string) {
	c.name = name
}

func (c *Catalog) Description() string {
	return c.description
}

func (c *Catalog) SetDescription(description string) {
	c.description = description
}

func (c *Catalog) Active() bool {
	return c.active
}

func (c *Catalog) SetActive(active bool) {
	c.active = active
}

func (c *Catalog) ReleaseDate() time.Time {
	return c.releaseDate
}

func (c *Catalog) SetReleaseDate(releaseDate time.Time) {
	c.releaseDate = releaseDate
}

func (c *Catalog) ExpirationDate() time.Time {
	return c.expirationDate
}

func (c *Catalog) SetExpirationDate(expirationDate time.Time) {
	c.expirationDate = expirationDate
}

type (
	ID struct {
		id string
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
