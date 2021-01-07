package _go

import (
	"fmt"
	"github.com/golang/glog"
	lua "github.com/yuin/gopher-lua"
)

//功能开启配置表
var FunctionStartConf map[string]*FunctionStartConfST

type FunctionStartConfST struct {
	Sid        string
	Name       string
	Condition1 string
	Reward     string
}

func init() {
	FunctionStartConf = make(map[string]*FunctionStartConfST)
	initFunctionStartConf()
}

func initFunctionStartConf() {
	L := lua.NewState()
	functionstartTmp := make(map[string]*FunctionStartConfST)
	defer L.Close()
	err := L.DoFile("./lua/function_proto.lua")
	if err != nil {
		fmt.Println("err", err.Error())
		panic(err)
	}
	luaTable := L.GetGlobal("function_proto")
	if tbl, ok := luaTable.(*lua.LTable); ok {
		tbl.ForEach(func(value lua.LValue, value2 lua.LValue) {
			data := new(FunctionStartConfST)
			data.Sid = value.String()
			if tbl2, ok2 := value2.(*lua.LTable); ok2 {
				tbl2.ForEach(func(value3 lua.LValue, value4 lua.LValue) {
					if "name" == value3.String() {
						data.Name = luaValueToString(value4)
					}
					if "condition1" == value3.String() {
						data.Condition1 = luaValueToString(value4)
					}
					if "reward" == value3.String() {
						data.Reward = luaValueToString(value4)
					}
				})
			}
			functionstartTmp[data.Sid] = data
		})
	}
	FunctionStartConf = functionstartTmp
	glog.Infof("load FunctionStartConfST size:%d", len(FunctionStartConf))
}
