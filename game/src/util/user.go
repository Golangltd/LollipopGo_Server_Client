package game_util

var GUser map[int64]string

func init()  {
	GUser = make(map[int64]string)
}

func SetGUserInfo(strOpenId string,account int64)  {
	GUser[account] = strOpenId
}

