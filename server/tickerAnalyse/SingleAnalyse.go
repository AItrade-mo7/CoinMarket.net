package tickerAnalyse

import (
	"fmt"

	"CoinMarket.net/server/okxInfo"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mTime"
)

/*
单个币种历史数据分析

需要分析的部分：
近1小时上涨情况
近2小时上涨情况
近3小时上涨情况
近4小时上涨情况
近5小时上涨情况

榜单整体上涨情况
*/

func SingleAnalyse(list []okxInfo.Kd) (resData AnalyseType) {
	resData = AnalyseType{}
	if len(list) < 3 {
		return
	}
	InstID := list[0].InstID

	resData.StartTime = list[0].Time
	resData.StartTimeUnix = list[0].TimeUnix
	resData.EndTime = list[len(list)-1].Time
	resData.EndTimeUnix = list[len(list)-1].TimeUnix
	resData.DiffHour = (resData.EndTimeUnix - resData.StartTimeUnix) / mTime.UnixTimeInt64.Hour

	fmt.Println(InstID)
	mJson.Println(resData)

	return
}
