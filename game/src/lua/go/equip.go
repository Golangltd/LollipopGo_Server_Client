package _go

import (
	"fmt"
	"github.com/golang/glog"
	lua "github.com/yuin/gopher-lua"
)

var EquipConf map[string]*EquipConfST

type EquipConfST struct {
	Sid           string
	Icon          string
	Name          string
	Bz            string
	Itemtype      string
	Quality       string
	Vocation      string
	Equip_sex     string
	Equiptype     string
	Position      string
	Bagtype       string
	Num_max       string
	Level         string
	Sale          string
	Camp_ad       string
	Strscore      string
	Maxstr        string
	Struse_score  string
	Struse_copper string
	Attribute1    string
	Attribute2    string
	Attribute3    string
	Attribute4    string
	Attribute5    string
}

func init() {
	EquipConf = make(map[string]*EquipConfST)
	initEquipConf()
}

func initEquipConf() {
	L := lua.NewState()
	equipTmp := make(map[string]*EquipConfST)
	defer L.Close()
	err := L.DoFile("./lua/equip_proto.lua")
	if err != nil {
		fmt.Println("err", err.Error())
		panic(err)
	}
	luaTable := L.GetGlobal("equip_proto")
	if tbl, ok := luaTable.(*lua.LTable); ok {
		tbl.ForEach(func(value lua.LValue, value2 lua.LValue) {
			data := new(EquipConfST)
			data.Sid = value.String()
			if tbl2, ok2 := value2.(*lua.LTable); ok2 {
				tbl2.ForEach(func(value3 lua.LValue, value4 lua.LValue) {
					if "icon" == value3.String() {
						data.Icon = luaValueToString(value4)
					}
					if "name" == value3.String() {
						data.Name = luaValueToString(value4)
					}
					if "bz" == value3.String() {
						data.Bz = luaValueToString(value4)
					}
					if "itemtype" == value3.String() {
						data.Itemtype = luaValueToString(value4)
					}
					if "quality" == value3.String() {
						data.Quality = luaValueToString(value4)
					}
					if "vocation" == value3.String() {
						data.Vocation = luaValueToString(value4)
					}
					if "equip_sex" == value3.String() {
						data.Equip_sex = luaValueToString(value4)
					}
					if "equiptype" == value3.String() {
						data.Equiptype = luaValueToString(value4)
					}
					if "position" == value3.String() {
						data.Position = luaValueToString(value4)
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
					if "camp_ad" == value3.String() {
						data.Camp_ad = luaValueToString(value4)
					}
					if "strscore" == value3.String() {
						data.Strscore = luaValueToString(value4)
					}
					if "maxstr" == value3.String() {
						data.Maxstr = luaValueToString(value4)
					}
					if "struse_score" == value3.String() {
						data.Struse_score = luaValueToString(value4)
					}
					if "struse_copper" == value3.String() {
						data.Struse_copper = luaValueToString(value4)
					}
					if "attribute1" == value3.String() {
						data.Attribute1 = luaValueToString(value4)
					}
					if "attribute2" == value3.String() {
						data.Attribute2 = luaValueToString(value4)
					}
					if "attribute3" == value3.String() {
						data.Attribute3 = luaValueToString(value4)
					}
					if "attribute4" == value3.String() {
						data.Attribute4 = luaValueToString(value4)
					}
					if "attribute5" == value3.String() {
						data.Attribute5 = luaValueToString(value4)
					}
				})
			}
			equipTmp[data.Sid] = data
		})
	}
	EquipConf = equipTmp
	glog.Infof("load EquipConfST size:%d", len(EquipConf))
}
