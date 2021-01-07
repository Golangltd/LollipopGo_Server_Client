package internal

import (
	"LollipopGo/log"
	"LollipopGo/tools/fs"
	"github.com/BurntSushi/toml"
)

type serverConf struct {
	WSAddr     string `toml:"WSAddr"`
	AppHostID  int64  `toml:"AppHostID"`
	DBAddr     string `toml:"DBAddr"`
	URL        string `toml:"URL"`
	HotReload  bool   `toml:"HotReload"`
	AesKey     string `toml:"AesKey"`
	AreaListId int    `toml:"AreaListId"`
	LoginHost  string `toml:"LoginHost"`
	LoginPort  string `toml:"LoginPort"`
	LoginPath  string `toml:"LoginPath"`
	ProxyHost  string `toml:"ProxyHost"`
	ProxyPort  string `toml:"ProxyPort"`
	ProxyPath  string `toml:"ProxyPath"`
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
		scp.config = &sc
	} else {
		//仅允许部分热加载
		scp.config.Log.Level = sc.Log.Level
	}
	scp.SetLastModifyTime(path, lastTs)
	scp.Unlock()
	return true
}

func (scp *ServerConfigParser) GetConfig() interface{} {
	scp.RLock()
	defer scp.RUnlock()
	return scp
}

//得到获取login数据的url拼接列表
func (scp *ServerConfigParser) GetLoginUrlList() []string {
	scp.RLock()
	defer scp.RUnlock()
	urlList := make([]string, 3) //这里可以写死，因为获取什么是已知的
	urlList[0] = scp.config.Server.LoginHost
	urlList[1] = scp.config.Server.LoginPort
	urlList[2] = scp.config.Server.LoginPath
	return urlList
}

//得到获取proxy数据的url拼接列表
func (scp *ServerConfigParser) GetProxyUrlList() []string {
	scp.RLock()
	defer scp.RUnlock()
	urlList := make([]string, 3) //这里可以写死，因为获取什么是已知的
	urlList[0] = scp.config.Server.ProxyHost
	urlList[1] = scp.config.Server.ProxyPort
	urlList[2] = scp.config.Server.ProxyPath
	return urlList
}

func (scp *ServerConfigParser) GetAppHostID() int64 {
	scp.RLock()
	defer scp.RUnlock()
	return scp.config.Server.AppHostID
}

func (scp *ServerConfigParser) GetHotReload() bool {
	scp.RLock()
	defer scp.RUnlock()
	return scp.config.Server.HotReload
}

func (scp *ServerConfigParser) GetServerUrl() string {
	scp.RLock()
	defer scp.RUnlock()
	return scp.config.Server.URL
}

func (scp *ServerConfigParser) GetServerWsAddr() string {
	scp.RLock()
	defer scp.RUnlock()
	return scp.config.Server.WSAddr
}

func (scp *ServerConfigParser) GetDBAddr() string {
	scp.RLock()
	defer scp.RUnlock()
	return scp.config.Server.DBAddr
}
