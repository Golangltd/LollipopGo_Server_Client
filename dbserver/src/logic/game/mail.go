package game


// 获取邮件列表
func (m *GameRPC)GetUserMailList(openid string,reply *interface{})  error{
	return nil
}

// 删除邮件
func (m *GameRPC)DelUserMail(mailid int64,reply *interface{}) error {
	return nil
}

// 领取附件
// 获取结构信息 -- 远程服务器传递过来的数据是结构
/*func (this *GameRPC)GetUserAttachment(openid string,mailid int64,reply *interface{}) error {
	return nil
}*/

// 一键领取所有附件
func (m *GameRPC)GetAllUserAttachment(openid string,reply *interface{}) error {
	return nil
}