package system

import (
	"context"
	"sync"

	"github.com/dddsphere/martello/internal/config"
	"github.com/dddsphere/martello/internal/log"
)

type (
	System interface {
		Init(cfg *config.Config, log log.Logger)
		Start(context.Context, Service) error
		Stop(context.Context) error
		Shutdown(context.Context) error
	}

	Subs struct {
		mu   sync.Mutex
		list []System
	}
)

func (ss *Subs) Add(s System) {
	ss.mu.Lock()
	ss.list = append(ss.list, s)
	ss.mu.Unlock()
}

func (ss *Subs) All() []System {
	return ss.list
}

type (
	BaseSystem struct {
		*BaseWorker
	}
)

func NewSystem(name string, cfg *config.Config, log log.Logger) *BaseSystem {
	return &BaseSystem{
		BaseWorker: NewWorker(name, cfg, log),
	}
}

func (bs *BaseSystem) Start(ctx context.Context, s Service) error {
	bs.Log().Infof("Default init triggered")
	return nil
}

func (bs *BaseSystem) Shutdown(ctx context.Context) error {
	bs.Log().Infof("Default shutdown triggered")
	return nil
}
