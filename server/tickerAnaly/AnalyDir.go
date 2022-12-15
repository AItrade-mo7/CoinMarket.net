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
				slices := Single[0]
				for _, SingleItem := range Single {
					if SingleItem.DiffHour > 6 {
						slices = SingleItem
						break
					}
				}
				MillionCoin = append(MillionCoin, slices)
			}
		}
	}

	return
}
