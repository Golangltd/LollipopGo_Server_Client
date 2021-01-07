package gamemsg

import (
	"LollipopGo/log"
	. "LollipopGo2.8x/proto/sr_proto"
	. "github.com/Golangltd/Twlib/proto"
	"github.com/mitchellh/mapstructure"
)

//再次获取学籍信息
type GetAgainSRData struct {
	Protocol  int
	Protocol2 int
	OpenId    string `json:"OpenId"`
}

//再次获取学籍数据
func NewGetAgainSRData(data map[string]interface{}) *GetAgainSRData {
	m := &GetAgainSRData{}
	if err := mapstructure.Decode(data, m); err != nil {
		log.Error("data[%v] decode GetAgainSRData error,error is [%v]! ", data, err)
	}
	m.Protocol = GGameBattleProto
	m.Protocol2 = C2SGetAgainSRDataReq
	return m
}
