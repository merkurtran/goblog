package auth

import (
	"errors"

	"github.com/merkurtran/goblog/app/models/user"
	"github.com/merkurtran/goblog/pkg/session"
	"gorm.io/gorm"
)

func _getUID() string {
	_uid := session.Get("uid")
	uid, ok := _uid.(string)
	if ok && len(uid) > 0 {
		return uid
	}
	return ""
}

func User() user.User {
	uid := _getUID()
	if len(uid) > 0 {
		_user, err := user.Get(uid)
		if err == nil {
			return _user
		}
	}
	return user.User{}
}

func Attempt(email string, password string) error {
	_user, err := user.GetByEmail(email)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("account notfound or password error")
		} else {
			return errors.New("Internal error, please try again after")
		}
	}
	if !_user.ComparePassword(password) {
		return errors.New("account notfound or password error")
	}

	session.Put("uid", _user.GetStringID())
	return nil
}

func Login(_user user.User) {
	session.Put("uid", _user.GetStringID())
}

func Logout() {
	session.Forget("uid")
}

func Check() bool {
	return len(_getUID()) > 0
}
