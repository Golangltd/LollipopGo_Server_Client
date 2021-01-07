package user

type Userrpc struct {
}

type User struct {
	ID         int     //用户ID
	Openid     string  //用户openID
	Mobile     string  //手机号
	Wechat     string  //微信号
	Nickname   string  //昵称
	Avatar     string  //头像
	IsGuest    int8    //是否游客
	DeviceID   string  //设备ID
	Password   string  //密码
	IdCcard    string  //身份证号码
	Wealth     float64 //金币
	MobileBind int8    //是否绑定手机号 0未绑定 1绑定
	WechatBind int8    //是否绑定微信 0未绑定 1绑定
	Status     int8    //状态 1封号 2正常
}

// 数据操作