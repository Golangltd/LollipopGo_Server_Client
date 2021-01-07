package _go

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
)

var Gsubject_proto map[string]*STsubject_proto

type STsubject_proto struct {
	Sid        string
	Desc       string
	College    string
	Camp       string
	Times      string
	Student_id string
}

func init() {
	Gsubject_proto = make(map[string]*STsubject_proto)
}

func Initsubject_proto() {
	L := lua.NewState()
	subject_prototmp := make(map[string]*STsubject_proto)
	defer L.Close()
	err := L.DoFile("./lua/subject_proto.lua")
	if err != nil {
		fmt.Println("err", err.Error())
		panic(err)
	}
	subject_proto := L.GetGlobal("subject_proto")
	if tbl, ok := subject_proto.(*lua.LTable); ok {
		fmt.Println(L.ObjLen(tbl))
		tbl.ForEach(func(value lua.LValue, value2 lua.LValue) {
			data := new(STsubject_proto)
			data.Sid = value.String()
			if tbl2, ok2 := value2.(*lua.LTable); ok2 {
				fmt.Println(L.ObjLen(tbl2))
				tbl2.ForEach(func(value3 lua.LValue, value4 lua.LValue) {
					if "desc" == value3.String() {
						data.Desc = value4.String()
					}
					if "college" == value3.String() {
						data.College = value4.String()
					}
					if "camp" == value3.String() {
						data.Camp = value4.String()
					}
					if "times" == value3.String() {
						data.Times = value4.String()
					}
					if "student_id" == value3.String() {
						data.Student_id = value4.String()
					}
				})
			}
			subject_prototmp[data.Sid] = data
		})
	}
	Gsubject_proto = subject_prototmp
}
