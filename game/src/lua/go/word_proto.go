package _go

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
)

var Gword_proto map[string]*STword_proto

type STword_proto struct {
	Sid string
	Cn  string
}

func init() {
	Gword_proto = make(map[string]*STword_proto)
}

func Initword_proto() {
	L := lua.NewState()
	word_prototmp := make(map[string]*STword_proto)
	defer L.Close()
	err := L.DoFile("./lua/word_proto.lua")
	if err != nil {
		fmt.Println("err", err.Error())
		panic(err)
	}
	word_proto := L.GetGlobal("word_proto")
	if tbl, ok := word_proto.(*lua.LTable); ok {
		fmt.Println(L.ObjLen(tbl))
		tbl.ForEach(func(value lua.LValue, value2 lua.LValue) {
			data := new(STword_proto)
			data.Sid = value.String()
			if tbl2, ok2 := value2.(*lua.LTable); ok2 {
				fmt.Println(L.ObjLen(tbl2))
				tbl2.ForEach(func(value3 lua.LValue, value4 lua.LValue) {
					if "cn" == value3.String() {
						data.Cn = value4.String()
					}
				})
			}
			word_prototmp[data.Sid] = data
		})
	}
	Gword_proto = word_prototmp
}
