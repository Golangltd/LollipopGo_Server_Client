package _go

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
)

var Gmail_proto map[string]*STmail_proto

type STmail_proto struct {
	Sid         string
	Title       string
	Description string
	Send        string
}

func init() {
	Gmail_proto = make(map[string]*STmail_proto)
}

func Initmail_proto() {
	L := lua.NewState()
	mail_prototmp := make(map[string]*STmail_proto)
	defer L.Close()
	err := L.DoFile("./lua/mail_proto.lua")
	if err != nil {
		fmt.Println("err", err.Error())
		panic(err)
	}
	mail_proto := L.GetGlobal("mail_proto")
	if tbl, ok := mail_proto.(*lua.LTable); ok {
		fmt.Println(L.ObjLen(tbl))
		tbl.ForEach(func(value lua.LValue, value2 lua.LValue) {
			data := new(STmail_proto)
			data.Sid = value.String()
			if tbl2, ok2 := value2.(*lua.LTable); ok2 {
				fmt.Println(L.ObjLen(tbl2))
				tbl2.ForEach(func(value3 lua.LValue, value4 lua.LValue) {
					if "title" == value3.String() {
						data.Title = value4.String()
					}
					if "description" == value3.String() {
						data.Description = value4.String()
					}
					if "send" == value3.String() {
						data.Send = value4.String()
					}
				})
			}
			mail_prototmp[data.Sid] = data
		})
	}
	Gmail_proto = mail_prototmp
}
