package tables

import (
	"LollipopGo/log"
	gameDB "LollipopGo2.8x/data"
	twLibDBTable "github.com/Golangltd/Twlib/dbtable"
	"github.com/golang/glog"
	"regexp"
	"strings"
)

// 全局表对应是离散表
type GlobalSt struct {
	Sid   string
	Data1 string // 参数1
	Data2 string // 参数2
	Data3 string // 参数3
	Data4 string // 参数4
	Data5 string // 参数5
	Data6 string // 参数6
	Data7 string // 参数7
}

var GMapGlobalTable map[string]*GlobalSt
var (
	GAttributePowerRate           map[int]int = make(map[int]int, 0) // 离散表220 属性类型，战力数值
	GCampAttributeAddVal          []int                              // 离散表221 阵营加成权重
	GCampAttributeAddRate         int                                // 离散表222 阵营基础属性加成系数(万分比)
	GFirstEnterGameMailTemplateID int                                // 离散表250 玩家角色初始化邮件模板ID
)

// 发送消息获取配置
func AddGlobalTable() {

	GMapGlobalTable = make(map[string]*GlobalSt)
	data := gameDB.GetCFGameData(twLibDBTable.Gl_ls_lssjb)
	if data == nil {
		log.Fatal("error! receive GMapGlobalTable is nil!!")
	}
	// 循环取数据
	for _, v := range data {
		rename := new(GlobalSt)
		rename.Sid = v.(map[string]interface{})["sid"].(string)
		rename.Data1 = v.(map[string]interface{})["data1"].(string)
		rename.Data2 = v.(map[string]interface{})["data2"].(string)
		rename.Data3 = v.(map[string]interface{})["data3"].(string)
		rename.Data4 = v.(map[string]interface{})["data4"].(string)
		rename.Data5 = v.(map[string]interface{})["data5"].(string)
		rename.Data6 = v.(map[string]interface{})["data6"].(string)
		rename.Data7 = v.(map[string]interface{})["data7"].(string)
		GMapGlobalTable[rename.Sid] = rename
	}
	initData()
}

// 初始化数据
func initData() {
	// 离散表220 初始化属性与战斗力转换比例
	initAttributePowerRate(GMapGlobalTable["220"].Data1)
	// 离散表221 阵营加成权重
	GCampAttributeAddVal = StringToIntSlice(GMapGlobalTable["221"].Data1)
	// 离散表222 阵营基础属性加成系数(万分比)
	GCampAttributeAddRate = StrToInt(GMapGlobalTable["222"].Data1)
	// 离散表250 玩家角色初始化邮件模板ID
	GFirstEnterGameMailTemplateID = StrToInt(GMapGlobalTable["250"].Data1)
}

// 初始化属性与战斗力转换比例
func initAttributePowerRate(data string) {
	reg := regexp.MustCompile("\\d+,\\d+")
	strArr := reg.FindAllStringSubmatch(data, -1)
	for _, temp := range strArr {
		arr := strings.Split(temp[0], ",")
		GAttributePowerRate[StrToInt(arr[0])] = StrToInt(arr[1])
	}
	glog.Info("离散220:", GAttributePowerRate)
}
