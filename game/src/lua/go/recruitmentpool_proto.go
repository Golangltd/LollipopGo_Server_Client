package _go

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
)

var Grecruitmentpool_proto map[string]*STrecruitmentpool_proto

type STrecruitmentpool_proto struct {
	Sid     string
	Pool_id string
	Cards   string
	Quality string
	Weight  string
	Up      string
}

func init() {
	Grecruitmentpool_proto = make(map[string]*STrecruitmentpool_proto)
	Initrecruitmentpool_proto()
}

func Initrecruitmentpool_proto() {
	L := lua.NewState()
	recruitmentpool_prototmp := make(map[string]*STrecruitmentpool_proto)
	defer L.Close()
	err := L.DoFile("./lua/recruitmentpool_proto.lua")
	if err != nil {
		fmt.Println("err", err.Error())
		panic(err)
	}
	recruitmentpool_proto := L.GetGlobal("recruitmentpool_proto")
	if tbl, ok := recruitmentpool_proto.(*lua.LTable); ok {
		tbl.ForEach(func(value lua.LValue, value2 lua.LValue) {
			data := new(STrecruitmentpool_proto)
			data.Sid = value.String()
			if tbl2, ok2 := value2.(*lua.LTable); ok2 {
				tbl2.ForEach(func(value3 lua.LValue, value4 lua.LValue) {
					if "pool_id" == value3.String() {
						data.Pool_id = luaValueToString(value4)
					}
					if "cards" == value3.String() {
						data.Cards = luaValueToString(value4)
					}
					if "quality" == value3.String() {
						data.Quality = luaValueToString(value4)
					}
					if "weight" == value3.String() {
						data.Weight = luaValueToString(value4)
					}
					if "up" == value3.String() {
						data.Up = luaValueToString(value4)
					}
				})
			}
			recruitmentpool_prototmp[data.Sid] = data
		})
	}
	Grecruitmentpool_proto = recruitmentpool_prototmp
}
