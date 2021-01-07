package _go

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
)

var Gjdjjc_proto map[string]*STjdjjc_proto

type STjdjjc_proto struct {
	Sid          string
	Integral_min string
	Integral_max string
	RandomM      string
	Floatk       string
}

func init() {
	Gjdjjc_proto = make(map[string]*STjdjjc_proto)
}

func Initjdjjc_proto() {
	L := lua.NewState()
	jdjjc_prototmp := make(map[string]*STjdjjc_proto)
	defer L.Close()
	err := L.DoFile("./lua/jdjjc_proto.lua")
	if err != nil {
		fmt.Println("err", err.Error())
		panic(err)
	}
	jdjjc_proto := L.GetGlobal("jdjjc_proto")
	if tbl, ok := jdjjc_proto.(*lua.LTable); ok {
		fmt.Println(L.ObjLen(tbl))
		tbl.ForEach(func(value lua.LValue, value2 lua.LValue) {
			data := new(STjdjjc_proto)
			data.Sid = value.String()
			if tbl2, ok2 := value2.(*lua.LTable); ok2 {
				fmt.Println(L.ObjLen(tbl2))
				tbl2.ForEach(func(value3 lua.LValue, value4 lua.LValue) {
					if "integral_min" == value3.String() {
						data.Integral_min = value4.String()
					}
					if "integral_max" == value3.String() {
						data.Integral_max = value4.String()
					}
					if "randomM" == value3.String() {
						data.RandomM = value4.String()
					}
					if "floatk" == value3.String() {
						data.Floatk = value4.String()
					}
				})
			}
			jdjjc_prototmp[data.Sid] = data
		})
	}
	Gjdjjc_proto = jdjjc_prototmp
}
