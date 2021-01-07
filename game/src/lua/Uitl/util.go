package lua_uitl

import (
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

func GetLuaDataFor3Row(data string) map[int]*DataSt {

	retData := make(map[int]*DataSt)
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
		retData[ikey] = data
	}
	return retData
}

func GetFirstLuaDataFor4Row(data string) []int {

	retData := make([]int, 0)
	re := regexp.MustCompile(`\d+,\d+,\d+,\d+`)
	strArr := re.FindAllStringSubmatch(data, -1)
	if len(strArr) < 1 {
		return retData
	} else {
		arr := strings.Split(strArr[0][0], ",")
		data0, _ := strconv.Atoi(arr[0])
		retData = append(retData, data0)
		data1, _ := strconv.Atoi(arr[1])
		retData = append(retData, data1)
		data2, _ := strconv.Atoi(arr[2])
		retData = append(retData, data2)
		data3, _ := strconv.Atoi(arr[3])
		retData = append(retData, data3)
	}
	return retData
}

func GetFirstLuaDataFor5Row(data string) []int {

	retData := make([]int, 0)
	re := regexp.MustCompile(`\d+,\d+,\d+,\d+,\d+`)
	strArr := re.FindAllStringSubmatch(data, -1)
	if len(strArr) < 1 {
		return retData
	} else {
		arr := strings.Split(strArr[0][0], ",")
		data0, _ := strconv.Atoi(arr[0])
		retData = append(retData, data0)
		data1, _ := strconv.Atoi(arr[1])
		retData = append(retData, data1)
		data2, _ := strconv.Atoi(arr[2])
		retData = append(retData, data2)
		data3, _ := strconv.Atoi(arr[3])
		retData = append(retData, data3)
		data4, _ := strconv.Atoi(arr[4])
		retData = append(retData, data4)
	}
	return retData
}
