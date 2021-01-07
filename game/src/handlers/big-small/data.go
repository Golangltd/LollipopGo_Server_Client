package bigsmall

import (
	impl "LollipopGo/network"
	"LollipopGo/util"
	bacall "LollipopGo2.8x/proto/big-small"
	"golang.org/x/net/websocket"
	"strconv"
)

//----------------------------------------------------------------------------------------------------------------------

// 房间结构
type RoomData struct {
	RoomID    string               // 房间ID
	Data      *bacall.EntryGame    // 结构
}

// 玩家结构
type PlayerST struct {
	MyAreaData []*ChipInST  // 索引是位置
}

// 押注结构
type ChipInST struct {
   MapData map[int]*SaveData  // key:金币类型
}

// 押注金币的数量结构
type SaveData struct {
	ICoinNum int       // 金币的数量
}

// 筹码
var (
	Lo = []int{10,20,30,50,100,200}
	Cj = []int{20,30,50,100,200,300}
	ZJ = []int{100,200,300,500,1000,2000}
	GJ = []int{200,300,5000,1000,2000,3000}
)

var (
	GRoom map[int]*RoomData
	GPlayer map[string]*PlayerST
	GPlayerBefore map[string]*PlayerST
	GPlayerBeforeOne map[string]int
	GPlayerBeforePos map[string]int
	GHistory map[int][]*bacall.HistoryData
	GPlayerOpenID map[string]int
	ConnBS *websocket.Conn
)

//----------------------------------------------------------------------------------------------------------------------

func init()  {

	GRoom = make(map[int]*RoomData)
	GPlayer = make(map[string]*PlayerST)
	GPlayerBefore = make(map[string]*PlayerST)
	GPlayerBeforeOne = make(map[string]int)
	GPlayerBeforePos = make(map[string]int)
	GHistory = make(map[int][]*bacall.HistoryData)
	GPlayerOpenID = make(map[string]int)

	for i:=0;i<4;i++{
		rcdata := new(RoomData)
		rcdata.RoomID = util.MD5_LollipopGO(strconv.Itoa(i+1))
		rcdataSt := new(bacall.EntryGame)
		rcdataSt.RoomType = rcdata.RoomID
		rcdataSt.GameState = bacall.Prepare
		rcdataSt.RoundID = "1"  // 唯一的数据
		rcdataSt.PlayerOl = 0
		rcdataSt.UserSt = nil
		// 游戏状态
		stalest :=new(bacall.PrepareST)
		stalest.NextStateTime = 1
		rcdataSt.Data = stalest
		// 排行榜
		stRank :=new(bacall.RankList)
		da:=new(bacall.EntryGame )
		rcdata.Data =da
		rcdata.Data.RankList = stRank
		// 筹码
		rcdataSt.CoinTypeList = GetCoinType(i+1)
		GRoom[i+1] = rcdata
	}
}

// 设置连接
func SetConn(conn *websocket.Conn)  {
	ConnBS = conn
}

// 获取筹码
func GetCoinType(roomType int)  []int{
	var coinData []int
	if roomType == 1{
		coinData = Lo
	}else if roomType == 2 {
		coinData = Cj
	}else if roomType == 3 {
		coinData = ZJ
	}else if roomType == 4 {
		coinData = GJ
	}
   return coinData
}

// 加入房间逻辑
func AddRoomFunc(user *bacall.B_SUser,roomType int)  {
	rcdata :=GRoom[roomType].Data
	rcdata.PlayerOl++
	rcdata.UserSt = append(rcdata.UserSt,user)
	return
}

// 获取房间状态
func GetRoomState(roomType int) (int,interface{}) {
	rcdata :=GRoom[roomType].Data
	return rcdata.GameState,rcdata.Data
}

// 更新房间状态
func UpdateRoomState(roomType int,leftTime int)  {
	rcdata :=GRoom[roomType].Data
	if  rcdata.GameState == bacall.Wait{
		rcdata.GameState = bacall.Prepare
	}else {
		rcdata.GameState++
	}
	UpdateGameStateST(rcdata,leftTime,roomType)
}

// 更新对应状态的结构数据
func UpdateGameStateST(roomData *bacall.EntryGame,leftTime int,roomType int)  {
	gameState:=roomData.GameState
	if gameState == bacall.Prepare{
		roomData.Data =bacall.PrepareST{
			NextStateTime:leftTime,
		}
		// 发消息
		for k,v:=range GPlayerOpenID{
			if v == roomType {
				SendAllPlayer(roomData.Data, k)
			}
		}
	}else if gameState == bacall.ChipIn{
		roomData.Data =bacall.ChipInST{
			NextStateTime:leftTime,
			AreaData:nil,
		}
		// 下注阶段
		for k,v:=range GPlayerOpenID{
			if v == roomType {
				SendAllPlayer(roomData.Data, k)
			}
		}
	}else if gameState == bacall.ChipInOver{
		roomData.Data =bacall.ChipInOverST{
			NextStateTime:leftTime,
			AreaData:nil,
			MyAreaData:nil,
		}
		// 下注结束几段
		for k,v:=range GPlayerOpenID{
			if v == roomType {
				SendAllPlayer(roomData.Data, k)
			}
		}
	}else if gameState == bacall.Settle{
		roomData.Data =bacall.SettleStateST{
			NextStateTime:leftTime,
			BonusData:nil,
			SettleData:nil,
		}
		// 保存开奖记录
		SaveChipInData(bacall.CreatResult(),roomType)
		// 保存上一局
		GPlayerBefore = GPlayer
		// 发消息
		for k,v:=range GPlayerOpenID{
			if v == roomType {
				SendAllPlayer(roomData.Data, k)
			}
		}
	}else if gameState == bacall.ChipIn{
		roomData.Data =bacall.WaitST{
			NextStateTime:leftTime,
		}
		for k,v:=range GPlayerOpenID{
			if v == roomType {
				SendAllPlayer(roomData.Data, k)
			}
		}
	}
}

//概率判断
func GetChipInRate(data []*bacall.HistoryData)  []int{
	var rateall []int
	allcount := len(data)*4
	hi:=0
	lo:=0
	one :=0
	two :=0
	three :=0
	four :=0
	five :=0
	six :=0

	for _,v:=range data{
		// 高低
      if v.HighLow == -1{
		  lo++
	  }else if v.HighLow == 0{
		  hi++
	  }
	  // 点数
	  if v.FirstData == One ||
		  v.SecondData == One ||
		  v.ThirdData == One  {
		  one++
	  }else if v.FirstData == Two ||
		  v.SecondData == Two ||
		  v.ThirdData == Two  {
		  two++
	  }else if v.FirstData == Three||
		  v.SecondData == Three ||
		  v.ThirdData == Three  {
		  three++
	  }else if v.FirstData == Four ||
		  v.SecondData == Four ||
		  v.ThirdData == Four  {
		  four++
	  }else if v.FirstData == Five ||
		  v.SecondData == Five ||
		  v.ThirdData == Five  {
		  five++
	  }else if v.FirstData == Six ||
		  v.SecondData == Six ||
		  v.ThirdData == Six  {
		  six++
	  }
	}
	// 计算概率
	rateall = append(rateall,one*10000/allcount)
	rateall = append(rateall,two*10000/allcount)
	rateall = append(rateall,three*10000/allcount)
	rateall = append(rateall,four*10000/allcount)
	rateall = append(rateall,five*10000/allcount)
	rateall = append(rateall,six*10000/allcount)
	return rateall
}

// 保存历史数据
func SaveChipInData(data []int,roomType int)  {

	heave :=new(bacall.HistoryData)
	val:=0
	for _,v:=range data{
		val+=v
	}

	if val <= 10 {
		heave.HighLow = -1
	} else {
		heave.HighLow = 0
	}

     for k,v:=range data{
     	if k == 0 {
			heave.FirstData = v
		}else if k ==1{
			heave.SecondData = v
		}else if k ==2{
			heave.ThirdData = v
		}
	 }

	 if len(GHistory[roomType]) >=60{
		 GHistory[roomType] = GHistory[roomType][1:]
		 GHistory[roomType] = append(GHistory[roomType], heave)
	 }else {
		 GHistory[roomType] = append(GHistory[roomType], heave)
	 }
}

// 发送
func SendAllPlayer(data interface{},strOpenId string)  {
		impl.PlayerSendToProxyServer(ConnBS, data, strOpenId)
}