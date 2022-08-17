package tickerAnaly

import (
	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/okxInfo"
	"github.com/EasyGolang/goTools/mOKX"
)

func Start() {
	// 基于  mOKX.MarketKdata  进行数据分析
	TickerAnalySingle := make(map[string][]mOKX.AnalySliceType)
	for _, list := range okxInfo.MarketKdata {
		Single := NewSingle(list)
		if len(Single.ResData) == len(config.SliceHour) {
			TickerAnalySingle[Single.Info.InstID] = Single.ResData
		} else {
			global.LogErr("tickerAnaly.Start  数据长度不足", Single.Info.InstID, len(Single.ResData))
		}
	}

	okxInfo.TickerAnalySingle = TickerAnalySingle

	WholeAnaly()

	AnalyDir()
}
