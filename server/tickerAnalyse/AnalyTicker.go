package tickerAnalyse

import (
	"fmt"
	"time"

	"CoinMarket.net/server/okxApi/okxInfo"
	"github.com/EasyGolang/goTools/mCount"
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

func WholeAnalyse() {
	for _, val := range okxInfo.TickerList {
		U_24_diff := mCount.Le(val.U_R24, "0")
		if U_24_diff > -1 {
			Up_Inst = append(Up_Inst, val.InstID)
		} else {
			Down_Inst = append(Down_Inst, val.InstID)
		}
	}

	fmt.Println("UP_Inst", Up_Inst)
	fmt.Println("Down_Inst", Down_Inst)
}

// 开始进行市场分析
