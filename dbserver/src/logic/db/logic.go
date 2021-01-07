package Mysyl_DB

import (
	"LollipopGo2.8x/conf"
	"database/sql"
)

var (
	Dbhostsip      =  ""
	Dbusername     =  ""
	Dbpassowrd     =  ""
	Dbname         =  ""
	DB             *mysql_db
	dbMaxOpenConns = 2000
	dbMaxIdleConns = 1000
)

type mysql_db struct {
	STdb map[string]*sql.DB
}

func SetConfig()  {
	Dbhostsip      = conf.GetConfig().DB.Host_IP
	Dbusername     = conf.GetConfig().DB.Host_Name
	Dbpassowrd     = conf.GetConfig().DB.Host_PW
	Dbname         = conf.GetConfig().DB.Host_DB
}


func (this *mysql_db) ReadGM_sysAllPlayerInfoDataByIDDali() {//([]*PlayerSTT, bool, int) {

/*	fmt.Println("ReadGM_sysAllPlayerInfoData")
	sql := "SELECT * FROM t_daili"
	fmt.Println(sql)
	bret := true
	rows, err := this.STdb.Query(sql)
	defer rows.Close()
	if err != nil {
		fmt.Println("error:", err)
		bret = false
	} else {
		fmt.Println("没有错误!")
	}

	requests := []*PlayerSTT{}
	for rows.Next() {
		LoginSTSt := new(DaliST)
		rows.Scan(&LoginSTSt.ID, &LoginSTSt.OpenID, &LoginSTSt.CreateTime)
		dada, _, _ := this.ReadGM_sysAllPlayerInfoDataByIDptr(LoginSTSt.OpenID)
		fmt.Println("++++++++++++====dadadadadada", dada)
		fmt.Println(LoginSTSt.ID)
		fmt.Println(LoginSTSt.OpenID)
		requests = append(requests, dada)
	}

	fmt.Println("++++++++++++====", requests)
	return requests, bret, this.Get_Total("t_daili")*/
}
