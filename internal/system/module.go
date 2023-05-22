package system

import (
	"context"
	"sync"

	"github.com/dddsphere/martello/internal/config"
	"github.com/dddsphere/martello/internal/log"
)

type (
	Module interface {
		Init(cfg *config.Config, log log.Logger)
		Start(context.Context, Service) error
		Stop(context.Context) error
		Shutdown(context.Context) error
	}

	Modules struct {
		mu   sync.Mutex
		list []Module
	}
)

func (mm *Modules) Add(s Module) {
	mm.mu.Lock()
	mm.list = append(mm.list, s)
	mm.mu.Unlock()
}

func (mm *Modules) All() []Module {
	return mm.list
}

type (
	BaseModule struct {
		*BaseWorker
	}
)

func NewSystem(name string, opts ...Option) *BaseModule {
	return &BaseModule{
		BaseWorker: NewWorker(name, opts...),
	}
}

func (bs *BaseModule) SetCfg(cfg *config.Config) {
	bs.cfg = cfg
}

func (bs *BaseModule) SetLog(log log.Logger) {
	bs.log = log
}

func (bs *BaseModule) Start(ctx context.Context, s Service) error {
	bs.Log().Infof("%s default init", bs.Name())
	return nil
}

func (bs *BaseModule) Shutdown(ctx context.Context) error {
	bs.Log().Infof("%s default shutdown", bs.Name())
	return nil
}
