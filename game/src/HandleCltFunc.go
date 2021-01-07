package main

import (
	impl "LollipopGo/network"
	"LollipopGo/util"
	"LollipopGo2.8x/conf"
	"LollipopGo2.8x/conf/g"
	. "LollipopGo2.8x/cxt"
	. "LollipopGo2.8x/data"
	gamedb "LollipopGo2.8x/data"
	"LollipopGo2.8x/handlers/OffLine"
	"LollipopGo2.8x/handlers/Round"
	"LollipopGo2.8x/models"
	. "LollipopGo2.8x/models"
	"LollipopGo2.8x/proto/comm_proto"
	ProtoGame "LollipopGo2.8x/proto/tw_proto"
	"LollipopGo2.8x/st"
	"LollipopGo2.8x/tables"
	"LollipopGo2.8x/temp"
	"LollipopGo2.8x/time/ticker"
	"LollipopGo2.8x/util"
	"LollipopGo2.8x/util/req"
	"encoding/json"
	"fmt"
	twChapter "github.com/Golangltd/Twlib/Chapter"
	"github.com/Golangltd/Twlib/DbSt"
	twVocation "github.com/Golangltd/Twlib/Vocation"
	"github.com/Golangltd/Twlib/errorcode"
	twLibItem "github.com/Golangltd/Twlib/item"
	twProto "github.com/Golangltd/Twlib/proto"
	twRewards "github.com/Golangltd/Twlib/rewards"
	twServer "github.com/Golangltd/Twlib/server"
	twUser "github.com/Golangltd/Twlib/user"
	"github.com/golang/glog"
	"github.com/robfig/cron"
	"golang.org/x/net/websocket"
	"strconv"
	"time"
)

// 卡牌图鉴表
func UserActivateGame(conn *websocket.Conn, ProtocolData map[string]interface{}) {

	strOpenId := ProtocolData["OpenId"].(string)

	data := &ProtoGame.GS2CUserActivateCard{
		Protocol:  twProto.GGameHallProto,
		Protocol2: ProtoGame.GS2CUserActivateCardProto2,
	}
	impl.PlayerSendToProxyServer(conn, data, strOpenId)

	// 更新图鉴表
	// TODO: 保存卡牌
}

func UserChooseMapChapter(conn *websocket.Conn, ProtocolData map[string]interface{}) {
	strOpenId := ProtocolData["OpenId"].(string)
	data := &ProtoGame.GS2CUserChooseMapChapter{
		Protocol:    twProto.GGameHallProto,
		Protocol2:   ProtoGame.GS2CUserChooseMapChapterProto2,
		ChapterList: nil,
	}

	// 在线人数
	var ddd []*twChapter.ChapterUserSt
	oldata := Round.GetUserOlInfo(0)
	for _, v := range oldata {
		chapterUserST := new(twChapter.ChapterUserSt)
		chapterUserST.RoleAvatar = v.RoleAvatar
		chapterUserST.RoleLev = v.RoleLev
		chapterUserST.RoleName = v.RoleName
		ddd = append(ddd, chapterUserST)
	}

	// 获取章节 Id
	vac, _ := M.Get(strOpenId + UserKey)
	val, _ := M.Get(strconv.Itoa(int(vac.(*models.Game).AccountId)) + "|User_R")
	iChapterId := val.(*models.Game).RoundInfo.ChapterId

	for i := 1; i <= iChapterId; i++ {
		var chapter = new(twChapter.ChapterSt)
		chapter.ChapterId = i
		chapter.ChapterState = twChapter.RoundUnLock
		// 通过章节ID 获取在线人数
		if true {

		}
		chapter.PlayerData = ddd
		if i == iChapterId {
			chapter.ChapterState = twChapter.ChapterCurrent
		}
		data.ChapterList = append(data.ChapterList, chapter)
	}
	impl.PlayerSendToProxyServer(conn, data, strOpenId)
}

func UserOffLineGet(conn *websocket.Conn, ProtocolData map[string]interface{}) {

	strOpenId := ProtocolData["OpenId"].(string)

	data := &ProtoGame.GS2CUserClickOnOffLine{
		Protocol:    twProto.GGameHallProto,
		Protocol2:   ProtoGame.GS2CUserClickOnOffLineProto2,
		OffLineTime: uint64(OffLine.Time),
		Rewards:     nil,
	}

	val, _ := M.Get(strOpenId + UserKey)
	if val != nil {

		inowtime := time.Now().Unix()
		glog.Info(" time.Now().Unix():", inowtime)
		itimereg := val.(*models.Game).UserInfo.RegisterTime
		glog.Info(" time.Now().Unix()itimereg:", itimereg)
		isubtime := (inowtime - itimereg) % 86400

		if isubtime < 43200 {
		} else {
			glog.Info("玩家数据大于12小时")
		}
		//data.OffLineTime = uint64(val.(*modules.Game).OfflineTime)
		data.OffLineTime = 0
		val.(*models.Game).OfflineTime = time.Now().Unix()

		dataRewards := SendDataDropInfo(val.(*models.Game), conn, strOpenId, isubtime)

		for _, v := range dataRewards {
			iRewatds := new(twRewards.RewardSt)
			iRewatds.ItemId = v.ItemId
			iRewatds.ItemNum = (v.ItemNum * int(isubtime)) / 3600.00
			iRewatds.ItemType = v.ItemType
			data.Rewards = append(data.Rewards, iRewatds)
		}
	}
	glog.Info("===============挂机的数量：", data.Rewards)
	impl.PlayerSendToProxyServer(conn, data, strOpenId)
}

// 更新掉落奖励
func SendDataDropInfo(game *models.Game, conn *websocket.Conn, stropenid string, isubtime int64) map[string]*tables.SRoundInfo {
	data := tables.GetRoundFor3Row(tables.RoundTables[game.RoundInfo.Round].ProfitInfo)
	glog.Info("SendDataDropInfo-----", data)
	items := make([]*twLibItem.ItemData, 0)
	for _, v := range data {
		if v.ItemType == twUser.ItemCommon { //  道具
			items = append(items, &twLibItem.ItemData{
				Type: twUser.ItemCommon,
				ID:   v.ItemId,
				Num:  (v.ItemNum / 60) * int(isubtime),
			})
		} else if v.ItemType == twUser.ItemEquipment { // 装备

			data := new(twUser.EquipSt)
			data.Num = (v.ItemNum / 60) * int(isubtime)
			data.ConfID = v.ItemId
			models.InitEquipPower(data)
			equiplist := []*twUser.EquipSt{data}
			comm_proto.UpdateEquip(conn, stropenid, equiplist)

		} else if v.ItemType == twUser.ItemCard { // 卡牌

		}
	}
	comm_proto.UpdateRoleItem(game, conn, stropenid, false, items)
	return data
}

func UserOffLine(ProtocolData map[string]interface{}) {
	strOpenId := ProtocolData["OpenId"].(string)
	data := &ProtoGame.GS2CUserOffLineBattle{
		Protocol:  twProto.GGameHallProto,
		Protocol2: ProtoGame.GS2CUserOffLineBattleProto2,
		Rewards:   nil,
	}
	val, _ := M.Get(strOpenId + UserKey)
	if val == nil {
		impl.PlayerSendToProxyServer(ConnXZ, data, strOpenId)
		return
	}
	roundID := val.(*models.Game).RoundInfo.Round
	dropData := tables.DropTable[roundID]
	if dropData == nil {
		return
	}

	inowtime := time.Now().Unix()
	//glog.Info(" time.Now().Unix():", inowtime)
	itimereg := val.(*models.Game).UserInfo.RegisterTime
	//glog.Info(" time.Now().Unix()itimereg:", itimereg)
	isubtime := (inowtime - itimereg) % 86400

	if isubtime < 43200 {
	} else {
		//glog.Info("玩家数据大于12小时")
	}

	if val.(*models.Game).OfflineTime != 0 {
		data.OffLineTime = uint64(inowtime - val.(*models.Game).OfflineTime)
	} else {
		data.OffLineTime = uint64(inowtime - val.(*models.Game).UserInfo.RegisterTime)
	}

	//glog.Info("玩家数据大于12小时", data.OffLineTime)

	data1 := tables.GetRoundFor3Row(tables.RoundTables[val.(*models.Game).RoundInfo.Round].ProfitInfo)
	// 获取数据
	for _, v := range data1 {
		rewardData := new(twRewards.RewardSt)
		rewardData.ItemType = v.ItemType
		rewardData.ItemId = v.ItemId
		rewardData.ItemNum = (v.ItemNum * int(isubtime)) / 3600.00
		OffLine.Rewatds = rewardData
		data.Rewards = append(data.Rewards, OffLine.Rewatds)
	}

	impl.PlayerSendToProxyServer(ConnXZ, data, strOpenId)
}

func SetSkill(conn *websocket.Conn, ProtocolData map[string]interface{}) {
	strOpenId := ProtocolData["OpenId"].(string)

	data := &ProtoGame.GS2CUserSetSkill{
		Protocol:  twProto.GGameHallProto,
		Protocol2: ProtoGame.GS2CUserSetSkillProto2,
	}
	impl.PlayerSendToProxyServer(conn, data, strOpenId)
}

func SkillUpGrade(conn *websocket.Conn, ProtocolData map[string]interface{}) {

	strOpenId := ProtocolData["OpenId"].(string)
	iCardUid := ProtocolData["CardUid"].(float64)
	iSkillId := ProtocolData["SkillId"].(float64)

	_ = iCardUid
	_ = iSkillId
	data := &ProtoGame.GS2CUserSkillUpGrade{
		Protocol:  twProto.GGameHallProto,
		Protocol2: ProtoGame.GS2CUserSetSkillProto2,
	}
	impl.PlayerSendToProxyServer(conn, data, strOpenId)
}

func UserCardBreak(conn *websocket.Conn, ProtocolData map[string]interface{}) {

	strOpenId := ProtocolData["OpenId"].(string)
	iCardId := ProtocolData["CardId"].(float64) // 数据Id

	_ = tables.RolesTable[int(iCardId)]
	// 卡牌突破
	data := &ProtoGame.GS2CUserCardUpGrades{
		Protocol:      twProto.GGameHallProto,
		Protocol2:     ProtoGame.GS2CUserCardUpGradeProto2,
		CardInfo:      nil,
		AttributeInfo: nil,
	}
	impl.PlayerSendMessageOfProxy(conn, data, strOpenId)
}

// 卡牌升级
func UserCardGrade(conn *websocket.Conn, ProtocolData map[string]interface{}) {
	strOpenId := ProtocolData["OpenId"].(string)
	iCardUid := ProtocolData["CardUid"].(float64)
	val, _ := M.Get(strOpenId + UserKey)
	//cardList := val.(*Game).UserInfo.CardList
	cardConfig := tables.WuShiTables
	fmt.Println("玩家拥有的卡牌：", val.(*Game).UserInfo.CardList)
	sendData := new(twUser.CardInfo)
	isBreak := false
	for _, v := range val.(*Game).UserInfo.CardList {
		if true{
			// 消耗 --
			strCost := cardConfig[v.Level].Cost
			constData := tables.GetWuShiFor3Row(strCost)
			fmt.Println("升级卡牌的消耗:", constData)
			roleData := tables.RolesTable[int(v.CardID)]
			// 是否突破
			if cardConfig[v.Level].LevelUp == "1" {
				isBreak = true
				v.Stars += 1
				if v.Stars > 5 {
					v.Stars = 5
				}
				// 获取突破的属性信息 error
			}
			v.Level++
			// 判断等级 是否解锁
			UnlockingNotice(conn, ProtocolData, int(v.CardID), v.Level)
			if !isBreak {
				if roleData.Vocation == twVocation.Warrior {
					data := tables.GetWuShiFor2Row(cardConfig[v.Level].Attribute1)
					v.AttributeInfo.HPPower += uint64(data[1])
					v.AttributeInfo.AttackPower += uint64(data[2])
					v.AttributeInfo.DefensePower += uint64(data[3])
				} else if roleData.Vocation == twVocation.Ranger {
					data := tables.GetWuShiFor2Row(cardConfig[v.Level].Attribute2)
					v.AttributeInfo.HPPower += uint64(data[1])
					v.AttributeInfo.AttackPower += uint64(data[2])
					v.AttributeInfo.DefensePower += uint64(data[3])
				} else if roleData.Vocation == twVocation.Mage {
					data := tables.GetWuShiFor2Row(cardConfig[v.Level].Attribute3)
					v.AttributeInfo.HPPower += uint64(data[1])
					v.AttributeInfo.AttackPower += uint64(data[2])
					v.AttributeInfo.DefensePower += uint64(data[3])
				} else if roleData.Vocation == twVocation.Assist {
					data := tables.GetWuShiFor2Row(cardConfig[v.Level].Attribute4)
					v.AttributeInfo.HPPower += uint64(data[1])
					v.AttributeInfo.AttackPower += uint64(data[2])
					v.AttributeInfo.DefensePower += uint64(data[3])
				}
			} else {
				// 突破属性加成
				if roleData.Vocation == twVocation.Warrior {
					fmt.Println("战士")
					data := tables.GetWuShiFor2Row(cardConfig[v.Level].Attribute5)
					v.AttributeInfo.HPPower += uint64(data[1])
					v.AttributeInfo.AttackPower += uint64(data[2])
					v.AttributeInfo.DefensePower += uint64(data[3])
				} else if roleData.Vocation == twVocation.Ranger {
					fmt.Println("游侠")
					data := tables.GetWuShiFor2Row(cardConfig[v.Level].Attribute6)
					v.AttributeInfo.HPPower += uint64(data[1])
					v.AttributeInfo.AttackPower += uint64(data[2])
					v.AttributeInfo.DefensePower += uint64(data[3])
				} else if roleData.Vocation == twVocation.Mage {
					fmt.Println("法师")
					data := tables.GetWuShiFor2Row(cardConfig[v.Level].Attribute7)
					v.AttributeInfo.HPPower += uint64(data[1])
					v.AttributeInfo.AttackPower += uint64(data[2])
					v.AttributeInfo.DefensePower += uint64(data[3])
				} else if roleData.Vocation == twVocation.Assist {
					fmt.Println("辅助")
					data := tables.GetWuShiFor2Row(cardConfig[v.Level].Attribute8)
					v.AttributeInfo.HPPower += uint64(data[1])
					v.AttributeInfo.AttackPower += uint64(data[2])
					v.AttributeInfo.DefensePower += uint64(data[3])
				}
			}
			sendData = v
			break
		} else {
			fmt.Println("玩家拥有的卡牌不存在：", iCardUid)
		}
	}
	// 获取卡牌升级
	data := &ProtoGame.GS2CUserCardUpGrades{
		Protocol:      twProto.GGameHallProto,
		Protocol2:     ProtoGame.GS2CUserCardUpGradeProto2,
		CardInfo:      sendData,
		AttributeInfo: sendData.AttributeInfo,
		IsBreak:       isBreak,
	}
	impl.PlayerSendToProxyServer(conn, data, strOpenId)
	// 保存数据库
	gamedb.UpdateUserCardInfo(sendData)
	return
}

func UnlockingNotice(conn *websocket.Conn, ProtocolData map[string]interface{}, carduid int, level int) {

	strOpenId := ProtocolData["OpenId"].(string)
	data := &ProtoGame.GS2CUserSkillUpGrade{
		Protocol:  twProto.GGameHallProto,
		Protocol2: ProtoGame.GS2CUserSkillUpGradeProto2,
		SkillInfo: nil,
	}

	dataskill := tables.RolesTable[carduid].Skills[level]
	if dataskill != 0 {
		skilldata := new(twUser.SkillInfo)
		skilldata.Position = -1
		skilldata.SkillId = int64(dataskill)
		skilldata.SkillLev = level
		impl.PlayerSendToProxyServer(conn, data, strOpenId)
	}
}

// 玩家创建角色数据
func UserRegisterInfo(conn *websocket.Conn, ProtocolData map[string]interface{}) {

	strOpenId := ProtocolData["OpenId"].(string)
	iSex := ProtocolData["Sex"].(float64)
	strName := ProtocolData["Name"].(string)
	val, _ := M.Get(strOpenId + UserKey)
	accountId := val.(*models.Game).AccountId
	// 数据库查询
	data := CreateGameRoleInfo(&twlib_DbSt.RoleSt{AccountId: accountId, OpenId: strOpenId, Sex: int(iSex), Name: strName})

	if data.RoleUid == -1 {
		senate := twProto.GameError{
			Protocol:  twProto.GErrorProto,
			Protocol2: 1,
			ErrorCode: errorcode.RepeatName,
		}
		impl.PlayerSendToProxyServer(conn, senate, strOpenId)
		return
	}

	data.CardList = []*twUser.CardInfo{}
	data.EquipData = new(twUser.EquipData)
	data.EquipData.EquipSts = []*twUser.EquipSt{}
	data.RegisterTime = data.RegisterTime
	val.(*models.Game).UserInfo = data


	//TODO:destroy --
	val.(*models.Game).UserTicker = make(map[ticker.TickerUid]*cron.Cron)
	val.(*models.Game).UserTicker[ticker.TickerUid(data.RoleUid)] = cron.New()

	time.AfterFunc(5*time.Second, func() {
		// 发送角色初始化邮件
		sendMailMsg := twProto.NewS2CSSendTemplateMailToPlayerMsg()
		sendMailMsg.OpenID = strOpenId
		sendMailMsg.UID = data.RoleUid
		sendMailMsg.MailTemplateID = tables.GFirstEnterGameMailTemplateID
		glog.Infof("发送初始化邮件请求:%+v\n", sendMailMsg)
		impl.PlayerSendMessageOfProxy(ConnXZ, sendMailMsg, util.MD5_LollipopGO(strconv.Itoa(twServer.CenterServerId)))
	})
	// 发送玩家数据到中心服
	userDataMsg := comm_proto.GS2S_GetUserSt{
		Protocol:  twProto.GGameHallProto,
		Protocol2: comm_proto.GS2S_GetUserStProto2,
		OpenId:    strOpenId,
		UserInfo:  data,
	}
	impl.PlayerSendMessageOfProxy(ConnXZ, userDataMsg, util.MD5_LollipopGO(strconv.Itoa(twServer.CenterServerId)))

	// 发送数据库，保存数据；确认数据库数据存在
	senate := ProtoGame.GS2CUserRegister{
		Protocol:  twProto.GGameHallProto,
		Protocol2: ProtoGame.GS2CUserRegisterProto2,
		PlayerSt:  data,
	}
	impl.PlayerSendToProxyServer(conn, senate, strOpenId)
}

// 获取卡牌信息
func UserClickCardGetInfo(ProtocolData map[string]interface{}) {
	strOpenId := ProtocolData["OpenId"].(string)
	strCardId := ProtocolData["CardId"].(float64)
	_ = strOpenId
	_ = strCardId
	// 数据更新操作
}

// 点击战斗按钮-- 进入挂机流程，挂机收益等
func UserClickBattleBTN(ProtocolData map[string]interface{}) {
	strOpenid := ProtocolData["OpenId"].(string)
	// 主角的随机，随机的生成规则/
	_ = strOpenid
	// 获取数据操作
}

// 玩家登录 --- 需要区DB验证
// token数据正确与否
func UserLoginGame(conn *websocket.Conn, ProtocolData map[string]interface{}) {
	// 解析客户端发过来的数据
	strOpenid := ProtocolData["OpenId"].(string) // 角色的唯一ID信息，server 加密的数据 MD5
	strToken := ProtocolData["Token"].(string)   // 从登陆服--游戏客户端--反向代理服务器--游戏主逻辑服务器
	// 需要校验token的正确性,http get 请求 login server 校验,获取GM数据
	loginURL := req.AddParamsToGetReq(g.Http, conf.ServerConfig().GetLoginUrlList(), map[string]string{"token": strToken})
	glog.Infof("connect to loginURL:%s\n", loginURL)
	rest := impl.Get(loginURL)
	glog.Info("impl.Get", rest)
	ulcerate := &twUser.UserSt{}
	err := json.Unmarshal([]byte(rest), ulcerate)
	if err != nil {
		glog.Error(fmt.Sprintf("dont't json.unmarshal rest,err is %s", err))
	}
	// 读取数据库,有无这个角色
	accountID := ulcerate.RoleUid // 获取有无角色数据,数据操作；应用操作！---- 数据操作
	roleInfo := GetGameRoleInfo(accountID)
	//返回游戏玩家的结构数据
	data := &ProtoGame.GS2CUserLogin{
		Protocol:  twProto.GGameHallProto,
		Protocol2: ProtoGame.GS2CUserLoginProto2,
		PlayerSt:  roleInfo, // 为空就是没有角色信息，需要创建角色数据
	}
	if len(temp.GMapBag[accountID]) == 0 {
		temp.GMapBag[accountID] = roleInfo.ItemList
	} else {
		roleInfo.ItemList = temp.GMapBag[accountID]
	}

	// 背包装备数据
	if _, ok := temp.GMapEquipBag[accountID]; !ok {
		temp.CreateTestBagEquip(accountID)
		temp.GMapEquipBag[accountID] = roleInfo.EquipData
	}
	Round.SetUserOlInfo(roleInfo)
	roleInfo.RegisterTime = roleInfo.RegisterTime
	impl.PlayerSendToProxyServer(ConnXZ, data, strOpenid)
	// --------------------------------------------------------------------------
	if roleInfo.RoleUid > 0 {
		// 发送玩家数据到中心服
		userDataMsg := comm_proto.GS2S_GetUserSt{
			Protocol:  twProto.GGameHallProto,
			Protocol2: comm_proto.GS2S_GetUserStProto2,
			OpenId:    strOpenid,
			UserInfo:  roleInfo,
		}
		impl.PlayerSendMessageOfProxy(ConnXZ, userDataMsg, util.MD5_LollipopGO(strconv.Itoa(twServer.CenterServerId)))
	}

	gameUserInfo := &models.Game{
		Connection:  conn,
		StrMD5:      strOpenid,
		MapSafe:     M,
		UserInfo:    roleInfo,
		AccountId:   accountID,
		RoundInfo:   new(st.RoundSt),
		OfflineTime: roleInfo.RegisterTime,
		UserTicker:  make(map[ticker.TickerUid]*cron.Cron),
		CacheDB:     CacheGame,
	}

	gameUserInfo.UserTicker[ticker.TickerUid(roleInfo.RoleUid)] = cron.New()
	gameUserInfo.SRModel = NewSRModel(gameUserInfo.StrMD5, gameUserInfo)
	//  初始化关卡数据
	val, _ := M.Get(strconv.Itoa(int(accountID)) + "|User_R")
	roundST := new(st.RoundSt)
	if roleInfo.RoleUid != 0 {
		//roleInfo.ChapterInfo = new(twUser.ChapterInfo)
		roundST.Sid = roleInfo.ChapterInfo.RoundId
		roundST.ChapterId = roleInfo.ChapterInfo.ChapterId
		roundST.ChapterId2 = roleInfo.ChapterInfo.ChapterId2
		roundST.Round = roleInfo.ChapterInfo.RoundId
	} else {
		roundST.Sid = 1
		roundST.ChapterId = 1
		roundST.ChapterId2 = 1
		roundST.Round = 1
	}
	if val != nil {
		roundST.Sid = val.(*models.Game).RoundInfo.Sid
		roundST.ChapterId = val.(*models.Game).RoundInfo.ChapterId
		roundST.ChapterId2 = val.(*models.Game).RoundInfo.ChapterId2
		roundST.Round = val.(*models.Game).RoundInfo.Round
	}
	gameUserInfo.RoundInfo = roundST
	_, _ = M.Put(strconv.Itoa(int(accountID))+"|User_R", gameUserInfo)
	_, _ = M.Put(strOpenid+UserKey, gameUserInfo)
	// 保存公用连接
	game_util.SetGUserInfo(strOpenid, accountID)
}

// 玩家选择入口，功能入口
func UserFunctionId(conn *websocket.Conn, ProtocolData map[string]interface{}) {
	//strOpenId := ProtocolData["OpenId"].(string)
	//
	//senate := ProtoGame.GS2CUserGetName{
	//	Protocol:  twProto.GGameHallProto,
	//	Protocol2: ProtoGame.GS2CUserGetNameProto2,
	//	Name:      "随机姓名测试" + strconv.Itoa(NameCount),
	//}
	//impl.PlayerSendMessageOfProxy(conn, senate, strOpenId)
	//NameCount++
	//return
}

// 玩家选择关卡
func UserChooseGameRound(conn *websocket.Conn, ProtocolData map[string]interface{}) {
	strOpenId := ProtocolData["OpenId"].(string)
	senate := ProtoGame.GS2CUserChooseRound{
		Protocol:  twProto.GGameHallProto,
		Protocol2: ProtoGame.GS2CUserChooseRoundProto2,
	}
	impl.PlayerSendToProxyServer(conn, senate, strOpenId)
	return
}

// 玩家选择战斗开始  game server 转发给 battle server
func UserChooseStartBattle(conn *websocket.Conn, ProtocolData map[string]interface{}) {

}

// battle server 发送数据到 game server 转发到 游戏客户端
func BattleSendDataGameStartBattle(conn *websocket.Conn, ProtocolData map[string]interface{}) {

}
