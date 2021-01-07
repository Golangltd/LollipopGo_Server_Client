package gamedb

import (
	"LollipopGo2.8x/conf"
	"github.com/golang/glog"
	"net/rpc"
	"net/rpc/jsonrpc"
)

var ConnRPC *rpc.Client

func DBInit() {
	client, err := jsonrpc.Dial("tcp", conf.GetConfig().GetDBAddr())
	if err != nil {
		glog.Info("dial error:", err)
		return
	}
	ConnRPC = client
}
