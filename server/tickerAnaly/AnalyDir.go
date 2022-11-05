package tickerAnaly

import (
	"github.com/EasyGolang/goTools/mCount"
	"github.com/EasyGolang/goTools/mOKX"
)

var million = "100000000"

type AnalyDirParam struct {
	TickerVol   []mOKX.TypeTicker
	AnalyWhole  []mOKX.TypeWholeTickerAnaly
	AnalySingle map[string][]mOKX.AnalySliceType
}

func AnalyDir(opt AnalyDirParam) (MillionCoin []mOKX.AnalySliceType) {
	// 初始化为空值

	// 筛选出过亿的币种
	MillionCoin = []mOKX.AnalySliceType{}
	for _, ticker := range opt.TickerVol {
		if mCount.Le(ticker.Volume, million) >= 0 {
			Single := opt.AnalySingle[ticker.InstID]
			if len(Single) > 0 {
				MillionCoin = append(MillionCoin, Single[0])
			}
		}
	}

	return
}
