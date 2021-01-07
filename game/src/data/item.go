package gamedb

import (
	twlib_user "github.com/Golangltd/Twlib/user"
	"github.com/golang/glog"
)

type SItem struct {
	RoleUID  int64
	Account  int64
	ItemInfo *twlib_user.ItemData
}

// 保存道具
func SaveItem(roleuid int64, account int64, itemData *twlib_user.ItemData) {
	data := &SItem{
		roleuid,
		account,
		itemData,
	}
	var result = new(bool)
	call := ConnRPC.Go("GameRPC.CreatItemInfo", data, result, nil)
	glog.Infof("result:%t\n", *result)
	replyCall := <-call.Done
	if call.Error != nil {
		glog.Errorf("db rpc err:%s\n", call.Error.Error())
	}
	glog.Info(replyCall.Reply)
}

func UpdateItem(roleuid int64, itemData *twlib_user.ItemSt) {

}

func DelItem(roleuid int64, itemData *twlib_user.ItemSt) {

}

func GetItem(roleuid int64, itemData *twlib_user.ItemSt) {

}
