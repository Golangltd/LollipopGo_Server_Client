package _go

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
)

var Gitemsubtype_proto map[string]*STitemsubtype_proto

type STitemsubtype_proto struct {
	Sid      string
	Typename string
}

func init() {
	Gitemsubtype_proto = make(map[string]*STitemsubtype_proto)
}

func Inititemsubtype_proto() {
	L := lua.NewState()
	itemsubtype_prototmp := make(map[string]*STitemsubtype_proto)
	defer L.Close()
	err := L.DoFile("./lua/itemsubtype_proto.lua")
	if err != nil {
		fmt.Println("err", err.Error())
		panic(err)
	}
	itemsubtype_proto := L.GetGlobal("itemsubtype_proto")
	if tbl, ok := itemsubtype_proto.(*lua.LTable); ok {
		fmt.Println(L.ObjLen(tbl))
		tbl.ForEach(func(value lua.LValue, value2 lua.LValue) {
			data := new(STitemsubtype_proto)
			data.Sid = value.String()
			if tbl2, ok2 := value2.(*lua.LTable); ok2 {
				fmt.Println(L.ObjLen(tbl2))
				tbl2.ForEach(func(value3 lua.LValue, value4 lua.LValue) {
					if "typename" == value3.String() {
						data.Typename = value4.String()
					}
				})
			}
			itemsubtype_prototmp[data.Sid] = data
		})
	}
	Gitemsubtype_proto = itemsubtype_prototmp
}
