package sr_proto

const (
	C2PSInSR                = 111 //客户端发送给代理服务器的玩家进入学籍系统的消息
	S2CSRInterfaceMsgProto2 = 112 //服务器发送给客户端的学籍界面消息
	C2SGetAgainSRDataReq    = 113 //客户端发送给服务器的再次获取学籍界面信息的消息
	C2SWizardUpgradeReq     = 114 //巫师等级升级请求
	S2CWizardUpgradeMsg     = 115 //服务器发送给客户端的巫师等级升级信息
	C2SGetCollegeDetail     = 116 //客户端发送给服务器获取学院详情信息的请求消息
	S2CCollegeDetail        = 117 //服务器发送给客户端的学院详情信息
	C2SCollegeUpgradeReq    = 118 //学院升级请求
	S2CCollegeUpgradeMsg    = 119 //服务器发送给客户端的学院等级升级消息
)
