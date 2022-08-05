package tickerAnaly

import (
	"CoinMarket.net/server/okxInfo"
	"github.com/EasyGolang/goTools/mOKX"
)

func Start() {
	// 基于  mOKX.TickerList  进行数据分析
	okxInfo.TickerAnalyWhole = mOKX.TypeWholeTickerAnaly{}
	WholeAnaly()

	// 基于  mOKX.MarketKdata  进行数据分析
	okxInfo.TickerAnalySingle = map[string][]mOKX.AnalySliceType{}

	// if config.AppEnv.RunMod == 1 {
	// 	list := okxInfo.MarketKdata["ETH-USDT"]
	// 	NewSingle(list)
	// 	return
	// }
	for _, list := range okxInfo.MarketKdata {
		NewSingle(list)
	}
}
