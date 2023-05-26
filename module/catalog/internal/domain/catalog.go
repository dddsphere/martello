package domain

import (
	"time"

	"github.com/dddsphere/martello/internal/core"
	"github.com/dddsphere/martello/internal/domain"
	"github.com/dddsphere/martello/internal/event"
)

type (
	Catalog struct {
		*domain.BaseAggregate
		regID          string
		items          []Item
		description    string
		status         string
		releaseDate    time.Time
		expirationDate time.Time
	}
)

func NewCatalog(id string) *Catalog {
	return &Catalog{
		BaseAggregate: domain.NewAggregate(id, aggregates.Catalog),
		status:        status.Initialized.Name(),
	}
}

func (c *Catalog) Init(regID string) (event.Event, error) {
	c.AddEvent(status.Initialized.Name(), &CatalogInitialized{
		RegID: regID,
	})

	return event.NewEvent(status.Initialized.Name(), c), nil
}

func (c *Catalog) Description() string {
	return c.description
}

func (c *Catalog) SetDescription(description string) {
	c.description = description
}

func (c *Catalog) Status() string {
	return c.status
}

func (c *Catalog) SetStatus(status string) {
	c.status = status
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
		core.ID
		description    string
		active         bool
		releaseDate    time.Time
		expirationDate time.Time
	}
)
