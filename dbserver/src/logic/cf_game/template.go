package cf_game

import (
	"database/sql"
	"fmt"
	twlib_dbtable "github.com/Golangltd/Twlib/dbtable"
	_ "github.com/go-sql-driver/mysql"
)

var GameMapTTL map[string]int

func init()  {
	GameMapTTL = make(map[string]int)
}

func (this *CfRPC)GetCfGameData(tableName string,reply *interface{}) error {
	if len(twlib_dbtable.GetDbTableEnum(tableName)) == 0{
		return nil
	}
	sql:=fmt.Sprintf("select * from "+ twlib_dbtable.GetDbTableEnum(tableName))
	fmt.Println(sql)
	rows, err := GetCfGameConn().Query(sql)
	defer rows.Close()
	if err!=nil{
		fmt.Println(err.Error())
		return err
	}
	cols, err := rows.Columns()
	if err!=nil{
		return err
	}
    fmt.Println("s所有的列：",len(cols))
	vals := make([][]byte, len(cols))
	scans := make([]interface{}, len(cols))
	for k, _ := range vals {
		scans[k] = &vals[k]
	}
	*reply = GetAll(rows)
	return nil
}

type s1 map[string]string
type s2 [] s1
func GetAll(rows *sql.Rows) s2 {
	if rows == nil {
		return nil
	}
	cols, err := rows.Columns()
	rawResult := make([][]byte, len(cols))
	result := s2{};
	dest := make([]interface{}, len(cols))
	for i, _ := range rawResult {
		dest[i] = &rawResult[i]
	}
	//rowsNum:=0;
	for rows.Next()  {
		err = rows.Scan(dest...)
		sresult:=make(s1, len(cols))
		for i, raw := range rawResult {
			if raw == nil {
				sresult[cols[i]] = ""
			} else {
				sresult[cols[i]] = string(raw)
			}
		}
		result=append(result,sresult)
	}
	_=err
	return result
}
