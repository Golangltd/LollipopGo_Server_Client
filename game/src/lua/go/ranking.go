package _go

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
)

var Granking map[string]*STranking

type STranking struct {
	Sid       string
	Condition string
	Listcount string
	Showtype  string
}

func init() {
	Granking = make(map[string]*STranking)
}

func Initranking() {
	L := lua.NewState()
	rankingtmp := make(map[string]*STranking)
	defer L.Close()
	err := L.DoFile("./lua/ranking.lua")
	if err != nil {
		fmt.Println("err", err.Error())
		panic(err)
	}
	ranking := L.GetGlobal("ranking")
	if tbl, ok := ranking.(*lua.LTable); ok {
		fmt.Println(L.ObjLen(tbl))
		tbl.ForEach(func(value lua.LValue, value2 lua.LValue) {
			data := new(STranking)
			data.Sid = value.String()
			if tbl2, ok2 := value2.(*lua.LTable); ok2 {
				fmt.Println(L.ObjLen(tbl2))
				tbl2.ForEach(func(value3 lua.LValue, value4 lua.LValue) {
					if "condition" == value3.String() {
						data.Condition = value4.String()
					}
					if "listcount" == value3.String() {
						data.Listcount = value4.String()
					}
					if "showtype" == value3.String() {
						data.Showtype = value4.String()
					}
				})
			}
			rankingtmp[data.Sid] = data
		})
	}
	Granking = rankingtmp
}
