package tables

import (
	"LollipopGo/log"
	"regexp"
	"strconv"
	"strings"
)

//string 转 int
func StrToInt(str string) int {
	str = strings.Replace(str, " ", "", -1)
	number, _ := strconv.Atoi(str)
	return number
}

// string 转 []int, string格式 {1,2,3,4,5}
func StringToIntSlice(str string) []int {
	str = strings.Replace(str, " ", "", -1)
	str = str[1 : len(str)-1]
	strArray := strings.Split(str, ",")
	slice := make([]int, 0)
	for i := 0; i < len(strArray); i++ {
		slice = append(slice, StrToInt(strArray[i]))
	}
	return slice
}

//将字符串转换为map
func StrToMap(str string) (result map[int]int) {
	result = make(map[int]int)
	re := regexp.MustCompile(`\d+,\d+`)
	strArr := re.FindAllStringSubmatch(str, -1)
	for _, temp := range strArr {
		list := getKeyAndValue(temp[0])
		if len(list) <= 1 {
			log.Error("An attribute is misConfigured!,attribute is:[v]", temp[0])
			continue
		}
		if _, ok := result[list[0]]; !ok {
			result[list[0]] = list[1]
		}
	}
	return
}

//获取key value
func getKeyAndValue(str string) []int {
	resultList := make([]int, 0)
	if len(str) <= 1 {
		resultList = append(resultList, StrToInt(str))
		return resultList
	}
	strList := strings.Split(str, ",")
	for _, v := range strList {
		resultList = append(resultList, StrToInt(v))
	}
	return resultList
}

//获取正则匹配后的二维数组
func GetRegexpArr(str string, regexpStr string) [][]string {
	re := regexp.MustCompile(regexpStr)
	strArr := re.FindAllStringSubmatch(str, -1)
	return strArr
}

func InitAttribute(str string, regexpStr string) (result map[int]int) {
	result = make(map[int]int)
	strArr := GetRegexpArr(str, regexpStr)
	for _, temp := range strArr {
		list := getKeyAndValue(temp[0])
		if len(list) <= 1 {
			log.Error("An attribute is misConfigured!,attribute is:[v]", temp[0])
			continue
		}
		if _, ok := result[list[0]]; !ok {
			result[list[0]] = list[1]
		}
	}
	return
}

//通过lua配置文件生成的结构体对象中的string类型带有括号与逗号应该去掉
func LuaStrToInt(str string) int {
	re := regexp.MustCompile(`\d+`)
	newStr := string(re.Find([]byte(str)))
	return StrToInt(newStr)
}

//获取cron表达式
func GetCronStr(str string) string {
	return str[1 : len(str)-2]
}
