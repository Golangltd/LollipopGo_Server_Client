package gamedb

import (
	twLibUser "github.com/Golangltd/Twlib/user"
	"github.com/golang/glog"
)

type UpdateEquipData struct {
	RoleUID int64
	Account int64
	Equip   *twLibUser.EquipSt
}

type CreateEquipsData struct {
	RoleUID int64
	Account int64
	Equips  []*twLibUser.EquipSt
}

// 创建装备
func CreateEquip(roleUID int64, accountID int64, equips []*twLibUser.EquipSt) {
	glog.Infof("CreateEquip num:%d\n", len(equips))
	if len(equips) < 1 {
		return
	}
	go func() {
		var createCount = 0
		data := &CreateEquipsData{
			RoleUID: roleUID,
			Account: accountID,
			Equips:  equips,
		}
		call := ConnRPC.Go("GameRPC.CreateEquips", data, &createCount, nil)
		replyCall := <-call.Done
		if replyCall.Error != nil {
			glog.Errorf("rpc CreateEquip err:%s\n", replyCall.Error.Error())
			return
		}
		glog.Infof("rpc create equip count:%d\n", createCount)
	}()
}

// 更新装备
func UpdateEquip(equip *twLibUser.EquipSt) {
	glog.Infof("UpdateEquip equip:%+v\n", equip)
	if equip == nil {
		return
	}
	go func() {
		var updateCount = 0
		call := ConnRPC.Go("GameRPC.UpdateEquips", equip, &updateCount, nil)
		replyCall := <-call.Done
		if replyCall.Error != nil {
			glog.Errorf("rpc UpdateEquip err:%s", replyCall.Error.Error())
			return
		}
		glog.Infof("rpc update equip count:%d\n", updateCount)
	}()
}

func DelEquips(equipUIDs []int64) {
	glog.Infof("DelEquips num:%d\n", len(equipUIDs))
	if len(equipUIDs) < 1 {
		return
	}
	go func() {
		var delCount = 0
		call := ConnRPC.Go("GameRPC.DelEquips", equipUIDs, &delCount, nil)
		replyCall := <-call.Done
		if replyCall.Error != nil {
			glog.Errorf("rpc DelEquips err:%s", replyCall.Error.Error())
			return
		}
		glog.Infof("rpc del equip count:%d\n", delCount)
	}()
}
