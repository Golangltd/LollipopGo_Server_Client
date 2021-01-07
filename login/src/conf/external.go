package conf

import (
	"LollipopGo/tools"
	"LollipopGo/tools/fs"
	"LollipopGo/tools/tz"
	"login/conf/g"
	"login/conf/internal"
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
	fullPathParsers = make(map[string]fs.IConfigParser, len(parsers))
	configBasePath = tools.GetConfigDir()
)

type allConfig struct {
	*internal.ServerConfig
}

func ServerConfig() *internal.ServerConfig {
	config := parsers["server"].GetConfig()
	return config.(*internal.ServerConfig)
}

func InitConfig() {
	var path string
	rand.Seed(time.Now().Unix())
	for label, parser := range parsers {
		path = filepath.Join(configBasePath, label+".toml")
		parser.ReloadConfig(path, true)
		fullPathParsers[path] = parser
	}
	if ServerConfig().Server.HotReload {
		tz.Schedule(func() {
			fs.WatchConfigFiles(fullPathParsers)
		}, 30, g.GameCloseChan)
	}
}

func GetConfig() *allConfig {
	return &allConfig{
		ServerConfig(),
	}
}