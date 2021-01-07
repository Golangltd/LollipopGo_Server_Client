package equip

import (
	impl "LollipopGo/network"
	"LollipopGo2.8x/conf/g"
	gamedb "LollipopGo2.8x/data"
	LogicCommon "LollipopGo2.8x/handlers/common"
	"LollipopGo2.8x/handlers/modules"
	util_handlers "LollipopGo2.8x/handlers/util"
	"LollipopGo2.8x/models"
	"LollipopGo2.8x/msg"
	"LollipopGo2.8x/proto/comm_proto"
	ProtoGame "LollipopGo2.8x/proto/tw_proto"
	"LollipopGo2.8x/tables"
	"errors"
	"fmt"
	TWLibItem "github.com/Golangltd/Twlib/item"
	twLibUser "github.com/Golangltd/Twlib/user"
	"github.com/golang/glog"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/net/websocket"
	"strconv"
)

// 注册handler
func InitEquipHandlers() {
	modules.HM.AddHandler(ProtoGame.C2GSEquipWearProto2, Wear)
	modules.HM.AddHandler(ProtoGame.C2GSEquipTakeOff, TakeOff)
	modules.HM.AddHandler(ProtoGame.C2GSEquipReplace, Replace)
	modules.HM.AddHandler(ProtoGame.C2GSEquipStrengthen, Strengthen)
}

/* 玩家穿戴/一键穿戴装备
将背包中的装备穿戴到卡牌身上
*/
func Wear(conn *websocket.Conn, ProtocolData map[string]interface{}) {
	var reqStrMsg gamemsg.C2GSEquipWearStrMsg
	if err := mapstructure.Decode(ProtocolData, &reqStrMsg); err != nil {
		panic(err)
	}
	reqMsg := &gamemsg.C2GSEquipWearMsg{
		Protocol:  reqStrMsg.Protocol,
		Protocol2: reqStrMsg.Protocol2,
		OpenId:    reqStrMsg.OpenId,
		CardUID:   reqStrMsg.CardUID,
	}
	uids := make([]int64, 0)
	for _, equipUID := range reqStrMsg.ItemUIDs {
		t, _ := strconv.ParseInt(equipUID, 10, 64)
		uids = append(uids, t)
	}
	reqMsg.ItemUIDs = uids

	// 获取玩家
	game, _, err := util_handlers.GetGameAndUser(reqMsg.OpenId)
	if err != nil {
		// TODO: 提示客户端重新登录
		LogicCommon.SendError(conn, reqMsg.OpenId, "0", "can find user", 1)
		panic(err)
	}

	// 获取卡牌
	card, err := game.GetCardInfo(reqMsg.CardUID)
	if err != nil {
		LogicCommon.SendError(conn, reqMsg.OpenId, "0", err.Error(), 1)
		panic(err)
	}

	cardConf := tables.RolesTable[int(card.CardID)]
	// 检测装备与卡牌性别匹配
	for _, equipUID := range reqMsg.ItemUIDs {
		e, err := game.GetUserBagEquip(equipUID)
		if err != nil {
			LogicCommon.SendError(conn, reqMsg.OpenId, "0", err.Error(), 1)
			panic(err)
		}
		equipConf, _ := tables.EquipTables[e.ConfID]
		if equipConf.UseSex != cardConf.Sex && equipConf.UseSex != g.SexTypeCommon {
			LogicCommon.SendError(conn, reqMsg.OpenId, "0", "equip sex error", 1)
			panic(errors.New("equip sex not match"))
		}
	}
	glog.Infof("穿戴前 装备背包数量:%d", GetBagEquipSize(game))
	glog.Infof("穿戴前 卡牌装备数量:%d", GetCardEquipSize(card))
	// 从背包中移除装备
	m := make(map[int64]int)
	for _, uid := range reqMsg.ItemUIDs {
		m[uid] = 1
	}
	delEquipUIDs, update, remove, err := game.RemoveBagEquips(m)
	if err != nil {
		panic(err)
	}

	// 执行穿戴
	addCardEquips := make([]*twLibUser.EquipSt, 6) // 新穿戴的装备

	var equipConf *tables.EquipConfSt // 配置档
	var addToBag []*twLibUser.EquipSt // 添加到背包的装备
	for _, equip := range remove {
		equipConf, _ = tables.EquipTables[equip.ConfID]
		if card.Equips[equipConf.Position-1] != nil { // 装备栏 有装备了
			if addToBag == nil {
				addToBag = make([]*twLibUser.EquipSt, 0)
			}
			addToBag = append(addToBag, card.Equips[equipConf.Position-1])
		}
		card.Equips[equipConf.Position-1] = equip // 卡牌穿戴装备
		glog.Infof("穿戴前 装备:%d, 战斗力:%d", equip.UID, equip.Power)
		InitCardEquipPower(card, equip) // 重新计算装备战斗力
		glog.Infof("穿戴后 装备:%d, 战斗力:%d", equip.UID, equip.Power)
		addCardEquips[equipConf.Position-1] = equip
	}
	// 添加到装备背包
	updates, adds, err := game.AddUserBagEquip(addToBag)
	if err != nil {
		LogicCommon.SendError(conn, reqMsg.OpenId, "0", err.Error(), 1)
		panic(err)
	}
	for _, u := range update {
		if updates[u.UID] == nil {
			updates = make(map[int64]*twLibUser.EquipSt)
			updates[u.UID] = u
		}
	}
	glog.Infof("穿戴后 装备背包数量:%d\n", GetBagEquipSize(game))
	glog.Infof("穿戴后 卡牌装备数量:%d\n", GetCardEquipSize(card))

	// 装备背包  更新消息
	updateBagEquipMsg := gamemsg.NewGS2CBagEquipUpdateMsg()
	if len(delEquipUIDs) > 0 { // 删除
		glog.Info("穿戴后 背包删除ID:", delEquipUIDs)
		updateBagEquipMsg.DelUIDs = delEquipUIDs
	}
	if len(update) > 0 { //更新
		updateBagEquipMsg.Updates = update
		for _, d := range update {
			glog.Infof("穿戴后 背包更新:%+v\n", d)
		}
	}
	if len(adds) > 0 {
		updateBagEquipMsg.Adds = adds
	}
	impl.PlayerSendToProxyServer(conn, updateBagEquipMsg, reqMsg.OpenId)

	// 卡牌装备更新
	cardEquipUpdateMsg := gamemsg.NewGS2CCardEquipUpdate()
	cardEquipSt := &gamemsg.CardEquipSt{
		Equips:  addCardEquips,
	}
	cardEquipUpdateMsg.CardsEquips = []*gamemsg.CardEquipSt{cardEquipSt}
	impl.PlayerSendToProxyServer(conn, cardEquipUpdateMsg, reqMsg.OpenId)
	// 更新数据
	gamedb.UpdateUserCardInfo(card)
}

/* 玩家脱下/一键脱下装备
从卡牌身上脱下装备
*/
func TakeOff(conn *websocket.Conn, ProtocolData map[string]interface{}) {
	var reqMsg gamemsg.C2GSEquipTakeOffMsg
	if err := mapstructure.Decode(ProtocolData, &reqMsg); err != nil {
		panic(err)
	}
	// 获取玩家
	game, _, err := util_handlers.GetGameAndUser(reqMsg.OpenId)
	if err != nil {
		// TODO: 提示客户端重新登录
		LogicCommon.SendError(conn, reqMsg.OpenId, "0", "user is nil", 1)
		panic(err)
	}
	// 获取卡牌
	card, err := game.GetCardInfo(reqMsg.CardUID)
	if err != nil {
		LogicCommon.SendError(conn, reqMsg.OpenId, "0", "card is nil", 1)
		panic(err)
	}

	glog.Infof("脱下前 装备背包数量:%d", GetBagEquipSize(game))
	glog.Infof("脱下前 卡牌装备数量:%d", GetCardEquipSize(card))

	addToBag := make([]*twLibUser.EquipSt, 0) // 记录脱下的装备
	delPos := make([]int, 0)
	for _, pos := range reqMsg.Positions {
		if card.Equips[pos-1] == nil {
			LogicCommon.SendError(conn, reqMsg.OpenId, "0", "no equip, pos:"+strconv.Itoa(pos), 1)
			panic(errors.New("no equip, pos:" + strconv.Itoa(pos)))
		}
		addToBag = append(addToBag, card.Equips[pos-1])
		glog.Infof("脱下前 装备:%d, 战斗力:%d", card.Equips[pos-1].UID, card.Equips[pos-1].Power)
		models.InitEquipPower(card.Equips[pos-1]) // 重新计算装备战斗力
		glog.Infof("脱下后 装备:%d, 战斗力:%d", card.Equips[pos-1].UID, card.Equips[pos-1].Power)
		card.Equips[pos-1] = nil // 卡牌装备置空
		delPos = append(delPos, pos)
	}
	// 添加到装备背包
	update, add, err := game.AddUserBagEquip(addToBag)
	if err != nil {
		LogicCommon.SendError(conn, reqMsg.OpenId, "0", err.Error(), 1)
		panic(err)
	}
	glog.Infof("脱下后 装备背包数量:%d", GetBagEquipSize(game))
	glog.Infof("脱下后 卡牌装备数量:%d", GetCardEquipSize(card))

	// 背包装备更新消息
	updateBagEquipMsg := gamemsg.NewGS2CBagEquipUpdateMsg()
	if len(update) > 0 {
		updateBagEquipMsg.Updates = make([]*twLibUser.EquipSt, 0)
		for _, v := range update {
			updateBagEquipMsg.Updates = append(updateBagEquipMsg.Updates, v)
		}
	}
	if len(add) > 0 {
		updateBagEquipMsg.Adds = add
	}
	impl.PlayerSendToProxyServer(conn, updateBagEquipMsg, reqMsg.OpenId)
	// 卡牌装备 删除消息
	delCardEquipMsg := gamemsg.NewGS2CCardEquipDelMsg()
	delCardEquipMsg.CardUID = reqMsg.CardUID
	delCardEquipMsg.Positions = delPos
	impl.PlayerSendToProxyServer(conn, delCardEquipMsg, reqMsg.OpenId)
	// 更新数据
	gamedb.UpdateUserCardInfo(card)
}

/* 替换装备
交换两张卡牌的装备
*/
func Replace(conn *websocket.Conn, ProtocolData map[string]interface{}) {
	var reqMsg gamemsg.C2GSEquipReplaceMsg
	if err := mapstructure.Decode(ProtocolData, &reqMsg); err != nil {
		panic(err)
	}
	// 获取玩家
	game, _, err := util_handlers.GetGameAndUser(reqMsg.OpenId)
	if err != nil {
		// TODO: 提示客户端重新登录
		LogicCommon.SendError(conn, reqMsg.OpenId, "0", "user is nil", 1)
		panic(err)
	}
	// 获取卡牌
	card, err := game.GetCardInfo(reqMsg.CardUID)
	if err != nil {
		LogicCommon.SendError(conn, reqMsg.OpenId, "0", "card is nil", 1)
		panic(err)
	}
	// 获取原卡牌
	srcCard, err := game.GetCardInfo(reqMsg.SrcCardUID)
	if err != nil {
		LogicCommon.SendError(conn, reqMsg.OpenId, "0", "srcCard is nil", 1)
		panic(err)
	}
	position := reqMsg.Position - 1
	srcPosition := reqMsg.SrcPosition - 1
	if card.Equips[position] == nil {
		LogicCommon.SendError(conn, reqMsg.OpenId, "0", "position is not nil", 1)
		panic(errors.New("position is not nil"))
	}
	if srcCard.Equips[srcPosition] == nil {
		LogicCommon.SendError(conn, reqMsg.OpenId, "0", "srcPosition is not nil", 1)
		panic(errors.New("srcPosition is not nil"))
	}
	// 交换装备
	card.Equips[position], srcCard.Equips[srcPosition] = srcCard.Equips[srcPosition], card.Equips[position]
	glog.Infof("交换前 装备:%d, 战斗力:%d", card.Equips[position].UID, card.Equips[position].Power)
	glog.Infof("交换前 装备:%d, 战斗力:%d", srcCard.Equips[srcPosition].UID, srcCard.Equips[srcPosition].Power)
	InitCardEquipPower(card, card.Equips[position])          // 重新计算卡牌装备战斗力
	InitCardEquipPower(srcCard, srcCard.Equips[srcPosition]) // 重新计算卡牌装备战斗力
	glog.Infof("交换后 装备:%d, 战斗力:%d", card.Equips[position].UID, card.Equips[position].Power)
	glog.Infof("交换后 装备:%d, 战斗力:%d", srcCard.Equips[srcPosition].UID, srcCard.Equips[srcPosition].Power)
	// 卡牌装备更新消息

	cardEquipUpdateMsg := gamemsg.NewGS2CCardEquipUpdate()
	cardsEquips := &gamemsg.CardEquipSt{
		Equips:  []*twLibUser.EquipSt{card.Equips[position]},
	}
	srcCardsEquips := &gamemsg.CardEquipSt{
		Equips:  []*twLibUser.EquipSt{srcCard.Equips[srcPosition]},
	}
	cardEquipUpdateMsg.CardsEquips = []*gamemsg.CardEquipSt{cardsEquips, srcCardsEquips}
	impl.PlayerSendToProxyServer(conn, cardEquipUpdateMsg, reqMsg.OpenId)

}

/* 玩家强化装备
强化背包中的装备或者卡牌身上的装备，通过请求消息中的CardID判断
*/
func Strengthen(conn *websocket.Conn, ProtocolData map[string]interface{}) {
	var reqStrMsg gamemsg.C2GSEquipStrengthenStrMsg
	if err := mapstructure.Decode(ProtocolData, &reqStrMsg); err != nil {
		panic(err)
	}

	reqMsg := &gamemsg.C2GSEquipStrengthenMsg{
		Protocol:      reqStrMsg.Protocol,
		Protocol2:     reqStrMsg.Protocol2,
		OpenId:        reqStrMsg.OpenId,
		CardUID:       reqStrMsg.CardUID,
		CostMaterials: reqStrMsg.CostMaterials,
		CostEquips:    nil,
	}
	equips := make([]*twLibUser.ItemData, 0)
	for _, e := range reqStrMsg.CostEquips {
		euid, _ := strconv.ParseInt(e.ItemUid, 10, 64)
		equips = append(equips, &twLibUser.ItemData{
			ItemUid:  euid,
			ItemId:   e.ItemId,
			ItemType: e.ItemType,
			ItemNum:  e.ItemNum,
		})
	}
	reqMsg.CostEquips = equips

	equipUID, err := strconv.ParseInt(reqStrMsg.EquipUID, 10, 64)
	reqMsg.EquipUID = equipUID
	if len(reqMsg.CostMaterials) < 1 && len(reqMsg.CostEquips) < 1 {
		LogicCommon.SendError(conn, reqMsg.OpenId, "0", "cost equip and material is nil", 1)
		panic(err)
	}

	// 获取玩家
	game, _, err := util_handlers.GetGameAndUser(reqMsg.OpenId)
	if err != nil {
		LogicCommon.SendError(conn, reqMsg.OpenId, "0", "user is nil", 1)
		panic(err)
	}

	var equip *twLibUser.EquipSt
	if reqMsg.CardUID <= 0 { // 背包装备
		equip, err = game.GetUserBagEquip(reqMsg.EquipUID)
		if err != nil {
			LogicCommon.SendError(conn, reqMsg.OpenId, "0", "bag equip is nil", 1)
			panic(err)
		}

	} else { // 卡牌装备
		equip, err = game.GetCardEquip(reqMsg.CardUID, reqMsg.EquipUID)
		if err != nil {
			LogicCommon.SendError(conn, reqMsg.OpenId, "0", "card equip is nil", 1)
			panic(err)
		}
	}
	equipConf, ok := tables.EquipTables[equip.ConfID]
	if !ok {
		LogicCommon.SendError(conn, reqMsg.OpenId, "0", "equip conf error", 1)
		panic(errors.New("cant find equip config:" + strconv.Itoa(equip.ConfID)))
	}

	if equip.Star >= equipConf.MaxStar { // 达到最大星级
		LogicCommon.SendError(conn, reqMsg.OpenId, "0", "had max star", 1)
		panic(errors.New("Strengthen, had max star:" + strconv.Itoa(equip.Star)))
	}
	costItems := make([]*TWLibItem.ItemData, 0)
	materialTotalExp := 0 // 所有强化材料提供的经验
	if len(reqMsg.CostMaterials) > 0 {
		// 检查强化材料是否足够
		for _, item := range reqMsg.CostMaterials {
			if game.GetUserItemNum(item.ItemId) < item.ItemNum {
				LogicCommon.SendError(conn, reqMsg.OpenId, "0", "item not enough", 1)
				panic(errors.New(fmt.Sprintf("Strengthen, item not enough:%d", item.ItemId)))
			}

		}
		// 检测强化材料所需的资源是否足够
		var itemConf *tables.ItemSt
		for _, item := range reqMsg.CostMaterials {
			// 获取材料的消耗资源
			itemConf = tables.ItemTables[item.ItemId]
			if itemConf == nil {
				LogicCommon.SendError(conn, reqMsg.OpenId, "0", "no such item config", 1)
				panic(errors.New(fmt.Sprintf("Strengthen, no such item config:%d", item.ItemId)))
			}
			// 添加强化材料消耗
			parameter := itemConf.Parameter.(*tables.ItemTypeEquipStrengthenMaterialParameter)
			costItems = append(costItems, &TWLibItem.ItemData{
				Type: parameter.CostItemType,
				ID:   parameter.CostItemID,
				Num:  -parameter.CostItemNum,
			})
			materialTotalExp += parameter.Exp * int(item.ItemNum)

			// 添加强化材料
			costItems = append(costItems, &TWLibItem.ItemData{
				Type: item.ItemType,
				ID:   item.ItemId,
				Num:  -int(item.ItemNum),
			})
		}
	}

	// 检测装备数量是否足够并返回作为材料时的消耗
	equipCostItems := make([]*TWLibItem.ItemData, 0)
	if len(reqMsg.CostEquips) > 0 {
		_, equipCostItems, err = checkCountAndGetEquip(game.UserInfo.EquipData, reqMsg.CostEquips)
		if err != nil {
			LogicCommon.SendError(conn, reqMsg.OpenId, "0", err.Error(), 1)
			panic(err)
		}
	}

	// 合并装备的消耗资源
	costItems = append(costItems, equipCostItems...)

	// 对玩家强化材料，及升星的消耗进行扣除
	if err = comm_proto.UpdateRoleItem(game, conn, reqMsg.OpenId, true, costItems); err != nil {
		LogicCommon.SendError(conn, reqMsg.OpenId, "0", err.Error(), 1)
		panic(err)
	}
	if materialTotalExp > 0 {
		equip.Exp += materialTotalExp
		for {
			if equip.Star >= equipConf.MaxStar { // 检测星级
				equip.Exp = 0
				break
			}
			if equip.Exp >= equipConf.StarUpExp[equip.Star] { // 超过升星经验
				equip.Exp = equip.Exp - equipConf.StarUpExp[equip.Star] // 扣除升星经验
				equip.Star = equip.Star + 1                             // 增加星级
			} else {
				break
			}
		}
	}

	delBagUIDs := make([]int64, 0)                   // 删除的背包装备
	updateBagEquips := make([]*twLibUser.EquipSt, 0) // 更新的背包装备
	var materialConf *tables.EquipConfSt             // 材料临时配置档

materialForLabel:
	for _, item := range reqMsg.CostEquips { // 请求材料
		if equip.Star >= equipConf.MaxStar { // 最大星级
			break
		}
		material, err := game.GetUserBagEquip(item.ItemUid) // 材料
		if err != nil {
			continue
		}
		materialConf, ok = tables.EquipTables[material.ConfID] // 材料配置档
		if !ok {
			LogicCommon.SendError(conn, reqMsg.OpenId, "0", "cant find equip config:"+strconv.Itoa(equip.ConfID), 1)
			panic(errors.New("cant find equip config:" + strconv.Itoa(material.ConfID)))
		}

		for { // 循环材料个数
			if item.ItemNum <= 0 {
				break
			}
			item.ItemNum = item.ItemNum - 1 // 扣除一个材料
			// 扣除背包材料
			remove := make(map[int64]int)
			remove[item.ItemUid] = 1
			del, update, _, err := game.RemoveBagEquips(remove)
			if err != nil {
				LogicCommon.SendError(conn, reqMsg.OpenId, "0", err.Error(), 1)
				panic(err)
			}
			if len(del) > 0 {
				delBagUIDs = append(delBagUIDs, del...)
			}
			if len(update) > 0 {
				updateBagEquips = append(updateBagEquips, update...)
			}
			// 装备增加经验
			equip.Exp = equip.Exp + materialConf.ConvertExp[material.Star] // 加上 材料星级对应转换的经验
			// 检测升星
			for {
				if equip.Star >= equipConf.MaxStar { // 检测星级
					equip.Exp = 0
					break materialForLabel
				}
				if equip.Exp >= equipConf.StarUpExp[equip.Star] { // 超过升星经验
					equip.Exp = equip.Exp - equipConf.StarUpExp[equip.Star] // 扣除升星经验
					equip.Star = equip.Star + 1                             // 增加星级
				} else {
					break
				}
			}
		}
	}

	updateBagEquipMsg := gamemsg.NewGS2CBagEquipUpdateMsg()
	if reqMsg.CardUID > 0 { // 卡片装备
		card, _ := game.GetCardInfo(reqMsg.CardUID)
		glog.Infof("强化前 装备:%d, 战斗力:%d", equip.UID, equip.Power)
		InitCardEquipPower(card, equip) // 重新计算卡牌装备战斗力
		glog.Infof("强化后 装备:%d, 战斗力:%d", equip.UID, equip.Power)
		// 卡牌装备 更新消息
		cardEquipUpdateMsg := gamemsg.NewGS2CCardEquipUpdate()
		updateEquips := make([]*twLibUser.EquipSt, 6)
		updateEquips[equipConf.Position-1] = equip
		cardsEquips := &gamemsg.CardEquipSt{
			CardUID: reqMsg.CardUID,
			Equips:  updateEquips,
		}
		cardEquipUpdateMsg.CardsEquips = []*gamemsg.CardEquipSt{cardsEquips}
		impl.PlayerSendToProxyServer(conn, cardEquipUpdateMsg, reqMsg.OpenId)
	} else {
		models.InitEquipPower(equip) // 计算背包装备战斗力
		//  背包装备
		updateBagEquips = append(updateBagEquips, equip)
	}

	//删除背包装备消息
	if len(delBagUIDs) > 0 {
		updateBagEquipMsg.DelUIDs = delBagUIDs
	}
	// 更新背包装备消息
	if len(updateBagEquips) > 0 {
		updateBagEquipMsg.Updates = append(updateBagEquipMsg.Updates, updateBagEquips...)
	}
	impl.PlayerSendToProxyServer(conn, updateBagEquipMsg, reqMsg.OpenId)
}

/* 获取背包装备数量（有装备的格子）
堆叠在一起的装备，算作1个装备。不统计堆叠的数量
*/
func GetBagEquipSize(game *models.Game) int {
	if game.UserInfo == nil {
		fmt.Println("PrintBagEquipSize, UserInfo is nil")
		return -1
	}
	if game.UserInfo.EquipData == nil {
		fmt.Println("PrintBagEquipSize, UserInfo.EquipData is nil")
		return -1
	}
	if game.UserInfo.EquipData.EquipSts == nil {
		fmt.Println("PrintBagEquipSize, UserInfo.EquipData.EquipSts is nil")
		return -1
	}
	return len(game.UserInfo.EquipData.EquipSts)
}

// 获取卡牌的装备数量
func GetCardEquipSize(card *twLibUser.CardInfo) int {
	count := 0
	for _, e := range card.Equips {
		if e != nil {
			count++
		}
	}
	return count
}

// 初始化卡牌身上的装备的战力
func InitCardEquipPower(card *twLibUser.CardInfo, equip *twLibUser.EquipSt) {
	equipConf := tables.EquipTables[equip.ConfID]
	if equipConf == nil {
		return
	}
	// 基础属性战力
	power := 0
	tempRate := 0
	for _, attribute := range equipConf.StarAttributes[0] {
		tempRate = tables.GAttributePowerRate[attribute.Type]
		if tempRate <= 0 {
			tempRate = 1
		}
		power = power + tempRate*attribute.Val
	}
	// 星级属性战力
	for _, attribute := range equipConf.StarAttributes[equip.Star] {
		tempRate = tables.GAttributePowerRate[attribute.Type]
		if tempRate <= 0 {
			tempRate = 1
		}
		power = power + tempRate*attribute.Val
	}
	// 阵营加成属性战斗力
	cardConf := tables.RolesTable[int(card.CardID)]
	if cardConf.Camp == equip.Camp {
		for _, attribute := range equipConf.StarAttributes[0] { // 基础属性
			tempRate = tables.GAttributePowerRate[attribute.Type]
			if tempRate <= 0 {
				tempRate = 1
			}
			power = power + tempRate*(int(float64(attribute.Val*tables.GCampAttributeAddRate)/float64(10000)))
		}
	}
	equip.Power = power
}

/* 检测数量是否足够, 并获取相应数量的背包中的装备,及作为升星材料时的消耗货币数量
返回:指定数量的装备(UID为0), 指定数量的装备作为材料时的消耗货币数量，失败错误
*/
func checkCountAndGetEquip(equipData *twLibUser.EquipData, items []*twLibUser.ItemData) (retEquips []*twLibUser.EquipSt, cost []*TWLibItem.ItemData, err error) {
	if items == nil || equipData == nil || equipData.EquipSts == nil || len(equipData.EquipSts) <= 0 {
		return nil, nil, errors.New("equipData or items is nil")
	}
	var tempConfig *tables.EquipConfSt
	for _, item := range items {
		find := false
		for _, equip := range equipData.EquipSts {
			if equip.UID == item.ItemUid {
				if equip.Num < int(item.ItemNum) {
					return nil, nil, errors.New("not enough UID:" + strconv.FormatInt(equip.UID, 10))
				}
				tempConfig = tables.EquipTables[equip.ConfID]
				if tempConfig == nil {
					return nil, nil, errors.New("cant find config, configID:" + strconv.Itoa(equip.ConfID))
				}
				if cost == nil {
					cost = make([]*TWLibItem.ItemData, 0)
				}
				cost = append(cost, tempConfig.AsMaterialCost...)
				if retEquips == nil {
					retEquips = make([]*twLibUser.EquipSt, 0)
				}
				e := &twLibUser.EquipSt{
					UID:      0,
					ConfID:   equip.ConfID,
					Star:     equip.Star,
					Camp:     equip.Camp,
					CampRate: equip.CampRate,
					Num:      int(item.ItemNum),
					Exp:      equip.Exp,
					Power:    equip.Power,
				}
				retEquips = append(retEquips, e)
				find = true
			}
		}
		if !find {
			return nil, nil, errors.New("can find UID:" + strconv.FormatInt(item.ItemUid, 10))
		}
	}
	return retEquips, cost, nil
}
