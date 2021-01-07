package drawcard

import (
	"LollipopGo/log"
	. "LollipopGo2.8x/conf/g"
	. "LollipopGo2.8x/proto/dc_proto"
	"LollipopGo2.8x/tables"
	. "github.com/Golangltd/Twlib/proto"
	. "github.com/Golangltd/Twlib/user"
)

//抽卡界面
type DCInterface struct {
	Protocol     int
	Protocol2    int
	PropList     []*PropData //显示的道具列表
	TenDrawsData *DrawSpend
	OneDrawData  *DrawSpend
}

func NewDCInterface() *DCInterface {
	m := &DCInterface{}
	m.Protocol = GGameHallProto
	m.Protocol2 = S2CDCInterface
	m.TenDrawsData = &DrawSpend{}
	m.OneDrawData = &DrawSpend{}
	return m
}

func (m *DCInterface) AddFieldData(userST *UserSt) {
	m.initPropList(userST)
	m.initDrawSpend(userST)
}

//初始化道具列表
func (m *DCInterface) initPropList(userST *UserSt) {
	m.PropList = make([]*PropData, 0) //初始化内存
	for _, propID := range tables.CurrencyList {
		if propID == Gold { //如果道具ID等于金币
			m.PropList = append(m.PropList, newPropData(propID, userST.Diamond))
		} else {
			m.PropList = append(m.PropList, newPropData(propID, getItemAmount(propID, userST)))
		}
	}
}

//初始化抽奖花费
func (m *DCInterface) initDrawSpend(userST *UserSt) {
	consume, err := tables.GetGoldConsumeConf()
	if err != nil {
		log.Error(err.Error())
		return
	}
	m.initOneDrawData(consume, userST)
	m.initTenDrawsData(consume, userST)
}

func (m *DCInterface) initOneDrawData(consume *tables.ConsumeMode, userST *UserSt) {
	m.OneDrawData.Spend = consume.DrawOneTime.ConsumeAmount
	if CDUsedFreeTimes[userST.RoleUid] < tables.GetCommonDrawFreeTimes() { //如果当前用户已经使用的免费次数小于配置中的免费次数,那么说明是可以抽奖的
		m.OneDrawData.IsDraw = true
	} else {
		if (userST.Diamond) >= int64(m.OneDrawData.Spend) { //当前用户的coin和钻石之和大于等于抽一次的花费,说明用户是可以抽奖的
			m.OneDrawData.IsDraw = true
		}
	}
}

func (m *DCInterface) initTenDrawsData(consume *tables.ConsumeMode, userST *UserSt) {
	m.TenDrawsData.Spend = consume.DrawTenTimes.ConsumeAmount
	if (tables.GetCommonDrawFreeTimes() - CDUsedFreeTimes[userST.RoleUid]) >= 10 { //这个10值一定是死值,允许的免费次数和10次相比较,如果大于等于10次，则允许抽奖
		m.TenDrawsData.IsDraw = true
	} else {
		if (userST.Diamond) >= int64(m.TenDrawsData.Spend) {
			m.TenDrawsData.IsDraw = true
		}
	}
}

//获取道具数量
func getItemAmount(propID int, userST *UserSt) (itemAmount int64) {
	for _, itemData := range userST.ItemList {
		if itemData.ItemId == propID {
			itemAmount = itemData.ItemNum
			break
		}
	}
	return
}
