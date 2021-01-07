package lua_uitl

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func GetLuaDataFor2Row(data string) map[int]int {
	retdata := make(map[int]int)
	re := regexp.MustCompile(`\d+,\d+`)
	strArr := re.FindAllStringSubmatch(data, -1)
	for i := 0; i < len(strArr); i++ {
		arr := strings.Split(strArr[i][0], ",")
		ikey, _ := strconv.Atoi(arr[0])
		ival, _ := strconv.Atoi(arr[1])
		retdata[ikey] = ival
	}
	return retdata
}

type DataSt struct {
	Id   int
	Type int
	Num  int64
}

func GetLuaDataFor3Row(data string) map[string]*DataSt {
	retData := make(map[string]*DataSt)
	re := regexp.MustCompile(`\d+,\d+,\d+`)
	strArr := re.FindAllStringSubmatch(data, -1)
	for i := 0; i < len(strArr); i++ {
		data := new(DataSt)
		arr := strings.Split(strArr[i][0], ",")
		itype, _ := strconv.Atoi(arr[0])
		ikey, _ := strconv.Atoi(arr[1])
		val, _ := strconv.Atoi(arr[2])
		data.Type = itype
		data.Id = ikey
		data.Num = int64(val)
		retData[arr[1]+"|"+arr[2]] = data
	}
	fmt.Println("------------",retData)
	return retData
}
