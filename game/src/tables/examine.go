package tables

import (
	. "LollipopGo2.8x/lua/go"
)

//考试配置表
var (
	ExamineTable  map[int]*ExamineSampleData
	examineHelper map[int]map[int][]*ExamineSampleData //每轮每种不同类型的配置数据
	ExamineLink   *ExamineNode                         //考核链表
)

type ExamineSampleData struct {
	ID               int        //任务ID
	TaskType         int        //任务类型
	Rounds           int        //轮数
	TaskLevel        int        //任务等级
	UpgradeExp       int        //升级经验
	DisplayCondition *Condition //显示条件
	AwardsCondition  *Condition //领奖条件
	HighAwards       []*Award   //高级奖励
	CommonAwards     []*Award   //初级奖励
}

//显示条件
type Condition struct {
	FunctionType int //功能类型
	ConditionLev int //条件等级
}

func newCondition(str string) *Condition {
	m := &Condition{}
	arr := StringToIntSlice(str)
	m.FunctionType = arr[0]
	m.ConditionLev = arr[1]
	return m
}

//领奖条件
type Award struct {
	ItemType   int //道具类型
	ItemID     int //道具ID
	ItemAmount int //道具数量
}

func newAward(dataList []int) *Award {
	m := &Award{}
	m.ItemType = dataList[0]
	m.ItemID = dataList[1]
	m.ItemAmount = dataList[2]
	return m
}

func init() {
	initExamineTable()
	examineHelperSort()
	initExamineLink()
}

//初始化考试表
func initExamineTable() {
	ExamineTable = make(map[int]*ExamineSampleData)
	examineHelper = make(map[int]map[int][]*ExamineSampleData)
	for _, sampleData := range ExamineConf {
		m := &ExamineSampleData{}
		m.ID = StrToInt(sampleData.Sid)
		m.TaskType = StrToInt(sampleData.Type)
		m.Rounds = StrToInt(sampleData.Turn)
		m.TaskLevel = StrToInt(sampleData.Level)
		m.UpgradeExp = StrToInt(sampleData.Score)
		m.DisplayCondition = newCondition(sampleData.Show_condition)
		m.AwardsCondition = newCondition(sampleData.Get_condition)
		m.HighAwards = initAwards(sampleData.Award_hight, `\d+,\d+,\d+`)
		m.CommonAwards = initAwards(sampleData.Award_common, `\d+,\d+,\d+`)
		ExamineTable[m.ID] = m
		if _, ok := examineHelper[m.Rounds]; !ok {
			examineHelper[m.Rounds] = make(map[int][]*ExamineSampleData)
		}
		examineHelper[m.Rounds][m.TaskType] = append(examineHelper[m.Rounds][m.TaskType], m)
	}
}

func initAwards(str string, regexpStr string) []*Award {
	awards := make([]*Award, 0)
	strArr := GetRegexpArr(str, regexpStr)
	for _, temp := range strArr {
		awards = append(awards, newAward(getKeyAndValue(temp[0])))
	}
	return awards
}

//考试helper排序
func examineHelperSort() {
	for _, v1 := range examineHelper {
		for _, v2 := range v1 {
			for i := 0; i < len(v2)-1; i++ {
				for j := i + 1; j > 0; j-- {
					if v2[j].TaskLevel < v2[j-1].TaskLevel {
						v2[j], v2[j-1] = v2[j-1], v2[j]
					} else {
						break
					}
				}
			}
		}
	}
}

/*
生成一个链表，方便后续业务逻辑操作
*/

//考核轮数数据
type ExamineNode struct {
	ERM      *ExamineTurnsData
	NextNode *ExamineNode
}

func newExamineRoundData(erm *ExamineTurnsData) *ExamineNode {
	m := &ExamineNode{}
	m.ERM = erm
	return m
}

func (m *ExamineNode) insert(nextNode *ExamineNode) {
	if m.NextNode == nil {
		m.NextNode = nextNode
	} else {
		p := m
		for p.NextNode != nil {
			p = p.NextNode
		}
		p.NextNode = nextNode
	}
}

//考试m每轮数据
type ExamineTurnsData struct {
	SampleDataList []*ExamineSampleData //每种类型的样本数据
	NextTurnData   *ExamineTurnsData
}

func newExamineRoundsModel(data []*ExamineSampleData) *ExamineTurnsData {
	m := &ExamineTurnsData{}
	m.SampleDataList = data
	return m
}

func (m *ExamineTurnsData) insert(nextNode *ExamineTurnsData) {
	if m.NextTurnData == nil {
		m.NextTurnData = nextNode
	} else {
		p := m
		for p.NextTurnData != nil {
			p = p.NextTurnData
		}
		p.NextTurnData = nextNode
	}
}

//初始化考核链表
func initExamineLink() {
	ExamineLink = &ExamineNode{}
	sortKeys := getSortKey()
	for _, keyValue := range sortKeys {
		tempArr := make([][]*ExamineSampleData, 0)
		for _, dataList := range examineHelper[keyValue] { //将每种任务类型的列表数据添加到tempArr中
			tempArr = append(tempArr, dataList)
		}
		tempArrSort(tempArr)
		examineTypeModel := initExamineTypeModel(tempArr)
		ExamineLink.insert(newExamineRoundData(examineTypeModel))
	}
}

func getSortKey() []int {
	keyList := make([]int, 0)
	for key, _ := range examineHelper {
		keyList = append(keyList, key)
	}
	for i := 0; i < len(keyList)-1; i++ {
		for j := i + 1; j > 0; j-- {
			if keyList[j] < keyList[j-1] {
				keyList[j], keyList[j-1] = keyList[j-1], keyList[j]
			} else {
				break
			}
		}
	}
	return keyList
}

func tempArrSort(tempArr [][]*ExamineSampleData) {
	for i := 0; i < len(tempArr)-1; i++ {
		for j := i + 1; j > 0; j-- {
			if tempArr[j][0].TaskType < tempArr[j-1][0].TaskType {
				tempArr[j], tempArr[j-1] = tempArr[j-1], tempArr[j]
			} else {
				break
			}
		}
	}
}

func initExamineTypeModel(arr [][]*ExamineSampleData) *ExamineTurnsData {
	m := &ExamineTurnsData{}
	m.SampleDataList = arr[0]
	for _, data := range arr {
		m.insert(newExamineRoundsModel(data))
	}
	//这里把这个链表设置为循环链表
	p := m
	for p.NextTurnData != nil {
		p = p.NextTurnData
	}
	p.NextTurnData = m //首尾相连
	return m
}
