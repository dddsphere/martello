package app

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"google.golang.org/grpc"

	"github.com/dddsphere/martello/internal/config"
	h "github.com/dddsphere/martello/internal/driver/http"
	"github.com/dddsphere/martello/internal/log"
	"github.com/dddsphere/martello/internal/module/analytic"
	"github.com/dddsphere/martello/internal/module/cart"
	"github.com/dddsphere/martello/internal/module/catalog"
	"github.com/dddsphere/martello/internal/module/order"
	"github.com/dddsphere/martello/internal/module/user"
	"github.com/dddsphere/martello/internal/module/warehouse"
	"github.com/dddsphere/martello/internal/system"
)

type App struct {
	sync.Mutex
	system.Worker
	system.Supervisor
	http    *h.Server
	modules system.Modules
}

func NewApp(name, namespace string, log log.Logger) (app *App) {
	cfg := config.Load(namespace)

	app = &App{
		Worker: system.NewWorker(name,
			system.WithConfig(cfg),
			system.WithLogger(log)),
		http:    h.NewServer("http-server", cfg, log),
		modules: system.Modules{},
	}

	app.EnableSupervisor()

	return app
}

func (app *App) EnableSupervisor() {
	name := fmt.Sprintf("%s-supervisor", app.Name())
	app.Supervisor = system.NewSupervisor(name, app.Cfg(), app.Log(), true)
}

func (app *App) Run() (err error) {
	ctx := context.Background()
	app.Init(ctx)
	return app.Start(ctx)
}

func (app *App) Init(ctx context.Context) {
	app.AddModule(analytic.NewService())
	app.AddModule(cart.NewService())
	app.AddModule(catalog.NewService())
	app.AddModule(order.NewService())
	app.AddModule(user.NewService())
	app.AddModule(warehouse.NewService())

	app.initModules()
}

func (app *App) Start(ctx context.Context) error {
	app.Log().Infof("%s starting...", app.Name())
	defer app.Log().Infof("%s stopped", app.Name())

	err := app.startModules()
	if err != nil {
		return err
	}

	app.Log().Infof("%s started!", app.Name())

	app.Supervisor.AddTasks(
		app.http.Start,
		//app.grpc.Start,
	)

	return app.Supervisor.Wait()
}

func (app *App) AddModule(sub system.Module) {
	app.Lock()
	app.modules.Add(sub)
	app.Unlock()
}

func (app *App) initModules() {
	mods := app.modules.All()
	wg := sync.WaitGroup{}
	wg.Add(len(mods))

	for _, mod := range app.modules.All() {
		go func(m system.Module) {
			m.Init(app.Cfg(), app.Log())
			wg.Done()
		}(mod)
	}

	wg.Wait()
}

func (app *App) startModules() error {
	mods := app.modules.All()
	wg := sync.WaitGroup{}
	wg.Add(len(mods))

	ctx := app.Supervisor.Context()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	errCh := make(chan error)

	for _, mod := range app.modules.All() {
		go func(m system.Module) {
			err := m.Start(ctx, app)
			if err != nil {
				errCh <- err
				cancel()
			}
		}(mod)
	}

	return nil
}

// Service interface

func (app *App) RegisterHTTPHandler(http.Handler) {
	app.Log().Infof("No registered HTTP handlers for %s", app.Name())
}

func (app *App) RegisterGRPCServer(srv *grpc.Server) {
	app.Log().Infof("No registered gRPC servers for %s", app.Name())
}

// Worker interface
func (app *App) Log() log.Logger {
	return app.Worker.Log()
}

// Worker interface

func (app *App) Name() string {
	return app.Worker.Name()
}

func (app *App) Cfg() *config.Config {
	return app.Worker.Cfg()
}
