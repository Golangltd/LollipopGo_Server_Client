package msg

import twlib_user "github.com/Golangltd/Twlib/user"

type LoginData struct {
	Token string
	Url string
	ServerList []*twlib_user.ServerList
	AesKey string
}
