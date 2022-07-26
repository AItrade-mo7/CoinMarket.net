package ready

import (
	"fmt"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/okxApi/binanceApi"
	"CoinMarket.net/server/okxApi/okxInfo"
	"CoinMarket.net/server/okxApi/restApi/tickers"
	"github.com/EasyGolang/goTools/mCount"
)

func SetTicker() {
	if len(binanceApi.BinanceTickerList) != 10 || len(tickers.OKXTickerList) != 10 {
		global.InstLog.Println("数据条目不正确", len(binanceApi.BinanceTickerList), len(tickers.OKXTickerList))
	}

	for i := 0; i < 10; i++ {
		fmt.Println(binanceApi.BinanceTickerList[i].InstID, tickers.OKXTickerList[i].InstID)
	}
}

func TickerCount(OKXTicker tickers.OKXTickerType, BinanceTicker binanceApi.BinanceTickerType) (Ticker okxInfo.TickerType) {
	Ticker = okxInfo.TickerType{}

	return Ticker
}

// 成交量排序
func BubbleVolume(arr []okxInfo.TickerType) []okxInfo.TickerType {
	size := len(arr)
	list := make([]okxInfo.TickerType, size)
	copy(list, arr)

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
