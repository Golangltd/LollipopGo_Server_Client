package tables

import (
	lua_uitl "LollipopGo2.8x/lua/Uitl"
	_go "LollipopGo2.8x/lua/go"
	twLibItem "github.com/Golangltd/Twlib/item"
	"github.com/golang/glog"
	"regexp"
	"strings"
)

var (
	EquipTables map[int]*EquipConfSt // 装备配置表
)

// 属性数据
type AttributeSt struct {
	Type int // 属性类型
	Val  int // 属性值
}

// 装备配置
type EquipConfSt struct {
	Sid            int                    // 配置ID
	ItemType       int                    // 道具类型
	Quality        int                    // 品质
	UseVocations   []int                  // 使用职业 0通用 1战士 2游侠 3法师 4辅助
	UseSex         int                    // 使用性别 0通用 1男 2女
	EquipType      int                    // 装备类型   0通用 1轻甲 2重甲
	Position       int                    // 装备部位 1武器 2头盔 3衣服 4鞋子
	BagType        int                    // 背包分类  0不放背包 1道具 2装备 3材料
	StackMaxNum    int                    // 叠加最大数量
	UseLv          int                    // 使用等级
	CanSale        bool                   // 能否出售
	Price          *twLibItem.ItemData    // 价格
	AddCampRate    int                    // 产生阵营加成概率 万分比
	ConvertExp     []int                  // 作为强化材料时，所增加的经验数量
	MaxStar        int                    // 强化的最高星级 0-4
	StarUpExp      []int                  // 升星所需的经验 1-4
	AsMaterialCost []*twLibItem.ItemData  // 所谓升星材料所需的货币消耗 map[星级]消耗
	StarAttributes map[int][]*AttributeSt // 星级对应的属性 map[星级]属性
}

func init() {
	ItemTables = make(map[int]*ItemSt)
	initEquipTable()
}

// 初始化装备表
func initEquipTable() {
	EquipTables = make(map[int]*EquipConfSt)
	for _, equip := range _go.EquipConf {
		t := &EquipConfSt{}
		t.Sid = StrToInt(equip.Sid)
		t.ItemType = StrToInt(equip.Itemtype)
		t.Quality = StrToInt(equip.Quality)
		t.UseVocations = parseUseVocations(equip.Vocation)
		t.UseSex = StrToInt(equip.Equip_sex)
		t.EquipType = StrToInt(equip.Equiptype)
		t.Position = StrToInt(equip.Position)
		t.BagType = StrToInt(equip.Bagtype)
		t.StackMaxNum = StrToInt(equip.Bagtype)
		t.UseLv = StrToInt(equip.Level)
		t.CanSale = StrToInt(equip.Sale) == 1
		// t.Price 暂无
		t.AddCampRate = StrToInt(equip.Camp_ad)
		t.MaxStar = StrToInt(equip.Maxstr)
		t.ConvertExp = parseConvertExp(equip.Strscore, t.MaxStar)
		t.StarUpExp = parseStarUpExp(equip.Struse_score, t.MaxStar)
		t.AsMaterialCost = parseAsMaterialCost(equip.Struse_copper)
		t.StarAttributes = make(map[int][]*AttributeSt)
		t.StarAttributes[0] = parseStarAttributes(equip.Attribute1)
		t.StarAttributes[1] = parseStarAttributes(equip.Attribute2)
		t.StarAttributes[2] = parseStarAttributes(equip.Attribute3)
		t.StarAttributes[3] = parseStarAttributes(equip.Attribute4)
		t.StarAttributes[4] = parseStarAttributes(equip.Attribute5)
		EquipTables[t.Sid] = t
	}
	glog.Infof("equip table size:%d", len(EquipTables))
}

// 解析价格
func parsePrice(str string) *twLibItem.ItemData {
	reg := regexp.MustCompile(`\d+,\d+,\d+`)
	strArr := reg.FindAllStringSubmatch(str, -1)
	for _, temp := range strArr {
		for _, t := range temp {
			arr := strings.Split(t, ",")
			r := &twLibItem.ItemData{
				Type: StrToInt(arr[0]),
				ID:   StrToInt(arr[1]),
				Num:  StrToInt(arr[2]),
			}
			return r
		}
	}
	return nil
}

// 解析作为材料时转化的经验
func parseConvertExp(str string, maxStar int) []int {
	if maxStar == 0 { // 不能强化
		exp := make([]int, 5)
		exp[0] = StrToInt(str)
		return exp
	}
	exp := make([]int, 0)
	result := lua_uitl.GetFirstLuaDataFor5Row(str)
	for _, v := range result {
		exp = append(exp, v)
	}
	return exp
}

// 解析升星所需的经验
func parseStarUpExp(str string, maxStar int) []int {
	if maxStar == 0 { // 不能强化
		exp := make([]int, 4)
		return exp
	}
	exp := make([]int, 0)
	result := lua_uitl.GetFirstLuaDataFor4Row(str)
	for _, v := range result {
		exp = append(exp, v)
	}
	return exp
}

// 解析作为材料时的消耗的货币
func parseAsMaterialCost(str string) []*twLibItem.ItemData {
	reg := regexp.MustCompile(`\d+,\d+,\d+`)
	strArr := reg.FindAllStringSubmatch(str, -1)
	m := make([]*twLibItem.ItemData, 0)
	for _, temp := range strArr {
		arr := strings.Split(temp[0], ",")
		m = append(m, &twLibItem.ItemData{
			Type: StrToInt(arr[0]),
			ID:   StrToInt(arr[1]),
			Num:  StrToInt(arr[2]),
		})
	}
	return m
}

// 解析使用职业
func parseUseVocations(str string) []int {
	rStarr := strings.Split(str, ",")
	ret := make([]int, 0)
	for _, v := range rStarr {
		ret = append(ret, StrToInt(v))
	}
	return ret
}

// 解析属性
func parseStarAttributes(str string) []*AttributeSt {
	reg := regexp.MustCompile(`\d+,\d+`)
	strArr := reg.FindAllStringSubmatch(str, -1)
	m := make([]*AttributeSt, 0)
	for _, temp := range strArr {
		arr := strings.Split(temp[0], ",")
		m = append(m, &AttributeSt{
			Type: StrToInt(arr[0]),
			Val:  StrToInt(arr[1]),
		})
	}
	return m
}
