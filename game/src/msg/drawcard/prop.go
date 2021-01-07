package drawcard

import "LollipopGo2.8x/conf/g"

//道具数据
type PropData struct {
	ItemType int
	ItemId   int   //道具ID
	ItemNum  int64 //道具数量
}

func newPropData(id int, amount int64) *PropData {
	m := &PropData{}
	m.ItemId = id
	m.ItemNum = amount
	m.ItemType = g.CommonItem
	return m
}
