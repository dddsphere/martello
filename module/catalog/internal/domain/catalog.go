package domain

import (
	"time"

	"github.com/dddsphere/martello/internal/domain"
)

type (
	Catalog struct {
		*domain.BaseAggregate
		items          []Item
		description    string
		active         bool
		releaseDate    time.Time
		expirationDate time.Time
	}
)

func NewCatalog(id string) *Catalog {
	return &Catalog{
		BaseAggregate: domain.NewAggregate(id, aggregates.Catalog),
		items:         []Item{},
	}
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

func (c *Catalog) AddItem(i *Item) {
	panic("not implemented yet")
}

func (c *Catalog) RemoveItem(id string) {
	panic("not implemented yet")
}

func (c *Catalog) Empty() {
	panic("not implemented yet")
}

type (
	Item struct {
		domain.ID
		description    string
		active         bool
		releaseDate    time.Time
		expirationDate time.Time
	}
)
