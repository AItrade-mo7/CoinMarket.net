package tickerAnaly

import (
	"CoinMarket.net/server/okxInfo"
	"github.com/EasyGolang/goTools/mOKX"
)

func Start() {
	// 基于  mOKX.MarketKdata  进行数据分析
	okxInfo.TickerAnalySingle = make(map[string][]mOKX.AnalySliceType)
	for _, list := range okxInfo.MarketKdata {
		Single := NewSingle(list)
		okxInfo.TickerAnalySingle[Single.Info.InstID] = Single.ResData
	}

	WholeAnaly()
}
