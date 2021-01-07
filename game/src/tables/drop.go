package tables

import (
	"LollipopGo/log"
	gameDB "LollipopGo2.8x/data"
	twLibDBTable "github.com/Golangltd/Twlib/dbtable"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

//掉落表
var (
	DropTable map[int]*DropSt
	dropMsg   []interface{}
	dropOnce  sync.Once
)

// 活动(Drop)配置
type DropSt struct {
	Sid   string
	Coin  string
	Item1 string
	Item2 string
	Item3 string
	Item4 string
	Item5 string
	Item6 string
	Item7 string
	Item8 string
}

// 读取章节数据
func AddDropTable() {
	data := gameDB.GetCFGameData(twLibDBTable.Gl_dl_dlpzb)
	if data == nil {
		log.Fatal("error! receive DropTable is nil!!")
	}
	dropMsg = data
	dropOnce.Do(initDropTable)
}

func initDropTable() {
	DropTable = make(map[int]*DropSt)
	for _, sample := range dropMsg {
		temp := sample.(map[string]interface{})
		m := &DropSt{}
		m.Sid = temp["sid"].(string)
		m.Coin = temp["coin"].(string)
		m.Item1 = temp["item1"].(string)
		m.Item2 = temp["item2"].(string)
		m.Item3 = temp["item3"].(string)
		m.Item4 = temp["item4"].(string)
		m.Item5 = temp["item5"].(string)
		m.Item6 = temp["item6"].(string)
		m.Item7 = temp["item7"].(string)
		m.Item8 = temp["item8"].(string)
		sid, _ := strconv.Atoi(m.Sid)
		DropTable[sid] = m
	}
}

type ItemtypeSt struct {
	ItemId   int
	ItemType int
	ItemNum  int
}

// 类型 道具 数量
func GetSureItemInfo(data string) map[string]*ItemtypeSt {
	retData := make(map[string]*ItemtypeSt)
	re := regexp.MustCompile(`\d+,\d+,\d+,\d`)
	strArr := re.FindAllStringSubmatch(data, -1)
	for i := 0; i < len(strArr); i++ {
		data:=new(ItemtypeSt)

		arr := strings.Split(strArr[i][0], ",")
		data.ItemType, _ = strconv.Atoi(arr[0])
		data.ItemId, _ = strconv.Atoi(arr[1])
		data.ItemNum, _ = strconv.Atoi(arr[2])
		retData[arr[0]+"|"+arr[1]] = data
	}
	return retData
}
