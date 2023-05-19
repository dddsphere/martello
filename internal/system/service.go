package system

import (
	"net/http"

	"google.golang.org/grpc"
)

type (
	Service interface {
		RegisterHTTPHandler(handler http.Handler)
		RegisterGRPCServer(server *grpc.Server)
	}
)
