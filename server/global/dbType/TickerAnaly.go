package dbType

import (
	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/tickerAnaly"
	"github.com/EasyGolang/goTools/mOKX"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mTime"
)

type AnalyTickerType struct {
	TickerVol   []mOKX.TypeTicker                `bson:"TickerVol"`   // 列表
	AnalyWhole  []mOKX.TypeWholeTickerAnaly      `bson:"AnalyWhole"`  // 大盘分析结果
	AnalySingle map[string][]mOKX.AnalySliceType `bson:"AnalySingle"` // 单个币种分析结果
	Unit        string                           `bson:"Unit"`
	WholeDir    int                              `bson:"WholeDir"`
	TimeUnix    int64                            `bson:"TimeUnix"`
	TimeStr     string                           `bson:"TimeStr"`
	TimeID      string                           `bson:"TimeID"`
}

func GetAnalyTicker(opt tickerAnaly.TickerAnalyParam) (resData AnalyTickerType) {
	resData = AnalyTickerType{}

	if len(opt.TickerVol) > 5 && len(opt.TickerKdata) > 5 && len(opt.TickerKdata["BTC-USDT"]) > 5 {
	} else {
		global.LogErr("dbType.GetAnalyTicker", len(opt.TickerVol), len(opt.TickerKdata))
		return
	}

	AnalyResult := tickerAnaly.GetAnaly(opt)

	resData.TickerVol = opt.TickerVol
	resData.AnalyWhole = AnalyResult.AnalyWhole
	resData.AnalySingle = AnalyResult.AnalySingle
	resData.WholeDir = AnalyResult.WholeDir
	resData.Unit = config.Unit
	resData.TimeUnix = mOKX.GetKdataTime(opt.TickerKdata)
	resData.TimeStr = mTime.UnixFormat(mStr.ToStr(resData.TimeUnix))
	resData.TimeID = mOKX.GetTimeID(resData.TimeUnix)

	return
}
