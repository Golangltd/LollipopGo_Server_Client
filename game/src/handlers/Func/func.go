package Func

func init() {
	NewFunc()
}

// 初始化
func NewFunc() *FuncSt {
	return &FuncSt{}
}

// 是否解锁
// 参数1：解锁条件，key：功能表ID，val:条件
// 参数2：true:需要全部满足；false:有一个满足就ok
func IsFuncUnlock(data map[int]int, iscondition bool) bool {

	for k, _ := range data {

		if k == RoleLev {

		}else if k == CardLev {

		}else if k == VIPLev {

		}else if k == LoginDay {

		}else if k == StartServerDay {

		}else if k == CustomsCount {

		}else if k == FinishMasterTask {

		}else if k == CustomsPaTa {

		}else if k == UnionLev {

		}else if k == CashNum {

		}else if k == Power {

		}else if k == StadiumRank {

		}else if k == SumCashDay {

		}else if k == SignInDay {

		}else if k == NeedItemId {

		}else if k == NeedCardId {

		}
	}

	return false
}
