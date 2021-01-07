package main

import (
	impl "LollipopGo/network"
	"LollipopGo/util"
	. "LollipopGo2.8x/cxt"
	gamedb "LollipopGo2.8x/data"
	"LollipopGo2.8x/models"
	"LollipopGo2.8x/proto/comm_proto"
	ProtoGame "LollipopGo2.8x/proto/tw_proto"
	"LollipopGo2.8x/tables"
	TwChapter "github.com/Golangltd/Twlib/Chapter"
	twLibItem "github.com/Golangltd/Twlib/item"
	twProto "github.com/Golangltd/Twlib/proto"
	twServer "github.com/Golangltd/Twlib/server"
	twUser "github.com/Golangltd/Twlib/user"
	"github.com/golang/glog"
	"golang.org/x/net/websocket"
	"strconv"
)

// 获取玩家数据
func UserStGet(conn *websocket.Conn, ProtocolData map[string]interface{}) {
	strOpenId := ProtocolData["OpenId"].(string)
	val, _ := M.Get(strOpenId + UserKey)
	senate := comm_proto.GS2S_GetUserSt{
		Protocol:  twProto.GGameHallProto,
		Protocol2: comm_proto.GS2S_GetUserStProto2,
		OpenId:    strOpenId,
		UserInfo:  val.(*models.Game).UserInfo,
	}
	impl.PlayerSendMessageOfProxy(conn, senate, util.MD5_LollipopGO(strconv.Itoa(twServer.CenterServerId)))
}

func UserStSet(conn *websocket.Conn, ProtocolData map[string]interface{}) {

	strOpenId := ProtocolData["OpenId"].(string)
	strUserInfo := ProtocolData["UserInfo"].(*twUser.UserSt)
	senate := comm_proto.GS2S_GetUserSt{
		Protocol:  twProto.GGameHallProto,
		Protocol2: comm_proto.GS2S_UpdateUserStProto2,
		OpenId:    strOpenId,
		UserInfo:  strUserInfo,
	}
	impl.PlayerSendToProxyServer(conn, senate, strOpenId)
	_, _ = M.Put(strOpenId+UserKey, strUserInfo)
}

// 玩家点击开始战斗
func UserPlayBattle(conn *websocket.Conn, ProtocolData map[string]interface{}) {
	strOpenId := ProtocolData["OpenId"].(string)
	var rite []*twUser.ItemSt
	senate := ProtoGame.C2GSUserPlayRet{
		Protocol:  twProto.GGameHallProto,
		Protocol2: ProtoGame.C2GSUserPlayRetProto2,
		IsVictory: true,
		Rewards:   rite,
		Exp:       0,
	}
	val1, _ := M.Get(strOpenId + UserKey)
	val, _ := M.Get(strconv.Itoa(int(val1.(*models.Game).AccountId)) + "|User_R")
	dataAward := tables.RoundTables[val.(*models.Game).RoundInfo.Sid].RewardInfo
	getter := tables.GetChapterFor3Row(dataAward)
	var daytime = new(twUser.ItemSt)
	for k, v := range getter {
		var daytime1 = new(twUser.ItemData)
		daytime1.ItemId = k
		daytime1.ItemNum = int64(v)
		daytime.ItemData = append(daytime.ItemData, daytime1)
		if k == twUser.ICardExp {
			senate.Exp = v
			val1.(*models.Game).UserInfo.RoleExp += senate.Exp
			// val1.(*models.Game).SRModel.UpGrade()
		}
	}
	rite = append(rite, daytime)
	impl.PlayerSendToProxyServer(conn, senate, strOpenId)
	val1.(*models.Game).UserInfo.RoleExp += senate.Exp
	if senate.IsVictory {
		// 通关奖励
		SendDataInfo(val.(*models.Game), ConnXZ, strOpenId)
		if true {
			valol, _ := M.Get(strconv.Itoa(val.(*models.Game).RoundInfo.ChapterId) + "|User_ROL")
			if valol != nil {
				data := valol.([]*TwChapter.ChapterUserSt)
				for k, v := range data {
					if v.RoleName == val1.(*models.Game).UserInfo.RoleName && len(v.RoleName) != 0 {
						data = append(data[:k], data[k+1:]...)
					}
				}
			}
		}

		val.(*models.Game).RoundInfo.Round++
		// val.(*models.Game).RoundInfo.Sid++ // 切换子章节时候 增加
		//fmt.Println("==================================val.(*models.Game).RoundInfo.Sid", val.(*models.Game).RoundInfo.Sid)
		if val.(*models.Game).RoundInfo.Round-2 != 0 {
			//fmt.Println("==================================ChapterId")
			//fmt.Println("==================================val.(*models.Game).RoundInfo.Round-1", val.(*models.Game).RoundInfo.Round-1)
			//fmt.Println("==================================tables.GetRoundNumFromChapter(val.(*models.Game).RoundInfo.Sid)", tables.GetRoundNumFromChapter(val.(*models.Game).RoundInfo.Sid))
			if (val.(*models.Game).RoundInfo.Round-1)%tables.GetRoundNumFromChapter(val.(*models.Game).RoundInfo.Sid) == 0 { // 累加++
				val.(*models.Game).RoundInfo.ChapterId2++
				val.(*models.Game).RoundInfo.Sid++
				//fmt.Println("==================================ChapterId2")
				if tables.GetChapter2NumFromChapter(val.(*models.Game).RoundInfo.Sid) == 1 {
					val.(*models.Game).RoundInfo.ChapterId++
					val.(*models.Game).RoundInfo.ChapterId2 = 1
					//fmt.Println("==================================ChapterId3")
				}
			}
		}

		if true {
			valol, _ := M.Get(strconv.Itoa(val.(*models.Game).RoundInfo.ChapterId) + "|User_ROL")
			data := valol.([]*TwChapter.ChapterUserSt)
			if valol != nil && len(data) < 10 {
				oluer := new(TwChapter.ChapterUserSt)
				oluer.RoleName = val1.(*models.Game).UserInfo.RoleName
				oluer.RoleAvatar = val1.(*models.Game).UserInfo.RoleAvatar
				oluer.RoleLev = val1.(*models.Game).UserInfo.RoleLev
				if oluer != nil {
					data = append(data, oluer)
				}
			}
		}

		savedb := new(gamedb.SChapter)
		savedb.RoleUID = val1.(*models.Game).UserInfo.RoleUid
		savedb.Account = val1.(*models.Game).AccountId
		savedb.ChapterInfo = new(twUser.ChapterInfo)
		savedb.ChapterInfo.ChapterId = val.(*models.Game).RoundInfo.ChapterId
		savedb.ChapterInfo.ChapterId2 = val.(*models.Game).RoundInfo.ChapterId2
		savedb.ChapterInfo.RoundId = val.(*models.Game).RoundInfo.Round
		gamedb.SaveChapterInfo(savedb)
		return
	}
}

// 更新通关奖励
func SendDataInfo(game *models.Game, conn *websocket.Conn, stropenid string) {
	data := tables.GetRoundFor3Row(tables.RoundTables[game.RoundInfo.Round].RewardInfo)
	glog.Info("SendDataInfo-----", data)
	items := make([]*twLibItem.ItemData, 0)
	// TODO:删除  装备强化道具
	items = append(items, &twLibItem.ItemData{
		Type: twUser.ItemCommon,
		ID:   48,
		Num:  10,
	})
	for _, v := range data {

		if v.ItemType == twUser.ItemCommon { //  道具
			items = append(items, &twLibItem.ItemData{
				Type: twUser.ItemCommon,
				ID:   v.ItemId,
				Num:  v.ItemNum,
			})
		} else if v.ItemType == twUser.ItemEquipment { // 装备

			data := new(twUser.EquipSt)
			data.Num = v.ItemNum
			data.ConfID = v.ItemId
			models.InitEquipPower(data)
			quips := []*twUser.EquipSt{data}
			comm_proto.UpdateEquip(conn, stropenid, quips)

		} else if v.ItemType == twUser.ItemCard { // 卡牌

		}
	}
	err := comm_proto.UpdateRoleItem(game, conn, stropenid, false, items)
	if err != nil {
		glog.Errorf("SendDataInfo err:%s\n", err.Error())
	}
}
