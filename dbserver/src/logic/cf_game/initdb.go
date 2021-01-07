package cf_game

import (
	Mysyl_DB "LollipopGo2.8x/logic/db"
	"database/sql"
	"fmt"
)

type CfRPC struct {
}

func GetCfGameConn() *sql.DB {
	dbptr :=Mysyl_DB.DB.STdb["cf_game"]
	if dbptr == nil{
		fmt.Println("GetCfGameConn is nil of cf_game")
		return nil
	}
	return	dbptr
}

