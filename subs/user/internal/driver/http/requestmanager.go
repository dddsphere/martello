package http

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"

	"github.com/dddsphere/martello/internal/config"
	"github.com/dddsphere/martello/internal/log"
	"github.com/dddsphere/martello/internal/system"
	"github.com/dddsphere/martello/subs/user/internal/domain/entity"
	ds "github.com/dddsphere/martello/subs/user/internal/domain/service"
)

type (
	RequestManager struct {
		*system.BaseWorker
		domainService *ds.User
	}
)

func NewRequestManager(cfg *config.Config, log log.Logger) (rm *RequestManager) {
	return &RequestManager{
		BaseWorker: system.NewWorker("request-manager",
			system.WithConfig(cfg),
			system.WithLogger(log)),
	}
}

// Dispatch commands
// Commands will return an HTTP 202 including headers with relevant metadata (i.e.: request ID).
func (rm *RequestManager) Dispatch(w http.ResponseWriter, r *http.Request, commandName string) {
	reqID := genReqID(r)
	var err error

	switch commandName {
	case "sign-up-user":
		// NOTE: Using a domain entity for now
		// Should be a DO (DataObject)
		user := rm.reqToUser(r)
		err := rm.domainService.Create(&user)
		if err != nil {
			rm.Log().Errorf("Req '%s' error: %s", reqID, err.Error())
		}

	default:
		err := fmt.Errorf("command '%s' not found", commandName)
		rm.Error(err, w)
	}

	if err != nil {
		rm.Log().Errorf("Dispatch command error: %w", err)
	}

	w.WriteHeader(http.StatusAccepted)
	_, err = w.Write([]byte("202 - Temporary OK message"))
	if err != nil {
		rm.Log().Errorf("Dispatch command error: %w", err)
	}
}

func (rm *RequestManager) Error(err error, w http.ResponseWriter) {
	rm.Log().Error(err.Error())
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func (rm *RequestManager) reqToUser(r *http.Request) entity.User {
	// WIP: Create User from request
	return entity.User{}
}

func genReqID(r *http.Request) (id string) {
	id = r.Header.Get("X-Request-ID")
	if id == "" {
		return uuid.New().String()
	}

	return id
}
