package _go

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
)

var Gmap_proto map[string]*STmap_proto

type STmap_proto struct {
	Sid  string
	Name string
	Map  string
}

func init() {
	Gmap_proto = make(map[string]*STmap_proto)
}

func Initmap_proto() {
	L := lua.NewState()
	map_prototmp := make(map[string]*STmap_proto)
	defer L.Close()
	err := L.DoFile("./lua/map_proto.lua")
	if err != nil {
		fmt.Println("err", err.Error())
		panic(err)
	}
	map_proto := L.GetGlobal("map_proto")
	if tbl, ok := map_proto.(*lua.LTable); ok {
		fmt.Println(L.ObjLen(tbl))
		tbl.ForEach(func(value lua.LValue, value2 lua.LValue) {
			data := new(STmap_proto)
			data.Sid = value.String()
			if tbl2, ok2 := value2.(*lua.LTable); ok2 {
				fmt.Println(L.ObjLen(tbl2))
				tbl2.ForEach(func(value3 lua.LValue, value4 lua.LValue) {
					if "name" == value3.String() {
						data.Name = value4.String()
					}
					if "map" == value3.String() {
						data.Map = value4.String()
					}
				})
			}
			map_prototmp[data.Sid] = data
		})
	}
	Gmap_proto = map_prototmp
}
