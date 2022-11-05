package ready

import (
	"time"

	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/okxApi/binanceApi"
	"CoinMarket.net/server/okxApi/restApi/kdata"
	"CoinMarket.net/server/okxInfo"
	"github.com/EasyGolang/goTools/mCount"
	"github.com/EasyGolang/goTools/mOKX"
)

func SetTickerKdata() {
	TickerKdata := make(map[string][]mOKX.TypeKd)
	TickerList := []mOKX.TypeTicker{}
	for _, item := range okxInfo.TickerVol {
		time.Sleep(time.Second / 2) // 1秒最多 2 次
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

		BinanceList := binanceApi.GetKdata(binanceApi.GetKdataParam{
			Symbol: item.Symbol,
			Size:   config.KdataLen,
		})
		OKXList := kdata.GetKdata(item.InstID, config.KdataLen)

		List := KdataMerge(DataMergeOpt{
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

type DataMergeOpt struct {
	OKXList     []mOKX.TypeKd
	BinanceList []mOKX.TypeKd
}

func KdataMerge(opt DataMergeOpt) []mOKX.TypeKd {
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
