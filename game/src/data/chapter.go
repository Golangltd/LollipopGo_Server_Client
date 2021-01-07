package gamedb

import (
	twlib_user "github.com/Golangltd/Twlib/user"
	"github.com/golang/glog"
)

type SChapter struct {
	RoleUID     int64
	Account     int64
	ChapterInfo *twlib_user.ChapterInfo
}

// 保存数据
func SaveChapterInfo(chapterinfo *SChapter) *int {
	glog.Info("---SaveChapterInfo:")
	iret := int(0)
	call := ConnRPC.Go("GameRPC.CreatChapterInfo", chapterinfo, &iret, nil)
	replyCall := <-call.Done
	glog.Info(call.Error)
	glog.Info(replyCall.Reply)
	return &iret
}
