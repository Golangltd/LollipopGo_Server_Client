package DB

import (
	"github.com/golang/glog"
	"login/conf"
	"net/rpc"
	"net/rpc/jsonrpc"
)

var ConnRPC *rpc.Client

func DBInit() {
	client, err := jsonrpc.Dial("tcp", conf.GetConfig().Server.DBAddr)
	if err != nil {
		glog.Info("dial error:", err)
		return
	}
	ConnRPC = client
}

// 获取链接信息
func GetConnRPC()*rpc.Client  {
	if ConnRPC != nil{
        return ConnRPC
	}else {
		client, err := jsonrpc.Dial("tcp", conf.GetConfig().Server.DBAddr)
		if err != nil {
			glog.Info("dial error:", err)
			return nil
		}
		return client
	}
}

// 实际操作信息
func DB_rpc_SysRole(RoleName, RoleDesc, Operator, authority string) interface{} {
	args := 1
	var reply bool
	// rpc固定的操作 操作!
	divCall := ConnRPC.Go("Arith.BinGeSysRole", args, &reply, nil)
	replyCall := <-divCall.Done
	glog.Info(replyCall.Reply)
	glog.Info("the arith.BinGeSysRole is :", reply)
	return reply
}