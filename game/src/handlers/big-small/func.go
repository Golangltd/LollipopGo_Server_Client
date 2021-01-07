package bigsmall

import (
	impl "LollipopGo/network"
	bailsman "LollipopGo2.8x/proto/big-small"
	twProto "github.com/Golangltd/Twlib/proto"
	"golang.org/x/net/websocket"
	"strconv"
)

// 进入游戏协议
func C2SEntryGameProto2Func(conn *websocket.Conn, ProtocolData map[string]interface{})  {

	strOpenId:= ProtocolData["OpenId"].(string)  // 玩家的校验码
	strTocken := ProtocolData["Tocken"].(string)  // 玩家的校验码
	iRoomType := int(ProtocolData["RoomType"].(float64)) // 房间类型：  1 2 3 4
	_=strTocken

	data := &bailsman.S2CEntryGame{
		Protocol:  twProto.GGameHallProto,
		Protocol2: bailsman.S2CEntryGameProto2,
	}

	GPlayerOpenID[strOpenId] = iRoomType
	// 加入房间逻辑
	userdata:=&bailsman.B_SUser{
		UUID:"1",
		RoleName:"demo1",
		RoleAvatar:1,
		RoleCoin:100000,
	}

	AddRoomFunc(userdata,iRoomType)
	// 验证房间的正确
	if iRoomType>0 && iRoomType < 5{
		data.Data = GRoom[iRoomType].Data
	}

	impl.PlayerSendToProxyServer(conn, data, strOpenId)
}

// 获取历史数据
func C2SGetHistoryDataProto2Func(conn *websocket.Conn, ProtocolData map[string]interface{})  {
	strOpenId:= ProtocolData["OpenId"].(string)
	iRoomType := int(ProtocolData["RoomType"].(float64))

	data := &bailsman.S2CGetHistoryData{
		Protocol:  twProto.GGameHallProto,
		Protocol2: bailsman.S2CGetHistoryDataProto2,
		HistoryData:GHistory[iRoomType],
	}
	impl.PlayerSendToProxyServer(conn, data, strOpenId)
}

// 玩家下注
func C2SUserChipInProto2Func(conn *websocket.Conn, ProtocolData map[string]interface{})  {

	strOpenId:= ProtocolData["OpenId"].(string)
	iRoomType := int(ProtocolData["RoomType"].(float64))
	iCoinType := int(ProtocolData["CoinType"].(float64))
	iChipInType := int(ProtocolData["ChipInType"].(float64))
	iChipInPos := int(ProtocolData["ChipInPos"].(float64))
	iChipInCoin := int(ProtocolData["ChipInCoin"].(float64))

	data:=bailsman.S2CUserChipIn{
		Protocol:  twProto.GGameHallProto,
		Protocol2: bailsman.S2CUserChipInProto2,
		IsMy:true, // 代表自己
	}

	if ProtocolData["ChipInType"] == nil {
		pasta :=GPlayer[strOpenId].MyAreaData[iChipInPos-1]
		if pasta == nil{
			damp := make(map[int]*SaveData)
			GPlayer[strOpenId+"|"+strconv.Itoa(iRoomType)].MyAreaData[iChipInPos-1].MapData = damp
			GPlayer[strOpenId+"|"+strconv.Itoa(iRoomType)].MyAreaData[iChipInPos-1].MapData[iCoinType].ICoinNum = iChipInCoin
		}else {
			GPlayer[strOpenId+"|"+strconv.Itoa(iRoomType)].MyAreaData[iChipInPos-1].MapData[iCoinType].ICoinNum += iChipInCoin
		}
		data.ChipInPos = iChipInPos
		data.ChipInCoin = iChipInCoin

		GPlayerBeforePos[strOpenId+"|"+strconv.Itoa(iRoomType)] = iChipInPos
		GPlayerBeforeOne[strOpenId+"|"+strconv.Itoa(iRoomType)] = iChipInCoin

	}else {

		if iChipInType == bailsman.Refund{
			data.ChipInPos = GPlayerBeforePos[strOpenId+"|"+strconv.Itoa(iRoomType)]
			data.ChipInCoin = -GPlayerBeforeOne[strOpenId+"|"+strconv.Itoa(iRoomType)]
		}else if iChipInType == bailsman.Cancel {
			GPlayer[strOpenId].MyAreaData = []*ChipInST{}
		}else if iChipInType == bailsman.Double {
             appdata :=GPlayer[strOpenId+"|"+strconv.Itoa(iRoomType)].MyAreaData
             for _,v:=range appdata {
             	for _,v1:=range  v.MapData{
             		v1.ICoinNum *=2
				}
			 }
		}else if iChipInType == bailsman.Before{
         	for k,v:=range GPlayerBefore[strOpenId+"|"+strconv.Itoa(iRoomType)].MyAreaData{
				data.BeforeCoin = append(data.BeforeCoin,v.MapData[k].ICoinNum)
			}
		}
	}

	impl.PlayerSendToProxyServer(conn, data, strOpenId)
}

// 获取概率
func C2SUserGetItemRateProto2Func(conn *websocket.Conn, ProtocolData map[string]interface{})  {

	strOpenId:= ProtocolData["OpenId"].(string)
	iRoomType := int(ProtocolData["RoomType"].(float64))
	data:=bailsman.S2CUserGetItemRate{
		Protocol:  twProto.GGameHallProto,
		Protocol2: bailsman.S2CUserGetItemRateProto2,
		ItemRate:GetChipInRate(GHistory[iRoomType]),
	}
	impl.PlayerSendToProxyServer(conn, data, strOpenId)
	return
}

// 获取游戏状态
// 切换场次的概念
func C2SUserGetGameStateProto2Func(conn *websocket.Conn, ProtocolData map[string]interface{})  {

	strOpenId:= ProtocolData["OpenId"].(string)
	iRoomType := int(ProtocolData["RoomType"].(float64))
	_,Data:=GetRoomState(iRoomType)

	data:=bailsman.S2CUserGetGameState{
		Protocol:  twProto.GGameHallProto,
		Protocol2: bailsman.S2CUserGetGameStateProto2,
		GameState: GRoom[iRoomType].Data.GameState,
		Data:&Data,
	}

	impl.PlayerSendToProxyServer(conn, data, strOpenId)
}