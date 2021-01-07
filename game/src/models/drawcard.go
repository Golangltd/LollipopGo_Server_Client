package models

import (
	. "LollipopGo2.8x/conf/error"
	"LollipopGo2.8x/conf/g"
	. "LollipopGo2.8x/msg/drawcard"
	"LollipopGo2.8x/tables"
)

/*抽卡系统model*/
type DrawCardModel struct {
	OpenId string
	*Game
	*DCInterface
	*DrawData
	SpendAmount int        //抽奖花费
	DrawType    g.DrawType //抽奖类型
}

//抽卡系统model
func NewDrawCardModel(openID string, game *Game) *DrawCardModel {
	m := &DrawCardModel{}
	m.OpenId = openID
	m.Game = game
	m.DCInterface = NewDCInterface()
	m.DrawData = NewDrawData()
	return m
}

//获取抽奖界面数据
func (m *DrawCardModel) GetDCInterfaceData() {
	m.DCInterface.AddFieldData(m.UserInfo)
}

//抽奖数据
func (m *DrawCardModel) GetDrawData(drawType int) {
	m.DrawData.InitDrawData() //将drawData数据初始化到原始状态
	m.DrawType = g.DrawType(drawType)
	drawCount, canDrawCard := m.addCodeAndGetCount()
	if !canDrawCard { //不可抽奖，直接返回
		return
	}
	for count := 0; count < drawCount; count++ {
		m.DrawData.AddDrawFieldData(g.CommonDraw, m.Game.UserInfo.RoleUid)
	}
	m.initSpendAmount(drawCount)
}

//添加错误码并且获取抽奖次数
func (m *DrawCardModel) addCodeAndGetCount() (count int, canDrawCard bool) {
	canDrawCard = true //默认为可以抽奖
	//是否能使用免费机会
	canUseFree := (tables.GetCommonDrawFreeTimes() - g.CDUsedFreeTimes[m.Game.UserInfo.RoleUid]) >= int(m.DrawType)
	if m.DrawType == g.OneDraws {
		count = 1
		if m.Game.UserInfo.Coin < int64(m.DCInterface.OneDrawData.Spend) && !canUseFree { //用户的金币小于一次抽奖花费的金币并且不能免费抽奖
			m.DrawData.Code = NoEnoughCoin
			canDrawCard = false
		}
	} else {
		count = 10
		if m.Game.UserInfo.Coin < int64(m.DCInterface.TenDrawsData.Spend) && !canUseFree { //用户的金币小于十次抽奖花费的金币并且不能免费抽奖
			m.DrawData.Code = NoEnoughCoin
			canDrawCard = false
		}
	}
	return
}

//初始化抽奖花费
//drawCount:10抽 1抽
func (m *DrawCardModel) initSpendAmount(drawCount int) {
	//免费次数大于用户已经使用的免费次数,说明此次用户抽奖是免费的
	if (tables.GetCommonDrawFreeTimes() - g.CDUsedFreeTimes[m.Game.UserInfo.RoleUid]) >= drawCount {
		g.CDUsedFreeTimes[m.Game.UserInfo.RoleUid] += drawCount //记录使用的免费抽奖次数
		return
	}
	if m.DrawType == g.OneDraws {
		m.SpendAmount = m.DCInterface.OneDrawData.Spend
	} else {
		m.SpendAmount = m.DCInterface.TenDrawsData.Spend
	}
}
