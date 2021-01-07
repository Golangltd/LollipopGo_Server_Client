package common

import (
	impl "LollipopGo/network"
	"LollipopGo2.8x/handlers/card"
	util_handlers "LollipopGo2.8x/handlers/util"
	"LollipopGo2.8x/models"
	"LollipopGo2.8x/msg/drawcard"
	"LollipopGo2.8x/proto/comm_proto"
	twlib_common "github.com/Golangltd/Twlib/common"
	twLibItem "github.com/Golangltd/Twlib/item"
	twlib_proto "github.com/Golangltd/Twlib/proto"
	twlib_user "github.com/Golangltd/Twlib/user"
	"github.com/golang/glog"
	"golang.org/x/net/websocket"
)

/**
发送错误提示
参数:连接,玩家OpenID,错误码,默认错误信息,提示类型
*/
func SendError(conn *websocket.Conn, openID string, errorCode string, errorDesc string, notifyType int) {
	msg := &twlib_common.ErrorMessage{
		Protocol:  twlib_proto.GErrorProto,
		Protocol2: twlib_common.ErrorMessageProto2,
		Code:      errorCode,
		Desc:      errorDesc,
		Type:      notifyType,
	}
	impl.PlayerSendMessageOfProxy(conn, msg, openID)
}

// 更新道具,装备，卡牌等
func UpdateItemOfServer(itemdata *twlib_user.ItemData, strOpenId string, conn *websocket.Conn, icardQuality int) {
	game, _, _ := util_handlers.GetGameAndUser(strOpenId)
	items := make([]*twLibItem.ItemData, 0)
	if itemdata.ItemType == twlib_user.ItemCommon { //  道具

		if itemdata.ItemId <= twlib_user.ICardBreakCoin {
			if itemdata.ItemId == twlib_user.ICardExp {
				game.UserInfo.RoleExp += int(itemdata.ItemNum)
			} else if itemdata.ItemId == twlib_user.ICoin {
				game.UserInfo.Coin += itemdata.ItemNum
			}
			return
		}
		items = append(items, &twLibItem.ItemData{
			Type: twlib_user.ItemCommon,
			ID:   itemdata.ItemId,
			Num:  int(itemdata.ItemNum),
		})
		err := comm_proto.UpdateRoleItem(game, conn, strOpenId, false, items)
		if err != nil {
			glog.Errorf("SendDataInfo err:%s\n", err.Error())
		}
	} else if itemdata.ItemType == twlib_user.ItemEquipment { // 装备
		data := new(twlib_user.EquipSt)
		data.Num = int(itemdata.ItemNum)
		data.ConfID = itemdata.ItemId
		models.InitEquipPower(data)
		equiplist := []*twlib_user.EquipSt{data}
		comm_proto.UpdateEquip(conn, strOpenId, equiplist)
	} else if itemdata.ItemType == twlib_user.ItemCard { // 卡牌
		cardList := []*drawcard.CardData{}
		carddata := new(drawcard.CardData)
		carddata.ItemType = twlib_user.ItemCard
		carddata.CardID = itemdata.ItemId
		carddata.Quality = icardQuality
		cardList = append(cardList, carddata)
		card.AddCardAndNotify(conn, strOpenId, game, cardList)
	}
}
