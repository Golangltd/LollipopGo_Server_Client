package msg

import (
	twlib_user "github.com/Golangltd/Twlib/user"
)

// 主协议 == 1
const (
	INITMsg = iota
	C2SUserLoginProto // C2SUserLoginProto ==  1 玩家登录协议
	S2CUserLoginProto // S2CUserLoginProto ==  2 服务器返回数据

	C2SUserGetServerListProto // C2SUserGetServerListProto == 3 获取登录列表
	S2CUserGetServerListProto // S2CUserGetServerListProto == 4 获取列表数据  目前设计的 30个数据一页
)

//---------------------------------------------------------------------------------------
// C2SUserGetServerListProto == 3 获取登录列表
type C2SUserGetServerList struct {
	PageNum int   // 第几页
}

// S2CUserGetServerListProto == 4 获取列表数据
// 30个数据一列
type S2CUserGetServerList struct {
	SeverList []*twlib_user.ServerList
}

//---------------------------------------------------------------------------------------
// C2SUserLoginProto == 1  客户端请求服务器的数据
// url := http://192.168.2.115:8867/BaBaLiuLiu
type C2SUserLogin struct {
	AccountName  string
	AccountPw string
	DeviceId string     // 游客登录，重新安装后就需要重置
}

// S2CUserLoginProto ==  2  服务器返回客户端登录数据
type S2CUserLogin struct {
	Protocol  int
	Protocol2 int
	//ProxyUrl string
	UserData *twlib_user.UserSt
/*	GameList  []ST.GameList  // 游戏列表*/
}
// ----------------------------------------------------------------------------------------