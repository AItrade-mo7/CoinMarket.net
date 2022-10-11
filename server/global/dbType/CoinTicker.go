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
	TimeStr   string                   `bson:"TimeStr"`
	TimeID    string                   `bson:"TimeID"`
}

func JoinCoinTicker() CoinTickerTable {
	var CoinTicker CoinTickerTable
	CoinTicker.TickerVol = okxInfo.TickerList
	CoinTicker.Kdata = okxInfo.MarketKdata
	CoinTicker.TimeUnix = CoinTicker.TickerVol[0].Ts
	CoinTicker.TimeStr = mTime.UnixFormat(mStr.ToStr(CoinTicker.TimeUnix))

	T := mTime.MsToTime(CoinTicker.TimeUnix, "0")
	timeStr := T.Format("2006-01-02T15:04")
	CoinTicker.TimeID = timeStr

	return CoinTicker
}
