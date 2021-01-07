package _go

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
)

var Gtips_proto map[string]*STtips_proto

type STtips_proto struct {
	Sid       string
	Tips_type string
}

func init() {
	Gtips_proto = make(map[string]*STtips_proto)
}

func Inittips_proto() {
	L := lua.NewState()
	tips_prototmp := make(map[string]*STtips_proto)
	defer L.Close()
	err := L.DoFile("./lua/tips_proto.lua")
	if err != nil {
		fmt.Println("err", err.Error())
		panic(err)
	}
	tips_proto := L.GetGlobal("tips_proto")
	if tbl, ok := tips_proto.(*lua.LTable); ok {
		fmt.Println(L.ObjLen(tbl))
		tbl.ForEach(func(value lua.LValue, value2 lua.LValue) {
			data := new(STtips_proto)
			data.Sid = value.String()
			if tbl2, ok2 := value2.(*lua.LTable); ok2 {
				fmt.Println(L.ObjLen(tbl2))
				tbl2.ForEach(func(value3 lua.LValue, value4 lua.LValue) {
					if "tips_type" == value3.String() {
						data.Tips_type = value4.String()
					}
				})
			}
			tips_prototmp[data.Sid] = data
		})
	}
	Gtips_proto = tips_prototmp
}
