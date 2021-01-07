package gamedb

import (
	twlib_user "github.com/Golangltd/Twlib/user"
	"github.com/golang/glog"
)

type CreateCardData struct {
	RoleUID int64
	Account int64
	Cards   []*twlib_user.CardInfo
}

// 创建卡牌
func CreateUserCardInfo(roleUID int64, accountID int64, cards []*twlib_user.CardInfo) {
	if len(cards) < 1 {
		return
	}
	glog.Infof("CreateUserCardInfo num:%d\n", len(cards))
	go func() {
		var createCount = 0
		data := &CreateCardData{
			RoleUID: roleUID,
			Account: accountID,
			Cards:   cards,
		}
		call := ConnRPC.Go("GameRPC.CreatCards", data, &createCount, nil)
		replyCall := <-call.Done
		if replyCall.Error != nil {
			glog.Errorf("rpc CreatCards err:%s\n", replyCall.Error.Error())
			return
		}
		glog.Infof("rpc CreatCards count:%d\n", createCount)
	}()
}

// 更新卡牌
func UpdateUserCardInfo(card *twlib_user.CardInfo) {
	if card == nil {
		return
	}
	glog.Infof("UpdateUserCardInfo card:%+v\n", card)
	go func() {
		var createCount = 0
		call := ConnRPC.Go("GameRPC.UpdateCard", card, &createCount, nil)
		replyCall := <-call.Done
		if replyCall.Error != nil {
			glog.Errorf("rpc UpdateCards err:%s\n", replyCall.Error.Error())
			return
		}
		glog.Infof("rpc UpdateCards count:%d\n", createCount)
	}()
}
