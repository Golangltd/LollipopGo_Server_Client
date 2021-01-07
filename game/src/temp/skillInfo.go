package temp

import twlib_user "github.com/Golangltd/Twlib/user"

func CreateSkillInfo() *twlib_user.SkillInfo {
	s := &twlib_user.SkillInfo{
		SkillId:  10060101,
		SkillLev: 1,
		Position: 0,
	}
	return s
}

