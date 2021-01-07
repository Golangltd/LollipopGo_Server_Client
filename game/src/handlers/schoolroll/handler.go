package schoolroll

import (
	"LollipopGo/log"
	impl "LollipopGo/network"
	"LollipopGo2.8x/cxt"
	. "LollipopGo2.8x/handlers/modules"
	"LollipopGo2.8x/models"
	msg "LollipopGo2.8x/msg"
	. "LollipopGo2.8x/proto/sr_proto"
	"golang.org/x/net/websocket"
)

func InitSRHandlers() {
	HM.AddHandler(C2PSInSR, SRHandleCallBack)
	HM.AddHandler(C2SGetAgainSRDataReq, GetAgainSRData)
	HM.AddHandler(C2SWizardUpgradeReq, WizardUpgradeHandle)
	HM.AddHandler(C2SGetCollegeDetail, CollegeDetailHandle)
	HM.AddHandler(C2SCollegeUpgradeReq, CollegeUpgradeHandle)
}

//获取学籍信息
func SRHandleCallBack(conn *websocket.Conn, ProtocolData map[string]interface{}) {
	var (
		value interface{}
		err   error
	)
	openID := ProtocolData["OpenId"].(string) //获取用户openID
	key := openID + cxt.UserKey               //拼接userKey,在内存中获取用户数据
	if value, err = cxt.M.Get(key); err != nil {
		log.Error("dont't get user data from M by current key[%v] !,err is [%v]", key, err)
		return
	}
	game := value.(*models.Game)
	game.SRModel = models.NewSRModel(openID, game)
	game.SRModel.GetSRInterfaceData()
	impl.PlayerSendToProxyServer(conn, game.SRModel.SRInterfaceMsg, game.SRModel.OpenId)
}

//再次获取学籍界面信息
func GetAgainSRData(conn *websocket.Conn, ProtocolData map[string]interface{}) {
	var (
		value interface{}
		err   error
	)
	data := msg.NewGetAgainSRData(ProtocolData)
	key := data.OpenId + cxt.UserKey //拼接userKey,在内存中获取用户数据
	if value, err = cxt.M.Get(key); err != nil {
		log.Error("dont't get user data from M by current key[%v] !,err is [%v]", key, err)
		return
	}
	game := value.(*models.Game)
	game.SRModel.GetSRInterfaceData()
	impl.PlayerSendToProxyServer(conn, game.SRModel.SRInterfaceMsg, game.SRModel.OpenId)
}

//巫师升级处理
func WizardUpgradeHandle(conn *websocket.Conn, ProtocolData map[string]interface{}) {
	var (
		value interface{}
		err   error
	)
	data := msg.NewWizardUpgradeReq(ProtocolData)
	key := data.OpenId + cxt.UserKey //拼接userKey,在内存中获取用户数据
	if value, err = cxt.M.Get(key); err != nil {
		log.Error("dont't get user data from M by current key[%v] !,err is [%v]", key, err)
		return
	}
	game := value.(*models.Game)
	game.SRModel.WizardLevelUpHandle()
	impl.PlayerSendToProxyServer(conn, game.SRModel.WizardLevelUpMsg, game.SRModel.OpenId)
}

//获取学院详情(学院详情处理)
func CollegeDetailHandle(conn *websocket.Conn, ProtocolData map[string]interface{}) {
	var (
		value interface{}
		err   error
	)
	data := msg.NewCollegeDetailReq(ProtocolData)
	key := data.OpenId + cxt.UserKey //拼接userKey,在内存中获取用户数据
	if value, err = cxt.M.Get(key); err != nil {
		log.Error("dont't get user data from M by current key[%v] !,err is [%v]", key, err)
		return
	}
	game := value.(*models.Game)
	game.SRModel.CollegeDetailHandle(data.CollegeID)
	impl.PlayerSendToProxyServer(conn, game.SRModel.CollegeDetailMsg, game.SRModel.OpenId)
}

func CollegeUpgradeHandle(conn *websocket.Conn, ProtocolData map[string]interface{}) {
	var (
		value interface{}
		err   error
	)
	data := msg.NewCollegeDetailReq(ProtocolData)
	key := data.OpenId + cxt.UserKey //拼接userKey,在内存中获取用户数据
	if value, err = cxt.M.Get(key); err != nil {
		log.Error("dont't get user data from M by current key[%v] !,err is [%v]", key, err)
		return
	}
	game := value.(*models.Game)
	game.SRModel.CollegeUpgradeHandle(data.CollegeID)
	impl.PlayerSendToProxyServer(conn, game.SRModel.CollegeUpgradeMsg, game.SRModel.OpenId)
}
