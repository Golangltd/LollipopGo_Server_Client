package _go

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
)

var Gquality_back map[string]*STquality_back

type STquality_back struct {
	Sid          string
	S_level      string
	C_level      string
	Back_level   string
	Back_amount1 string
	Back_amount2 string
	Back_cost    string
	Camp1_item   string
	Camp2_item   string
	Camp3_item   string
	Camp4_item   string
}

func init() {
	Gquality_back = make(map[string]*STquality_back)
}

func Initquality_back() {
	L := lua.NewState()
	quality_backtmp := make(map[string]*STquality_back)
	defer L.Close()
	err := L.DoFile("./lua/quality_back.lua")
	if err != nil {
		fmt.Println("err", err.Error())
		panic(err)
	}
	quality_back := L.GetGlobal("quality_back")
	if tbl, ok := quality_back.(*lua.LTable); ok {
		fmt.Println(L.ObjLen(tbl))
		tbl.ForEach(func(value lua.LValue, value2 lua.LValue) {
			data := new(STquality_back)
			data.Sid = value.String()
			if tbl2, ok2 := value2.(*lua.LTable); ok2 {
				fmt.Println(L.ObjLen(tbl2))
				tbl2.ForEach(func(value3 lua.LValue, value4 lua.LValue) {
					if "s_level" == value3.String() {
						data.S_level = value4.String()
					}
					if "c_level" == value3.String() {
						data.C_level = value4.String()
					}
					if "back_level" == value3.String() {
						data.Back_level = value4.String()
					}
					if "back_amount1" == value3.String() {
						data.Back_amount1 = value4.String()
					}
					if "back_amount2" == value3.String() {
						data.Back_amount2 = value4.String()
					}
					if "back_cost" == value3.String() {
						data.Back_cost = value4.String()
					}
					if "camp1_item" == value3.String() {
						data.Camp1_item = value4.String()
					}
					if "camp2_item" == value3.String() {
						data.Camp2_item = value4.String()
					}
					if "camp3_item" == value3.String() {
						data.Camp3_item = value4.String()
					}
					if "camp4_item" == value3.String() {
						data.Camp4_item = value4.String()
					}
				})
			}
			quality_backtmp[data.Sid] = data
		})
	}
	Gquality_back = quality_backtmp
}
