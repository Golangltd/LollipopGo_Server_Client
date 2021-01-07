package _go

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
)

var Gskill_trigger map[string]*STskill_trigger

type STskill_trigger struct {
	Sid         string
	Name        string
	Description string
}

func init() {
	Gskill_trigger = make(map[string]*STskill_trigger)
}

func Initskill_trigger() {
	L := lua.NewState()
	skill_triggertmp := make(map[string]*STskill_trigger)
	defer L.Close()
	err := L.DoFile("./lua/skill_trigger.lua")
	if err != nil {
		fmt.Println("err", err.Error())
		panic(err)
	}
	skill_trigger := L.GetGlobal("skill_trigger")
	if tbl, ok := skill_trigger.(*lua.LTable); ok {
		fmt.Println(L.ObjLen(tbl))
		tbl.ForEach(func(value lua.LValue, value2 lua.LValue) {
			data := new(STskill_trigger)
			data.Sid = value.String()
			if tbl2, ok2 := value2.(*lua.LTable); ok2 {
				fmt.Println(L.ObjLen(tbl2))
				tbl2.ForEach(func(value3 lua.LValue, value4 lua.LValue) {
					if "name" == value3.String() {
						data.Name = value4.String()
					}
					if "description" == value3.String() {
						data.Description = value4.String()
					}
				})
			}
			skill_triggertmp[data.Sid] = data
		})
	}
	Gskill_trigger = skill_triggertmp
}
