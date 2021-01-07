package _go

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
)

var Gtalent_proto map[string]*STtalent_proto

type STtalent_proto struct {
	Sid     string
	Type    string
	Class   string
	Level   string
	Exp     string
	Per_exp string
	Cost    string
}

func init() {
	Gtalent_proto = make(map[string]*STtalent_proto)
}

func Inittalent_proto() {
	L := lua.NewState()
	talent_prototmp := make(map[string]*STtalent_proto)
	defer L.Close()
	err := L.DoFile("./lua/talent_proto.lua")
	if err != nil {
		fmt.Println("err", err.Error())
		panic(err)
	}
	talent_proto := L.GetGlobal("talent_proto")
	if tbl, ok := talent_proto.(*lua.LTable); ok {
		fmt.Println(L.ObjLen(tbl))
		tbl.ForEach(func(value lua.LValue, value2 lua.LValue) {
			data := new(STtalent_proto)
			data.Sid = value.String()
			if tbl2, ok2 := value2.(*lua.LTable); ok2 {
				fmt.Println(L.ObjLen(tbl2))
				tbl2.ForEach(func(value3 lua.LValue, value4 lua.LValue) {
					if "type" == value3.String() {
						data.Type = value4.String()
					}
					if "class" == value3.String() {
						data.Class = value4.String()
					}
					if "level" == value3.String() {
						data.Level = value4.String()
					}
					if "exp" == value3.String() {
						data.Exp = value4.String()
					}
					if "per_exp" == value3.String() {
						data.Per_exp = value4.String()
					}
					if "cost" == value3.String() {
						data.Cost = value4.String()
					}
				})
			}
			talent_prototmp[data.Sid] = data
		})
	}
	Gtalent_proto = talent_prototmp
}
