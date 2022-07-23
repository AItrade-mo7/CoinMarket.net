package tickers

import (
	"fmt"
	"strings"

	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/okxApi/okxInfo"
	"github.com/EasyGolang/goTools/mCount"
)

func TickerCount(data okxInfo.TickerType, BinanceTicker okxInfo.BinanceTickerType) (Ticker okxInfo.TickerType) {
	Ticker = data

	fmt.Printf("Binance:%-15s  OKX:%-15s \n", BinanceTicker.InstID, Ticker.InstID)

	// 24 小时 涨幅
	Ticker.U_R24 = mCount.RoseCent(Ticker.Last, Ticker.Open24H)

	// 币种的名称
	Ticker.CcyName = strings.Replace(Ticker.InstID, config.SPOT_suffix, "", -1)

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
