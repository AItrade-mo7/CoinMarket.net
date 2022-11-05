package ready

import (
	"time"

	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/okxApi"
	"CoinMarket.net/server/okxInfo"
	"github.com/EasyGolang/goTools/mOKX"
)

func SetTickerNowKdata() {
	TickerKdata := make(map[string][]mOKX.TypeKd)
	TickerList := []mOKX.TypeTicker{}
	for _, item := range okxInfo.TickerVol {
		time.Sleep(time.Second / 2) // 1秒最多 2 次

		List := okxApi.GetKdata(okxApi.GetKdataOpt{
			InstID: item.InstID,
			Size:   config.KdataLen,
		})

		if len(List) >= config.KdataLen {
			TickerList = append(TickerList, item)
			TickerKdata[item.InstID] = List
		}

	}

	okxInfo.TickerKdata = make(map[string][]mOKX.TypeKd)
	okxInfo.TickerKdata = TickerKdata

	okxInfo.TickerVol = []mOKX.TypeTicker{}
	okxInfo.TickerVol = TickerList
}
