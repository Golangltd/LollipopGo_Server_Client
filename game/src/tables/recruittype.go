package tables

import (
	. "LollipopGo2.8x/conf/g"
	. "LollipopGo2.8x/lua/go"
	"errors"
)

/*巫师招募类型表*/
var (
	RecruitTypeTable map[int]*RecruitTypeSample
)

func init() {
	initRecruitTypeTable()
}

type RecruitTypeSample struct {
	RecruitType       int          //招募类型
	FirstConsumeMode  *ConsumeMode //第一种消耗方式
	SecondConsumeMode *ConsumeMode //第二种消耗方式
	Number            *Number
	FreeTimes         int    //免费次数
	TimesResetTime    string //次数重置时间
	Point             *Point //积分
	PoolType          []int  //卡池类型
}

//消耗方式
type ConsumeMode struct {
	DrawOneTime  *DrawConsume //抽一次
	DrawTenTimes *DrawConsume //十连抽
}

func newConsumeMode(drawConsumeList []*DrawConsume) *ConsumeMode {
	m := &ConsumeMode{}
	m.DrawOneTime = drawConsumeList[0]
	m.DrawTenTimes = drawConsumeList[1]
	return m
}

//抽奖消耗
type DrawConsume struct {
	ItemType      int //道具类型
	ItemID        int //道具ID
	ConsumeAmount int //消耗数量
}

func newDrawConsume(strList []int) *DrawConsume {
	m := &DrawConsume{}
	m.ItemType = strList[0]
	m.ItemID = strList[1]
	m.ConsumeAmount = strList[2]
	return m
}

//对应次数
type Number struct {
	FirstDegree  int //第一种次数
	SecondDegree int //第二种次数
}

func newNumber(strList []int) *Number {
	m := &Number{}
	m.FirstDegree = strList[0]
	m.SecondDegree = strList[1]
	return m
}

type Point struct {
	ItemType    int //道具类型
	ItemID      int //道具ID
	PointAmount int //积分数量
}

func newPoint(strList []int) *Point {
	m := &Point{}
	m.ItemType = strList[0]
	m.ItemID = strList[1]
	m.PointAmount = strList[2]
	return m
}

func initRecruitTypeTable() {
	RecruitTypeTable = make(map[int]*RecruitTypeSample)
	for _, v := range Grecruitmenttype_proto {
		m := &RecruitTypeSample{}
		m.RecruitType = StrToInt(v.Sid)
		m.FirstConsumeMode = initConsumeMode(v.Price, `\d+,\d+,\d+`)
		m.SecondConsumeMode = initConsumeMode(v.Gold, `\d+,\d+,\d+`)
		m.Number = newNumber(getKeyAndValue(v.Number))
		m.FreeTimes = StrToInt(v.Free_number)
		m.TimesResetTime = GetCronStr(v.Reset)
		m.Point = initPoint(v.Point, `\d+,\d+,\d+`)
		m.PoolType = getKeyAndValue(v.Pool_id[1 : len(v.Pool_id)-2])
		RecruitTypeTable[m.RecruitType] = m
	}
}

func initConsumeMode(str string, regexpStr string) *ConsumeMode {
	if len(str) == 0 {
		return nil
	}
	strArr := GetRegexpArr(str, regexpStr)
	drawConsumeList := make([]*DrawConsume, 0)
	for _, temp := range strArr {
		drawConsumeList = append(drawConsumeList, newDrawConsume(getKeyAndValue(temp[0])))
	}
	return newConsumeMode(drawConsumeList)
}

func initPoint(str string, regexpStr string) *Point {
	strArr := GetRegexpArr(str, regexpStr)
	for _, temp := range strArr {
		return newPoint(getKeyAndValue(temp[0]))
	}
	return nil
}

//获取金币消耗conf
func GetGoldConsumeConf() (*ConsumeMode, error) {
	sampleData := RecruitTypeTable[int(CommonDraw)]
	consumeList := []*ConsumeMode{sampleData.FirstConsumeMode, sampleData.SecondConsumeMode}
	for _, consume := range consumeList {
		if consume.DrawOneTime.ItemID == Gold { //随意判断consume的itemID,如果道具ID为金币,那么就返回此ConsumeMode
			return consume, nil
		}
	}
	return nil, errors.New("don't find gold consume conf")
}

//获取普通抽奖的免费次数
func GetCommonDrawFreeTimes() int {
	return RecruitTypeTable[int(CommonDraw)].FreeTimes
}
