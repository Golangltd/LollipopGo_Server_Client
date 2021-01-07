package game

import (
	"LollipopGo/util"
	"fmt"
	twlib_user "github.com/Golangltd/Twlib/user"
	"github.com/golang/glog"
	"strconv"
)

type SChapter struct {
	RoleUID int64
	Account int64
	ChapterInfo *twlib_user.ChapterInfo
}

//  创建章节数据
func (m *GameRPC)CreatChapterInfo(data SChapter,reply *interface{}) error {

	// 更新
	if true{
		sqlStr := `update gl_chapter set chapterid=? ,chapterid2=?, roundid = ? where roleuid = ?`
		ret, err := GetGameConn().Exec(sqlStr, data.ChapterInfo.ChapterId,data.ChapterInfo.ChapterId2,data.ChapterInfo.RoundId,data.RoleUID)
		if err != nil {
			fmt.Printf("update failed, err:%v\n", err)
			return err
		}
		n, err := ret.RowsAffected() //RowsAffected 受影响的行数
		if err != nil {
			fmt.Printf("get id failed,err:%v\n", err)
			return err
		}
		if n>0{
			*reply = 1
			return nil
		}
	}


	sql:=  `insert into gl_chapter(roleuid, chapterid,chapterid2,roundid,createtime) 
           values(`
	val:= strconv.Itoa(int(data.RoleUID))+",'"+strconv.Itoa(int(data.ChapterInfo.ChapterId))+"','"+strconv.Itoa(int(data.ChapterInfo.ChapterId))+
		"',"+strconv.Itoa(int(data.ChapterInfo.RoundId))+ ","+util.UTCTime_LollipopGO()+")"
	sql=sql+val
	fmt.Println("-------------------->:",sql)
	ret, err := GetGameConn().Exec(sql)
	if err !=nil{
		fmt.Println(err.Error())
		return err
	}
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get id failed,err:%v\n", err)
		return err
	}
	*reply = id
	return nil
}

// 获取章节数据
func GetChapterInfo(RoleUid int64)  *twlib_user.ChapterInfo{

	sql := "select chapterid,chapterid2,roundid from gl_chapter where roleuid = '" +strconv.FormatInt(RoleUid, 10)+"'"
	rows, err := GetGameConn().Query(sql)
	glog.Info("-----",sql)
	if err != nil {
		fmt.Printf("exec %s query failed, err:%v\n", sql, err)
		return nil
	}
	// 3. 一定要关闭rows，才会释放连接（数据库的连接）
	defer rows.Close()
	// 4. 循环取值
	datatmp  := new(twlib_user.ChapterInfo)
	for rows.Next() {
		chapterid:=0
		chapterid2:=0
		roundid := 0
		err := rows.Scan(&chapterid, &chapterid2, &roundid)
		if err != nil {
			fmt.Printf("scan failed,err:%v\n", err)
		}
		glog.Info("======uid",chapterid)
		datatmp.ChapterId = chapterid
		datatmp.ChapterId2 = chapterid2
		datatmp.RoundId = roundid
	}
	glog.Info("======",datatmp)
	if datatmp.ChapterId == 0{
		datatmp.ChapterId = 1
		datatmp.ChapterId2 = 1
		datatmp.RoundId = 1
	}
	return datatmp
}