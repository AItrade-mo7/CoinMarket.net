package tickers

import (
	"github.com/EasyGolang/goTools/mCount"
)

// 成交量排序
func VolumeSort(arr []OKXTickerType) []OKXTickerType {
	size := len(arr)
	list := make([]OKXTickerType, size)
	copy(list, arr)

	var swapped bool
	for i := size - 1; i > 0; i-- {
		swapped = false
		for j := 0; j < i; j++ {
			a := list[j+1].VolCcy24H
			b := list[j].VolCcy24H
			if mCount.Le(a, b) < 0 {
				list[j], list[j+1] = list[j+1], list[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return list
}

func Reverse(arr []OKXTickerType) []OKXTickerType {
	list := make(
		[]OKXTickerType,
		len(arr),
		len(arr)*2,
	)

	j := 0
	for i := len(arr) - 1; i > -1; i-- {
		list[j] = arr[i]
		j++
	}

	return list
}
