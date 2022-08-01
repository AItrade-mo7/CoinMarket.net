package ready

import (
	"time"

	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/okxApi/restApi/kdata"
	"CoinMarket.net/server/okxInfo"
	"CoinMarket.net/server/tickerAnalyse"
)

func TickerKdata() {
	if config.AppEnv.RunMod == 1 {
		item := okxInfo.TickerList[0]
		list := kdata.GetKdata(item.InstID)
		tickerAnalyse.NewSingle(list)
		return
	}

	for _, item := range okxInfo.TickerList {
		time.Sleep(time.Second / 5) // 1秒最多 5 次
		list := kdata.GetKdata(item.InstID)
		tickerAnalyse.NewSingle(list)
		okxInfo.MarketKdata[item.InstID] = list
	}
}
