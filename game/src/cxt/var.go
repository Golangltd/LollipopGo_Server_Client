package cxt

import (
	"github.com/Golangltd/cache2go"
	concurrent "github.com/fanliao/go-concurrentMap"
	"github.com/yuin/gopher-lua"
	"golang.org/x/net/websocket"
)

var (
	CacheGame  *cache2go.CacheTable
	M      *concurrent.ConcurrentMap
	ConnXZ *websocket.Conn
	L      *lua.LState
)

type MapKey = string

const (
	UserKey MapKey = "|User" //通过全局key获取user的信息
)

