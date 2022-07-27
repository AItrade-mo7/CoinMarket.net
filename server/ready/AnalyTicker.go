package ready

import (
	"fmt"

	"CoinMarket.net/server/okxApi/okxInfo"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mTime"
)

type AnalyseType struct {
	StartTime int64 `json:"StartTime"`
	EndTime   int64 `json:"EndTime"`
	DiffHour  int64 `json:"DiffHour"`
}

// 开始进行市场分析
func AnalyseTicker() {
	for key, val := range okxInfo.MarketKdata {
		fmt.Println(key)
		AnalyseTicker_single(val)
	}
}

func AnalyseTicker_single(list []okxInfo.Kd) (resData AnalyseType) {
	resData = AnalyseType{}

	resData.StartTime = list[0].TimeUnix
	resData.EndTime = list[len(list)-1].TimeUnix
	resData.DiffHour = (resData.EndTime - resData.StartTime) / mTime.UnixTimeInt64.Hour

	mJson.Println(resData)

	return
}
