package dbType

import (
	"strings"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/tickerAnaly"
	"github.com/EasyGolang/goTools/mOKX"
	"github.com/EasyGolang/goTools/mTime"
)

type CoinTickerTable struct {
	TickerVol []mOKX.TypeTicker        `bson:"TickerVol"` // 成交量排序
	Kdata     map[string][]mOKX.TypeKd `bson:"Kdata"`     //
	TimeUnix  int64                    `bson:"TimeUnix"`
	TimeStr   string                   `bson:"TimeStr"`
	TimeID    string                   `bson:"TimeID"`
}

func JoinCoinTicker(opt tickerAnaly.TickerAnalyParam) CoinTickerTable {
	if len(opt.TickerVol) > 3 && len(opt.TickerKdata) == len(opt.TickerVol) && len(opt.TickerKdata["BTC-USDT"]) == config.KdataLen {
	} else {
		global.LogErr("dbType.GetAnalyTicker", len(opt.TickerVol), len(opt.TickerKdata))
		return CoinTickerTable{}
	}

	var CoinTicker CoinTickerTable
	CoinTicker.TickerVol = opt.TickerVol
	CoinTicker.Kdata = make(map[string][]mOKX.TypeKd)

	for key, val := range opt.TickerKdata {
		find := strings.Contains(key, "-SWAP")
		if !find {
			if len(val) >= config.KdataLen {
				CoinTicker.Kdata[key] = val
			}
		}
	}

	CoinTicker.TimeUnix = opt.TickerVol[0].TimeUnix
	CoinTicker.TimeStr = mTime.UnixFormat(CoinTicker.TimeUnix)
	CoinTicker.TimeID = mOKX.GetTimeID(CoinTicker.TimeUnix)

	return CoinTicker
}
