package sort

import (
	"github.com/EasyGolang/goTools/mCount"
	"github.com/EasyGolang/goTools/mOKX"
)

// 成交量排序
func Volume(data []mOKX.TickerType) []mOKX.TickerType {
	size := len(data)
	list := make([]mOKX.TickerType, size)
	copy(list, data)

	var swapped bool
	for i := size - 1; i > 0; i-- {
		swapped = false
		for j := 0; j < i; j++ {
			a := list[j+1].Volume
			b := list[j].Volume
			if mCount.Le(a, b) < 0 {
				list[j], list[j+1] = list[j+1], list[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}

	// 设置 VolIdx 并翻转

	listIDX := []mOKX.TickerType{}
	j := 0
	for i := len(list) - 1; i > -1; i-- {
		Ticker := list[i]
		Ticker.VolIdx = j + 1
		listIDX = append(listIDX, Ticker)
		j++
	}

	return listIDX
}
