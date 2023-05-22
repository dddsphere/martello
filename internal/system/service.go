package system

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

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
		*BaseModule
	}
)

var (
	runes = []rune("abcdefghijklmnopqrstuvwxyz0123456789")
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func NewService(name string, opts ...Option) *BaseService {
	return &BaseService{
		BaseModule: NewSystem(name, opts...),
	}
}

func (bs *BaseService) SetCfg(cfg *config.Config) {
	bs.BaseWorker.SetCfg(cfg)
}

func (bs *BaseService) SetLog(log log.Logger) {
	bs.BaseWorker.SetLog(log)
}

func (bs *BaseService) Init(cfg *config.Config, log log.Logger) {
	bs.SetCfg(cfg)
	bs.SetLog(log)
}

func WithSuffix(name string, n int) string {
	suffix := make([]rune, n)
	for i := range suffix {
		suffix[i] = runes[rand.Intn(len(runes))]
	}
	return fmt.Sprintf("%s-%s", name, string(suffix))
}

type IgnoreUnimplementedRegistration struct{}

var _ Service = (*IgnoreUnimplementedRegistration)(nil)

func (IgnoreUnimplementedRegistration) RegisterHTTPHandler(handler http.Handler) {}

func (IgnoreUnimplementedRegistration) RegisterGRPCServer(server *grpc.Server) {}
