package temp

import (
	"LollipopGo2.8x/tables"
	twlib_user "github.com/Golangltd/Twlib/user"
	"strconv"
	"strings"
)

// 背包临时数据
var GMapBag map[int64][]*twlib_user.ItemData
var GMapCardBag map[int64][]*twlib_user.CardInfo
var GMapEquipBag map[int64]*twlib_user.EquipData
var GmapRoleUID map[int64]string
var GmapRoleCoin map[int64]int64
var GmapRoleDiom map[int64]int64

func init() {
	GMapBag = make(map[int64][]*twlib_user.ItemData)
	GMapCardBag = make(map[int64][]*twlib_user.CardInfo)
	GMapEquipBag = make(map[int64]*twlib_user.EquipData)
	GmapRoleUID = make(map[int64]string)
	GmapRoleCoin = make(map[int64]int64)
	GmapRoleDiom = make(map[int64]int64)

	/*
	       var Item1 []*twlib_user.ItemSt
	   	var Item2 []*twlib_user.ItemData
	   	// 增加数据
	   	Item2 = append(Item2, &twlib_user.ItemData{
	   		1,
	   		4405,
	   		20,
	   	})
	   	Item1 = append(Item1, &twlib_user.ItemSt{FunctionId: 41, ItemData: Item2})  // 背包
	   	GMapBag["1"] = Item1
	*/

	//------------------------------------------------------------------------------------------------------------------
}

//// 创建
//func CreateItem(accountid int64, roleuid int64) []*twlib_user.ItemData {
//	datas := tables.GMapGlobalTable["16"]
//	stringdaya := FenGe(datas.Data1)
//	var Item5 []*twlib_user.ItemData
//	for _, v := range stringdaya {
//		var Item3 []*twlib_user.ItemData
//		var item4 = new(twlib_user.ItemData)
//
//		// 离散表
//		icardid, _ := strconv.Atoi(v)
//		uid := gameDB.SaveItem(roleuid, accountid, item4)
//		Item3 = append(Item3, &twlib_user.ItemData{
//			int64(uid),
//			icardid,
//			1,
//			int64(10),
//		})
//		Item5 = append(Item5, Item3...) // 卡牌背包
//	}
//
//	GMapBag[accountid] = Item5
//
//	return GMapBag[accountid]
//}

// 创建测试卡牌
func CreateTestCard(accountid int64) []*twlib_user.CardInfo {
	var Item3 []*twlib_user.CardInfo
	// 属性加成
	att := &twlib_user.AttributeSt{
		BattlePower:  1000, // 战斗力
		HPPower:      100,  // 血量
		AttackPower:  100,  // 攻击力
		DefensePower: 100,  // 防御力
	}

	cardid := 0
	datas := tables.GMapGlobalTable["16"]
	stringdaya := FenGe(datas.Data1)

	for _, v := range stringdaya {
		icardid, _ := strconv.Atoi(v)
		cardid = icardid
		if cardid == 0 {
			cardid = 1006
		}

		c := &twlib_user.CardInfo{
			// 卡牌唯一ID
			Level:         1,                                // 卡牌等级
			Skills:        make([]*twlib_user.SkillInfo, 0), // 技能列表
			Equips:        make([]*twlib_user.EquipSt, 6),
			AttributeInfo: att, // 战斗力
			Stars:         0,
		}
		c.Skills = append(c.Skills, CreateSkillInfo())
		Item3 = append(Item3, c)
	}
	GMapCardBag[accountid] = Item3
	//---------------------------------------------------------------------------------------
	return GMapCardBag[accountid]
}

// 创建测试装备
func CreateTestBagEquip(accountid int64) {
	equips := make([]*twlib_user.EquipSt, 0)
	userEquip := &twlib_user.EquipData{
		FunctionId: 101,
		EquipSts:   equips,
	}
	GMapEquipBag[accountid] = userEquip
}

func FenGe(data string) []string {
	strArr := strings.Split(data, `,`)
	// 输出 字符串数组 中的 字符串
	for _, str := range strArr {
		println(str)
	}
	return strArr
}

/*
   1. 数据更新操作，应用开始！
   2. 数据个别新
   3. 数据操作
   4. 数据更新，数据更新操作，数据操作；应用更新
   5. 数据跟新... ...
   6. 测试数据跟操作，数据操纵！
   7. 数据更新操作，主要的应用更新！数据变化更新，数据分析更新操作，更新操作！
   8. 数据更新操作，主要的开发设备!更新操作应用，数据分析操作，应用更新！更新操纵数据分析变化！
   9. 测试ok，数据更新操纵，主要的测试数据
*/

func GetDataNowFive() {
}
