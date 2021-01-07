package gamemsg

import (
	"LollipopGo2.8x/conf/g"
	. "LollipopGo2.8x/proto/sr_proto"
	"LollipopGo2.8x/tables"
	. "github.com/Golangltd/Twlib/proto"
	. "github.com/Golangltd/Twlib/user"
)

/*
学籍界面消息
*/
type SRInterfaceMsg struct {
	Protocol    int            //主协议
	Protocol2   int            //子协议
	UserName    string         //用户名
	StudentID   int64          //学号
	GradeInfo   *GradeMsg      //评级信息
	Association string         //协会
	BattlePower int            //战力
	Colleges    []*CollegeData //学院列表信息
}

func NewSRInterfaceMsg() *SRInterfaceMsg {
	m := &SRInterfaceMsg{}
	m.Protocol = GGameBattleProto
	m.Protocol2 = S2CSRInterfaceMsgProto2
	return m
}

//发送给客户端的学院数据
type CollegeData struct {
	ID          int //学院ID
	CollegeType int //学院类型
	Level       int //学院等级
}

func newCollegeData(info *CollegeInfo) *CollegeData {
	m := &CollegeData{}
	m.ID = info.CollegeID                                 //学院ID
	m.Level = tables.CollegeTable[m.ID].CollegeLevel      //学院等级
	m.CollegeType = tables.CollegeTable[m.ID].CollegeType //学院类型
	return m
}

func (m *SRInterfaceMsg) InitFieldsData(st *UserSt) {
	m.UserName = st.RoleName
	m.StudentID = st.RoleUid
	m.GradeInfo = newGradeMsg(st) //生成评级信息
	m.Association = st.Association
	m.BattlePower = st.TotalPower
	m.InitColleges(st.CollegesInfo)
}

//初始化学院列表信息
func (m *SRInterfaceMsg) InitColleges(collegesInfo map[g.CollegeID]*CollegeInfo) {
	m.Colleges = make([]*CollegeData, 0)
	for _, c := range collegesInfo {
		m.Colleges = append(m.Colleges, newCollegeData(c))
	}
	m.sortColleges()
}

//排序学院列表
func (m *SRInterfaceMsg) sortColleges() {
	if len(m.Colleges) <= 1 {
		return
	}
	for i := 0; i < len(m.Colleges)-1; i++ {
		for j := i + 1; j > 0; j-- {
			if m.Colleges[j].ID < m.Colleges[i].ID {
				m.Colleges[j], m.Colleges[i] = m.Colleges[i], m.Colleges[j]
			} else {
				break
			}
		}
	}
}
