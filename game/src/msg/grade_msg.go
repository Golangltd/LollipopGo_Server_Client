package gamemsg

import (
	"LollipopGo2.8x/conf/g"
	"LollipopGo2.8x/tables"
	twlib_user "github.com/Golangltd/Twlib/user"
)

/*
评级消息
*/

//用户评级消息
type GradeMsg struct {
	CurrentLv       int       //当前学籍等级
	CurrentLVName   int       //等级名称(文本ID)
	CurrentExp      int       //当前经验
	UpgradeTotalExp int       //升级需要的总经验
	NextLV          int       //下级级数
	NextLVName      int       //下级等级名称
	IsMeetUpgrade   bool      //是否满足升级条件
	TrailMsg        *TrialMsg //试炼条件
}

//生成评级消息
func newGradeMsg(st *twlib_user.UserSt) *GradeMsg {
	m := &GradeMsg{}
	m.CurrentLv = st.RoleLev
	m.CurrentLVName = tables.SchoolRollTable[m.CurrentLv].SRName //学籍名称(关联文本ID)
	m.CurrentExp = st.RoleExp                                    //当前经验值
	m.UpgradeTotalExp = tables.SchoolRollTable[m.CurrentLv].Exp  //升级到下一级需要的总经验
	isFullLv := m.isFullLevel()
	m.addNextLVData(isFullLv)
	if !isFullLv { //如果没有满级
		m.upgradeConditionHandle(st.ClearanceDuplicates, st)
	}
	return m
}

func (m *GradeMsg) isFullLevel() bool {
	return m.CurrentLv >= g.WizardMaxLevel
}

func (m *GradeMsg) addNextLVData(isFullLV bool) {
	if !isFullLV { //没有满级
		m.NextLV = m.CurrentLv + 1
		m.NextLVName = tables.SchoolRollTable[m.NextLV].SRName
	} else { //满级
		m.NextLV = m.CurrentLv
		m.NextLVName = m.CurrentLVName
	}
}

//升级条件处理
func (m *GradeMsg) upgradeConditionHandle(duplicatesInfo map[int]bool, userST *twlib_user.UserSt) {
	challengeID := tables.SchoolRollTable[m.CurrentLv].ChallengeID //副本ID
	if m.CurrentExp >= m.UpgradeTotalExp && challengeID <= 0 {     //经验满足升级经验并且当前无挑战副本直接升级
		m.IsMeetUpgrade = true
	}
	if m.CurrentExp >= m.UpgradeTotalExp && challengeID > 0 { //如果经验满足但是需要挑战副本时
		if duplicatesInfo[challengeID] { //如果通过此副本,那么就满足升级要求.否则就默认不满足升级条件
			m.IsMeetUpgrade = true
		}
	}
	m.trialConditionHandle(userST)
}

//试炼条件处理
func (m *GradeMsg) trialConditionHandle(userST *twlib_user.UserSt) {
	challengeID := tables.SchoolRollTable[m.CurrentLv].ChallengeID //学籍试炼ID
	sampleData := tables.SRTrialTable[challengeID]
	trialCondition := sampleData.TrialCondition //试炼条件
	for _, collegeInfo := range userST.CollegesInfo {
		collegeSample := tables.CollegeTable[collegeInfo.CollegeID] //学院样本数据
		if collegeSample.CollegeType == trialCondition.CollegeType {
			tc := newTrialCondition(trialCondition.CollegeType, collegeSample.CollegeLevel, trialCondition.CollegeLV)
			m.TrailMsg = newTrialMsg(collegeSample.CollegeLevel >= trialCondition.CollegeLV, tc)
			break
		}
	}
}
