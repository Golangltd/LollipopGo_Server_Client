package examsystem

import "LollipopGo2.8x/tables"

//考试界面
type ExamInterface struct {
	IsUnlockHigh  bool //是否解锁高级考试
	RemainingDays int  //考试剩余天数
	ExpModel      *ExpModel
	AwardModel    *AwardModel //奖励模块
}

//经验条(经验模块)
type ExpModel struct {
	PracticeLV int  //实践等级
	CurrentExp int  //当前经验
	UpgradeExp int  //升级经验
	IsFullLV   bool //是否满级
}

//奖励model
type AwardModel struct {
	AwardDetailList []*AwardDetail //奖励详情列表
}

//奖励详情
type AwardDetail struct {
	ConditionLV int        //条件等级
	CommonAward *AwardInfo //普通奖励
	HighAward   *AwardInfo //高级奖励
}

//奖励信息
type AwardInfo struct {
	IsReceive bool //是否领取
	Awards    []*tables.Award
}

