package models

import (
	. "LollipopGo/network"
	"LollipopGo2.8x/conf/g"
	. "LollipopGo2.8x/cxt"
	. "LollipopGo2.8x/msg"
	"LollipopGo2.8x/tables"
)

type SRModel struct {
	OpenId string
	*Game
	*SRInterfaceMsg    //学籍界面信息
	*WizardLevelUpMsg  //巫师等级提升
	*CollegeDetailMsg  //学院详细信息
	*CollegeUpgradeMsg //学院升级信息
}

//学籍处理model
func NewSRModel(openID string, game *Game) *SRModel {
	m := &SRModel{}
	m.OpenId = openID
	m.Game = game
	m.SRInterfaceMsg = NewSRInterfaceMsg()
	m.WizardLevelUpMsg = NewWizardLevelUpInfo()
	m.CollegeDetailMsg = NewCollegeDetailInfo()
	m.CollegeUpgradeMsg = NewCollegeUpgradeMsg()
	return m
}

//巫师升级处理handler
func (m *SRModel) UpGrade() {
	if m.Game.UserInfo.RoleLev >= g.WizardMaxLevel { //如果用户当前的巫师等级处于最大值,那么就不用升级了
		return
	}
	sampleData := tables.SchoolRollTable[m.UserInfo.RoleLev]
	if m.UserInfo.RoleExp < sampleData.Exp { //用户经验小于升级所需经验,直接返回
		return
	}
	if sampleData.ChallengeID <= 0 { //无挑战关卡,直接升级
		m.WizardLevelUpHandle() //巫师升级处理
	} else {
		if !m.UserInfo.ClearanceDuplicates[sampleData.ChallengeID] { //没有通过试炼
			return
		}
		m.WizardLevelUpHandle()
	}
	PlayerSendToProxyServer(ConnXZ, m.WizardCurrentInfo, m.OpenId)
}

//获取学籍数据
func (m *SRModel) GetSRInterfaceData() {
	m.SRInterfaceMsg.InitFieldsData(m.Game.UserInfo)
}

//巫师升级处理
func (m *SRModel) WizardLevelUpHandle() {
	m.WizardLevelUpMsg.AddWizardLevelUpInfo(m.Game.UserInfo.GradeInfo.SRLevel)
	nameText := tables.SchoolRollTable[m.WizardCurrentInfo.SRLevel].SRName //获取等级关联的文本
	m.Game.UserInfo.UpdateGradeInfo(m.WizardCurrentInfo.SRLevel, nameText)
	m.Game.UserInfo.UpdateTotalPower(m.WizardCurrentInfo.BattlePower)
}

//学院详情处理
func (m *SRModel) CollegeDetailHandle(collegeID int) {
	collegeInfo := m.Game.UserInfo.CollegesInfo[collegeID]
	m.CollegeDetailMsg.AddDataInfo(collegeInfo)
}

//学院升级处理
func (m *SRModel) CollegeUpgradeHandle(collegeID int) {
	var (
		pastCollegeID    int
		currentCollegeID int
	)
	pastCollegeID = collegeID                      //没升级之前的ID
	pastCSD := tables.CollegeTable[collegeID]      //当前学院ID所处的学院样本数据
	if pastCSD.CollegeLevel == g.CollegeMaxLevel { //如果当前学院等级已经是最高等级了,直接返回
		return
	}
	m.CollegeUpgradeMsg.PastCollege = &College{
		Lv:          pastCSD.CollegeLevel,
		ClassRoomLV: pastCSD.ClassRoomLV,
		SeatAmount:  pastCSD.SeatAmount,
		CourseAward: pastCSD.CourseAward,
	}
	collegeID++
	currentCollegeID = collegeID                 //升级之后的ID
	currentCSD := tables.CollegeTable[collegeID] //升级之后的学院样本数据
	m.CollegeUpgradeMsg.CurrentCollege = &College{
		Lv:          currentCSD.CollegeLevel,
		ClassRoomLV: currentCSD.ClassRoomLV,
		SeatAmount:  currentCSD.SeatAmount,
		CourseAward: currentCSD.CourseAward,
	}
	m.Game.UserInfo.UpgradeHandle(pastCollegeID, currentCollegeID)
}
