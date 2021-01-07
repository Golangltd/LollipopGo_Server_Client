package Shop

import (
	"LollipopGo2.8x/proto/shop_proto"
	"fmt"
	"github.com/golang/glog"
	"golang.org/x/net/websocket"
	"runtime/debug"
)

// 商城系统处理
func HandleCltProtocol2S(conn *websocket.Conn, protocol2 interface{}, protocolData map[string]interface{}) {

	defer func() {
		if err := recover(); err != nil {
			glog.Errorln(fmt.Sprintf("ERROR:[%s]\nSTACK:[%s\n]", err, string(debug.Stack())))
		}
	}()

	switch protocol2 {
	case float64(shop_proto.C2GS_GetShopInfoProto2):
		GetShopInfo(conn, protocolData, false)
	case float64(shop_proto.C2GS_UpdateShopInfoProto2):
		//UpdateShopInfo(conn, protocolData, false)
	case float64(shop_proto.C2GS_GetBuyGoodsProto2):
		GetBuyGoods(conn, protocolData)
	default:
		fmt.Println("协议错误")
	}
}
