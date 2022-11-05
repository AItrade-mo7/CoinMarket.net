package okxApi

import (
	"fmt"

	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/okxApi/binanceApi"
	"CoinMarket.net/server/okxApi/restApi/kdata"
	"CoinMarket.net/server/okxInfo"
	"github.com/EasyGolang/goTools/mCount"
	"github.com/EasyGolang/goTools/mOKX"
)

type GetKdataOpt struct {
	InstID  string `bson:"InstID"`
	Current int    `bson:"Current"` // 当前页码 0 为
	After   int64  `bson:"After"`   // 时间 默认为当前时间
	Size    int    `bson:"Size"`    // 数量 默认为100
}

func GetKdata(opt GetKdataOpt) (KdataList []mOKX.TypeKd) {
	KdataList = []mOKX.TypeKd{}
	SPOT := okxInfo.Inst[opt.InstID]
	if len(SPOT.InstID) < 3 {
		return
	}

	BinanceList := binanceApi.GetKdata(binanceApi.GetKdataParam{
		Symbol:  SPOT.Symbol,
		Current: opt.Current,
		After:   opt.After,
		Size:    config.KdataLen,
	})
	OKXList := []mOKX.TypeKd{}
	if opt.After == 0 {
		OKXList = kdata.GetKdata(SPOT.InstID, config.KdataLen)
	}

	List := KdataMerge(KdataMergeOpt{
		OKXList:     OKXList,
		BinanceList: BinanceList,
	})

	fmt.Println(len(List))

	return
}

type KdataMergeOpt struct {
	OKXList     []mOKX.TypeKd
	BinanceList []mOKX.TypeKd
}

func KdataMerge(opt KdataMergeOpt) []mOKX.TypeKd {
	OKXList := opt.OKXList
	BinanceList := opt.BinanceList
	Kdata := []mOKX.TypeKd{}

	fmt.Println(len(OKXList), len(BinanceList))
	fmt.Println(OKXList[0].TimeStr, OKXList[0].TimeStr == BinanceList[0].TimeStr)
	fmt.Println(OKXList[48].TimeStr, OKXList[48].TimeStr == BinanceList[48].TimeStr)
	fmt.Println(OKXList[86].TimeStr, OKXList[86].TimeStr == BinanceList[86].TimeStr)
	fmt.Println(OKXList[config.KdataLen-1].TimeStr, OKXList[config.KdataLen-1].TimeStr == BinanceList[config.KdataLen-1].TimeStr)

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
