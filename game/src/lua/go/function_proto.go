package _go

import (
	"fmt"
	"github.com/golang/glog"
	lua "github.com/yuin/gopher-lua"
)

var Function_protoConfMap map[string]*Function_protoConfST

type Function_protoConfST struct {
	Sid        string
	Name       string
	Condition1 string
	Reward     string
}

func init() {
	Function_protoConfMap = make(map[string]*Function_protoConfST)
	initfunction_protoConf()
}

func initfunction_protoConf() {
	L := lua.NewState()
	function_protoTmp := make(map[string]*Function_protoConfST)
	defer L.Close()
	err := L.DoFile("./lua/function_proto.lua")
	if err != nil {
		fmt.Println("err", err.Error())
		panic(err)
	}
	luaTable := L.GetGlobal("function_proto")
	if tbl, ok := luaTable.(*lua.LTable); ok {
		tbl.ForEach(func(value lua.LValue, value2 lua.LValue) {
			data := new(Function_protoConfST)
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
			function_protoTmp[data.Sid] = data
		})
	}
	Function_protoConfMap = function_protoTmp
	glog.Infof("load function_protoConfST size:%d", len(Function_protoConfMap))
}
