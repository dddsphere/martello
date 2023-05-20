package system

import (
	"net/http"

	"google.golang.org/grpc"

	"github.com/dddsphere/martello/internal/config"
	"github.com/dddsphere/martello/internal/log"
)

type (
	Service interface {
		RegisterHTTPHandler(handler http.Handler)
		RegisterGRPCServer(server *grpc.Server)
	}

	BaseService struct {
		*BaseSystem
	}
)

func NewService(name string, cfg *config.Config, log log.Logger) *BaseService {
	return &BaseService{
		BaseSystem: NewSystem(name, cfg, log),
	}
}

type IgnoreUnimplementedRegistration struct{}

var _ Service = (*IgnoreUnimplementedRegistration)(nil)

func (IgnoreUnimplementedRegistration) RegisterHTTPHandler(handler http.Handler) {}

func (IgnoreUnimplementedRegistration) RegisterGRPCServer(server *grpc.Server) {}
