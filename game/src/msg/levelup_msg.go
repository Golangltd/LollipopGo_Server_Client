package gamemsg

import (
	"LollipopGo/log"
	"LollipopGo2.8x/conf/g"
	. "LollipopGo2.8x/proto/sr_proto"
	"LollipopGo2.8x/tables"
	. "github.com/Golangltd/Twlib/proto"
	"github.com/mitchellh/mapstructure"
)

//巫师等级升级请求
type WizardUpgradeReq struct {
	Protocol  int    //主协议
	Protocol2 int    //子协议
	OpenId    string `json:"OpenId"`
}

//初始化巫师等级升级请求模块
func NewWizardUpgradeReq(data map[string]interface{}) *WizardUpgradeReq {
	m := &WizardUpgradeReq{}
	if err := mapstructure.Decode(data, m); err != nil {
		log.Error("data[%v] decode WizardUpgradeReq error,error is [%v]! ", data, err)
	}
	m.Protocol = GGameBattleProto
	m.Protocol2 = C2SWizardUpgradeReq
	return m
}

//巫师等级升级信息
type WizardLevelUpMsg struct {
	Protocol          int         //主协议
	Protocol2         int         //子协议
	WizardPastInfo    *WizardInfo //巫师过去的信息
	WizardCurrentInfo *WizardInfo //巫师当前的信息
}

//初始化巫师升级模块
func NewWizardLevelUpInfo() *WizardLevelUpMsg {
	m := &WizardLevelUpMsg{}
	m.Protocol = GGameBattleProto
	m.Protocol2 = S2CWizardUpgradeMsg
	return m
}

//巫师信息
type WizardInfo struct {
	SRLevel      int //学籍等级
	BattlePower  int //战斗力
	Life         int //生命
	AttackValue  int //攻击力
	DefenseValue int //防御力
}

func NewWizardInfo(level int) *WizardInfo {
	m := &WizardInfo{}
	m.SRLevel = level
	//TODO:通过接口获取战斗力
	m.BattlePower = 0
	sRSampleData := tables.SchoolRollTable[level]            //获取样本数据
	m.Life = sRSampleData.SRLevelAtr[g.LevelLife]            //添加生命
	m.AttackValue = sRSampleData.SRLevelAtr[g.LevelAttack]   //添加攻击力
	m.DefenseValue = sRSampleData.SRLevelAtr[g.LevelDefense] //添加防御力
	return m
}

/*
添加巫师等级升级信息
currentLevel:当前的等级
*/
func (m *WizardLevelUpMsg) AddWizardLevelUpInfo(currentLevel int) {
	m.WizardPastInfo = NewWizardInfo(currentLevel)
	if currentLevel < g.WizardMaxLevel { //升级
		currentLevel++
	}
	m.WizardCurrentInfo = NewWizardInfo(currentLevel)
}

//学院升级请求
type CollegeUpgradeReq struct {
	Protocol  int    //主协议
	Protocol2 int    //子协议
	OpenId    string `json:"OpenId"`
	CollegeID int    `json:"CollegeID"`
}

//初始化学院升级请求
func NewCollegeUpgradeReq(data map[string]interface{}) *CollegeUpgradeReq {
	m := &CollegeUpgradeReq{}
	if err := mapstructure.Decode(data, m); err != nil {
		log.Error("data[%v] decode CollegeUpgradeReq error,error is [%v]! ", data, err)
	}
	m.Protocol = GGameBattleProto
	m.Protocol2 = C2SCollegeUpgradeReq
	return m
}

type CollegeUpgradeMsg struct {
	Protocol       int      //主协议
	Protocol2      int      //子协议
	PastCollege    *College //过去的学院信息
	CurrentCollege *College //当前升级后的学院信息
}

func NewCollegeUpgradeMsg() *CollegeUpgradeMsg {
	m := &CollegeUpgradeMsg{}
	m.Protocol = GGameBattleProto
	m.Protocol2 = S2CCollegeUpgradeMsg
	return m
}

type College struct {
	Lv          int //学院等级
	ClassRoomLV int //教室等级
	SeatAmount  int //教室数量
	CourseAward *tables.CourseAward
}
