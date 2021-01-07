package conf

import (
	"LollipopGo/tools"
	"LollipopGo/tools/fs"
	"LollipopGo2.8x/conf/internal"
	"github.com/golang/glog"
	"math/rand"
	"path/filepath"
	"time"
)

/*
   整个的服务器的配置的文件的管理，
   1. 读取服务器的配置，加载到内存
   2. 底层支持“热更”， 实现配置修改，服务器定期更新内存（启动定时器实现）
   3. 通用模板的实现，通过结构体实现“类”的性质
   4. 读取配置的第三方库：https://github.com/BurntSushi/toml
*/

var (
	// 服务器配置文件
	parsers = map[string]fs.IConfigParser{
		"server": &internal.ServerConfigParser{},
	}

	// 所有配置文件内存表现
	fullPathParsers = make(map[string]fs.IConfigParser, len(parsers))
	// 配置文件路径
	configBasePath = tools.GetConfigDir()
)

type allConfig struct {
	*internal.ServerConfigParser
}

// 服务配置
func ServerConfig() *internal.ServerConfigParser {
	config := parsers["server"].GetConfig()
	return config.(*internal.ServerConfigParser)
}

// 初始化服务配置
func InitConfig() {
	glog.Info("InitConfig")
	var path string
	// 设置随机种子
	rand.Seed(time.Now().Unix())
	// 加载各个游戏配置
	for label, parser := range parsers {
		// 配置文件路径
		path = filepath.Join(configBasePath, label+".toml")
		// 加载配置文件
		parser.ReloadConfig(path, true)
		fullPathParsers[path] = parser
	}

	// 配置文件热加载
	//if ServerConfig().Server.HotReload {
	//	tz.Schedule(func() {
	//		fs.WatchConfigFiles(fullPathParsers)
	//	}, g.ReloadCycle, g.GameCloseChan)
	//}
}

// 获取所有配置
func GetConfig() *allConfig {
	return &allConfig{
		ServerConfig(),
	}
}
