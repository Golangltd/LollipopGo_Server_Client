package st

// 邮件类型
const (
	 MAILTYPE = iota
	 MAILSYS       // MAILSYS == 1 系统邮件
	 MAIACTIVITY   // MAIACTIVITY == 2 活动邮件
	 MAILUNION     // MAILUNION == 3 工会邮件
	 MAILPERSONSAL // MAILPERSONSAL == 4 个人邮件
)

// 邮件状态
const (
	MAILSTYATE = iota // 未读状态
	MAILREADED        // MAILREADED ==  1 已读未领取
	MAILREADEDDRAW    // MAILREADEDDRAW == 2 已读已经领取
)
