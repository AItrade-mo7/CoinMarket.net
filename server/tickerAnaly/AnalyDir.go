package tickerAnaly

import (
	"github.com/EasyGolang/goTools/mCount"
	"github.com/EasyGolang/goTools/mOKX"
	"github.com/EasyGolang/goTools/mStr"
)

func AnalyDir(TickerAnalyWhole []mOKX.TypeWholeTickerAnaly) int {
	// 初始化为空值
	WholeDir := 0

	upDir := []string{}
	downDir := []string{}
	zeroDir := []string{}
	allDir := []string{}

	for key, item := range TickerAnalyWhole {
		fade := len(TickerAnalyWhole) - key
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

	WholeDir = mCount.Le(upFade, downFade)

	if mCount.Le(zeroPer, "50") > 0 {
		WholeDir = WholeDir * 2
	}

	return WholeDir
}
