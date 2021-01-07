package _go

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
)

var Gfunction_type_proto map[string]*STfunction_type_proto

type STfunction_type_proto struct {
	Sid    string
	Reward string
}

func init() {
	Gfunction_type_proto = make(map[string]*STfunction_type_proto)
}

func Initfunction_type_proto() {
	L := lua.NewState()
	function_type_prototmp := make(map[string]*STfunction_type_proto)
	defer L.Close()
	err := L.DoFile("./lua/function_type_proto.lua")
	if err != nil {
		fmt.Println("err", err.Error())
		panic(err)
	}
	function_type_proto := L.GetGlobal("function_type_proto")
	if tbl, ok := function_type_proto.(*lua.LTable); ok {
		fmt.Println(L.ObjLen(tbl))
		tbl.ForEach(func(value lua.LValue, value2 lua.LValue) {
			data := new(STfunction_type_proto)
			data.Sid = value.String()
			if tbl2, ok2 := value2.(*lua.LTable); ok2 {
				fmt.Println(L.ObjLen(tbl2))
				tbl2.ForEach(func(value3 lua.LValue, value4 lua.LValue) {
					if "reward" == value3.String() {
						data.Reward = value4.String()
					}
				})
			}
			function_type_prototmp[data.Sid] = data
		})
	}
	Gfunction_type_proto = function_type_prototmp
}
