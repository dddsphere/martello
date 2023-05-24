package http

import (
	"github.com/dddsphere/martello/internal/system"
	"github.com/dddsphere/martello/module/user/internal/domain/service"
)

type (
	Endpoint struct {
		*system.BaseWorker
		service *service.User
	}
)
