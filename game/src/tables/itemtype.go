package tables

import (
	"LollipopGo/log"
	gameDB "LollipopGo2.8x/data"
	twLibDBTable "github.com/Golangltd/Twlib/dbtable"
	"strconv"
	"sync"
)

// 道具类型表
type ItemTypeSt struct {
	Sid       string // 道具Id
	ParaMeter string // 参数
	FuncId    string // 功能Id
}

var (
	ItemTypeTable map[int]*ItemTypeSt
	itemTypeMsg   []interface{}
	itemTypeOnce  sync.Once
)

// 读取章节数据
func AddItemTypeTable() {
	data := gameDB.GetCFGameData(twLibDBTable.Gl_dj_djlxb)
	if data == nil {
		log.Fatal("error! receive ItemTypeTable is nil!!")
	}
	itemTypeMsg = data
	itemTypeOnce.Do(initItemTypeTable)
}

func initItemTypeTable() {
	ItemTypeTable = make(map[int]*ItemTypeSt)
	for _, sample := range itemTypeMsg {
		temp := sample.(map[string]interface{})
		m := &ItemTypeSt{}
		m.Sid = temp["sid"].(string)
		m.ParaMeter = temp["parameter"].(string)
		m.FuncId = temp["funcid"].(string)
		sid, _ := strconv.Atoi(m.Sid)
		ItemTypeTable[sid] = m
	}
}
