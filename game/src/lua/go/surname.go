package _go

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
)

var Gsurname map[string]*STsurname

type STsurname struct {
	Sid         string
	Chs_surname string
	Eng_surname string
}

func init() {
	Gsurname = make(map[string]*STsurname)
}

func Initsurname() {
	L := lua.NewState()
	surnametmp := make(map[string]*STsurname)
	defer L.Close()
	err := L.DoFile("./lua/surname.lua")
	if err != nil {
		fmt.Println("err", err.Error())
		panic(err)
	}
	surname := L.GetGlobal("surname")
	if tbl, ok := surname.(*lua.LTable); ok {
		fmt.Println(L.ObjLen(tbl))
		tbl.ForEach(func(value lua.LValue, value2 lua.LValue) {
			data := new(STsurname)
			data.Sid = value.String()
			if tbl2, ok2 := value2.(*lua.LTable); ok2 {
				fmt.Println(L.ObjLen(tbl2))
				tbl2.ForEach(func(value3 lua.LValue, value4 lua.LValue) {
					if "chs_surname" == value3.String() {
						data.Chs_surname = value4.String()
					}
					if "eng_surname" == value3.String() {
						data.Eng_surname = value4.String()
					}
				})
			}
			surnametmp[data.Sid] = data
		})
	}
	Gsurname = surnametmp
}
