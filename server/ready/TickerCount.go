package ready

import (
	"strings"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/okxApi/okxInfo"
	"github.com/EasyGolang/goTools/mCount"
)

func SetTicker() {
	if len(okxInfo.BinanceTickerList) != 10 || len(okxInfo.OKXTickerList) != 10 {
		global.InstLog.Println("数据条目不正确", len(okxInfo.BinanceTickerList), len(okxInfo.OKXTickerList))
	}

	for _, okx := range okxInfo.OKXTickerList {
		for _, binance := range okxInfo.BinanceTickerList {
			if okx.InstID == binance.InstID {
				TickerCount(okx, binance)
				break
			}
		}
	}
}

func TickerCount(OKXTicker okxInfo.OKXTickerType, BinanceTicker okxInfo.BinanceTickerType) (Ticker okxInfo.TickerType) {
	Ticker = okxInfo.TickerType{}
	Ticker.InstID = OKXTicker.InstID
	Ticker.CcyName = strings.Replace(Ticker.InstID, config.SPOT_suffix, "", -1)
	Ticker.Last = OKXTicker.Last
	Ticker.Open24H = OKXTicker.Open24H
	Ticker.High24H = OKXTicker.High24H
	Ticker.Low24H = OKXTicker.Low24H
	Ticker.OKXVol24H = OKXTicker.VolCcy24H
	Ticker.BinanceVol24H = BinanceTicker.QuoteVolume
	Ticker.U_R24 = mCount.RoseCent(OKXTicker.Last, OKXTicker.Open24H)
	Ticker.Volume = mCount.Add(OKXTicker.VolCcy24H, BinanceTicker.QuoteVolume)
	Ticker.OkxVolRose = mCount.PerCent(Ticker.OKXVol24H, Ticker.Volume)
	Ticker.BinanceVolRose = mCount.PerCent(Ticker.BinanceVol24H, Ticker.Volume)

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
