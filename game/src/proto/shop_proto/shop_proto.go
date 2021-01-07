package shop_proto

// 主协议 game == 6

const (
	C2GS_GetShopInfoProto2 = 151 // 获取商城消息
	GS2C_GetShopInfoProto2 = 152 // 返回数据

	C2GS_UpdateShopInfoProto2 = 153 // 刷新商城数据

	C2GS_GetBuyGoodsProto2 = 154 // 获取商城消息
	GS2C_GetBuyGoodsProto2 = 155 // 返回数据
)

// 商店结构数据
type ShopInfo struct {
	ShopType int // 对应表的配置类型
}

// 商品数据结构
type GoodsInfo struct {
	Sid  int // 商品Id
	Num  int // 系统的拥有的数量
	Left int // 剩余次数
}

// 购买成功后
type GoodsData struct {
	Id    int64
	Itype int64
	Num   int64
}

//-----------------------------------------------------------------------------------------
// C2GS_GetBuyGoodsProto2  == 154
type C2GS_GetBuyGoods struct {
	Protocol  int
	Protocol2 int
	OpenId    string
	GoodId    int // 商品Id
}

// GS2C_GetBuyGoodsProto2  == 155
type GS2C_GetBuyGoods struct {
	Protocol  int
	Protocol2 int
	Goods     []*GoodsData
}

//-----------------------------------------------------------------------------------------
// C2GS_UpdateShopInfoProto2 == 153
type C2GS_UpdateShopInfo struct {
	Protocol  int
	Protocol2 int
	OpenId    string
	ShopType  int // 对应表的数据
}

//-----------------------------------------------------------------------------------------
// C2GS_GetShopInfoProto2 == 151
type C2GS_GetShopInfo struct {
	Protocol  int
	Protocol2 int
	OpenId    string
	ShopType  int // 对应表的数据
}

// GS2C_GetShopInfoProto2 == 152
type GS2C_GetShopInfo struct {
	Protocol     int
	Protocol2    int
	GoodsInfo    []*GoodsInfo
	RefreshTime  int64 // 时间戳 暂时不用
	RefreshCount int64 // 已刷新次数
	ShopType     int
}

//-----------------------------------------------------------------------------------------
