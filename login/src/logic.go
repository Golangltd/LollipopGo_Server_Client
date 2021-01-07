package main

import (
	"LollipopGo/util"
	"encoding/json"
	"fmt"
	twlib_proto "github.com/Golangltd/Twlib/proto"
	twlib_user "github.com/Golangltd/Twlib/user"
	"github.com/golang/glog"
	"login/DB"
	"login/conf"
	"login/msg"
	"net/http"
)

var GMapUser map[string]*twlib_user.UserSt
var GMapToken map[string]string


// 返回给htt请求的
func replyJson(resp http.ResponseWriter, content interface{}) {
	resp.Header().Set("Content-Type", "application/json")
	s, err := json.Marshal(content)
	if err != nil {
		glog.Error("can't marshal response to json:%v", err.Error())
		return
	}
	_, _ = resp.Write(s)
}

// 注册操作
func register(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json
	req.Header.Add("content-type", "charset=UTF-8")
	// 获取账号密码
	fmt.Println("req:",req)
	if req.Method == "GET" {// 正式上线 POST  header
		if accountable := req.FormValue("AccountName"); accountable != "" {
			if accountPw := req.FormValue("AccountPw"); accountPw != "" {
				// 获取玩家数据
		/*		data := twlib_proto.C2SUserLogin{
					AccountName: accountable,
					AccountPw:   util.MD5_LollipopGO(accountPw),
				}*/
				// 发送给DB sever 去验证数据
				ulcerate := &twlib_user.UserSt{
					RoleUid:int64(10+len(GMapUser)),
					RoleName:accountable,
					RoleAvatar:1,
					RoleLev:1,
					RoleSex:0,
				}//DbLogin(&data)
				fmt.Println(ulcerate)
				// token生成的机制 -- 19位时间戳 md5
				stricken := util.GetNowtimeMD5_LollipopGO() // 时间19位时间
				GMapUser[stricken] = ulcerate
				rcdata := &msg.LoginData{
					Token:      stricken,
					Url:        ulcerate.LatestArea, // 默认新区,或者近期登录
				}
				// 确保登录正确
				if ulcerate.RoleUid > 0 {
					rcdata.AesKey = conf.GetConfig().Server.AesKey
				}
				// 返回客户端
				sendData, _ := json.Marshal(rcdata)
				replyJson(w, sendData)
				return
			}
		}
		replyJson(w, "parameter  is wrong!")
	}
}

// 登录操作
func Login(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json
	req.Header.Add("content-type", "charset=UTF-8")
	// 获取账号密码
	fmt.Println("req:",req)
	if req.Method == "GET" {// 正式上线 POST  header
		if accountable := req.FormValue("AccountName"); accountable != "" {
			if accountPw := req.FormValue("AccountPw"); accountPw != "" {
				// 获取玩家数据
				data := twlib_proto.C2SUserLogin{
					AccountName: accountable,
					AccountPw:   util.MD5_LollipopGO(accountPw),
				}
				// 发送给DB sever 去验证数据
				ulcerate := DbLogin(&data)
				fmt.Println("+++++++++++++++++++++++============== ::---",ulcerate)
				// token生成的机制 -- 19位时间戳 md5
				stricken := util.GetNowtimeMD5_LollipopGO() // 时间19位时间
				GMapUser[stricken] = ulcerate
				rcdata := &msg.LoginData{
					Token:      stricken,
					Url:        ulcerate.LatestArea, // 默认新区,或者近期登录
				}
				// 确保登录正确
				if ulcerate.RoleUid > 0 {
					rcdata.AesKey = conf.GetConfig().Server.AesKey
				}
				// 返回客户端
				sendData, _ := json.Marshal(rcdata)
				replyJson(w, sendData)
				return
			}
		}
		replyJson(w, "parameter  is wrong!")
	}
}

// 获取游戏服务器区域列表
func servers(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json
	req.Header.Add("content-type", "charset=UTF-8")
	// 获取账号密码
	if req.Method == "GET" {
		if PageNum := req.FormValue("PageNum"); PageNum != "" {

			data1 := twlib_proto.C2SUserLogin{
				AccountName: "001",
				AccountPw:   util.MD5_LollipopGO("123456"),
			}
			// 发送给DB sever 去验证数据
			ulcerate := DbLogin(&data1)
			// 发送给DB sever 去验证数据
			//ulcerate := DbServerList(PageNum)
			data := msg.S2CUserGetServerList{
			}
			if ulcerate == nil{
				/*var datas []*twlib_user.ServerList
				server1:=&twlib_user.ServerList{
					ServerId:1,
					ServerName:"稳定测试服",
					Url        :"192.168.2.199:8888",
					State      :1,
				}
				datas = append(datas,server1)
				server2:=&twlib_user.ServerList{
					ServerId:1,
					ServerName:"功能验证服",
					Url        :"192.168.2.115:8888",
					State      :3,
				}
				datas = append(datas,server2)*/
				data.SeverList = ulcerate.ServerList
			}else {
				data.SeverList = ulcerate.ServerList
			}
			sendData, _ := json.Marshal(data)
			fmt.Println("-----------------------:sendData",data.SeverList[0].UserInfo)
			replyJson(w, sendData)
		} else {
			// 返回错误数据
		}
	}
}

// 登录操作
func GameToken(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Accss-Control-Allow-Origin", "*")
	req.Header.Add("content-type", "charset=UTF-8")
	if req.Method == "GET" {
		if stricken := req.FormValue("token"); stricken != "" {
			data := msg.S2CUserLogin{
				UserData: GMapUser[stricken], // 返回区域列表连接数据，game server 拿到数据后，去区域服务器的db server 获取角色信息
			}
			// 返回给game server 数据
			sendData, _ := json.Marshal(data.UserData)
			fmt.Fprint(w,string(sendData))
			//replyJson(w, sendData)
			return
		} else {
			// 统一错误处理
		}
	}
}

// 登录流程
// DB 需要完成相应的注册函数,获取玩家结构数据，
/*func DbLogin(data *twlib_proto.C2SUserLogin) *string {
	var proxyUrl *string
	call := DB.ConnRPC.Go("AcRPC.GetUserLogin", data, &proxyUrl, nil)
	replyCall := <-call.Done
	glog.Info(replyCall.Reply)
	return proxyUrl
}*/

// db 反向代理返回--- 玩家账号信息，玩家的名字，等级
func DbLogin(data *twlib_proto.C2SUserLogin) *twlib_user.UserSt {
	var user *twlib_user.UserSt
	call := DB.ConnRPC.Go("AcRPC.GetUserLogin", data, &user, nil)
	replyCall := <-call.Done
	glog.Info(replyCall.Reply)
	return user
}

// 获取游戏服务器列表，分页流程 30数据一页
func DbServerList(data string) *[]*twlib_user.ServerList {
	var servers *[]*twlib_user.ServerList
	call := DB.ConnRPC.Go("AcRPC.GetAreaUrl", data, &servers, nil)
	replyCall := <-call.Done
	glog.Info(replyCall.Reply)
	return servers
}
