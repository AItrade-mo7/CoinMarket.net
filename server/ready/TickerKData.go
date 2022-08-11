package ready

import (
	"time"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/okxApi/binanceApi"
	"CoinMarket.net/server/okxApi/restApi/kdata"
	"CoinMarket.net/server/okxInfo"
	"github.com/EasyGolang/goTools/mCount"
	"github.com/EasyGolang/goTools/mOKX"
)

func TickerKdata() {
	okxInfo.MarketKdata = make(map[string][]mOKX.TypeKd)
	TickerList := []mOKX.TypeTicker{}
	for _, item := range okxInfo.TickerList {
		time.Sleep(time.Second) // 1秒最多 1 次

		OKXList := kdata.GetKdata(item.InstID)
		BinanceList := binanceApi.GetKdata(item.Symbol)

		List := DataMerge(DataMergeOpt{
			OKXList:     OKXList,
			BinanceList: BinanceList,
		})

		if len(List) == 300 {
			TickerList = append(TickerList, item)
			okxInfo.MarketKdata[item.InstID] = List
		} else {
			global.LogErr("ready.TickerKdata", "长度不正确", item.InstID, len(List), len(OKXList), len(BinanceList))
		}
	}
	okxInfo.TickerList = make([]mOKX.TypeTicker, len(TickerList))
	okxInfo.TickerList = TickerList
}

type DataMergeOpt struct {
	OKXList     []mOKX.TypeKd
	BinanceList []mOKX.TypeKd
}

func DataMerge(opt DataMergeOpt) []mOKX.TypeKd {
	OKXList := opt.OKXList
	BinanceList := opt.BinanceList
	Kdata := []mOKX.TypeKd{}
	for _, item := range OKXList {
		OkxItem := item

		for _, BinanceItem := range BinanceList {
			if OkxItem.TimeUnix == BinanceItem.TimeUnix {
				VolCcy := mCount.Add(BinanceItem.VolCcy, OkxItem.VolCcy)
				OkxItem.VolCcy = VolCcy
				Vol := mCount.Add(BinanceItem.Vol, OkxItem.Vol)
				OkxItem.Vol = Vol
				break
			}
		}

		Kdata = append(Kdata, OkxItem)
	}
	return Kdata
}
