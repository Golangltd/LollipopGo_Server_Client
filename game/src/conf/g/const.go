package g

import "time"

/*g模块声明用于全局访问的常量*/

const (
	ConfigDir           = "_config"
	ReloadInterval      = 60 * time.Minute //重置的时间间隔
	ReloadCycle         = 30 * time.Minute //热更时间
	ReloadShopConfigLua = 5 * time.Second  //热更时间
)

type TransportType string //传输类型

const (
	Http  TransportType = "http"
	Https TransportType = "https"
	Ws    TransportType = "ws"
)
