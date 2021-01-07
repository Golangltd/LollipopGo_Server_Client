package game

import (
	"LollipopGo/util"
	"fmt"
	twLibUser "github.com/Golangltd/Twlib/user"
	"github.com/golang/glog"
	"strconv"
)

type SEquip struct {
	RoleUID int64
	Account int64
	Equip    *twLibUser.EquipSt
}

// 创建装备
func (m *GameRPC) CreatEquip(data *SEquip, reply *interface{}) error {

	if data.Equip == nil{
		glog.Info("-----------------------------------------------------------CreatEquip")
		*reply = -1
		return nil
	}
	// 更新
	if true{
		sqlStr := `update gl_equip set num=num+? where equipid = ? AND  roleuid = ?`
		ret, err := GetGameConn().Exec(sqlStr, data.Equip.Num,data.Equip.ConfID,data.RoleUID)
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
			*reply = n
			return nil
		}
	}

	sql:=  `insert into gl_equip(roleuid, accountid,equipid,num,createtime,star,camp,camprate,exp,power) 
           values(`

	val:= strconv.Itoa(int(data.RoleUID))+",'"+strconv.Itoa(int(data.Account))+"','"+strconv.Itoa(data.Equip.ConfID)+
		"',"+strconv.Itoa(data.Equip.Num)+ ","+util.UTCTime_LollipopGO()+","+strconv.Itoa(data.Equip.Star)+
		","+strconv.Itoa(data.Equip.Camp)+
		","+strconv.Itoa(data.Equip.CampRate)+
		","+strconv.Itoa(data.Equip.Exp)+
		","+strconv.Itoa(data.Equip.Power)+
		")"
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

	// 更新UID
	if true{
		sqlStr := `update gl_equip set equipuid=? where id = ?`
		ret, err = GetGameConn().Exec(sqlStr, id,id)
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
	}

	*reply = id
	return nil
}

// 获取装备
func GetEquipInfo(RoleUid int64)   []*twLibUser.EquipSt{

	sql := "select equipid,num,star,camp,camprate,exp,power,equipuid from gl_equip where roleuid = '" +strconv.FormatInt(RoleUid, 10)+"'"
	rows, err := GetGameConn().Query(sql)
	glog.Info("-----",sql)
	if err != nil {
		fmt.Printf("exec %s query failed, err:%v\n", sql, err)
		return nil
	}
	// 3. 一定要关闭rows，才会释放连接（数据库的连接）
	defer rows.Close()
	// 4. 循环取值
	data := make(map[int64][]*twLibUser.EquipSt)
	dataslice := []*twLibUser.EquipSt{}
	for rows.Next() {
		datatmp  := new(twLibUser.EquipSt)
		err := rows.Scan(&datatmp.ConfID, &datatmp.Num, &datatmp.Star,&datatmp.Camp,&datatmp.CampRate,&datatmp.Exp,&datatmp.Power,&datatmp.UID)
		if err != nil {
			fmt.Printf("scan failed,err:%v\n", err)
		}

		dataslice = append(dataslice,datatmp)
	}
	data[RoleUid] = dataslice
	return data[RoleUid]
}