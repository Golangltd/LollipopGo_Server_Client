package tables

import (
	"LollipopGo/log"
	gameDB "LollipopGo2.8x/data"
	twLibDBTable "github.com/Golangltd/Twlib/dbtable"
	"sync"
)

/*
英雄卡牌图鉴表
*/
var (
	CardTJOnce   sync.Once
	CardTJMsg    []interface{}
	CardTJTables map[int]*CardTJSample
)

type CardTJSample struct {
	ID            int
	TJName        int         //图鉴名字
	CardID        int         //卡牌ID
	ActivationAtr map[int]int //激活属性
}

func AddCardTJTable() {
	data := gameDB.GetCFGameData(twLibDBTable.Gl_tj_yxkptjb)
	if data == nil {
		log.Fatal("error! receive skillTables is nil!!")
	}
	CardTJMsg = data
	CardTJOnce.Do(initCardTJTable)

}

func initCardTJTable() {
	CardTJTables = make(map[int]*CardTJSample)
	for _, sample := range CardTJMsg {
		temp := sample.(map[string]interface{})
		m := &CardTJSample{}
		m.ID = StrToInt(temp["sid"].(string))
		m.TJName = StrToInt(temp["handname"].(string))
		m.CardID = StrToInt(temp["character_id"].(string))
		m.ActivationAtr = InitAttribute(temp["award"].(string), `\d+,\d+`)
		CardTJTables[m.ID] = m
	}
	CardTJMsg = nil
}
