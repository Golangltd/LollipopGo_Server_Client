package conf

import (
	"LollipopGo/tools"
	"LollipopGo/tools/fs"
	"LollipopGo2.8x/conf/internal"
	"math/rand"
	"path/filepath"
	"time"
)

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
	*internal.ServerConfig
}

// 服务配置
func ServerConfig() *internal.ServerConfig {
	config := parsers["server"].GetConfig()
	return config.(*internal.ServerConfig)
}

// 初始化服务配置
func InitConfig() {
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
}

// 获取所有配置
func GetConfig() *allConfig {
	return &allConfig{
		ServerConfig(),
	}
}
