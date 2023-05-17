package rest

import (
	"net/http"

	"github.com/dddsphere/martello/internal/config"
	"github.com/dddsphere/martello/internal/log"
	"github.com/dddsphere/martello/internal/system"
)

type (
	Server struct {
		system.Worker
		router http.Handler
	}
)

func NewServer(name string, cfg *config.Config, log log.Logger) (server *Server) {
	return &Server{
		Worker: system.NewWorker(name, cfg, log),
	}
}

func (srv *Server) SetRouter(router http.Handler) {
	srv.router = router
}
