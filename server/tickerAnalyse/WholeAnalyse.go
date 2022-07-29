package tickerAnalyse

import (
	"time"

	"CoinMarket.net/server/okxInfo"
	"github.com/EasyGolang/goTools/mCount"
	"github.com/EasyGolang/goTools/mStr"
)

type AnalyseType struct {
	StartTime     time.Time `json:"StartTime"`
	StartTimeUnix int64     `json:"StartTimeUnix"`
	EndTime       time.Time `json:"EndTime"`
	EndTimeUnix   int64     `json:"EndTimeUnix"`
	DiffHour      int64     `json:"DiffHour"`
}

/*

近 24 小时 币种排行榜的上涨和下跌情况

*/

var (
	Up_Inst []string
	Up_Num  []string
	Up_UR   string

	Down_Inst []string
	Down_Num  []string
	Down_UR   string
)

func WholeAnalyse() (resData okxInfo.WholeTickerAnalyseType) {
	resData = okxInfo.WholeTickerAnalyseType{}
	okxInfo.WholeTickerAnalyse = resData

	if len(okxInfo.TickerList) < 3 {
		return
	}

	for _, val := range okxInfo.TickerList {
		U_24_diff := mCount.Le(val.U_R24, "0")
		if U_24_diff > -1 {
			Up_Inst = append(Up_Inst, val.InstID)
			Up_Num = append(Up_Num, val.U_R24)
		} else {
			Down_Inst = append(Down_Inst, val.InstID)
			Down_Num = append(Down_Num, val.U_R24)
		}
	}

	// 上涨指数
	upN := mStr.ToStr(len(Up_Inst))
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

	okxInfo.WholeTickerAnalyse = resData
	return resData
}
