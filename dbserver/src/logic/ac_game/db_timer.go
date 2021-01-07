package ac_game

import (
	"LollipopGo/timer"
	twlib_server "github.com/Golangltd/Twlib/server"
	"time"
)

func init()  {
	timelier :=LollipopGo_timer.NewDispatcher(twlib_server.DBServerId)
	timelier.AfterFunc(time.Second*60*60,ReloadConfigDB)
}

// 数据操作, 应用操作
func ReloadConfigDB()  {
  // 提前加载配置表，
}