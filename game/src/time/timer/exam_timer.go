package timer

import "time"

/*
考试timer
*/

func init() {

}

//考试模块启动Timer
func ExamModelStartTimer() {
	timer := time.NewTimer(30 * time.Second)
	defer timer.Stop()
	select {
	case <-timer.C:

	}
}
