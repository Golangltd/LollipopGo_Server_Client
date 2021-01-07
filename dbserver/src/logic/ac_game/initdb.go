package ac_game

import (
	Mysyl_DB "LollipopGo2.8x/logic/db"
	"database/sql"
	"encoding/json"
	"fmt"
	twlib_proto "github.com/Golangltd/Twlib/proto"
	twlib_user "github.com/Golangltd/Twlib/user"
	"github.com/golang/glog"
	"strconv"
)

type AcRPC struct {
}

func GetAcGameConn() *sql.DB {
	return Mysyl_DB.DB.STdb["ac_game"]
}

// 获取账号Id
func (this *AcRPC) GetUserLogin(data *twlib_proto.C2SUserLogin, reply *twlib_user.UserSt) error {

	// 测试阶段数据操作
	defer func() {
		if err := recover(); err != nil {
			serer := fmt.Sprintf("%s", err)
			glog.Errorln(serer)
		}
	}()

	sql := "select id,avatar,name,sex,lev,areacur from ac_account where loginname='" + data.AccountName + "' AND loginpw='" + data.AccountPw + "'"
	fmt.Println(sql)
	rows, err := GetAcGameConn().Query(sql)
	defer rows.Close()
	if err != nil {
		return err
	}
	recliner := twlib_user.UserSt{}
	for rows.Next() {
		rows.Scan(&recliner.RoleUid, &recliner.RoleAvatar, &recliner.RoleName, &recliner.RoleSex, &recliner.RoleLev,
			&recliner.RoleAvatar)
		recliner.ChannelId, _ = strconv.Atoi(recliner.LatestArea)
	}
	fmt.Println(recliner)
	// 没有账号数据
	if recliner.RoleUid == 0{
		recliner = *CreateRoleInfo(data.AccountName)
	}

	recliner.ServerList = GetAreacur(&recliner)
	fmt.Println("------",recliner.ServerList)
	for k,v:=range recliner.ServerList{
		fmt.Println("------",k)
		fmt.Println("------",v.UserInfo)
	}
	recliner.LatestArea = "192.168.1.5:8888"
	*reply = recliner
	return nil
}

func CheckReName(name string) bool {
	sql := "select id from ac_account where loginname='" + name + "'"
	fmt.Println(sql)
	rows, err := GetAcGameConn().Query(sql)
	defer rows.Close()
	if err != nil {
		return false
	}
	recliner := twlib_user.UserSt{}
	for rows.Next() {
		rows.Scan(&recliner.RoleUid)
	}
	// 用户名存在
	if recliner.RoleUid !=0{
		return true
	}
	return false
}

func GetAreacur( accounting *twlib_user.UserSt) []*twlib_user.ServerList {
	sql := "select channel,name,state,url from ac_gamelist"
	fmt.Println(sql)
	rows, err := GetAcGameConn().Query(sql)
	if err != nil {
		return nil
	}
	var regret []*twlib_user.ServerList
	serverless :=GetAccountIdOfAllServerList()
	for rows.Next() {
		repleted := new(twlib_user.ServerList)
		rows.Scan(&repleted.ServerId, &repleted.ServerName, &repleted.State, &repleted.Url)
		da:= serverless[strconv.FormatInt(accounting.RoleUid,10)+"|"+strconv.Itoa(repleted.ServerId)]
		//if repleted.ServerId == 2{
		//	da:= serverless[strconv.FormatInt(1,10)+"|"+strconv.Itoa(1)]
			if da!=nil{
				repleted.UserInfo = da
			}
		//}

		regret = append(regret, repleted)
	}
	if len(regret) > 0 {
		fmt.Println(regret[0])
	}
	fmt.Println(regret)
	return regret
}

// 查询玩家所在的服所创建的所有的游戏服务器账号
func GetAccountIdOfAllServerList() map[string]*twlib_user.UserSt {
	sql := "select accountid,serverid,userinfo from ac_serverlist"
	fmt.Println(sql)
	rows, err := GetAcGameConn().Query(sql)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	data := make(map[string]*twlib_user.UserSt)
	for rows.Next() {
		account := int64(0)
		server := 0
		userinfo := ""
		rows.Scan(account, server, userinfo)
		fmt.Println(userinfo)
		ushering := &twlib_user.UserSt{}
		err := json.Unmarshal([]byte(userinfo), ushering)
		if err != nil {
			fmt.Println(err)
			continue
		}
		data[strconv.FormatInt(account,10)+"|"+strconv.Itoa(server)] = ushering
	}
	fmt.Println("============================data===============",data)
	return data
}

// 获取区域服的数据
func (this *AcRPC) GetAreaUrl(pagenum int, reply *[]*twlib_user.ServerList) error {

	sql := "select channel,name,state,url from ac_gamelist"
	fmt.Println(sql)
	rows, err := GetAcGameConn().Query(sql)
	if err != nil {
		return nil
	}
	var regret []*twlib_user.ServerList
	for rows.Next() {
		repleted := new(twlib_user.ServerList)
		rows.Scan(&repleted.ServerId, &repleted.ServerName, &repleted.State, &repleted.Url)
		regret = append(regret, repleted)
	}
	if len(regret) > 0 {
		fmt.Println(regret[0])
	}
	fmt.Println(regret)
	reply = &regret
	return nil
}
