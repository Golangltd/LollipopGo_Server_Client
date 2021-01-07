package _go

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
)

var Gskill map[string]*STskill

type STskill struct {
	Sid           string
	Cooldown      string
	Spexpend      string
	Skill_type    string
	Basic_type    string
	Dmgparameter1 string
	Dmgparameter2 string
	Target        string
	Trigger       string
	Effect        string
	Extratarget   string
	Maxtarget     string
}

func init() {
	Gskill = make(map[string]*STskill)
}

func Initskill() {
	L := lua.NewState()
	skilltmp := make(map[string]*STskill)
	defer L.Close()
	err := L.DoFile("./lua/skill.lua")
	if err != nil {
		fmt.Println("err", err.Error())
		panic(err)
	}
	skill := L.GetGlobal("skill")
	if tbl, ok := skill.(*lua.LTable); ok {
		fmt.Println(L.ObjLen(tbl))
		tbl.ForEach(func(value lua.LValue, value2 lua.LValue) {
			data := new(STskill)
			data.Sid = value.String()
			if tbl2, ok2 := value2.(*lua.LTable); ok2 {
				fmt.Println(L.ObjLen(tbl2))
				tbl2.ForEach(func(value3 lua.LValue, value4 lua.LValue) {
					if "cooldown" == value3.String() {
						data.Cooldown = value4.String()
					}
					if "spexpend" == value3.String() {
						data.Spexpend = value4.String()
					}
					if "skill_type" == value3.String() {
						data.Skill_type = value4.String()
					}
					if "basic_type" == value3.String() {
						data.Basic_type = value4.String()
					}
					if "dmgparameter1" == value3.String() {
						data.Dmgparameter1 = value4.String()
					}
					if "dmgparameter2" == value3.String() {
						data.Dmgparameter2 = value4.String()
					}
					if "target" == value3.String() {
						data.Target = value4.String()
					}
					if "trigger" == value3.String() {
						data.Trigger = value4.String()
					}
					if "effect" == value3.String() {
						data.Effect = value4.String()
					}
					if "extratarget" == value3.String() {
						data.Extratarget = value4.String()
					}
					if "maxtarget" == value3.String() {
						data.Maxtarget = value4.String()
					}
				})
			}
			skilltmp[data.Sid] = data
		})
	}
	Gskill = skilltmp
}
