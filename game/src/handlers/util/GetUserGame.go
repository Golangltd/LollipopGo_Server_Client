package util_handlers

import (
	. "LollipopGo2.8x/cxt"
	"LollipopGo2.8x/models"
	"errors"
	twLibUser "github.com/Golangltd/Twlib/user"
	concurrent "github.com/fanliao/go-concurrentMap"
	"github.com/golang/glog"
)

// 获取玩家Game和UserInfo
func GetGameAndUser(openID string) (*models.Game, *twLibUser.UserSt, error) {
	// 获取玩家
	vac, err := M.Get(openID + UserKey)
	if err != nil {
		glog.Error("GetGameAndUser, cant find modules.Game, err:", err)
		return nil, nil, errors.New("GetGameAndUser, cant find modules.Game, err:" + err.Error())
	}

	game := vac.(*models.Game)
	if game.UserInfo == nil {
		return nil, nil, errors.New("GetGameAndUser, cant find game.UserInfo")
	}
	return game, game.UserInfo, nil
}

// 修改玩家 GameInfo
func SetGameInfo(openID string, savedata models.Game) {
	vac, err := M.Get(openID + UserKey)
	if err != nil {
		glog.Error("GetGameAndUser, cant find modules.Game, err:", err)
		return
	}
	*vac.(*models.Game) = savedata
}

func GetGameAllUser() *concurrent.ConcurrentMap {
	return M
}
