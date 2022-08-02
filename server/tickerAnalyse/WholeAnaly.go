package tickerAnalyse

import (
	"CoinMarket.net/server/okxInfo"
	"github.com/EasyGolang/goTools/mCount"
	"github.com/EasyGolang/goTools/mOKX"
	"github.com/EasyGolang/goTools/mStr"
)

/*

近 24 小时 币种排行榜的上涨和下跌情况

*/

func WholeAnaly() (resData mOKX.TypeWholeTickerAnaly) {
	resData = mOKX.TypeWholeTickerAnaly{}
	okxInfo.TickerAnalyWhole = resData

	if len(okxInfo.TickerList) < 3 {
		return
	}

	// 开始

	var (
		Up_Num   []string
		Down_Num []string
	)

	for _, val := range okxInfo.TickerList {
		U_24_diff := mCount.Le(val.U_R24, "0")
		if U_24_diff > -1 {
			Up_Num = append(Up_Num, val.U_R24)
		} else {
			Down_Num = append(Down_Num, val.U_R24)
		}
	}

	// 上涨指数
	upN := mStr.ToStr(len(Up_Num))
	allN := mStr.ToStr(len(okxInfo.TickerList))

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

	resData.Ts = okxInfo.TickerList[0].Ts

	resData.MaxUP = okxInfo.TickerU_R24[0]
	resData.MaxDown = okxInfo.TickerU_R24[len(okxInfo.TickerU_R24)-1]

	okxInfo.TickerAnalyWhole = resData
	return resData
}
