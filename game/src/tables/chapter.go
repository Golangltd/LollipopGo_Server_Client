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

var (
	ChapterTables map[int]*ChapterSt
	chapterMsg    []interface{}
	chapterOnce   sync.Once
)

// 章节配置
type ChapterSt struct {
	Sid        string
	Name       string // 章节的Id信息，因为文本表
	SubChap    string // 子章节配置
	LastLevel  string // 关卡数据
	FightMap   string
	Award      string
	InstallMap string
	Event      string
	TjAward    string
	AwardShow  string
	Dexc       string
	ItemReward string
}

// 读取章节数据
func GetChapterInfo() {
	data := gameDB.GetCFGameData(twLibDBTable.Gl_gk_zjpzb)
	if data == nil {
		log.Fatal("error! receive skillTables is nil!!")
	}
	chapterMsg = data
	chapterOnce.Do(initChapterTables)
}

func initChapterTables() {
	ChapterTables = make(map[int]*ChapterSt)
	for _, sample := range chapterMsg {
		temp := sample.(map[string]interface{})
		m := &ChapterSt{}
		m.Sid = temp["sid"].(string)
		m.Name = temp["chap_name"].(string)
		m.SubChap = temp["sub_chap"].(string)
		m.LastLevel = temp["last_level"].(string)
		m.FightMap = temp["fight_map"].(string)
		m.Award = temp["award"].(string)
		m.InstallMap = temp["tj_award"].(string)
		m.Event = temp["events"].(string)
		m.AwardShow = temp["award_show"].(string)
		m.Dexc = temp["dexc"].(string)
		m.ItemReward = temp["item_reward"].(string)
		sid, _ := strconv.Atoi(m.Sid)
		ChapterTables[sid] = m
	}
}

// 获取奖励
func GetChapterFor3Row(data string) map[int]int {
	retData := make(map[int]int)
	re := regexp.MustCompile(`\d+,\d+,\d+`)
	strArr := re.FindAllStringSubmatch(data, -1)
	for i := 0; i < len(strArr); i++ {
		arr := strings.Split(strArr[i][0], ",")
		ikey, _ := strconv.Atoi(arr[1])
		val, _ := strconv.Atoi(arr[2])
		retData[ikey] = val
	}
	return retData
}

// 处理处理 获取管卡数
func GetRoundNumFromChapter(sid int) int {
	LastLv := ChapterTables[sid].LastLevel
	if ChapterTables[sid+1] != nil {
		strlastlevnext := ChapterTables[sid+1].LastLevel
		ilastlev, _ := strconv.Atoi(LastLv)
		ilastlevnext, _ := strconv.Atoi(strlastlevnext)
		return ilastlevnext - ilastlev
	} else {
		// 通关
		return -1
	}
}

// 获取子章节的数量
func GetChapter2NumFromChapter(sid int) int {
	tempStr := ChapterTables[sid].SubChap
	subChap, _ := strconv.Atoi(tempStr)
	return subChap
}

// 处理处理 获取管卡数
func GetRoundNumFromChaptersub(sid int) int {
	LastLv := ChapterTables[sid].LastLevel
	ilastlevnext, _ := strconv.Atoi(LastLv)
	return ilastlevnext
}
