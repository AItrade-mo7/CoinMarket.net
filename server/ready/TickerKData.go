package ready

import (
	"time"

	"CoinMarket.net/server/okxApi/restApi/kdata"
	"CoinMarket.net/server/okxInfo"
)

func TickerKdata() {
	for _, item := range okxInfo.TickerList {
		time.Sleep(time.Second / 5) // 1秒最多 5 次
		list := kdata.GetKdata(item.InstID)
		okxInfo.MarketKdata[item.InstID] = list
	}
}
