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
	"github.com/dddsphere/martello/internal/system"
	"github.com/dddsphere/martello/subs/analytic"
	"github.com/dddsphere/martello/subs/cart"
	"github.com/dddsphere/martello/subs/catalog"
	"github.com/dddsphere/martello/subs/order"
	"github.com/dddsphere/martello/subs/user"
	"github.com/dddsphere/martello/subs/warehouse"
)

type App struct {
	sync.Mutex
	system.Worker
	system.Supervisor
	http *h.Server
	subs system.Subs
}

func NewApp(name, namespace string, log log.Logger) (app *App) {
	cfg := config.Load(namespace)

	app = &App{
		Worker: system.NewWorker(name,
			system.WithConfig(cfg),
			system.WithLogger(log)),
		http: h.NewServer("http-server", cfg, log),
		subs: system.Subs{},
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
	app.AddSub(analytic.NewService())
	app.AddSub(cart.NewService())
	app.AddSub(catalog.NewService())
	app.AddSub(order.NewService())
	app.AddSub(user.NewService())
	app.AddSub(warehouse.NewService())

	app.initSubs()
}

func (app *App) Start(ctx context.Context) error {
	app.Log().Infof("%s starting...", app.Name())
	defer app.Log().Infof("%s stopped", app.Name())

	err := app.startSubs()
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

func (app *App) AddSub(sub system.System) {
	app.Lock()
	app.subs.Add(sub)
	app.Unlock()
}

func (app *App) initSubs() {
	for _, sub := range app.subs.All() {
		sub.Init(app.Cfg(), app.Log())
	}
}

func (app *App) startSubs() error {
	ctx := app.Supervisor.Context()

	for _, sub := range app.subs.All() {
		err := sub.Start(ctx, app)
		if err != nil {
			return err
		}
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
