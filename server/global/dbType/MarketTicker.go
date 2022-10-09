package dbType

import (
	"github.com/EasyGolang/goTools/mOKX"
)

type MarketTickerTable struct {
	List           []mOKX.TypeTicker                `bson:"List"`        // 成交量排序列表
	ListU_R24      []mOKX.TypeTicker                `bson:"ListU_R24"`   // 涨跌幅排序列表
	AnalyWhole     []mOKX.TypeWholeTickerAnaly      `bson:"AnalyWhole"`  // 大盘分析结果
	AnalySingle    map[string][]mOKX.AnalySliceType `bson:"AnalySingle"` // 单个币种分析结果
	Unit           string                           `bson:"Unit"`
	WholeDir       int                              `bson:"WholeDir"`
	TimeUnix       int64                            `bson:"TimeUnix"`
	Time           string                           `bson:"Time"`
	CreateTimeUnix int64                            `bson:"CreateTimeUnix"`
	CreateTime     string                           `bson:"CreateTime"`
}
