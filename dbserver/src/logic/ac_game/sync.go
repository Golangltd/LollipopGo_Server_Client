package ac_game

import (
	"encoding/json"
	"fmt"
	twlib_user "github.com/Golangltd/Twlib/user"
	"strconv"
)

// 主要是同步数据---> 区域服数据同步到 账号服务器
// 同步创建角色数据
func SyncCreateRoleInfo(serverid int,accountid int64,userst *twlib_user.UserSt)error{
	data,_:=json.Marshal(userst)
	sql:=`insert into ac_serverlist(accountid, serverid,userinfo)
     values(`+strconv.FormatInt(accountid,10) +","+strconv.Itoa(serverid)+",'"+string(data)+"')"
	fmt.Println(sql)
	ret, err := GetAcGameConn().Exec(sql)
	if err !=nil{
		fmt.Println(err.Error())
		return err
	}
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get id failed,err:%v\n", err)
		return err
	}
	fmt.Println("---------id",id)
	return nil
}