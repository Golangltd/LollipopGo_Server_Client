package _go

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
)

var Gfashions_proto map[string]*STfashions_proto

type STfashions_proto struct {
	Sid      string
	Position string
	Resouce  string
}

func init() {
	Gfashions_proto = make(map[string]*STfashions_proto)
}

func Initfashions_proto() {
	L := lua.NewState()
	fashions_prototmp := make(map[string]*STfashions_proto)
	defer L.Close()
	err := L.DoFile("./lua/fashions_proto.lua")
	if err != nil {
		fmt.Println("err", err.Error())
		panic(err)
	}
	fashions_proto := L.GetGlobal("fashions_proto")
	if tbl, ok := fashions_proto.(*lua.LTable); ok {
		fmt.Println(L.ObjLen(tbl))
		tbl.ForEach(func(value lua.LValue, value2 lua.LValue) {
			data := new(STfashions_proto)
			data.Sid = value.String()
			if tbl2, ok2 := value2.(*lua.LTable); ok2 {
				fmt.Println(L.ObjLen(tbl2))
				tbl2.ForEach(func(value3 lua.LValue, value4 lua.LValue) {
					if "position" == value3.String() {
						data.Position = value4.String()
					}
					if "resouce" == value3.String() {
						data.Resouce = value4.String()
					}
				})
			}
			fashions_prototmp[data.Sid] = data
		})
	}
	Gfashions_proto = fashions_prototmp
}
