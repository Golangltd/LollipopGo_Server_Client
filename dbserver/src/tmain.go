package main

import (
	lua_uitl "LollipopGo2.8x/Lua/Uitl"
	_go "LollipopGo2.8x/Lua/go"
	"LollipopGo2.8x/conf"
	"LollipopGo2.8x/logic/ac_game"
	"LollipopGo2.8x/logic/cf_game"
	Mysyl_DB "LollipopGo2.8x/logic/db"
	"LollipopGo2.8x/logic/game"
	"fmt"
	"github.com/golang/glog"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
)

func main() {
	conf.InitConfig()
	Mysyl_DB.Init()
	MainListener(conf.GetConfig().Server.WSAddr)
}

func MainListener(strport string) {
	rpcRegister()
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":"+strport)
	checkError(err)
	Listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	glog.Infof("db listen to port:%s\n", strport)
	glog.Info("db start ok")
	for {
		defer func() {
			if err := recover(); err != nil {
				strerr := fmt.Sprintf("%s", err)
				fmt.Println("异常捕获:", strerr)
			}
		}()
		conn, err := Listener.Accept()
		if err != nil {
			fmt.Fprint(os.Stderr, "accept err: %s", err.Error())
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}

func rpcRegister() {
	_ = rpc.Register(new(ac_game.AcRPC))
	_ = rpc.Register(new(cf_game.CfRPC))
	_ = rpc.Register(new(game.GameRPC))
}

func checkError(err error) {
	if err != nil {
		fmt.Fprint(os.Stderr, "Usage: %s", err.Error())
	}
}

func init() {
	_go.Initvariable_proto()
	conf_data := _go.Gvariable_proto["253"]
	fmt.Println("--------------253", conf_data.Data1)
	lua_uitl.GetLuaDataFor3Row(conf_data.Data1)
}
