package proto_net

// 主协议 GameNet_Proto == 3       游戏的NET主协议
const (
	INIT_PROTO_NET_INIT = iota //
	Net_HeartBeatProto         //  Net_HeartBeatProto == 1  心跳协议
	Net_RelinkProto            //  Net_RelinkProto == 2  断线重连
)

//------------------------------------------------------------------------------
// 断线重连
type Net_Relink struct {
	Protocol  int
	Protocol2 int
	OpenId    string
	GameState int  // 玩家在服务器的状态
}

//------------------------------------------------------------------------------
// 心跳协议
type Net_HeartBeat struct {
	Protocol  int
	Protocol2 int
	OpenId    string
	TimeStamp int64  // 服务器时间戳
}

//------------------------------------------------------------------------------
