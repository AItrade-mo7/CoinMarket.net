package tickerAnaly

import (
	"fmt"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"github.com/EasyGolang/goTools/mCount"
	"github.com/EasyGolang/goTools/mOKX"
	"github.com/EasyGolang/goTools/mStr"
)

/*

近 24 小时 币种排行榜的上涨和下跌情况

*/

func WholeAnaly(TickerAnalySingle map[string][]mOKX.AnalySliceType) []mOKX.TypeWholeTickerAnaly {
	TickerAnalyWhole := []mOKX.TypeWholeTickerAnaly{}

	TickerSingle := make(map[int][]mOKX.AnalySliceType)

	TickerVolumeList := make(map[int][]mOKX.AnalySliceType)

	TickerURList := make(map[int][]mOKX.AnalySliceType)

	for _, Slice := range TickerAnalySingle {
		for _, Single := range Slice {
			TickerSingle[Single.DiffHour] = append(TickerSingle[Single.DiffHour], Single)
		}
	}

	if len(TickerSingle) != len(config.SliceHour) {
		global.LogErr("tickerAnaly.WholeAnaly  config.SliceHour 长度不正确", len(TickerSingle), len(config.SliceHour))
		return nil
	}

	for key, list := range TickerSingle {
		TickerVolumeList[key] = mOKX.SortAnalySlice_Volume(list)
		TickerURList[key] = mOKX.SortAnalySlice_UR(list)
	}

	for key, list := range TickerVolumeList {
		URList := TickerURList[key]
		Analy := TickerWholeAnaly(list, URList)
		TickerAnalyWhole = append(TickerAnalyWhole, Analy)
	}

	return mOKX.Sort_DiffHour(TickerAnalyWhole)
}

func TickerWholeAnaly(list, URList []mOKX.AnalySliceType) (resData mOKX.TypeWholeTickerAnaly) {
	resData = mOKX.TypeWholeTickerAnaly{}

	// 开始
	var (
		Up_Num   []string // 上涨幅度的集合
		Down_Num []string // 下跌幅度的集合
	)

	for _, val := range list {
		U_R_diff := mCount.Le(val.RosePer, "0")

		if U_R_diff >= 0 {
			Up_Num = append(Up_Num, val.RosePer)
		} else {
			Down_Num = append(Down_Num, val.RosePer)
		}
	}

	fmt.Println(Up_Num, Down_Num)

	// 上涨指数
	upN := mStr.ToStr(len(Up_Num))
	allN := mStr.ToStr(len(list))

	resData.UPIndex = mCount.PerCent(upN, allN)

	// 涨跌均值
	upAvg := mCount.Average(Up_Num)
	downAvg := mCount.Average(Down_Num)
	UDAvg := mCount.Add(upAvg, downAvg)
	resData.UDAvg = mCount.Cent(UDAvg, 2)

	// 上涨指数 计算
	resData.UPLe = mCount.Le(resData.UPIndex, "50")
	// 综合涨幅均值 计算
	resData.UDLe = mCount.Le(resData.UDAvg, "0")

	resData.DirIndex = 0

	if resData.UPLe > 0 && resData.UDLe > 0 {
		resData.DirIndex = 1
	} else if resData.UPLe < 0 && resData.UDLe < 0 {
		resData.DirIndex = -1
	}

	resData.StartTimeStr = list[0].StartTimeStr
	resData.StartTimeUnix = list[0].StartTimeUnix

	resData.EndTimeStr = list[0].EndTimeStr
	resData.EndTimeUnix = list[0].EndTimeUnix

	resData.MaxUP = URList[0] // 这里要按照涨跌幅 排序
	resData.MaxDown = URList[len(URList)-1]

	resData.DiffHour = list[0].DiffHour
	return resData
}
