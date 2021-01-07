package tables

import (
	"LollipopGo/log"
	gameDB "LollipopGo2.8x/data"
	twLibDBTable "github.com/Golangltd/Twlib/dbtable"
	"strconv"
	"sync"
)

var (
	ActivityTable map[int]*ActivitySt
	activityMsg   []interface{}
	activityOnce  sync.Once
)

// 活动(activity)配置
type ActivitySt struct {
	Sid       string
	OpenType  string
	BeginTime string
	EndTime   string
	CloseTime string
	Interval  string
	Rewards   string
	Condition string
	Complete  string
}

// 读取章节数据
func GetActivityInfo() {
	data := gameDB.GetCFGameData(twLibDBTable.Gl_gk_zjpzb)
	if data == nil {
		log.Fatal("error! receive ActivityTable is nil!!")
	}
	activityMsg = data
	activityOnce.Do(initActivityTables)
}

func initActivityTables() {
	ActivityTable = make(map[int]*ActivitySt)
	for _, sample := range activityMsg {
		temp := sample.(map[string]interface{})
		m := &ActivitySt{}
		m.Sid = temp["sid"].(string)
		// 把每个技能对象存到map中
		sid, _ := strconv.Atoi(m.Sid)
		ActivityTable[sid] = m
	}
	activityMsg = nil
}
