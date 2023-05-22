package http

import (
	"github.com/dddsphere/martello/internal/module/user/internal/domain/service"
	"github.com/dddsphere/martello/internal/system"
)

type (
	Endpoint struct {
		*system.BaseWorker
		service *service.User
	}
)
