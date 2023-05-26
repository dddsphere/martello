package domain

import (
	"github.com/pkg/errors"

	"github.com/dddsphere/martello/internal/domain"
)

var (
	ErrEmptyName     = errors.Wrap(domain.ErrDomain, "product name cannot be blank")
	ErrEmptyQuantity = errors.Wrap(domain.ErrDomain, "product quantity cannot be minor than 1")
)
