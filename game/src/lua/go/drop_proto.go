package _go

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
)

var Gdrop_proto map[string]*STdrop_proto

type STdrop_proto struct {
	Sid   string
	Coin  string
	Item1 string
	Item2 string
}

func init() {
	Gdrop_proto = make(map[string]*STdrop_proto)
}

func Initdrop_proto() {
	L := lua.NewState()
	drop_prototmp := make(map[string]*STdrop_proto)
	defer L.Close()
	err := L.DoFile("./lua/drop_proto.lua")
	if err != nil {
		fmt.Println("err", err.Error())
		panic(err)
	}
	drop_proto := L.GetGlobal("drop_proto")
	if tbl, ok := drop_proto.(*lua.LTable); ok {
		fmt.Println(L.ObjLen(tbl))
		tbl.ForEach(func(value lua.LValue, value2 lua.LValue) {
			data := new(STdrop_proto)
			data.Sid = value.String()
			if tbl2, ok2 := value2.(*lua.LTable); ok2 {
				fmt.Println(L.ObjLen(tbl2))
				tbl2.ForEach(func(value3 lua.LValue, value4 lua.LValue) {
					if "coin" == value3.String() {
						data.Coin = value4.String()
					}
					if "item1" == value3.String() {
						data.Item1 = value4.String()
					}
					if "item2" == value3.String() {
						data.Item2 = value4.String()
					}
				})
			}
			drop_prototmp[data.Sid] = data
		})
	}
	Gdrop_proto = drop_prototmp
}
