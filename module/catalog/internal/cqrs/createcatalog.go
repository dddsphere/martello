package cqrs

import (
	"context"

	"github.com/pkg/errors"

	"github.com/dddsphere/martello/internal/event"
	"github.com/dddsphere/martello/module/catalog/internal/repo"
)

type CreateCatalog struct {
	ID    string
	RegID string
}

type CreateCatalogHandler struct {
	repo         repo.Catalog
	eventManager event.Manager[event.Event]
}

func NewCreateCatalogHandler(repo repo.Catalog, manager event.Manager[event.Event]) CreateCatalogHandler {
	return CreateCatalogHandler{
		repo:         repo,
		eventManager: manager,
	}
}

func (cch CreateCatalogHandler) CreateCatalog(ctx context.Context, cmd CreateCatalog) error {
	catalog, err := cch.repo.Get(ctx, cmd.ID)
	if err != nil {
		return err
	}

	evt, err := catalog.Init(cmd.RegID)
	if err != nil {
		return errors.Wrap(err, "create catalog command error")
	}

	err = cch.repo.Save(ctx, &catalog)
	if err != nil {
		return errors.Wrap(err, "order creation")
	}

	return cch.eventManager.Publish(ctx, evt)
}
