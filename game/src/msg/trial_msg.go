package gamemsg

/*
试炼消息
*/

//试炼信息
type TrialMsg struct {
	IsMeetTrial    bool            //是否满足试炼
	TrialCondition *TrialCondition //试炼条件
}

func newTrialMsg(isMeetTrial bool, trialCondition *TrialCondition) *TrialMsg {
	m := &TrialMsg{}
	m.IsMeetTrial = isMeetTrial
	m.TrialCondition = trialCondition
	return m
}

//试炼条件
type TrialCondition struct {
	CollegeType      int //学院类型
	CollegeCurrentLv int //学院当前等级
	ConditionLv      int //条件等级
}

func newTrialCondition(collegeType int, collegeCurrentLV int, conditionLv int) *TrialCondition {
	m := &TrialCondition{}
	m.CollegeType = collegeType
	m.CollegeCurrentLv = collegeCurrentLV
	m.ConditionLv = conditionLv
	return m
}
