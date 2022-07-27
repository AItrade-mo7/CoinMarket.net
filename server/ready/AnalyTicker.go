package ready

import (
	"fmt"

	"CoinMarket.net/server/okxApi/okxInfo"
)

type AnalyseType struct {
	StartTime int64 `json:"StartTime"`
	EndTime   int64 `json:"EndTime"`
}

// 开始进行市场分析
func AnalyseTicker() {
	for key, val := range okxInfo.MarketKdata {
		fmt.Println(key)
		AnalyseTicker_single(val)
	}
}

func AnalyseTicker_single(list []okxInfo.Kd) {
	StartTime := list[0].TimeUnix
	EndTime := list[len(list)-1].TimeUnix

	fmt.Println(StartTime, EndTime)
}
