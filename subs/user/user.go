package user

import (
	http2 "net/http"

	"google.golang.org/grpc"

	"github.com/dddsphere/martello/internal/system"
)

type (
	User struct {
		*system.BaseSystem
	}
)

func (u *User) RegisterHTTPHandler(h http2.Handler) {
	u.Log().Infof("No registered HTTP handlers for %s", u.Name())
}

func (u *User) RegisterGRPCServer(srv grpc.Server) {
	u.Log().Infof("No registered gRPC servers for %s", u.Name())
}
