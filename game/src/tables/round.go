package tables

import (
	"LollipopGo/log"
	gameDB "LollipopGo2.8x/data"
	twLibDBTable "github.com/Golangltd/Twlib/dbtable"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

var (
	roundOnce   sync.Once
	roundMsg    []interface{}
	RoundTables map[int]*RoundSt
)

// 关卡表配置
type RoundSt struct {
	Sid            int    // 关卡的Id
	NextId         int    // 下关卡Id
	RoundType      int    // 关卡类型 1 == 普通，3 == boss 关卡
	BattleMax      int    // 最大战斗回合
	SkipBattle     int    // 跳过回合数  0 == 不可以跳过
	MonsterGroupId int    // 怪物组Id
	RewardInfo     string // 奖励数据--通关奖励
	ProfitInfo     string // 收益数据--挂机奖励
}

// 获取关卡
func AddRoundTable() {
	data := gameDB.GetCFGameData(twLibDBTable.Gl_gk_gkpzb)
	if data == nil {
		log.Fatal("error! receive skillTables is nil!!")
	}
	roundMsg = data
	roundOnce.Do(initRoundTable)
}

func initRoundTable() {
	RoundTables = make(map[int]*RoundSt)
	for _, sample := range roundMsg {
		temp := sample.(map[string]interface{})
		m := &RoundSt{}
		m.Sid = StrToInt(temp["sid"].(string))
		m.NextId = StrToInt(temp["next_id"].(string))
		m.RoundType = StrToInt(temp["level_type"].(string))
		m.BattleMax = StrToInt(temp["round_max"].(string))
		m.SkipBattle = StrToInt(temp["skip_round"].(string))
		m.MonsterGroupId = StrToInt(temp["monster_group"].(string))
		m.RewardInfo = temp["award"].(string)
		m.ProfitInfo = temp["install_award"].(string) // 每小时收益奖励
		// 把每个技能对象存到map中
		RoundTables[m.Sid] = m
	}
	roundMsg = nil
}

type SRoundInfo struct {
	ItemId   int
	ItemType int
	ItemNum  int
}

// 获取奖励
func GetRoundFor3Row(data string) map[string]*SRoundInfo {
	retData := make(map[string]*SRoundInfo)
	re := regexp.MustCompile(`\d+,\d+,\d+`)
	strArr := re.FindAllStringSubmatch(data, -1)
	for i := 0; i < len(strArr); i++ {
		data := new(SRoundInfo)
		arr := strings.Split(strArr[i][0], ",")
		itype, _ := strconv.Atoi(arr[0])
		ikey, _ := strconv.Atoi(arr[1])
		val, _ := strconv.Atoi(arr[2])
		data.ItemId = ikey
		data.ItemType = itype
		data.ItemNum = val
		retData[arr[0]+"|"+arr[1]] = data
	}
	return retData
}

/*// 更新通关奖励
func SendDataInfo(roundid int, conn *websocket.Conn, stropenid string, userst *twlib_user.UserSt) {
	data := GetRoundFor3Row(RoundTables[roundid].RewardInfo)
	glog.Info("SendDataInfo-----", data)
	for _, v := range data {
		if v.ItemId == 1 { // 金币
			comm_proto.UpdateRoleCoin(conn, stropenid, v.ItemNum)
		} else if v.ItemId == 2 { // 砖石
			comm_proto.UpdateRoleDiamond(conn, stropenid, v.ItemNum)
		} else if v.ItemId == 3 { // 战队经验
		} else if v.ItemId == 4 { // 卡牌经验
			userst.RoleExp += v.ItemNum
		} else { // 道具
		}
	}
}
*/
