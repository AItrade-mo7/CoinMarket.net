package ready

import (
	"time"

	"CoinMarket.net/server/okxApi/okxInfo"
	"CoinMarket.net/server/okxApi/restApi/kdata"
)

func TickerKdata() {
	MaxNum := len(okxInfo.TickerList) / 2 // 去除前 6 条数据
	for key, item := range okxInfo.TickerList {
		time.Sleep(time.Second / 6) // 1秒最多 6 次
		list := kdata.GetKdata(item.InstID)
		okxInfo.MarketKdata[item.InstID] = list
		if key > MaxNum-2 {
			break
		}
	}
}
