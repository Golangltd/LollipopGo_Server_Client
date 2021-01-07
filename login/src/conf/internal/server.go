package internal

import (
	"LollipopGo/leaf"
	"LollipopGo/log"
	"LollipopGo/tools/fs"
	"github.com/BurntSushi/toml"
)

type roomConfig struct {
	ID       int32
	AI       bool
	BotIndex int32
	BotLimit [2]int32
}

type serverConf struct {
	RPCAddr           string
	HTTPAddr          string
	DBAddr            string
	URL               string
	ProfilePort       int
	RoomConfig        []roomConfig
	CreditRecordLimit int32
	HotReload         bool
	AesKey            string
}

type dbConf struct {
	Host    string
	IsDebug bool `toml:"Debug"`
}

type redisConf struct {
	Host     string
	Password string
	Index    int
}

type logConf struct {
	Path  string
	Level string
}

type mongoConf struct {
	Host string
}

type ServerConfig struct {
	Server  *serverConf
	DB      *dbConf
	Mongo   *mongoConf
	Log     *logConf
	Redis   *redisConf
	IMRedis *redisConf
}

type ServerConfigParser struct {
	fs.ParserMixIn
	config *ServerConfig
}

func (scp *ServerConfigParser) ReloadConfig(path string, init bool) bool {
	var sc ServerConfig
	modified, lastTs := scp.CheckModify(path)
	if !modified {
		return false
	}
	if _, err := toml.DecodeFile(path, &sc); err != nil {
		if init {
			panic(err)
		} else {
			log.Error("can't decode %v", path)
			return false
		}
	}
	scp.Lock()
	if init {
		leaf.EnableProfile(sc.Server.ProfilePort)
		if sc.Log.Level == "debug" {
			leaf.ConfigLog(true)
		} else {
			leaf.ConfigLog(false)
		}
		scp.config = &sc
	} else {
		//仅允许部分热加载
		scp.config.Log.Level = sc.Log.Level
		scp.config.Server.RoomConfig = sc.Server.RoomConfig
		scp.config.Server.CreditRecordLimit = sc.Server.CreditRecordLimit
	}
	scp.SetLastModifyTime(path, lastTs)
	scp.Unlock()
	return true
}

func (scp *ServerConfigParser) GetConfig() interface{} {
	scp.RLock()
	defer scp.RUnlock()
	return scp.config
}
