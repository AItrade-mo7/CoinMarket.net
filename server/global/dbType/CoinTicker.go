package dbType

import "github.com/EasyGolang/goTools/mOKX"

type CoinTickerTable struct {
	TickerVol []mOKX.TypeTicker        `bson:"TickerVol"` // 成交量排序
	Kdata     map[string][]mOKX.TypeKd `bson:"Kdata"`     //
	TimeUnix  int64                    `bson:"TimeUnix"`
	TimeStr   string                   `bson:"Time"`
}
