package app

import (
	"context"

	"github.com/dddsphere/martello/internal/config"
	"github.com/dddsphere/martello/internal/log"
	"github.com/dddsphere/martello/internal/system"
)

type app struct {
	system.Worker
	subs []Subsystem
}

type Subsystem interface {
	Setup(context.Context, Service) error
	Start(context.Context, Service) error
	Shutdown(context.Context, Service) error
}

type Service interface {
	Cfg() *config.Config
	Log() log.Logger
}
