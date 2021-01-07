package main

import (
	"LollipopGo"
	Proto_Proxy "LollipopGo/Proxy_Server/Proto"
	impl "LollipopGo/network"
	"LollipopGo/util"
	"LollipopGo2.8x/conf"
	"LollipopGo2.8x/conf/g"
	. "LollipopGo2.8x/cxt"
	bigsmall "LollipopGo2.8x/handlers/big-small"
	"LollipopGo2.8x/models"
	"LollipopGo2.8x/tables"
	"LollipopGo2.8x/util/req"
	"encoding/base64"
	"encoding/json"
	twlib_proto "github.com/Golangltd/Twlib/proto"
	twlib_server "github.com/Golangltd/Twlib/server"
	"github.com/Golangltd/cache2go"
	concurrent "github.com/fanliao/go-concurrentMap"
	"github.com/golang/glog"
	"golang.org/x/net/websocket"
	"strconv"
	"strings"
)

func init() {
	LollipopGo.Run()
	impl.IMsg = new(models.Game)
	CacheGame = cache2go.Cache("GameCache")
	M = concurrent.NewConcurrentMap()
	conf.InitConfig()
	InitProxyNet()
}

// 连接proxy服务器
func InitProxyNet() {
	proxyURL := req.AddParamsToGetReq(g.Ws, conf.ServerConfig().GetProxyUrlList(), map[string]string{"data": "{ID:1}"})
	glog.Infof("connect to proxy addr:%s\n", proxyURL)
	conn, err := websocket.Dial(proxyURL, "", "test://golang/")
	if err != nil {
		glog.Errorln("err:", err.Error())
		return
	}
	ConnXZ = conn
	// 保存内部到配置模块连接
	tables.SetConfigConn(ConnXZ)
	// bs
	bigsmall.SetConn(conn)
	//--------------------------------------------------------------------------
	// 发送保存链接
	data := Proto_Proxy.G2Proxy_ConnData{
		Protocol:  twlib_proto.GameDataProto,
		Protocol2: Proto_Proxy.G2Proxy_ConnDataProto,
		ServerID:  util.MD5_LollipopGO(strconv.Itoa(twlib_server.GameServerId)),
	}
	impl.PlayerSendToServer(conn, data)
	go GameServerReceive(ConnXZ)
}

func GameServerReceive(ws *websocket.Conn) {
	for {
		var content string
		err := websocket.Message.Receive(ws, &content)
		if err != nil {
			continue
		}
		//glog.Info(strings.Trim("", "\""))
		//glog.Info(content)
		content = strings.Replace(content, "\"", "", -1)
		contentstr, errr := base64Decode([]byte(content))
		if errr != nil {
			glog.Errorln(errr)
			continue
		}
		//glog.Info("收到数据：", string(contentstr))
		go SyncMessageFun(string(contentstr))
	}
}

func SyncMessageFun(content string) {
	var r Requestbody
	r.req = content

	if ProtocolData, err := r.Json2map(); err == nil {
		HandleCltProtocolXL(ProtocolData["Protocol"], ProtocolData["Protocol2"], ProtocolData)
	} else {
		glog.Info("解析失败：", err.Error())
	}
}

func base64Decode(src []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(src))
}

// 结构体数据类型
type Requestbody struct {
	req string
}

func (r *Requestbody) Json2map() (s map[string]interface{}, err error) {
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(r.req), &result); err != nil {
		glog.Error("Json2map:", err.Error())
		return nil, err
	}
	return result, nil
}
