package sort

import (
	"CoinMarket.net/server/okxInfo"
	"github.com/EasyGolang/goTools/mCount"
)

// 按照最高价排序
func H_sort(data []okxInfo.Kd) []okxInfo.Kd {
	size := len(data)
	list := make([]okxInfo.Kd, size)
	copy(list, data)

	var swapped bool
	for i := size - 1; i > 0; i-- {
		swapped = false
		for j := 0; j < i; j++ {
			a := list[j+1].H
			b := list[j].H
			if mCount.Le(a, b) < 0 {
				list[j], list[j+1] = list[j+1], list[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	// 设置 Idx 并翻转
	listIDX := []okxInfo.Kd{}
	j := 0
	for i := len(list) - 1; i > -1; i-- {
		Ticker := list[i]
		Ticker.H_idx = j + 1
		listIDX = append(listIDX, Ticker)
		j++
	}
	return listIDX
}
