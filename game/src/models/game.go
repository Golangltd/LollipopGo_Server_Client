package models

import (
	"LollipopGo2.8x/conf/g"
	. "LollipopGo2.8x/cxt"
	gamedb "LollipopGo2.8x/data"
	"LollipopGo2.8x/st"
	"LollipopGo2.8x/tables"
	"LollipopGo2.8x/time/ticker"
	"encoding/json"
	"errors"
	twLibProto "github.com/Golangltd/Twlib/proto"
	twLibRewards "github.com/Golangltd/Twlib/rewards"
	"github.com/Golangltd/Twlib/user"
	twLibUser "github.com/Golangltd/Twlib/user"
	"github.com/Golangltd/cache2go"
	concurrent "github.com/fanliao/go-concurrentMap"
	"github.com/golang/glog"
	"github.com/nsqio/go-nsq"
	"github.com/robfig/cron"
	"golang.org/x/net/websocket"
	"strconv"
)

// 游戏的结构信息
type Game struct {
	Connection    *websocket.Conn
	StrMD5        string
	MapSafe       *concurrent.ConcurrentMap
	UserInfo      *twLibUser.UserSt
	AccountId     int64
	RoundInfo     *st.RoundSt
	OfflineReward []*twLibRewards.RewardSt
	OfflineTime   int64
	SRModel       *SRModel
	DrawCardModel *DrawCardModel
	CacheDB       *cache2go.CacheTable
	UserTicker    map[ticker.TickerUid]*cron.Cron
	Producer      map[string]*nsq.Producer
}

func (m *Game) CloseEOF(closeEvent interface{}) int {
	return 0
}

func (m *Game) HandleCltProtocol(protocol interface{}, protocol2 interface{}, ProtocolData map[string]interface{}, Connection interface{}) interface{} {
//func (m *Game) HandleCltProtocol(protocol interface{}, protocol2 interface{}, ProtocolData map[string]interface{}, Connection *websocket.Conn) interface{} {
	switch protocol {
	case float64(twLibProto.GameDataProto):
		{
			m.HandleCltProtocol2(protocol2, ProtocolData, Connection.(*websocket.Conn))
		}
	default:
		glog.Info("protocol default")
	}
	return 0
}

// 子协议处理
func (m *Game) HandleCltProtocol2(protocol2 interface{}, ProtocolData map[string]interface{}, Connection interface{}) interface{} {
//func (m *Game) HandleCltProtocol2(protocol2 interface{}, ProtocolData map[string]interface{}, Connection *websocket.Conn) interface{} {
	ConnectionData := &Game{
		Connection: Connection.(*websocket.Conn),
		MapSafe:    M,
	}
	_ = ConnectionData
	switch protocol2 {
	default:
		glog.Info("protocol2 default")
	}
	return 0
}

//------------------------------------------------------------------------------
func (m *Game) User_Login(ProtocolData map[string]interface{}) {
	OpenID := ""
	onlineUser := &Game{
		Connection: m.Connection,
		MapSafe:    m.MapSafe,
	}
	m.MapSafe.Put(OpenID+UserKey, onlineUser)
	glog.Info("User_Login")
	data := &Tsts{
		Id: 2,
	}
	m.PlayerSendMessage(data)
}
func (m *Game) Server_Login(ProtocolData map[string]interface{}) {
	onlineUser := &Game{
		Connection: m.Connection,
		MapSafe:    m.MapSafe,
	}
	serverid := ""
	m.MapSafe.Put(serverid+"|Server", onlineUser)
	glog.Info("User_Login")
	data := &Tsts{
		Id: 2,
	}
	m.PlayerSendMessage(data)
}

func (m *Game) PlayerSendMessage(senddata interface{}) int {

	b, err1 := json.Marshal(senddata)
	if err1 != nil {
		glog.Error("PlayerSendMessage json.Marshal _config fail ! err:", err1.Error())
		glog.Flush()
		return 1
	}
	err := websocket.JSON.Send(m.Connection, b)
	if err != nil {
		glog.Error("PlayerSendMessage send _config fail ! err:", err.Error())
		glog.Flush()
		return 2
	}
	return 0
}

// 获取玩家背包装备
func (m *Game) GetUserBagEquip(equipUid int64) (*twLibUser.EquipSt, error) {
	if m.UserInfo == nil {
		return nil, errors.New("GetUserBagEquip, m.UserInfo is nil")
	}
	if m.UserInfo.EquipData == nil {
		return nil, errors.New("GetUserBagEquip, m.UserInfo.EquipData is nil")
	}
	if m.UserInfo.EquipData.EquipSts == nil {
		return nil, errors.New("GetUserBagEquip, m.UserInfo.EquipData.EquipSts is nil")
	}
	for _, equip := range m.UserInfo.EquipData.EquipSts {
		if equip.UID == equipUid {
			return equip, nil
		}
	}
	return nil, errors.New("GetUserBagEquip, cant find bag equip:" + strconv.FormatInt(equipUid, 10))
}

// 移除背包中的装备
func (m *Game) RemoveBagEquips(delMap map[int64]int) (delUIDs []int64, update []*twlib_user.EquipSt, remove []*twlib_user.EquipSt, err error) {
	// TODO: 检查数量
	remove = make([]*twlib_user.EquipSt, 0) // 移除的装备
	delUIDs = make([]int64, 0)              // 数量减为0,需要删除
	update = make([]*twlib_user.EquipSt, 0) // 数量大于0,需要更新
	var bagEquip *twlib_user.EquipSt
	for delUid, delCount := range delMap {
		find := false
		for i := 0; i < len(m.UserInfo.EquipData.EquipSts); i++ {
			bagEquip = m.UserInfo.EquipData.EquipSts[i]
			if bagEquip.UID == delUid {
				bagEquip.Num -= delCount

				if bagEquip.Num < 0 {
					bagEquip.Num = 0
				}
				e := &twlib_user.EquipSt{
					UID:      bagEquip.UID,
					ConfID:   bagEquip.ConfID,
					Star:     bagEquip.Star,
					Camp:     bagEquip.Camp,
					CampRate: bagEquip.CampRate,
					Num:      1,
					Exp:      bagEquip.Exp,
					Power:    0,
				}
				remove = append(remove, e)
				InitEquipPower(e)
				if bagEquip.Num <= 0 {
					// 数量为0，删除
					delUIDs = append(delUIDs, bagEquip.UID)
					m.UserInfo.EquipData.EquipSts = append(m.UserInfo.EquipData.EquipSts[:i], m.UserInfo.EquipData.EquipSts[i+1:]...) // 背包张移除装备
				} else {
					// 数量大于0 更新
					update = append(update, bagEquip)
				}
				find = true
				break
			}
		}
		if !find {
			return nil, nil, nil, errors.New("找不到装备UID:" + strconv.FormatInt(delUid, 10))
		}
	}
	if len(delUIDs) > 0 {
		gamedb.DelEquips(delUIDs)
	}
	if len(update) > 0 {
		for _, e := range update {
			gamedb.UpdateEquip(e)
		}
	}
	return delUIDs, update, remove, nil
}

//// 移除背包中的装备
//func (m *Game) RemoveBagEquip(equipUID int64, count int) (del []int64, update []*twlib_user.EquipSt, remove []*twlib_user.EquipSt, err error) {
//
//	return m.RemoveBagEquips(rmap)
//}

// 添加装备到装备背包(进行堆叠逻辑)
func (m *Game) AddUserBagEquip(addEquips []*twlib_user.EquipSt) (update map[int64]*twlib_user.EquipSt, add []*twlib_user.EquipSt, err error) {
	add = make([]*twlib_user.EquipSt, 0)
addEquip:
	for _, addEquip := range addEquips {
		for _, bagEquip := range m.UserInfo.EquipData.EquipSts {
			if bagEquip.ConfID != addEquip.ConfID {
				continue
			}
			if CanStack(bagEquip, addEquip) { // 是否可以叠加
				newStack := addEquipCount(bagEquip, 1)
				if update == nil {
					update = make(map[int64]*twLibUser.EquipSt)
				}
				update[bagEquip.UID] = bagEquip
				if len(newStack) > 0 {
					m.UserInfo.EquipData.EquipSts = append(m.UserInfo.EquipData.EquipSts, newStack...)
					add = append(add, newStack...)
				}
				continue addEquip
			}
		}
		// 背包中没有找到 根据装备数据，堆叠数量，新增一组装备
		newStack := AddNewStackBagEquip(addEquip, 1)
		m.UserInfo.EquipData.EquipSts = append(m.UserInfo.EquipData.EquipSts, newStack...)
		add = append(add, newStack...)
	}
	// 更新数据库装备
	for _, v := range update {
		gamedb.UpdateEquip(v)
	}
	// 新增数据库装备
	gamedb.CreateEquip(m.UserInfo.RoleUid, m.AccountId, add)
	return update, add, nil
}

// 通过卡牌唯一ID获取玩家卡牌
func (m *Game) GetCardInfo(cardUID int64) (*twLibUser.CardInfo, error) {
	if m.UserInfo == nil {
		return nil, errors.New("GetCardInfo, m.UserInfo is nil")
	}
	if m.UserInfo.CardList == nil {
		return nil, errors.New("GetCardInfo, m.UserInfo.CardList is nil")
	}
	return nil, errors.New("GetCardInfo, cant find cardUID:" + strconv.FormatInt(cardUID, 10))
}

// 通过卡牌唯一ID获取玩家卡牌装备
func (m *Game) GetCardEquip(cardUID int64, equipUID int64) (*twLibUser.EquipSt, error) {
	if m.UserInfo == nil {
		return nil, errors.New("GetCardInfo, m.UserInfo is nil")
	}
	if m.UserInfo.CardList == nil {
		return nil, errors.New("GetCardInfo, m.UserInfo.CardList is nil")
	}

	return nil, errors.New("GetCardInfo, cant find cardUID:" + strconv.FormatInt(cardUID, 10))
}

func (m *Game) GetUserEquipData() (*twLibUser.EquipData, error) {
	if m.UserInfo == nil {
		return nil, errors.New("ReduceGroupUserBagEquip, m.UserInfo is nil")
	}
	if m.UserInfo.EquipData == nil {
		return nil, errors.New("ReduceGroupUserBagEquip, m.UserInfo.EquipData is nil")
	}
	return m.UserInfo.EquipData, nil
}

// 获取玩家道具数量
func (m *Game) GetUserItemNum(itemID int) int64 {
	var itemNum int64 = 0
	if itemID == g.ItemTypeDiamond {
		return m.UserInfo.Diamond
	} else if itemID == g.ItemTypeCoin {
		return m.UserInfo.Coin
	} else if itemID == g.ItemTypeSchollRollExp {

	} else if itemID == g.ItemTypeCardExp {
		return int64(m.UserInfo.RoleExp)
	}
	// TODO:确认所有道具类型
	for _, item := range m.UserInfo.ItemList {
		if item != nil && item.ItemId == itemID {
			itemNum += item.ItemNum
		}
	}
	return itemNum
}

// 初始化装备战力
func InitEquipPower(equip *twLibUser.EquipSt) {
	conf := tables.EquipTables[equip.ConfID]
	if conf == nil {
		return
	}
	power := 0
	tempRate := 0
	for _, attribute := range conf.StarAttributes[equip.Star] {
		tempRate = tables.GAttributePowerRate[attribute.Type]
		if tempRate <= 0 {
			tempRate = 1
		}
		power = power + tempRate*attribute.Val
	}
	equip.Power = power
}

// 增加装备数量
func addEquipCount(oldEquip *twLibUser.EquipSt, count int) (add []*twLibUser.EquipSt) {
	if oldEquip == nil {
		return nil
	}
	conf, ok := tables.EquipTables[oldEquip.ConfID]
	if !ok {
		return nil
	}
	// 叠加原来的装备数量
	left := conf.StackMaxNum - oldEquip.Num // 当前装备剩余叠加数量
	if left >= count {                      // 可堆叠剩余数大于添加数
		oldEquip.Num += count
		return nil
	} else { // 可堆叠剩余数小于添加数量
		oldEquip.Num = conf.StackMaxNum
		count -= left
	}
	newStack := AddNewStackBagEquip(oldEquip, count)
	return newStack
}

// 通知指定装备数据增加 新堆叠的装备
func AddNewStackBagEquip(equip *twLibUser.EquipSt, count int) (add []*twLibUser.EquipSt) {
	conf, ok := tables.EquipTables[equip.ConfID]
	if !ok {
		return nil
	}
	newStack := make([]*twLibUser.EquipSt, 0)
	var n *twLibUser.EquipSt
	for {
		if count <= 0 {
			break
		}
		n = &twLibUser.EquipSt{
			ConfID:   equip.ConfID,
			Star:     equip.Star,
			Camp:     equip.Camp,
			CampRate: equip.CampRate,
			Num:      0,
			Exp:      equip.Exp,
			Power:    0,
		}
		InitEquipPower(n)
		newStack = append(newStack, n)
		if count <= conf.StackMaxNum {
			n.Num = count
			break
		} else {
			n.Num = conf.StackMaxNum
			count -= conf.StackMaxNum
		}

	}
	return newStack
}

// 判断两件装备是否可以叠加
func CanStack(equip1 *twLibUser.EquipSt, equip2 *twLibUser.EquipSt) bool {
	if equip1 == nil || equip2 == nil {
		return false
	}
	// 配置档ID相同, 星级相同, 阵营加成相同，阵营加成值相同, 经验相同
	if equip1.ConfID == equip2.ConfID && equip1.Star == equip2.Star && equip1.Camp == equip2.Camp && equip1.CampRate == equip2.CampRate && equip1.Exp == equip2.Exp {
		return true
	}
	return false
}
