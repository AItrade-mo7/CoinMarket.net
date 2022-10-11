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
	okxInfo.TickerKdata = make(map[string][]mOKX.TypeKd)
	TickerList := []mOKX.TypeTicker{}
	for _, item := range okxInfo.TickerList {
		time.Sleep(time.Second / 3) // 1秒最多 3 次
		// 开始设置 SWAP
		SwapInst := mOKX.TypeInst{}
		for _, SWAP := range okxInfo.SWAP_inst {
			if SWAP.Uly == item.InstID {
				SwapInst = SWAP
				break
			}
		}
		if len(SwapInst.InstID) < 3 {
			continue
		}

		OKXList := kdata.GetKdata(item.InstID)
		BinanceList := binanceApi.GetKdata(item.Symbol)

		SWAPList := kdata.GetKdata(SwapInst.InstID)

		List := DataMerge(DataMergeOpt{
			OKXList:     OKXList,
			BinanceList: BinanceList,
		})

		if len(List) == 300 {
			TickerList = append(TickerList, item)
			okxInfo.TickerKdata[item.InstID] = List
			okxInfo.TickerKdata[SwapInst.InstID] = SWAPList
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
