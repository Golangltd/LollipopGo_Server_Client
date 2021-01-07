package game

import (
	"LollipopGo/log"
	"LollipopGo/util"
	lua_uitl "LollipopGo2.8x/Lua/Uitl"
	_go "LollipopGo2.8x/Lua/go"
	"LollipopGo2.8x/conf"
	"LollipopGo2.8x/logic/ac_game"
	"fmt"
	"github.com/Golangltd/Twlib/DbSt"
	twlib_user "github.com/Golangltd/Twlib/user"
	"strconv"
)

// 通过账户ID在角色表中获取角色信息
func (m *GameRPC) GetRoleInfo(accountId int64, reply *twlib_user.UserSt) (err error) {
	sql := "select roleuid,avatar,name,lev,channelid,sex,coin,diamond,creattime from gl_role where accountid=?"
	err = GetGameConn().QueryRow(sql, accountId).Scan(&reply.RoleUid, &reply.RoleAvatar, &reply.RoleName, &reply.RoleLev,
		&reply.ChannelId, &reply.RoleSex, &reply.Coin, &reply.Diamond, &reply.RegisterTime)
	if err != nil {
		log.Error("get role data failed!,err is %v", err)
		return
	}
	// 获取道具
	iteminfo := GetItemInfotmp(reply.RoleUid)
	if reply.ItemList == nil{
		reply.ItemList = []*twlib_user.ItemData{}
	}
	// 章节，关卡数据
	reply.ChapterInfo = GetChapterInfo(reply.RoleUid)
    // 道具
	reply.ItemList = iteminfo[reply.RoleUid]
	// 卡牌
	reply.CardList = GetCardInfo(reply.RoleUid)
	// 获取装备
	reply.EquipData = new(twlib_user.EquipData)
	reply.EquipData.EquipSts = GetEquipInfo(reply.RoleUid)
	return
}

func (m *GameRPC) CreatRoleInfo(RoleSt twlib_DbSt.RoleSt, reply *interface{}) error {
	if true {
		sql := "select roleuid,avatar,name,lev,sex,coin,diamond,creattime from gl_role where accountid=" + strconv.FormatInt(RoleSt.AccountId, 10)
		fmt.Println(sql)
		rows, err := GetGameConn().Query(sql)
		defer rows.Close()
		if err != nil {
			return err
		}
		recliner := twlib_user.UserSt{}
		for rows.Next() {
			rows.Scan(&recliner.RoleUid, &recliner.RoleAvatar, &recliner.RoleName, &recliner.RoleLev, &recliner.RoleSex,
				&recliner.Coin, &recliner.Diamond,&recliner.RegisterTime)
		}
		if recliner.RoleUid != 0 {
			recliner.RoleUid = -1
			*reply = recliner
			return nil //fmt.Errorf("role is exit")
		}
	}

	if true {
		sql := "select roleuid from gl_role where name='" + RoleSt.Name + "'"
		fmt.Println(sql)
		rows, err := GetGameConn().Query(sql)
		defer rows.Close()
		if err != nil {
			return err
		}
		recliner := twlib_user.UserSt{}
		for rows.Next() {
			rows.Scan(&recliner.RoleUid)
		}
		// 获取道具
		iteminfo :=GetItemInfotmp(recliner.RoleUid)
		if recliner.ItemList == nil{
			recliner.ItemList = []*twlib_user.ItemData{}
		}
		recliner.ItemList = iteminfo[recliner.RoleUid]
		if recliner.RoleUid != 0 {
			recliner.RoleUid = -1
			*reply = recliner
			return nil//fmt.Errorf("role is exit")
		}
	}

	datainfo := &twlib_user.UserSt{
		RoleName:   RoleSt.Name,
		RoleAvatar: 1,
		RoleLev:    1,
		RoleSex:    RoleSt.Sex,
		Coin:       0,
		Diamond:    0,
	}

	rtime :=util.GetNowUnix_LollipopGo()
	sql := `insert into gl_role(roleuid, avatar,name,lev,channelid,accountid,sex,coin,diamond,creattime) 
           values(`
	val := strconv.Itoa(0) + "," + "1,'" + RoleSt.Name + "'," + "1" + "," + "1" + "," + strconv.FormatInt(RoleSt.AccountId, 10) + "," +
		strconv.Itoa(RoleSt.Sex) + "," + "0" + "," + "0" + "," + strconv.FormatInt(rtime,10) + ")"
	sql = sql + val
	fmt.Println("-------------------->:", sql)
	ret, err := GetGameConn().Exec(sql)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get id failed,err:%v\n", err)
		return err
	}
	sqlStr := `update gl_role set roleuid=? where id = ?`
	ret, err = GetGameConn().Exec(sqlStr, id, id)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return err
	}
	n, err := ret.RowsAffected() //RowsAffected 受影响的行数
	if err != nil {
		fmt.Printf("get id failed,err:%v\n", err)
		return err
	}
	fmt.Printf("更新了%d行数据\n", n)
	datainfo.RoleUid = id
	// 创建道具
	var itemdatatmp = SItem{}
	itemdatatmp.Account = RoleSt.AccountId
	itemdatatmp.RoleUID = datainfo.RoleUid
	// 离散表
	conf_data:=_go.Gvariable_proto["253"]
	// 道具类型,道具Id,道具梳理
	itemmap:=lua_uitl.GetLuaDataFor3Row(conf_data.Data1)
	// fmt.Println("-----------------itemmap:",itemmap)
	for _,v:=range itemmap {
		itemdatatmp.ItemInfo = &twlib_user.ItemData{
			ItemUid:  0, // 唯一ID
			ItemId:   v.Id,
			ItemType: v.Type,
			ItemNum:  v.Num, // 道具的数量
		}
		fmt.Println("-----------------v.Nu====++++++++++++++++++++++++++++++++++++++++++m:",itemdatatmp.ItemInfo.ItemId)
		if itemdatatmp.ItemInfo.ItemId== 1{
			UpdateRoleCoin(datainfo.RoleUid,v.Num)
			datainfo.Coin+=v.Num
		}else if itemdatatmp.ItemInfo.ItemId== 2{
			UpdateRoleDiamond(datainfo.RoleUid,v.Num)
			datainfo.Diamond+=v.Num
		}
		retdata:=UpdateRoleItem(datainfo.RoleUid,v.Num, v.Id)
		// fmt.Printlng("-----------------v.Num:",v.Num)
		if retdata == 0 || retdata == -1{
			var iditem interface{}
			m.CreatItemInfo(itemdatatmp, &iditem)
			itemdatatmp.ItemInfo.ItemUid = iditem.(int64)
		}
		datainfo.ItemList = append(datainfo.ItemList,itemdatatmp.ItemInfo)
	}
	// 章节，关卡数据
	datainfo.ChapterInfo = GetChapterInfo(datainfo.RoleUid)
	// fmt.Println("-----------------item:",datainfo.ItemList)
	datainfo.RegisterTime = rtime
	*reply = datainfo
	// 发送给账号服务器 --> 数据保存操作 --- 通过http发送过去
	ac_game.SyncCreateRoleInfo(conf.GetConfig().Server.AreaListId, RoleSt.AccountId, datainfo)
	return nil
}

// 更新金币
func UpdateRoleCoin(roleuid int64 ,number int64){
	sqlStr := `update gl_role set coin=coin + ? where roleuid = ?`
	ret, err := GetGameConn().Exec(sqlStr, number, roleuid)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() //RowsAffected 受影响的行数
	if err != nil {
		fmt.Printf("get id failed,err:%v\n", err)
		return
	}
	fmt.Printf("更新了%d行数据\n", n)
}

// 更新砖石
func UpdateRoleDiamond(roleuid int64 ,number int64){
	sqlStr := `update gl_role set diamond=diamond+? where roleuid = ?`
	ret, err := GetGameConn().Exec(sqlStr, number, roleuid)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() //RowsAffected 受影响的行数
	if err != nil {
		fmt.Printf("get id failed,err:%v\n", err)
		return
	}
	fmt.Printf("更新了%d行数据\n", n)
}

// 更新道具
func UpdateRoleItem(roleuid int64 ,number int64,itemid int) int{
	sqlStr := `update gl_item set itemnum=itemnum +? where roleuid = ? and itemid =?`
	ret, err := GetGameConn().Exec(sqlStr, number, roleuid,itemid)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return -1
	}
	n, err := ret.RowsAffected() //RowsAffected 受影响的行数
	if err != nil {
		fmt.Printf("get id failed,err:%v\n", err)
		return -1
	}
	fmt.Printf("更新了%d行数据\n", n)
	return int(n)
}

