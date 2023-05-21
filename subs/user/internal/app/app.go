package app

import (
	"github.com/dddsphere/martello/internal/config"
	h "github.com/dddsphere/martello/internal/driver/http"
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
		Worker: system.NewWorker(name,
			system.WithConfig(cfg),
			system.WithLogger(log)),
		http: h.NewServer("http-server", cfg, log),
	}

	return app
}
