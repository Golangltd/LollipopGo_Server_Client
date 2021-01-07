package game

import (
	"LollipopGo/util"
	"encoding/json"
	"fmt"
	twLibUser "github.com/Golangltd/Twlib/user"
	"github.com/golang/glog"
	"strconv"
)

type SCard struct {
	RoleUID int64
	Account int64
	Card    *twLibUser.CardInfo
}

func (m *GameRPC) CreatCards(data SCard, reply *interface{}) error {

	sql:=  `insert into gl_card(roleuid, accountid,cardid,lev,createtime,quality,skilllist,equiplist,attribute,isshow,star) 
           values(`

	// 技能
	savemap:=make(map[int]*twLibUser.SkillInfo)
	for _,v:=range data.Card.Skills{
		savemap[v.Position] = v
	}
	savestr,_:=json.Marshal(savemap)

	// 装备
	savemape:=make(map[int64]*twLibUser.EquipSt)
	for _,v:=range data.Card.Equips{
		if v == nil || v.UID ==0{
			continue
		}
		savemape[v.UID] = v
	}
	savestre,_:=json.Marshal(savemape)

	// 属性
	savestrea,_:=json.Marshal(data.Card.AttributeInfo)

	// 展示
	isshow := 0
	if data.Card.IsShow {
		isshow =  1
	}

	val:= strconv.Itoa(int(data.RoleUID))+",'"+strconv.Itoa(int(data.Account))+"','"+strconv.Itoa(int(data.Card.CardID))+
		"',"+strconv.Itoa(data.Card.Level)+ ","+util.UTCTime_LollipopGO()+","+strconv.Itoa(data.Card.Quality)+
		",'"+string(savestr)+
		"','"+string(savestre)+
		"','"+string(savestrea)+
		"',"+strconv.Itoa(isshow)+
		","+strconv.Itoa(data.Card.Stars)+
		")"
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

	// 更新UID
	if true{
		sqlStr := `update gl_card set carduid=? where id = ?`
		ret, err = GetGameConn().Exec(sqlStr, id,id)
		if err != nil {
			fmt.Printf("update failed, err:%v\n", err)
			return err
		}
		n, err := ret.RowsAffected() //RowsAffected 受影响的行数
		if err != nil {
			fmt.Printf("get id failed,err:%v\n", err)
			return err
		}
		fmt.Printf("更新了%d行数据\n", n)
	}

	*reply = id
	return nil
}

// 获取道具信息
func GetCardInfo(RoleUid int64)   []*twLibUser.CardInfo{

	sql := "select cardid, carduid,lev ,quality,skilllist,equiplist,attribute,star,isshow from gl_card where roleuid = '" +strconv.FormatInt(RoleUid, 10)+"'"
	rows, err := GetGameConn().Query(sql)
	glog.Info("-----",sql)
	if err != nil {
		fmt.Printf("exec %s query failed, err:%v\n", sql, err)
		return nil
	}
	// 3. 一定要关闭rows，才会释放连接（数据库的连接）
	defer rows.Close()
	// 4. 循环取值
	data := make(map[int64][]*twLibUser.CardInfo)
	dataslice := []*twLibUser.CardInfo{}
	for rows.Next() {
		datatmp  := new(twLibUser.CardInfo)
		dataSkills:=""
		dataEquips:=""
		dataAttributeInfo:=""
		datatmp.AttributeInfo = new(twLibUser.AttributeSt)
		err := rows.Scan(&datatmp.CardID, &datatmp.CardUid, &datatmp.Level,&datatmp.Quality,&dataSkills,&dataEquips,&dataAttributeInfo,&datatmp.Stars,&datatmp.IsShow)
		if err != nil {
			fmt.Printf("scan failed,err:%v\n", err)
		}
		json.Unmarshal([]byte(dataAttributeInfo),datatmp.AttributeInfo)
        dataresp:=[]*twLibUser.EquipSt{}
		err=json.Unmarshal([]byte(dataEquips),&dataresp)
		if err !=nil{
			fmt.Printf("exec failed, err:%v\n", err)
		}
		datatmp.Equips = dataresp
		dataslice = append(dataslice,datatmp)
	}
	data[RoleUid] = dataslice
	return data[RoleUid]
}

// 更新卡牌装备
func  (m *GameRPC)UpdateCards(data SCard, bret *bool) error {
	sqlStr := `update gl_card set equiplist=? where carduid = ?`
	savedata,_:=json.Marshal(data.Card.Equips)
	ret, err := GetGameConn().Exec(sqlStr,string(savedata),data.Card.CardUid)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return err
	}
	n, err := ret.RowsAffected() //RowsAffected 受影响的行数
	if err != nil {
		fmt.Printf("get id failed,err:%v\n", err)
		return err
	}
	fmt.Printf("更新了%d行数据\n", n)
	*bret = true
	return nil
}