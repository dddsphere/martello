package system

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

type (
	Task     func(ctx context.Context) error
	Teardown func()
)

type (
	Supervisor interface {
		Add(tasks ...Task)
		Teardown(tasks ...Teardown)
		Wait() error
		Context() context.Context
		CancelFunc() context.CancelFunc
	}

	supervisor struct {
		tasks    []Task
		teardown []Teardown
		ctx      context.Context
		cancel   context.CancelFunc
	}

	SupervisorCfg func(cfg *supervisorCfg)

	supervisorCfg struct {
		parentCtx context.Context
		notify    bool
	}
)

func NewSupervisor(configs ...SupervisorCfg) Supervisor {
	cfg := &supervisorCfg{
		parentCtx: context.Background(),
		notify:    false,
	}

	for _, apply := range configs {
		apply(cfg)
	}

	sv := &supervisor{
		tasks:    []Task{},
		teardown: []Teardown{},
	}

	sv.ctx, sv.cancel = context.WithCancel(cfg.parentCtx)
	if cfg.notify {
		sv.ctx, sv.cancel = signal.NotifyContext(sv.ctx, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	}

	return sv
}

func (sv *supervisor) Add(tasks ...Task) {
	sv.tasks = append(sv.tasks, tasks...)
}

func (sv *supervisor) Teardown(tasks ...Teardown) {
	sv.teardown = append(sv.teardown, tasks...)
}

func (sv *supervisor) Wait() (err error) {
	eg, ctx := errgroup.WithContext(sv.ctx)
	eg.Go(sv.contextDone(ctx))

	for _, t := range sv.tasks {
		task := t
		eg.Go(func() error { return task(ctx) })
	}

	for _, tt := range sv.teardown {
		teardown := tt
		defer teardown()
	}

	return eg.Wait()
}

func (sv *supervisor) contextDone(ctx context.Context) func() error {
	return func() error {
		<-ctx.Done()
		sv.cancel()
		return nil
	}
}

func (sv *supervisor) Context() context.Context {
	return sv.ctx
}

func (sv *supervisor) CancelFunc() context.CancelFunc {
	return sv.cancel
}
