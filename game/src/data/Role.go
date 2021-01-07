package gamedb

import (
	"github.com/Golangltd/Twlib/DbSt"
	twlib_user "github.com/Golangltd/Twlib/user"
	"github.com/golang/glog"
)

// 获取角色信息
func GetGameRoleInfo(accountId int64) *twlib_user.UserSt {
	RoleData := &twlib_user.UserSt{}
	call := ConnRPC.Go("GameRPC.GetRoleInfo", accountId, &RoleData, nil)
	replyCall := <-call.Done
	glog.Info(call.Error) // 错误处理
	glog.Info(replyCall.Reply)
	return RoleData
}

// 创建角色信息  "role is exit"
func CreateGameRoleInfo(roleInfo *twlib_DbSt.RoleSt) *twlib_user.UserSt {
	var RoleData *twlib_user.UserSt
	call := ConnRPC.Go("GameRPC.CreatRoleInfo", roleInfo, &RoleData, nil)
	replyCall := <-call.Done
	glog.Info(call.Error) // 错误处理
	if RoleData.RoleUid == -1 {
		return RoleData
	}
	glog.Info(replyCall.Reply)
	return RoleData
}
