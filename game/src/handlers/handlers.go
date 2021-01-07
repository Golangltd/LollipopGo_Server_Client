package handlers

import (
	"LollipopGo2.8x/handlers/drawcard"
	"LollipopGo2.8x/handlers/equip"
	"LollipopGo2.8x/handlers/schoolroll"
)

func init() {
	drawcard.InitDCHandlers()
	schoolroll.InitSRHandlers()
	equip.InitEquipHandlers()
}
