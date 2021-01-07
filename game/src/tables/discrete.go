package tables

import (
	. "LollipopGo2.8x/conf/g"
	. "LollipopGo2.8x/lua/go"
	"strings"
)

//考试升级配置数据
type ExamineUpgrade struct {
	ItemType int
	ItemID   int
}

/*离散表*/
var (
	CurrencyList []int //显示货币的List
	EUpgrade     *ExamineUpgrade
)

//初始化进入抽卡界面显示货币列表数据
func InitCurrencyList(strList []string) {
	CurrencyList = make([]int, 0)
	for _, currencyStr := range strList {
		CurrencyList = append(CurrencyList, LuaStrToInt(currencyStr))
	}
}

func initExamineUpgrade(str string) {
	EUpgrade = &ExamineUpgrade{}
	strArr := strings.Split(str[1:len(str)-2], ",")
	EUpgrade.ItemType = StrToInt(strArr[0])
	EUpgrade.ItemID = StrToInt(strArr[1])
}

func init() {
	sampleData := Gvariable_proto[DisPlayCurrencyConfID]
	InitCurrencyList([]string{sampleData.Data1, sampleData.Data2, sampleData.Data3, sampleData.Data4})

	initExamineUpgrade(Gvariable_proto[ExamineConfID].Data1)
}
