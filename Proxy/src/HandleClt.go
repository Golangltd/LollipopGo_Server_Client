package main

import (
	Proto_Proxy "LollipopGo/Proxy_Server/Proto"
	"fmt"
	twlib_proto "github.com/Golangltd/Twlib/proto"
	"github.com/golang/glog"
	"golang.org/x/net/websocket"
)


// 主协议处理 HandleCltProtocol
func (this *ProxyServer) HandleCltProtocol(protocol interface{}, protocol2 interface{}, ProtocolData map[string]interface{}, Connection interface{}) interface{} {
	fmt.Println("-----进入消息：",protocol)
	fmt.Println("-----进入消息2：",protocol2)
	switch protocol {
	case float64(twlib_proto.GGameBattleProto):
		{
			this.HandleCltProtocol2(protocol2, ProtocolData, Connection.(*websocket.Conn))
		}
	case float64(twlib_proto.GameDataProto):
		{
			this.HandleCltProtocol22(protocol2, ProtocolData, Connection.(*websocket.Conn))
		}

	default:
		glog.Info("protocol default")
	}
	return 0
}

// 代理服务器处理
func (this *ProxyServer) HandleCltProtocol22(protocol2 interface{}, ProtocolData map[string]interface{}, Connection *websocket.Conn) interface{} {
	ConnectionData := &ProxyServer{
		Connection: Connection,
		MapSafe:    M,
	}
	switch protocol2 {
	case float64(Proto_Proxy.C2Proxy_ConnDataProto):
		{ // 客户端连接协议
			ConnectionData.User_Login(ProtocolData)
		}
	case float64(Proto_Proxy.G2Proxy_ConnDataProto):
		{ // 服务器连接协议
			ConnectionData.Server_Login(ProtocolData)
		}
	case float64(Proto_Proxy.C2Proxy_SendDataProto):
		{ // 转发协议 客户端发--->转服务器
			ConnectionData.ClientSendDataToServer(ProtocolData)
		}
	case float64(Proto_Proxy.G2Proxy_SendDataProto):
		{ // 转发协议 服务器发--->转客户端
			ConnectionData.ServerSendDataToClient(ProtocolData)
		}

	default:
		glog.Info("protocol2 default", ProtocolData)
	}
	return 0
}

// 子协议处理
func (this *ProxyServer) HandleCltProtocol2(protocol2 interface{}, ProtocolData map[string]interface{}, Connection interface{}) interface{} {
    /*
		ConnectionData := &ProxyServer{
			Connection: Connection,
			MapSafe:    M,
		}
	*/
	switch protocol2 {
	default:
		glog.Info("protocol2 default")
	}
	return 0
}
