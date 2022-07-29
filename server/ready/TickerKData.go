package ready

import (
	"time"

	"CoinMarket.net/server/okxApi/okxInfo"
	"CoinMarket.net/server/okxApi/restApi/kdata"
)

func TickerKdata() {
	for _, item := range okxInfo.TickerList {
		time.Sleep(time.Second / 6) // 1秒最多 6 次
		list := kdata.GetKdata(item.InstID)
		okxInfo.MarketKdata[item.InstID] = list
	}

	AnalyseTicker_single(okxInfo.MarketKdata["ETH-USDT"])
}
