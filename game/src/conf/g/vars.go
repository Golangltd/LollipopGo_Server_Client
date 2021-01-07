package g

var (
	WizardMaxLevel       int                                   //巫师最大等级
	CollegeMaxLevel      int                                   //学院的最大等级
	CDUsedFreeTimes      map[int64]int                         //普通抽奖单日已经使用的免费次数 key-->用户ID
	UsersCardsPoolUpRate map[int64]map[RecruitType]map[int]int //保存用户不同招募类型的不同卡池的幸运增长比
)

func init() {
	CDUsedFreeTimes = make(map[int64]int) //初始化内存
	UsersCardsPoolUpRate = make(map[int64]map[RecruitType]map[int]int)
}


