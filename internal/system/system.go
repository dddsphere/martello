package system

import (
	"context"
	"sync"

	"github.com/dddsphere/martello/internal/config"
	"github.com/dddsphere/martello/internal/log"
)

type (
	System interface {
		Worker
		Init(context.Context, Service) error
		Shutdown(context.Context, Service) error
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
	Service interface {
		Cfg() *config.Config
		Log() log.Logger
	}
)
