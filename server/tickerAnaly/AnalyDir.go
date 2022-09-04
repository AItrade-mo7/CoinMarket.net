package tickerAnaly

import (
	"CoinMarket.net/server/okxInfo"
	"github.com/EasyGolang/goTools/mCount"
	"github.com/EasyGolang/goTools/mStr"
)

func AnalyDir() {
	// 初始化为空值
	okxInfo.WholeDir = 0

	upDir := []string{}
	downDir := []string{}
	zeroDir := []string{}
	allDir := []string{}

	for key, item := range okxInfo.TickerAnalyWhole {
		fade := len(okxInfo.TickerAnalyWhole) - key
		fadeStr := mStr.ToStr(fade)
		if item.DirIndex > 0 {
			upDir = append(upDir, fadeStr)
		}
		if item.DirIndex == 0 {
			zeroDir = append(zeroDir, fadeStr)
		}
		if item.DirIndex < 0 {
			downDir = append(downDir, fadeStr)
		}

		allDir = append(allDir, fadeStr)

	}

	upFade := mCount.Sum(upDir)
	downFade := mCount.Sum(downDir)
	zeroFade := mCount.Sum(zeroDir)
	allFade := mCount.Sum(allDir)

	zeroPer := mCount.PerCent(zeroFade, allFade)

	okxInfo.WholeDir = mCount.Le(upFade, downFade)

	if mCount.Le(zeroPer, "50") > 0 {
		okxInfo.WholeDir = okxInfo.WholeDir * 2
	}
}
