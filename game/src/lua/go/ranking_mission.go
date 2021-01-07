package _go

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
)

var Granking_mission map[string]*STranking_mission

type STranking_mission struct {
	Sid          string
	Mission_type string
	Condition    string
	Reward       string
}

func init() {
	Granking_mission = make(map[string]*STranking_mission)
}

func Initranking_mission() {
	L := lua.NewState()
	ranking_missiontmp := make(map[string]*STranking_mission)
	defer L.Close()
	err := L.DoFile("./lua/ranking_mission.lua")
	if err != nil {
		fmt.Println("err", err.Error())
		panic(err)
	}
	ranking_mission := L.GetGlobal("ranking_mission")
	if tbl, ok := ranking_mission.(*lua.LTable); ok {
		fmt.Println(L.ObjLen(tbl))
		tbl.ForEach(func(value lua.LValue, value2 lua.LValue) {
			data := new(STranking_mission)
			data.Sid = value.String()
			if tbl2, ok2 := value2.(*lua.LTable); ok2 {
				fmt.Println(L.ObjLen(tbl2))
				tbl2.ForEach(func(value3 lua.LValue, value4 lua.LValue) {
					if "mission_type" == value3.String() {
						data.Mission_type = value4.String()
					}
					if "condition" == value3.String() {
						data.Condition = value4.String()
					}
					if "reward" == value3.String() {
						data.Reward = value4.String()
					}
				})
			}
			ranking_missiontmp[data.Sid] = data
		})
	}
	Granking_mission = ranking_missiontmp
}
