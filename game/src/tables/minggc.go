package tables

import (
	gamedb "LollipopGo2.8x/data"
	twlib_dbtable "github.com/Golangltd/Twlib/dbtable"
)

var GmapMingGC map[string]*SensitivitySt

// 敏感词过滤表
type SensitivitySt struct {
	Sid  int
	Name string
}

// 获取敏感词函数
func GetMingGCInfo() {
	data := gamedb.GetCFGameData(twlib_dbtable.Gl_mg_mgc)
	_=data
}
