package main

import (
	"LollipopGo/network"
	"LollipopGo2.8x/conf"
	"github.com/Golangltd/cache2go"
	concurrent "github.com/fanliao/go-concurrentMap"
	"github.com/golang/glog"
	"golang.org/x/net/websocket"
	"net/http"
	"runtime"
)

type ProxyServer struct {
	Connection *websocket.Conn
	StrMD5     string
	MapSafe    *concurrent.ConcurrentMap
}

var (
	cache *cache2go.CacheTable
	M     *concurrent.ConcurrentMap
)

func main()  {
	conf.InitConfig()
	runtime.GOMAXPROCS(runtime.NumCPU())
	glog.Info("conf.GetConfig().Server.WSAddr",conf.GetConfig().Server.WSAddr)
	http.Handle("/"+conf.GetConfig().Server.URL, websocket.Handler(BuildConnection))
	if err := http.ListenAndServe(conf.GetConfig().Server.WSAddr, nil); err != nil {
		glog.Info("Entry nil", err.Error())
		glog.Flush()
		return
	}
}

func BuildConnection(ws *websocket.Conn) {
	data := ws.Request().URL.Query().Get("data")
	if data == "" {
		glog.Info("data is Nil")
		glog.Flush()
		return
	}
	impl.InitConnection(ws)
}