package _go

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
)

var Gcharacter map[string]*STcharacter

type STcharacter struct {
	Sid             string
	Name            string
	Description     string
	Vocation        string
	Camp            string
	Quality         string
	Maxquality      string
	Skill1          string
	Skill2          string
	Skill3          string
	Skill4          string
	Skill5          string
	Resouce         string
	Attribute1      string
	Gender          string
	Handbookhero_id string
}

func init() {
	Gcharacter = make(map[string]*STcharacter)
}

func Initcharacter() {
	L := lua.NewState()
	charactertmp := make(map[string]*STcharacter)
	defer L.Close()
	err := L.DoFile("./lua/character.lua")
	if err != nil {
		fmt.Println("err", err.Error())
		panic(err)
	}
	character := L.GetGlobal("character")
	if tbl, ok := character.(*lua.LTable); ok {
		fmt.Println(L.ObjLen(tbl))
		tbl.ForEach(func(value lua.LValue, value2 lua.LValue) {
			data := new(STcharacter)
			data.Sid = value.String()
			if tbl2, ok2 := value2.(*lua.LTable); ok2 {
				fmt.Println(L.ObjLen(tbl2))
				tbl2.ForEach(func(value3 lua.LValue, value4 lua.LValue) {
					if "name" == value3.String() {
						data.Name = value4.String()
					}
					if "description" == value3.String() {
						data.Description = value4.String()
					}
					if "vocation" == value3.String() {
						data.Vocation = value4.String()
					}
					if "camp" == value3.String() {
						data.Camp = value4.String()
					}
					if "quality" == value3.String() {
						data.Quality = value4.String()
					}
					if "maxquality" == value3.String() {
						data.Maxquality = value4.String()
					}
					if "skill1" == value3.String() {
						data.Skill1 = value4.String()
					}
					if "skill2" == value3.String() {
						data.Skill2 = value4.String()
					}
					if "skill3" == value3.String() {
						data.Skill3 = value4.String()
					}
					if "skill4" == value3.String() {
						data.Skill4 = value4.String()
					}
					if "skill5" == value3.String() {
						data.Skill5 = value4.String()
					}
					if "resouce" == value3.String() {
						data.Resouce = value4.String()
					}
					if "attribute1" == value3.String() {
						data.Attribute1 = value4.String()
					}
					if "gender" == value3.String() {
						data.Gender = value4.String()
					}
					if "handbookhero_id" == value3.String() {
						data.Handbookhero_id = value4.String()
					}
				})
			}
			charactertmp[data.Sid] = data
		})
	}
	Gcharacter = charactertmp
}
