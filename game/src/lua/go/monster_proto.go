package _go

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
)

var Gmonster_proto map[string]*STmonster_proto

type STmonster_proto struct {
	Sid          string
	Modelid      string
	Level        string
	Monster_type string
	Skills       string
	Gift         string
	Attribute    string
}

func init() {
	Gmonster_proto = make(map[string]*STmonster_proto)
}

func Initmonster_proto() {
	L := lua.NewState()
	monster_prototmp := make(map[string]*STmonster_proto)
	defer L.Close()
	err := L.DoFile("./lua/monster_proto.lua")
	if err != nil {
		fmt.Println("err", err.Error())
		panic(err)
	}
	monster_proto := L.GetGlobal("monster_proto")
	if tbl, ok := monster_proto.(*lua.LTable); ok {
		fmt.Println(L.ObjLen(tbl))
		tbl.ForEach(func(value lua.LValue, value2 lua.LValue) {
			data := new(STmonster_proto)
			data.Sid = value.String()
			if tbl2, ok2 := value2.(*lua.LTable); ok2 {
				fmt.Println(L.ObjLen(tbl2))
				tbl2.ForEach(func(value3 lua.LValue, value4 lua.LValue) {
					if "modelid" == value3.String() {
						data.Modelid = value4.String()
					}
					if "level" == value3.String() {
						data.Level = value4.String()
					}
					if "monster_type" == value3.String() {
						data.Monster_type = value4.String()
					}
					if "skills" == value3.String() {
						data.Skills = value4.String()
					}
					if "gift" == value3.String() {
						data.Gift = value4.String()
					}
					if "attribute" == value3.String() {
						data.Attribute = value4.String()
					}
				})
			}
			monster_prototmp[data.Sid] = data
		})
	}
	Gmonster_proto = monster_prototmp
}
