package modules

import (
	"errors"
	"fmt"
	"github.com/golang/glog"
	"golang.org/x/net/websocket"
	"sync"
)

type HandlersMap struct {
	sync.RWMutex //加入锁来控制写入
	Handlers     map[float64]func(conn *websocket.Conn, reqData map[string]interface{})
}

var HM *HandlersMap

func init() {
	HM = &HandlersMap{}
	HM.Handlers = make(map[float64]func(conn *websocket.Conn, reqData map[string]interface{}))
}

func (m *HandlersMap) AddHandler(protoType int, handler func(conn *websocket.Conn, reqData map[string]interface{})) {
	HM.RLock()
	defer HM.RUnlock()
	if _, ok := HM.Handlers[float64(protoType)]; ok { // 如果存在key的handler,就返回错误！
		err := errors.New(fmt.Sprintf("this protoType is [%v], but handler already exists!", protoType))
		glog.Fatal(err.Error())
	}
	HM.Handlers[float64(protoType)] = handler
	return
}

//是否存在此协议的handler
func (m *HandlersMap) IsExistHandler(protoType float64) (isExist bool) {
	_, isExist = m.Handlers[protoType]
	return
}
