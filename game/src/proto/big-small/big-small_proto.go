package big_small


// 协议号 1000-1050
const (
	C2SEntryGameProto2 = 1000  // 玩家进入游戏 C2SEntryGameProto2
	S2CEntryGameProto2 = 1001 //  S2CEntryGameProto2

	C2SGetHistoryDataProto2 = 1002 // 客户端请求获取历史数据 C2SGetHistoryDataProto2  --游戏下方的
	S2CGetHistoryDataProto2 = 1003 // 返回请求历史数据 S2CGetHistoryDataProto2

	C2SUserChipInProto2 = 1004 // 客户端请求下注 C2SChipInProto2
	S2CUserChipInProto2 = 1005 // 返回下注信息 S2CUserChipInProto2  --- 广播协议

	C2SUserGetItemRateProto2 = 1006 // 客户端请求item概率协议 C2SUserGetItemRateProto2
	S2CUserGetItemRateProto2 = 1007 // 返回item概率协议信息  S2CUserGetItemRateProto2

	C2SUserGetGameStateProto2 = 1008  // 客户端获取游戏状态协议 C2SUserGetGameStateProto2
	S2CUserGetGameStateProto2 = 1009  //                   S2CUserGetGameStateProto2

)

// 商家结构
type B_SUser struct {
	UUID       string
	RoleName   string
	RoleLev    string
	RoleAvatar  int
	RoleSex     int
	RoleCoin    int64
}

// 历史数据
type HistoryData struct {
	HighLow    int  // -1:低(小)；0：高(大)
	FirstData  int  // 1-6 数值，对应骰子的数值
	SecondData int  // 1-6 数值，对应骰子的数值
	ThirdData  int  // 1-6 数值，对应骰子的数值
}

// 排行信息
type RankData struct {
	 Name    string
	 Coin    int
}

// 总排行结构
type RankList struct {
	Richest     []*RankData     // 富豪榜  0位置：第一名，1位置：第二名，2位置：第三名
	BigWinners  []*RankData     // 神算子  0位置：第一名，1位置：第二名，2位置：第三名
}

// 下注类型
const (
	ChipInTypeInit = iota
	Refund  // 退回下注，上一次下注     Refund == 1
	Cancel  // 取消下注，所有的下注     Cancel == 2
	Double  // 双倍下注，上一次下注     Double == 3
	Before  // 上一局压住             Before == 4
)

//----------------------------------------------------------------------------------------------------------------------

// 结算结构
type SettleST struct {
	AreaID int // 1-36
	Times  int // 倍率
	Coin   int // 总钱(单注*倍率)
}

// 红利的结构
type BonusST struct {
	AreaID int // 1-36
	Times  int // 倍率
}

// 游戏状态结构体
const (
	GameStateInit = iota
	Prepare     // 准备阶段
	ChipIn      // 下注阶段
	ChipInOver  // 下注结束
	Settle      // 结算状态
	Wait        // 等待状态
)

//----------------------------------------------------------------------------------------------------------------------

//----------------------------------------------------------------------------------------------------------------------
// 客户端获取游戏状态协议 C2SUserGetGameStateProto2 == 9
type C2SUserGetGameState struct {
	Protocol   int
	Protocol2  int
	RoomType   string  // 房间类型
}

// S2CUserGetGameStateProto2 == 10
type S2CUserGetGameState struct {
	Protocol   int
	Protocol2  int
	GameState  int          // 游戏状态
	Data      *interface{}  // 数据结构
}

//----------------------------------------------------------------------------------------------------------------------
// 获取概率协议：C2SUserGetItemRateProto2
type C2SUserGetItemRate struct {
	Protocol   int
	Protocol2  int
	OpenId     string
	RoomType   string  // 房间类型
}

// S2CUserGetItemRateProto2 服务器返回
type S2CUserGetItemRate struct {
	Protocol   int
	Protocol2  int
	ItemRate   []int  // 服务器发整数，客户端自己加%
}
//----------------------------------------------------------------------------------------------------------------------
// 下注信息
type C2SUserChipIn struct {
	Protocol   int
	Protocol2  int
	OpenId     string
	RoomType   string  // 房间类型
	CoinType   int     // 压住类型 直接发筹码上的数据
	ChipInType int     // 1:退回下注,2:取消下注,3:双倍下注
	ChipInPos  int     // 下注位置
	ChipInCoin int     // 下注金额
}

// 广播协议  -- 可以不广播
type S2CUserChipIn struct{
	Protocol   int
	Protocol2  int
	IsMy       bool  // 是否是自己下注
	ChipInPos  int   // 下注位置
	ChipInCoin int   // 下注金额
	BeforeCoin []int  // 上一局所有筹码
}
//----------------------------------------------------------------------------------------------------------------------
// 请求历史数据
type C2SGetHistoryData struct {
	Protocol  int
	Protocol2 int
	OpenId     string
	RoomType   string  // 房间类型
}

// 返回历史数据：游戏底部历史数据
type S2CGetHistoryData struct {
	Protocol  int
	Protocol2 int
	HistoryData []*HistoryData
}

//----------------------------------------------------------------------------------------------------------------------
// 玩家进入游戏
type C2SEntryGame struct {
	Protocol  int
	Protocol2 int
	Tocken    string   // 玩家的校验码
	OpenId    string   // 用户唯一Id
	RoomType  int      // 房间类型：  1 2 3 4
}

// 返回进入游戏的状态信息
type S2CEntryGame struct {
	Protocol   int
	Protocol2  int
	UserSt     *B_SUser     // 玩家的结构信息
    Data       *EntryGame
}

// 返回进入游戏的状态信息
type EntryGame struct {
	UserSt         []*B_SUser   // 玩家的结构信息
	RankList       *RankList    // 富豪版+神算子
	PlayerOl       int          // 在线人数
	RoomType       string       // 房间类型
	GameState      int          // 游戏状态  --> 对应不同的状态有不同的结构体
	RoundID        string       // 局的ID
	Data           interface{}  // 对应游戏状态的数据结构，例如游戏准备状态 --> 对应准备的结构数据
	CoinTypeList   []int        // 筹码数据
}

//----------------------------------------------------------------------------------------------------------------------
// 游戏准备的结构
type PrepareST struct {
	NextStateTime int   // 毫秒
}

// 下注阶段的结构体
type ChipInST struct {
	NextStateTime int   // 毫秒
	AreaData []int      // 每个区域的钱
}

// 下注结束
type ChipInOverST struct {
	NextStateTime int        // 毫秒
	AreaData      []int      // 每个区域的钱
	MyAreaData    []int      // 玩家压的每个区域的钱
}

// 结算结构
type SettleStateST struct {
	NextStateTime int      // 毫秒
	BonusData  []*BonusST  // 红利结构
	SettleData []*SettleST // 结算结构
}

// 等待状态
type WaitST struct {
	NextStateTime int   // 毫秒
}
//----------------------------------------------------------------------------------------------------------------------