package dbType

import (
	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/tickerAnaly"
	"github.com/EasyGolang/goTools/mOKX"
)

type AnalyTickerType struct {
	TickerVol   []mOKX.TypeTicker                `bson:"TickerVol"`   // 列表
	AnalyWhole  []mOKX.TypeWholeTickerAnaly      `bson:"AnalyWhole"`  // 大盘分析结果
	AnalySingle map[string][]mOKX.AnalySliceType `bson:"AnalySingle"` // 单个币种分析结果
	MillionCoin []mOKX.AnalySliceType            `bson:"MillionCoin"`
	Version     int                              `bson:"Version"`
	Unit        string                           `bson:"Unit"`
	TimeUnix    int64                            `bson:"TimeUnix"`
	TimeStr     string                           `bson:"TimeStr"`
	TimeID      string                           `bson:"TimeID"`
}

func GetAnalyTicker(opt tickerAnaly.TickerAnalyParam) (resData AnalyTickerType) {
	resData = AnalyTickerType{}
	resData.Version = 3

	if len(opt.TickerVol) > 3 && len(opt.TickerKdata) == len(opt.TickerVol) && len(opt.TickerKdata["BTC-USDT"]) >= config.KdataLen {
	} else {
		global.LogErr("dbType.GetAnalyTicker", len(opt.TickerVol), len(opt.TickerKdata), len(opt.TickerKdata["BTC-USDT"]))
		return
	}

	AnalyResult := tickerAnaly.GetAnaly(opt)

	resData.TickerVol = opt.TickerVol
	resData.AnalyWhole = AnalyResult.AnalyWhole
	resData.AnalySingle = AnalyResult.AnalySingle
	resData.Unit = config.Unit
	resData.MillionCoin = AnalyResult.MillionCoin
	resData.TimeUnix = AnalyResult.TimeUnix
	resData.TimeStr = AnalyResult.TimeStr
	resData.TimeID = AnalyResult.TimeID

	return
}
