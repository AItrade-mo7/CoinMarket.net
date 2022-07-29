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
		global.InstLog.Println("TickerList 数据条目不正确", len(okxInfo.BinanceTickerList), len(okxInfo.OKXTickerList))
	}

	tickerList := []okxInfo.TickerType{}

	for _, okx := range okxInfo.OKXTickerList {
		for _, binance := range okxInfo.BinanceTickerList {
			if okx.InstID == binance.InstID {
				ticker := TickerCount(okx, binance)
				tickerList = append(tickerList, ticker)
				break
			}
		}
	}

	VolumeSortList := VolumeSort(tickerList)
	okxInfo.TickerList = VolumeSortList
	okxInfo.TickerU_R24 = U_R24Sort(VolumeSortList)
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
	Ticker.Ts = OKXTicker.Ts

	return Ticker
}

// 成交量排序
func VolumeSort(data []okxInfo.TickerType) []okxInfo.TickerType {
	size := len(data)
	list := make([]okxInfo.TickerType, size)
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

	listIDX := []okxInfo.TickerType{}
	j := 0
	for i := len(list) - 1; i > -1; i-- {
		Ticker := list[i]
		Ticker.VolIdx = j + 1
		listIDX = append(listIDX, Ticker)
		j++
	}

	return listIDX
}

// 涨跌幅排序
func U_R24Sort(data []okxInfo.TickerType) []okxInfo.TickerType {
	size := len(data)
	list := make([]okxInfo.TickerType, size)
	copy(list, data)

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

	// 设置 U_RIdx 并翻转
	listIDX := []okxInfo.TickerType{}
	j := 0
	for i := len(list) - 1; i > -1; i-- {
		Ticker := list[i]
		Ticker.U_RIdx = j + 1
		listIDX = append(listIDX, Ticker)
		j++
	}
	return listIDX
}
