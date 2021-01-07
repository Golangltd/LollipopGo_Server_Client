package _go

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
)

var Gskill_target map[string]*STskill_target

type STskill_target struct {
	Sid        string
	Type       string
	Parameter1 string
	Parameter2 string
	Parameter3 string
	Parameter4 string
}

func init() {
	Gskill_target = make(map[string]*STskill_target)
}

func Initskill_target() {
	L := lua.NewState()
	skill_targettmp := make(map[string]*STskill_target)
	defer L.Close()
	err := L.DoFile("./lua/skill_target.lua")
	if err != nil {
		fmt.Println("err", err.Error())
		panic(err)
	}
	skill_target := L.GetGlobal("skill_target")
	if tbl, ok := skill_target.(*lua.LTable); ok {
		fmt.Println(L.ObjLen(tbl))
		tbl.ForEach(func(value lua.LValue, value2 lua.LValue) {
			data := new(STskill_target)
			data.Sid = value.String()
			if tbl2, ok2 := value2.(*lua.LTable); ok2 {
				fmt.Println(L.ObjLen(tbl2))
				tbl2.ForEach(func(value3 lua.LValue, value4 lua.LValue) {
					if "type" == value3.String() {
						data.Type = value4.String()
					}
					if "parameter1" == value3.String() {
						data.Parameter1 = value4.String()
					}
					if "parameter2" == value3.String() {
						data.Parameter2 = value4.String()
					}
					if "parameter3" == value3.String() {
						data.Parameter3 = value4.String()
					}
					if "parameter4" == value3.String() {
						data.Parameter4 = value4.String()
					}
				})
			}
			skill_targettmp[data.Sid] = data
		})
	}
	Gskill_target = skill_targettmp
}
