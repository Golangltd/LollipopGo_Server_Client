package main

import (
	twlib_user "github.com/Golangltd/Twlib/user"
	"github.com/golang/glog"
	"login/DB"
	"login/conf"
	"net/http"
)

func init()  {
	GMapUser = make(map[string]*twlib_user.UserSt)
	GMapToken = make(map[string]string)
}

func main() {
	conf.InitConfig()
	DB.DBInit()
	http.HandleFunc("/"+conf.GetConfig().Server.URL+"/client/login", Login)     // 登录流程
	http.HandleFunc("/"+conf.GetConfig().Server.URL+"/client/register", register)  // 注册流程
	http.HandleFunc("/"+conf.GetConfig().Server.URL+"/server/list", servers)    // 获取区域列表
	http.HandleFunc("/"+conf.GetConfig().Server.URL+"/server/login", GameToken) // 服务器校验，token数据流程
	err := http.ListenAndServe(conf.GetConfig().Server.HTTPAddr, nil)
	if err != nil {
		glog.Errorln("http.ListenAndServe()函数执行错误,错误为:%v\n", err)
		return
	}
}