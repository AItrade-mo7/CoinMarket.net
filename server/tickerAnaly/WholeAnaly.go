package tickerAnaly

import (
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
		TickerVolumeList[key] = mOKX.SortAnalySlice_Volume(list) // 成交量排序
		TickerURList[key] = mOKX.SortAnalySlice_UR(list)         // 涨跌幅排序
	}

	for key, list := range TickerVolumeList {
		URList := TickerURList[key]
		Analy := TickerWholeAnaly(list, URList)
		TickerAnalyWhole = append(TickerAnalyWhole, Analy)
	}

	return mOKX.Sort_DiffHour(TickerAnalyWhole)
}

// list1 :成交量排序数组   list2 涨跌幅排序数组
func TickerWholeAnaly(list1, list2 []mOKX.AnalySliceType) (resData mOKX.TypeWholeTickerAnaly) {
	resData = mOKX.TypeWholeTickerAnaly{}

	VolList := make([]mOKX.AnalySliceType, len(list1)) // 成交量排序数组
	copy(VolList, list1)

	URList := make([]mOKX.AnalySliceType, len(list2)) // 涨跌幅排序数组
	copy(URList, list2)

	inciseURList := URList[1 : len(URList)-1] // 去除一个最高值和最低值
	newURList := make([]mOKX.AnalySliceType, len(inciseURList))
	copy(newURList, inciseURList)

	// 开始
	var (
		Up_Num   []string // 上涨幅度的集合
		Down_Num []string // 下跌幅度的集合
	)

	for _, val := range newURList {
		U_R_diff := mCount.Le(val.RosePer, "0") // 涨幅的正负

		if U_R_diff > 0 {
			Up_Num = append(Up_Num, val.RosePer)
		}
		if U_R_diff < 0 {
			Down_Num = append(Down_Num, val.RosePer)
		}
	}

	// 上涨指数
	upN := mStr.ToStr(len(Up_Num))
	allN := mStr.ToStr(len(Up_Num) + len(Down_Num))

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

	resData.StartTimeStr = VolList[0].StartTimeStr
	resData.StartTimeUnix = VolList[0].StartTimeUnix

	resData.EndTimeStr = VolList[0].EndTimeStr
	resData.EndTimeUnix = VolList[0].EndTimeUnix

	resData.MaxUP = URList[0] // 这里要按照涨跌幅 排序
	resData.MaxDown = URList[len(URList)-1]

	resData.DiffHour = VolList[0].DiffHour
	return resData
}
