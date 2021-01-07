package _go

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
)

var Gwizardlv_proto map[string]*STwizardlv_proto

type STwizardlv_proto struct {
	Sid        string
	Cost       string
	Levelup    string
	Attribute1 string
	Attribute2 string
	Attribute3 string
	Attribute4 string
}

func init() {
	Gwizardlv_proto = make(map[string]*STwizardlv_proto)
}

func Initwizardlv_proto() {
	L := lua.NewState()
	wizardlv_prototmp := make(map[string]*STwizardlv_proto)
	defer L.Close()
	err := L.DoFile("./lua/wizardlv_proto.lua")
	if err != nil {
		fmt.Println("err", err.Error())
		panic(err)
	}
	wizardlv_proto := L.GetGlobal("wizardlv_proto")
	if tbl, ok := wizardlv_proto.(*lua.LTable); ok {
		fmt.Println(L.ObjLen(tbl))
		tbl.ForEach(func(value lua.LValue, value2 lua.LValue) {
			data := new(STwizardlv_proto)
			data.Sid = value.String()
			if tbl2, ok2 := value2.(*lua.LTable); ok2 {
				fmt.Println(L.ObjLen(tbl2))
				tbl2.ForEach(func(value3 lua.LValue, value4 lua.LValue) {
					if "cost" == value3.String() {
						data.Cost = value4.String()
					}
					if "levelup" == value3.String() {
						data.Levelup = value4.String()
					}
					if "attribute1" == value3.String() {
						data.Attribute1 = value4.String()
					}
					if "attribute2" == value3.String() {
						data.Attribute2 = value4.String()
					}
					if "attribute3" == value3.String() {
						data.Attribute3 = value4.String()
					}
					if "attribute4" == value3.String() {
						data.Attribute4 = value4.String()
					}
				})
			}
			wizardlv_prototmp[data.Sid] = data
		})
	}
	Gwizardlv_proto = wizardlv_prototmp
}
