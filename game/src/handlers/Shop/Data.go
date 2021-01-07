package Shop

import (
	_go "LollipopGo2.8x/lua/go"
	"LollipopGo2.8x/tables"
	"fmt"
)

// 玩家限购
const (
	ErrorLimitBuy          = -1   // 不限制购买
	NoLimitBuy             = iota // 不限制购买
	PerRefreshTimeLimitBuy        // 单次修改时间限制
	PerDayLimitBuy                // 每日限购
	PerWeekLimitBuy               // 每周限购
	OnePersonLimitBuy             // 终身限购
)

// 服务器限购
const (
	SErrorLimitBuy       = -1   // 不限制购买
	SNoLimitBuy          = iota // 不限制购买
	SPerReshTimeLimitBuy        // 单次修改时间限制
	SPerDayLimitBuy             // 每日限购
	SPerWeekLimitBuy            // 每周限购
)

// 商城使用
const (
	Day1   = 60 * 60 * 24 * 1   // 1天
	Day7   = 60 * 60 * 24 * 7   // 7天
	Day365 = 60 * 60 * 24 * 365 // 365天
)

var GshopBase map[int]*_go.STshop_proto      // 商店数据(配置lua数据shop_proto.lua)
var Gshop map[int][]*_go.STgoods_proto       // 商店类型key，对应相应类型的商品
var GGoodsBase map[string]*_go.STgoods_proto // 所有商品(配置lua数据goods_proto.lua)
var GUserR map[int64]int

func init() {
	GshopBase = make(map[int]*_go.STshop_proto)
	Gshop = make(map[int][]*_go.STgoods_proto)
	GGoodsBase = make(map[string]*_go.STgoods_proto)
	GUserR = make(map[int64]int)

	GshopBase = getAllConfigDataShop()
	Gshop = getGoodsInfoFromShopType()
	GGoodsBase = getAllConfigDataGoods()
}

// 获取lua配置数据
func getAllConfigDataShop() map[int]*_go.STshop_proto {
	data := make(map[int]*_go.STshop_proto)
	for k, v := range _go.Gshop_proto {
		data[tables.StrToInt(k)] = v
	}
	if data == nil {
		fmt.Println("getAllConfigDataShop lua data is nil")
		return nil
	}
	return data
}

//------------------------------------------------------------------------------------------------
// 获取lua配置数据
func getAllConfigDataGoods() map[string]*_go.STgoods_proto {
	return _go.Ggoods_proto
}

// 根据商品类型存储不同商品
// key ==  商店类型
// val ==  商品
func getGoodsInfoFromShopType() map[int][]*_go.STgoods_proto {

	data_save := make(map[int][]*_go.STgoods_proto)
	shopdata := getAllConfigDataShop()
	data := getAllConfigDataGoods()

	for k1, _ := range shopdata {
		data_save1 := []*_go.STgoods_proto{}
		for _, v := range data {
			if k1 == tables.StrToInt(v.Shopid) {
				data_save2 := new(_go.STgoods_proto)
				// 保存到相应的数据结构里
				data_save2 = v
				data_save1 = append(data_save1, data_save2)
			}
		}
		data_save[k1] = data_save1
	}
	return data_save
}
