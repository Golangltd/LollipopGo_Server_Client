package _go

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
)

var Gmission_proto map[string]*STmission_proto

type STmission_proto struct {
	Sid         string
	Type        string
	Sontype     string
	Missiontype string
	Param       string
	Reward      string
	Funcid      string
	Times       string
}

func init() {
	Gmission_proto = make(map[string]*STmission_proto)
}

func Initmission_proto() {
	L := lua.NewState()
	mission_prototmp := make(map[string]*STmission_proto)
	defer L.Close()
	err := L.DoFile("./lua/mission_proto.lua")
	if err != nil {
		fmt.Println("err", err.Error())
		panic(err)
	}
	mission_proto := L.GetGlobal("mission_proto")
	if tbl, ok := mission_proto.(*lua.LTable); ok {
		fmt.Println(L.ObjLen(tbl))
		tbl.ForEach(func(value lua.LValue, value2 lua.LValue) {
			data := new(STmission_proto)
			data.Sid = value.String()
			if tbl2, ok2 := value2.(*lua.LTable); ok2 {
				fmt.Println(L.ObjLen(tbl2))
				tbl2.ForEach(func(value3 lua.LValue, value4 lua.LValue) {
					if "type" == value3.String() {
						data.Type = value4.String()
					}
					if "sontype" == value3.String() {
						data.Sontype = value4.String()
					}
					if "missiontype" == value3.String() {
						data.Missiontype = value4.String()
					}
					if "param" == value3.String() {
						data.Param = value4.String()
					}
					if "reward" == value3.String() {
						data.Reward = value4.String()
					}
					if "funcid" == value3.String() {
						data.Funcid = value4.String()
					}
					if "times" == value3.String() {
						data.Times = value4.String()
					}
				})
			}
			mission_prototmp[data.Sid] = data
		})
	}
	Gmission_proto = mission_prototmp
}
