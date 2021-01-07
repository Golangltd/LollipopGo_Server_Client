package Mysyl_DB

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func Init() {
	DB = &mysql_db{}
	SetConfig()
	DB.mysql_open()
	return
}

// 链接数据库ok
func (this *mysql_db) mysql_open() {

	dynamite :=[]string{"ac_game","cf_game","game"}
	decennial :=make(map[string]*sql.DB)
	for i:=0;i<len(dynamite);i++ {
		Odb, err := sql.Open("mysql", Dbusername+":"+Dbpassowrd+"@tcp("+Dbhostsip+")/"+dynamite[i])
		if err != nil {
			fmt.Println("链接失败",dynamite[i])
			continue
		}
		fmt.Println("链接数据库成功...........",dynamite[i])
		// 设置连接池
		Odb.SetMaxOpenConns(dbMaxOpenConns)
		Odb.SetMaxIdleConns(dbMaxIdleConns)
		Odb.Ping()
		// defer Odb.Close()
		decennial[dynamite[i]] = Odb
		// this.STdb = Odb
	}
	this.STdb = decennial
	fmt.Println("链接数据库成功...........",this.STdb)
	// 设置链接池
    /*	this.STdb.SetMaxOpenConns(dbMaxOpenConns)
	this.STdb.SetMaxIdleConns(dbMaxIdleConns)
	this.STdb.Ping()
	defer Odb.Close()*/
}