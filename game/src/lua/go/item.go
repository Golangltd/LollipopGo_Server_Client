package _go

import (
	"fmt"
	"github.com/golang/glog"
	lua "github.com/yuin/gopher-lua"
)

var ItemConf map[string]*ItemConfST

type ItemConfST struct {
	Sid       string
	Itemtype  string
	Bagtype   string
	Num_max   string
	Level     string
	Sale      string
	Price     string
	Parameter string
}

func init() {
	ItemConf = make(map[string]*ItemConfST)
	initItemConf()
}

func initItemConf() {
	L := lua.NewState()
	itemTmp := make(map[string]*ItemConfST)
	defer L.Close()
	err := L.DoFile("./lua/item_proto.lua")
	if err != nil {
		fmt.Println("err", err.Error())
		panic(err)
	}
	luaTable := L.GetGlobal("item_proto")
	if tbl, ok := luaTable.(*lua.LTable); ok {
		tbl.ForEach(func(value lua.LValue, value2 lua.LValue) {
			data := new(ItemConfST)
			data.Sid = value.String()
			if tbl2, ok2 := value2.(*lua.LTable); ok2 {
				tbl2.ForEach(func(value3 lua.LValue, value4 lua.LValue) {
					if "itemtype" == value3.String() {
						data.Itemtype = luaValueToString(value4)
					}
					if "bagtype" == value3.String() {
						data.Bagtype = luaValueToString(value4)
					}
					if "num_max" == value3.String() {
						data.Num_max = luaValueToString(value4)
					}
					if "level" == value3.String() {
						data.Level = luaValueToString(value4)
					}
					if "sale" == value3.String() {
						data.Sale = luaValueToString(value4)
					}
					if "price" == value3.String() {
						data.Price = luaValueToString(value4)
					}
					if "parameter" == value3.String() {
						data.Parameter = luaValueToString(value4)
					}
				})
			}
			itemTmp[data.Sid] = data
		})
	}
	ItemConf = itemTmp
	glog.Infof("load ItemConfST size:%d", len(ItemConf))
}
