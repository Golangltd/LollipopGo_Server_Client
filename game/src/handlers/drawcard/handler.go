package drawcard

import (
	"LollipopGo/log"
	impl "LollipopGo/network"
	. "LollipopGo2.8x/conf/error"
	"LollipopGo2.8x/cxt"
	"LollipopGo2.8x/handlers/card"
	. "LollipopGo2.8x/handlers/modules"
	"LollipopGo2.8x/models"
	. "LollipopGo2.8x/msg/drawcard"
	"LollipopGo2.8x/proto/comm_proto"
	. "LollipopGo2.8x/proto/dc_proto"
	"golang.org/x/net/websocket"
)

func InitDCHandlers() {
	HM.AddHandler(C2SUserIntoDCReq, DCHandleCallBack)
	HM.AddHandler(C2STenDrawReq, DrawHandle)
}

//用户进入抽奖界面
func DCHandleCallBack(conn *websocket.Conn, ProtocolData map[string]interface{}) {
	var (
		value interface{}
		err   error
	)
	openID := ProtocolData["OpenId"].(string) //获取用户openID
	key := openID + cxt.UserKey               //拼接userKey,在内存中获取用户数据
	if value, err = cxt.M.Get(key); err != nil {
		log.Error("dont't get user data from M by current key[%v] !,err is [%v]", key, err)
		return
	}
	game := value.(*models.Game)
	game.DrawCardModel = models.NewDrawCardModel(openID, game)
	game.DrawCardModel.GetDCInterfaceData()
	impl.PlayerSendToProxyServer(conn, game.DrawCardModel.DCInterface, game.DrawCardModel.OpenId)
}

//抽奖处理
func DrawHandle(conn *websocket.Conn, ProtocolData map[string]interface{}) {
	var (
		value interface{}
		err   error
	)
	data := NewDrawReqData(ProtocolData)
	key := data.OpenId + cxt.UserKey //拼接userKey,在内存中获取用户数据
	if value, err = cxt.M.Get(key); err != nil {
		log.Error("dont't get user data from M by current key[%v] !,err is [%v]", key, err)
		return
	}
	game := value.(*models.Game)
	game.DrawCardModel.GetDrawData(data.DrawType)
	if game.DrawCardModel.DrawData.Code <= CorrectCode { //code为0,说明已经抽奖,进行用户卡牌数据的更新和装备背包的更新
		comm_proto.UpdateRoleCoin(conn, game.DrawCardModel.OpenId, -game.DrawCardModel.SpendAmount, game.UserInfo)
		card.AddCardAndNotify(conn, game.DrawCardModel.OpenId, game, game.DrawCardModel.DrawData.CardList)
	}
	impl.PlayerSendToProxyServer(conn, game.DrawCardModel.DrawData, game.DrawCardModel.OpenId)
}
