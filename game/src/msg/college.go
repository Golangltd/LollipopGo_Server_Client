package gamemsg

import (
	"LollipopGo/log"
	. "LollipopGo2.8x/proto/sr_proto"
	"LollipopGo2.8x/tables"
	. "github.com/Golangltd/Twlib/proto"
	. "github.com/Golangltd/Twlib/user"
	"github.com/mitchellh/mapstructure"
)

//获取学院详情请求
type CollegeDetailReq struct {
	Protocol  int
	Protocol2 int
	OpenId    string
	CollegeID int //学院ID
}

func NewCollegeDetailReq(data map[string]interface{}) *CollegeDetailReq {
	m := &CollegeDetailReq{}
	if err := mapstructure.Decode(data, m); err != nil {
		log.Error("data[%v] decode SchoolRollDetailReq error,error is [%v]! ", data, err)
	}
	m.Protocol = GGameBattleProto
	m.Protocol2 = C2SGetCollegeDetail
	return m
}

//学院详情信息
type CollegeDetailMsg struct {
	Protocol          int
	Protocol2         int
	CollegeTypeID     int              //学院类型ID
	Level             int              //学院等级
	IsCanUpgrade      bool             //是否可以升级
	UpgradeDetailList []*UpgradeDetail //升级详情列表
	MapTalents        []*MapTalent     //图鉴天赋列表
}

//升级详情
type UpgradeDetail struct {
	WizardLevel     int //巫师等级
	CurrentAmount   int //当前巫师的个数
	ConditionAmount int //升级条件的巫师个数
}

func NewCollegeDetailInfo() *CollegeDetailMsg {
	m := &CollegeDetailMsg{}
	m.Protocol = GGameBattleProto
	m.Protocol2 = S2CCollegeDetail
	return m
}

//添加数据信息
func (m *CollegeDetailMsg) AddDataInfo(c *CollegeInfo) {
	collegeSampleData := tables.CollegeTable[c.CollegeID]
	m.CollegeTypeID = collegeSampleData.CollegeType //添加学院类型ID
	m.Level = collegeSampleData.CollegeLevel        //添加学院等级
	m.IsCanUpgrade = true                           //默认满足升级条件
	m.initUpgradeDetail(c, collegeSampleData)
	c.SortMapTalent() //图鉴排一下序
	m.MapTalents = c.MapTalents
}

//初始化升级详情
func (m *CollegeDetailMsg) initUpgradeDetail(c *CollegeInfo, sampleData *tables.CollegeSample) {
	m.UpgradeDetailList = make([]*UpgradeDetail, 0)
	//遍历升级所需要的条件
	for _, term := range sampleData.UpgradeTerms {
		upgradeDetail := &UpgradeDetail{}
		upgradeDetail.WizardLevel = term.TalentLevel //巫师等级
		upgradeDetail.ConditionAmount = term.Amount  //获取升级的条件个数
		upgradeDetail.CurrentAmount = m.GetMapAmountByCampAndLV(c.MapTalents, term.Camp, term.TalentLevel)
		if upgradeDetail.CurrentAmount < upgradeDetail.ConditionAmount { //如果当前的巫师图鉴个数小于条件要求的个数,那么是否可以升级直接置为false
			m.IsCanUpgrade = false
		}
		m.UpgradeDetailList = append(m.UpgradeDetailList, upgradeDetail)
	}
}

/*
通过阵营和等级条件获取个数
mapTalents:图鉴列表 camp:阵营 lv:等级
*/
func (m *CollegeDetailMsg) GetMapAmountByCampAndLV(mapTalents []*MapTalent, camp int, lv int) (amount int) {
	/*	for _, mapTalent := range mapTalents {
		cardTJSampleData := tables.CardTJTables[mapTalent.MapID]
		if tables.RolesTable[cardTJSampleData.CardID].Camp == camp && mapTalent.TalentLevel == lv {
			amount++
		}
	}*/
	return
}
