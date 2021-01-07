package chat_proto

/*
   聊天系统
   1. 玩家上线后可以接受到消息，除了私聊，世界频道等全部是打开后收到消息
*/

const (
	ChatInit = iota

	C2GSOpenChatSysProto2 = 251 // 玩家打开界面
	GS2COpenChatSysProto2 = 252

	C2GSSendChatSysProto2 = 253 // 发送数据
)

// 聊天的玩家的结构数据信息
type ChatUserSt struct {
	RoleName    string      //角色名字
	RoleAvatar  int         //角色头像
	RoleLev     int         //角色等级
	RoleSex     int         //角色性别
	//TODO:根据调试再增加
}

// 聊天数据结构
type ChatSt struct {
	UserInfo *ChatUserSt
	Content  string
}

//--------------------------------------------------------------------------------------------------------
//  C2GSSendChatSysProto2 = 253
type C2GSSendChatSys struct {
	Protocol  int
	Protocol2 int
	OpenId    string
	ChatType  int      // twlib_chat
	Content   string
}

//--------------------------------------------------------------------------------------------------------
// C2GSOpenChatSysProto2 = 251
type C2GSOpenChatSys struct {
	Protocol  int
	Protocol2 int
	OpenId    string
	ChatType  int      // twlib_chat
}

// GS2COpenChatSysProto2 = 252
type GS2COpenChatSys struct {
	Protocol  int
	Protocol2 int
	ChatInfo  []*ChatSt
}

//--------------------------------------------------------------------------------------------------------
