package hotList

import (
	"CoinMarket.net/okxInfo"
	"github.com/EasyGolang/goTools/mCount"
)

// 成交额排序
func SortAmount(data []okxInfo.HotInfo) []okxInfo.HotInfo {
	size := len(data)
	list := make([]okxInfo.HotInfo, size)
	copy(list, data)

	reList := BubbleAmount(list)

	return Reverse(reList)
}

// 涨跌幅排序
func SortU_R(data []okxInfo.HotInfo) []okxInfo.HotInfo {
	size := len(data)
	list := make([]okxInfo.HotInfo, size)
	copy(list, data)

	reList := BubbleU_R24(list, "")

	return Reverse(reList)
}

// 涨跌幅绝对值排序
func SortU_RAbs(data []okxInfo.HotInfo) []okxInfo.HotInfo {
	size := len(data)
	list := make([]okxInfo.HotInfo, size)
	copy(list, data)

	reList := BubbleU_R24(list, "Abs")

	return Reverse(reList)
}

// 匹配现货数据
func FindSPOTInst(data okxInfo.HotInfo) okxInfo.InstInfoType {
	SPOT := okxInfo.InstInfoType{}
	for _, val := range okxInfo.InstList.SPOT {
		if val.BaseCcy == data.CcyName {
			SPOT = val
			break
		}
	}
	return SPOT
}

// 匹配合约数据
func FindSWAPInst(data okxInfo.HotInfo) okxInfo.InstInfoType {
	SWAP := okxInfo.InstInfoType{}
	for _, val := range okxInfo.InstList.SWAP {
		if val.CtValCcy == data.CcyName {
			SWAP = val
			break
		}
	}
	return SWAP
}

// 成交量排序
func BubbleAmount(arr []okxInfo.HotInfo) []okxInfo.HotInfo {
	size := len(arr)
	list := make([]okxInfo.HotInfo, size)
	copy(list, arr)

	var swapped bool
	for i := size - 1; i > 0; i-- {
		swapped = false
		for j := 0; j < i; j++ {
			a := list[j+1].Amount
			b := list[j].Amount
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

// 涨跌幅排序
func BubbleU_R24(arr []okxInfo.HotInfo, lType string) []okxInfo.HotInfo {
	size := len(arr)
	list := make([]okxInfo.HotInfo, size)
	copy(list, arr)

	var swapped bool
	for i := size - 1; i > 0; i-- {
		swapped = false
		for j := 0; j < i; j++ {
			a := list[j+1].U_R24
			b := list[j].U_R24
			if lType == "Abs" {
				a = mCount.Abs(list[j+1].U_R24)
				b = mCount.Abs(list[j].U_R24)
			}
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

func Reverse(arr []okxInfo.HotInfo) []okxInfo.HotInfo {
	list := make(
		[]okxInfo.HotInfo,
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
