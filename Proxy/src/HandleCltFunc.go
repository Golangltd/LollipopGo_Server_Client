package main

import (
	Proto_Proxy "LollipopGo/Proxy_Server/Proto"
	impl "LollipopGo/network"
	"LollipopGo/util"
	"fmt"
)

// 客户端发消息---> 服务器保存
func (this *ProxyServer) ClientSendDataToServer(ProtocolData map[string]interface{}) {
	if ProtocolData["ServerID"] == nil || ProtocolData["Data"] == nil {
		panic("error")
	}
	strServerID := ProtocolData["ServerID"].(string)
	strData := ProtocolData["Data"].(map[string]interface{})
	val, _ := this.MapSafe.Get(strServerID + "|Server")
	if val != nil {
		val.(*ProxyServer).PlayerSendMessage(strData)
	}
	return
}

// 客户端发消息---> 服务器保存
func (this *ProxyServer) ServerSendDataToClient(ProtocolData map[string]interface{}) {
	if ProtocolData["OpenID"] == nil || ProtocolData["Data"] == nil {
		panic("error")
	}
	strOpenID := ProtocolData["OpenID"].(string)
	strData := ProtocolData["Data"].(map[string]interface{})
	val, err := this.MapSafe.Get(strOpenID + "|User")
	if err != nil {
		fmt.Println("ServerSendDataToClient:", err)
	}
	if val != nil {
		val.(*ProxyServer).PlayerSendMessage(strData)
	}
	return
}

// 客户端连接
func (this *ProxyServer) User_Login(ProtocolData map[string]interface{}) {
	strOpenID := ""
	if ProtocolData["OpenID"] == nil {
		strOpenID = util.MD5_LollipopGO(util.UTCTime_LollipopGO())
	} else {
		strOpenID = ProtocolData["OpenID"].(string)
	}
	onlineUser := &ProxyServer{
		Connection: this.Connection,
		MapSafe:    this.MapSafe,
		StrMD5:     strOpenID,
	}
	this.MapSafe.Put(strOpenID+"|User", onlineUser)
	data := Proto_Proxy.Proxy2C_ConnData{
		Protocol:  1,
		Protocol2: Proto_Proxy.Proxy2C_ConnDataProto,
		OpenID:    strOpenID,
	}
	this.PlayerSendMessage(data)
	return
}

// 服务器连接
func (this *ProxyServer) Server_Login(ProtocolData map[string]interface{}) {
	if ProtocolData["ServerID"] == nil {
		panic("error")
	}
	onlineUser := &ProxyServer{
		Connection: this.Connection,
		MapSafe:    this.MapSafe,
	}
	fmt.Println("test data:", ProtocolData)
	serverid := ProtocolData["ServerID"].(string)
	onlineUser.StrMD5 = serverid
	this.MapSafe.Put(serverid+"|Server", onlineUser)
	data := Proto_Proxy.Proxy2C_ConnData{
		Protocol:  1,
		Protocol2: Proto_Proxy.Proxy2G_ConnDataProto,
	}
	this.PlayerSendMessage(data)
	val, _ := this.MapSafe.Get(serverid + "|Server")
	if val == nil {
	} else {
		impl.PlayerSendToProxyServer(val.(interface{}).(*ProxyServer).Connection, data, "www.8866.ltd:client")
		val.(interface{}).(*ProxyServer).PlayerSendMessage(nil)
		return
	}
	return
}

func (this *ProxyServer) Get_UserConn(strOpenID string) {
	val, _ := this.MapSafe.Get(strOpenID + "|connect")
	if val == nil {
	} else {
		val.(interface{}).(*ProxyServer).PlayerSendMessage("test")
		return
	}
}
