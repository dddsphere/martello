package app

import (
	"context"

	"github.com/dddsphere/martello/internal/config"
	"github.com/dddsphere/martello/internal/log"
)

type app struct {
	cfg        *config.Config
	log 			 log.Logger
	subsystems []Subsystem
}

func (app *app) Cfg() *config.Config {
	return app.cfg
}

func (app *app) Log() log.Logger {
	return app.log
}

type Subsystem interface {
	Setup(context.Context, Service) error
	Start(context.Context, Service) error
	Shutdown(context.Context, Service) error
}

type Service interface {
	Cfg() config.Config
	Log() log.Logger()
}
