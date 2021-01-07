package OffLine

import (
	ProtoGame "LollipopGo2.8x/proto/tw_proto"
	game_util "LollipopGo2.8x/util"
	twlib_proto "github.com/Golangltd/Twlib/proto"
	twlib_rewards "github.com/Golangltd/Twlib/rewards"
	"time"
)

func InitOffLine() {
	go Timer()
}

// 获取挂机的时间
// 玩家注册的时间 和现在的时间：96400 （24小时）取余，然后处理数据

func Timer() {
	Goldmines := time.NewTicker(time.Millisecond * 1000)
	icount := 0
	for {
		select {
		case <-Goldmines.C:
			icount++
			for _, v := range game_util.GUser {
				SendOffLineInfoAward(v, icount)
			}
		}
	}
}

// 针对的是玩家的数据结构
var Rewatds = new(twlib_rewards.RewardSt)
var Rewatds1 = new(twlib_rewards.RewardSt)
var Rewatds2 = new(twlib_rewards.RewardSt)
var Rewatds3 = new(twlib_rewards.RewardSt)
var Time int = 0

func SendOffLineInfoAward(stropenid string, icount int) {
	// conn:=game_util.GetConnALl()
	data := &ProtoGame.GS2CUserOffLineBattle{
		Protocol:    twlib_proto.GGameHallProto,
		Protocol2:   ProtoGame.GS2CUserOffLineBattleProto2,
		OffLineTime: uint64(icount),
		Rewards:     nil,
	}
	Time++
	// 收益--
	// 1. 金币收益
	Rewatds.ItemId = 1
	Rewatds.ItemType = 1
	JinBiNum += JinBi / 60
	Rewatds.ItemNum = JinBiNum
	data.Rewards = append(data.Rewards, Rewatds)
	// 2. 砖石收益
	Rewatds1.ItemId = 2
	Rewatds1.ItemType = 1
	YinBiNum += YinBi / 60
	Rewatds1.ItemNum = YinBiNum
	data.Rewards = append(data.Rewards, Rewatds1)
	// 3. 战队经验收益
	Rewatds2.ItemId = 3
	Rewatds2.ItemType = 1
	ZhanDuiNum += ZhanDui / 60
	Rewatds2.ItemNum = ZhanDuiNum
	data.Rewards = append(data.Rewards, Rewatds2)
	// 4. 开牌经验收益
	Rewatds3.ItemId = 4
	Rewatds3.ItemType = 1
	CardNum += Card / 60
	Rewatds3.ItemNum = CardNum
	Rewatds3.ItemNum = Card / 60
	data.Rewards = append(data.Rewards, Rewatds3)
	/*
		if conn == nil{
			return
		}
		impl.PlayerSendToProxyServer(conn,data,stropenid)
	*/
}
