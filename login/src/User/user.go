package User


// 玩家的结构】
type UserSt struct {
	AccountId int64  // 账号ID
	RoleName string  // 账号对应角色的名字
	RoleSex  int     // 角色的性别
	RoleLev int      // 角色等级
	AreaCurrent int  // 角色当前的区
	AreaList map[int]*AreaSt //  所有的区的列表
	CardList map[int]*CardSt //  卡牌列表,所属区的，登录成功后下发

	// 技能配置， 不花费  3技能操作
	// 天赋效果，天赋自己升级 累加
	// 记录来源, 卡牌自身属性+养成属性 装备，携带身上
	// 数据操作，应用操作
}

// 游戏区的结构
type AreaSt struct {
	AreaId  int     // 区域服的ID数据
	AreaName string // 区域服的名字
	AreaUrl string  // 区域服务器的地址，玩家连接的地址
	AreaState int   // 区域服的状态
}

// 卡牌
type CardSt struct {
	CardId int
	CardLev int
	Skill map[int]*SkillSt
}

// 技能
type SkillSt struct {
	SkillId  int  // 技能的ID
}