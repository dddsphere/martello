package domain

import (
	"github.com/dddsphere/martello/internal/event"
)

type CatalogInitialized struct {
	event.Event
	RegID string
}
