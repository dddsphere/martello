package app

import (
	"context"
	"fmt"
	"net/http"

	"google.golang.org/grpc"

	h "github.com/dddsphere/martello/internal/driver/http"
	"github.com/dddsphere/martello/subs/user"

	"github.com/dddsphere/martello/internal/config"
	"github.com/dddsphere/martello/internal/log"
	"github.com/dddsphere/martello/internal/system"
)

type App struct {
	system.Worker
	system.Supervisor
	http *h.Server
	subs system.Subs
}

func NewApp(name, namespace string, log log.Logger) (app *App) {
	cfg := config.Load(namespace)

	app = &App{
		Worker: system.NewWorker(name, cfg, log),
		http:   h.NewServer("http-server", cfg, log),
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
	app.subs.Add(user.User{})
}

func (app *App) Run() (err error) {
	err = app.startSubsystems()
	if err != nil {
		return err
	}

	app.Log().Infof("%s started", app.Name())
	defer app.Log().Infof("%s stopped", app.Name())

	app.Supervisor.AddTasks(
		app.http.Start,
	)

	return app.Supervisor.Wait()
}

func (app *App) startSubsystems() error {
	ctx := app.Supervisor.Context()

	for _, sub := range app.subs.All() {
		err := sub.Init(ctx, app)
		if err != nil {
			return err
		}
	}

	return nil
}

func (app *App) RegisterHTTPHandler(handler http.Handler) {
	panic("not implemented yet")
}

func (app *App) RegisterGRPCServer(server *grpc.Server) {
	panic("not implemented yet")
}
