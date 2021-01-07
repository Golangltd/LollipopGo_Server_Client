package models

import "container/list"

// 模块化数据接口
type IModel interface {
	RegisterModel(modeName interface{})
	SaveModelInfo(modeName interface{},data interface{})
    GetModelInfo(modeName interface{})interface{}
	SendOtherModel(modeName interface{},data interface{})
	ReceiveOtherModel()
}

// 订阅主题
type HotSubject struct {
	title string
	l     *list.List
}

func (sub *HotSubject) RegisterModel(o Observer) {
	sub.l.PushBack(o)
}