package models

import "sync"

var equipID int64 = 0 // 装备增长ID
var l sync.Mutex

// 获取增长指定步长的装备ID
func GetAndAddEquipID(add int64) int64 {
	l.Lock()
	defer l.Unlock()
	equipID = equipID + add
	return equipID
}

// 获取自增1的装备ID
func GetAndIncEquipID() int64 {
	return GetAndAddEquipID(1)
}