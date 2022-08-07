package tickerAnaly

import (
	"CoinMarket.net/server/okxInfo"
	"github.com/EasyGolang/goTools/mCount"
	"github.com/EasyGolang/goTools/mOKX"
	"github.com/EasyGolang/goTools/mStr"
)

/*

近 24 小时 币种排行榜的上涨和下跌情况

*/

func WholeAnaly() {
	TickerAnalyWhole := []mOKX.TypeWholeTickerAnaly{}

	okxInfo.TickerAnalyWhole = []mOKX.TypeWholeTickerAnaly{}

	TickerSingle := make(map[int][]mOKX.AnalySliceType)

	TickerVolumeList := make(map[int][]mOKX.AnalySliceType)

	/*
		币种时间切片
		okxInfo.TickerAnalySingle
	*/

	for _, Slice := range okxInfo.TickerAnalySingle {
		for _, Single := range Slice {
			TickerSingle[Single.DiffHour] = append(TickerSingle[Single.DiffHour], Single)
		}
	}

	for key, list := range TickerSingle {
		TickerVolumeList[key] = mOKX.SortAnalySlice_Volume(list)
	}
	for _, list := range TickerVolumeList {
		Analy := TickerWholeAnaly(list)
		TickerAnalyWhole = append(TickerAnalyWhole, Analy)
	}

	okxInfo.TickerAnalyWhole = mOKX.Sort_DiffHour(TickerAnalyWhole)
}

func TickerWholeAnaly(list []mOKX.AnalySliceType) (resData mOKX.TypeWholeTickerAnaly) {
	resData = mOKX.TypeWholeTickerAnaly{}

	if len(list) != len(okxInfo.TickerList) {
		return
	}

	// 开始
	var (
		Up_Num   []string // 上涨幅度的集合
		Down_Num []string // 下跌幅度的集合
	)

	for _, val := range list {
		U_R_diff := mCount.Le(val.RosePer, "0")
		if U_R_diff > -1 {
			Up_Num = append(Up_Num, val.RosePer)
		} else {
			Down_Num = append(Down_Num, val.RosePer)
		}
	}

	// 上涨指数
	upN := mStr.ToStr(len(Up_Num))
	allN := mStr.ToStr(len(list))

	resData.UPIndex = mCount.PerCent(upN, allN)

	// 涨跌均值
	upAvg := mCount.Average(Up_Num)
	downAvg := mCount.Average(Down_Num)
	UDAvg := mCount.Add(upAvg, downAvg)
	resData.UDAvg = mCount.Cent(UDAvg, 3)

	// 上涨趋势方向
	resData.UPLe = mCount.Le(resData.UPIndex, "50")
	// 涨幅均值方向
	resData.UDLe = mCount.Le(resData.UDAvg, "0")

	resData.DirIndex = 0
	if resData.UPLe > 0 && resData.UDLe > 0 {
		resData.DirIndex = 1
	} else if resData.UPLe < 0 && resData.UDLe < 0 {
		resData.DirIndex = -1
	}

	resData.Ts = list[0].EndTimeUnix

	resData.MaxUP = list[0]
	resData.MaxDown = list[len(list)-1]
	resData.DiffHour = list[0].DiffHour

	return resData
}
