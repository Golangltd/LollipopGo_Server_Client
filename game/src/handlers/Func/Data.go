package Func


// Function_protoConfMap  -- 配置文件

// 功能开启类型
const (
	FuncInit         = iota
	RoleLev          = 1  // 角色等级
	CardLev          = 2  // 卡牌等级
	VIPLev           = 3  // VIP等级
	LoginDay         = 4  // 登录天数
	StartServerDay   = 5  // 开服天数
	CustomsCount     = 6  // 关卡通关
	FinishMasterTask = 7  // 完成主线任务
	CustomsPaTa      = 8  // 通关爬塔
	UnionLev         = 9  // 工会等级
	CashNum          = 10 // 现金数量，达到
	Power            = 11 // 战斗力数值，达到
	StadiumRank      = 12 // 竞技场排名
	SumCashDay       = 13 // 累计充值天数
	SignInDay        = 14 // 累计签到天数
	NeedItemId       = 15 // 需要道具的ID
	NeedCardId       = 16 // 需要卡牌的ID
)

// 功能开启结构
type FuncSt struct {

}