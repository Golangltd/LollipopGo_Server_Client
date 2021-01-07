package comm_proto

import (
	impl "LollipopGo/network"
	"LollipopGo2.8x/conf/g"
	gamedb "LollipopGo2.8x/data"
	util_handlers "LollipopGo2.8x/handlers/util"
	"LollipopGo2.8x/models"
	gamemsg "LollipopGo2.8x/msg"
	"errors"
	"fmt"
	twLibItem "github.com/Golangltd/Twlib/item"
	twProto "github.com/Golangltd/Twlib/proto"
	twlib_user "github.com/Golangltd/Twlib/user"
	"github.com/golang/glog"
	"golang.org/x/net/websocket"
)

type UpdateComm struct {
	RoleUid   int64
	AccountId int64
	Num       int64
}

// 更新金币
func UpdateRoleDiamond(conn *websocket.Conn, stropenid string, num int, userinfo *twlib_user.UserSt) int {

	if num == 0 || userinfo.Coin+int64(num) < 0 {
		// 余额不足
		return -1
	}

	data := &GS2C_UpdateCoin{
		Protocol:  twProto.GGameHallProto,
		Protocol2: GS2C_UpdateCoinProto2,
		CoinNum:   num,
	}
	impl.PlayerSendToProxyServer(conn, data, stropenid)

	// 保存数据库
	savedata := UpdateComm{
		userinfo.RoleUid,
		0,
		int64(num),
	}

	iret := UpdateDBRoleCoinInfo(savedata)
	if *iret == 1 {
		userinfo.Coin += int64(num)
		return 0
	}
	return -2 // 扣费失败
}

// 更新砖石
func UpdateRoleCoin(conn *websocket.Conn, stropenid string, num int, userinfo *twlib_user.UserSt) int {

	if num == 0 || userinfo.Diamond+int64(num) < 0 {
		// 余额不足
		return -1
	}
	data := &GS2C_UpdateDiamond{
		Protocol:   twProto.GGameHallProto,
		Protocol2:  GS2C_UpdateDiamondProto2,
		DiamondNum: num,
	}
	impl.PlayerSendToProxyServer(conn, data, stropenid)

	savedata := UpdateComm{
		userinfo.RoleUid,
		0,
		int64(num),
	}
	iret := UpdateDBRoleDiamondInfo(savedata)
	if *iret == 1 {
		userinfo.Diamond += int64(num)
		return 0
	}
	return -2 // 扣费失败
}

func UpdateRoleItemMap(game *models.Game, conn *websocket.Conn, openID string, checkReduceNum bool, itemMap map[int]int) error {
	// 检查数量是否足够
	if checkReduceNum {
		for k, v := range itemMap {
			if v > 0 {
				continue
			}
			if value, ok := itemMap[g.ItemTypeDiamond]; ok {
				if checkReduceNum && value < 0 && game.UserInfo.Diamond < int64(-value) {
					return errors.New(fmt.Sprintf("diamond not enough,curent:%d, update:%d\n", game.UserInfo.Diamond, value))
				}
			}
			if value, ok := itemMap[g.ItemTypeCoin]; ok {
				if checkReduceNum && value < 0 && game.UserInfo.Coin < int64(-value) {
					return errors.New(fmt.Sprintf("coin not enough,curent:%d, update:%d\n", game.UserInfo.Coin, value))
				}
			}
			// TODO:学籍经验
			if _, ok := itemMap[g.ItemTypeSchollRollExp]; ok {
			}
			// TODO:巫师经验
			if value, ok := itemMap[g.ItemTypeCardExp]; ok {
				if checkReduceNum && value < 0 && game.UserInfo.RoleExp < -value {
					return errors.New(fmt.Sprintf("role exp not enough,curent:%d, update:%d\n", game.UserInfo.Coin, value))
				}
			}
			exist := false
			for _, item := range game.UserInfo.ItemList {
				if item.ItemId == k {
					exist = true
					if item.ItemNum < int64(-v) {
						return errors.New(fmt.Sprintf("item not enough,curent:%d, update:%d\n", item.ItemNum, v))
					}
				}
			}
			if !exist {
				return errors.New(fmt.Sprintf("item not enough,curent:0, update:%d\n", v))
			}
		}
	}

	// 更新
	// 金加隆(钻石 充值货币)
	if value, ok := itemMap[g.ItemTypeDiamond]; ok {
		UpdateRoleCoin(conn, openID, value, game.UserInfo)
		delete(itemMap, g.ItemTypeDiamond)
	}
	// 铜纳特(金币 游戏货币)
	if value, ok := itemMap[g.ItemTypeCoin]; ok {
		UpdateRoleDiamond(conn, openID, value, game.UserInfo)
		delete(itemMap, g.ItemTypeCoin)
	}
	// TODO:学籍经验
	if _, ok := itemMap[g.ItemTypeSchollRollExp]; ok {
		delete(itemMap, g.ItemTypeSchollRollExp)
	}
	// TODO:巫师经验
	if value, ok := itemMap[g.ItemTypeCardExp]; ok {
		game.UserInfo.RoleExp += value
		delete(itemMap, g.ItemTypeCardExp)
	}
	// 集合中剩余为道具背包中的道具
	if len(itemMap) > 0 {
		if err := updateRoleItem(conn, openID, false, game, itemMap); err != nil {
			return err
		}
	}
	return nil
}

func UpdateRoleItem(game *models.Game, conn *websocket.Conn, openID string, checkReduceNum bool, items []*twLibItem.ItemData) error {

	// 添加的道具转换成map
	var itemMap = make(map[int]int) // map[itemID]itemNum1
	for _, item := range items {
		if item.Type != twlib_user.ItemCommon {
			return errors.New("UpdateRoleItem, but item type error")
		}
		if _, ok := itemMap[item.ID]; ok {
			itemMap[item.ID] = itemMap[item.ID] + item.Num
		} else {
			itemMap[item.ID] = item.Num
		}
	}
	return UpdateRoleItemMap(game, conn, openID, checkReduceNum, itemMap)
}

func updateRoleItem(conn *websocket.Conn, openID string, checkReduceNum bool, game *models.Game, itemMap map[int]int) error {
	if checkReduceNum {
		for k, v := range itemMap {
			if v > 0 {
				continue
			}
			for _, item := range game.UserInfo.ItemList {
				if item.ItemId == k && item.ItemNum < int64(-v) {
					return errors.New(fmt.Sprintf("item not enough,curent:%d, update:%d\n", item.ItemNum, v))
				}
			}
		}
	}
	var updateItem *twlib_user.ItemData
	updateItems := make([]*twlib_user.ItemData, 0)
	for k, v := range itemMap {
		exist := false
		for _, item := range game.UserInfo.ItemList {
			if item != nil && item.ItemId == k {
				exist = true
				item.ItemNum += int64(v)
				if item.ItemNum < 0 {
					v = int(-item.ItemNum)
					item.ItemNum = 0
				}
				updateItem = &twlib_user.ItemData{
					ItemUid:  item.ItemUid,
					ItemId:   item.ItemId,
					ItemType: item.ItemType,
					ItemNum:  int64(v),
				}
				break
			}
		}
		if !exist && v > 0 {
			updateItem = &twlib_user.ItemData{
				ItemId:   k,
				ItemType: twlib_user.ItemCommon,
				ItemNum:  int64(v),
			}
			game.UserInfo.ItemList = append(game.UserInfo.ItemList, updateItem)
		}
		gamedb.SaveItem(game.UserInfo.RoleUid, game.AccountId, updateItem)
		updateItems = append(updateItems, updateItem)
	}
	data := &GS2C_UpDataItem41Data{
		Protocol:  twProto.GGameHallProto,
		Protocol2: GS2C_UpDataItem41DataProto2,
		ItemList:  updateItems,
	}
	impl.PlayerSendToProxyServer(conn, data, openID)
	return nil
}

// 更新装备
func UpdateEquip(conn *websocket.Conn, stropenid string, iteminfo []*twlib_user.EquipSt) {
	game, _, _ := util_handlers.GetGameAndUser(stropenid)
	update, add, _ := game.AddUserBagEquip(iteminfo)
	updateBagEquipMsg := gamemsg.NewGS2CBagEquipUpdateMsg()
	if len(update) > 0 {
		updateBagEquipMsg.Updates = make([]*twlib_user.EquipSt, 0)
		for _, v := range update {
			updateBagEquipMsg.Updates = append(updateBagEquipMsg.Updates, v)
		}
	}
	if len(add) > 0 {
		updateBagEquipMsg.Adds = add
	}
	impl.PlayerSendToProxyServer(conn, updateBagEquipMsg, stropenid)
}

//---save db
func UpdateDBRoleCoinInfo(data UpdateComm) *int {
	iint := new(int)
	call := gamedb.ConnRPC.Go("GameRPC.UpdateRoleCoinComm", data, &iint, nil)
	replyCall := <-call.Done
	glog.Info(call.Error) // 错误处理
	glog.Info(replyCall.Reply)
	return iint
}

func UpdateDBRoleDiamondInfo(data UpdateComm) *int {
	iint := new(int)
	call := gamedb.ConnRPC.Go("GameRPC.UpdateRoleDiamondComm", data, &iint, nil)
	replyCall := <-call.Done
	glog.Info(call.Error) // 错误处理
	glog.Info(replyCall.Reply)
	return iint
}

// 卡牌更新消息(包含更新，添加，删除)
func UpdateCard(conn *websocket.Conn, openID string, update []*twlib_user.CardInfo, add []*twlib_user.CardInfo, del []uint64) {
	msg := NewGS2CUpdateCardMsg()
	msg.OpenId = openID
	if len(update) > 0 {
		msg.Updates = update
	}
	if len(add) > 0 {
		msg.Adds = add
	}
	if len(del) > 0 {
		msg.DelUIDs = del
	}
	impl.PlayerSendToProxyServer(conn, msg, openID)
}
