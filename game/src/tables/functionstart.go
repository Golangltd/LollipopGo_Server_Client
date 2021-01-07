package tables

import (
	. "LollipopGo2.8x/conf/g"
	. "LollipopGo2.8x/lua/go"
)

/*
解析功能开启表数据
*/

//这里只获取巫师考核的配置信息(sid=420)
var ExamModelStartTime int //考试模块开启天数(单位:dy)

func init() {
	sampleData := FunctionStartConf[ExamModelStartID].Condition1
	m := StrToMap(sampleData)
	ExamModelStartTime = m[int(ServiceDays)]
}
