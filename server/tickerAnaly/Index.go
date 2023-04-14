package tickerAnaly

import (
	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"github.com/EasyGolang/goTools/mOKX"
	"github.com/EasyGolang/goTools/mTime"
)

/*
18  80   340

*/

type TickerAnalyParam struct {
	TickerVol   []mOKX.TypeTicker
	TickerKdata map[string][]mOKX.TypeKd
}

type AnalyResult struct {
	AnalySingle map[string][]mOKX.AnalySliceType
	AnalyWhole  []mOKX.TypeWholeTickerAnaly
	MillionCoin []mOKX.AnalySliceType
	TimeUnix    int64
	TimeStr     string
	TimeID      string
}

func GetAnaly(opt TickerAnalyParam) AnalyResult {
	// 进行数据分析和计算
	global.Run.Println(
		"== 开始分析 ==",
		len(opt.TickerVol),
		len(opt.TickerKdata),
		mOKX.GetTimeID(opt.TickerVol[0].TimeUnix),
	)

	TickerAnalySingle := make(map[string][]mOKX.AnalySliceType)
	for _, item := range opt.TickerVol {
		list := opt.TickerKdata[item.InstID]
		Single := NewSingle(list)
		if len(Single.ResData) == len(config.SliceHour) {
			TickerAnalySingle[Single.Info.InstID] = Single.ResData
		} else {
			global.LogErr("tickerAnaly.Start  数据长度不足", Single.Info.InstID, len(Single.ResData))
		}
	}

	TickerAnalyWhole := WholeAnaly(TickerAnalySingle)
	MillionCoin := AnalyDir(AnalyDirParam{
		TickerVol:   opt.TickerVol,
		AnalyWhole:  TickerAnalyWhole,
		AnalySingle: TickerAnalySingle,
	})

	var Analy AnalyResult
	Analy.AnalySingle = TickerAnalySingle
	Analy.AnalyWhole = TickerAnalyWhole
	Analy.MillionCoin = MillionCoin // 过亿
	Analy.TimeUnix = opt.TickerVol[0].TimeUnix
	Analy.TimeStr = mTime.UnixFormat(Analy.TimeUnix)
	Analy.TimeID = mOKX.GetTimeID(Analy.TimeUnix)

	global.Run.Println(
		"== 分析结束 ==",
		len(Analy.AnalyWhole),
		len(Analy.AnalySingle),
		config.Unit,
		Analy.TimeID,
	)

	return Analy
}
