package main

import (
	impl "LollipopGo/network"
	. "LollipopGo2.8x/cxt"
	gameDB "LollipopGo2.8x/data"
	"LollipopGo2.8x/models"
	"LollipopGo2.8x/proto/comm_proto"
	"LollipopGo2.8x/temp"
	twProto "github.com/Golangltd/Twlib/proto"
	twUser "github.com/Golangltd/Twlib/user"
	"golang.org/x/net/websocket"
	"strconv"
)

// 中心服道具
func UpdateItemData(conn *websocket.Conn, ProtocolData map[string]interface{}) {

	strOpenId := ProtocolData["OpenId"].(string)
	iType := int(ProtocolData["Type"].(float64))
	ItemInfo := ProtocolData["ItemInfo"]
	CardInfo := ProtocolData["CardInfo"]
	EquipInfo := ProtocolData["EquipInfo"]
	val, _ := M.Get(strOpenId + UserKey)
	if ItemInfo != nil { // 道具
		var Item2 []*twUser.ItemData
		for _, v := range ProtocolData["ItemInfo"].([]interface{}) {
			itemid := v.(map[string]interface{})["ItemId"].(float64)
			itemnum := v.(map[string]interface{})["ItemNum"].(float64)
			itemtype := v.(map[string]interface{})["ItemType"].(float64)
			//uid := 11111//twlib_uniqueid.NextID()
			// 更新数据库，返回道具唯一
			gameDB.SaveItem(val.(*models.Game).UserInfo.RoleUid, val.(*models.Game).AccountId, &twUser.ItemData{

				ItemId:   int(itemid),
				ItemType: int(itemtype),
				ItemNum:  int64(itemnum),
			})
			Item2 = append(Item2, &twUser.ItemData{

				ItemId:   int(itemid),
				ItemType: int(itemtype),
				ItemNum:  int64(itemnum),
			})
		}
		temp.GMapBag[val.(*models.Game).AccountId] = Item2
	}

	if CardInfo != nil { //卡牌

	}

	if EquipInfo != nil { // 装备

	}

	if iType == 1 { // 增加
		/*CreateItem(val.(*modules.Game).AccountId)*/
	} else if iType == 2 { // 删除
	}

	senate := comm_proto.GS2C_UpDataItemData{
		Protocol:  twProto.GGameHallProto,
		Protocol2: comm_proto.GS2C_UpDataItemDataProto2,
	}
	impl.PlayerSendToProxyServer(conn, senate, strOpenId)
}

func GetRoundData(conn *websocket.Conn, ProtocolData map[string]interface{}) {
	strOpenId := ProtocolData["OpenId"].(string)
	valc, _ := M.Get(strOpenId + UserKey)
	val, _ := M.Get(strconv.Itoa(int(valc.(*models.Game).AccountId)) + "|User_R")
	senate := comm_proto.GS2C_GetFunctionData{
		Protocol:  twProto.GGameHallProto,
		Protocol2: comm_proto.GS2C_GetFunctionDataProto2,
		Data: strconv.Itoa(val.(*models.Game).RoundInfo.ChapterId) + "," + strconv.Itoa(val.(*models.Game).RoundInfo.ChapterId2) + "," +
			strconv.Itoa(val.(*models.Game).RoundInfo.Round), //"1,2,10", // 注：大章节，小章杰，关卡数
	}
	impl.PlayerSendToProxyServer(conn, senate, strOpenId)
}
