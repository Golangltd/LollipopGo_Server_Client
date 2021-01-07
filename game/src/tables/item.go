package tables

import (
	"LollipopGo2.8x/conf/g"
	lua_uitl "LollipopGo2.8x/lua/Uitl"
	_go "LollipopGo2.8x/lua/go"
)

// 道具表
type ItemSt struct {
	Sid       int         // 道具Id
	ItemType  int         // 道具类型
	BagType   int         // 背包类型
	Sale      int         // 是否可出售
	Parameter interface{} // 属性，不同类型道具有不同的属性
}

// 装备强化材料属性
type ItemTypeEquipStrengthenMaterialParameter struct {
	Exp          int // 强化石提供的经验
	CostItemType int // 消耗的道具类型
	CostItemID   int // 消耗的道具ID
	CostItemNum  int // 消耗的道具数量
}

var (
	ItemTables map[int]*ItemSt
)

func init() {
	ItemTables = make(map[int]*ItemSt)
	initItemTable()
}

func initItemTable() {
	ItemTables = make(map[int]*ItemSt)
	for _, sample := range _go.ItemConf {
		m := &ItemSt{}
		m.Sid = StrToInt(sample.Sid)
		m.ItemType = StrToInt(sample.Itemtype)
		m.BagType = StrToInt(sample.Bagtype)
		m.Sale = StrToInt(sample.Sale)
		// 解析属性
		if m.ItemType == g.ItemTypeEquipStrengthenMaterial { // 强化材料
			parameter := lua_uitl.GetFirstLuaDataFor4Row(sample.Parameter)
			//fmt.Println(parameter)
			if len(parameter) >= 4 {
				pst := &ItemTypeEquipStrengthenMaterialParameter{
					Exp:          parameter[0],
					CostItemType: parameter[1],
					CostItemID:   parameter[2],
					CostItemNum:  parameter[3],
				}
				m.Parameter = pst
			}
		}
		ItemTables[m.Sid] = m
	}
}
