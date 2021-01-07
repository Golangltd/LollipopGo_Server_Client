package main

import (
	"LollipopGo/network"
	"LollipopGo/tools/sample"
	"LollipopGo2.8x/conf"
	"LollipopGo2.8x/cxt"
	Data "LollipopGo2.8x/data"
	_ "LollipopGo2.8x/handlers"
	_ "LollipopGo2.8x/time/ticker"
	"github.com/Golangltd/cache2go"
	"github.com/golang/glog"
	"github.com/nsqio/go-nsq"
	"golang.org/x/net/websocket"
	"net/http"
	"runtime"
	"time"
)

var producer *nsq.Producer

func main() {
	// 初始化本服唯一ID生成器
	glog.Infof("appHostID:%d\n", conf.GetConfig().GetAppHostID())
	// 初始化db rpc
	Data.DBInit()
	InitCache()
	// go LoadTableConfig()
	// go OffLine.InitOffLine()
	sample.InitRand()
	runtime.GOMAXPROCS(runtime.NumCPU())
	http.Handle("/"+conf.GetConfig().GetServerUrl(), websocket.Handler(BuildConnection))
	glog.Infof("game listen to:[%s]\n", conf.GetConfig().GetServerWsAddr())
	glog.Info("game start ok")
	if err := http.ListenAndServe(conf.GetConfig().GetServerWsAddr(), nil); err != nil {
		glog.Info("Entry nil", err.Error())
		glog.Flush()
		return
	}
}

func BuildConnection(ws *websocket.Conn) {
	data := ws.Request().URL.Query().Get("_config")
	if data == "" {
		glog.Info("_config is Nil")
		glog.Flush()
		return
	}
	impl.InitConnection(ws)
}

// 加载配置
func LoadTableConfig() {
	time.AfterFunc(time.Second*5, AddTableConfig)
}

// 初始化 Cache
func InitCache() {
	cxt.CacheGame = cache2go.Cache("TW.Enjoy")
}

// 获取配置
func AddTableConfig() {
	//config.GetFuncBeginInfo()// 获取功能开启配置表
	//tables.AddGlobalTable() // 获取离散表
	//tables.AddSkillTable()  // 获取技能总表
	//tables.AddRoleTable()   // 获取角色表
	//tables.GetMingGCInfo()  // 获取敏感词表
	//tables.GetChapterInfo() // 获取章节表
	//tables.AddRoundTable()  // 获取关卡表
	//tables.AddCardTJTable() // 获取图鉴表
	//tables.AddItemTable()       // 获取道具表
	//tables.AddDropTable()       // 获取掉落表
	//tables.AddCollegeTable()    // 获取学院表
	//tables.AddSchoolRollTable() // 获取学籍表
	//tables.AddSRTrialTable()    //获取学籍试炼表
	//tables.LoadEquipConf()      // 获取装备表
	//tables.GetWuShiInfo() // 巫师经验表
	//tables.AddSRTrialTable()    // 获取学籍试炼表
}

