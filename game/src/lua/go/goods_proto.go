package _go

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
)

var Ggoods_proto map[string]*STgoods_proto

type STgoods_proto struct {
	Sid                string
	Num                string
	Shopid             string
	Item_group         string
	Price              string
	Player_limit       string
	System_limit       string
	Sale_condition_min string
	Sale_condition_max string
	Buy_condition      string
	Weight             string
}

func init() {
	Ggoods_proto = make(map[string]*STgoods_proto)
	Initgoods_proto()
}

func Initgoods_proto() {
	L := lua.NewState()
	goods_prototmp := make(map[string]*STgoods_proto)
	defer L.Close()
	err := L.DoFile("./lua/goods_proto.lua")
	if err != nil {
		fmt.Println("err", err.Error())
		panic(err)
	}
	goods_proto := L.GetGlobal("goods_proto")
	if tbl, ok := goods_proto.(*lua.LTable); ok {
		tbl.ForEach(func(value lua.LValue, value2 lua.LValue) {
			data := new(STgoods_proto)
			data.Sid = value.String()
			if tbl2, ok2 := value2.(*lua.LTable); ok2 {
				tbl2.ForEach(func(value3 lua.LValue, value4 lua.LValue) {
					if "shopid" == value3.String() {
						data.Shopid = luaValueToString(value4)
					}
					if "item_group" == value3.String() {
						data.Item_group = luaValueToString(value4)
					}
					if "num" == value3.String() {
						data.Num = luaValueToString(value4)
					}
					if "price" == value3.String() {
						data.Price = luaValueToString(value4)
					}
					if "player_limit" == value3.String() {
						data.Player_limit = luaValueToString(value4)
					}
					if "system_limit" == value3.String() {
						data.System_limit = luaValueToString(value4)
					}
					if "sale_condition_min" == value3.String() {
						data.Sale_condition_min = luaValueToString(value4)
					}

					if "sale_condition_max" == value3.String() {
						data.Sale_condition_max = luaValueToString(value4)
					}
					if "buy_condition" == value3.String() {
						data.Buy_condition = luaValueToString(value4)
					}
					if "weight" == value3.String() {
						data.Weight = luaValueToString(value4)
					}
				})
			}
			goods_prototmp[data.Sid] = data
		})
	}
	Ggoods_proto = goods_prototmp
}
