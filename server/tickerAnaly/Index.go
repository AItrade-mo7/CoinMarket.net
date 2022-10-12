package tickerAnaly

import (
	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"github.com/EasyGolang/goTools/mOKX"
)

type TickerAnalyParam struct {
	TickerList  []mOKX.TypeTicker
	MarketKdata map[string][]mOKX.TypeKd
}

type AnalyResult struct {
	AnalySingle map[string][]mOKX.AnalySliceType
	AnalyWhole  []mOKX.TypeWholeTickerAnaly
	WholeDir    int
}

func GetAnaly(opt TickerAnalyParam) AnalyResult {
	// 基于  mOKX.MarketKdata  进行数据分析
	TickerAnalySingle := make(map[string][]mOKX.AnalySliceType)

	for _, item := range opt.TickerList {
		list := opt.MarketKdata[item.InstID]

		Single := NewSingle(list)
		if len(Single.ResData) == len(config.SliceHour) {
			TickerAnalySingle[Single.Info.InstID] = Single.ResData
		} else {
			global.LogErr("tickerAnaly.Start  数据长度不足", Single.Info.InstID, len(Single.ResData))
		}
	}

	TickerAnalyWhole := WholeAnaly(TickerAnalySingle)
	WholeDir := AnalyDir(TickerAnalyWhole)

	var Analy AnalyResult
	Analy.AnalySingle = TickerAnalySingle
	Analy.AnalyWhole = TickerAnalyWhole
	Analy.WholeDir = WholeDir

	return Analy
}
