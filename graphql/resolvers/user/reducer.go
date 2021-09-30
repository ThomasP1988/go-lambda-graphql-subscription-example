package user

import (
	"shared/repositories/user"
)

type UserAPI = user.User

func Reducer(usr *user.User) UserAPI {
	return UserAPI(*usr)
}
