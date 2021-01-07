package tables

import (
	"LollipopGo/log"
	"LollipopGo2.8x/conf/g"
	gameDB "LollipopGo2.8x/data"
	twLibDBTable "github.com/Golangltd/Twlib/dbtable"
	"strings"
	"sync"
)

// 学院配置表
var (
	collegeOnce  sync.Once
	collegeMsg   []interface{}
	CollegeTable map[int]*CollegeSample
)

//学院样本数据
type CollegeSample struct {
	ID           int
	CollegeType  int            //学院类型
	CollegeLevel int            //学院等级
	UpgradeTerms []*UpgradeTerm //升级所需要的条件
	ClassRoomLV  int            //教室等级
	SeatAmount   int            //座位数量
	CourseAward  *CourseAward   //课程收益
}

type UpgradeTerm struct {
	Amount      int //数量
	Camp        int //阵营
	TalentLevel int //天赋等级
}

type CourseAward struct {
	PropType   int //道具类型
	PropID     int //道具ID
	PropAmount int //道具数量
}

func newCourseAward(strArr []string) *CourseAward {
	m := &CourseAward{}
	m.PropType = StrToInt(strArr[0])
	m.PropID = StrToInt(strArr[1])
	m.PropAmount = StrToInt(strArr[2])
	return m
}

func newUpgradeTerm(strArr []string) *UpgradeTerm {
	m := &UpgradeTerm{}
	m.Amount = StrToInt(strArr[0])
	m.Camp = StrToInt(strArr[1])
	m.TalentLevel = StrToInt(strArr[2])
	return m
}

// 添加学院表
func AddCollegeTable() {
	data := gameDB.GetCFGameData(twLibDBTable.Gl_School)
	if data == nil {
		log.Fatal("error! receive CollegeTable is nil!!")
	}
	collegeMsg = data
	collegeOnce.Do(initCollegeTable)
}

//初始化学院表
func initCollegeTable() {
	CollegeTable = make(map[int]*CollegeSample)
	for _, sample := range collegeMsg {
		temp := sample.(map[string]interface{})
		m := &CollegeSample{}
		m.ID = StrToInt(temp["sid"].(string))
		m.CollegeType = StrToInt(temp["type"].(string))
		m.CollegeLevel = StrToInt(temp["level"].(string))
		if m.CollegeLevel > g.CollegeMaxLevel {
			g.CollegeMaxLevel = m.CollegeLevel
		}
		m.initUpgradeTerm(temp["missiontype"].(string), `\d+,\d+,\d+`)
		m.ClassRoomLV = StrToInt(temp["classlevel"].(string))
		m.SeatAmount = StrToInt(temp["seatnumber"].(string))
		m.initCourseAward(temp["reward"].(string), `\d+,\d+,\d+`)
		CollegeTable[m.ID] = m
	}
}

func (m *CollegeSample) initUpgradeTerm(str string, regexpStr string) {
	m.UpgradeTerms = make([]*UpgradeTerm, 0)
	regexpArr := GetRegexpArr(str, regexpStr)
	for _, tempArr := range regexpArr {
		strArr := strings.Split(tempArr[0], ",")
		if len(strArr) != 3 {
			log.Error("college conf's missiontype is error!,this missiontype is %v", tempArr[0])
		}
		m.UpgradeTerms = append(m.UpgradeTerms, newUpgradeTerm(strArr))
	}
}

func (m *CollegeSample) initCourseAward(str string, regexpStr string) {
	regexpArr := GetRegexpArr(str, regexpStr)
	for _, tempArr := range regexpArr {
		strArr := strings.Split(tempArr[0], ",")
		if len(strArr) != 3 {
			log.Error("college conf's reward is error!,this reward is %v", tempArr[0])
		}
		m.CourseAward = newCourseAward(strArr)
	}
}
