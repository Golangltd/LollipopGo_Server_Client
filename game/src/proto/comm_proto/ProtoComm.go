package comm_proto

import (
	twlibproto "github.com/Golangltd/Twlib/proto"
	twlib_user "github.com/Golangltd/Twlib/user"
)

// 通用更新消息,主协议：6
const (
	GS2C_UpdateCoinProto2    = 10001 //  更新金币道具等协议
	GS2C_UpdateDiamondProto2 = 10002 //  更新砖石协议

	S2GS_GetUserStProto2 = 10003 //  服务器间获取玩家数据，仅限服务器内部访问
	GS2S_GetUserStProto2 = 10004 //  返回玩家数据

	S2GS_UpdateUserStProto2 = 10005 //  更新玩家数据，不同服务器间玩家数据更新
	GS2S_UpdateUserStProto2 = 10006 //  返回玩家数据

	C2GS_GetFunctionDataProto2 = 10007 //  获取玩家数据
	GS2C_GetFunctionDataProto2 = 10008 //  返回数据

	C2GS_UpDataItemDataProto2 = 10009 // 更新道具             --- 中心服 发给游戏主逻辑服务器
	GS2C_UpDataItemDataProto2 = 10010 // 返回数据更新成功

	C2GS_UpDataItem41DataProto2 = 10011 // 更新道具 功能ID=41
	GS2C_UpDataItem41DataProto2 = 10012 // 返回数据更新成功

	GS2C_UpDateUserExpDataProto2 = 10013 // 更新经验

	GS2C_UpdateCardProto2 = 10014 // 更新卡牌
)

//-------------------------------------------------------------------------------------------
// GS2C_UpDateUserExpDataProto2
type GS2C_UpDateUserExpData struct {
	Protocol  int //  主协议：6
	Protocol2 int
	UserExp   int // 更新玩家经验
}

//-------------------------------------------------------------------------------------------
// GS2C_UpDataItem41DataProto2
type GS2C_UpDataItem41Data struct {
	Protocol  int //  主协议：6
	Protocol2 int
	ItemList  []*twlib_user.ItemData
}

//-------------------------------------------------------------------------------------------
// 更新道具
type C2GS_UpDataItemData struct {
	Protocol  int //  主协议：6
	Protocol2 int
	OpenId    string // 更新中心服的道具到玩家身上
	Type      int    // 更新类型，1：增加，2：删除
	ItemInfo  []*twlib_user.ItemSt
	CardInfo  []*twlib_user.CardInfo
	EquipInfo []*twlib_user.EquipData
}

// GS2C_UpDataItemDataProto2
type GS2C_UpDataItemData struct {
	Protocol  int //  主协议：6
	Protocol2 int
}

//-------------------------------------------------------------------------------------------
type GS2C_UpdateCoinOrOtherItem struct {
	Protocol  int //  主协议：6
	Protocol2 int
	ItemType  int   // 对应道具类型表
	ItemId    int64 // 道具Id
	ItemNum   int64 // 道具数量
}

//-------------------------------------------------------------------------------------------
type C2GS_GetFunctionData struct {
	Protocol  int //  主协议：6
	Protocol2 int
	OpenId    string
	IType     int
}

type GS2C_GetFunctionData struct {
	Protocol  int //  主协议：6
	Protocol2 int
	Data      interface{}
}

// 获取类型
const (
	GetFuncType = iota
	GetRound
)

//-------------------------------------------------------------------------------------------
//  更新玩家数据，不同服务器间玩家数据更新
type S2GS_UpdateUserSts struct {
	Protocol  int //  主协议：6
	Protocol2 int
	UserInfo  *twlib_user.UserSt
}

// 回则表示成功
type GS2S_UpdateUserSt struct {
	Protocol  int //  主协议：6
	Protocol2 int
	UserInfo  *twlib_user.UserSt
}

//-------------------------------------------------------------------------------------------
// 服务器间获取玩家数据，仅限服务器内部访问
type S2GS_GetUserSt struct {
	Protocol  int //  主协议：6
	Protocol2 int
	OpenId    string // 玩家唯一Id
}

type GS2S_GetUserSt struct {
	Protocol  int //  主协议：6
	Protocol2 int
	OpenId    string // 玩家唯一Id
	UserInfo  *twlib_user.UserSt
}

//-------------------------------------------------------------------------------------------
//  更新砖石协议
type GS2C_UpdateDiamond struct {
	Protocol   int
	Protocol2  int
	DiamondNum int // 砖石数量
}

//-------------------------------------------------------------------------------------------
//  更新金币协议
type GS2C_UpdateCoin struct {
	Protocol  int
	Protocol2 int
	CoinNum   int // 金币数量
}

//-------------------------------------------------------------------------------------------
// 卡牌更新
type GS2CUpdateCardMsg struct {
	Protocol  int
	Protocol2 int
	OpenId    string                 // 玩家唯一ID
	Updates   []*twlib_user.CardInfo // 更新的卡牌
	Adds      []*twlib_user.CardInfo // 新增的卡牌
	DelUIDs   []uint64               // 删除的卡牌UID
}

func NewGS2CUpdateCardMsg() *GS2CUpdateCardMsg {
	m := &GS2CUpdateCardMsg{
		Protocol:  twlibproto.GGameHallProto,
		Protocol2: GS2C_UpdateCardProto2,
	}
	return m
}
