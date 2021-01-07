package chat

import (
	impl "LollipopGo/network"
	"LollipopGo2.8x/proto/chat_proto"
	twProto "github.com/Golangltd/Twlib/proto"
	"golang.org/x/net/websocket"
)

// 获取聊天数据 -- 更新操作
func GetChatInfo(conn *websocket.Conn, ProtocolData map[string]interface{}) {
	strOpenId := ProtocolData["OpenId"].(string)

	senate := chat_proto.GS2COpenChatSys{
		Protocol:  twProto.GGameHallProto,
		Protocol2: chat_proto.GS2COpenChatSysProto2,
		ChatInfo:  nil,
	}
	impl.PlayerSendToProxyServer(conn, senate, strOpenId)

	return
}

// 发送消息 数据操作
func SendChatInfo(conn *websocket.Conn, ProtocolData map[string]interface{}) {
	strOpenId := ProtocolData["OpenId"].(string)

	senate := chat_proto.GS2COpenChatSys{
		Protocol:  twProto.GGameHallProto,
		Protocol2: chat_proto.GS2COpenChatSysProto2,
		ChatInfo:  nil,
	}
	impl.PlayerSendToProxyServer(conn, senate, strOpenId)

	return
}
