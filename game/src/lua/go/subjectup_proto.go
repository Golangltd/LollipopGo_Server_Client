package _go

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
)

var Gsubjectup_proto map[string]*STsubjectup_proto

type STsubjectup_proto struct {
	Sid           string
	Subject_id    string
	Subjectlevel  string
	College_level string
	Seat          string
	Reword1       string
	Reword2       string
}

func init() {
	Gsubjectup_proto = make(map[string]*STsubjectup_proto)
}

func Initsubjectup_proto() {
	L := lua.NewState()
	subjectup_prototmp := make(map[string]*STsubjectup_proto)
	defer L.Close()
	err := L.DoFile("./lua/subjectup_proto.lua")
	if err != nil {
		fmt.Println("err", err.Error())
		panic(err)
	}
	subjectup_proto := L.GetGlobal("subjectup_proto")
	if tbl, ok := subjectup_proto.(*lua.LTable); ok {
		fmt.Println(L.ObjLen(tbl))
		tbl.ForEach(func(value lua.LValue, value2 lua.LValue) {
			data := new(STsubjectup_proto)
			data.Sid = value.String()
			if tbl2, ok2 := value2.(*lua.LTable); ok2 {
				fmt.Println(L.ObjLen(tbl2))
				tbl2.ForEach(func(value3 lua.LValue, value4 lua.LValue) {
					if "subject_id" == value3.String() {
						data.Subject_id = value4.String()
					}
					if "subjectlevel" == value3.String() {
						data.Subjectlevel = value4.String()
					}
					if "college_level" == value3.String() {
						data.College_level = value4.String()
					}
					if "seat" == value3.String() {
						data.Seat = value4.String()
					}
					if "reword1" == value3.String() {
						data.Reword1 = value4.String()
					}
					if "reword2" == value3.String() {
						data.Reword2 = value4.String()
					}
				})
			}
			subjectup_prototmp[data.Sid] = data
		})
	}
	Gsubjectup_proto = subjectup_prototmp
}
