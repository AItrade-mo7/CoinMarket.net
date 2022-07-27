package analyse

import (
	"CoinMarket.net/server/okxApi/okxInfo"
	"github.com/EasyGolang/goTools/mCount"
)

// 构造新的Kdata
func NewKdata(now okxInfo.Kd, list []okxInfo.Kd) (kdata okxInfo.Kd) {
	kdata = okxInfo.Kd{
		InstID:   now.InstID,
		TimeUnix: now.TimeUnix,
		Time:     now.Time,
		O:        now.O,
		H:        now.H,
		L:        now.L,
		C:        now.C,
		Vol:      now.Vol,
		VolCcy:   now.VolCcy,
		Type:     now.Type,
	}

	if mCount.Le("0", now.C) > -1 {
		return
	}

	kdata.Dir = mCount.Le(kdata.C, kdata.O)

	Center := mCount.Average([]string{now.C, now.O, now.H, now.L})
	kdata.Center = mCount.PriceCent(Center, now.C)

	HLPer := mCount.Rose(now.H, now.L)
	kdata.HLPer = mCount.PriceCent(HLPer, now.C)

	SolidPer := mCount.Rose(now.C, now.O)
	kdata.SolidPer = mCount.PriceCent(SolidPer, now.C)

	U_shade, D_shade := ShadeCount(kdata)
	kdata.U_shade = mCount.PriceCent(U_shade, now.C)
	kdata.D_shade = mCount.PriceCent(D_shade, now.C)

	if len(list) < 1 {
		return
	}
	Pre := list[len(list)-1]
	RosePer := mCount.Rose(now.C, Pre.C)
	kdata.RosePer = mCount.PriceCent(RosePer, now.C)
	kdata.C_dir = C_dirCount(kdata, Pre)

	return
}

func ShadeCount(now okxInfo.Kd) (U_shade, D_shade string) {
	if now.Dir > 0 { // 上涨时
		// 最高 - 收盘价 = 上影线
		U_shade = mCount.Rose(now.H, now.C)
		// 最低 - 开盘价 = 下影线
		D_shade = mCount.Rose(now.O, now.L)
	} else { // 下跌时
		// 最高 - 开盘价 = 上影线
		U_shade = mCount.Rose(now.H, now.O)
		// 最低 - 收盘价 = 下影线
		D_shade = mCount.Rose(now.C, now.L)
	}
	return
}

func C_dirCount(now, pre okxInfo.Kd) int {
	// 格子方向
	C_dir := mCount.Le(now.Center, pre.Center) // 以中心点为基准来计算，当前-过去的
	return C_dir
}
