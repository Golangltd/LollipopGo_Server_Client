package gm_proto

//  GGameGMProto == 8       游戏GM管理系统

const (
	GMINIT = iota
	C2GS_GMProto2   // C2GS_GMProto2 == 1  子协议 客户端GM命令
)

/* 类型说明
	OpType{
		1:增/减金币
		2:增/减砖石
		3:增/减道具
		4:增/减卡牌
		5:增/减装备
	}
*/

// 对应的结构体信息
type C2GS_GM struct {
	Protocol  int
	Protocol2 int
	OpenId    string // 玩家唯一Id
	OpType    int    // 操作类型
	ItemId    int    // 道具Id,策划对应的配置表
	ItemNum   int    // 变化数量，正式为增加，负数减少
}