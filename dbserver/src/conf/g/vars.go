package g

import (
	"github.com/globalsign/mgo"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	"sync"
)

var (
	GameCloseChan = make(chan bool)
	GameCloseWG   = &sync.WaitGroup{}
	RoomCloseWG   = &sync.WaitGroup{}
	AllRoom       = &sync.Map{}
	AllAgents     = &sync.Map{}
	ServerClosed  int32
	IMPool          *redis.Pool
	ORM             *gorm.DB
	RedisPool       *redis.Pool
	MG              *mgo.Session
)
