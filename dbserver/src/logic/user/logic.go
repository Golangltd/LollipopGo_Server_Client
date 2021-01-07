package user

//根据uid获取用户
func (r *Userrpc) GetUserByUID(uid int, reply *User) error {
	/*glog.Info("GetUserByOpenID uid:", uid)
	sqlstr := "select id, openid, mobile, wechat, nickname, avatar, is_guest, id_card, wealth, mobile_bind, wechat_bind, status from gl_user where id= ? "
	row := Mysyl_DB.DB.STdb.QueryRow(sqlstr, uid)

	user := new(User)
	err := row.Scan(&user.ID, &user.Openid, &user.Mobile, &user.Wechat, &user.Nickname, &user.Avatar, &user.IsGuest, &user.IdCcard, &user.Wealth, &user.MobileBind, &user.WechatBind, &user.Status)
	if err != nil {
		glog.Error("scan failed, err:%v", err)
		return nil
	}
	*reply = *user*/
	return nil
}