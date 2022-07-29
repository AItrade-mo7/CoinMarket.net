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

/*
需要分析的部分：
近1小时上涨情况
近2小时上涨情况
近3小时上涨情况
近4小时上涨情况
近5小时上涨情况

榜单整体上涨情况

*/

// 开始进行市场分析
func AnalyseTicker_single(list []okxInfo.Kd) (resData AnalyseType) {
	InstID := list[0].InstID
	resData = AnalyseType{}

	resData.StartTime = list[0].TimeUnix
	resData.EndTime = list[len(list)-1].TimeUnix
	resData.DiffHour = (resData.EndTime - resData.StartTime) / mTime.UnixTimeInt64.Hour

	fmt.Println(InstID)
	mJson.Println(resData)

	return
}
