package ready

import (
	"fmt"
	"time"

	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/okxApi"
	"CoinMarket.net/server/okxApi/binanceApi"
	"CoinMarket.net/server/okxApi/restApi/kdata"
	"CoinMarket.net/server/okxInfo"
	"github.com/EasyGolang/goTools/mOKX"
)

func SetTickerNowKdata() {
	TickerKdata := make(map[string][]mOKX.TypeKd)
	TickerList := []mOKX.TypeTicker{}
	for _, item := range okxInfo.TickerVol {
		time.Sleep(time.Second / 2) // 1秒最多 2 次

		fmt.Println(item.Symbol,item.InstID)

		BinanceList := binanceApi.GetKdata(binanceApi.GetKdataParam{
			Symbol: item.Symbol,
			Size:   config.KdataLen,
		})
		OKXList := kdata.GetKdata(item.InstID, config.KdataLen)

		List := okxApi.KdataMerge(okxApi.KdataMergeOpt{
			OKXList:     OKXList,
			BinanceList: BinanceList,
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
