package app

import (
	"context"
	"fmt"
	"sync"

	"github.com/dddsphere/martello/internal/driver/http"

	"github.com/dddsphere/martello/internal/config"
	"github.com/dddsphere/martello/internal/log"
	"github.com/dddsphere/martello/internal/system"
)

type App struct {
	system.Worker
	system.Supervisor
	http *http.Server
	subs Subsystems
}

func NewApp(name, namespace string, log log.Logger) (app *App) {
	cfg := config.Load(namespace)

	app = &App{
		Worker: system.NewWorker(name, cfg, log),
		http:   http.NewServer("http-server", cfg, log),
		subs:   Subsystems{},
	}

	app.EnableSupervisor()

	return app
}

func (app *App) Name() string {
	return app.Worker.Name()
}

func (app *App) Cfg() *config.Config {
	return app.Worker.Cfg()
}

func (app *App) Log() log.Logger {
	return app.Worker.Log()
}

func (app *App) EnableSupervisor() {
	name := fmt.Sprintf("%s-supervisor", app.Name())
	app.Supervisor = system.NewSupervisor(name, app.Cfg(), app.Log(), true)
}

func (app *App) Init(ctx context.Context) {
	//app.subs.Add(a)
}

func (app *App) Run() (err error) {
	err = app.startSubsystems()
	if err != nil {
		return err
	}

	app.Log().Infof("%s started", app.Name())
	defer app.Log().Infof("%s stoped", app.Name())

	app.Supervisor.AddTasks(
		app.http.StartTask(context.Background()),
	)

	return app.Supervisor.Wait()
}

func (app *App) startSubsystems() error {
	ctx := app.Supervisor.Context()

	for _, sub := range app.subs.All() {
		err := sub.Start(ctx, app)
		if err != nil {
			return err
		}
	}

	return nil
}

type (
	Subsystem interface {
		Setup(context.Context, Service) error
		Start(context.Context, Service) error
		Shutdown(context.Context, Service) error
	}

	Subsystems struct {
		mu   sync.Mutex
		list []Subsystem
	}
)

func (ss *Subsystems) Add(s Subsystem) {
	ss.mu.Lock()
	ss.list = append(ss.list, s)
	ss.mu.Unlock()
}

func (ss *Subsystems) All() []Subsystem {
	return ss.list
}

type (
	Service interface {
		Cfg() *config.Config
		Log() log.Logger
	}
)
