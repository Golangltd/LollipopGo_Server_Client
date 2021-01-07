package ProtoGame

import (
	"github.com/Golangltd/Twlib/Chapter"
	twlib_rewards "github.com/Golangltd/Twlib/rewards"
	twlib_user "github.com/Golangltd/Twlib/user"
)

// 主协议 6 ，游戏主要逻辑

const (
	INIT                = iota // INIT == 0
	C2GSUserLoginProto2        // C2GSUserLoginProto2 == 1 玩家登录协议
	GS2CUserLoginProto2        // GS2CUserLoginProto2 == 2 返回玩家数据操作

	C2GSUserFunctionProto2 // C2GSUserFunctionProto2 == 3  玩家发送功能Id,进入不同功能的入口
	GS2CUserFunctionProto2 // GS2CUserFunctionProto2 == 4  暂时不用

	C2GSUserChooseMapChapterProto2 // C2GSUserChooseMapChapterProto2 == 5 玩家选择地图,大地图章节
	GS2CUserChooseMapChapterProto2 // GS2CUserChooseMapChapterProto2 == 6 返回玩家数据操作

	C2GSUserChooseChapterProto2 // C2GSUserChooseChapterProto2 == 7 玩家选择章节
	GS2CUserChooseChapterProto2 // GS2CUserChooseChapterProto2 == 8 返回玩家数据操作本章节的所有关卡

	C2GSUserChooseRoundProto2 // C2GSUserChooseRoundProto2 == 9 玩家选择关卡
	GS2CUserChooseRoundProto2 // GS2CUserChooseRoundProto2 == 10 返回玩家数据操作

	G2BSUserBattleProto2 // G2BSUserBattleProto2 == 11 主逻辑服务器发送数据到战斗服
	BS2GUserBattleProto2 // BS2GUserBattleProto2 == 12

	C2GSUserStartBattleProto2 // C2GSUserStartBattleProto2 == 13  玩家开始战斗 ---> 触发 7号协议
	GS2CUserStartBattleProto2 // GS2CUserStartBattleProto2 == 14  返回

	C2GSUserOffLineBattleProto2 // C2GSUserOffLineBattleProto2 == 15  进入挂机系统
	GS2CUserOffLineBattleProto2 // GS2CUserOffLineBattleProto2 == 16  服务器返回挂机数据

	C2GSUserBagProto2 // C2GSUserBagProto2 == 17 游戏客户端点击背包
	GS2CUserBagProto2 // GS2CUserBagProto2 == 18 发送全部道具信息

	C2GSUserChickCardProto2 // C2GSUserChickCardProto2 == 19 游戏客户端点击卡牌
	GS2CUserChickCardProto2 // GS2CUserChickCardProto2 == 20 发送卡牌属性

	C2GSUserRegisterProto2 // C2GSUserRegisterProto2 == 21 玩家创建角色
	GS2CUserRegisterProto2 // GS2CUserRegisterProto2 == 22 返回创建数据信息

	GS2CUserOpAddItemProto2    // GS2CUserOpAddItemProto2 == 23 道具增加
	GS2CUserOpDelItemProto2    // GS2CUserOpDelItemProto2 == 24 道具删除
	GS2CUserOpChangeItemProto2 // GS2CUserOpChangeItemProto2 == 25 道具改变

	C2GSUserPlayProto2 // C2GSUserPlayProto2 == 26 游戏客户端点击【Play】开始战斗，战斗服务器回复战斗数据了

	C2GSUserGetNameProto2 // C2GSUserGetNameProto2 == 27 获取姓名，第一次进来可以直接发
	GS2CUserGetNameProto2 // GS2CUserGetNameProto2 == 28 返回获取姓名信息

	C2GSUserCheckNameProto2 // C2GSUserCheckNameProto2 == 29 玩家输入姓名后，前端发送检测
	GS2CUserCheckNameProto2 // GS2CUserCheckNameProto2 == 30 返回获取是否重名

	C2GSUserCardUpGradeProto2 // C2GSUserCardUpGradeProto2 == 31 点击卡牌升级
	GS2CUserCardUpGradeProto2 // GS2CUserCardUpGradeProto2 == 32 返回卡牌升级,突破信息

	C2GSUserCardUpGradeBreakProto2 // C2GSUserCardUpGradeBreakProto2 == 33 点击卡牌突破

	C2GSUserClickOnOffLineProto2 // C2GSUserClickOnOffLineProto2 == 34 点击获取奖励
	GS2CUserClickOnOffLineProto2 // GS2CUserClickOnOffLineProto2 == 35 返回奖励信息

	C2GSUserActivateCardProto2 // C2GSUserActivateCardProto2 == 36 游戏客户端点击图鉴系统--激活按钮
	GS2CUserActivateCardProto2 // GS2CUserActivateCardProto2 == 37 返回数据加成--成功

	C2GSUserPlayRetProto2 // C2GSUserPlayProto2 == 38 游戏客户端点击【Play】开始战斗，战斗服务器回复战斗数据了,返回战斗结算数据

	C2GSUserSkillUpGradeProto2 // C2GSUserSkillUpGradeProto2 == 39 点击卡牌技能升级
	GS2CUserSkillUpGradeProto2 // GS2CUserSkillUpGradeProto2 == 40 返回卡牌技能升级

	C2GSUserSetSkillProto2 // C2GSUserSetSkillProto2 == 41 设置技能
	GS2CUserSetSkillProto2 // GS2CUserSetSkillProto2 == 42

	// 50 - 100 装备系统
	Proto2EquipBegin    = 50 // ------装备消息开始值
	C2GSEquipWearProto2 = 50 // 客户端请求 穿戴/一键穿戴
	C2GSEquipTakeOff    = 51 // 客户端请求 脱下/一键脱下
	C2GSEquipReplace    = 52 // 客户端请求 替换
	C2GSEquipStrengthen = 53 // 客户端请求 强化
	GS2CCardEquipUpdate = 54 // 服务器返回 卡牌装备更新
	GS2CCardEquipDel    = 55 // 服务器返回 卡牌装备删除
	// 100-150 装备背包
	GS2CEquipBagUpdate = 101 // 服务器返回 装备背包 装备更新
	Proto2EquipEnd     = 150 // ------装备消息结束值

)
//-----------------------------------------------------------------------------
// C2GSUserSetSkillProto2 == 41 设置技能
type C2GSUserSetSkill struct {
	Protocol  int
	Protocol2 int
	OpenId    string
	CardUid   uint64
	SkillList []int64  // 技能Id
}

// GS2CUserSetSkillProto2 == 42
type GS2CUserSetSkill struct {
	Protocol  int
	Protocol2 int
}
//-----------------------------------------------------------------------------
 // C2GSUserSkillUpGradeProto2 == 39 点击卡牌技能升级
type C2GSUserSkillUpGrade struct {
	Protocol  int
	Protocol2 int
	OpenId    string
	CardUid   uint64
	SkillId   uint64
}
 // GS2CUserSkillUpGradeProto2 == 40 返回卡牌技能升级
type GS2CUserSkillUpGrade struct {
	Protocol  int
	Protocol2 int
	SkillInfo []*twlib_user.SkillInfo
}

//------------------------------------------------------------------------------
// C2GSUserPlayProto2 == 38 游戏客户端点击【Play】开始战斗，-,返回战斗结算数据
type C2GSUserPlayRet struct {
	Protocol  int
	Protocol2 int
	IsVictory bool
	Rewards   []*twlib_user.ItemSt
	Exp       int
}

//------------------------------------------------------------------------------
// C2GSUserActivateCardProto2 == 36 游戏客户端点击图鉴系统
type C2GSUserActivateCard struct {
	Protocol  int
	Protocol2 int
	OpenId    string
	CardId    uint64 // 卡牌id,查找图鉴Id,返回奖励
}

// GS2CUserActivateCardProto2 == 37 返回数据加成成功
type GS2CUserActivateCard struct {
	Protocol  int
	Protocol2 int
}

//------------------------------------------------------------------------------
// C2GSUserClickOnOffLineProto2 == 34 点击获取奖励
type C2GSUserClickOnOffLine struct {
	Protocol  int
	Protocol2 int
	OpenId    string
}

// GS2CUserClickOnOffLineProto2 == 35 返回奖励信息
type GS2CUserClickOnOffLine struct {
	Protocol    int
	Protocol2   int
	OffLineTime uint64
	Rewards     []*twlib_rewards.RewardSt
}

//------------------------------------------------------------------------------
// C2GSUserCardUpGradeBreakProto2 == 33 点击卡牌突破
type C2GSUserCardUpGradeBreak struct {
	Protocol  int
	Protocol2 int
	OpenId    string // 唯一Id
	CardId    int    // 卡牌Id
}

//------------------------------------------------------------------------------
// C2GSUserCardUpGradeProto2 == 31 点击卡牌升级
type C2GSUserCardUpGrade struct {
	Protocol  int
	Protocol2 int
	OpenId    string // 唯一Id
	CardUid   int    // 卡牌Iid
}

// GS2CUserCardUpGradeProto2 == 32 返回卡牌升级信息
type GS2CUserCardUpGrades struct {
	Protocol      int
	Protocol2     int
	CardInfo      *twlib_user.CardInfo
	AttributeInfo *twlib_user.AttributeSt //  属性加成变化的
	IsBreak       bool
}

//------------------------------------------------------------------------------
// C2GSUserCheckNameProto2 == 29 玩家输入姓名后，前端发送检测
type C2GSUserCheckName struct {
	Protocol  int
	Protocol2 int
	OpenId    string // 唯一Id
	Name      string // 玩家输入的姓名
}

// GS2CUserCheckNameProto2 == 28 返回获取是否重名
type GS2CUserCheckName struct {
	Protocol  int
	Protocol2 int
	BRegister bool // BRegister = true 可以注册；BRegister = false 不可以注册
}

//------------------------------------------------------------------------------
// C2GSUserGetNameProto2 == 27 获取姓名，第一次进来可以直接发
type C2GSUserGetName struct {
	Protocol  int
	Protocol2 int
	Sex       int    // 0:表示选择女角色，1:表示选择男角色
	OpenId    string // 唯一Id
}

// GS2CUserGetNameProto2 == 28 返回获取姓名信息
type GS2CUserGetName struct {
	Protocol  int
	Protocol2 int
	Name      string // 随机的名字
}

//------------------------------------------------------------------------------
// C2GSUserPlayProto2 == 26 游戏客户端点击【Play】开始战斗，战斗服务器回复战斗数据了
type C2GSUserPlays struct {
	Protocol  int
	Protocol2 int
	IsPvP     bool      // 是否是pvp
	UserA     *UserInfo // 玩家A(一定是正式玩家)    --- 需要保存数据库 --- 记录下卡牌的数据  --- 上阵阵容
	UserB     *UserInfo // 玩家B(可以是玩家，也可以是怪物)
}

type UserInfo struct {
	OpenId string
	Level  int
	Cards  map[int]*CardInfo // 卡牌信息
}

type CardInfo struct {
	CardID   uint64
	Level    int
	RoleID   int
	Position int
	Skills   []int
	Sex      int // 性别 1男 2女
}

//------------------------------------------------------------------------------
// GS2CUserOpAddItemProto2 == 23 道具操作
type GS2CUserOpAddItem struct {
	Protocol   int
	Protocol2  int
	FunctionId int // 功能Id，每个功能Id都是不一样的，且是唯一的标识
	ItemId     int // 道具Id
	ItemNum    int // 道具数量
	ItemType   int // 道具类型
}

// GS2CUserOpDelItemProto2 == 24 道具减少
type GS2CUserOpDelItem struct {
	Protocol   int
	Protocol2  int
	FunctionId int // 功能Id，每个功能Id都是不一样的，且是唯一的标识
	ItemId     int // 道具Id
	ItemNum    int // 道具数量
	ItemType   int // 道具类型
}

// GS2CUserOpChangeItemProto2 == 25 道具改变
type GS2CUserOpChangeItem struct {
	Protocol   int
	Protocol2  int
	FunctionId int // 功能Id，每个功能Id都是不一样的，且是唯一的标识
	ItemId     int // 道具Id
	ItemNum    int // 道具数量
	ItemType   int // 道具类型
}

//------------------------------------------------------------------------------
// C2GSUserRegisterProto2 == 21 玩家创建角色
type C2GSUserRegister struct {
	Protocol  int
	Protocol2 int
	OpenId    string // 通信的唯一Id
	Sex       int    // 性别
	Name      string // 名字,名字表
}

// 返回注册结果
type GS2CUserRegister struct {
	Protocol  int
	Protocol2 int
	PlayerSt  *twlib_user.UserSt // 玩家的结构体数据
}

//------------------------------------------------------------------------------
// C2GSUserChickCardProto2 == 19 游戏客户端点击卡牌
type C2GSUserChickCard struct {
	Protocol   int
	Protocol2  int
	OpenId     string // 玩家唯一Id
	FunctionId int    // 功能Id，每个功能Id都是不一样的，且是唯一的标识
	CardId     int    // 卡牌Id
}

// GS2CUserChickCardProto2 == 20 发送卡牌属性信息
type GS2CUserChickCard struct {
	Protocol  int
	Protocol2 int
	CardInfo  *twlib_user.CardInfo // 卡牌信息，数据操作
}

//------------------------------------------------------------------------------
// C2GSUserFunctionProto2 == 3 玩家发送功能Id,进入不同功能的入口; 根据不同的Id返回不同的功能
type C2GSUserFunction struct {
	Protocol   int
	Protocol2  int
	OpenId     string // 玩家唯一Id
	FunctionId int    // 功能Id，每个功能Id都是不一样的，且是唯一的标识
}

//------------------------------------------------------------------------------
// C2GSUserBagProto2 == 15 游戏客户端点击背包
type C2GSUserBag struct {
	Protocol  int
	Protocol2 int
	OpenId    string // 玩家唯一Id
}

// GS2CUserBagProto2 == 16  服务器返回数据,数据操作
type GS2CUserBag struct {
	Protocol      int
	Protocol2     int
	CardAttribute []*twlib_user.ItemSt // 卡牌属性数据
}

//------------------------------------------------------------------------------
// C2GSUserChooseChapterProto2 == 5 玩家选择关卡  --- 进入战斗系统
type C2GSUserChooseChapter struct {
	Protocol  int
	Protocol2 int
	OpenId    string // 玩家唯一Id
	ChapterId int    // 章节的Id
}

// GS2CUserChooseChapterProto2 == 6 返回玩家数据操作
type GS2CUserChooseChapter struct {
	Protocol  int
	Protocol2 int
	RoundList []*Twlib_Chapter.RoundSt // 关卡的结构体
}

//------------------------------------------------------------------------------
// C2GSUserChooseRoundProto2 == 7 玩家选择关卡  --- 进入战斗系统
type C2GSUserChooseRounds struct {
	Protocol  int
	Protocol2 int
	OpenId    string // 玩家唯一Id
	RoundId   int    // 关卡的Id
}

// GS2CUserChooseRoundProto2 == 8 返回玩家数据操作，进入战前选择怪状态
type GS2CUserChooseRound struct {
	Protocol  int
	Protocol2 int
}

//------------------------------------------------------------------------------
// C2GSUserChooseMapChapterProto2 == 3 玩家选择地图,大地图章节,进入章节的地图
type C2GSUserChooseMapChapter struct {
	Protocol   int
	Protocol2  int
	OpenId     string // 玩家唯一Id
	FunctionId int    // 功能ID = 211
}

// GS2CUserChooseMapChapterProto2 == 4 返回玩家数据操作，章节列表
type GS2CUserChooseMapChapter struct {
	Protocol    int
	Protocol2   int
	ChapterList []*Twlib_Chapter.ChapterSt // 章节的结构体
}

//------------------------------------------------------------------------------
// C2GSUserOffLineBattleProto2 == 15  进入挂机系统
type C2GSUserOffLineBattle struct {
	Protocol  int
	Protocol2 int
	OpenId    string // 玩家唯一Id
}

// GS2CUserOffLineBattleProto2 == 16  服务器返回挂机奖励数据
type GS2CUserOffLineBattle struct {
	Protocol    int
	Protocol2   int
	OffLineTime uint64                    // 时间戳
	Rewards     []*twlib_rewards.RewardSt // 奖励数据
}

//------------------------------------------------------------------------------
// C2GSUserLoginProto2 == 1  游戏玩家登录
type C2GSUserLogin struct {
	Protocol  int
	Protocol2 int
	OpenId    string // 玩家唯一Id,临时的
	Token     string
}

// GS2CUserLoginProto2 == 2 游戏客户端返回给玩家的数据
type GS2CUserLogin struct {
	Protocol  int
	Protocol2 int
	PlayerSt  *twlib_user.UserSt
}

//------------------------------------------------------------------------------
// C2GSUserStartBattleProto2 == 9  玩家开始战斗 ---> 触发 7号协议
type C2GSUserStartBattle struct {
	Protocol   int
	Protocol2  int
	OpenID     string // 玩家唯一Id,全服唯一
	FunctionId int    // 功能Id，每个功能Id都是不一样的，且是唯一的标识
}

// GS2CUserStartBattleProto2 == 10 返回数据
type GS2CUserStartBattle struct {
	Protocol  int
	Protocol2 int
}

//------------------------------------------------------------------------------
// C2GSUserChooseRoundProto2 == 5 玩家选择关卡
type C2GSUserChooseRound struct {
	Protocol  int
	Protocol2 int
	MapId     int                    // 章节Id
	RoundId   int                    // 关卡Id
	OpenId    string                 // 玩家唯一Id
	CardList  []*twlib_user.CardInfo // 玩家的card 的属性,战斗前的展位
}

// GS2CUserChooseRoundProto2 == 6 返回玩家数据操作
type GS2CUserChooseRounds struct {
	Protocol  int
	Protocol2 int
	MarkId    string
}

//------------------------------------------------------------------------------
// G2BSUserBattleProto2 == 7 主逻辑服务器发送数据到战斗服
type G2BSUserBattle struct {
	Protocol   int
	Protocol2  int
	BattleType int
	PlayerList []*twlib_user.UserInfo // 玩家结构数据
	RoundId    int
	MarkId     string
}

// GS2CUserBattleProto2 == 8  结算数据
type GS2CUserBattle struct {
	Protocol int

	Protocol2 int
	MarkId    string
	Rewards   []*twlib_rewards.RewardSt
}

//------------------------------------------------------------------------------
