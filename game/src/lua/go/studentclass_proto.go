package _go

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
)

var Gstudentclass_proto map[string]*STstudentclass_proto

type STstudentclass_proto struct {
	Sid         string
	Type        string
	Level       string
	Missiontype string
}

func init() {
	Gstudentclass_proto = make(map[string]*STstudentclass_proto)
}

func Initstudentclass_proto() {
	L := lua.NewState()
	studentclass_prototmp := make(map[string]*STstudentclass_proto)
	defer L.Close()
	err := L.DoFile("./lua/studentclass_proto.lua")
	if err != nil {
		fmt.Println("err", err.Error())
		panic(err)
	}
	studentclass_proto := L.GetGlobal("studentclass_proto")
	if tbl, ok := studentclass_proto.(*lua.LTable); ok {
		fmt.Println(L.ObjLen(tbl))
		tbl.ForEach(func(value lua.LValue, value2 lua.LValue) {
			data := new(STstudentclass_proto)
			data.Sid = value.String()
			if tbl2, ok2 := value2.(*lua.LTable); ok2 {
				fmt.Println(L.ObjLen(tbl2))
				tbl2.ForEach(func(value3 lua.LValue, value4 lua.LValue) {
					if "type" == value3.String() {
						data.Type = value4.String()
					}
					if "level" == value3.String() {
						data.Level = value4.String()
					}
					if "missiontype" == value3.String() {
						data.Missiontype = value4.String()
					}
				})
			}
			studentclass_prototmp[data.Sid] = data
		})
	}
	Gstudentclass_proto = studentclass_prototmp
}
