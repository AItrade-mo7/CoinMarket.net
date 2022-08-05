package ready

import (
	"time"

	"CoinMarket.net/server/okxApi/binanceApi"
	"CoinMarket.net/server/okxApi/restApi/kdata"
	"CoinMarket.net/server/okxInfo"
	"github.com/EasyGolang/goTools/mCount"
	"github.com/EasyGolang/goTools/mOKX"
)

func TickerKdata() {
	for _, item := range okxInfo.TickerList {
		time.Sleep(time.Second / 5) // 1秒最多 5 次
		List := DataMerge(DataMergeOpt{
			OKXList:     kdata.GetKdata(item.InstID),
			BinanceList: binanceApi.GetKdata(item.Symbol),
		})
		okxInfo.MarketKdata[item.InstID] = List
	}
}

type DataMergeOpt struct {
	OKXList     []mOKX.TypeKd
	BinanceList []mOKX.TypeKd
}

func DataMerge(opt DataMergeOpt) []mOKX.TypeKd {
	OKXList := opt.OKXList
	BinanceList := opt.BinanceList
	Kdata := []mOKX.TypeKd{}
	for key, item := range OKXList {
		OkxItem := item
		BinanceItem := BinanceList[key]
		if OkxItem.TimeUnix == BinanceItem.TimeUnix {
			VolCcy := mCount.Add(BinanceItem.VolCcy, OkxItem.VolCcy)
			OkxItem.VolCcy = VolCcy
			Vol := mCount.Add(BinanceItem.Vol, OkxItem.Vol)
			OkxItem.Vol = Vol
			Kdata = append(Kdata, OkxItem)
		} else {
			Kdata = []mOKX.TypeKd{} // 在这里清除
			break
		}
	}
	return Kdata
}
