package tables

import (
	"LollipopGo/log"
	gameDB "LollipopGo2.8x/data"
	twLibDBTable "github.com/Golangltd/Twlib/dbtable"
	"strconv"
	"sync"
)

var (
	CardExpTable map[int]*CardExpSt
	cardExpMsg   []interface{}
	cardExpOnce  sync.Once
)

// 巫师经验表
type CardExpSt struct {
	Sid        string // 自然唯一ID
	Cost       string // 升级花费     for 3
	LevelUp    string // 是否突破     1:是突破
	Attribute1 string // 战士等级成长属性 for 2
	Attribute2 string // 游侠等级成长属性 for 2
	Attribute3 string // 法师等级成长属性 for 2
	Attribute4 string // 辅助等级成长属性 for 2
	Attribute5 string // 战士突破成长额外属性 for 2
	Attribute6 string // 游侠突破成长额外属性 for 2
	Attribute7 string // 法师突破成长额外属性 for 2
	Attribute8 string // 辅助突破成长额外属性 for 2
}

func AddCardExpTable() {
	data := gameDB.GetCFGameData(twLibDBTable.Gl_wse_wsjyb)
	if data == nil {
		log.Fatal("error! receive CardExpTable is nil!!")
	}
	cardExpMsg = data
	cardExpOnce.Do(initCardExpTable)
}

func initCardExpTable() {
	CardExpTable = make(map[int]*CardExpSt)
	for _, sample := range cardExpMsg {
		temp := sample.(map[string]interface{})
		m := &CardExpSt{}
		m.Sid = temp["sid"].(string)
		m.Cost = temp["cost"].(string)
		m.LevelUp = temp["levelup"].(string)
		m.Attribute1 = temp["attribute1"].(string)
		m.Attribute2 = temp["attribute2"].(string)
		m.Attribute3 = temp["attribute3"].(string)
		m.Attribute4 = temp["attribute4"].(string)
		m.Attribute5 = temp["attribute5"].(string) // 数据操作，主要应用操作！！！
		m.Attribute6 = temp["attribute6"].(string)
		m.Attribute7 = temp["attribute7"].(string)
		m.Attribute8 = temp["attribute8"].(string)
		sid, _ := strconv.Atoi(m.Sid)
		CardExpTable[sid] = m
	}
	cardExpMsg = nil
}
