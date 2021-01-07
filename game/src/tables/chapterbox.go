package tables

// 章节宝箱
type ChapterBoxSt struct {
	Sid            int              // id 信息
	BoxResId       string           // 宝箱资源Id
	BosResEffectId string           // 资源特效Id
	AwardList      map[int]*AwardSt // 奖励数组
}

// 奖励结构类型
type AwardSt struct {
	ItemType int
	ItemId   int
	ItemNum  int   
}