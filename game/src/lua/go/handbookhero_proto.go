package _go

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
)

var Ghandbookhero_proto map[string]*SThandbookhero_proto

type SThandbookhero_proto struct {
	Sid          string
	Handname     string
	Character_id string
	Award        string
}

func init() {
	Ghandbookhero_proto = make(map[string]*SThandbookhero_proto)
}

func Inithandbookhero_proto() {
	L := lua.NewState()
	handbookhero_prototmp := make(map[string]*SThandbookhero_proto)
	defer L.Close()
	err := L.DoFile("./lua/handbookhero_proto.lua")
	if err != nil {
		fmt.Println("err", err.Error())
		panic(err)
	}
	handbookhero_proto := L.GetGlobal("handbookhero_proto")
	if tbl, ok := handbookhero_proto.(*lua.LTable); ok {
		fmt.Println(L.ObjLen(tbl))
		tbl.ForEach(func(value lua.LValue, value2 lua.LValue) {
			data := new(SThandbookhero_proto)
			data.Sid = value.String()
			if tbl2, ok2 := value2.(*lua.LTable); ok2 {
				fmt.Println(L.ObjLen(tbl2))
				tbl2.ForEach(func(value3 lua.LValue, value4 lua.LValue) {
					if "handname" == value3.String() {
						data.Handname = value4.String()
					}
					if "character_id" == value3.String() {
						data.Character_id = value4.String()
					}
					if "award" == value3.String() {
						data.Award = value4.String()
					}
				})
			}
			handbookhero_prototmp[data.Sid] = data
		})
	}
	Ghandbookhero_proto = handbookhero_prototmp
}
