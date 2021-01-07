package _go

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
)

var Gskill_effect map[string]*STskill_effect

type STskill_effect struct {
	Sid        string
	Type       string
	Time       string
	Maxnum     string
	Isdispel   string
	Parameter1 string
}

func init() {
	Gskill_effect = make(map[string]*STskill_effect)
}

func Initskill_effect() {
	L := lua.NewState()
	skill_effecttmp := make(map[string]*STskill_effect)
	defer L.Close()
	err := L.DoFile("./lua/skill_effect.lua")
	if err != nil {
		fmt.Println("err", err.Error())
		panic(err)
	}
	skill_effect := L.GetGlobal("skill_effect")
	if tbl, ok := skill_effect.(*lua.LTable); ok {
		fmt.Println(L.ObjLen(tbl))
		tbl.ForEach(func(value lua.LValue, value2 lua.LValue) {
			data := new(STskill_effect)
			data.Sid = value.String()
			if tbl2, ok2 := value2.(*lua.LTable); ok2 {
				fmt.Println(L.ObjLen(tbl2))
				tbl2.ForEach(func(value3 lua.LValue, value4 lua.LValue) {
					if "type" == value3.String() {
						data.Type = value4.String()
					}
					if "time" == value3.String() {
						data.Time = value4.String()
					}
					if "maxnum" == value3.String() {
						data.Maxnum = value4.String()
					}
					if "isdispel" == value3.String() {
						data.Isdispel = value4.String()
					}
					if "parameter1" == value3.String() {
						data.Parameter1 = value4.String()
					}
				})
			}
			skill_effecttmp[data.Sid] = data
		})
	}
	Gskill_effect = skill_effecttmp
}
