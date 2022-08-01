package tickerAnalyse

import (
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/okxInfo"
)

var Single map[string]SingleType

func Start() {
	// 基于  okxInfo.TickerList  进行数据分析
	WholeAnalyse()

	// 基于  okxInfo.MarketKdata  进行数据分析
	Single = make(map[string]SingleType)

	if config.AppEnv.RunMod == 1 {
		list := okxInfo.MarketKdata["ETC-USDT"]
		NewSingle(list)
		return
	}

	for _, list := range okxInfo.MarketKdata {
		NewSingle(list)
	}
}
