package conf

import (
	"LollipopGo/tools"
	"LollipopGo/tools/fs"
	"LollipopGo/tools/tz"
	"LollipopGo2.8x/conf/g"
	"LollipopGo2.8x/conf/internal"
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

//	sc := ServerConfig()
//	_, _, dc, _ := sc.Redis, sc.IMRedis, sc.DB, sc.Mongo
/*	sc := ServerConfig()
	rc, ic, dc, mc := sc.Redis, sc.IMRedis, sc.DB, sc.Mongo
	g.RedisPool = database.NewRedisPool(rc.Host, rc.Password, rc.Index)
	g.IMPool = database.NewRedisPool(ic.Host, ic.Password, ic.Index)
	g.ORM = database.NewMysqlConn(dc.Host, dc.IsDebug)
	g.MG = database.NewMongoSession(mc.Host)*/
//	g.ORM = database.NewMysqlConn(dc.Host, dc.IsDebug)

	// 配置文件热加载
	if ServerConfig().Server.HotReload {
		tz.Schedule(func() {
			fs.WatchConfigFiles(fullPathParsers)
		}, g.ReloadCycle, g.GameCloseChan)
	}
}

// 获取所有配置
func GetConfig() *allConfig {
	return &allConfig{
		ServerConfig(),
	}
}
