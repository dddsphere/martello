package domain

import (
	"time"

	"github.com/dddsphere/martello/internal/domain"
)

type (
	Catalog struct {
		*domain.BaseAggregate
		description    string
		active         bool
		releaseDate    time.Time
		expirationDate time.Time
	}
)

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
