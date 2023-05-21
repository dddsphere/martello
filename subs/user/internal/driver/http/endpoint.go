package http

import (
	"github.com/dddsphere/martello/internal/system"
	"github.com/dddsphere/martello/subs/user/internal/domain/service"
)

type (
	Endpoint struct {
		*system.BaseWorker
		service *service.User
	}
)
