package chat

import (
	"LollipopGo2.8x/proto/chat_proto"
	"fmt"
	"github.com/golang/glog"
	"golang.org/x/net/websocket"
	"runtime/debug"
)

// 商城系统处理
func HandleCltProtocol2Chat(conn *websocket.Conn, protocol2 interface{}, protocolData map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			glog.Errorln(fmt.Sprintf("ERROR:[%s]\nSTACK:[%s\n]", err, string(debug.Stack())))
		}
	}()
	switch protocol2 {
	case float64(chat_proto.C2GSOpenChatSysProto2):
		GetChatInfo(conn, protocolData)
	case float64(chat_proto.C2GSSendChatSysProto2):
		SendChatInfo(conn, protocolData)
	default:
		fmt.Println("协议错误")
	}
}
