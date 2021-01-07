package tables


// 场景表
type SceneSt struct {
	Sid      int
	Name     string
	MapId    int
	MapType  int      // 场景类型
	Point    *PointSt // 出生点
	MusicId  int      // 资源Id
	NpcPoint string   // NPC的坐标点[1，10，10]坐标点
}

// 进入场景的坐标
type PointSt struct {
	PosX int
	PosY int
}

// NPC坐标点
type NPCPointSt struct {
	Sid   int      // NPCId
	Point *PointSt // 出生点
}
