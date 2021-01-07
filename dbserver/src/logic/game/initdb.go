package game

import (
	Mysyl_DB "LollipopGo2.8x/logic/db"
	"database/sql"
)

type GameRPC struct {
}

func GetGameConn() *sql.DB {
	return Mysyl_DB.DB.STdb["game"]
}

