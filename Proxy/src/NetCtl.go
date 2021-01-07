package main

import (
	"LollipopGo"
	impl "LollipopGo/network"
	"encoding/json"
	"github.com/Golangltd/cache2go"
	concurrent "github.com/fanliao/go-concurrentMap"
	"github.com/golang/glog"
	"golang.org/x/net/websocket"
	"strings"
)

// 初始化数据
func init()  {
	LollipopGo.Run()
	impl.IMsg = new(ProxyServer)
	cache = cache2go.Cache("ProxyCache")
	M = concurrent.NewConcurrentMap()
}

func (this *ProxyServer) PlayerSendMessage(senddata interface{}) int {
	b, err1 := json.Marshal(senddata)
	if err1 != nil {
		glog.Error("PlayerSendMessage json.Marshal data fail ! err:", err1.Error())
		glog.Flush()
		return 1
	}
	err := websocket.JSON.Send(this.Connection, b)
	if err != nil {
		glog.Error("PlayerSendMessage send data fail ! err:", err.Error())
		glog.Flush()
		return 2
	}
	return 0
}

// 玩家退出
type PlayerExit struct {
	Protocol  int
	Protocol2 int
	OpenId    string
}

// 除代理服务器外，其他服务器不用实现
func (this *ProxyServer)CloseEOF(closeEvent interface{}) int {
    for itr :=M.Iterator();itr.HasNext();{
    	k,val,_:=itr.Next()
    	if val.(*ProxyServer).Connection == closeEvent{
    		M.Remove(k)
			val.(*ProxyServer).Connection.Close()
    		glog.Info("玩家退出：",val.(*ProxyServer).StrMD5)
            if true {
				for itrs := M.Iterator(); itrs.HasNext(); {
					ks, vals, _ := itrs.Next()
					if strings.Contains(ks.(string), "|Server") {
						data := &PlayerExit{
							Protocol:  1005,
							Protocol2: 1005,
							OpenId:    val.(*ProxyServer).StrMD5,
						}
						impl.PlayerSendMessageOfExit(vals.(*ProxyServer).Connection, data, vals.(*ProxyServer).StrMD5)
					}
				}
			}
		}
	}
	return 0
}