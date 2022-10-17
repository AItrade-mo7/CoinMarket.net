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

func AnalyDir(opt AnalyDirParam) int {
	// 初始化为空值
	WholeDir := 0
	if len(opt.AnalyWhole) > 0 {
		WholeDir = opt.AnalyWhole[0].DirIndex
	}

	// 筛选出过亿的币种
	millionCoin := []string{}
	for _, ticker := range opt.TickerVol {
		if mCount.Le(ticker.Volume, million) >= 0 {
			millionCoin = append(millionCoin, ticker.InstID)
		}
	}
	// 过亿币种数量小于4 ， 交易量小的上涨不作数的
	if len(millionCoin) < 4 {
		// 如果此时判断为涨，修订为0
		if WholeDir > 0 {
			WholeDir = 0
		}
	}

	return WholeDir
}
