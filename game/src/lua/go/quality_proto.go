package _go

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
)

var Gquality_proto map[string]*STquality_proto

type STquality_proto struct {
	Sid        string
	Camp_type  string
	Levelup    string
	Level      string
	Next       string
	Card       string
	Samecard1  string
	Samecard2  string
	Cardlevel1 string
	Cardlevel2 string
	Cardgroup  string
}

func init() {
	Gquality_proto = make(map[string]*STquality_proto)
}

func Initquality_proto() {
	L := lua.NewState()
	quality_prototmp := make(map[string]*STquality_proto)
	defer L.Close()
	err := L.DoFile("./lua/quality_proto.lua")
	if err != nil {
		fmt.Println("err", err.Error())
		panic(err)
	}
	quality_proto := L.GetGlobal("quality_proto")
	if tbl, ok := quality_proto.(*lua.LTable); ok {
		fmt.Println(L.ObjLen(tbl))
		tbl.ForEach(func(value lua.LValue, value2 lua.LValue) {
			data := new(STquality_proto)
			data.Sid = value.String()
			if tbl2, ok2 := value2.(*lua.LTable); ok2 {
				fmt.Println(L.ObjLen(tbl2))
				tbl2.ForEach(func(value3 lua.LValue, value4 lua.LValue) {
					if "camp_type" == value3.String() {
						data.Camp_type = value4.String()
					}
					if "levelup" == value3.String() {
						data.Levelup = value4.String()
					}
					if "level" == value3.String() {
						data.Level = value4.String()
					}
					if "next" == value3.String() {
						data.Next = value4.String()
					}
					if "card" == value3.String() {
						data.Card = value4.String()
					}
					if "samecard1" == value3.String() {
						data.Samecard1 = value4.String()
					}
					if "samecard2" == value3.String() {
						data.Samecard2 = value4.String()
					}
					if "cardlevel1" == value3.String() {
						data.Cardlevel1 = value4.String()
					}
					if "cardlevel2" == value3.String() {
						data.Cardlevel2 = value4.String()
					}
					if "cardgroup" == value3.String() {
						data.Cardgroup = value4.String()
					}
				})
			}
			quality_prototmp[data.Sid] = data
		})
	}
	Gquality_proto = quality_prototmp
}
