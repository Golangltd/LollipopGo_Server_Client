package test

import (
	"LollipopGo2.8x/logic/game"
	"fmt"
	twLibUser "github.com/Golangltd/Twlib/user"
)

func GetRoleDataTest() {
	m := &game.GameRPC{}
	if err := m.GetRoleInfo(1, &twLibUser.UserSt{}); err != nil {
		fmt.Println(err.Error())
	}
}
