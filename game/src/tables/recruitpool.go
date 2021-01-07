package tables

/*招募卡池*/

import (
	"LollipopGo/tools/sample"
	. "LollipopGo2.8x/lua/go"
	. "LollipopGo2.8x/conf/g"
)

var (
	RecruitPoolTable map[int]*RecruitPoolSample                 //招募卡池表
	CardsPoolHelper  map[RecruitType]map[int]*RecruitPoolSample //卡池Helper
)

func init() {
	initRecruitPoolTable()
}

type RecruitPoolSample struct {
	ID          int
	PoolType    int           //卡池类型
	WizardCards []*WizardCard //巫师卡牌集群
	Quality     int           //品质
	Weight      int           //概率
	UpRate      int           //幸运增长比
}

//随机一张巫师巫师卡牌
func (m *RecruitPoolSample) RandWizardCard() *WizardCard {
	weights := make([]int, 0)
	for _, wizardCard := range m.WizardCards {
		weights = append(weights, wizardCard.Rate)
	}
	index := sample.WeightedChoice(weights)
	return m.WizardCards[index]
}

func (m *RecruitPoolSample) IsUpRate() bool {
	return m.UpRate > 0
}

//巫师卡牌
type WizardCard struct {
	PropType int //类型
	PropID   int //道具索引
	Amount   int //数量
	Rate     int //概率
}

func newWizardCard(dataList []int) *WizardCard {
	m := &WizardCard{}
	m.PropType = dataList[0]
	m.PropID = dataList[1]
	m.Amount = dataList[2]
	m.Rate = dataList[3]
	return m
}

func initRecruitPoolTable() {
	RecruitPoolTable = make(map[int]*RecruitPoolSample)
	CardsPoolHelper = make(map[RecruitType]map[int]*RecruitPoolSample)
	for _, v := range Grecruitmentpool_proto {
		m := &RecruitPoolSample{}
		m.ID = StrToInt(v.Sid)
		m.PoolType = StrToInt(v.Pool_id)
		m.Quality = StrToInt(v.Quality)
		m.Weight = StrToInt(v.Weight)
		m.WizardCards = initWizardCards(v.Cards, `\d+,\d+,\d+,\d+`)
		m.UpRate = StrToInt(v.Up)
		RecruitPoolTable[m.ID] = m
		if _, ok := CardsPoolHelper[RecruitType(m.PoolType)]; !ok { //没有这个key初始化内存
			CardsPoolHelper[RecruitType(m.PoolType)] = make(map[int]*RecruitPoolSample)
		}
		CardsPoolHelper[RecruitType(m.PoolType)][m.ID] = m //存入CardsPoolHelper
	}
}

func initWizardCards(str string, regexpStr string) []*WizardCard {
	wizardCards := make([]*WizardCard, 0)
	strArr := GetRegexpArr(str, regexpStr)
	for _, temp := range strArr {
		wizardCards = append(wizardCards, newWizardCard(getKeyAndValue(temp[0])))
	}
	return wizardCards
}

//获取奖池样本数据
func GetRecruitPoolSampleByUserID(userID int64, rType RecruitType) *RecruitPoolSample {
	var (
		ok     bool
		upRate map[int]int
	)
	if upRate, ok = UsersCardsPoolUpRate[userID][rType]; !ok { //根据用户ID和招募类型获取用户在此招募类型下的所有奖池的幸运增长比
		upRate = make(map[int]int)
	}
	keyList := make([]int, 0) //索引list
	weightList := make([]int, 0)
	for _, sampleData := range CardsPoolHelper[rType] {
		keyList = append(keyList, sampleData.ID)
		weightList = append(weightList, sampleData.Weight+upRate[sampleData.ID])
	}
	index := sample.WeightedChoice(weightList) //获取索引
	poolKey := keyList[index]                  //获取对应的奖池
	return RecruitPoolTable[poolKey]
}
