package _go

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
)

var Gshop_proto map[string]*STshop_proto

type STshop_proto struct {
	Sid               string
	Shop_type         string
	Money             string
	Refresh_time      string
	Refresh_byself    string
	Refresh_free      string
	Refresh_price     string
	Refresh_price_add string
	Goods_num_max     string
}

func init() {
	Gshop_proto = make(map[string]*STshop_proto)
	Initshop_proto()
}

func Initshop_proto() {
	L := lua.NewState()
	shop_prototmp := make(map[string]*STshop_proto)
	defer L.Close()
	err := L.DoFile("./lua/shop_proto.lua")
	if err != nil {
		fmt.Println("err", err.Error())
		panic(err)
	}
	shop_proto := L.GetGlobal("shop_proto")
	if tbl, ok := shop_proto.(*lua.LTable); ok {
		tbl.ForEach(func(value lua.LValue, value2 lua.LValue) {
			data := new(STshop_proto)
			data.Sid = value.String()
			if tbl2, ok2 := value2.(*lua.LTable); ok2 {
				tbl2.ForEach(func(value3 lua.LValue, value4 lua.LValue) {
					if "shop_type" == value3.String() {
						data.Shop_type = luaValueToString(value4)
					}
					if "money" == value3.String() {
						data.Money = luaValueToString(value4)
					}
					if "refresh_time" == value3.String() {
						data.Refresh_time = luaValueToString(value4)
					}
					if "refresh_byself" == value3.String() {
						data.Refresh_byself = luaValueToString(value4)
					}
					if "refresh_free" == value3.String() {
						data.Refresh_free = luaValueToString(value4)
					}
					if "refresh_price" == value3.String() {
						data.Refresh_price = luaValueToString(value4)
					}
					if "refresh_price_add" == value3.String() {
						data.Refresh_price_add = luaValueToString(value4)
					}
					if "goods_num_max" == value3.String() {
						data.Goods_num_max = luaValueToString(value4)
					}
				})
			}
			shop_prototmp[data.Sid] = data
		})
	}
	Gshop_proto = shop_prototmp
}
