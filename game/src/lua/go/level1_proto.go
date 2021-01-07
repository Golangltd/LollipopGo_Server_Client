package _go

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
)

var Glevel1_proto map[string]*STlevel1_proto

type STlevel1_proto struct {
	Sid           string
	Next_id       string
	Level_type    string
	Round_max     string
	Skip_round    string
	Monster_group string
	Award         string
	Install_award string
}

func init() {
	Glevel1_proto = make(map[string]*STlevel1_proto)
}

func Initlevel1_proto() {
	L := lua.NewState()
	level1_prototmp := make(map[string]*STlevel1_proto)
	defer L.Close()
	err := L.DoFile("./lua/level1_proto.lua")
	if err != nil {
		fmt.Println("err", err.Error())
		panic(err)
	}
	level1_proto := L.GetGlobal("level1_proto")
	if tbl, ok := level1_proto.(*lua.LTable); ok {
		fmt.Println(L.ObjLen(tbl))
		tbl.ForEach(func(value lua.LValue, value2 lua.LValue) {
			data := new(STlevel1_proto)
			data.Sid = value.String()
			if tbl2, ok2 := value2.(*lua.LTable); ok2 {
				fmt.Println(L.ObjLen(tbl2))
				tbl2.ForEach(func(value3 lua.LValue, value4 lua.LValue) {
					if "next_id" == value3.String() {
						data.Next_id = value4.String()
					}
					if "level_type" == value3.String() {
						data.Level_type = value4.String()
					}
					if "round_max" == value3.String() {
						data.Round_max = value4.String()
					}
					if "skip_round" == value3.String() {
						data.Skip_round = value4.String()
					}
					if "monster_group" == value3.String() {
						data.Monster_group = value4.String()
					}
					if "award" == value3.String() {
						data.Award = value4.String()
					}
					if "install_award" == value3.String() {
						data.Install_award = value4.String()
					}
				})
			}
			level1_prototmp[data.Sid] = data
		})
	}
	Glevel1_proto = level1_prototmp
}
