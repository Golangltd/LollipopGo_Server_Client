package tables

import (
	"LollipopGo/log"
	"strings"
	"sync"
)

/*
学籍试炼表
*/

var (
	srTrialOnce  sync.Once
	srTrialMsg   []interface{}
	SRTrialTable map[int]*SRTrialSample
)

type SRTrialSample struct {
	ID             int
	TrialCondition *TrialCondition
	RoundMax       int //最大战斗回合
	SkipRound      int //跳过回合
	MonsterGroup   int //怪物组
}

//试炼条件
type TrialCondition struct {
	CollegeType int //学院类型
	CollegeLV   int //学院等级
}

func newTrialCondition(strArr []string) *TrialCondition {
	m := &TrialCondition{}
	m.CollegeType = StrToInt(strArr[0])
	m.CollegeLV = StrToInt(strArr[1])
	return m
}

// 添加试炼表
func AddSRTrialTable() {

}

func InitSRTrialTable() {
	SRTrialTable = make(map[int]*SRTrialSample)
	for _, sample := range srTrialMsg {
		temp := sample.(map[string]interface{})
		m := &SRTrialSample{}
		m.ID = StrToInt(temp["sid"].(string))
		m.initTrialCondition(temp["condition"].(string), `\d+,\d+`)
		m.RoundMax = StrToInt(temp["round_max"].(string))
		m.SkipRound = StrToInt(temp["skip_round"].(string))
		m.MonsterGroup = StrToInt(temp["monstergroup"].(string))
		SRTrialTable[m.ID] = m
	}
	srTrialMsg = nil
}

func (m *SRTrialSample) initTrialCondition(str string, regexpStr string) {
	regexpArr := GetRegexpArr(str, regexpStr)
	for _, tempArr := range regexpArr {
		strArr := strings.Split(tempArr[0], ",")
		if len(strArr) != 2 {
			log.Error("schoolRollTrial conf's TrialCondition is error!,this TrialCondition is %v", tempArr[0])
		}
		m.TrialCondition = newTrialCondition(strArr)
	}
}
