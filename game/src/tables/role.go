package tables

import (
	"LollipopGo/log"
	gameDB "LollipopGo2.8x/data"
	twLibDBTable "github.com/Golangltd/Twlib/dbtable"
	"sync"
)

var (
	roleOnce   sync.Once
	roleMsg    []interface{}
	RolesTable map[int]*RoleSample //角色表
)

//角色样本
type RoleSample struct {
	RoleID         int         //角色表
	Sex            int         //性别
	Name           string      //名称
	Desc           string      //描述
	Vocation       int         //职业
	Camp           int         //阵营
	Quality        int         //品质
	MaxQuality     int         //最高品质
	Talent         string      //天赋
	Skills         map[int]int //技能
	Resource       string      //美术资源
	Attributes     map[int]int //基础属性
	LevelGrowth    string      //等级成长
	StarGrowth     string      //星级成长
	QualityGrowth  string      //品质成长
	HandBookHeroID int         //图鉴ID
}

// 获取数据库配置数据
func AddRoleTable() {
	data := gameDB.GetCFGameData(twLibDBTable.RoleTable)
	if data == nil {
		log.Fatal("error! receive roleTables is nil!!")
	}
	roleMsg = data
	roleOnce.Do(initRoleTable)
}

//初始化技能
func (m *RoleSample) initSkills(skills ...int) {
	m.Skills = make(map[int]int)
	for id, skill := range skills {
		m.Skills[id] = skill
	}
}

func initRoleTable() {
	RolesTable = make(map[int]*RoleSample)
	for _, sample := range roleMsg {
		temp := sample.(map[string]interface{}) //断言
		m := &RoleSample{}
		m.RoleID = StrToInt(temp["sid"].(string))
		m.Name = temp["name"].(string)
		m.Desc = temp["description"].(string)
		m.Vocation = StrToInt(temp["vocation"].(string))
		m.Camp = StrToInt(temp["camp"].(string))
		m.Quality = StrToInt(temp["quality"].(string))
		m.MaxQuality = StrToInt(temp["maxquality"].(string))
		m.Talent = temp["talent"].(string)
		m.initSkills(StrToInt(temp["skill1"].(string)), StrToInt(temp["skill2"].(string)), StrToInt(temp["skill3"].(string)),
			StrToInt(temp["skill4"].(string)), StrToInt(temp["skill5"].(string)))
		//m.Resource = temp["resouce"].(string)
		m.Attributes = StrToMap(temp["attribute1"].(string))
		m.LevelGrowth = temp["attribute2"].(string)
		m.StarGrowth = temp["attribute3"].(string)
		m.QualityGrowth = temp["attribute4"].(string)
		m.Sex = StrToInt(temp["gender"].(string))
		m.HandBookHeroID = StrToInt(temp["handbookhero_id"].(string))
		//把每个角色对象添加到map中
		RolesTable[m.RoleID] = m
	}
	roleMsg = nil //删除
}
