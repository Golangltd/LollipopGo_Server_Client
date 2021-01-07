package tables

import (
	"LollipopGo/log"
	"LollipopGo2.8x/conf/g"
	gameDB "LollipopGo2.8x/data"
	twLibDBTable "github.com/Golangltd/Twlib/dbtable"
	"sync"
)

// 学籍配置表
var (
	SchoolRollTable map[int]*SchoolRollSample
	schoolRollMsg   []interface{}
	schoolRollOnce  sync.Once
)

//学籍样本数据
type SchoolRollSample struct {
	Level       int         //等级
	SRName      int         //学籍名称(学籍关联的文本)
	Exp         int         //升级所需经验
	ChallengeID int         //突破试炼副本ID
	SRLevelAtr  map[int]int //学籍等级属性
}

func AddSchoolRollTable() {
	data := gameDB.GetCFGameData(twLibDBTable.Gl_SchoolRool)
	if data == nil {
		log.Fatal("error! receive SchoolRollTable is nil!!")
	}
	schoolRollMsg = data
	schoolRollOnce.Do(initSchoolRollTable)
}

func initSchoolRollTable() {
	SchoolRollTable = make(map[int]*SchoolRollSample)
	tempArr := make([]int, 0)
	for _, sample := range schoolRollMsg {
		temp := sample.(map[string]interface{})
		m := &SchoolRollSample{}
		m.Level = StrToInt(temp["sid"].(string))
		g.WizardMaxLevel = m.Level //更新最大等级数
		m.SRName = StrToInt(temp["name"].(string))
		m.Exp = StrToInt(temp["exp"].(string))
		m.ChallengeID = StrToInt(temp["challengeID"].(string))
		if m.ChallengeID <= 0 {
			tempArr = append(tempArr, m.Level) //将此等级添加进入临时列表
		} else {
			addChallengeID(m.ChallengeID, tempArr)
			tempArr = make([]int, 0) //重新初始化
		}
		m.SRLevelAtr = InitAttribute(temp["atr1"].(string), `\d+,\d+`)
		SchoolRollTable[m.Level] = m
	}
	schoolRollMsg = nil
	tempArr = nil
}

//给每个学籍等级添加挑战ID
//challengeID:挑战ID lvArr:等级列表
func addChallengeID(challengeID int, lvArr []int) {
	for _, lv := range lvArr {
		SchoolRollTable[lv].ChallengeID = challengeID
	}
}
