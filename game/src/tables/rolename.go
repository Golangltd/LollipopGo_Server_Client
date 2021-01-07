package tables

import (
	"LollipopGo/log"
	gamedb "LollipopGo2.8x/data"
	twlib_dbtable "github.com/Golangltd/Twlib/dbtable"
)

// 角色姓名表
type RoleName struct {
	Name string // 配置的姓名
}

var GMapRoleNameTable map[string][]*RoleName

// 发送消息获取配置
func AddRoleNameTable() {
	GMapRoleNameTable = make(map[string][]*RoleName)
	data := gamedb.GetCFGameData(twlib_dbtable.Gl_xm_sjxmb)
	if data == nil {
		log.Fatal("error! receive GMapRoleNameTable is nil!!")
	}
	// 循环取数据
	for _, v := range data {
		rename := new(RoleName)
		rename.Name = v.(map[string]interface{})["chs_F_name"].(string)
		rename.Name = v.(map[string]interface{})["chs_M_name"].(string)
		rename.Name = v.(map[string]interface{})["chs_F_name"].(string)
		rename.Name = v.(map[string]interface{})["chs_M_name"].(string)
	}
}
