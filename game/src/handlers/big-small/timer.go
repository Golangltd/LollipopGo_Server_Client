package bigsmall

import (
	bismark "LollipopGo2.8x/proto/big-small"
	"time"
)

var TickerCount  = 0  // 更新记时
var TickerALLCount  = 50

func init()  {
   go GameStateTurn()
}

// 状态切换
func GameStateTurn()  {
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ticker.C:
			TickerCount++
			for i:=1;i<5;i++{
				aid :=TickerCount%TickerALLCount
				imitate,_ :=GetRoomState(i)
				if (imitate == bismark.Prepare)&&(aid >=2){
					UpdateRoomState(i,15*1000)
				}else if (imitate == bismark.ChipIn)&&(aid >=17){
					UpdateRoomState(i,12*1000)
				}else if (imitate == bismark.ChipInOver)&&(aid >=25){
					UpdateRoomState(i,11*1000)
				}else if (imitate == bismark.Settle)&&(aid >=36){
					UpdateRoomState(i,13*1000)
				}else if (imitate == bismark.Wait)&&(aid >=49){
					UpdateRoomState(i,1*1000)
				}
			}
		default:
		}
	}
}