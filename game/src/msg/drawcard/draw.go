package drawcard

import (
	"LollipopGo/log"
	. "LollipopGo2.8x/conf/error"
	"LollipopGo2.8x/conf/g"
	. "LollipopGo2.8x/proto/dc_proto"
	. "LollipopGo2.8x/tables"
	. "LollipopGo2.8x/util/uint64"
	. "github.com/Golangltd/Twlib/proto"
	"github.com/mitchellh/mapstructure"
)

//抽卡请求(十抽和一抽)
type DrawReqData struct {
	Protocol  int
	Protocol2 int
	OpenId    string
	DrawType  int //抽卡类型(0：1抽 1：10抽)
}

func NewDrawReqData(data map[string]interface{}) *DrawReqData {
	m := &DrawReqData{}
	if err := mapstructure.Decode(data, m); err != nil {
		log.Error("data[%v] decode DrawReqData error,error is [%v]! ", data, err)
	}
	m.Protocol = GGameHallProto
	m.Protocol2 = C2STenDrawReq
	return m
}

//抽奖数据
type DrawData struct {
	Protocol  int
	Protocol2 int
	CardList  []*CardData
	Code      Code //Code码
}

func NewDrawData() *DrawData {
	m := &DrawData{}
	m.Protocol = GGameHallProto
	m.Protocol2 = S2CDrawData
	return m
}

//卡牌数据
type CardData struct {
	CardUid  uint64 //卡牌唯一ID
	ItemType int    //道具类型
	CardID   int    //卡牌ID
	Quality  int    //品质
}

func newCardData(sampleData *RecruitPoolSample) *CardData {
	m := &CardData{}
	wizardCard := sampleData.RandWizardCard()
	m.CardUid = GetUint64()
	m.ItemType = wizardCard.PropType
	m.CardID = wizardCard.PropID
	m.Quality = sampleData.Quality
	return m
}

func (m *DrawData) InitDrawData() {
	m.CardList = make([]*CardData, 0)
	m.Code = CorrectCode //code码初始化为0
}

//添加抽奖字段数据
func (m *DrawData) AddDrawFieldData(recruitType g.RecruitType, userID int64) {
	sampleData := GetRecruitPoolSampleByUserID(userID, recruitType)
	m.CardList = append(m.CardList, newCardData(sampleData))
	m.resetPoolUpRate(sampleData, recruitType, userID)
	m.addUsersUpdate(recruitType, sampleData.ID, userID)
}

//重置某类型奖池的幸运增长比(通过recruitType(招募类型)重置该用户此招募类型下的幸运增长比)
func (m *DrawData) resetPoolUpRate(sampleData *RecruitPoolSample, recruitType g.RecruitType, userID int64) {
	if sampleData.IsUpRate() { //如果此奖池存在幸运增长比
		if g.UsersCardsPoolUpRate[userID][recruitType] != nil { //如果此抽奖类型存在幸运增长比就重置
			g.UsersCardsPoolUpRate[userID][recruitType] = make(map[int]int) //重置
		}
	}
}

//如果同类型的某个奖池有幸运增长比,那么就进行增加幸运增长比，相对的其他奖池减少幸运增长比
//hitPoolID:随机被击中的奖池ID
func (m *DrawData) addUsersUpdate(recruitType g.RecruitType, hitPoolID int, userID int64) {
	poolsSampleData := CardsPoolHelper[recruitType]
	for _, data := range poolsSampleData {
		if data.ID != hitPoolID { //说明这是同类型卡池没有随机击中的卡池,说明可以去进行判断它是否有幸运增长比
			if data.IsUpRate() { //判断是否拥有幸运增长比
				m.updateUpRate(recruitType, userID, data.ID, poolsSampleData)
			}
		}
	}
}

//update 幸运增长比
func (m *DrawData) updateUpRate(recruitType g.RecruitType, userID int64, hitPoolID int, data map[int]*RecruitPoolSample) {
	if g.UsersCardsPoolUpRate[userID][recruitType] == nil { //没有初始化内存
		g.UsersCardsPoolUpRate[userID] = make(map[g.RecruitType]map[int]int) //初始内存
		g.UsersCardsPoolUpRate[userID][recruitType] = make(map[int]int)
	}
	reduceRate := m.getReduceRate(hitPoolID, data) //
	for _, sampleData := range data {              //遍历此招募类型下的所有奖池数据
		if sampleData.ID != hitPoolID {
			g.UsersCardsPoolUpRate[userID][recruitType][sampleData.ID] -= reduceRate
		} else {
			g.UsersCardsPoolUpRate[userID][recruitType][sampleData.ID] += data[hitPoolID].UpRate
		}
	}
}

//获取减去的rate(其它同类型的奖池被减概率应该平均分配)
func (m *DrawData) getReduceRate(hitPoolID int, data map[int]*RecruitPoolSample) int {
	return data[hitPoolID].UpRate / (len(data) - 1)
}
