package game

import (
	"fmt"
)

// 公用结构
type UpdateComm struct {
	RoleUid int64
	AccountId int64
	Num int64
}

// 更新金币
func (m *GameRPC)UpdateRoleCoinComm(data UpdateComm,replydata *int)  error {

	sqlStr := `update gl_role set coin=coin + ? where roleuid = ?`
	ret, err := GetGameConn().Exec(sqlStr, data.Num, data.RoleUid)
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
	*replydata = 1
	return nil
}

// 更新砖石
func (m *GameRPC)UpdateRoleDiamondComm(data UpdateComm,replydata *int)  error {
	sqlStr := `update gl_role set diamond=diamond+? where roleuid = ?`
	ret, err := GetGameConn().Exec(sqlStr, data.Num, data.RoleUid)
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
    *replydata = 1
	return nil
}