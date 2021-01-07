package _go

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
)

var Gplayername map[string]*STplayername

type STplayername struct {
	Sid        string
	Chs_F_name string
	Chs_M_name string
	Eng_F_name string
	Eng_M_name string
}

func init() {
	Gplayername = make(map[string]*STplayername)
}

func Initplayername() {
	L := lua.NewState()
	playernametmp := make(map[string]*STplayername)
	defer L.Close()
	err := L.DoFile("./lua/playername.lua")
	if err != nil {
		fmt.Println("err", err.Error())
		panic(err)
	}
	playername := L.GetGlobal("playername")
	if tbl, ok := playername.(*lua.LTable); ok {
		fmt.Println(L.ObjLen(tbl))
		tbl.ForEach(func(value lua.LValue, value2 lua.LValue) {
			data := new(STplayername)
			data.Sid = value.String()
			if tbl2, ok2 := value2.(*lua.LTable); ok2 {
				fmt.Println(L.ObjLen(tbl2))
				tbl2.ForEach(func(value3 lua.LValue, value4 lua.LValue) {
					if "chs_F_name" == value3.String() {
						data.Chs_F_name = value4.String()
					}
					if "chs_M_name" == value3.String() {
						data.Chs_M_name = value4.String()
					}
					if "eng_F_name" == value3.String() {
						data.Eng_F_name = value4.String()
					}
					if "eng_M_name" == value3.String() {
						data.Eng_M_name = value4.String()
					}
				})
			}
			playernametmp[data.Sid] = data
		})
	}
	Gplayername = playernametmp
}
