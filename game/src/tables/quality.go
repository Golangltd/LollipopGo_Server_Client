package tables

import (
	"LollipopGo/log"
	gamedb "LollipopGo2.8x/data"
	twlib_dbtable "github.com/Golangltd/Twlib/dbtable"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

var (
	PinZhiTables map[int]*PinZhiSt
	PinZhiMsg    []interface{}
	PinZhiOnce   sync.Once
)

type PinZhiSt struct {
	Sid        string
	Camp_type  string // 1 是普通，2 特殊
	LevelUp    string // 是否可以进行生品
	Level      string // 品质ID
	Next       string // 下一级品质ID
	Card       string // 需要副卡的数量
	SameCard1  string // 副卡1是否同名
	SameCard2  string // 副卡2是否同名
	CardLevel1 string //
	CardLevel2 string //
	CardGroup  string // 是否同阵营
}

// 读取章节数据
func GetPinZhiInfo() {
	data := gamedb.GetCFGameData(twlib_dbtable.Gl_wse_wsjyb)
	AddWuShiTables(data)
}

func AddPinZhiTables(msg []interface{}) {
	if msg == nil {
		log.Fatal("error! receive skillTables is nil!!")
	}
	wuShiMsg = msg
	wuShiOnce.Do(initWuShiTables)
}

func initPinZhiTables() {
	PinZhiTables = make(map[int]*PinZhiSt)
	for _, sample := range PinZhiMsg {
		temp := sample.(map[string]interface{})
		m := &PinZhiSt{}
		m.Sid = temp["sid"].(string)
		m.Camp_type = temp["camp_type"].(string)
		m.LevelUp = temp["levelup"].(string)
		m.Level = temp["level"].(string)
		m.Next = temp["next"].(string)
		m.Card = temp["card"].(string)
		m.SameCard1 = temp["samecard1"].(string)
		m.SameCard2 = temp["samecard2"].(string)
		m.CardLevel1 = temp["cardlevel1"].(string)
		m.CardLevel2 = temp["cardlevel2"].(string)
		m.CardGroup = temp["cardgroup"].(string)
		sid, _ := strconv.Atoi(m.Sid)
		PinZhiTables[sid] = m
	}
}

// 获取奖励
func GetPinZhiFor2Row(data string) map[int]int {
	retdata := make(map[int]int)
	re := regexp.MustCompile(`\d+,\d+`)
	strArr := re.FindAllStringSubmatch(data, -1)
	for i := 0; i < len(strArr); i++ {
		arr := strings.Split(strArr[i][0], ",")
		ikey, _ := strconv.Atoi(arr[0])
		ival, _ := strconv.Atoi(arr[1])
		retdata[ikey] = ival
	}
	return retdata
}

// 获取奖励
func GetPinZhiFor3Row(data string) map[int]int {
	retdata := make(map[int]int)
	re := regexp.MustCompile(`\d+,\d+`)
	strArr := re.FindAllStringSubmatch(data, -1)
	for i := 0; i < len(strArr); i++ {
		arr := strings.Split(strArr[i][0], ",")
		ikey, _ := strconv.Atoi(arr[0])
		ival, _ := strconv.Atoi(arr[1])
		retdata[ikey] = ival
	}
	return retdata
}
