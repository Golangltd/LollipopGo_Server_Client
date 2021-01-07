package Shop

import (
	impl "LollipopGo/network"
	"LollipopGo/tools/sample"
	"LollipopGo2.8x/handlers/Func"
	"LollipopGo2.8x/handlers/common"
	util_handlers "LollipopGo2.8x/handlers/util"
	lua_uitl "LollipopGo2.8x/lua/Uitl"
	"LollipopGo2.8x/proto/comm_proto"
	"LollipopGo2.8x/proto/shop_proto"
	"LollipopGo2.8x/tables"
	twProto "github.com/Golangltd/Twlib/proto"
	twlib_user "github.com/Golangltd/Twlib/user"
	"github.com/golang/glog"
	"golang.org/x/net/websocket"
	"strconv"
	"strings"
	"time"
)

// 商品数据结构
type GoodsInfo struct {
	Sid  int // 商品Id
	Num  int // 系统的拥有的数量
	Left int // 剩余次数
}

//上架,下架 条件
func GoodsCondition(sale_condition_min string, sale_condition_max string, lev int) bool {
	datamap := lua_uitl.GetLuaDataFor2Row(sale_condition_min)
	for _, v := range datamap {
		if v <= lev {
			if len(sale_condition_max) == 0 {
				return true
			}
			for _, v1 := range lua_uitl.GetLuaDataFor2Row(sale_condition_max) {
				if v1 >= lev {
					return true
				}
			}
		}
	}
	return false
}

func buyGoodsCondition(goodsid int) bool {

	data := GGoodsBase[strconv.Itoa(goodsid)]
	if data == nil {
		glog.Error("无对应的商品，获取的商品ID：", goodsid)
		return false
	}
	// 获取功能开启数据
	if len(data.Buy_condition) != 0 {
		datamap := lua_uitl.GetLuaDataFor2Row(data.Buy_condition)
		return Func.IsFuncUnlock(datamap, strings.Contains(data.Buy_condition, "&&"))
	}
	return true
}

// 获取限购数据--配置数据
// return：限购类型，限购数量
func GetLimitBuyInfo(goodsid int) (int, int, int) {

	data := GGoodsBase[strconv.Itoa(goodsid)]
	if data == nil {
		glog.Error("无对应的商品，获取的商品ID：", goodsid)
		return ErrorLimitBuy, ErrorLimitBuy, GetLimitBuyInfoServer(goodsid)
	}

	if len(data.Player_limit) == 0 {
		return NoLimitBuy, NoLimitBuy, GetLimitBuyInfoServer(goodsid)
	}

	datamap := lua_uitl.GetLuaDataFor2Row(data.Player_limit)
	for k, v := range datamap {
		return k, v, GetLimitBuyInfoServer(goodsid)
	}

	return ErrorLimitBuy, ErrorLimitBuy, GetLimitBuyInfoServer(goodsid)
}

// 服务器限购
func GetLimitBuyInfoServer(goodsid int) int {

	return 0

	data := GGoodsBase[strconv.Itoa(goodsid)]
	if data == nil {
		glog.Error("无对应的商品，获取的商品ID：", goodsid)
		return SErrorLimitBuy
	}

	if len(data.System_limit) == 0 {
		return SErrorLimitBuy
	}

	datamap := lua_uitl.GetLuaDataFor2Row(data.System_limit)
	for _, v := range datamap {
		return v
	}

	return SErrorLimitBuy
}

func getShopConditionInfo(iShopType int, RoleLev int) map[int]int {

	conf_lua := GshopBase[iShopType]
	goods_num_max, _ := strconv.Atoi(conf_lua.Goods_num_max)
	if goods_num_max == 0 { // 不限购
		glog.Info("商店不限上架数量！")
	} else {
		glog.Info("商店限上架数量！goods_num_max: ", goods_num_max)
	}

	conf_luagoods := Gshop[iShopType]
	retarr := make(map[int]int)

	sidmap := make(map[int][]int)    // sid 数组
	Weightmap := make(map[int][]int) // 权重数组

	if conf_lua.Shop_type == "1" {
		for i := 0; i < goods_num_max; i++ {
			sidarr := []int{}
			Weightarr := []int{}
			for _, v := range conf_luagoods {
				if v.Shopid == strconv.Itoa(iShopType) && strconv.Itoa(i+1) == v.Num {
					if !GoodsCondition(GGoodsBase[v.Sid].Sale_condition_min,
						GGoodsBase[v.Sid].Sale_condition_max,
						RoleLev) {
						continue
					}
					sidarr = append(sidarr, tables.StrToInt(v.Sid))
					if tables.StrToInt(v.Weight) == 0 {
						continue
					}
					Weightarr = append(Weightarr, tables.StrToInt(v.Weight))
				}
			}
			sidmap[i] = sidarr
			Weightmap[i] = Weightarr
		}

		for i := 0; i < goods_num_max; i++ {
			if len(Weightmap[i]) == 0 {
				continue
			}
			iret := sample.WeightedChoice(Weightmap[i])
			retarr[i] = sidmap[i][iret]
		}
	} else {
		for i := 0; i < goods_num_max; i++ {
			sidarr := []int{}
			for _, v := range conf_luagoods {
				if v.Shopid == strconv.Itoa(iShopType) {
					if !GoodsCondition(GGoodsBase[v.Sid].Sale_condition_min,
						GGoodsBase[v.Sid].Sale_condition_max,
						RoleLev) {
						continue
					}
					sidarr = append(sidarr, tables.StrToInt(v.Sid))
				}
			}
			sidmap[i] = sidarr
		}
		for i := 0; i < len(sidmap[i]); i++ {
			if sidmap[i] == nil {
				continue
			}
			retarr[i] = sidmap[i][i]
		}
	}
	return retarr
}

//获取商城协议
func GetShopInfo(conn *websocket.Conn, ProtocolData map[string]interface{}, update bool) {


	return
}

// 获取缓存商店数据

// 玩家购买商品
func GetBuyGoods(conn *websocket.Conn, ProtocolData map[string]interface{}) {

	strOpenId := ProtocolData["OpenId"].(string)
	iGoodId := int(ProtocolData["GoodId"].(float64))
	goodsdata := GGoodsBase[strconv.Itoa(iGoodId)]
	shoptype, _ := strconv.Atoi(goodsdata.Shopid)

	data := &shop_proto.GS2C_GetBuyGoods{
		Protocol:  twProto.GGameHallProto,
		Protocol2: shop_proto.GS2C_GetBuyGoodsProto2,
	}

	if !buyGoodsCondition(iGoodId) {
		return
	}

	irettype, iretnum, snum := GetLimitBuyInfo(iGoodId)
	glog.Info("限购类型：", irettype, iretnum)
	game, _, _ := util_handlers.GetGameAndUser(strOpenId)

	buylimit := 0
	if irettype == PerDayLimitBuy { // 每天限制
		buylimit = Day1
	} else if irettype == PerWeekLimitBuy { // 每周限制
		buylimit = Day7
	} else if irettype == OnePersonLimitBuy { // 终身限制
		buylimit = Day365
	}

	// 记录玩家的次数购买
	val, err := game.CacheDB.Value(strconv.Itoa(int(game.AccountId)) + "|" + strconv.Itoa(int(iGoodId)) + "|" + "BuyCount")
	if err != nil {
	} else {
		if val.Data().(int) >= iretnum {
			data.Goods = nil
			impl.PlayerSendToProxyServer(conn, data, strOpenId)
			return
		}
	}

	// 记录玩家的次数购买
	vals, err := game.CacheDB.Value(strconv.Itoa(int(iGoodId)) + "|" + "BuyCount")
	if snum != 0 {
		if err != nil {
		} else {
			if vals.Data().(int) >= snum {
				data.Goods = nil
				impl.PlayerSendToProxyServer(conn, data, strOpenId)
				return
			}
		}
	}

	iteminfomap := lua_uitl.GetLuaDataFor3Row(goodsdata.Item_group)
	priceinfomap := lua_uitl.GetLuaDataFor3Row(goodsdata.Price)

	send_info := new(shop_proto.GoodsData)
	itemDatatmp := new(twlib_user.ItemData)
	for _, v := range iteminfomap {
		send_info.Id = int64(v.Id)
		send_info.Num = v.Num
		send_info.Itype = int64(v.Type)

		itemDatatmp.ItemId = v.Id
		itemDatatmp.ItemNum = v.Num
		itemDatatmp.ItemType = v.Type
	}
	data.Goods = append(data.Goods, send_info)

	for _, v := range priceinfomap {
		if v.Type == twlib_user.ICoin {
			if comm_proto.UpdateRoleCoin(conn, strOpenId, int(-v.Num), game.UserInfo) == -1 {
				return
			}
			break
		} else if v.Type == twlib_user.ISilver {
			if comm_proto.UpdateRoleDiamond(conn, strOpenId, int(-v.Num), game.UserInfo) == -1 {
				return
			}
			break
		}
	}

	impl.PlayerSendToProxyServer(conn, data, strOpenId)
	// 更新道具
	common.UpdateItemOfServer(itemDatatmp, strOpenId, conn, 0)
	if iretnum != 0 {
		if val != nil {
			game.CacheDB.Add(strconv.Itoa(int(game.AccountId))+"|"+strconv.Itoa(int(iGoodId))+"|"+"BuyCount", time.Duration(buylimit), val.Data().(int)+1)
		} else {
			game.CacheDB.Add(strconv.Itoa(int(game.AccountId))+"|"+strconv.Itoa(int(iGoodId))+"|"+"BuyCount", time.Duration(buylimit), 1)
		}
	}
	if snum != 0 {
		if vals != nil {
			game.CacheDB.Add(strconv.Itoa(int(iGoodId))+"|"+"BuyCount", time.Duration(Day365), vals.Data().(int)+1)
		} else {
			game.CacheDB.Add(strconv.Itoa(int(iGoodId))+"|"+"BuyCount", time.Duration(Day365), 1)
		}
	}
	// 更新下列表
	ProtocolDatatmp := make(map[string]interface{})
	ProtocolDatatmp["OpenId"] = strOpenId
	ProtocolDatatmp["ShopType"] = float64(shoptype)
	GetShopInfo(conn, ProtocolDatatmp, false)
	return
}
