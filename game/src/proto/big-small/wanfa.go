package big_small

import "LollipopGo/tools/sample"

// 高低界限：10点
const (
	HiLo = 10
	Hi = "Hi"
	Lo = "Lo"
)

// 位置对应赔率
var RateItem = []int{
	1,1,1,1,1,1,
	5,5,5,5,5,5,
	5,5,5,5,1,1,
	2,4,5,5,4,2,
	5,5,5,2,5,5,
	3,3,2,5,5,5,
}

// 组合对应位置--> 获取位置-->通过位置得到赔率
var ZuHePos = []string{
	"1","2","3","4","5","6",
	"1+4","1+5","1+6","2+4","2+5","2+6",
	"1+3","3+5","3+6","4+6","Lo","Hi",
	"2+Lo","4+Lo","5+Lo","6+Lo","3+Hi","5+Hi",
	"11","1+2+3","4+5+6","1+Lo","1+2","2+3",
	"3+Lo","4+Hi","6+Hi","5+6","4+5","3+4",
}

// 特殊组合 全部是1X
var ZuHePos2 = []string{
	"1+2","1+3","2+3",
	"4+5","4+6","5+6",
}

// 随机骰子
var Item = []int{
	1,2,3,4,5,6,
}

// 创建每局的唯一ID
func CreatResult()[]int{
	var ret = []int{}
	for i:=0;i<3;i++{
		var mresult = sample.RandInt(1,Item[len(Item)-1])
		ret = append(ret,mresult)
	}
	return ret
}

// 创建红利 --> 开始于 投注结果
func CreatHLResult() []int{
	var ret = []int{}
	for i:=0;i<3;i++{
		var mresult = sample.RandInt(0,len(RateItem)-1)
		ret = append(ret,mresult)
	}
	return ret
}
