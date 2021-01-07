package Shop

import (
	util_handlers "LollipopGo2.8x/handlers/util"
	"LollipopGo2.8x/models"
	"LollipopGo2.8x/tables"
	"github.com/golang/glog"
	"github.com/robfig/cron"
)

func init() {
	//tz.Schedule(func() {
	refreshShop()
	//}, g.ReloadShopConfigLua, nil)
}

func refreshShop() {
	for _, v := range GshopBase {
		if tables.RecruitTypeTable[tables.StrToInt(v.Sid)] == nil {
			continue
		}
		c := cron.New()
		err := c.AddFunc(tables.RecruitTypeTable[tables.StrToInt(v.Sid)].TimesResetTime, func() {
			m := util_handlers.GetGameAllUser()
			for itr := m.Iterator(); itr.HasNext(); {
				_, vbg, _ := itr.Next()
				if vbg.(interface{}).(*models.Game) == nil {
					continue
				}
				ProtocolData := make(map[string]interface{})
				ProtocolData["OpenId"] = vbg.(interface{}).(*models.Game).StrMD5
				ProtocolData["ShopType"] = float64(tables.StrToInt(v.Sid))
				//UpdateShopInfo(cxt.ConnXZ, ProtocolData, true)
			}
		})
		if err != nil {
			glog.Error("ticker error!,error is [%v]", err.Error())
		}
		go c.Start()
		defer c.Stop()
	}
}
