package game

import (
	"LollipopGo/util"
	"fmt"
	twlib_user "github.com/Golangltd/Twlib/user"
	"github.com/golang/glog"
	"strconv"
)


type SItem struct {
	RoleUID int64
	Account int64
	ItemInfo *twlib_user.ItemData
}

//  保存道具，首先看看道具是否存在
func (m *GameRPC)CreatItemInfo(data SItem,reply *interface{}) error {

	// 更新
	if true{
		sqlStr := `update gl_item set itemnum=itemnum+? where itemid = ? and  roleuid = ?`
		ret, err := GetGameConn().Exec(sqlStr, data.ItemInfo.ItemNum,data.ItemInfo.ItemId,data.RoleUID)
		if err != nil {
			fmt.Printf("update failed, err:%v\n", err)
			return err
		}
		n, err := ret.RowsAffected() //RowsAffected 受影响的行数
		if err != nil {
			fmt.Printf("get id failed,err:%v\n", err)
			return err
		}
		if n>0{
			*reply ,_= ret.LastInsertId()
			return nil
		}
	}

	sql:=  `insert into gl_item(uid, accountid,roleuid,itemid,itemtype,itemnum,creattime) 
           values(`
	val:= strconv.Itoa(int(data.RoleUID))+",'"+strconv.Itoa(int(data.Account))+"','"+strconv.Itoa(int(data.RoleUID))+"',"+strconv.Itoa(int(data.ItemInfo.ItemId))+
		","+strconv.Itoa(int(data.ItemInfo.ItemType))+","+strconv.Itoa(int(data.ItemInfo.ItemNum)) +","+util.UTCTime_LollipopGO()+")"
	sql=sql+val
	fmt.Println("-------------------->:",sql)
	ret, err := GetGameConn().Exec(sql)
	if err !=nil{
		fmt.Println(err.Error())
		return err
	}
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get id failed,err:%v\n", err)
		return err
	}
	sqlStr := `update gl_item set uid=? where id = ?`
	GetGameConn().Exec(sqlStr, id, id)
	*reply = id
	return nil
}

// 获取道具信息
func (m *GameRPC)  GetItemInfo(RoleUid int64,replydata *map[int64][]*twlib_user.ItemData)  error{

	sql := "select uid, itemid,itemtype,itemnum from gl_item where roleuid = '" +strconv.FormatInt(RoleUid, 10)+"'"
	rows, err := GetGameConn().Query(sql)
	glog.Info("-----",sql)
	if err != nil {
		fmt.Printf("exec %s query failed, err:%v\n", sql, err)
		return nil
	}
	// 3. 一定要关闭rows，才会释放连接（数据库的连接）
	defer rows.Close()
	// 4. 循环取值
	data := make(map[int64][]*twlib_user.ItemData)
	dataslice := []*twlib_user.ItemData{}
	for rows.Next() {
		datatmp  := new(twlib_user.ItemData)
		uid:=0
		id:=0
		itype := 0
		inum:=0
		err := rows.Scan(&uid, &id, &itype,&inum)
		if err != nil {
			fmt.Printf("scan failed,err:%v\n", err)
		}
		glog.Info("======uid",uid)

		datatmp.ItemUid = int64(uid)
		datatmp.ItemId = id
		datatmp.ItemType = itype
		datatmp.ItemNum = int64(inum)

		dataslice = append(dataslice,datatmp)
	}
	data[RoleUid] = dataslice
	replydata = &data
	glog.Info("======",dataslice)
	return nil
}

// 获取道具信息
func GetItemInfotmp(RoleUid int64)   map[int64][]*twlib_user.ItemData{

	sql := "select uid, itemid,itemtype,itemnum from gl_item where roleuid = '" +strconv.FormatInt(RoleUid, 10)+"'"
	rows, err := GetGameConn().Query(sql)
	glog.Info("-----",sql)
	if err != nil {
		fmt.Printf("exec %s query failed, err:%v\n", sql, err)
		return nil
	}
	// 3. 一定要关闭rows，才会释放连接（数据库的连接）
	defer rows.Close()
	// 4. 循环取值
	data := make(map[int64][]*twlib_user.ItemData)
	dataslice := []*twlib_user.ItemData{}
	for rows.Next() {
		datatmp  := new(twlib_user.ItemData)
		uid:=0
		id:=0
		itype := 0
		inum:=0
		err := rows.Scan(&uid, &id, &itype,&inum)
		if err != nil {
			fmt.Printf("scan failed,err:%v\n", err)
		}
		glog.Info("======uid",uid)
		datatmp.ItemUid = int64(uid)
		datatmp.ItemId = id
		datatmp.ItemType = itype
		datatmp.ItemNum = int64(inum)

		dataslice = append(dataslice,datatmp)
	}
	data[RoleUid] = dataslice
	glog.Info("======",dataslice)
	return data
}