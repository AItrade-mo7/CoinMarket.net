package dbType

import (
	"strings"

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

func JoinCoinTicker(TickerList []mOKX.TypeTicker, KdataList map[string][]mOKX.TypeKd) CoinTickerTable {
	var CoinTicker CoinTickerTable
	CoinTicker.TickerVol = TickerList
	CoinTicker.Kdata = make(map[string][]mOKX.TypeKd)

	for key, val := range KdataList {
		find := strings.Contains(key, "-SWAP")
		if !find {
			CoinTicker.Kdata[key] = val
		}
	}

	CoinTicker.TimeUnix = CoinTicker.TickerVol[0].Ts
	CoinTicker.TimeStr = mTime.UnixFormat(mStr.ToStr(CoinTicker.TimeUnix))
	CoinTicker.TimeID = mOKX.GetTimeID(CoinTicker.TimeUnix)

	return CoinTicker
}
