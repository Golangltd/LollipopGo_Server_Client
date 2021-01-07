package _go

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
)

var Gwand_proto map[string]*STwand_proto

type STwand_proto struct {
	Sid              string
	Attri_type       string
	Quality          string
	Position         string
	Bagtype          string
	Num_max          string
	Level            string
	Attribute1       string
	Resolve          string
	Entry1_condition string
	Ritio1           string
}

func init() {
	Gwand_proto = make(map[string]*STwand_proto)
}

func Initwand_proto() {
	L := lua.NewState()
	wand_prototmp := make(map[string]*STwand_proto)
	defer L.Close()
	err := L.DoFile("./lua/wand_proto.lua")
	if err != nil {
		fmt.Println("err", err.Error())
		panic(err)
	}
	wand_proto := L.GetGlobal("wand_proto")
	if tbl, ok := wand_proto.(*lua.LTable); ok {
		fmt.Println(L.ObjLen(tbl))
		tbl.ForEach(func(value lua.LValue, value2 lua.LValue) {
			data := new(STwand_proto)
			data.Sid = value.String()
			if tbl2, ok2 := value2.(*lua.LTable); ok2 {
				fmt.Println(L.ObjLen(tbl2))
				tbl2.ForEach(func(value3 lua.LValue, value4 lua.LValue) {
					if "attri_type" == value3.String() {
						data.Attri_type = value4.String()
					}
					if "quality" == value3.String() {
						data.Quality = value4.String()
					}
					if "position" == value3.String() {
						data.Position = value4.String()
					}
					if "bagtype" == value3.String() {
						data.Bagtype = value4.String()
					}
					if "num_max" == value3.String() {
						data.Num_max = value4.String()
					}
					if "level" == value3.String() {
						data.Level = value4.String()
					}
					if "attribute1" == value3.String() {
						data.Attribute1 = value4.String()
					}
					if "resolve" == value3.String() {
						data.Resolve = value4.String()
					}
					if "entry1_condition" == value3.String() {
						data.Entry1_condition = value4.String()
					}
					if "ritio1" == value3.String() {
						data.Ritio1 = value4.String()
					}
				})
			}
			wand_prototmp[data.Sid] = data
		})
	}
	Gwand_proto = wand_prototmp
}
