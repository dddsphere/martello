package system

import (
	"context"
	"sync"
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
	BaseSystem struct {
	}
)

func NewSystem() *BaseSystem {
	return &BaseSystem{}
}

func (bs *BaseSystem) Init(ctx context.Context, s Service) error {
	return nil
}

func (bs *BaseSystem) Shutdown(ctx context.Context, s Service) error {
	return nil
}
