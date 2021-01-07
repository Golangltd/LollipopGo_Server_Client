package g

//抽卡
const (
	DisPlayCurrencyConfID = "126" //显示货币配置
	Gold                  = 1     //金币
)

type RecruitType int //招募类型
const (
	CommonDraw          RecruitType = 1 //普通抽
	FriendShipPointDraw RecruitType = 2 //友情点抽
	RaceDraw            RecruitType = 3 //种族抽
	IntegralDraw        RecruitType = 4 //积分抽
)

type DrawType int

const (
	OneDraws DrawType = 0 //1抽
	TenDraws DrawType = 1 //10抽
)
