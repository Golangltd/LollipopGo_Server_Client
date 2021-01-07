package tables

import "golang.org/x/net/websocket"

var configConn *websocket.Conn

func SetConfigConn(connMain *websocket.Conn) {
	if connMain != nil {
		configConn = connMain
	}
}
