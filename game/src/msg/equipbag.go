package gamemsg

import (
	ProtoGame "LollipopGo2.8x/proto/tw_proto"
	twlib_proto "github.com/Golangltd/Twlib/proto"
	twlib_user "github.com/Golangltd/Twlib/user"
)

// 服务器推送 背包装备更新
type GS2CBagEquipUpdateMsg struct {
	Protocol  int
	Protocol2 int
	Updates   []*twlib_user.EquipSt // 更新的装备
	Adds      []*twlib_user.EquipSt // 新增装备
	DelUIDs   []int64               // 删除装备的UID
}

func NewGS2CBagEquipUpdateMsg() *GS2CBagEquipUpdateMsg {
	m := &GS2CBagEquipUpdateMsg{
		Protocol:  twlib_proto.GGameHallProto,
		Protocol2: ProtoGame.GS2CEquipBagUpdate,
	}
	return m
}
