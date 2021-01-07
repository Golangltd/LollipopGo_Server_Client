package _go

import (
	"fmt"
	"github.com/golang/glog"
	lua "github.com/yuin/gopher-lua"
)

var ExamineConf map[string]*ExamineConfST

type ExamineConfST struct {
	Sid            string
	Type           string
	Turn           string
	Level          string
	Score          string
	Show_condition string
	Get_condition  string
	Award_hight    string
	Award_common   string
}

func init() {
	ExamineConf = make(map[string]*ExamineConfST)
	initExamineConf()
}

func initExamineConf() {
	L := lua.NewState()
	examineTmp := make(map[string]*ExamineConfST)
	defer L.Close()
	err := L.DoFile("./lua/examine_proto.lua")
	if err != nil {
		fmt.Println("err", err.Error())
		panic(err)
	}
	luaTable := L.GetGlobal("examine_proto")
	if tbl, ok := luaTable.(*lua.LTable); ok {
		tbl.ForEach(func(value lua.LValue, value2 lua.LValue) {
			data := new(ExamineConfST)
			data.Sid = value.String()
			if tbl2, ok2 := value2.(*lua.LTable); ok2 {
				tbl2.ForEach(func(value3 lua.LValue, value4 lua.LValue) {
					if "type" == value3.String() {
						data.Type = luaValueToString(value4)
					}
					if "turn" == value3.String() {
						data.Turn = luaValueToString(value4)
					}
					if "level" == value3.String() {
						data.Level = luaValueToString(value4)
					}
					if "score" == value3.String() {
						data.Score = luaValueToString(value4)
					}
					if "show_condition" == value3.String() {
						data.Show_condition = luaValueToString(value4)
					}
					if "get_condition" == value3.String() {
						data.Get_condition = luaValueToString(value4)
					}
					if "award_hight" == value3.String() {
						data.Award_hight = luaValueToString(value4)
					}
					if "award_common" == value3.String() {
						data.Award_common = luaValueToString(value4)
					}
				})
			}
			examineTmp[data.Sid] = data
		})
	}
	ExamineConf = examineTmp
	glog.Infof("load ExamineConfST size:%d", len(ExamineConf))
}
