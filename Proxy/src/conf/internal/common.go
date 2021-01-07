package internal

import (
	"LollipopGo/log"
	"LollipopGo/tools/fs"
	"github.com/BurntSushi/toml"
)

type roomConf struct {
	ID         int32
	Name       string
	Base       int32
	Chips      []int32
	EnterLimit int32
	Status     int32
}

type CommonConfig struct {
	RoomBase []roomConf
}

type CommonConfigParser struct {
	fs.ParserMixIn
	config *CommonConfig
}

func (ccp *CommonConfigParser) ReloadConfig(path string, init bool) bool {
	modified, lastTs := ccp.CheckModify(path)
	if !modified {
		return false
	}
	var cc CommonConfig
	if _, err := toml.DecodeFile(path, &cc); err != nil {
		if init {
			panic(err)
		} else {
			log.Error("fail to reload %v for %v", path, err)
			return false
		}
	}
	ccp.Lock()
	ccp.config = &cc
	ccp.SetLastModifyTime(path, lastTs)
	ccp.Unlock()
	return true
}

func (ccp *CommonConfigParser) GetConfig() interface{} {
	ccp.RLock()
	defer ccp.RUnlock()
	return ccp.config
}
