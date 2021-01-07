package _go

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
)

var Gwandcondition_proto map[string]*STwandcondition_proto

type STwandcondition_proto struct {
	Sid           string
	Struse_item   string
	Str_attribute string
}

func init() {
	Gwandcondition_proto = make(map[string]*STwandcondition_proto)
}

func Initwandcondition_proto() {
	L := lua.NewState()
	wandcondition_prototmp := make(map[string]*STwandcondition_proto)
	defer L.Close()
	err := L.DoFile("./lua/wandcondition_proto.lua")
	if err != nil {
		fmt.Println("err", err.Error())
		panic(err)
	}
	wandcondition_proto := L.GetGlobal("wandcondition_proto")
	if tbl, ok := wandcondition_proto.(*lua.LTable); ok {
		fmt.Println(L.ObjLen(tbl))
		tbl.ForEach(func(value lua.LValue, value2 lua.LValue) {
			data := new(STwandcondition_proto)
			data.Sid = value.String()
			if tbl2, ok2 := value2.(*lua.LTable); ok2 {
				fmt.Println(L.ObjLen(tbl2))
				tbl2.ForEach(func(value3 lua.LValue, value4 lua.LValue) {
					if "struse_item" == value3.String() {
						data.Struse_item = value4.String()
					}
					if "str_attribute" == value3.String() {
						data.Str_attribute = value4.String()
					}
				})
			}
			wandcondition_prototmp[data.Sid] = data
		})
	}
	Gwandcondition_proto = wandcondition_prototmp
}
