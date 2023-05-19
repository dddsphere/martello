package user

import (
	http2 "net/http"

	"google.golang.org/grpc"

	"github.com/dddsphere/martello/internal/system"
)

type (
	User struct {
		*system.BaseWorker
		*system.BaseSystem
	}
)

func (u *User) RegisterHTTPHandler(h http2.Handler) {
	panic("not implemented yet")
}

func (u *User) RegisterGRPCServer(srv grpc.Server) {
	panic("not implemented yet")
}
