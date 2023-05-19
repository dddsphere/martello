package system

import (
	"net/http"
	"sync"

	"google.golang.org/grpc"
)

type (
	registry struct {
		httpHandlers []http.Handler
		grpcServers  []*grpc.Server
	}
)

var lock = &sync.Mutex{}
var instance *registry

func Instance() *registry {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()

		if instance == nil {
			instance = &registry{}
		}
	}

	return instance
}

func (r *registry) AddHTTPHandler(h http.Handler) {
	r.httpHandlers = append(r.httpHandlers, h)
}

func (r *registry) AddGRPCServer(s *grpc.Server) {
	r.grpcServers = append(r.grpcServers, s)
}
