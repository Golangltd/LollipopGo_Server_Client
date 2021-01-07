package ac_game

import (
	"LollipopGo/util"
	"fmt"
	twlib_user "github.com/Golangltd/Twlib/user"
	"github.com/golang/glog"
)

/*
    data： 数据处理
	var req float64 = 3.0
	var respSync *float64
	//异步的调用方式
	syncCall := client.Go("MathUtil.CalculateCircleArea", req, &respSync, nil)
	replayDone := <-syncCall.Done
	fmt.Println(replayDone)
	fmt.Println(*respSync)

	RoleUid    int64
	RoleName   string
	RoleAvatar int
	RoleLev    int
	RoleSex    int
	Coin       int64         // 金币
	Diamond    int64         // 砖石
	CardList   []*CardSt     // 角色拥有的卡牌
	LatestArea string        // 上一次的最新登录的区   区的url：ip+port
	ItemList   []*ItemSt     // 角色身上的道具，包括装备
	ChannelId  int           // 渠道Id
	ServerList []*ServerList // 整个游戏的所有区列表，从上线开始  1-30  29

*/
// 如果获取玩家数据失败的情况下,
func CreateRoleInfo(name string)  *twlib_user.UserSt {
    fmt.Println("--------------:CreateRoleInfo")
	defer func() {
		if err := recover(); err != nil {
			serer := fmt.Sprintf("%s", err)
			glog.Errorln(serer)
		}
	}()

	user := &twlib_user.UserSt{
		RoleName:name,
		RoleAvatar:1,
		RoleLev:1,
		RoleSex:1,
		Coin:0,
		Diamond:0,
	}

	sql:=  `insert into ac_account(loginname,loginpw,avatar,name,sex,lev,areacur,createtime) 
           values('`+user.RoleName+"','"+"e10adc3949ba59abbe56e057f20f883e"+"',"+"1"+","+"name"+","+"1"+","+"1"+","+"1"+""+",'"+
		util.GetTime_LollipopGO()+"')"
	fmt.Println(sql)
	ret, err := GetAcGameConn().Exec(sql)
	if err !=nil{
		fmt.Println(err.Error())
		return nil
	}
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get id failed,err:%v\n", err)
		return nil
	}
	fmt.Println("id:",id)
	user.RoleUid = id
	return user
}