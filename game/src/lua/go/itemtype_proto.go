package _go

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
)

var Gitemtype_proto map[string]*STitemtype_proto

type STitemtype_proto struct {
	Sid       string
	Type_name string
}

func init() {
	Gitemtype_proto = make(map[string]*STitemtype_proto)
}

func Inititemtype_proto() {
	L := lua.NewState()
	itemtype_prototmp := make(map[string]*STitemtype_proto)
	defer L.Close()
	err := L.DoFile("./lua/itemtype_proto.lua")
	if err != nil {
		fmt.Println("err", err.Error())
		panic(err)
	}
	itemtype_proto := L.GetGlobal("itemtype_proto")
	if tbl, ok := itemtype_proto.(*lua.LTable); ok {
		fmt.Println(L.ObjLen(tbl))
		tbl.ForEach(func(value lua.LValue, value2 lua.LValue) {
			data := new(STitemtype_proto)
			data.Sid = value.String()
			if tbl2, ok2 := value2.(*lua.LTable); ok2 {
				fmt.Println(L.ObjLen(tbl2))
				tbl2.ForEach(func(value3 lua.LValue, value4 lua.LValue) {
					if "type_name" == value3.String() {
						data.Type_name = value4.String()
					}
				})
			}
			itemtype_prototmp[data.Sid] = data
		})
	}
	Gitemtype_proto = itemtype_prototmp
}
