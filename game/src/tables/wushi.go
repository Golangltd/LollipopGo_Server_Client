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
	WuShiTables map[int]*WuShiSt
	wuShiMsg    []interface{}
	wuShiOnce   sync.Once
)

// 巫师经验表
type WuShiSt struct {
	Sid        string
	Cost       string // 花费
	LevelUp    string // 是否突破
	Attribute1 string // 战士等级成长属性
	Attribute2 string // 游侠等级成长属性
	Attribute3 string // 法师等级成长属性
	Attribute4 string // 辅助等级成长属性
	Attribute5 string // 战士【突破】额外属性
	Attribute6 string // 游侠【突破】额外属性
	Attribute7 string // 法师【突破】额外属性
	Attribute8 string // 辅助【突破】额外属性
}

// 读取章节数据
func GetWuShiInfo() {
	data := gamedb.GetCFGameData(twlib_dbtable.Gl_wse_wsjyb)
	AddWuShiTables(data)
}

func AddWuShiTables(msg []interface{}) {
	if msg == nil {
		log.Fatal("error! receive skillTables is nil!!")
	}
	wuShiMsg = msg
	wuShiOnce.Do(initWuShiTables)
}

func initWuShiTables() {
	WuShiTables = make(map[int]*WuShiSt)
	for _, sample := range wuShiMsg {
		temp := sample.(map[string]interface{})
		m := &WuShiSt{}
		m.Sid = temp["sid"].(string)
		m.Cost = temp["cost"].(string)
		m.LevelUp = temp["levelup"].(string)
		m.Attribute1 = temp["attribute1"].(string)
		m.Attribute2 = temp["attribute2"].(string)
		m.Attribute3 = temp["attribute3"].(string)
		m.Attribute4 = temp["attribute4"].(string)
		m.Attribute5 = temp["attribute5"].(string)
		m.Attribute6 = temp["attribute6"].(string)
		m.Attribute7 = temp["attribute7"].(string)
		m.Attribute8 = temp["attribute8"].(string)
		sid, _ := strconv.Atoi(m.Sid)
		WuShiTables[sid] = m
	}
}

// 获取奖励
func GetWuShiFor2Row(data string) map[int]int {
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
func GetWuShiFor3Row(data string) map[int]int {
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
