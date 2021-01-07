package bigsmall

import (
	ball "LollipopGo2.8x/proto/big-small"
	"fmt"
	"github.com/golang/glog"
	"golang.org/x/net/websocket"
	"runtime/debug"
)

//  泰式比高低函数处理
func HandleCltProtocol2BS(conn *websocket.Conn, protocol2 interface{}, protocolData map[string]interface{}) {

	defer func() {
		if err := recover(); err != nil {
			glog.Errorln(fmt.Sprintf("ERROR:[%s]\nSTACK:[%s\n]", err, string(debug.Stack())))
		}
	}()

	switch protocol2 {
	case float64(ball.C2SEntryGameProto2):
		C2SEntryGameProto2Func(conn,protocolData)
	case float64(ball.C2SGetHistoryDataProto2):
		C2SGetHistoryDataProto2Func(conn,protocolData)
	case float64(ball.C2SUserChipInProto2):
		C2SUserChipInProto2Func(conn,protocolData)
	case float64(ball.C2SUserGetItemRateProto2):
		C2SUserGetItemRateProto2Func(conn,protocolData)
	case float64(ball.C2SUserGetGameStateProto2):
		C2SUserGetGameStateProto2Func(conn,protocolData)
	default:
		fmt.Println("协议错误")
	}

	return
}
