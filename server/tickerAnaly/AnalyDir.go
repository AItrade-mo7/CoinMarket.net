package tickerAnaly

import (
	"github.com/EasyGolang/goTools/mCount"
	"github.com/EasyGolang/goTools/mOKX"
)

var million = "100000000"

type AnalyDirParam struct {
	TickerVol  []mOKX.TypeTicker
	AnalyWhole []mOKX.TypeWholeTickerAnaly
}

func AnalyDir(opt AnalyDirParam) (MillionCoin []string, DirIndex int) {
	// 初始化为空值
	DirIndex = 0
	if len(opt.AnalyWhole) > 0 {
		DirIndex = opt.AnalyWhole[0].DirIndex
	}

	// 筛选出过亿的币种
	MillionCoin = []string{}
	for _, ticker := range opt.TickerVol {
		if mCount.Le(ticker.Volume, million) >= 0 {
			MillionCoin = append(MillionCoin, ticker.InstID)
		}
	}

	return
}
