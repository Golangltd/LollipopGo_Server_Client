package Round

import twlib_user "github.com/Golangltd/Twlib/user"

var UserOl map[int64]*twlib_user.UserSt

func init() {
	UserOl = make(map[int64]*twlib_user.UserSt)
}

func SetUserOlInfo(userdata *twlib_user.UserSt) {
	if len(UserOl) < 10 {
		UserOl[userdata.RoleUid] = userdata
	}
}

func GetUserOlInfo(iChapterId int) map[int64]*twlib_user.UserSt {
	return UserOl
}
