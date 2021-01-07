package main

import (
	impl "LollipopGo/network"
	"LollipopGo/util"
	"LollipopGo2.8x/cxt"
	"LollipopGo2.8x/proto/proto_net"
	"fmt"
	twProto "github.com/Golangltd/Twlib/proto"
	"github.com/golang/glog"
	"golang.org/x/net/websocket"
)

func HandleCltProtocol2Net(protocol2 interface{}, ProtocolData map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			serer := fmt.Sprintf("%s", err)
			glog.Errorln(serer)
		}
	}()
	switch protocol2 {
	case float64(proto_net.Net_HeartBeatProto):
		{
			HeartBeat(cxt.ConnXZ,ProtocolData)
		}
	}
}

func HeartBeat(conn *websocket.Conn, ProtocolData map[string]interface{})  {

	strOpenId := ProtocolData["OpenId"].(string)

	senate := proto_net.Net_HeartBeat{
		Protocol:  twProto.GameNetProto,
		Protocol2: proto_net.Net_HeartBeatProto,
		TimeStamp: util.GetNowUnix_LollipopGo(),
	}
	impl.PlayerSendToProxyServer(conn, senate, strOpenId)
	return
}
