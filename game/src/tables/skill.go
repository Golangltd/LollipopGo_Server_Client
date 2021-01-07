package tables

import (
	"LollipopGo/log"
	gameDB "LollipopGo2.8x/data"
	twLibDBTable "github.com/Golangltd/Twlib/dbtable"
	"strings"
	"sync"
)

var (
	skillOnce   sync.Once
	skillMsg    []interface{}
	SkillTables map[int]*SkillSample
)

type AttributeType = int //属性类型

//技能总表每条样本数据
type SkillSample struct {
	ID            int   //id
	CoolDown      int   //冷却时间
	SPExpend      int   //SP消耗
	Type          int   //技能类型
	Attribute     int   //技能属性
	Function      int   //功能
	Multiplier    int   //技能倍率
	DefaultDmg    int   //默认伤害
	Target        int   //有效目标类型
	TriggerID     int   //触发机制表中对应的ID
	EffectID      int   //特殊效果表中对应的ID
	ExtraTargetID int   //额外目标表中对应的ID
	TargetCap     int   //目标上限
	Additional    []int //附加节能id，可能有多个这种技能
}

func AddSkillTable() {

	data := gameDB.GetCFGameData(twLibDBTable.SkillTable4)
	if data == nil {
		log.Fatal("error! receive skillTables is nil!!")
	}
	skillMsg = data
	skillOnce.Do(initSkillTable)
}

func initSkillTable() {
	SkillTables = make(map[AttributeType]*SkillSample)
	for _, sample := range skillMsg {
		temp := sample.(map[string]interface{})
		m := &SkillSample{}
		m.ID = StrToInt(temp["sid"].(string))
		m.CoolDown = StrToInt(temp["cooldown"].(string))
		m.SPExpend = StrToInt(temp["spexpend"].(string))
		m.Type = StrToInt(temp["skill_type"].(string))
		m.Attribute = StrToInt(temp["attribute"].(string))
		m.Function = StrToInt(temp["basic_type"].(string))
		m.Multiplier = StrToInt(temp["dmgparameter1"].(string))
		m.DefaultDmg = StrToInt(temp["dmgparameter2"].(string))
		m.Target = StrToInt(temp["target"].(string))
		m.TriggerID = StrToInt(temp["trigger"].(string))
		m.EffectID = StrToInt(temp["effect"].(string))
		m.ExtraTargetID = StrToInt(temp["extratarget"].(string))
		m.TargetCap = StrToInt(temp["maxtarget"].(string))
		m.Additional = addAdditional(temp["additional"].(string))
		//把每个技能对象存到map中
		SkillTables[m.ID] = m
	}
	skillMsg = nil

}

func addAdditional(str string) []int {
	if len(str) == 0 {
		return nil
	}
	strArr := strings.Split(str, ",")
	additional := make([]int, len(strArr))
	for index, v := range strArr {
		additional[index] = StrToInt(v)
	}
	return additional
}
