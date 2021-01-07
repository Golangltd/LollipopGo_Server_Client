package _go

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
)

var Gmonstergroup_proto map[string]*STmonstergroup_proto

type STmonstergroup_proto struct {
	Sid      string
	Type     string
	Monsters string
}

func init() {
	Gmonstergroup_proto = make(map[string]*STmonstergroup_proto)
}

func Initmonstergroup_proto() {
	L := lua.NewState()
	monstergroup_prototmp := make(map[string]*STmonstergroup_proto)
	defer L.Close()
	err := L.DoFile("./lua/monstergroup_proto.lua")
	if err != nil {
		fmt.Println("err", err.Error())
		panic(err)
	}
	monstergroup_proto := L.GetGlobal("monstergroup_proto")
	if tbl, ok := monstergroup_proto.(*lua.LTable); ok {
		fmt.Println(L.ObjLen(tbl))
		tbl.ForEach(func(value lua.LValue, value2 lua.LValue) {
			data := new(STmonstergroup_proto)
			data.Sid = value.String()
			if tbl2, ok2 := value2.(*lua.LTable); ok2 {
				fmt.Println(L.ObjLen(tbl2))
				tbl2.ForEach(func(value3 lua.LValue, value4 lua.LValue) {
					if "type" == value3.String() {
						data.Type = value4.String()
					}
					if "monsters" == value3.String() {
						data.Monsters = value4.String()
					}
				})
			}
			monstergroup_prototmp[data.Sid] = data
		})
	}
	Gmonstergroup_proto = monstergroup_prototmp
}
