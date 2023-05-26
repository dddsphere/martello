package repo

import (
	"context"

	"github.com/dddsphere/martello/module/catalog/internal/domain"
)

type (
	Catalog interface {
		Get(ctx context.Context, id string) (domain.Catalog, error)
		Save(ctx context.Context, c *domain.Catalog) error
	}
)
