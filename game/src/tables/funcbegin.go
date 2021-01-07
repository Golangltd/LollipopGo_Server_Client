package tables

import (
	"LollipopGo/log"
	gameDB "LollipopGo2.8x/data"
	twLibDBTable "github.com/Golangltd/Twlib/dbtable"
	"strconv"
	"sync"
)

var (
	FuncBeginTable map[int]*FuncBeginSt
	funcBeginMsg   []interface{}
	funcBeginOnce  sync.Once
)

type FuncBeginSt struct {
	Sid        string
	Name       string // 章节的Id信息，因为文本表
	Condition1 string // 子章节配置
	Reward     string // 关卡数据
}

// 读取章节数据
func AddFuncBeginTable() {
	data := gameDB.GetCFGameData(twLibDBTable.Gl_gk_zjpzb)
	if data == nil {
		log.Fatal("error! receive FuncBeginTable is nil!!")
	}
	funcBeginMsg = data
	funcBeginOnce.Do(initFuncBeginTable)
}

func initFuncBeginTable() {
	FuncBeginTable = make(map[int]*FuncBeginSt)
	for _, sample := range funcBeginMsg {
		temp := sample.(map[string]interface{})
		m := &FuncBeginSt{}
		m.Sid = temp["sid"].(string)
		m.Name = temp["name"].(string)
		m.Condition1 = temp["condition1"].(string)
		m.Reward = temp["reward"].(string)
		sid, _ := strconv.Atoi(m.Sid)
		FuncBeginTable[sid] = m
	}
	funcBeginMsg = nil
}
