package service

import (
	"github.com/dddsphere/martello/internal/module/user/internal/domain/entity"
	"github.com/dddsphere/martello/internal/system"
)

type (
	User struct {
		*system.BaseWorker
		// NOTE: Add required dependencies
	}
)

// WIP: Accepting User entity for now, should be a user data object.
func (svc *User) Create(u *entity.User) error {
	svc.Log().Infof("Creating the user '%s'", u.Username())
	return nil
}
