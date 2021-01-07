package _go

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
)

var Gledermodle_proto map[string]*STledermodle_proto

type STledermodle_proto struct {
	Sid   string
	Model string
}

func init() {
	Gledermodle_proto = make(map[string]*STledermodle_proto)
}

func Initledermodle_proto() {
	L := lua.NewState()
	ledermodle_prototmp := make(map[string]*STledermodle_proto)
	defer L.Close()
	err := L.DoFile("./lua/ledermodle_proto.lua")
	if err != nil {
		fmt.Println("err", err.Error())
		panic(err)
	}
	ledermodle_proto := L.GetGlobal("ledermodle_proto")
	if tbl, ok := ledermodle_proto.(*lua.LTable); ok {
		fmt.Println(L.ObjLen(tbl))
		tbl.ForEach(func(value lua.LValue, value2 lua.LValue) {
			data := new(STledermodle_proto)
			data.Sid = value.String()
			if tbl2, ok2 := value2.(*lua.LTable); ok2 {
				fmt.Println(L.ObjLen(tbl2))
				tbl2.ForEach(func(value3 lua.LValue, value4 lua.LValue) {
					if "model" == value3.String() {
						data.Model = value4.String()
					}
				})
			}
			ledermodle_prototmp[data.Sid] = data
		})
	}
	Gledermodle_proto = ledermodle_prototmp
}
