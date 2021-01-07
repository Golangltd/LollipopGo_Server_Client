package ticker

import (
	"LollipopGo/log"
	. "LollipopGo2.8x/conf/g"
	"LollipopGo2.8x/tables"
	"github.com/robfig/cron"
)

type TickerUid int64

//重置服务器抽奖免费次数的ticker
func init() {
	cronHandler()
}

//重置免费抽
func cronHandler() {
	c := cron.New()
	err := c.AddFunc(tables.RecruitTypeTable[int(CommonDraw)].TimesResetTime, func() {
		CDUsedFreeTimes = make(map[int64]int) //定时重置次数
	})
	if err != nil {
		log.Error("ticker error!,error is [%v]", err.Error())
	}
	go c.Start()
	defer c.Stop()
}
