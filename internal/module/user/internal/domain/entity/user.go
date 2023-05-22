package entity

import (
	"github.com/dddsphere/martello/internal/module/user/internal/domain/vo"
)

type (
	User struct {
		username string
		email    string
		locale   vo.Locale
	}
)

func (u *User) Username() string {
	return u.username
}

func (u *User) SetUsername(username string) {
	u.username = username
}

func (u *User) Email() string {
	return u.email
}

func (u *User) SetEmail(email string) {
	u.email = email
}

func (u *User) Locale() vo.Locale {
	return u.locale
}

func (u *User) SetLocale(locale vo.Locale) {
	u.locale = locale
}
