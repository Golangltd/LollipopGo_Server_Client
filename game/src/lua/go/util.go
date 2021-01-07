package _go

import (
	lua "github.com/yuin/gopher-lua"
	"strings"
)

func luaValueToString(value lua.LValue) string {
	sb := &strings.Builder{}
	count, _ := recursiveLuaValueToString(true, 0, sb, value)
	str := sb.String()
	if count > 1 {
		str = str[1 : len(str)-3]
	}
	return str
}

func recursiveLuaValueToString(newTable bool, count int, sb *strings.Builder, value lua.LValue) (int, bool) {
	if value.Type() == lua.LTTable {
		count++
		sb.WriteString(`{`)
		newTable = true
		value.(*lua.LTable).ForEach(func(key lua.LValue, value lua.LValue) {
			count, newTable = recursiveLuaValueToString(newTable, count, sb, value)
		})
		sb.WriteString("}")
		sb.WriteString(",")
	} else {
		if !newTable {
			sb.WriteString(",")
		}
		newTable = false
		sb.WriteString(value.String())
	}
	return count, newTable
}
