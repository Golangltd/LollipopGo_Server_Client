package game_util

import "golang.org/x/net/websocket"

var ConnAll *websocket.Conn

func SetConnALl(connall *websocket.Conn)  {
	if connall!=nil{
		ConnAll = connall
	}
}

