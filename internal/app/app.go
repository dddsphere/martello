package app

import (
	"context"

	"github.com/dddsphere/martello/internal/config"
)

type app struct {
	cfg        *config.Config
	subsystems []Subsystem
}

func (app *app) Cfg() *config.Config {
	return app.cfg
}

type Subsystem interface {
	Setup(context.Context, Service) error
	Start(context.Context, Service) error
	Shutdown(context.Context, Service) error
}

type Service interface {
	Cfg() config.Config
}
