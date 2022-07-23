package tickers

import (
	"strings"

	"CoinMarket.net/server/okxApi/okxInfo"
	"github.com/EasyGolang/goTools/mCount"
)

func TickerCount(data okxInfo.TickerType) (Ticker okxInfo.TickerType) {
	Ticker = data

	// 24 小时均价
	Avg24 := mCount.Average([]string{
		Ticker.Open24H,
		Ticker.High24H,
		Ticker.Low24H,
	})
	Ticker.Avg24 = mCount.PriceCent(Avg24, Ticker.Last)

	// 24 小时成交额
	Amount := mCount.Mul(
		Ticker.Avg24,
		Ticker.VolCcy24H,
	)
	Ticker.Amount = mCount.PriceCent(Amount, Ticker.Last)

	// 24 小时 涨幅
	Ticker.U_R24 = mCount.RoseCent(Ticker.Last, Ticker.Open24H)

	// 币种的名称
	Ticker.CcyName = strings.Replace(Ticker.InstID, "-USDT", "", -1)

	return Ticker
}

// 成交量排序
func BubbleAmount(arr []okxInfo.TickerType) []okxInfo.TickerType {
	size := len(arr)
	list := make([]okxInfo.TickerType, size)
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
func BubbleU_R24(arr []okxInfo.TickerType) []okxInfo.TickerType {
	size := len(arr)
	list := make([]okxInfo.TickerType, size)
	copy(list, arr)

	var swapped bool
	for i := size - 1; i > 0; i-- {
		swapped = false
		for j := 0; j < i; j++ {

			a := list[j+1].U_R24
			b := list[j].U_R24
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

func Reverse(arr []okxInfo.TickerType) []okxInfo.TickerType {
	list := make(
		[]okxInfo.TickerType,
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
