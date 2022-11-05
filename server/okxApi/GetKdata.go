package okxApi

import (
	"fmt"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/okxApi/binanceApi"
	"CoinMarket.net/server/okxApi/restApi/kdata"
	"CoinMarket.net/server/okxInfo"
	"github.com/EasyGolang/goTools/mCount"
	"github.com/EasyGolang/goTools/mOKX"
	"github.com/EasyGolang/goTools/mTime"
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

	nowUnix := mTime.GetUnixInt64() - mTime.UnixTimeInt64.Minute*15
	if opt.After > nowUnix {
		opt.After = 0 // 当前
	} else {
		// 历史
		if opt.Size > 100 {
			opt.Size = 100
		}
	}

	if opt.Size < config.KdataLen {
		opt.Size = config.KdataLen
	}

	BinanceList := binanceApi.GetKdata(binanceApi.GetKdataParam{
		Symbol:  SPOT.Symbol,
		Current: opt.Current,
		After:   opt.After,
		Size:    opt.Size,
	})

	var OKXList []mOKX.TypeKd
	if (opt.After) > 0 {
		OKXList = kdata.GetHistoryKdata(kdata.HistoryKdataParam{
			InstID:  SPOT.InstID,
			Current: opt.Current,
			After:   opt.After,
			Size:    opt.Size,
		})
	} else {
		OKXList = kdata.GetKdata(SPOT.InstID, 100)
	}

	List, err := KdataMerge(KdataMergeOpt{
		OKXList:     OKXList,
		BinanceList: BinanceList,
	})
	if err != nil {
		global.LogErr(err)
		return
	}

	KdataList = List

	return
}

type KdataMergeOpt struct {
	OKXList     []mOKX.TypeKd
	BinanceList []mOKX.TypeKd
}

func KdataMerge(opt KdataMergeOpt) (Kdata []mOKX.TypeKd, resErr error) {
	OKXList := opt.OKXList
	BinanceList := opt.BinanceList
	Kdata = []mOKX.TypeKd{}
	resErr = nil

	if len(OKXList) != len(BinanceList) {
		resErr = fmt.Errorf("okxApi.KdataMerge len %+v %+v", len(OKXList), len(BinanceList))
		return
	}

	if OKXList[0].TimeStr != BinanceList[0].TimeStr {
		resErr = fmt.Errorf("okxApi.KdataMerge [0] %+v %+v", OKXList[0].TimeStr, BinanceList[0].TimeStr)
		return
	}

	if OKXList[48].TimeStr != BinanceList[48].TimeStr {
		resErr = fmt.Errorf("okxApi.KdataMerge [48] %+v %+v", OKXList[48].TimeStr, BinanceList[48].TimeStr)
		return
	}

	if OKXList[len(OKXList)-1].TimeStr != BinanceList[len(BinanceList)-1].TimeStr {
		resErr = fmt.Errorf("okxApi.KdataMerge [last] %+v %+v", OKXList[len(OKXList)-1].TimeStr, BinanceList[len(BinanceList)-1].TimeStr)
		return
	}

	for _, item := range OKXList {
		OkxItem := item

		for _, BinanceItem := range BinanceList {
			if OkxItem.TimeUnix == BinanceItem.TimeUnix {
				VolCcy := mCount.Add(BinanceItem.VolCcy, OkxItem.VolCcy)
				OkxItem.VolCcy = VolCcy
				Vol := mCount.Add(BinanceItem.Vol, OkxItem.Vol)
				OkxItem.Vol = Vol
				OkxItem.DataType = "Merge"
				break
			}
		}

		Kdata = append(Kdata, OkxItem)
	}

	return
}
