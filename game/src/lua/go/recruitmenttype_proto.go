package _go

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
)

var Grecruitmenttype_proto map[string]*STrecruitmenttype_proto

type STrecruitmenttype_proto struct {
	Sid         string
	Price       string
	Gold        string
	Number      string
	Free_number string
	Reset       string
	Point       string
	Pool_id     string
}

func init() {
	Grecruitmenttype_proto = make(map[string]*STrecruitmenttype_proto)
	Initrecruitmenttype_proto()
}

func Initrecruitmenttype_proto() {
	L := lua.NewState()
	recruitmenttype_prototmp := make(map[string]*STrecruitmenttype_proto)
	defer L.Close()
	err := L.DoFile("./lua/recruitmenttype_proto.lua")
	if err != nil {
		fmt.Println("err", err.Error())
		panic(err)
	}
	recruitmenttype_proto := L.GetGlobal("recruitmenttype_proto")
	if tbl, ok := recruitmenttype_proto.(*lua.LTable); ok {
		tbl.ForEach(func(value lua.LValue, value2 lua.LValue) {
			data := new(STrecruitmenttype_proto)
			data.Sid = value.String()
			if tbl2, ok2 := value2.(*lua.LTable); ok2 {
				tbl2.ForEach(func(value3 lua.LValue, value4 lua.LValue) {
					if "price" == value3.String() {
						data.Price = luaValueToString(value4)
					}
					if "gold" == value3.String() {
						data.Gold = luaValueToString(value4)
					}
					if "number" == value3.String() {
						data.Number = luaValueToString(value4)
					}
					if "free_number" == value3.String() {
						data.Free_number = luaValueToString(value4)
					}
					if "reset" == value3.String() {
						data.Reset = luaValueToString(value4)
					}
					if "point" == value3.String() {
						data.Point = luaValueToString(value4)
					}
					if "pool_id" == value3.String() {
						data.Pool_id = luaValueToString(value4)
					}
				})
			}
			recruitmenttype_prototmp[data.Sid] = data
		})
	}
	Grecruitmenttype_proto = recruitmenttype_prototmp
}
