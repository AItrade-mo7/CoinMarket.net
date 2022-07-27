package analyse

// 构造新的Kdata
func NewKdata(pre, now okxInfo.Kd) okxInfo.Kd {
	Kdata := okxInfo.Kd{
		InstID:   now.InstID,
		TimeUnix: now.TimeUnix,
		Time:     now.Time,
		O:        now.O,
		H:        now.H,
		L:        now.L,
		C:        now.C,
		DataType: now.DataType,
		KRunType: now.KRunType,
	}

	if len(Kdata.C) < 1 {
		return Kdata
	}

	// 基础数据计算结果
	Kdata.Center = MA([]string{now.C, now.O, now.H, now.L}, 4)
	Kdata.Dir = count.Le(now.C, now.O)

	HLPer := count.Rose(now.H, now.L)
	Kdata.HLPer = count.CentPrice(HLPer, now.C)

	SolidPer := count.Rose(now.C, now.O)
	Kdata.SolidPer = count.CentPrice(SolidPer, now.C)

	RosePer := count.Rose(now.C, pre.C)
	Kdata.RosePer = count.CentPrice(RosePer, now.C)

	// 额外数据
	Kdata.C_dir = C_dirCount(Kdata, pre)

	U_shade, D_shade := ShadeCount(Kdata)
	Kdata.U_shade = count.CentPrice(U_shade, now.C)
	Kdata.D_shade = count.CentPrice(D_shade, now.C)

	EmaGather := EmaCount(Kdata)

	Kdata.Ema5 = EmaGather.Ema5
	Kdata.Ema15 = EmaGather.Ema15
	Kdata.Ema60 = EmaGather.Ema60

	// 均线相交
	Kdata.Ema_5A15 = Ema_5A15_func(pre, Kdata)
	Kdata.Ema_5A60 = Ema_5A60_func(pre, Kdata)

	// 均线(中点)相交
	Kdata.Ema_KA5 = Ema_KA5_func(pre, Kdata)
	Kdata.Ema_KA15 = Ema_KA15_func(pre, Kdata)
	Kdata.Ema_KA60 = Ema_KA60_func(pre, Kdata)

	// SAR 技术指标计算
	Sar, SarDir := SARCount(Kdata)

	Kdata.Sar = Sar
	Kdata.SarDir = SarDir

	// SAR 指标买卖点
	Kdata.Sar_K = Sar_KFun(pre, Kdata)

	return Kdata
}

func ShadeCount(now okxInfo.Kd) (U_shade, D_shade string) {
	if now.Dir > 0 { // 上涨时
		// 最高 - 收盘价 = 上影线
		U_shade = count.Rose(now.H, now.C)
		// 最低 - 开盘价 = 下影线
		D_shade = count.Rose(now.O, now.L)
	} else { // 下跌时
		// 最高 - 开盘价 = 上影线
		U_shade = count.Rose(now.H, now.O)
		// 最低 - 收盘价 = 下影线
		D_shade = count.Rose(now.C, now.L)
	}
	return
}

func C_dirCount(now, pre okxInfo.Kd) int {
	// 格子方向
	C_dir := count.Le(now.Center, pre.Center) // 以中心点为基准来计算，当前-过去的
	return C_dir
}

type KdataEmaGatherType struct {
	Ema5  string
	Ema15 string
	Ema60 string
}

// Ema 技术指标计算
func EmaCount(now okxInfo.Kd) KdataEmaGatherType {
	KdataC := make([]string, len(okxInfo.TradeKRunC))
	copy(KdataC, okxInfo.TradeKRunC)

	if now.KRunType == "dirKRun" {
		KdataC = okxInfo.DirKRunC
	}

	KdataC = append(KdataC, now.C)

	var Ema5 string
	var Ema15 string
	var Ema60 string

	Ema5 = EMA(KdataC, 5)
	Ema15 = EMA(KdataC, 15)
	Ema60 = EMA(KdataC, 60)

	return KdataEmaGatherType{
		Ema5:  Ema5,
		Ema15: Ema15,
		Ema60: Ema60,
	}
}

// SAR 技术指标计算
func SARCount(now okxInfo.Kd) (Value string, Dir int) {
	KdataList := make([]okxInfo.Kd, len(okxInfo.TradeKRunList))
	copy(KdataList, okxInfo.TradeKRunList)

	if now.KRunType == "dirKRun" {
		KdataList = make([]okxInfo.Kd, len(okxInfo.DirKRunList))
		copy(KdataList, okxInfo.DirKRunList)
	}

	KdataList = append(KdataList, now)

	Value, Dir = SAR(KdataList)

	return
}

// 5和15均线相交
func Ema_5A15_func(pre, now okxInfo.Kd) int {
	nowStatus := count.Le(now.Ema5, now.Ema15)
	preStatus := count.Le(pre.Ema5, pre.Ema15)

	Ema_5A15 := 0
	if nowStatus != preStatus {
		Ema_5A15 = nowStatus
	}
	// 消除噪点
	if Ema_5A15 > 0 {
		if count.Le(now.Ema15, pre.Ema15) > 0 && count.Le(now.C, now.Ema5) > 0 {
			// 如果 5 上穿 15 则 当前的 15 均线一定 大于 上一条 15 均线， K 线一定大于 Ema5， 今天的收盘价一定大于 EMA5
		} else {
			Ema_5A15 = 0
		}
	}

	if Ema_5A15 < 0 {
		if count.Le(now.Ema15, pre.Ema15) < 0 && count.Le(now.C, now.Ema5) < 0 {
			// 如果 5 下穿 15 则 当前的 15 均线一定 小于 上一条 15 均线，  K 线一定小于 Ema5，今天的收盘价一定小于 EMA5
		} else {
			Ema_5A15 = 0
		}
	}

	return Ema_5A15
}

// 5和60均线相交
func Ema_5A60_func(pre, now okxInfo.Kd) int {
	nowStatus := count.Le(now.Ema5, now.Ema60)
	preStatus := count.Le(pre.Ema5, pre.Ema60)

	Ema_5A60 := 0
	if nowStatus != preStatus {
		Ema_5A60 = nowStatus
	}
	return Ema_5A60
}

// K和5均线相交
func Ema_KA5_func(pre, now okxInfo.Kd) int {
	nowStatus := count.Le(now.Center, now.Ema5)
	preStatus := count.Le(pre.Center, pre.Ema5)

	Ema_KA5 := 0
	if nowStatus != preStatus {
		Ema_KA5 = nowStatus
	}
	return Ema_KA5
}

// K和15均线相交
func Ema_KA15_func(pre, now okxInfo.Kd) int {
	nowStatus := count.Le(now.Center, now.Ema15)
	preStatus := count.Le(pre.Center, pre.Ema15)

	Ema_KA15 := 0
	if nowStatus != preStatus {
		Ema_KA15 = nowStatus
	}
	return Ema_KA15
}

// K和60均线相交
func Ema_KA60_func(pre, now okxInfo.Kd) int {
	nowStatus := count.Le(now.Center, now.Ema60)
	preStatus := count.Le(pre.Center, pre.Ema60)

	Ema_KA60 := 0
	if nowStatus != preStatus {
		Ema_KA60 = nowStatus
	}
	return Ema_KA60
}

func Sar_KFun(pre, now okxInfo.Kd) int {
	Sar_K := 0
	if pre.SarDir != now.SarDir {
		Sar_K = now.SarDir
	}

	return Sar_K
}
