package analytic

import (
	"net/http"

	"google.golang.org/grpc"

	"github.com/dddsphere/martello/internal/system"
)

type (
	Service struct {
		*system.BaseService
	}
)

const (
	name = "analytic-service"
)

func NewService(opts ...system.Option) *Service {
	name := system.WithSuffix(name, 8)
	return &Service{
		BaseService: system.NewService(name, opts...),
	}
}

func (s *Service) RegisterHTTPHandler(h http.Handler) {
	s.Log().Infof("No registered HTTP handlers for %s", s.Name())
}

func (s *Service) RegisterGRPCServer(srv grpc.Server) {
	s.Log().Infof("No registered gRPC servers for %s", s.Name())
}
