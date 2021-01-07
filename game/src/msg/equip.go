package gamemsg

import (
	ProtoGame "LollipopGo2.8x/proto/tw_proto"
	twlib_proto "github.com/Golangltd/Twlib/proto"
	twlib_user "github.com/Golangltd/Twlib/user"
)

// 道具数据
type ItemDataStr struct {
	ItemUid  string // 唯一ID
	ItemId   int
	ItemType int
	ItemNum  int64 // 道具的数量
}

// 客户端请求 穿戴装备/一键穿戴
type C2GSEquipWearMsg struct {
	Protocol  int
	Protocol2 int
	OpenId    string  // 玩家唯一ID
	CardUID   int64   // 卡牌唯一ID
	ItemUIDs  []int64 // 装备背包中的装备唯一ID
}
type C2GSEquipWearStrMsg struct {
	Protocol  int
	Protocol2 int
	OpenId    string   // 玩家唯一ID
	CardUID   int64    // 卡牌唯一ID
	ItemUIDs  []string // 装备背包中的装备唯一ID
}

// 客户端请求 脱下装备/一键脱下
type C2GSEquipTakeOffMsg struct {
	Protocol  int
	Protocol2 int
	OpenId    string // 玩家唯一ID
	CardUID   int64  // 卡牌唯一ID
	Positions []int  // 装备部位
}

// 客户端请求 装备替换
type C2GSEquipReplaceMsg struct {
	Protocol    int
	Protocol2   int
	OpenId      string // 玩家唯一ID
	CardUID     int64  // 当前操作的卡牌唯一ID
	Position    int    // 当前操作的装备部位
	SrcCardUID  int64  // 来源卡牌ID
	SrcPosition int    // 来源部位
}

// 客户端请求 装备强化
type C2GSEquipStrengthenMsg struct {
	Protocol      int
	Protocol2     int
	OpenId        string                 // 玩家唯一ID
	CardUID       int64                  // 卡牌唯一ID 为0时，表示装备背包中的装备
	EquipUID      int64                  // 装备唯一ID
	CostMaterials []*twlib_user.ItemData // 使用的材料
	CostEquips    []*twlib_user.ItemData // 使用的装备
}
type C2GSEquipStrengthenStrMsg struct {
	Protocol      int
	Protocol2     int
	OpenId        string                 // 玩家唯一ID
	CardUID       int64                  // 卡牌唯一ID 为0时，表示装备背包中的装备
	EquipUID      string                 // 装备唯一ID
	CostMaterials []*twlib_user.ItemData // 使用的材料
	CostEquips    []*ItemDataStr         // 使用的装备
}

type CardEquipSt struct {
	CardUID int64                 // 卡牌唯一ID
	Equips  []*twlib_user.EquipSt // 装备
}

// 服务器返回 卡牌装备更新
type GS2CCardEquipUpdateMsg struct {
	Protocol    int
	Protocol2   int
	CardsEquips []*CardEquipSt
}

func NewGS2CCardEquipUpdate() *GS2CCardEquipUpdateMsg {
	m := &GS2CCardEquipUpdateMsg{
		Protocol:  twlib_proto.GGameHallProto,
		Protocol2: ProtoGame.GS2CCardEquipUpdate,
	}
	return m
}

// 服务器返回 卡牌装备脱下
type GS2CCardEquipDelMsg struct {
	Protocol  int
	Protocol2 int
	CardUID   int64 // 卡牌ID
	Positions []int // 移除的部位
}

func NewGS2CCardEquipDelMsg() *GS2CCardEquipDelMsg {
	m := &GS2CCardEquipDelMsg{
		Protocol:  twlib_proto.GGameHallProto,
		Protocol2: ProtoGame.GS2CCardEquipDel,
	}
	return m
}
