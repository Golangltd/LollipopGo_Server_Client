package models

type Tsts struct {
	Id int
}

//---------------------------------------------------------------------
// 测试架构数据
type TJiaGou struct {
	SID  int
	Name string
	*Game  // 继承基类
}

// 初始化调用
// Msg_data = new(TJiaGou)

// 方法调用
func (m *TJiaGou)Test(){

}

//-------------各个模块【活动】 例子 观察者模式-----------------------------------
//抽象主题
type Msg_data interface {
	Add(o Observer)
    Send(str string)
}

//抽象观察者
type Observer interface {
    Receive(str string)
}

