package dbType

import (
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/okxInfo"
	"github.com/EasyGolang/goTools/mOKX"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mTime"
)

type MarketTickerTable struct {
	List        []mOKX.TypeTicker                `bson:"List"`        // 成交量排序列表
	ListU_R24   []mOKX.TypeTicker                `bson:"ListU_R24"`   // 涨跌幅排序列表
	AnalyWhole  []mOKX.TypeWholeTickerAnaly      `bson:"AnalyWhole"`  // 大盘分析结果
	AnalySingle map[string][]mOKX.AnalySliceType `bson:"AnalySingle"` // 单个币种分析结果
	Unit        string                           `bson:"Unit"`
	WholeDir    int                              `bson:"WholeDir"`
	TimeUnix    int64                            `bson:"TimeUnix"`
	Date        string                           `bson:"Date"`
}

// 拼接数据
func GetTickerDB() MarketTickerTable {
	TickerRes := MarketTickerTable{}
	TickerRes.List = okxInfo.TickerList
	TickerRes.ListU_R24 = okxInfo.TickerList
	TickerRes.AnalyWhole = okxInfo.TickerAnalyWhole
	TickerRes.AnalySingle = okxInfo.TickerAnalySingle
	TickerRes.Unit = config.Unit
	TickerRes.WholeDir = okxInfo.WholeDir
	TickerRes.TimeUnix = okxInfo.TickerList[0].Ts

	TickerDate := mTime.UnixFormat(mStr.ToStr(TickerRes.TimeUnix))
	TickerRes.Date = TickerDate

	return TickerRes
}
