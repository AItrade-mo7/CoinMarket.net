package dbType

import (
	"CoinMarket.net/server/okxInfo"
	"github.com/EasyGolang/goTools/mOKX"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mTime"
)

type CoinTickerTable struct {
	TickerVol []mOKX.TypeTicker        `bson:"TickerVol"` // 成交量排序
	Kdata     map[string][]mOKX.TypeKd `bson:"Kdata"`     //
	TimeUnix  int64                    `bson:"TimeUnix"`
	TimeStr   string                   `bson:"Time"`
}

func JoinCoinTicker() CoinTickerTable {
	var CoinTicker CoinTickerTable
	CoinTicker.TickerVol = okxInfo.TickerList
	CoinTicker.Kdata = okxInfo.MarketKdata
	CoinTicker.TimeUnix = CoinTicker.TickerVol[0].Ts
	CoinTicker.TimeStr = mTime.UnixFormat(mStr.ToStr(CoinTicker.TimeUnix))

	return CoinTicker
}
