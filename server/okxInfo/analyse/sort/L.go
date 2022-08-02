package sort

import (
	"github.com/EasyGolang/goTools/mCount"
	"github.com/EasyGolang/goTools/mOKX"
)

// 按照 最低 价排序
func L_sort(data []mOKX.Kd) []mOKX.Kd {
	size := len(data)
	list := make([]mOKX.Kd, size)
	copy(list, data)

	var swapped bool
	for i := size - 1; i > 0; i-- {
		swapped = false
		for j := 0; j < i; j++ {
			a := list[j+1].L
			b := list[j].L
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
	listIDX := []mOKX.Kd{}
	j := 0
	for i := len(list) - 1; i > -1; i-- {
		Ticker := list[i]
		Ticker.L_idx = j + 1
		listIDX = append(listIDX, Ticker)
		j++
	}
	return listIDX
}
