package _go

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
)

var Gchap_proto map[string]*STchap_proto

type STchap_proto struct {
	Sid         string
	Chap_name   string
	Sub_chap    string
	Last_level  string
	Fight_map   string
	Award       string
	Install_map string
	Events      string
	Tj_award    string
	Dexc        string
	Item_reward string
}

func init() {
	Gchap_proto = make(map[string]*STchap_proto)
}

func Initchap_proto() {
	L := lua.NewState()
	chap_prototmp := make(map[string]*STchap_proto)
	defer L.Close()
	err := L.DoFile("./lua/chap_proto.lua")
	if err != nil {
		fmt.Println("err", err.Error())
		panic(err)
	}
	chap_proto := L.GetGlobal("chap_proto")
	if tbl, ok := chap_proto.(*lua.LTable); ok {
		fmt.Println(L.ObjLen(tbl))
		tbl.ForEach(func(value lua.LValue, value2 lua.LValue) {
			data := new(STchap_proto)
			data.Sid = value.String()
			if tbl2, ok2 := value2.(*lua.LTable); ok2 {
				fmt.Println(L.ObjLen(tbl2))
				tbl2.ForEach(func(value3 lua.LValue, value4 lua.LValue) {
					if "chap_name" == value3.String() {
						data.Chap_name = value4.String()
					}
					if "sub_chap" == value3.String() {
						data.Sub_chap = value4.String()
					}
					if "last_level" == value3.String() {
						data.Last_level = value4.String()
					}
					if "fight_map" == value3.String() {
						data.Fight_map = value4.String()
					}
					if "award" == value3.String() {
						data.Award = value4.String()
					}
					if "install_map" == value3.String() {
						data.Install_map = value4.String()
					}
					if "events" == value3.String() {
						data.Events = value4.String()
					}
					if "tj_award" == value3.String() {
						data.Tj_award = value4.String()
					}
					if "dexc" == value3.String() {
						data.Dexc = value4.String()
					}
					if "item_reward" == value3.String() {
						data.Item_reward = value4.String()
					}
				})
			}
			chap_prototmp[data.Sid] = data
		})
	}
	Gchap_proto = chap_prototmp
}
