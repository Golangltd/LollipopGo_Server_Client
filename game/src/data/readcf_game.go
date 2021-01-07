package gamedb

import "github.com/golang/glog"

// 获取配置数据库数据,战斗服务器通过消息获取数据
func GetCFGameData(tableName string) []interface{} {
	var tableData []interface{}
	call := ConnRPC.Go("CfRPC.GetCfGameData", tableName, &tableData, nil)
	replyCall := <-call.Done
	if replyCall.Error != nil {
		glog.Info(replyCall.Error)
	}
	return tableData
}
