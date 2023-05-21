package http

import (
	"fmt"
	"net/http"
)

const (
	userRes = "user"
)

func (ep *Endpoint) InitSignUpUser(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "Nothing meaningful for now!")

	if err != nil {
		ep.Log().Error(err.Error())
	}
}
